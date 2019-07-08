package cluster

import (
	"context"
	"fmt"
	"github.com/kubermatic/kubermatic/api/pkg/resources"

	"github.com/go-kit/kit/endpoint"
	"github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	"github.com/kubermatic/kubermatic/api/pkg/handler/middleware"
	"github.com/kubermatic/kubermatic/api/pkg/handler/v1/common"
	"github.com/kubermatic/kubermatic/api/pkg/provider"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func GetOidcKubeconfigEndpoint(projectProvider provider.ProjectProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(common.GetClusterReq)
		clusterProvider := ctx.Value(middleware.ClusterProviderContextKey).(provider.ClusterProvider)
		userInfo := ctx.Value(middleware.UserInfoContextKey).(*provider.UserInfo)
		_, err := projectProvider.Get(userInfo, req.ProjectID, &provider.ProjectGetOptions{})
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		cluster, err := clusterProvider.Get(userInfo, req.ClusterID, &provider.ClusterGetOptions{})
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}
		adminClientCfg, err := clusterProvider.GetAdminKubeconfigForCustomerCluster(cluster)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		clientCmdAuth := clientcmdapi.NewAuthInfo()
		clientCmdAuthProvider := &clientcmdapi.AuthProviderConfig{Config: map[string]string{}}
		clientCmdAuthProvider.Name = "oidc"
		clientCmdAuthProvider.Config["idp-issuer-url"] = cluster.Spec.OIDC.IssuerURL
		clientCmdAuthProvider.Config["client-id"] = cluster.Spec.OIDC.ClientID
		if cluster.Spec.OIDC.ClientSecret != "" {
			clientCmdAuthProvider.Config["client-secret"] = cluster.Spec.OIDC.ClientSecret
		}
		if cluster.Spec.OIDC.ExtraScopes != "" {
			clientCmdAuthProvider.Config["extra-scopes"] = cluster.Spec.OIDC.ExtraScopes
		}
		clientCmdAuth.AuthProvider = clientCmdAuthProvider

		adminClientCfg.AuthInfos = map[string]*clientcmdapi.AuthInfo{}
		adminClientCfg.AuthInfos["default"] = clientCmdAuth

		return &encodeKubeConifgResponse{clientCfg: adminClientCfg, filePrefix: "oidc"}, nil
	}
}

// This function will replace the 'default' context and user in a kubeconfig with sanitized names. Error will for now
// always be nil, but added the return in case we want to add checks e.g. for cluster or project in the future
func SanitizeKubeconfigContext(clientConfig *clientcmdapi.Config, cluster *v1.Cluster, project *v1.Project) (*clientcmdapi.Config, error) {
	sanitizedUser := fmt.Sprintf("admin-%s", cluster.Name)
	sanitizedContext := fmt.Sprintf("admin@%s/%s", project.Spec.Name, cluster.Spec.HumanReadableName)

	clientConfig.AuthInfos[sanitizedUser] = clientConfig.AuthInfos[resources.KubeconfigDefaultContextKey]
	clientConfig.AuthInfos[resources.KubeconfigDefaultContextKey] = nil

	clientConfig.Contexts[sanitizedContext] = clientConfig.Contexts[resources.KubeconfigDefaultContextKey]
	clientConfig.Contexts[sanitizedContext].AuthInfo = sanitizedUser
	clientConfig.Contexts[resources.KubeconfigDefaultContextKey] = nil

	clientConfig.CurrentContext = sanitizedContext

	return clientConfig, nil
}
