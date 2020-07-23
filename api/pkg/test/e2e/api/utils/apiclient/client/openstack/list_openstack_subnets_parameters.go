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

// NewListOpenstackSubnetsParams creates a new ListOpenstackSubnetsParams object
// with the default values initialized.
func NewListOpenstackSubnetsParams() *ListOpenstackSubnetsParams {
	var ()
	return &ListOpenstackSubnetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListOpenstackSubnetsParamsWithTimeout creates a new ListOpenstackSubnetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListOpenstackSubnetsParamsWithTimeout(timeout time.Duration) *ListOpenstackSubnetsParams {
	var ()
	return &ListOpenstackSubnetsParams{

		timeout: timeout,
	}
}

// NewListOpenstackSubnetsParamsWithContext creates a new ListOpenstackSubnetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListOpenstackSubnetsParamsWithContext(ctx context.Context) *ListOpenstackSubnetsParams {
	var ()
	return &ListOpenstackSubnetsParams{

		Context: ctx,
	}
}

// NewListOpenstackSubnetsParamsWithHTTPClient creates a new ListOpenstackSubnetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListOpenstackSubnetsParamsWithHTTPClient(client *http.Client) *ListOpenstackSubnetsParams {
	var ()
	return &ListOpenstackSubnetsParams{
		HTTPClient: client,
	}
}

/*ListOpenstackSubnetsParams contains all the parameters to send to the API endpoint
for the list openstack subnets operation typically these are written to a http.Request
*/
type ListOpenstackSubnetsParams struct {

	/*Credential*/
	Credential *string
	/*DatacenterName*/
	DatacenterName *string
	/*Domain*/
	Domain *string
	/*Password*/
	Password *string
	/*Tenant*/
	Tenant *string
	/*TenantID*/
	TenantID *string
	/*Username*/
	Username *string
	/*NetworkID*/
	NetworkID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) WithTimeout(timeout time.Duration) *ListOpenstackSubnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) WithContext(ctx context.Context) *ListOpenstackSubnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) WithHTTPClient(client *http.Client) *ListOpenstackSubnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) WithCredential(credential *string) *ListOpenstackSubnetsParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithDatacenterName adds the datacenterName to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) WithDatacenterName(datacenterName *string) *ListOpenstackSubnetsParams {
	o.SetDatacenterName(datacenterName)
	return o
}

// SetDatacenterName adds the datacenterName to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) SetDatacenterName(datacenterName *string) {
	o.DatacenterName = datacenterName
}

// WithDomain adds the domain to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) WithDomain(domain *string) *ListOpenstackSubnetsParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) SetDomain(domain *string) {
	o.Domain = domain
}

// WithPassword adds the password to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) WithPassword(password *string) *ListOpenstackSubnetsParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) SetPassword(password *string) {
	o.Password = password
}

// WithTenant adds the tenant to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) WithTenant(tenant *string) *ListOpenstackSubnetsParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantID adds the tenantID to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) WithTenantID(tenantID *string) *ListOpenstackSubnetsParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithUsername adds the username to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) WithUsername(username *string) *ListOpenstackSubnetsParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) SetUsername(username *string) {
	o.Username = username
}

// WithNetworkID adds the networkID to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) WithNetworkID(networkID *string) *ListOpenstackSubnetsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the list openstack subnets params
func (o *ListOpenstackSubnetsParams) SetNetworkID(networkID *string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *ListOpenstackSubnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}

	}

	if o.DatacenterName != nil {

		// header param DatacenterName
		if err := r.SetHeaderParam("DatacenterName", *o.DatacenterName); err != nil {
			return err
		}

	}

	if o.Domain != nil {

		// header param Domain
		if err := r.SetHeaderParam("Domain", *o.Domain); err != nil {
			return err
		}

	}

	if o.Password != nil {

		// header param Password
		if err := r.SetHeaderParam("Password", *o.Password); err != nil {
			return err
		}

	}

	if o.Tenant != nil {

		// header param Tenant
		if err := r.SetHeaderParam("Tenant", *o.Tenant); err != nil {
			return err
		}

	}

	if o.TenantID != nil {

		// header param TenantID
		if err := r.SetHeaderParam("TenantID", *o.TenantID); err != nil {
			return err
		}

	}

	if o.Username != nil {

		// header param Username
		if err := r.SetHeaderParam("Username", *o.Username); err != nil {
			return err
		}

	}

	if o.NetworkID != nil {

		// query param network_id
		var qrNetworkID string
		if o.NetworkID != nil {
			qrNetworkID = *o.NetworkID
		}
		qNetworkID := qrNetworkID
		if qNetworkID != "" {
			if err := r.SetQueryParam("network_id", qNetworkID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
