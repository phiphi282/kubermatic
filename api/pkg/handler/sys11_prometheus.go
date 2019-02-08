package handler

import (
	"context"
	"fmt"
	goerrors "errors"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/kubermatic/kubermatic/api/pkg/handler/middleware"
	"github.com/kubermatic/kubermatic/api/pkg/handler/v1/common"
	"github.com/kubermatic/kubermatic/api/pkg/provider"
	k8cerrors "github.com/kubermatic/kubermatic/api/pkg/util/errors"
	prometheusapi "github.com/prometheus/client_golang/api"
	"time"

	"github.com/gorilla/mux"
	"net/http"
)

// RawResponse is the default representation of a raw (proxied) HTTP response
type RawResponse struct {
	Header http.Header
	Body   []byte
}

func (r Routing) prometheusProxyHandler() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			r.oidcAuthenticator.Verifier(),
			r.userSaverMiddleware(),
			middleware.Datacenter(r.clusterProviders, r.datacenters),
			r.userInfoMiddleware(),
		)(getPrometheusProxyEndpoint()),
		decodePrometheusProxyReq,
		encodeRawResponse,
		r.defaultServerOptions()...,
	)
}

// GetPrometheusProxyReq represents a request to the Prometheus proxy route
type GetPrometheusProxyReq struct {
	common.GetClusterReq
	PrometheusQueryPath string `json:"prometheus_query_path"`
	PrometheusRawQuery  string `json:"prometheus_raw_query"`
	RequestHeaders      http.Header
}

func decodePrometheusProxyReq(c context.Context, r *http.Request) (interface{}, error) {
	var req GetPrometheusProxyReq

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
	req.PrometheusQueryPath = mux.Vars(r)["query_path"]
	if req.PrometheusQueryPath != "query" && req.PrometheusQueryPath != "query_range" {
		return nil, k8cerrors.New(http.StatusBadRequest, "invalid Prometheus query path")
	}

	req.PrometheusRawQuery = r.URL.RawQuery
	req.RequestHeaders = r.Header

	return req, nil
}


func getPrometheusProxyEndpoint() endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (interface{}, error) {
		userInfo := ctx.Value(middleware.UserInfoContextKey).(*provider.UserInfo)
		clusterProvider := ctx.Value(middleware.ClusterProviderContextKey).(provider.ClusterProvider)
		req := request.(GetPrometheusProxyReq)
		c, err := clusterProvider.Get(userInfo, req.ClusterID, &provider.ClusterGetOptions{})
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()

		var prometheusClient prometheusapi.Client
		if prometheusClient, err = prometheusapi.NewClient(prometheusapi.Config{
			Address: fmt.Sprintf(`http://prometheus.cluster-%s.svc.cluster.local:9090`, c.Name),
		}); err != nil {
			return nil, err
		}

		promURL := prometheusClient.URL("/api/v1/"+req.PrometheusQueryPath, map[string]string{})
		promURL.RawQuery = req.PrometheusRawQuery

		promReq, err := http.NewRequest(http.MethodGet, promURL.String(), nil)
		if err != nil {
			return nil, err
		}

		promResp, body, err := prometheusClient.Do(ctx, promReq)
		if err != nil {
			return nil, err
		}

		resp := RawResponse{
			Header: http.Header{
				"Content-Type": []string{promResp.Header.Get("Content-Type")},
			},
			Body: body,
		}

		return resp, nil
	}
}

func encodeRawResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	if resp, ok := response.(RawResponse); ok {
		for field, values := range resp.Header {
			for _, value := range values {
				w.Header().Set(field, value)
			}
		}
		_, err := w.Write(resp.Body)
		return err
	}
	return goerrors.New("internal error (unexpected raw response object)")
}

