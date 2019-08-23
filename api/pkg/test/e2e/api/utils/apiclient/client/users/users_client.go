// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddUserToProject Adds the given user to the given project
*/
func (a *Client) AddUserToProject(params *AddUserToProjectParams, authInfo runtime.ClientAuthInfoWriter) (*AddUserToProjectCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddUserToProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addUserToProject",
		Method:             "POST",
		PathPattern:        "/api/v1/projects/{project_id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddUserToProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddUserToProjectCreated), nil

}

/*
DeleteUserFromProject Removes the given member from the project
*/
func (a *Client) DeleteUserFromProject(params *DeleteUserFromProjectParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUserFromProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserFromProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUserFromProject",
		Method:             "DELETE",
		PathPattern:        "/api/v1/projects/{project_id}/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserFromProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUserFromProjectOK), nil

}

/*
EditUserInProject Changes membership of the given user for the given project
*/
func (a *Client) EditUserInProject(params *EditUserInProjectParams, authInfo runtime.ClientAuthInfoWriter) (*EditUserInProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditUserInProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "editUserInProject",
		Method:             "PUT",
		PathPattern:        "/api/v1/projects/{project_id}/users/{user_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EditUserInProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EditUserInProjectOK), nil

}

/*
GetCurrentUser returns information about the current user
*/
func (a *Client) GetCurrentUser(params *GetCurrentUserParams, authInfo runtime.ClientAuthInfoWriter) (*GetCurrentUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrentUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCurrentUser",
		Method:             "GET",
		PathPattern:        "/api/v1/me",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCurrentUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCurrentUserOK), nil

}

/*
GetUsersForProject Get list of users for the given project
*/
func (a *Client) GetUsersForProject(params *GetUsersForProjectParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersForProjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersForProjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsersForProject",
		Method:             "GET",
		PathPattern:        "/api/v1/projects/{project_id}/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersForProjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUsersForProjectOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
