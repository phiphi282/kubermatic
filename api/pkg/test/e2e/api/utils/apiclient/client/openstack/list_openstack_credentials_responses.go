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

// ListOpenstackCredentialsReader is a Reader for the ListOpenstackCredentials structure.
type ListOpenstackCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListOpenstackCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListOpenstackCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListOpenstackCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListOpenstackCredentialsOK creates a ListOpenstackCredentialsOK with default headers values
func NewListOpenstackCredentialsOK() *ListOpenstackCredentialsOK {
	return &ListOpenstackCredentialsOK{}
}

/*ListOpenstackCredentialsOK handles this case with default header values.

CredentialList
*/
type ListOpenstackCredentialsOK struct {
	Payload *models.CredentialList
}

func (o *ListOpenstackCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/openstack/credentials][%d] listOpenstackCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListOpenstackCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListOpenstackCredentialsDefault creates a ListOpenstackCredentialsDefault with default headers values
func NewListOpenstackCredentialsDefault(code int) *ListOpenstackCredentialsDefault {
	return &ListOpenstackCredentialsDefault{
		_statusCode: code,
	}
}

/*ListOpenstackCredentialsDefault handles this case with default header values.

ErrorResponse is the default representation of an error
*/
type ListOpenstackCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorDetails
}

// Code gets the status code for the list openstack credentials default response
func (o *ListOpenstackCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListOpenstackCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/openstack/credentials][%d] listOpenstackCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListOpenstackCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
