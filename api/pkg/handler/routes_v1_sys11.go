package handler

import (
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/kubermatic/kubermatic/api/pkg/handler/middleware"
	"github.com/kubermatic/kubermatic/api/pkg/handler/v1/cluster"
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

	mux.Methods(http.MethodGet).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/oidckubeconfig").
		Handler(r.getOidcClusterKubeconfig())
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
			middleware.TokenVerifier(r.tokenVerifiers),
			middleware.UserSaver(r.userProvider),
		)(provider.OpenstackImageEndpoint(r.cloudProviders)),
		provider.DecodeOpenstackReq,
		encodeJSON,
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
			middleware.TokenVerifier(r.tokenVerifiers),
			middleware.UserSaver(r.userProvider),
		)(provider.OpenstackQuotaLimitEndpoint(r.cloudProviders)),
		provider.DecodeOpenstackReq,
		encodeJSON,
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
			middleware.TokenVerifier(r.tokenVerifiers),
			middleware.UserSaver(r.userProvider),
			middleware.Datacenter(r.clusterProviders, r.datacenters),
			middleware.UserInfoExtractor(r.userProjectMapper),
		)(provider.OpenstackImageNoCredentialsEndpoint(r.projectProvider, r.cloudProviders)),
		provider.DecodeOpenstackNoCredentialsReq,
		encodeJSON,
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
			middleware.TokenVerifier(r.tokenVerifiers),
			middleware.UserSaver(r.userProvider),
			middleware.Datacenter(r.clusterProviders, r.datacenters),
			middleware.UserInfoExtractor(r.userProjectMapper),
		)(provider.OpenstackQuotaLimitNoCredentialsEndpoint(r.projectProvider, r.cloudProviders)),
		provider.DecodeOpenstackNoCredentialsReq,
		encodeJSON,
		r.defaultServerOptions()...,
	)
}

// swagger:route GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/oidckubeconfig project getOidcClusterKubeconfig
//
//     Gets the kubeconfig for the specified cluster with oidc authentication.
//
//     Produces:
//     - application/yaml
//
//     Responses:
//       default: errorResponse
//       200: Kubeconfig
//       401: empty
//       403: empty
func (r Routing) getOidcClusterKubeconfig() http.Handler {
	privilegedUserGroups := map[string]bool{
		"owners":  true,
		"editors": true,
		"viewers": false,
	}
	return httptransport.NewServer(
		endpoint.Chain(
			middleware.TokenVerifier(r.tokenVerifiers),
			middleware.UserSaver(r.userProvider),
			middleware.Datacenter(r.clusterProviders, r.datacenters),
			middleware.UserInfoExtractor(r.userProjectMapper),
			middleware.PrivilegedUserGroupVerifier(r.userProjectMapper, privilegedUserGroups),
		)(cluster.GetOidcKubeconfigEndpoint(r.projectProvider)),
		cluster.DecodeGetAdminKubeconfig,
		cluster.EncodeKubeconfig,
		r.defaultServerOptions()...,
	)
}
