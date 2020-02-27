package handler

import (
	"github.com/kubermatic/kubermatic/api/pkg/handler/v1/addon"
	"github.com/kubermatic/kubermatic/api/pkg/handler/v1/node"
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

	mux.Methods(http.MethodGet).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/kubeloginkubeconfig").
		Handler(r.getKubeLoginClusterKubeconfig())

	mux.Methods(http.MethodPost).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons").
		Handler(r.createAddon())

	mux.Methods(http.MethodGet).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons").
		Handler(r.listAddons())

	mux.Methods(http.MethodGet).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}").
		Handler(r.getAddon())

	mux.Methods(http.MethodPatch).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}").
		Handler(r.patchAddon())

	mux.Methods(http.MethodDelete).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}").
		Handler(r.deleteAddon())

	mux.Methods(http.MethodPost).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests").
		Handler(r.createNodeDeploymentRequest())

	mux.Methods(http.MethodGet).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests").
		Handler(r.listNodeDeploymentRequests())

	mux.Methods(http.MethodGet).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests/{ndrequest_id}").
		Handler(r.getNodeDeploymentRequest())

	mux.Methods(http.MethodPatch).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests/{ndrequest_id}").
		Handler(r.patchNodeDeploymentRequest())

	mux.Methods(http.MethodDelete).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests/{ndrequest_id}").
		Handler(r.deleteNodeDeploymentRequest())

	// Metakube Vault Unsealing
	mux.Methods(http.MethodPost).
		Path("/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/metakube-vault/unseal").
		Handler(r.unsealVaultAddon())
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
			middleware.UserInfoExtractor(r.userProjectMapper),
		)(provider.OpenstackImageEndpoint(r.seedsGetter, r.presetsManager)),
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
			middleware.UserInfoExtractor(r.userProjectMapper),
		)(provider.OpenstackQuotaLimitEndpoint(r.seedsGetter, r.presetsManager)),
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
			middleware.SetClusterProvider(r.clusterProviderGetter, r.seedsGetter),
			middleware.UserInfoExtractor(r.userProjectMapper),
		)(provider.OpenstackImageNoCredentialsEndpoint(r.projectProvider, r.seedsGetter, r.presetsManager)),
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
			middleware.SetClusterProvider(r.clusterProviderGetter, r.seedsGetter),
			middleware.UserInfoExtractor(r.userProjectMapper),
		)(provider.OpenstackQuotaLimitNoCredentialsEndpoint(r.projectProvider, r.seedsGetter, r.presetsManager)),
		provider.DecodeOpenstackNoCredentialsReq,
		encodeJSON,
		r.defaultServerOptions()...,
	)
}

// swagger:route GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/kubeloginkubeconfig project getKubeLoginClusterKubeconfig
//
//     Gets the kubeconfig for the specified cluster with oidc authentication that works nicely with kube-login.
//
//     Produces:
//     - application/yaml
//
//     Responses:
//       default: errorResponse
//       200: Kubeconfig
//       401: empty
//       403: empty
func (r Routing) getKubeLoginClusterKubeconfig() http.Handler {
	privilegedUserGroups := map[string]bool{
		"owners":  true,
		"editors": true,
		"viewers": false,
	}
	return httptransport.NewServer(
		endpoint.Chain(
			middleware.TokenVerifier(r.tokenVerifiers),
			middleware.UserSaver(r.userProvider),
			middleware.SetClusterProvider(r.clusterProviderGetter, r.seedsGetter),
			middleware.UserInfoExtractor(r.userProjectMapper),
			middleware.Keycloak(r.keycloakFacade),
			middleware.PrivilegedUserGroupVerifier(r.userProjectMapper, privilegedUserGroups),
		)(cluster.GetKubeLoginKubeconfigEndpoint(r.projectProvider)),
		cluster.DecodeGetAdminKubeconfig,
		cluster.EncodeKubeconfig,
		r.defaultServerOptions()...,
	)
}

// swagger:route POST /projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/metakube-vault/unseal
//
//    Unseals the vault addon in the cluster.
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: empty
//       401: empty
//       403: empty
func (r Routing) unsealVaultAddon() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			middleware.TokenVerifier(r.tokenVerifiers),
			middleware.UserSaver(r.userProvider),
			middleware.SetClusterProvider(r.clusterProviderGetter, r.seedsGetter),
			middleware.Addons(r.addonProviderGetter, r.seedsGetter),
			middleware.UserInfoExtractor(r.userProjectMapper),
		)(addon.UnsealVaultAddonEndpoint(r.seedsGetter)),
		addon.DecodeUnsealVaultAddon,
		encodeJSON,
		r.defaultServerOptions()...,
	)
}

