package handler

import (
	"context"
	"fmt"
	"github.com/kubermatic/kubermatic/api/pkg/resources"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/kubermatic/kubermatic/api/pkg/handler/middleware"
	"github.com/kubermatic/kubermatic/api/pkg/handler/v1/common"
	"github.com/kubermatic/kubermatic/api/pkg/provider"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r Routing) resourceUsageCollectorProxyHandler() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			middleware.TokenVerifier(r.tokenVerifiers),
			middleware.UserSaver(r.userProvider),
			middleware.SetClusterProvider(r.clusterProviderGetter, r.seedsGetter),
		)(getResourceUsageCollectorProxyEndpoint(r.seedsGetter, r.userInfoGetter)),
		decodeResourceUsageCollectorProxyReq,
		encodeRawResponse,
		r.defaultServerOptions()...,
	)
}

// GetResourceUsageCollectorProxyReq represents a request to the ResourceUsageCollector proxy route
type GetResourceUsageCollectorProxyReq struct {
	common.GetClusterReq
	ResourceUsageCollectorQueryPath string              `json:"prometheus_query_path"`
	ResourceUsageCollectorQuery     map[string][]string `json:"prometheus_raw_query"`
	RequestHeaders      http.Header
}

func decodeResourceUsageCollectorProxyReq(c context.Context, r *http.Request) (interface{}, error) {
	var req GetResourceUsageCollectorProxyReq

	clusterID, err := common.DecodeClusterID(c, r)
	if err != nil {
		return nil, err
	}

	dcr, err := common.DecodeDcReq(c, r)
	if err != nil {
		return nil, err
	}

	req.ClusterID = clusterID
	req.DCReq = dcr.(common.DCReq)

	return req, nil
}

func getResourceUsageCollectorProxyEndpoint(seedsGetter provider.SeedsGetter, userInfoGetter provider.UserInfoGetter) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		clusterProvider := ctx.Value(middleware.ClusterProviderContextKey).(provider.ClusterProvider)
		req := request.(GetResourceUsageCollectorProxyReq)
		userInfo, err := userInfoGetter(ctx, req.ProjectID)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}
		c, err := clusterProvider.Get(userInfo, req.ClusterID, &provider.ClusterGetOptions{})
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		seed, _, err := provider.DatacenterFromSeedMap(userInfo, seedsGetter, c.Spec.Cloud.DatacenterName)
		if err != nil {
			return nil, fmt.Errorf("error getting dc: %v", err)
		}

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

		seedConfig, err := getSeedKubeconfig(seed.Name, seedConfigSecret.Data[resources.KubeconfigSecretKey])
		if err != nil {
			return nil, fmt.Errorf("failed to get seed kubeconfig: %v", err)
		}

		seedKubeClient, err := kubernetes.NewForConfig(seedConfig)
		if err != nil {
			return nil, fmt.Errorf("failed to create a seed kubernetes client: %v", err)
		}

		resourceUsageCollectorQuery := map[string]string{}

		// prometheus does not support the same query parameter twice, so we can just use the first
		for k, v := range req.ResourceUsageCollectorQuery {
			resourceUsageCollectorQuery[k] = v[0]
		}

		proxyRequest := seedKubeClient.CoreV1().Services("resource-usage-collector").ProxyGet(
			"tcp",
			"resource-usage-collector",
			"http",
			"/v1/usage",
			resourceUsageCollectorQuery,
		).(*rest.Request)

		body, err := proxyRequest.DoRaw()

		if err != nil {
			return nil, err
		}

		resp := RawResponse{
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			Body: body,
		}

		return resp, nil
	}
}
