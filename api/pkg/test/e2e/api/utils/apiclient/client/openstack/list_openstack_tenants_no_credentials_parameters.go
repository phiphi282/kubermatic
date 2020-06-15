// Code generated by go-swagger; DO NOT EDIT.

package openstack

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

// NewListOpenstackTenantsNoCredentialsParams creates a new ListOpenstackTenantsNoCredentialsParams object
// with the default values initialized.
func NewListOpenstackTenantsNoCredentialsParams() *ListOpenstackTenantsNoCredentialsParams {
	var ()
	return &ListOpenstackTenantsNoCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListOpenstackTenantsNoCredentialsParamsWithTimeout creates a new ListOpenstackTenantsNoCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOpenstackTenantsNoCredentialsParamsWithTimeout(timeout time.Duration) *ListOpenstackTenantsNoCredentialsParams {
	var ()
	return &ListOpenstackTenantsNoCredentialsParams{

		timeout: timeout,
	}
}

// NewListOpenstackTenantsNoCredentialsParamsWithContext creates a new ListOpenstackTenantsNoCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOpenstackTenantsNoCredentialsParamsWithContext(ctx context.Context) *ListOpenstackTenantsNoCredentialsParams {
	var ()
	return &ListOpenstackTenantsNoCredentialsParams{

		Context: ctx,
	}
}

// NewListOpenstackTenantsNoCredentialsParamsWithHTTPClient creates a new ListOpenstackTenantsNoCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOpenstackTenantsNoCredentialsParamsWithHTTPClient(client *http.Client) *ListOpenstackTenantsNoCredentialsParams {
	var ()
	return &ListOpenstackTenantsNoCredentialsParams{
		HTTPClient: client,
	}
}

/*ListOpenstackTenantsNoCredentialsParams contains all the parameters to send to the API endpoint
for the list openstack tenants no credentials operation typically these are written to a http.Request
*/
type ListOpenstackTenantsNoCredentialsParams struct {

	/*ClusterID*/
	ClusterID string
	/*Dc*/
	DC string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list openstack tenants no credentials params
func (o *ListOpenstackTenantsNoCredentialsParams) WithTimeout(timeout time.Duration) *ListOpenstackTenantsNoCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list openstack tenants no credentials params
func (o *ListOpenstackTenantsNoCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list openstack tenants no credentials params
func (o *ListOpenstackTenantsNoCredentialsParams) WithContext(ctx context.Context) *ListOpenstackTenantsNoCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list openstack tenants no credentials params
func (o *ListOpenstackTenantsNoCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list openstack tenants no credentials params
func (o *ListOpenstackTenantsNoCredentialsParams) WithHTTPClient(client *http.Client) *ListOpenstackTenantsNoCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list openstack tenants no credentials params
func (o *ListOpenstackTenantsNoCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list openstack tenants no credentials params
func (o *ListOpenstackTenantsNoCredentialsParams) WithClusterID(clusterID string) *ListOpenstackTenantsNoCredentialsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list openstack tenants no credentials params
func (o *ListOpenstackTenantsNoCredentialsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the list openstack tenants no credentials params
func (o *ListOpenstackTenantsNoCredentialsParams) WithDC(dc string) *ListOpenstackTenantsNoCredentialsParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list openstack tenants no credentials params
func (o *ListOpenstackTenantsNoCredentialsParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the list openstack tenants no credentials params
func (o *ListOpenstackTenantsNoCredentialsParams) WithProjectID(projectID string) *ListOpenstackTenantsNoCredentialsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list openstack tenants no credentials params
func (o *ListOpenstackTenantsNoCredentialsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListOpenstackTenantsNoCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
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
