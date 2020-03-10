// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListClusterRoleNamesParams creates a new ListClusterRoleNamesParams object
// with the default values initialized.
func NewListClusterRoleNamesParams() *ListClusterRoleNamesParams {
	var ()
	return &ListClusterRoleNamesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListClusterRoleNamesParamsWithTimeout creates a new ListClusterRoleNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListClusterRoleNamesParamsWithTimeout(timeout time.Duration) *ListClusterRoleNamesParams {
	var ()
	return &ListClusterRoleNamesParams{

		timeout: timeout,
	}
}

// NewListClusterRoleNamesParamsWithContext creates a new ListClusterRoleNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListClusterRoleNamesParamsWithContext(ctx context.Context) *ListClusterRoleNamesParams {
	var ()
	return &ListClusterRoleNamesParams{

		Context: ctx,
	}
}

// NewListClusterRoleNamesParamsWithHTTPClient creates a new ListClusterRoleNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListClusterRoleNamesParamsWithHTTPClient(client *http.Client) *ListClusterRoleNamesParams {
	var ()
	return &ListClusterRoleNamesParams{
		HTTPClient: client,
	}
}

/*ListClusterRoleNamesParams contains all the parameters to send to the API endpoint
for the list cluster role names operation typically these are written to a http.Request
*/
type ListClusterRoleNamesParams struct {

	/*ClusterID*/
	ClusterID string
	/*Dc*/
	Dc string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list cluster role names params
func (o *ListClusterRoleNamesParams) WithTimeout(timeout time.Duration) *ListClusterRoleNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list cluster role names params
func (o *ListClusterRoleNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list cluster role names params
func (o *ListClusterRoleNamesParams) WithContext(ctx context.Context) *ListClusterRoleNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list cluster role names params
func (o *ListClusterRoleNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list cluster role names params
func (o *ListClusterRoleNamesParams) WithHTTPClient(client *http.Client) *ListClusterRoleNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list cluster role names params
func (o *ListClusterRoleNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list cluster role names params
func (o *ListClusterRoleNamesParams) WithClusterID(clusterID string) *ListClusterRoleNamesParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list cluster role names params
func (o *ListClusterRoleNamesParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDc adds the dc to the list cluster role names params
func (o *ListClusterRoleNamesParams) WithDc(dc string) *ListClusterRoleNamesParams {
	o.SetDc(dc)
	return o
}

// SetDc adds the dc to the list cluster role names params
func (o *ListClusterRoleNamesParams) SetDc(dc string) {
	o.Dc = dc
}

// WithProjectID adds the projectID to the list cluster role names params
func (o *ListClusterRoleNamesParams) WithProjectID(projectID string) *ListClusterRoleNamesParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list cluster role names params
func (o *ListClusterRoleNamesParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListClusterRoleNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.Dc); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
