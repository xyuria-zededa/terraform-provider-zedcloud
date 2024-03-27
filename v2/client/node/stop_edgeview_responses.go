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

// EdgeNodeConfigurationStopEdgeviewEdgeNodeReader is a Reader for the EdgeNodeConfigurationStopEdgeviewEdgeNode structure.
type EdgeNodeConfigurationStopEdgeviewEdgeNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationStopEdgeviewEdgeNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNodeConfigurationStopEdgeviewEdgeNodeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationStopEdgeviewEdgeNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationStopEdgeviewEdgeNodeOK creates a EdgeNodeConfigurationStopEdgeviewEdgeNodeOK with default headers values
func NewEdgeNodeConfigurationStopEdgeviewEdgeNodeOK() *EdgeNodeConfigurationStopEdgeviewEdgeNodeOK {
	return &EdgeNodeConfigurationStopEdgeviewEdgeNodeOK{}
}

/*
EdgeNodeConfigurationStopEdgeviewEdgeNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationStopEdgeviewEdgeNodeOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration stop edgeview edge node o k response has a 2xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration stop edgeview edge node o k response has a 3xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration stop edgeview edge node o k response has a 4xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration stop edgeview edge node o k response has a 5xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration stop edgeview edge node o k response a status code equal to that given
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/disable][%d] edgeNodeConfigurationStopEdgeviewEdgeNodeOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeOK) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/disable][%d] edgeNodeConfigurationStopEdgeviewEdgeNodeOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized creates a EdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized with default headers values
func NewEdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized() *EdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized {
	return &EdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized{}
}

/*
EdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration stop edgeview edge node unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration stop edgeview edge node unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration stop edgeview edge node unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration stop edgeview edge node unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration stop edgeview edge node unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/disable][%d] edgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/disable][%d] edgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden creates a EdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden with default headers values
func NewEdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden() *EdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden {
	return &EdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden{}
}

/*
EdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration stop edgeview edge node forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration stop edgeview edge node forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration stop edgeview edge node forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration stop edgeview edge node forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration stop edgeview edge node forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/disable][%d] edgeNodeConfigurationStopEdgeviewEdgeNodeForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/disable][%d] edgeNodeConfigurationStopEdgeviewEdgeNodeForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound creates a EdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound with default headers values
func NewEdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound() *EdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound {
	return &EdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound{}
}

/*
EdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration stop edgeview edge node not found response has a 2xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration stop edgeview edge node not found response has a 3xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration stop edgeview edge node not found response has a 4xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration stop edgeview edge node not found response has a 5xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration stop edgeview edge node not found response a status code equal to that given
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/disable][%d] edgeNodeConfigurationStopEdgeviewEdgeNodeNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/disable][%d] edgeNodeConfigurationStopEdgeviewEdgeNodeNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStopEdgeviewEdgeNodeConflict creates a EdgeNodeConfigurationStopEdgeviewEdgeNodeConflict with default headers values
func NewEdgeNodeConfigurationStopEdgeviewEdgeNodeConflict() *EdgeNodeConfigurationStopEdgeviewEdgeNodeConflict {
	return &EdgeNodeConfigurationStopEdgeviewEdgeNodeConflict{}
}

/*
EdgeNodeConfigurationStopEdgeviewEdgeNodeConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge node record.
*/
type EdgeNodeConfigurationStopEdgeviewEdgeNodeConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration stop edgeview edge node conflict response has a 2xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration stop edgeview edge node conflict response has a 3xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration stop edgeview edge node conflict response has a 4xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration stop edgeview edge node conflict response has a 5xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration stop edgeview edge node conflict response a status code equal to that given
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeConflict) IsCode(code int) bool {
	return code == 409
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/disable][%d] edgeNodeConfigurationStopEdgeviewEdgeNodeConflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeConflict) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/disable][%d] edgeNodeConfigurationStopEdgeviewEdgeNodeConflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError creates a EdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError with default headers values
func NewEdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError() *EdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError {
	return &EdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError{}
}

/*
EdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration stop edgeview edge node internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration stop edgeview edge node internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration stop edgeview edge node internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration stop edgeview edge node internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration stop edgeview edge node internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/disable][%d] edgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/disable][%d] edgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout creates a EdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout with default headers values
func NewEdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout() *EdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout {
	return &EdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout{}
}

/*
EdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration stop edgeview edge node gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration stop edgeview edge node gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration stop edgeview edge node gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration stop edgeview edge node gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration stop edgeview edge node gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/disable][%d] edgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/disable][%d] edgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStopEdgeviewEdgeNodeDefault creates a EdgeNodeConfigurationStopEdgeviewEdgeNodeDefault with default headers values
func NewEdgeNodeConfigurationStopEdgeviewEdgeNodeDefault(code int) *EdgeNodeConfigurationStopEdgeviewEdgeNodeDefault {
	return &EdgeNodeConfigurationStopEdgeviewEdgeNodeDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationStopEdgeviewEdgeNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationStopEdgeviewEdgeNodeDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration stop edgeview edge node default response
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node configuration stop edgeview edge node default response has a 2xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration stop edgeview edge node default response has a 3xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration stop edgeview edge node default response has a 4xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration stop edgeview edge node default response has a 5xx status code
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration stop edgeview edge node default response a status code equal to that given
func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/disable][%d] EdgeNodeConfiguration_StopEdgeviewEdgeNode default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeDefault) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/edgeview/disable][%d] EdgeNodeConfiguration_StopEdgeviewEdgeNode default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationStopEdgeviewEdgeNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}