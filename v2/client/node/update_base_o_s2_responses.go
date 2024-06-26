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

// EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Reader is a Reader for the EdgeNodeConfigurationUpdateEdgeNodeBaseOS2 structure.
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK() *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s2 o k response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration update edge node base o s2 o k response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s2 o k response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration update edge node base o s2 o k response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration update edge node base o s2 o k response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/publish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS2OK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/publish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS2OK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized() *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s2 unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s2 unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s2 unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration update edge node base o s2 unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration update edge node base o s2 unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/publish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/publish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden() *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s2 forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s2 forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s2 forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration update edge node base o s2 forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration update edge node base o s2 forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/publish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/publish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound() *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s2 not found response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s2 not found response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s2 not found response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration update edge node base o s2 not found response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration update edge node base o s2 not found response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/publish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/publish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict() *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge node record.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s2 conflict response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s2 conflict response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s2 conflict response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration update edge node base o s2 conflict response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration update edge node base o s2 conflict response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict) IsCode(code int) bool {
	return code == 409
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/publish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/publish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError() *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s2 internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s2 internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s2 internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration update edge node base o s2 internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration update edge node base o s2 internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/publish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/publish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout() *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout{}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration update edge node base o s2 gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration update edge node base o s2 gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration update edge node base o s2 gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration update edge node base o s2 gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration update edge node base o s2 gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/publish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/publish][%d] edgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2GatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default creates a EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default with default headers values
func NewEdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default(code int) *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default {
	return &EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration update edge node base o s2 default response
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node configuration update edge node base o s2 default response has a 2xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration update edge node base o s2 default response has a 3xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration update edge node base o s2 default response has a 4xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration update edge node base o s2 default response has a 5xx status code
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration update edge node base o s2 default response a status code equal to that given
func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/publish][%d] EdgeNodeConfiguration_UpdateEdgeNodeBaseOS2 default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/publish][%d] EdgeNodeConfiguration_UpdateEdgeNodeBaseOS2 default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationUpdateEdgeNodeBaseOS2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
