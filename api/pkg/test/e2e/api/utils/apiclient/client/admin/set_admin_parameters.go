// Code generated by go-swagger; DO NOT EDIT.

package admin

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

	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// NewSetAdminParams creates a new SetAdminParams object
// with the default values initialized.
func NewSetAdminParams() *SetAdminParams {
	var ()
	return &SetAdminParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetAdminParamsWithTimeout creates a new SetAdminParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetAdminParamsWithTimeout(timeout time.Duration) *SetAdminParams {
	var ()
	return &SetAdminParams{

		timeout: timeout,
	}
}

// NewSetAdminParamsWithContext creates a new SetAdminParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetAdminParamsWithContext(ctx context.Context) *SetAdminParams {
	var ()
	return &SetAdminParams{

		Context: ctx,
	}
}

// NewSetAdminParamsWithHTTPClient creates a new SetAdminParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetAdminParamsWithHTTPClient(client *http.Client) *SetAdminParams {
	var ()
	return &SetAdminParams{
		HTTPClient: client,
	}
}

/*SetAdminParams contains all the parameters to send to the API endpoint
for the set admin operation typically these are written to a http.Request
*/
type SetAdminParams struct {

	/*Body*/
	Body *models.Admin

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set admin params
func (o *SetAdminParams) WithTimeout(timeout time.Duration) *SetAdminParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set admin params
func (o *SetAdminParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set admin params
func (o *SetAdminParams) WithContext(ctx context.Context) *SetAdminParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set admin params
func (o *SetAdminParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set admin params
func (o *SetAdminParams) WithHTTPClient(client *http.Client) *SetAdminParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set admin params
func (o *SetAdminParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set admin params
func (o *SetAdminParams) WithBody(body *models.Admin) *SetAdminParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set admin params
func (o *SetAdminParams) SetBody(body *models.Admin) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SetAdminParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
