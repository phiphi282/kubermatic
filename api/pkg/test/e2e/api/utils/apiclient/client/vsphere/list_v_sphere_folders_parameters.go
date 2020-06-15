// Code generated by go-swagger; DO NOT EDIT.

package vsphere

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

// NewListVSphereFoldersParams creates a new ListVSphereFoldersParams object
// with the default values initialized.
func NewListVSphereFoldersParams() *ListVSphereFoldersParams {

	return &ListVSphereFoldersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListVSphereFoldersParamsWithTimeout creates a new ListVSphereFoldersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListVSphereFoldersParamsWithTimeout(timeout time.Duration) *ListVSphereFoldersParams {

	return &ListVSphereFoldersParams{

		timeout: timeout,
	}
}

// NewListVSphereFoldersParamsWithContext creates a new ListVSphereFoldersParams object
// with the default values initialized, and the ability to set a context for a request
func NewListVSphereFoldersParamsWithContext(ctx context.Context) *ListVSphereFoldersParams {

	return &ListVSphereFoldersParams{

		Context: ctx,
	}
}

// NewListVSphereFoldersParamsWithHTTPClient creates a new ListVSphereFoldersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListVSphereFoldersParamsWithHTTPClient(client *http.Client) *ListVSphereFoldersParams {

	return &ListVSphereFoldersParams{
		HTTPClient: client,
	}
}

/*ListVSphereFoldersParams contains all the parameters to send to the API endpoint
for the list v sphere folders operation typically these are written to a http.Request
*/
type ListVSphereFoldersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list v sphere folders params
func (o *ListVSphereFoldersParams) WithTimeout(timeout time.Duration) *ListVSphereFoldersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list v sphere folders params
func (o *ListVSphereFoldersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list v sphere folders params
func (o *ListVSphereFoldersParams) WithContext(ctx context.Context) *ListVSphereFoldersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list v sphere folders params
func (o *ListVSphereFoldersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list v sphere folders params
func (o *ListVSphereFoldersParams) WithHTTPClient(client *http.Client) *ListVSphereFoldersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list v sphere folders params
func (o *ListVSphereFoldersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListVSphereFoldersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
