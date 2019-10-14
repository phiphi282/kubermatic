package cluster

import (
	"context"
	"fmt"
	"strings"

	"github.com/kubermatic/kubermatic/api/pkg/keycloak"

	v1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
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

		oidcSettings, err := getOidcSettings(ctx, cluster)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		clientCmdAuth := clientcmdapi.NewAuthInfo()
		clientCmdAuthProvider := &clientcmdapi.AuthProviderConfig{Config: map[string]string{}}
		clientCmdAuthProvider.Name = "oidc"
		clientCmdAuthProvider.Config["idp-issuer-url"] = oidcSettings.IssuerURL
		clientCmdAuthProvider.Config["client-id"] = oidcSettings.ClientID
		if oidcSettings.ClientSecret != "" {
			clientCmdAuthProvider.Config["client-secret"] = oidcSettings.ClientSecret
		}
		if oidcSettings.ExtraScopes != "" {
			clientCmdAuthProvider.Config["extra-scopes"] = oidcSettings.ExtraScopes
		}
		clientCmdAuth.AuthProvider = clientCmdAuthProvider

		oidcClientCfg.AuthInfos = map[string]*clientcmdapi.AuthInfo{}
		oidcClientCfg.AuthInfos[resources.KubeconfigDefaultContextKey] = clientCmdAuth

		if req.UseUniqueNames {
			return getKubeconfigWithUniqueName(project, cluster, oidcClientCfg)
		}
		return &encodeKubeConifgResponse{clientCfg: oidcClientCfg, filePrefix: "oidc"}, nil
	}
}

func GetKubeLoginKubeconfigEndpoint(projectProvider provider.ProjectProvider) endpoint.Endpoint {
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
		kubeloginClientCfg, err := clusterProvider.GetAdminKubeconfigForCustomerCluster(cluster)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		oidcSettings, err := getOidcSettings(ctx, cluster)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		clientCmdAuth := clientcmdapi.NewAuthInfo()
		execConfig := &clientcmdapi.ExecConfig{
			APIVersion: "client.authentication.k8s.io/v1beta1",
			Command:    "kubectl",
			Args: []string{
				"oidc-login",
				"get-token",
				fmt.Sprintf("--oidc-issuer-url=%s", oidcSettings.IssuerURL),
				fmt.Sprintf("--oidc-client-id=%s", oidcSettings.ClientID),
			},
		}
		if oidcSettings.ClientSecret != "" {
			execConfig.Args = append(execConfig.Args, fmt.Sprintf("--oidc-client-secret=%s", oidcSettings.ClientSecret))
		}
		if oidcSettings.ExtraScopes != "" {
			for _, scope := range strings.Split(oidcSettings.ExtraScopes, " ") {
				execConfig.Args = append(execConfig.Args, fmt.Sprintf("--oidc-extra-scope=%s", scope))
			}
		}
		clientCmdAuth.Exec = execConfig

		kubeloginClientCfg.AuthInfos = map[string]*clientcmdapi.AuthInfo{}
		kubeloginClientCfg.AuthInfos[resources.KubeconfigDefaultContextKey] = clientCmdAuth

		if req.UseUniqueNames {
			return getKubeconfigWithUniqueName(project, cluster, kubeloginClientCfg)
		}
		return &encodeKubeConifgResponse{clientCfg: kubeloginClientCfg, filePrefix: "kubelogin"}, nil
	}
}

func getKubeconfigWithUniqueName(project *v1.Project, cluster *v1.Cluster, clientConfig *clientcmdapi.Config) (interface{}, error) {
	kubeloginClientCfg, err := NoDefaultsKubeconfig(clientConfig, "oidc", cluster, project)
	if err != nil {
		return nil, kcerrors.NewBadRequest("failed to replace default names in kubelogin kubeconfig: %v", err)
	}
	return &encodeKubeConifgResponse{
		clientCfg:   kubeloginClientCfg,
		filePrefix:  "kubelogin",
		clusterName: fmt.Sprintf("%s-%s", project.Spec.Name, cluster.Spec.HumanReadableName),
	}, nil
}

func getOidcSettings(ctx context.Context, cluster *v1.Cluster) (*v1.OIDCSettings, error) {
	keycloakFacade := ctx.Value(middleware.KeycloakFacadeContextKey).(keycloak.Facade)

	if cluster.Spec.Sys11Auth.Realm != "" {
		settings, err := keycloak.Sys11AuthToOidcSettings(cluster.Spec.Sys11Auth, keycloakFacade)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}
		return settings, nil
	} else if cluster.Spec.OIDC.IssuerURL != "" && cluster.Spec.OIDC.ClientID != "" {
		return &cluster.Spec.OIDC, nil
	} else {
		return nil, kcerrors.NewNotFound("OIDC settings for", cluster.Name)
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
