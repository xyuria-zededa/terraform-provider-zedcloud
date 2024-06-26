package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xyuria-zededa/terraform-provider-zedcloud/v2/models"
)

// EdgeNodeConfigurationActivateEdgeNodeReader is a Reader for the EdgeNodeConfigurationActivateEdgeNode structure.
type EdgeNodeConfigurationActivateEdgeNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationActivateEdgeNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationActivateEdgeNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationActivateEdgeNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationActivateEdgeNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationActivateEdgeNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNodeConfigurationActivateEdgeNodeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationActivateEdgeNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationActivateEdgeNodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationActivateEdgeNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationActivateEdgeNodeOK creates a EdgeNodeConfigurationActivateEdgeNodeOK with default headers values
func NewEdgeNodeConfigurationActivateEdgeNodeOK() *EdgeNodeConfigurationActivateEdgeNodeOK {
	return &EdgeNodeConfigurationActivateEdgeNodeOK{}
}

/*
EdgeNodeConfigurationActivateEdgeNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationActivateEdgeNodeOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration activate edge node o k response has a 2xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration activate edge node o k response has a 3xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration activate edge node o k response has a 4xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration activate edge node o k response has a 5xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration activate edge node o k response a status code equal to that given
func (o *EdgeNodeConfigurationActivateEdgeNodeOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeConfigurationActivateEdgeNodeOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] edgeNodeConfigurationActivateEdgeNodeOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationActivateEdgeNodeOK) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] edgeNodeConfigurationActivateEdgeNodeOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationActivateEdgeNodeOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationActivateEdgeNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationActivateEdgeNodeUnauthorized creates a EdgeNodeConfigurationActivateEdgeNodeUnauthorized with default headers values
func NewEdgeNodeConfigurationActivateEdgeNodeUnauthorized() *EdgeNodeConfigurationActivateEdgeNodeUnauthorized {
	return &EdgeNodeConfigurationActivateEdgeNodeUnauthorized{}
}

/*
EdgeNodeConfigurationActivateEdgeNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationActivateEdgeNodeUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration activate edge node unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration activate edge node unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration activate edge node unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration activate edge node unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration activate edge node unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationActivateEdgeNodeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeConfigurationActivateEdgeNodeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] edgeNodeConfigurationActivateEdgeNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationActivateEdgeNodeUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] edgeNodeConfigurationActivateEdgeNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationActivateEdgeNodeUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationActivateEdgeNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationActivateEdgeNodeForbidden creates a EdgeNodeConfigurationActivateEdgeNodeForbidden with default headers values
func NewEdgeNodeConfigurationActivateEdgeNodeForbidden() *EdgeNodeConfigurationActivateEdgeNodeForbidden {
	return &EdgeNodeConfigurationActivateEdgeNodeForbidden{}
}

/*
EdgeNodeConfigurationActivateEdgeNodeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationActivateEdgeNodeForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration activate edge node forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration activate edge node forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration activate edge node forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration activate edge node forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration activate edge node forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationActivateEdgeNodeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeConfigurationActivateEdgeNodeForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] edgeNodeConfigurationActivateEdgeNodeForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationActivateEdgeNodeForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] edgeNodeConfigurationActivateEdgeNodeForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationActivateEdgeNodeForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationActivateEdgeNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationActivateEdgeNodeNotFound creates a EdgeNodeConfigurationActivateEdgeNodeNotFound with default headers values
func NewEdgeNodeConfigurationActivateEdgeNodeNotFound() *EdgeNodeConfigurationActivateEdgeNodeNotFound {
	return &EdgeNodeConfigurationActivateEdgeNodeNotFound{}
}

/*
EdgeNodeConfigurationActivateEdgeNodeNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationActivateEdgeNodeNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration activate edge node not found response has a 2xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration activate edge node not found response has a 3xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration activate edge node not found response has a 4xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration activate edge node not found response has a 5xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration activate edge node not found response a status code equal to that given
func (o *EdgeNodeConfigurationActivateEdgeNodeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeConfigurationActivateEdgeNodeNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] edgeNodeConfigurationActivateEdgeNodeNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationActivateEdgeNodeNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] edgeNodeConfigurationActivateEdgeNodeNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationActivateEdgeNodeNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationActivateEdgeNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationActivateEdgeNodeConflict creates a EdgeNodeConfigurationActivateEdgeNodeConflict with default headers values
func NewEdgeNodeConfigurationActivateEdgeNodeConflict() *EdgeNodeConfigurationActivateEdgeNodeConflict {
	return &EdgeNodeConfigurationActivateEdgeNodeConflict{}
}

/*
EdgeNodeConfigurationActivateEdgeNodeConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge node record.
*/
type EdgeNodeConfigurationActivateEdgeNodeConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration activate edge node conflict response has a 2xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration activate edge node conflict response has a 3xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration activate edge node conflict response has a 4xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration activate edge node conflict response has a 5xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration activate edge node conflict response a status code equal to that given
func (o *EdgeNodeConfigurationActivateEdgeNodeConflict) IsCode(code int) bool {
	return code == 409
}

