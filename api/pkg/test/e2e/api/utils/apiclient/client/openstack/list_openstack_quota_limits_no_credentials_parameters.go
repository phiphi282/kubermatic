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

	strfmt "github.com/go-openapi/strfmt"
)

// NewListOpenstackQuotaLimitsNoCredentialsParams creates a new ListOpenstackQuotaLimitsNoCredentialsParams object
// with the default values initialized.
func NewListOpenstackQuotaLimitsNoCredentialsParams() *ListOpenstackQuotaLimitsNoCredentialsParams {
	var ()
	return &ListOpenstackQuotaLimitsNoCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListOpenstackQuotaLimitsNoCredentialsParamsWithTimeout creates a new ListOpenstackQuotaLimitsNoCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOpenstackQuotaLimitsNoCredentialsParamsWithTimeout(timeout time.Duration) *ListOpenstackQuotaLimitsNoCredentialsParams {
	var ()
	return &ListOpenstackQuotaLimitsNoCredentialsParams{

		timeout: timeout,
	}
}

// NewListOpenstackQuotaLimitsNoCredentialsParamsWithContext creates a new ListOpenstackQuotaLimitsNoCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOpenstackQuotaLimitsNoCredentialsParamsWithContext(ctx context.Context) *ListOpenstackQuotaLimitsNoCredentialsParams {
	var ()
	return &ListOpenstackQuotaLimitsNoCredentialsParams{

		Context: ctx,
	}
}

// NewListOpenstackQuotaLimitsNoCredentialsParamsWithHTTPClient creates a new ListOpenstackQuotaLimitsNoCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOpenstackQuotaLimitsNoCredentialsParamsWithHTTPClient(client *http.Client) *ListOpenstackQuotaLimitsNoCredentialsParams {
	var ()
	return &ListOpenstackQuotaLimitsNoCredentialsParams{
		HTTPClient: client,
	}
}

/*ListOpenstackQuotaLimitsNoCredentialsParams contains all the parameters to send to the API endpoint
for the list openstack quota limits no credentials operation typically these are written to a http.Request
*/
type ListOpenstackQuotaLimitsNoCredentialsParams struct {

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

// WithTimeout adds the timeout to the list openstack quota limits no credentials params
func (o *ListOpenstackQuotaLimitsNoCredentialsParams) WithTimeout(timeout time.Duration) *ListOpenstackQuotaLimitsNoCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list openstack quota limits no credentials params
func (o *ListOpenstackQuotaLimitsNoCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list openstack quota limits no credentials params
func (o *ListOpenstackQuotaLimitsNoCredentialsParams) WithContext(ctx context.Context) *ListOpenstackQuotaLimitsNoCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list openstack quota limits no credentials params
func (o *ListOpenstackQuotaLimitsNoCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list openstack quota limits no credentials params
func (o *ListOpenstackQuotaLimitsNoCredentialsParams) WithHTTPClient(client *http.Client) *ListOpenstackQuotaLimitsNoCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list openstack quota limits no credentials params
func (o *ListOpenstackQuotaLimitsNoCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list openstack quota limits no credentials params
func (o *ListOpenstackQuotaLimitsNoCredentialsParams) WithClusterID(clusterID string) *ListOpenstackQuotaLimitsNoCredentialsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list openstack quota limits no credentials params
func (o *ListOpenstackQuotaLimitsNoCredentialsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDc adds the dc to the list openstack quota limits no credentials params
func (o *ListOpenstackQuotaLimitsNoCredentialsParams) WithDc(dc string) *ListOpenstackQuotaLimitsNoCredentialsParams {
	o.SetDc(dc)
	return o
}

// SetDc adds the dc to the list openstack quota limits no credentials params
func (o *ListOpenstackQuotaLimitsNoCredentialsParams) SetDc(dc string) {
	o.Dc = dc
}

// WithProjectID adds the projectID to the list openstack quota limits no credentials params
func (o *ListOpenstackQuotaLimitsNoCredentialsParams) WithProjectID(projectID string) *ListOpenstackQuotaLimitsNoCredentialsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list openstack quota limits no credentials params
func (o *ListOpenstackQuotaLimitsNoCredentialsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListOpenstackQuotaLimitsNoCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
