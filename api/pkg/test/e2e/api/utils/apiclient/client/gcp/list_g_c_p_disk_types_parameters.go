// Code generated by go-swagger; DO NOT EDIT.

package gcp

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

// NewListGCPDiskTypesParams creates a new ListGCPDiskTypesParams object
// with the default values initialized.
func NewListGCPDiskTypesParams() *ListGCPDiskTypesParams {
	var ()
	return &ListGCPDiskTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListGCPDiskTypesParamsWithTimeout creates a new ListGCPDiskTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListGCPDiskTypesParamsWithTimeout(timeout time.Duration) *ListGCPDiskTypesParams {
	var ()
	return &ListGCPDiskTypesParams{

		timeout: timeout,
	}
}

// NewListGCPDiskTypesParamsWithContext creates a new ListGCPDiskTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListGCPDiskTypesParamsWithContext(ctx context.Context) *ListGCPDiskTypesParams {
	var ()
	return &ListGCPDiskTypesParams{

		Context: ctx,
	}
}

// NewListGCPDiskTypesParamsWithHTTPClient creates a new ListGCPDiskTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListGCPDiskTypesParamsWithHTTPClient(client *http.Client) *ListGCPDiskTypesParams {
	var ()
	return &ListGCPDiskTypesParams{
		HTTPClient: client,
	}
}

/*ListGCPDiskTypesParams contains all the parameters to send to the API endpoint
for the list g c p disk types operation typically these are written to a http.Request
*/
type ListGCPDiskTypesParams struct {

	/*Credential*/
	Credential *string
	/*ServiceAccount*/
	ServiceAccount *string
	/*Zone*/
	Zone *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list g c p disk types params
func (o *ListGCPDiskTypesParams) WithTimeout(timeout time.Duration) *ListGCPDiskTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list g c p disk types params
func (o *ListGCPDiskTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list g c p disk types params
func (o *ListGCPDiskTypesParams) WithContext(ctx context.Context) *ListGCPDiskTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list g c p disk types params
func (o *ListGCPDiskTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list g c p disk types params
func (o *ListGCPDiskTypesParams) WithHTTPClient(client *http.Client) *ListGCPDiskTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list g c p disk types params
func (o *ListGCPDiskTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list g c p disk types params
func (o *ListGCPDiskTypesParams) WithCredential(credential *string) *ListGCPDiskTypesParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list g c p disk types params
func (o *ListGCPDiskTypesParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithServiceAccount adds the serviceAccount to the list g c p disk types params
func (o *ListGCPDiskTypesParams) WithServiceAccount(serviceAccount *string) *ListGCPDiskTypesParams {
	o.SetServiceAccount(serviceAccount)
	return o
}

// SetServiceAccount adds the serviceAccount to the list g c p disk types params
func (o *ListGCPDiskTypesParams) SetServiceAccount(serviceAccount *string) {
	o.ServiceAccount = serviceAccount
}

// WithZone adds the zone to the list g c p disk types params
func (o *ListGCPDiskTypesParams) WithZone(zone *string) *ListGCPDiskTypesParams {
	o.SetZone(zone)
	return o
}

// SetZone adds the zone to the list g c p disk types params
func (o *ListGCPDiskTypesParams) SetZone(zone *string) {
	o.Zone = zone
}

// WriteToRequest writes these params to a swagger request
func (o *ListGCPDiskTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Credential != nil {

		// query param Credential
		var qrCredential string
		if o.Credential != nil {
			qrCredential = *o.Credential
		}
		qCredential := qrCredential
		if qCredential != "" {
			if err := r.SetQueryParam("Credential", qCredential); err != nil {
				return err
			}
		}

	}

	if o.ServiceAccount != nil {

		// header param ServiceAccount
		if err := r.SetHeaderParam("ServiceAccount", *o.ServiceAccount); err != nil {
			return err
		}

	}

	if o.Zone != nil {

		// header param Zone
		if err := r.SetHeaderParam("Zone", *o.Zone); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