func (o *EdgeNodeConfigurationActivateEdgeNodeConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] edgeNodeConfigurationActivateEdgeNodeConflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationActivateEdgeNodeConflict) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] edgeNodeConfigurationActivateEdgeNodeConflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationActivateEdgeNodeConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationActivateEdgeNodeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationActivateEdgeNodeInternalServerError creates a EdgeNodeConfigurationActivateEdgeNodeInternalServerError with default headers values
func NewEdgeNodeConfigurationActivateEdgeNodeInternalServerError() *EdgeNodeConfigurationActivateEdgeNodeInternalServerError {
	return &EdgeNodeConfigurationActivateEdgeNodeInternalServerError{}
}

/*
EdgeNodeConfigurationActivateEdgeNodeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationActivateEdgeNodeInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration activate edge node internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration activate edge node internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration activate edge node internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration activate edge node internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration activate edge node internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationActivateEdgeNodeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeConfigurationActivateEdgeNodeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] edgeNodeConfigurationActivateEdgeNodeInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationActivateEdgeNodeInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] edgeNodeConfigurationActivateEdgeNodeInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationActivateEdgeNodeInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationActivateEdgeNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationActivateEdgeNodeGatewayTimeout creates a EdgeNodeConfigurationActivateEdgeNodeGatewayTimeout with default headers values
func NewEdgeNodeConfigurationActivateEdgeNodeGatewayTimeout() *EdgeNodeConfigurationActivateEdgeNodeGatewayTimeout {
	return &EdgeNodeConfigurationActivateEdgeNodeGatewayTimeout{}
}

/*
EdgeNodeConfigurationActivateEdgeNodeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationActivateEdgeNodeGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration activate edge node gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration activate edge node gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration activate edge node gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration activate edge node gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration activate edge node gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationActivateEdgeNodeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeConfigurationActivateEdgeNodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] edgeNodeConfigurationActivateEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationActivateEdgeNodeGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] edgeNodeConfigurationActivateEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationActivateEdgeNodeGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationActivateEdgeNodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationActivateEdgeNodeDefault creates a EdgeNodeConfigurationActivateEdgeNodeDefault with default headers values
func NewEdgeNodeConfigurationActivateEdgeNodeDefault(code int) *EdgeNodeConfigurationActivateEdgeNodeDefault {
	return &EdgeNodeConfigurationActivateEdgeNodeDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationActivateEdgeNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationActivateEdgeNodeDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration activate edge node default response
func (o *EdgeNodeConfigurationActivateEdgeNodeDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node configuration activate edge node default response has a 2xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration activate edge node default response has a 3xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration activate edge node default response has a 4xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration activate edge node default response has a 5xx status code
func (o *EdgeNodeConfigurationActivateEdgeNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration activate edge node default response a status code equal to that given
func (o *EdgeNodeConfigurationActivateEdgeNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeConfigurationActivateEdgeNodeDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] EdgeNodeConfiguration_ActivateEdgeNode default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationActivateEdgeNodeDefault) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/activate][%d] EdgeNodeConfiguration_ActivateEdgeNode default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationActivateEdgeNodeDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationActivateEdgeNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
