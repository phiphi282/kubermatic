package handler

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"

	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/kubermatic/kubermatic/api/pkg/handler/middleware"
	"github.com/kubermatic/kubermatic/api/pkg/handler/v1/common"
	"github.com/kubermatic/kubermatic/api/pkg/provider"
	k8cerrors "github.com/kubermatic/kubermatic/api/pkg/util/errors"

)

var resourceUsageCollectorAuthToken = os.Getenv("RESOURCE_USAGE_COLLECTOR_AUTH_TOKEN")

// swagger:route GET /projects/{project_id}/dc/{dc}/cluster/{cluster_id}/v1/usage
//
// Gets resource usage data for the cluster
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: Quotas
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
	RequestHeaders http.Header
	ResourceUsageCollectorQueryPath string              `json:"resource_usage_collector_query_path"`
	ResourceUsageCollectorQuery     map[string][]string `json:"resource_usage_collector_raw_query"`
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
	req.ResourceUsageCollectorQueryPath = mux.Vars(r)["usage"]
	if req.ResourceUsageCollectorQueryPath != "usage" {
		return nil, k8cerrors.New(http.StatusBadRequest, "invalid Prometheus query path")
	}

	req.ResourceUsageCollectorQuery = r.URL.Query()
	req.RequestHeaders = r.Header

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

		resourceUsageCollectorQuery := map[string]string{}

		// Fixed Query parameters can be generated from the given url
		resourceUsageCollectorQuery["seed_cluster_name"] = seed.Name
		resourceUsageCollectorQuery["cluster_id"] = req.ClusterID
		resourceUsageCollectorQuery["project_id"] = req.ProjectID

		// Add additional query parameters
		for k, v := range req.ResourceUsageCollectorQuery {
			resourceUsageCollectorQuery[k] = v[0]
		}

		proxyRequest := masterKubeClient.CoreV1().Services("resource-usage-collector").ProxyGet(
			"http",
			"resource-usage-collector",
			"http",
			"/v1/"+req.ResourceUsageCollectorQueryPath,
			resourceUsageCollectorQuery,
		).(*rest.Request)

		proxyRequest.SetHeader("x-auth", resourceUsageCollectorAuthToken)
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
