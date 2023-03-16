// Code generated by go-swagger; DO NOT EDIT.

package edge_node_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// EdgeNodeConfigurationGetEdgeNodeReader is a Reader for the EdgeNodeConfigurationGetEdgeNode structure.
type EdgeNodeConfigurationGetEdgeNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationGetEdgeNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationGetEdgeNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationGetEdgeNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationGetEdgeNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationGetEdgeNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationGetEdgeNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationGetEdgeNodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationGetEdgeNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationGetEdgeNodeOK creates a EdgeNodeConfigurationGetEdgeNodeOK with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeOK() *EdgeNodeConfigurationGetEdgeNodeOK {
	return &EdgeNodeConfigurationGetEdgeNodeOK{}
}

/*
EdgeNodeConfigurationGetEdgeNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationGetEdgeNodeOK struct {
	Payload *models.Node
}

// IsSuccess returns true when this edge node configuration get edge node o k response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration get edge node o k response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node o k response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration get edge node o k response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get edge node o k response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeConfigurationGetEdgeNodeOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] edgeNodeConfigurationGetEdgeNodeOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] edgeNodeConfigurationGetEdgeNodeOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeOK) GetPayload() *models.Node {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeUnauthorized creates a EdgeNodeConfigurationGetEdgeNodeUnauthorized with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeUnauthorized() *EdgeNodeConfigurationGetEdgeNodeUnauthorized {
	return &EdgeNodeConfigurationGetEdgeNodeUnauthorized{}
}

/*
EdgeNodeConfigurationGetEdgeNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationGetEdgeNodeUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration get edge node unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get edge node unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeConfigurationGetEdgeNodeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] edgeNodeConfigurationGetEdgeNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] edgeNodeConfigurationGetEdgeNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeForbidden creates a EdgeNodeConfigurationGetEdgeNodeForbidden with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeForbidden() *EdgeNodeConfigurationGetEdgeNodeForbidden {
	return &EdgeNodeConfigurationGetEdgeNodeForbidden{}
}

/*
EdgeNodeConfigurationGetEdgeNodeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationGetEdgeNodeForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration get edge node forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get edge node forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeConfigurationGetEdgeNodeForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] edgeNodeConfigurationGetEdgeNodeForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] edgeNodeConfigurationGetEdgeNodeForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeNotFound creates a EdgeNodeConfigurationGetEdgeNodeNotFound with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeNotFound() *EdgeNodeNotFound {
	return &EdgeNodeNotFound{}
}

/*
EdgeNodeNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node not found response has a 2xx status code
func (o *EdgeNodeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node not found response has a 3xx status code
func (o *EdgeNodeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node not found response has a 4xx status code
func (o *EdgeNodeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration get edge node not found response has a 5xx status code
func (o *EdgeNodeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get edge node not found response a status code equal to that given
func (o *EdgeNodeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] edgeNodeConfigurationGetEdgeNodeNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeNotFound) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] edgeNodeConfigurationGetEdgeNodeNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeInternalServerError creates a EdgeNodeConfigurationGetEdgeNodeInternalServerError with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeInternalServerError() *EdgeNodeConfigurationGetEdgeNodeInternalServerError {
	return &EdgeNodeConfigurationGetEdgeNodeInternalServerError{}
}

/*
EdgeNodeConfigurationGetEdgeNodeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationGetEdgeNodeInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration get edge node internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration get edge node internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeConfigurationGetEdgeNodeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] edgeNodeConfigurationGetEdgeNodeInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] edgeNodeConfigurationGetEdgeNodeInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeGatewayTimeout creates a EdgeNodeConfigurationGetEdgeNodeGatewayTimeout with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeGatewayTimeout() *EdgeNodeConfigurationGetEdgeNodeGatewayTimeout {
	return &EdgeNodeConfigurationGetEdgeNodeGatewayTimeout{}
}

/*
EdgeNodeConfigurationGetEdgeNodeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationGetEdgeNodeGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get edge node gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get edge node gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get edge node gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration get edge node gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration get edge node gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeConfigurationGetEdgeNodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] edgeNodeConfigurationGetEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] edgeNodeConfigurationGetEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetEdgeNodeDefault creates a EdgeNodeConfigurationGetEdgeNodeDefault with default headers values
func NewEdgeNodeConfigurationGetEdgeNodeDefault(code int) *EdgeNodeConfigurationGetEdgeNodeDefault {
	return &EdgeNodeConfigurationGetEdgeNodeDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationGetEdgeNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationGetEdgeNodeDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration get edge node default response
func (o *EdgeNodeConfigurationGetEdgeNodeDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node configuration get edge node default response has a 2xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration get edge node default response has a 3xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration get edge node default response has a 4xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration get edge node default response has a 5xx status code
func (o *EdgeNodeConfigurationGetEdgeNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration get edge node default response a status code equal to that given
func (o *EdgeNodeConfigurationGetEdgeNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeConfigurationGetEdgeNodeDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] EdgeNodeConfiguration_GetEdgeNode default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeDefault) String() string {
	return fmt.Sprintf("[GET /v1/devices/id/{id}][%d] EdgeNodeConfiguration_GetEdgeNode default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationGetEdgeNodeDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetEdgeNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}