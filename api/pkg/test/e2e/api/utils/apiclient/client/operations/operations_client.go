// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Addon Lists names of addons that can be configured inside the user clusters
*/
func (a *Client) Addon(params *AddonParams, authInfo runtime.ClientAuthInfoWriter) (*AddonOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddonParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addon",
		Method:             "POST",
		PathPattern:        "/api/v1/addons",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddonReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddonOK), nil

}

/*
CreateOIDCKubeconfig Starts OIDC flow and generates kubeconfig, the generated config
contains OIDC provider authentication info
*/
func (a *Client) CreateOIDCKubeconfig(params *CreateOIDCKubeconfigParams, authInfo runtime.ClientAuthInfoWriter) (*CreateOIDCKubeconfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateOIDCKubeconfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createOIDCKubeconfig",
		Method:             "GET",
		PathPattern:        "/api/v1/kubeconfig",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateOIDCKubeconfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateOIDCKubeconfigOK), nil

}

/*
GetAddonConfig returns specified addon config
*/
func (a *Client) GetAddonConfig(params *GetAddonConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetAddonConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAddonConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAddonConfig",
		Method:             "GET",
		PathPattern:        "/api/v1/addonconfigs/{addon_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAddonConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAddonConfigOK), nil

}

/*
ListAddonConfigs returns all available addon configs
*/
func (a *Client) ListAddonConfigs(params *ListAddonConfigsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAddonConfigsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAddonConfigsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAddonConfigs",
		Method:             "GET",
		PathPattern:        "/api/v1/addonconfigs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAddonConfigsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListAddonConfigsOK), nil

}

/*
ListSystemLabels List restricted system labels
*/
func (a *Client) ListSystemLabels(params *ListSystemLabelsParams, authInfo runtime.ClientAuthInfoWriter) (*ListSystemLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSystemLabelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSystemLabels",
		Method:             "PATCH",
		PathPattern:        "/api/v1/labels/system",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListSystemLabelsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSystemLabelsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
