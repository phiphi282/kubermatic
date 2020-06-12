// Code generated by go-swagger; DO NOT EDIT.

package alibaba

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

// NewListAlibabaZonesNoCredentialsParams creates a new ListAlibabaZonesNoCredentialsParams object
// with the default values initialized.
func NewListAlibabaZonesNoCredentialsParams() *ListAlibabaZonesNoCredentialsParams {
	var ()
	return &ListAlibabaZonesNoCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAlibabaZonesNoCredentialsParamsWithTimeout creates a new ListAlibabaZonesNoCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAlibabaZonesNoCredentialsParamsWithTimeout(timeout time.Duration) *ListAlibabaZonesNoCredentialsParams {
	var ()
	return &ListAlibabaZonesNoCredentialsParams{

		timeout: timeout,
	}
}

// NewListAlibabaZonesNoCredentialsParamsWithContext creates a new ListAlibabaZonesNoCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAlibabaZonesNoCredentialsParamsWithContext(ctx context.Context) *ListAlibabaZonesNoCredentialsParams {
	var ()
	return &ListAlibabaZonesNoCredentialsParams{

		Context: ctx,
	}
}

// NewListAlibabaZonesNoCredentialsParamsWithHTTPClient creates a new ListAlibabaZonesNoCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAlibabaZonesNoCredentialsParamsWithHTTPClient(client *http.Client) *ListAlibabaZonesNoCredentialsParams {
	var ()
	return &ListAlibabaZonesNoCredentialsParams{
		HTTPClient: client,
	}
}

/*ListAlibabaZonesNoCredentialsParams contains all the parameters to send to the API endpoint
for the list alibaba zones no credentials operation typically these are written to a http.Request
*/
type ListAlibabaZonesNoCredentialsParams struct {

	/*Region*/
	Region *string
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

// WithTimeout adds the timeout to the list alibaba zones no credentials params
func (o *ListAlibabaZonesNoCredentialsParams) WithTimeout(timeout time.Duration) *ListAlibabaZonesNoCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list alibaba zones no credentials params
func (o *ListAlibabaZonesNoCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list alibaba zones no credentials params
func (o *ListAlibabaZonesNoCredentialsParams) WithContext(ctx context.Context) *ListAlibabaZonesNoCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list alibaba zones no credentials params
func (o *ListAlibabaZonesNoCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list alibaba zones no credentials params
func (o *ListAlibabaZonesNoCredentialsParams) WithHTTPClient(client *http.Client) *ListAlibabaZonesNoCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list alibaba zones no credentials params
func (o *ListAlibabaZonesNoCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegion adds the region to the list alibaba zones no credentials params
func (o *ListAlibabaZonesNoCredentialsParams) WithRegion(region *string) *ListAlibabaZonesNoCredentialsParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the list alibaba zones no credentials params
func (o *ListAlibabaZonesNoCredentialsParams) SetRegion(region *string) {
	o.Region = region
}

// WithClusterID adds the clusterID to the list alibaba zones no credentials params
func (o *ListAlibabaZonesNoCredentialsParams) WithClusterID(clusterID string) *ListAlibabaZonesNoCredentialsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list alibaba zones no credentials params
func (o *ListAlibabaZonesNoCredentialsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the list alibaba zones no credentials params
func (o *ListAlibabaZonesNoCredentialsParams) WithDC(dc string) *ListAlibabaZonesNoCredentialsParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the list alibaba zones no credentials params
func (o *ListAlibabaZonesNoCredentialsParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the list alibaba zones no credentials params
func (o *ListAlibabaZonesNoCredentialsParams) WithProjectID(projectID string) *ListAlibabaZonesNoCredentialsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list alibaba zones no credentials params
func (o *ListAlibabaZonesNoCredentialsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAlibabaZonesNoCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Region != nil {

		// header param Region
		if err := r.SetHeaderParam("Region", *o.Region); err != nil {
			return err
		}

	}

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
