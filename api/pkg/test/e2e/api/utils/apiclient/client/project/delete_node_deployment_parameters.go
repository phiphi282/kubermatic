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

// NewDeleteNodeDeploymentParams creates a new DeleteNodeDeploymentParams object
// with the default values initialized.
func NewDeleteNodeDeploymentParams() *DeleteNodeDeploymentParams {
	var ()
	return &DeleteNodeDeploymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNodeDeploymentParamsWithTimeout creates a new DeleteNodeDeploymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNodeDeploymentParamsWithTimeout(timeout time.Duration) *DeleteNodeDeploymentParams {
	var ()
	return &DeleteNodeDeploymentParams{

		timeout: timeout,
	}
}

// NewDeleteNodeDeploymentParamsWithContext creates a new DeleteNodeDeploymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNodeDeploymentParamsWithContext(ctx context.Context) *DeleteNodeDeploymentParams {
	var ()
	return &DeleteNodeDeploymentParams{

		Context: ctx,
	}
}

// NewDeleteNodeDeploymentParamsWithHTTPClient creates a new DeleteNodeDeploymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNodeDeploymentParamsWithHTTPClient(client *http.Client) *DeleteNodeDeploymentParams {
	var ()
	return &DeleteNodeDeploymentParams{
		HTTPClient: client,
	}
}

/*DeleteNodeDeploymentParams contains all the parameters to send to the API endpoint
for the delete node deployment operation typically these are written to a http.Request
*/
type DeleteNodeDeploymentParams struct {

	/*ClusterID*/
	ClusterID string
	/*Dc*/
	Dc string
	/*NodedeploymentID*/
	NodedeploymentID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete node deployment params
func (o *DeleteNodeDeploymentParams) WithTimeout(timeout time.Duration) *DeleteNodeDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete node deployment params
func (o *DeleteNodeDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete node deployment params
func (o *DeleteNodeDeploymentParams) WithContext(ctx context.Context) *DeleteNodeDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete node deployment params
func (o *DeleteNodeDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete node deployment params
func (o *DeleteNodeDeploymentParams) WithHTTPClient(client *http.Client) *DeleteNodeDeploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete node deployment params
func (o *DeleteNodeDeploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the delete node deployment params
func (o *DeleteNodeDeploymentParams) WithClusterID(clusterID string) *DeleteNodeDeploymentParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete node deployment params
func (o *DeleteNodeDeploymentParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDc adds the dc to the delete node deployment params
func (o *DeleteNodeDeploymentParams) WithDc(dc string) *DeleteNodeDeploymentParams {
	o.SetDc(dc)
	return o
}

// SetDc adds the dc to the delete node deployment params
func (o *DeleteNodeDeploymentParams) SetDc(dc string) {
	o.Dc = dc
}

// WithNodedeploymentID adds the nodedeploymentID to the delete node deployment params
func (o *DeleteNodeDeploymentParams) WithNodedeploymentID(nodedeploymentID string) *DeleteNodeDeploymentParams {
	o.SetNodedeploymentID(nodedeploymentID)
	return o
}

// SetNodedeploymentID adds the nodedeploymentId to the delete node deployment params
func (o *DeleteNodeDeploymentParams) SetNodedeploymentID(nodedeploymentID string) {
	o.NodedeploymentID = nodedeploymentID
}

// WithProjectID adds the projectID to the delete node deployment params
func (o *DeleteNodeDeploymentParams) WithProjectID(projectID string) *DeleteNodeDeploymentParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete node deployment params
func (o *DeleteNodeDeploymentParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNodeDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param nodedeployment_id
	if err := r.SetPathParam("nodedeployment_id", o.NodedeploymentID); err != nil {
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
