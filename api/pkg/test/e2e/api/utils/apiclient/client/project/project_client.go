// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new project API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for project API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AssignSSHKeyToCluster Assigns an existing ssh key to the given cluster
*/
func (a *Client) AssignSSHKeyToCluster(params *AssignSSHKeyToClusterParams, authInfo runtime.ClientAuthInfoWriter) (*AssignSSHKeyToClusterCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssignSSHKeyToClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "assignSSHKeyToCluster",
		Method:             "PUT",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssignSSHKeyToClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AssignSSHKeyToClusterCreated), nil

}

/*
CreateCluster creates a cluster for the given project
*/
func (a *Client) CreateCluster(params *CreateClusterParams, authInfo runtime.ClientAuthInfoWriter) (*CreateClusterCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCluster",
		Method:             "POST",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClusterCreated), nil

}

/*
CreateNodeDeployment Creates a node deployment that will belong to the given cluster
*/
func (a *Client) CreateNodeDeployment(params *CreateNodeDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*CreateNodeDeploymentCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNodeDeploymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createNodeDeployment",
		Method:             "POST",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateNodeDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateNodeDeploymentCreated), nil

}

/*
CreateNodeForClusterLegacy deprecateds creates a node that will belong to the given cluster

This endpoint is deprecated, please create a Node Deployment instead.
Use POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments
*/
func (a *Client) CreateNodeForClusterLegacy(params *CreateNodeForClusterLegacyParams, authInfo runtime.ClientAuthInfoWriter) (*CreateNodeForClusterLegacyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNodeForClusterLegacyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createNodeForClusterLegacy",
		Method:             "POST",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateNodeForClusterLegacyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateNodeForClusterLegacyCreated), nil

}

/*
CreateProject creates a brand new project

Note that this endpoint can be consumed by every authenticated user.
*/
func (a *Client) CreateProject(params *CreateProjectParams, authInfo runtime.ClientAuthInfoWriter) (*CreateProjectCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createProject",
		Method:             "POST",
		PathPattern:        "/api/v1/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateProjectCreated), nil

}

/*
CreateSSHKey adds the given SSH key to the specified project
*/
func (a *Client) CreateSSHKey(params *CreateSSHKeyParams, authInfo runtime.ClientAuthInfoWriter) (*CreateSSHKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSSHKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSSHKey",
		Method:             "POST",
		PathPattern:        "/api/v1/projects/{project_id}/sshkeys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateSSHKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSSHKeyOK), nil

}

/*
DeleteAddon deletes the given addon that belongs to the cluster
*/
func (a *Client) DeleteAddon(params *DeleteAddonParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAddonOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAddonParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteAddon",
		Method:             "DELETE",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/addons/{addon_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAddonReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAddonOK), nil

}

/*
DeleteCluster Deletes the specified cluster
*/
func (a *Client) DeleteCluster(params *DeleteClusterParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCluster",
		Method:             "DELETE",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClusterOK), nil

}

/*
DeleteNodeDeployment deletes the given node deployment that belongs to the cluster
*/
func (a *Client) DeleteNodeDeployment(params *DeleteNodeDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNodeDeploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodeDeploymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNodeDeployment",
		Method:             "DELETE",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNodeDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNodeDeploymentOK), nil

}

/*
DeleteNodeForClusterLegacy deprecateds deletes the given node that belongs to the cluster

This endpoint is deprecated, please create a Node Deployment instead.
*/
func (a *Client) DeleteNodeForClusterLegacy(params *DeleteNodeForClusterLegacyParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNodeForClusterLegacyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodeForClusterLegacyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNodeForClusterLegacy",
		Method:             "DELETE",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNodeForClusterLegacyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNodeForClusterLegacyOK), nil

}

/*
DeleteProject deletes the project with the given ID
*/
func (a *Client) DeleteProject(params *DeleteProjectParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteProject",
		Method:             "DELETE",
		PathPattern:        "/api/v1/projects/{project_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteProjectOK), nil

}

/*
DeleteSSHKey removes the given SSH key from the system
*/
func (a *Client) DeleteSSHKey(params *DeleteSSHKeyParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteSSHKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSSHKeyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSSHKey",
		Method:             "DELETE",
		PathPattern:        "/api/v1/projects/{project_id}/sshkeys/{key_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteSSHKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSSHKeyOK), nil

}

