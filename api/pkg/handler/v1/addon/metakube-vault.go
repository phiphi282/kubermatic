package addon

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kubermatic/kubermatic/api/pkg/resources"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/kubermatic/kubermatic/api/pkg/handler/middleware"
	"github.com/kubermatic/kubermatic/api/pkg/handler/v1/common"
	"github.com/kubermatic/kubermatic/api/pkg/provider"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// unsealVaultReq defines HTTP request for unsealVaultAddon
// swagger:parameters unsealVaultAddon
type unsealVaultReq struct {
	common.GetClusterReq
	Body unsealKeys
}

type unsealKeys struct {
	Keys []string `json:"keys"`
}

type RawResponse struct {
	Header http.Header
	Body   []byte
}

func DecodeUnsealVaultAddon(c context.Context, r *http.Request) (interface{}, error) {
	var req unsealVaultReq

	cr, err := common.DecodeGetClusterReq(c, r)
	if err != nil {
		return nil, err
	}

	req.GetClusterReq = cr.(common.GetClusterReq)

	if err := json.NewDecoder(r.Body).Decode(&req.Body); err != nil {
		return nil, err
	}

	return req, nil
}

func UnsealVaultAddonEndpoint(datacenters map[string]provider.DatacenterMeta) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		userInfo := ctx.Value(middleware.UserInfoContextKey).(*provider.UserInfo)
		clusterProvider := ctx.Value(middleware.ClusterProviderContextKey).(provider.ClusterProvider)
		req := request.(unsealVaultReq)
		c, err := clusterProvider.Get(userInfo, req.ClusterID, &provider.ClusterGetOptions{})
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
		defer cancel()

		masterConfig, err := rest.InClusterConfig()
		if err != nil {
			//For running outside of a cluster
			masterConfig, err = clientcmd.
				NewNonInteractiveDeferredLoadingClientConfig(
					clientcmd.NewDefaultClientConfigLoadingRules(),
					&clientcmd.ConfigOverrides{},
				).ClientConfig()
			if err != nil {
				return nil, fmt.Errorf("failed to create in cluster client config: %v", err)
			}
		}

		masterKubeClient, err := kubernetes.NewForConfig(masterConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create a master kubernetes client: %v", err)
		}

		seedConfigSecret, err := masterKubeClient.CoreV1().Secrets("kubermatic").Get("kubeconfig", metaV1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to get kubeconfig secret: %v", err)
		}

		seedConfig, err := getKubeconfig(datacenters[c.Spec.Cloud.DatacenterName].Seed, seedConfigSecret.Data[resources.KubeconfigSecretKey])
		if err != nil {
			return nil, fmt.Errorf("failed to get seed kubeconfig: %v", err)
		}

		seedKubeClient, err := kubernetes.NewForConfig(seedConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create a seed kubernetes client: %v", err)
		}

		clusterKubeConfigSecret, err := seedKubeClient.CoreV1().Secrets(fmt.Sprintf("cluster-%s", c.Name)).Get("admin-kubeconfig", metaV1.GetOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed to get cluster kubeconfig secret: %v", err)
		}

		clusterKubeConfig, err := getKubeconfig(resources.KubeconfigDefaultContextKey, clusterKubeConfigSecret.Data[resources.KubeconfigSecretKey])
		if err != nil {
			return nil, fmt.Errorf("failed to get cluster kubeconfig: %v", err)
		}

		clusterKubeClient, err := kubernetes.NewForConfig(clusterKubeConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create a cluster kubernetes client: %v", err)
		}

		// find metakube-vault addon pods
		requirement, err := labels.NewRequirement("app.kubernetes.io/instance", selection.Equals, []string{"syseleven-vault"})
		if err != nil {
			return nil, fmt.Errorf("failed to build label selector: %v", err)
		}
		options := metaV1.ListOptions{
			LabelSelector: requirement.String(),
		}
		endpoints, err := clusterKubeClient.CoreV1().Pods("metakube-vault").List(options)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch vault pods: %v", err)
		}

		for _, pod := range endpoints.Items {
			for _, key := range req.Body.Keys {
				proxyRequest := clusterKubeClient.CoreV1().RESTClient().Put().
					Namespace(pod.Namespace).
					Resource("pods").
					Name(fmt.Sprintf("%s:8200", pod.Name)).
					SubResource("proxy").
					Suffix("sys/unseal").
					Body(fmt.Sprintf(`{ "key": "%s" }`, key))
				body, err := proxyRequest.DoRaw()
				if err != nil {
					return nil, err
				}
				// TODO why body null?
			}
		}

		resp := RawResponse{
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			Body: nil,
		}

		return resp, nil
	}
}

func getKubeconfig(contextName string, cfg []byte) (*rest.Config, error) {
	clusterClientCfg, err := clientcmd.Load(cfg)
	if err != nil {
		return nil, err
	}

	return clientcmd.NewNonInteractiveClientConfig(
		*clusterClientCfg,
		contextName,
		&clientcmd.ConfigOverrides{},
		nil,
	).ClientConfig()
}
