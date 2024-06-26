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

// EdgeNodeConfigurationStopDebugEdgeNodeReader is a Reader for the EdgeNodeConfigurationStopDebugEdgeNode structure.
type EdgeNodeConfigurationStopDebugEdgeNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationStopDebugEdgeNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationStopDebugEdgeNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationStopDebugEdgeNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationStopDebugEdgeNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationStopDebugEdgeNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNodeConfigurationStopDebugEdgeNodeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationStopDebugEdgeNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationStopDebugEdgeNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationStopDebugEdgeNodeOK creates a EdgeNodeConfigurationStopDebugEdgeNodeOK with default headers values
func NewEdgeNodeConfigurationStopDebugEdgeNodeOK() *EdgeNodeConfigurationStopDebugEdgeNodeOK {
	return &EdgeNodeConfigurationStopDebugEdgeNodeOK{}
}

/*
EdgeNodeConfigurationStopDebugEdgeNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationStopDebugEdgeNodeOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration stop debug edge node o k response has a 2xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration stop debug edge node o k response has a 3xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration stop debug edge node o k response has a 4xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration stop debug edge node o k response has a 5xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration stop debug edge node o k response a status code equal to that given
func (o *EdgeNodeConfigurationStopDebugEdgeNodeOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/disable][%d] edgeNodeConfigurationStopDebugEdgeNodeOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeOK) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/disable][%d] edgeNodeConfigurationStopDebugEdgeNodeOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStopDebugEdgeNodeUnauthorized creates a EdgeNodeConfigurationStopDebugEdgeNodeUnauthorized with default headers values
func NewEdgeNodeConfigurationStopDebugEdgeNodeUnauthorized() *EdgeNodeConfigurationStopDebugEdgeNodeUnauthorized {
	return &EdgeNodeConfigurationStopDebugEdgeNodeUnauthorized{}
}

/*
EdgeNodeConfigurationStopDebugEdgeNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationStopDebugEdgeNodeUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration stop debug edge node unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration stop debug edge node unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration stop debug edge node unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration stop debug edge node unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration stop debug edge node unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationStopDebugEdgeNodeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/disable][%d] edgeNodeConfigurationStopDebugEdgeNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/disable][%d] edgeNodeConfigurationStopDebugEdgeNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStopDebugEdgeNodeForbidden creates a EdgeNodeConfigurationStopDebugEdgeNodeForbidden with default headers values
func NewEdgeNodeConfigurationStopDebugEdgeNodeForbidden() *EdgeNodeConfigurationStopDebugEdgeNodeForbidden {
	return &EdgeNodeConfigurationStopDebugEdgeNodeForbidden{}
}

/*
EdgeNodeConfigurationStopDebugEdgeNodeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationStopDebugEdgeNodeForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration stop debug edge node forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration stop debug edge node forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration stop debug edge node forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration stop debug edge node forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration stop debug edge node forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationStopDebugEdgeNodeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/disable][%d] edgeNodeConfigurationStopDebugEdgeNodeForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/disable][%d] edgeNodeConfigurationStopDebugEdgeNodeForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStopDebugEdgeNodeNotFound creates a EdgeNodeConfigurationStopDebugEdgeNodeNotFound with default headers values
func NewEdgeNodeConfigurationStopDebugEdgeNodeNotFound() *EdgeNodeConfigurationStopDebugEdgeNodeNotFound {
	return &EdgeNodeConfigurationStopDebugEdgeNodeNotFound{}
}

/*
EdgeNodeConfigurationStopDebugEdgeNodeNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationStopDebugEdgeNodeNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration stop debug edge node not found response has a 2xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration stop debug edge node not found response has a 3xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration stop debug edge node not found response has a 4xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration stop debug edge node not found response has a 5xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration stop debug edge node not found response a status code equal to that given
func (o *EdgeNodeConfigurationStopDebugEdgeNodeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/disable][%d] edgeNodeConfigurationStopDebugEdgeNodeNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/disable][%d] edgeNodeConfigurationStopDebugEdgeNodeNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStopDebugEdgeNodeConflict creates a EdgeNodeConfigurationStopDebugEdgeNodeConflict with default headers values
func NewEdgeNodeConfigurationStopDebugEdgeNodeConflict() *EdgeNodeConfigurationStopDebugEdgeNodeConflict {
	return &EdgeNodeConfigurationStopDebugEdgeNodeConflict{}
}

/*
EdgeNodeConfigurationStopDebugEdgeNodeConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge node record.
*/
type EdgeNodeConfigurationStopDebugEdgeNodeConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration stop debug edge node conflict response has a 2xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration stop debug edge node conflict response has a 3xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration stop debug edge node conflict response has a 4xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration stop debug edge node conflict response has a 5xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration stop debug edge node conflict response a status code equal to that given
func (o *EdgeNodeConfigurationStopDebugEdgeNodeConflict) IsCode(code int) bool {
	return code == 409
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/disable][%d] edgeNodeConfigurationStopDebugEdgeNodeConflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeConflict) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/disable][%d] edgeNodeConfigurationStopDebugEdgeNodeConflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStopDebugEdgeNodeInternalServerError creates a EdgeNodeConfigurationStopDebugEdgeNodeInternalServerError with default headers values
func NewEdgeNodeConfigurationStopDebugEdgeNodeInternalServerError() *EdgeNodeConfigurationStopDebugEdgeNodeInternalServerError {
	return &EdgeNodeConfigurationStopDebugEdgeNodeInternalServerError{}
}

/*
EdgeNodeConfigurationStopDebugEdgeNodeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationStopDebugEdgeNodeInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration stop debug edge node internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration stop debug edge node internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration stop debug edge node internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration stop debug edge node internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration stop debug edge node internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationStopDebugEdgeNodeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/disable][%d] edgeNodeConfigurationStopDebugEdgeNodeInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/disable][%d] edgeNodeConfigurationStopDebugEdgeNodeInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout creates a EdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout with default headers values
func NewEdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout() *EdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout {
	return &EdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout{}
}

/*
EdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration stop debug edge node gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration stop debug edge node gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration stop debug edge node gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration stop debug edge node gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration stop debug edge node gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/disable][%d] edgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/disable][%d] edgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStopDebugEdgeNodeDefault creates a EdgeNodeConfigurationStopDebugEdgeNodeDefault with default headers values
func NewEdgeNodeConfigurationStopDebugEdgeNodeDefault(code int) *EdgeNodeConfigurationStopDebugEdgeNodeDefault {
	return &EdgeNodeConfigurationStopDebugEdgeNodeDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationStopDebugEdgeNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationStopDebugEdgeNodeDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration stop debug edge node default response
func (o *EdgeNodeConfigurationStopDebugEdgeNodeDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node configuration stop debug edge node default response has a 2xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration stop debug edge node default response has a 3xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration stop debug edge node default response has a 4xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration stop debug edge node default response has a 5xx status code
func (o *EdgeNodeConfigurationStopDebugEdgeNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration stop debug edge node default response a status code equal to that given
func (o *EdgeNodeConfigurationStopDebugEdgeNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/disable][%d] EdgeNodeConfiguration_StopDebugEdgeNode default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeDefault) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/disable][%d] EdgeNodeConfiguration_StopDebugEdgeNode default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationStopDebugEdgeNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
