// Code generated by go-swagger; DO NOT EDIT.

package preset

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

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// NewUpdatePresetParams creates a new UpdatePresetParams object
// with the default values initialized.
func NewUpdatePresetParams() *UpdatePresetParams {
	var ()
	return &UpdatePresetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePresetParamsWithTimeout creates a new UpdatePresetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePresetParamsWithTimeout(timeout time.Duration) *UpdatePresetParams {
	var ()
	return &UpdatePresetParams{

		timeout: timeout,
	}
}

// NewUpdatePresetParamsWithContext creates a new UpdatePresetParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePresetParamsWithContext(ctx context.Context) *UpdatePresetParams {
	var ()
	return &UpdatePresetParams{

		Context: ctx,
	}
}

// NewUpdatePresetParamsWithHTTPClient creates a new UpdatePresetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePresetParamsWithHTTPClient(client *http.Client) *UpdatePresetParams {
	var ()
	return &UpdatePresetParams{
		HTTPClient: client,
	}
}

/*UpdatePresetParams contains all the parameters to send to the API endpoint
for the update preset operation typically these are written to a http.Request
*/
type UpdatePresetParams struct {

	/*Body*/
	Body *models.Preset
	/*ProviderName*/
	ProviderName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update preset params
func (o *UpdatePresetParams) WithTimeout(timeout time.Duration) *UpdatePresetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update preset params
func (o *UpdatePresetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update preset params
func (o *UpdatePresetParams) WithContext(ctx context.Context) *UpdatePresetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update preset params
func (o *UpdatePresetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update preset params
func (o *UpdatePresetParams) WithHTTPClient(client *http.Client) *UpdatePresetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update preset params
func (o *UpdatePresetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update preset params
func (o *UpdatePresetParams) WithBody(body *models.Preset) *UpdatePresetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update preset params
func (o *UpdatePresetParams) SetBody(body *models.Preset) {
	o.Body = body
}

// WithProviderName adds the providerName to the update preset params
func (o *UpdatePresetParams) WithProviderName(providerName string) *UpdatePresetParams {
	o.SetProviderName(providerName)
	return o
}

// SetProviderName adds the providerName to the update preset params
func (o *UpdatePresetParams) SetProviderName(providerName string) {
	o.ProviderName = providerName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePresetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param provider_name
	if err := r.SetPathParam("provider_name", o.ProviderName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}