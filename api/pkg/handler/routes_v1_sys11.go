package handler

import (
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/kubermatic/kubermatic/api/pkg/handler/middleware"
	"github.com/kubermatic/kubermatic/api/pkg/handler/v1/provider"

	"github.com/gorilla/mux"
)

func (r Routing) RegisterV1SysEleven(mux *mux.Router) {
	mux.Methods(http.MethodGet).
		Path("/projects/{project_id}/dc/{dc}/cluster/{cluster_id}/prometheus/{query_path}").
		Handler(r.prometheusProxyHandler())

	mux.Methods(http.MethodGet).
		Path("/providers/openstack/images").
		Handler(r.listOpenstackImages())

	mux.Methods(http.MethodGet).
		Path("/providers/openstack/quotalimits").
		Handler(r.listOpenstackQuotaLimits())

	mux.Methods(http.MethodGet).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/images").
		Handler(r.listOpenstackImagesNoCredentials())

	mux.Methods(http.MethodGet).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/quotalimits").
		Handler(r.listOpenstackQuotaLimitsNoCredentials())
}

// swagger:route GET /api/v1/providers/openstack/images openstack listOpenstackImages
//
// Lists images from openstack
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: []Image
func (r Routing) listOpenstackImages() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			r.oidcAuthenticator.Verifier(),
			middleware.UserSaver(r.userProvider),
		)(provider.OpenstackImageEndpoint(r.cloudProviders)),
		provider.DecodeOpenstackReq,
		EncodeJSON,
		r.defaultServerOptions()...,
	)
}

// swagger:route GET /api/v1/providers/openstack/quotalimits openstack listOpenstackQuotaLimits
//
// Lists quotalimits for tenant from openstack
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: Quotas
func (r Routing) listOpenstackQuotaLimits() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			r.oidcAuthenticator.Verifier(),
			middleware.UserSaver(r.userProvider),
		)(provider.OpenstackQuotaLimitEndpoint(r.cloudProviders)),
		provider.DecodeOpenstackReq,
		EncodeJSON,
		r.defaultServerOptions()...,
	)
}

// swagger:route GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/images openstack listOpenstackImagesNoCredentials
//
// Lists images from openstack
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: []Image
func (r Routing) listOpenstackImagesNoCredentials() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			r.oidcAuthenticator.Verifier(),
			middleware.UserSaver(r.userProvider),
			middleware.Datacenter(r.clusterProviders, r.datacenters),
			middleware.UserInfo(r.userProjectMapper),
		)(provider.OpenstackImageNoCredentialsEndpoint(r.projectProvider, r.cloudProviders)),
		provider.DecodeOpenstackNoCredentialsReq,
		EncodeJSON,
		r.defaultServerOptions()...,
	)
}

// swagger:route GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/openstack/quotalimits openstack listOpenstackQuotaLimitsNoCredentials
//
// Lists quotalimits for tenant from openstack
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: Quotas
func (r Routing) listOpenstackQuotaLimitsNoCredentials() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			r.oidcAuthenticator.Verifier(),
			middleware.UserSaver(r.userProvider),
			middleware.Datacenter(r.clusterProviders, r.datacenters),
			middleware.UserInfo(r.userProjectMapper),
		)(provider.OpenstackQuotaLimitNoCredentialsEndpoint(r.projectProvider, r.cloudProviders)),
		provider.DecodeOpenstackNoCredentialsReq,
		EncodeJSON,
		r.defaultServerOptions()...,
	)
}
