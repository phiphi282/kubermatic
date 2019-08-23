// Code generated by go-swagger; DO NOT EDIT.

package openstack

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// ListOpenstackQuotaLimitsReader is a Reader for the ListOpenstackQuotaLimits structure.
type ListOpenstackQuotaLimitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOpenstackQuotaLimitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListOpenstackQuotaLimitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListOpenstackQuotaLimitsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOpenstackQuotaLimitsOK creates a ListOpenstackQuotaLimitsOK with default headers values
func NewListOpenstackQuotaLimitsOK() *ListOpenstackQuotaLimitsOK {
	return &ListOpenstackQuotaLimitsOK{}
}

/*ListOpenstackQuotaLimitsOK handles this case with default header values.

Quotas
*/
type ListOpenstackQuotaLimitsOK struct {
	Payload *models.Quotas
}

func (o *ListOpenstackQuotaLimitsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/openstack/quotalimits][%d] listOpenstackQuotaLimitsOK  %+v", 200, o.Payload)
}

func (o *ListOpenstackQuotaLimitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Quotas)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenstackQuotaLimitsDefault creates a ListOpenstackQuotaLimitsDefault with default headers values
func NewListOpenstackQuotaLimitsDefault(code int) *ListOpenstackQuotaLimitsDefault {
	return &ListOpenstackQuotaLimitsDefault{
		_statusCode: code,
	}
}

/*ListOpenstackQuotaLimitsDefault handles this case with default header values.

ErrorResponse is the default representation of an error
*/
type ListOpenstackQuotaLimitsDefault struct {
	_statusCode int

	Payload *models.ErrorDetails
}

// Code gets the status code for the list openstack quota limits default response
func (o *ListOpenstackQuotaLimitsDefault) Code() int {
	return o._statusCode
}

func (o *ListOpenstackQuotaLimitsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/openstack/quotalimits][%d] listOpenstackQuotaLimits default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenstackQuotaLimitsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
