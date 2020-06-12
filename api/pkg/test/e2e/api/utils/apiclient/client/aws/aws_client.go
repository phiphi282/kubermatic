// Code generated by go-swagger; DO NOT EDIT.

package aws

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new aws API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for aws API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListAWSSizes(params *ListAWSSizesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAWSSizesOK, error)

	ListAWSSizesNoCredentials(params *ListAWSSizesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAWSSizesNoCredentialsOK, error)

	ListAWSSubnets(params *ListAWSSubnetsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAWSSubnetsOK, error)

	ListAWSSubnetsNoCredentials(params *ListAWSSubnetsNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAWSSubnetsNoCredentialsOK, error)

	ListAWSVPCS(params *ListAWSVPCSParams, authInfo runtime.ClientAuthInfoWriter) (*ListAWSVPCSOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListAWSSizes lists available a w s sizes
*/
func (a *Client) ListAWSSizes(params *ListAWSSizesParams, authInfo runtime.ClientAuthInfoWriter) (*ListAWSSizesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAWSSizesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAWSSizes",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/aws/sizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAWSSizesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAWSSizesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAWSSizesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListAWSSizesNoCredentials Lists available AWS sizes
*/
func (a *Client) ListAWSSizesNoCredentials(params *ListAWSSizesNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAWSSizesNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAWSSizesNoCredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAWSSizesNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/aws/sizes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAWSSizesNoCredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAWSSizesNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAWSSizesNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListAWSSubnets Lists available AWS subnets
*/
func (a *Client) ListAWSSubnets(params *ListAWSSubnetsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAWSSubnetsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAWSSubnetsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAWSSubnets",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/aws/{dc}/subnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAWSSubnetsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAWSSubnetsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAWSSubnetsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListAWSSubnetsNoCredentials Lists available AWS subnets
*/
func (a *Client) ListAWSSubnetsNoCredentials(params *ListAWSSubnetsNoCredentialsParams, authInfo runtime.ClientAuthInfoWriter) (*ListAWSSubnetsNoCredentialsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAWSSubnetsNoCredentialsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAWSSubnetsNoCredentials",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/aws/subnets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAWSSubnetsNoCredentialsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAWSSubnetsNoCredentialsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAWSSubnetsNoCredentialsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListAWSVPCS Lists available AWS vpc's
*/
func (a *Client) ListAWSVPCS(params *ListAWSVPCSParams, authInfo runtime.ClientAuthInfoWriter) (*ListAWSVPCSOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAWSVPCSParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAWSVPCS",
		Method:             "GET",
		PathPattern:        "/api/v1/providers/aws/{dc}/vpcs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListAWSVPCSReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAWSVPCSOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAWSVPCSDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
