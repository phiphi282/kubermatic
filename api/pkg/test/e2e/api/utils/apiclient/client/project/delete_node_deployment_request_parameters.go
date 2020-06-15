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
	"github.com/go-openapi/strfmt"
)

// NewDeleteNodeDeploymentRequestParams creates a new DeleteNodeDeploymentRequestParams object
// with the default values initialized.
func NewDeleteNodeDeploymentRequestParams() *DeleteNodeDeploymentRequestParams {
	var ()
	return &DeleteNodeDeploymentRequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNodeDeploymentRequestParamsWithTimeout creates a new DeleteNodeDeploymentRequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNodeDeploymentRequestParamsWithTimeout(timeout time.Duration) *DeleteNodeDeploymentRequestParams {
	var ()
	return &DeleteNodeDeploymentRequestParams{

		timeout: timeout,
	}
}

// NewDeleteNodeDeploymentRequestParamsWithContext creates a new DeleteNodeDeploymentRequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNodeDeploymentRequestParamsWithContext(ctx context.Context) *DeleteNodeDeploymentRequestParams {
	var ()
	return &DeleteNodeDeploymentRequestParams{

		Context: ctx,
	}
}

// NewDeleteNodeDeploymentRequestParamsWithHTTPClient creates a new DeleteNodeDeploymentRequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNodeDeploymentRequestParamsWithHTTPClient(client *http.Client) *DeleteNodeDeploymentRequestParams {
	var ()
	return &DeleteNodeDeploymentRequestParams{
		HTTPClient: client,
	}
}

/*DeleteNodeDeploymentRequestParams contains all the parameters to send to the API endpoint
for the delete node deployment request operation typically these are written to a http.Request
*/
type DeleteNodeDeploymentRequestParams struct {

	/*ClusterID*/
	ClusterID string
	/*Dc*/
	DC string
	/*NdrequestID*/
	NdrName string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete node deployment request params
func (o *DeleteNodeDeploymentRequestParams) WithTimeout(timeout time.Duration) *DeleteNodeDeploymentRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete node deployment request params
func (o *DeleteNodeDeploymentRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete node deployment request params
func (o *DeleteNodeDeploymentRequestParams) WithContext(ctx context.Context) *DeleteNodeDeploymentRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete node deployment request params
func (o *DeleteNodeDeploymentRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete node deployment request params
func (o *DeleteNodeDeploymentRequestParams) WithHTTPClient(client *http.Client) *DeleteNodeDeploymentRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete node deployment request params
func (o *DeleteNodeDeploymentRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the delete node deployment request params
func (o *DeleteNodeDeploymentRequestParams) WithClusterID(clusterID string) *DeleteNodeDeploymentRequestParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete node deployment request params
func (o *DeleteNodeDeploymentRequestParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the delete node deployment request params
func (o *DeleteNodeDeploymentRequestParams) WithDC(dc string) *DeleteNodeDeploymentRequestParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the delete node deployment request params
func (o *DeleteNodeDeploymentRequestParams) SetDC(dc string) {
	o.DC = dc
}

// WithNdrName adds the ndrequestID to the delete node deployment request params
func (o *DeleteNodeDeploymentRequestParams) WithNdrName(ndrequestID string) *DeleteNodeDeploymentRequestParams {
	o.SetNdrName(ndrequestID)
	return o
}

// SetNdrName adds the ndrequestId to the delete node deployment request params
func (o *DeleteNodeDeploymentRequestParams) SetNdrName(ndrequestID string) {
	o.NdrName = ndrequestID
}

// WithProjectID adds the projectID to the delete node deployment request params
func (o *DeleteNodeDeploymentRequestParams) WithProjectID(projectID string) *DeleteNodeDeploymentRequestParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete node deployment request params
func (o *DeleteNodeDeploymentRequestParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNodeDeploymentRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param ndrequest_id
	if err := r.SetPathParam("ndrequest_id", o.NdrName); err != nil {
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
