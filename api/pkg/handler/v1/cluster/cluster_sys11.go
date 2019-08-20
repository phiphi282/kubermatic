package cluster

import (
	"context"
	"fmt"

	"github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	"github.com/kubermatic/kubermatic/api/pkg/handler/middleware"
	"github.com/kubermatic/kubermatic/api/pkg/handler/v1/common"
	"github.com/kubermatic/kubermatic/api/pkg/provider"
	"github.com/kubermatic/kubermatic/api/pkg/resources"
	kcerrors "github.com/kubermatic/kubermatic/api/pkg/util/errors"

	"github.com/go-kit/kit/endpoint"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func GetOidcKubeconfigEndpoint(projectProvider provider.ProjectProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetClusterKubeconfigRequest)
		clusterProvider := ctx.Value(middleware.ClusterProviderContextKey).(provider.ClusterProvider)
		userInfo := ctx.Value(middleware.UserInfoContextKey).(*provider.UserInfo)
		project, err := projectProvider.Get(userInfo, req.ProjectID, &provider.ProjectGetOptions{})
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		cluster, err := clusterProvider.Get(userInfo, req.ClusterID, &provider.ClusterGetOptions{})
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}
		oidcClientCfg, err := clusterProvider.GetAdminKubeconfigForCustomerCluster(cluster)
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

		oidcClientCfg.AuthInfos = map[string]*clientcmdapi.AuthInfo{}
		oidcClientCfg.AuthInfos[resources.KubeconfigDefaultContextKey] = clientCmdAuth

		if req.UseUniqueNames {
			oidcClientCfg, err = NoDefaultsKubeconfig(oidcClientCfg, "oidc", cluster, project)
			if err != nil {
				return nil, kcerrors.NewBadRequest("failed to replace default names in oidc kubeconfig: %v", err)
			}
			return &encodeKubeConifgResponse{
				clientCfg:   oidcClientCfg,
				filePrefix:  "oidc",
				clusterName: fmt.Sprintf("%s-%s", project.Spec.Name, cluster.Spec.HumanReadableName),
			}, nil
		}
		return &encodeKubeConifgResponse{clientCfg: oidcClientCfg, filePrefix: "oidc"}, nil
	}
}

type GetClusterKubeconfigRequest struct {
	common.GetClusterReq
	// in: path
	UseUniqueNames bool `json:"use_unique_names"`
}

// This function will replace the 'default' context and user in a kubeconfig with sanitized names. Error will for now
// always be nil, but added the return in case we want to add checks e.g. for cluster or project in the future
func NoDefaultsKubeconfig(clientConfig *clientcmdapi.Config, username string, cluster *v1.Cluster, project *v1.Project) (*clientcmdapi.Config, error) {

	newUsername := fmt.Sprintf("%s-%s", username, cluster.Name)
	newContext := fmt.Sprintf("%s@%s/%s", username, project.Spec.Name, cluster.Spec.HumanReadableName)

	clientConfig.AuthInfos[newUsername] = clientConfig.AuthInfos[resources.KubeconfigDefaultContextKey]
	delete(clientConfig.AuthInfos, resources.KubeconfigDefaultContextKey)

	clientConfig.Contexts[newContext] = clientConfig.Contexts[resources.KubeconfigDefaultContextKey]
	clientConfig.Contexts[newContext].AuthInfo = newUsername
	delete(clientConfig.Contexts, resources.KubeconfigDefaultContextKey)

	clientConfig.CurrentContext = newContext

	return clientConfig, nil
}