/*
DetachSSHKeyFromCluster Unassignes an ssh key from the given cluster
*/
func (a *Client) DetachSSHKeyFromCluster(params *DetachSSHKeyFromClusterParams, authInfo runtime.ClientAuthInfoWriter) (*DetachSSHKeyFromClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetachSSHKeyFromClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "detachSSHKeyFromCluster",
		Method:             "DELETE",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys/{key_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DetachSSHKeyFromClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DetachSSHKeyFromClusterOK), nil

}

/*
GetCluster Gets the cluster with the given name
*/
func (a *Client) GetCluster(params *GetClusterParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCluster",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterOK), nil

}

/*
GetClusterEvents gets the events related to the specified cluster
*/
func (a *Client) GetClusterEvents(params *GetClusterEventsParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterEvents",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/events",
		ProducesMediaTypes: []string{"application/yaml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterEventsOK), nil

}

/*
GetClusterHealth Returns the cluster's component health status
*/
func (a *Client) GetClusterHealth(params *GetClusterHealthParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterHealthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterHealthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterHealth",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/health",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterHealthReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterHealthOK), nil

}

/*
GetClusterKubeconfig gets the kubeconfig for the specified cluster
*/
func (a *Client) GetClusterKubeconfig(params *GetClusterKubeconfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterKubeconfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterKubeconfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterKubeconfig",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/kubeconfig",
		ProducesMediaTypes: []string{"application/yaml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterKubeconfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterKubeconfigOK), nil

}

/*
GetClusterMetrics Gets cluster metrics
*/
func (a *Client) GetClusterMetrics(params *GetClusterMetricsParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterMetricsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterMetrics",
		Method:             "GET",
		PathPattern:        "/api/v1alpha/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/metrics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterMetricsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterMetricsOK), nil

}

/*
GetClusterUpgrades Gets possible cluster upgrades
*/
func (a *Client) GetClusterUpgrades(params *GetClusterUpgradesParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterUpgradesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterUpgradesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterUpgrades",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/upgrades",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterUpgradesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterUpgradesOK), nil

}

/*
GetNodeDeployment gets a node deployment that is assigned to the given cluster
*/
func (a *Client) GetNodeDeployment(params *GetNodeDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodeDeploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodeDeploymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNodeDeployment",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNodeDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodeDeploymentOK), nil

}

/*
GetNodeForClusterLegacy deprecateds gets a node that is assigned to the given cluster

This endpoint is deprecated, please create a Node Deployment instead.
*/
func (a *Client) GetNodeForClusterLegacy(params *GetNodeForClusterLegacyParams, authInfo runtime.ClientAuthInfoWriter) (*GetNodeForClusterLegacyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodeForClusterLegacyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNodeForClusterLegacy",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/{node_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNodeForClusterLegacyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodeForClusterLegacyOK), nil

}

/*
GetOidcClusterKubeconfig gets the kubeconfig for the specified cluster with oidc authentication
*/
func (a *Client) GetOidcClusterKubeconfig(params *GetOidcClusterKubeconfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetOidcClusterKubeconfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOidcClusterKubeconfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getOidcClusterKubeconfig",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/oidckubeconfig",
		ProducesMediaTypes: []string{"application/yaml"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOidcClusterKubeconfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetOidcClusterKubeconfigOK), nil

}

/*
GetProject Gets the project with the given ID
*/
func (a *Client) GetProject(params *GetProjectParams, authInfo runtime.ClientAuthInfoWriter) (*GetProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProject",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProjectOK), nil

}

/*
ListClusters lists clusters for the specified project and data center
*/
func (a *Client) ListClusters(params *ListClustersParams, authInfo runtime.ClientAuthInfoWriter) (*ListClustersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListClustersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listClusters",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListClustersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListClustersOK), nil

}

/*
ListClustersForProject lists clusters for the specified project
*/
func (a *Client) ListClustersForProject(params *ListClustersForProjectParams, authInfo runtime.ClientAuthInfoWriter) (*ListClustersForProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListClustersForProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listClustersForProject",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListClustersForProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListClustersForProjectOK), nil

}

/*
ListNodeDeploymentNodes lists nodes that belong to the given node deployment
*/
func (a *Client) ListNodeDeploymentNodes(params *ListNodeDeploymentNodesParams, authInfo runtime.ClientAuthInfoWriter) (*ListNodeDeploymentNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNodeDeploymentNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listNodeDeploymentNodes",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNodeDeploymentNodesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListNodeDeploymentNodesOK), nil

}

