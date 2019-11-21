// Code generated by go-swagger; DO NOT EDIT.

package versions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// GetMetaKubeVersionsReader is a Reader for the GetMetaKubeVersions structure.
type GetMetaKubeVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetaKubeVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMetaKubeVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetMetaKubeVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMetaKubeVersionsOK creates a GetMetaKubeVersionsOK with default headers values
func NewGetMetaKubeVersionsOK() *GetMetaKubeVersionsOK {
	return &GetMetaKubeVersionsOK{}
}

/*GetMetaKubeVersionsOK handles this case with default header values.

MetaKubeVersions
*/
type GetMetaKubeVersionsOK struct {
	Payload *models.KubermaticVersions
}

func (o *GetMetaKubeVersionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/version][%d] getMetaKubeVersionsOK  %+v", 200, o.Payload)
}

func (o *GetMetaKubeVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KubermaticVersions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetaKubeVersionsDefault creates a GetMetaKubeVersionsDefault with default headers values
func NewGetMetaKubeVersionsDefault(code int) *GetMetaKubeVersionsDefault {
	return &GetMetaKubeVersionsDefault{
		_statusCode: code,
	}
}

/*GetMetaKubeVersionsDefault handles this case with default header values.

errorResponse
*/
type GetMetaKubeVersionsDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get meta kube versions default response
func (o *GetMetaKubeVersionsDefault) Code() int {
	return o._statusCode
}

func (o *GetMetaKubeVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/version][%d] getMetaKubeVersions default  %+v", o._statusCode, o.Payload)
}

func (o *GetMetaKubeVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
