package node

import (
	"context"
	"encoding/json"
	"fmt"
	jsonpatch "github.com/evanphx/json-patch"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/gorilla/mux"

	apiv1 "github.com/kubermatic/kubermatic/api/pkg/api/v1"
	kubermaticapiv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	"github.com/kubermatic/kubermatic/api/pkg/handler/middleware"
	"github.com/kubermatic/kubermatic/api/pkg/handler/v1/common"
	"github.com/kubermatic/kubermatic/api/pkg/provider"
)

// ndrReq defines HTTP request for getAddon and deleteAddon
// swagger:parameters getAddon deleteAddon
type ndrReq struct {
	common.GetClusterReq
	// in: path
	NdrName string `json:"ndrequest_id"`
}

// listReq defines HTTP request for listAddons endpoint
// swagger:parameters listAddons
type listReq struct {
	common.GetClusterReq
}

// createReq defines HTTP request for createAddon endpoint
// swagger:parameters createAddon
type createReq struct {
	common.GetClusterReq
	// in: body
	Body apiv1.NodeDeploymentRequest
}

// patchReq defines HTTP request for patchAddon endpoint
// swagger:parameters patchAddon
type patchReq struct {
	ndrReq
	// in: body
	Patch []byte
}

func DecodeGetNodeDeploymentRequest(c context.Context, r *http.Request) (interface{}, error) {
	var req ndrReq

	cr, err := common.DecodeGetClusterReq(c, r)
	if err != nil {
		return nil, err
	}

	ndrName, err := decodeNdrName(c, r)
	if err != nil {
		return nil, err
	}

	req.GetClusterReq = cr.(common.GetClusterReq)
	req.NdrName = ndrName

	return req, nil
}

func DecodeListNodeDeploymentRequests(c context.Context, r *http.Request) (interface{}, error) {
	var req listReq

	cr, err := common.DecodeGetClusterReq(c, r)
	if err != nil {
		return nil, err
	}

	req.GetClusterReq = cr.(common.GetClusterReq)

	return req, nil
}

func DecodeCreateNodeDeploymentRequest(c context.Context, r *http.Request) (interface{}, error) {
	var req createReq

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

func DecodePatchNodeDeploymentRequest(c context.Context, r *http.Request) (interface{}, error) {
	var req patchReq

	gr, err := DecodeGetNodeDeploymentRequest(c, r)
	if err != nil {
		return nil, err
	}

	req.ndrReq = gr.(ndrReq)

	if req.Patch, err = ioutil.ReadAll(r.Body); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeNdrName(c context.Context, r *http.Request) (string, error) {
	ndrName := mux.Vars(r)["ndrequest_id"]
	if ndrName == "" {
		return "", fmt.Errorf("'ndrequest_id' parameter is required but was not provided")
	}

	return ndrName, nil
}

func GetNodeDeploymentRequestEndpoint(projectProvider provider.ProjectProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ndrReq)
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

		mdrProvider := ctx.Value(middleware.MachineDeploymentRequestProviderContextKey).(provider.MachineDeploymentRequestProvider)
		mdr, err := mdrProvider.Get(userInfo, cluster, req.NdrName)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		result, err := convertMdrToNdr(mdr)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}
		return result, nil
	}
}

func ListNodeDeploymentRequestEndpoint(projectProvider provider.ProjectProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(listReq)
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

		mdrProvider := ctx.Value(middleware.MachineDeploymentRequestProviderContextKey).(provider.MachineDeploymentRequestProvider)
		mdrs, err := mdrProvider.List(userInfo, cluster)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		result, err := convertMdrToNdrList(mdrs)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}
		return result, nil
	}
}

func CreateNodeDeploymentRequestEndpoint(projectProvider provider.ProjectProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createReq)
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

		mdr, err := convertNdrToMdr(&req.Body)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		mdrProvider := ctx.Value(middleware.MachineDeploymentRequestProviderContextKey).(provider.MachineDeploymentRequestProvider)
		mdr, err = mdrProvider.New(userInfo, cluster, mdr.Name, &mdr.Spec.MdSpec)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		result, err := convertMdrToNdr(mdr)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}
		return result, nil
	}
}

func PatchNodeDeploymentRequestEndpoint(projectProvider provider.ProjectProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(patchReq)
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

		mdrProvider := ctx.Value(middleware.MachineDeploymentRequestProviderContextKey).(provider.MachineDeploymentRequestProvider)
		mdr, err := mdrProvider.Get(userInfo, cluster, req.NdrName)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		ndr, err := convertMdrToNdr(mdr)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		ndrJSON, err := json.Marshal(ndr)
		if err != nil {
			return nil, fmt.Errorf("cannot decode existing node deployment request: %v", err)
		}

		patchedNdrJSON, err := jsonpatch.MergePatch(ndrJSON, req.Patch)
		if err != nil {
			return nil, fmt.Errorf("cannot patch node deployment request: %v", err)
		}

		var patchedNdr *apiv1.NodeDeploymentRequest
		if err := json.Unmarshal(patchedNdrJSON, &patchedNdr); err != nil {
			return nil, fmt.Errorf("cannot decode patched cluster: %v", err)
		}

		// TODO validate patchedNdr here; see node.go / PatchNodeDeployment

		patchedMdr, err := convertNdrToMdr(patchedNdr)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		_, err = mdrProvider.Update(userInfo, cluster, patchedMdr)
		if err != nil {
			return nil, common.KubernetesErrorToHTTPError(err)
		}

		return patchedNdr, nil
	}
}

func DeleteNodeDeploymentRequestEndpoint(projectProvider provider.ProjectProvider) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ndrReq)
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

		mdrProvider := ctx.Value(middleware.MachineDeploymentRequestProviderContextKey).(provider.MachineDeploymentRequestProvider)

		return nil, common.KubernetesErrorToHTTPError(mdrProvider.Delete(userInfo, cluster, req.NdrName))
	}
}

func convertMdrToNdr(mdr *kubermaticapiv1.MachineDeploymentRequest) (*apiv1.NodeDeploymentRequest, error) {
	ndSpec, err := mdSpecToNdSpec(&mdr.Spec.MdSpec)
	if err != nil {
		return nil, err
	}

	result := &apiv1.NodeDeploymentRequest{
		ObjectMeta: apiv1.ObjectMeta{
			ID:                mdr.Name,
			Name:              mdr.Name,
			CreationTimestamp: apiv1.NewTime(mdr.CreationTimestamp.Time),
			DeletionTimestamp: func() *apiv1.Time {
				if mdr.DeletionTimestamp != nil {
					deletionTimestamp := apiv1.NewTime(mdr.DeletionTimestamp.Time)
					return &deletionTimestamp
				}
				return nil
			}(),
		},
		Spec: apiv1.NodeDeploymentRequestSpec{
			NdSpec: *ndSpec,
		},
	}
	return result, nil
}

func convertMdrToNdrList(mdrs []*kubermaticapiv1.MachineDeploymentRequest) ([]*apiv1.NodeDeploymentRequest, error) {
	result := []*apiv1.NodeDeploymentRequest{}

	for _, mdr := range mdrs {
		converted, err := convertMdrToNdr(mdr)
		if err != nil {
			return nil, err
		}
		result = append(result, converted)
	}

	return result, nil
}

func convertNdrToMdr(mdr *apiv1.NodeDeploymentRequest) (*kubermaticapiv1.MachineDeploymentRequest, error) {
	// TODO impl
	return nil, nil
}
