// Code generated by go-swagger; DO NOT EDIT.

package gcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// ListGCPSizesNoCredentialsReader is a Reader for the ListGCPSizesNoCredentials structure.
type ListGCPSizesNoCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGCPSizesNoCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListGCPSizesNoCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListGCPSizesNoCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListGCPSizesNoCredentialsOK creates a ListGCPSizesNoCredentialsOK with default headers values
func NewListGCPSizesNoCredentialsOK() *ListGCPSizesNoCredentialsOK {
	return &ListGCPSizesNoCredentialsOK{}
}

/*ListGCPSizesNoCredentialsOK handles this case with default header values.

GCPMachineSizeList
*/
type ListGCPSizesNoCredentialsOK struct {
	Payload models.GCPMachineSizeList
}

func (o *ListGCPSizesNoCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/sizes][%d] listGCPSizesNoCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListGCPSizesNoCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGCPSizesNoCredentialsDefault creates a ListGCPSizesNoCredentialsDefault with default headers values
func NewListGCPSizesNoCredentialsDefault(code int) *ListGCPSizesNoCredentialsDefault {
	return &ListGCPSizesNoCredentialsDefault{
		_statusCode: code,
	}
}

/*ListGCPSizesNoCredentialsDefault handles this case with default header values.

ErrorResponse is the default representation of an error
*/
type ListGCPSizesNoCredentialsDefault struct {
	_statusCode int

	Payload *models.ErrorDetails
}

// Code gets the status code for the list g c p sizes no credentials default response
func (o *ListGCPSizesNoCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *ListGCPSizesNoCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/providers/gcp/sizes][%d] listGCPSizesNoCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *ListGCPSizesNoCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