// swagger:route POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests ndrequest createNodeDeploymentRequest
//
//     Creates a NodeDeploymentRequest that will belong to the given cluster
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       201: NodeDeploymentRequest
//       401: empty
//       403: empty
func (r Routing) createNodeDeploymentRequest() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			middleware.TokenVerifier(r.tokenVerifiers),
			middleware.UserSaver(r.userProvider),
			middleware.SetClusterProvider(r.clusterProviderGetter, r.seedsGetter),
			middleware.MachineDeploymentRequests(r.mdRequestProviderGetter, r.seedsGetter),
			middleware.UserInfoExtractor(r.userProjectMapper),
		)(node.CreateNodeDeploymentRequestEndpoint(r.projectProvider)),
		node.DecodeCreateNodeDeploymentRequest,
		setStatusCreatedHeader(encodeJSON),
		r.defaultServerOptions()...,
	)
}

// swagger:route GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests ndrequest listNodeDeploymentRequests
//
//     Lists NodeDeploymentRequests that belong to the given cluster
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: []NodeDeploymentRequest
//       401: empty
//       403: empty
func (r Routing) listNodeDeploymentRequests() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			middleware.TokenVerifier(r.tokenVerifiers),
			middleware.UserSaver(r.userProvider),
			middleware.SetClusterProvider(r.clusterProviderGetter, r.seedsGetter),
			middleware.MachineDeploymentRequests(r.mdRequestProviderGetter, r.seedsGetter),
			middleware.UserInfoExtractor(r.userProjectMapper),
		)(node.ListNodeDeploymentRequestEndpoint(r.projectProvider)),
		node.DecodeListNodeDeploymentRequests,
		encodeJSON,
		r.defaultServerOptions()...,
	)
}

// swagger:route GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests/{ndrequest_id} ndrequest getNodeDeploymentRequest
//
//     Gets a NodeDeploymentRequest that is assigned to the given cluster.
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: NodeDeploymentRequest
//       401: empty
//       403: empty
func (r Routing) getNodeDeploymentRequest() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			middleware.TokenVerifier(r.tokenVerifiers),
			middleware.UserSaver(r.userProvider),
			middleware.SetClusterProvider(r.clusterProviderGetter, r.seedsGetter),
			middleware.MachineDeploymentRequests(r.mdRequestProviderGetter, r.seedsGetter),
			middleware.UserInfoExtractor(r.userProjectMapper),
		)(node.GetNodeDeploymentRequestEndpoint(r.projectProvider)),
		node.DecodeGetNodeDeploymentRequest,
		encodeJSON,
		r.defaultServerOptions()...,
	)
}

// swagger:route PATCH /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests/{ndrequest_id} ndrequest patchNodeDeploymentRequest
//
//     Patches a NodeDeploymentRequest that is assigned to the given cluster.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: NodeDeploymentRequest
//       401: empty
//       403: empty
func (r Routing) patchNodeDeploymentRequest() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			middleware.TokenVerifier(r.tokenVerifiers),
			middleware.UserSaver(r.userProvider),
			middleware.SetClusterProvider(r.clusterProviderGetter, r.seedsGetter),
			middleware.MachineDeploymentRequests(r.mdRequestProviderGetter, r.seedsGetter),
			middleware.UserInfoExtractor(r.userProjectMapper),
		)(node.PatchNodeDeploymentRequestEndpoint(r.projectProvider)),
		node.DecodePatchNodeDeploymentRequest,
		encodeJSON,
		r.defaultServerOptions()...,
	)
}

// swagger:route DELETE /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests/{ndrequest_id} ndrequest deleteNodeDeploymentRequest
//
//    Deletes the given NodeDeploymentRequest that belongs to the cluster.
//
//     Produces:
//     - application/json
//
//     Responses:
//       default: errorResponse
//       200: empty
//       401: empty
//       403: empty
func (r Routing) deleteNodeDeploymentRequest() http.Handler {
	return httptransport.NewServer(
		endpoint.Chain(
			middleware.TokenVerifier(r.tokenVerifiers),
			middleware.UserSaver(r.userProvider),
			middleware.SetClusterProvider(r.clusterProviderGetter, r.seedsGetter),
			middleware.MachineDeploymentRequests(r.mdRequestProviderGetter, r.seedsGetter),
			middleware.UserInfoExtractor(r.userProjectMapper),
		)(node.DeleteNodeDeploymentRequestEndpoint(r.projectProvider)),
		node.DecodeGetNodeDeploymentRequest,
		encodeJSON,
		r.defaultServerOptions()...,
	)
}
