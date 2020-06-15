// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// UpdateAdmissionPluginReader is a Reader for the UpdateAdmissionPlugin structure.
type UpdateAdmissionPluginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAdmissionPluginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAdmissionPluginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateAdmissionPluginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateAdmissionPluginForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateAdmissionPluginDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAdmissionPluginOK creates a UpdateAdmissionPluginOK with default headers values
func NewUpdateAdmissionPluginOK() *UpdateAdmissionPluginOK {
	return &UpdateAdmissionPluginOK{}
}

/*UpdateAdmissionPluginOK handles this case with default header values.

AdmissionPlugin
*/
type UpdateAdmissionPluginOK struct {
	Payload *models.AdmissionPlugin
}

func (o *UpdateAdmissionPluginOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/admission/plugins/{name}][%d] updateAdmissionPluginOK  %+v", 200, o.Payload)
}

func (o *UpdateAdmissionPluginOK) GetPayload() *models.AdmissionPlugin {
	return o.Payload
}

func (o *UpdateAdmissionPluginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdmissionPlugin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAdmissionPluginUnauthorized creates a UpdateAdmissionPluginUnauthorized with default headers values
func NewUpdateAdmissionPluginUnauthorized() *UpdateAdmissionPluginUnauthorized {
	return &UpdateAdmissionPluginUnauthorized{}
}

/*UpdateAdmissionPluginUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type UpdateAdmissionPluginUnauthorized struct {
}

func (o *UpdateAdmissionPluginUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/admission/plugins/{name}][%d] updateAdmissionPluginUnauthorized ", 401)
}

func (o *UpdateAdmissionPluginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAdmissionPluginForbidden creates a UpdateAdmissionPluginForbidden with default headers values
func NewUpdateAdmissionPluginForbidden() *UpdateAdmissionPluginForbidden {
	return &UpdateAdmissionPluginForbidden{}
}

/*UpdateAdmissionPluginForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type UpdateAdmissionPluginForbidden struct {
}

func (o *UpdateAdmissionPluginForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/admission/plugins/{name}][%d] updateAdmissionPluginForbidden ", 403)
}

func (o *UpdateAdmissionPluginForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAdmissionPluginDefault creates a UpdateAdmissionPluginDefault with default headers values
func NewUpdateAdmissionPluginDefault(code int) *UpdateAdmissionPluginDefault {
	return &UpdateAdmissionPluginDefault{
		_statusCode: code,
	}
}

/*UpdateAdmissionPluginDefault handles this case with default header values.

errorResponse
*/
type UpdateAdmissionPluginDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update admission plugin default response
func (o *UpdateAdmissionPluginDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAdmissionPluginDefault) Error() string {
	return fmt.Sprintf("[PATCH /api/v1/admin/admission/plugins/{name}][%d] updateAdmissionPlugin default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAdmissionPluginDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UpdateAdmissionPluginDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