/*
ListNodeDeploymentNodesEvents lists node deployment events if query parameter type is set to warning then only warning events are retrieved

If the value is 'normal' then normal events are returned. If the query parameter is missing method returns all events.
*/
func (a *Client) ListNodeDeploymentNodesEvents(params *ListNodeDeploymentNodesEventsParams, authInfo runtime.ClientAuthInfoWriter) (*ListNodeDeploymentNodesEventsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNodeDeploymentNodesEventsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listNodeDeploymentNodesEvents",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}/nodes/events",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNodeDeploymentNodesEventsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListNodeDeploymentNodesEventsOK), nil

}

/*
ListNodeDeployments Lists node deployments that belong to the given cluster
*/
func (a *Client) ListNodeDeployments(params *ListNodeDeploymentsParams, authInfo runtime.ClientAuthInfoWriter) (*ListNodeDeploymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNodeDeploymentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listNodeDeployments",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNodeDeploymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListNodeDeploymentsOK), nil

}

/*
ListNodesForClusterLegacy deprecateds lists nodes that belong to the given cluster

This endpoint is deprecated, please create a Node Deployment instead.
*/
func (a *Client) ListNodesForClusterLegacy(params *ListNodesForClusterLegacyParams, authInfo runtime.ClientAuthInfoWriter) (*ListNodesForClusterLegacyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNodesForClusterLegacyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listNodesForClusterLegacy",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNodesForClusterLegacyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListNodesForClusterLegacyOK), nil

}

/*
ListProjects lists projects that an authenticated user is a member of
*/
func (a *Client) ListProjects(params *ListProjectsParams, authInfo runtime.ClientAuthInfoWriter) (*ListProjectsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListProjectsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listProjects",
		Method:             "GET",
		PathPattern:        "/api/v1/projects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListProjectsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListProjectsOK), nil

}

/*
ListSSHKeys lists SSH keys that belong to the given project

The returned collection is sorted by creation timestamp.
*/
func (a *Client) ListSSHKeys(params *ListSSHKeysParams, authInfo runtime.ClientAuthInfoWriter) (*ListSSHKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSSHKeysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSSHKeys",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/sshkeys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSSHKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSSHKeysOK), nil

}

/*
ListSSHKeysAssignedToCluster Lists ssh keys that are assigned to the cluster
The returned collection is sorted by creation timestamp.
*/
func (a *Client) ListSSHKeysAssignedToCluster(params *ListSSHKeysAssignedToClusterParams, authInfo runtime.ClientAuthInfoWriter) (*ListSSHKeysAssignedToClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSSHKeysAssignedToClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSSHKeysAssignedToCluster",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/sshkeys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSSHKeysAssignedToClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSSHKeysAssignedToClusterOK), nil

}

/*
PatchCluster patches the given cluster using JSON merge patch method https tools ietf org html rfc7396
*/
func (a *Client) PatchCluster(params *PatchClusterParams, authInfo runtime.ClientAuthInfoWriter) (*PatchClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchCluster",
		Method:             "PATCH",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchClusterReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchClusterOK), nil

}

/*
PatchNodeDeployment Patches a node deployment that is assigned to the given cluster. Please note that at the moment only
node deployment's spec can be updated by a patch, no other fields can be changed using this endpoint.
*/
func (a *Client) PatchNodeDeployment(params *PatchNodeDeploymentParams, authInfo runtime.ClientAuthInfoWriter) (*PatchNodeDeploymentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchNodeDeploymentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchNodeDeployment",
		Method:             "PATCH",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodedeployments/{nodedeployment_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchNodeDeploymentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchNodeDeploymentOK), nil

}

/*
RevokeClusterAdminToken Revokes the current admin token
*/
func (a *Client) RevokeClusterAdminToken(params *RevokeClusterAdminTokenParams, authInfo runtime.ClientAuthInfoWriter) (*RevokeClusterAdminTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevokeClusterAdminTokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "revokeClusterAdminToken",
		Method:             "PUT",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RevokeClusterAdminTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RevokeClusterAdminTokenOK), nil

}

/*
UpdateProject Updates the given project
*/
func (a *Client) UpdateProject(params *UpdateProjectParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateProject",
		Method:             "PUT",
		PathPattern:        "/api/v1/projects/{project_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateProjectOK), nil

}

/*
UpgradeClusterNodeDeployments Upgrades node deployments in a cluster
*/
func (a *Client) UpgradeClusterNodeDeployments(params *UpgradeClusterNodeDeploymentsParams, authInfo runtime.ClientAuthInfoWriter) (*UpgradeClusterNodeDeploymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpgradeClusterNodeDeploymentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "upgradeClusterNodeDeployments",
		Method:             "PUT",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/nodes/upgrades",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpgradeClusterNodeDeploymentsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpgradeClusterNodeDeploymentsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
