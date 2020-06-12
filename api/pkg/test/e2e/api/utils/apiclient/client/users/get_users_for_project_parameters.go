// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetUsersForProjectParams creates a new GetUsersForProjectParams object
// with the default values initialized.
func NewGetUsersForProjectParams() *GetUsersForProjectParams {
	var ()
	return &GetUsersForProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsersForProjectParamsWithTimeout creates a new GetUsersForProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsersForProjectParamsWithTimeout(timeout time.Duration) *GetUsersForProjectParams {
	var ()
	return &GetUsersForProjectParams{

		timeout: timeout,
	}
}

// NewGetUsersForProjectParamsWithContext creates a new GetUsersForProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsersForProjectParamsWithContext(ctx context.Context) *GetUsersForProjectParams {
	var ()
	return &GetUsersForProjectParams{

		Context: ctx,
	}
}

// NewGetUsersForProjectParamsWithHTTPClient creates a new GetUsersForProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsersForProjectParamsWithHTTPClient(client *http.Client) *GetUsersForProjectParams {
	var ()
	return &GetUsersForProjectParams{
		HTTPClient: client,
	}
}

/*GetUsersForProjectParams contains all the parameters to send to the API endpoint
for the get users for project operation typically these are written to a http.Request
*/
type GetUsersForProjectParams struct {

	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get users for project params
func (o *GetUsersForProjectParams) WithTimeout(timeout time.Duration) *GetUsersForProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get users for project params
func (o *GetUsersForProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get users for project params
func (o *GetUsersForProjectParams) WithContext(ctx context.Context) *GetUsersForProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get users for project params
func (o *GetUsersForProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get users for project params
func (o *GetUsersForProjectParams) WithHTTPClient(client *http.Client) *GetUsersForProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get users for project params
func (o *GetUsersForProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the get users for project params
func (o *GetUsersForProjectParams) WithProjectID(projectID string) *GetUsersForProjectParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get users for project params
func (o *GetUsersForProjectParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsersForProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
