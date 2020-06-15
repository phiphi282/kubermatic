// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// CreateNodeDeploymentRequestReader is a Reader for the CreateNodeDeploymentRequest structure.
type CreateNodeDeploymentRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNodeDeploymentRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNodeDeploymentRequestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateNodeDeploymentRequestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateNodeDeploymentRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateNodeDeploymentRequestDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateNodeDeploymentRequestCreated creates a CreateNodeDeploymentRequestCreated with default headers values
func NewCreateNodeDeploymentRequestCreated() *CreateNodeDeploymentRequestCreated {
	return &CreateNodeDeploymentRequestCreated{}
}

/*CreateNodeDeploymentRequestCreated handles this case with default header values.

NodeDeploymentRequest
*/
type CreateNodeDeploymentRequestCreated struct {
	Payload *models.NodeDeploymentRequest
}

func (o *CreateNodeDeploymentRequestCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests][%d] createNodeDeploymentRequestCreated  %+v", 201, o.Payload)
}

func (o *CreateNodeDeploymentRequestCreated) GetPayload() *models.NodeDeploymentRequest {
	return o.Payload
}

func (o *CreateNodeDeploymentRequestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeDeploymentRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNodeDeploymentRequestUnauthorized creates a CreateNodeDeploymentRequestUnauthorized with default headers values
func NewCreateNodeDeploymentRequestUnauthorized() *CreateNodeDeploymentRequestUnauthorized {
	return &CreateNodeDeploymentRequestUnauthorized{}
}

/*CreateNodeDeploymentRequestUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type CreateNodeDeploymentRequestUnauthorized struct {
}

func (o *CreateNodeDeploymentRequestUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests][%d] createNodeDeploymentRequestUnauthorized ", 401)
}

func (o *CreateNodeDeploymentRequestUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateNodeDeploymentRequestForbidden creates a CreateNodeDeploymentRequestForbidden with default headers values
func NewCreateNodeDeploymentRequestForbidden() *CreateNodeDeploymentRequestForbidden {
	return &CreateNodeDeploymentRequestForbidden{}
}

/*CreateNodeDeploymentRequestForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type CreateNodeDeploymentRequestForbidden struct {
}

func (o *CreateNodeDeploymentRequestForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests][%d] createNodeDeploymentRequestForbidden ", 403)
}

func (o *CreateNodeDeploymentRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateNodeDeploymentRequestDefault creates a CreateNodeDeploymentRequestDefault with default headers values
func NewCreateNodeDeploymentRequestDefault(code int) *CreateNodeDeploymentRequestDefault {
	return &CreateNodeDeploymentRequestDefault{
		_statusCode: code,
	}
}

/*CreateNodeDeploymentRequestDefault handles this case with default header values.

errorResponse
*/
type CreateNodeDeploymentRequestDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the create node deployment request default response
func (o *CreateNodeDeploymentRequestDefault) Code() int {
	return o._statusCode
}

func (o *CreateNodeDeploymentRequestDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/projects/{project_id}/dc/{dc}/clusters/{cluster_id}/ndrequests][%d] createNodeDeploymentRequest default  %+v", o._statusCode, o.Payload)
}

func (o *CreateNodeDeploymentRequestDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateNodeDeploymentRequestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
