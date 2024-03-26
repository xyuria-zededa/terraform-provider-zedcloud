package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/xyuria-zededa/terraform-provider-zedcloud/v2/models"
)

// EdgeNodeConfigurationStartDebugEdgeNodeReader is a Reader for the EdgeNodeConfigurationStartDebugEdgeNode structure.
type EdgeNodeConfigurationStartDebugEdgeNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationStartDebugEdgeNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationStartDebugEdgeNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationStartDebugEdgeNodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationStartDebugEdgeNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationStartDebugEdgeNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeNodeConfigurationStartDebugEdgeNodeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationStartDebugEdgeNodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationStartDebugEdgeNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationStartDebugEdgeNodeOK creates a EdgeNodeConfigurationStartDebugEdgeNodeOK with default headers values
func NewEdgeNodeConfigurationStartDebugEdgeNodeOK() *EdgeNodeConfigurationStartDebugEdgeNodeOK {
	return &EdgeNodeConfigurationStartDebugEdgeNodeOK{}
}

/*
EdgeNodeConfigurationStartDebugEdgeNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationStartDebugEdgeNodeOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration start debug edge node o k response has a 2xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration start debug edge node o k response has a 3xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration start debug edge node o k response has a 4xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration start debug edge node o k response has a 5xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration start debug edge node o k response a status code equal to that given
func (o *EdgeNodeConfigurationStartDebugEdgeNodeOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/enable][%d] edgeNodeConfigurationStartDebugEdgeNodeOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeOK) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/enable][%d] edgeNodeConfigurationStartDebugEdgeNodeOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStartDebugEdgeNodeUnauthorized creates a EdgeNodeConfigurationStartDebugEdgeNodeUnauthorized with default headers values
func NewEdgeNodeConfigurationStartDebugEdgeNodeUnauthorized() *EdgeNodeConfigurationStartDebugEdgeNodeUnauthorized {
	return &EdgeNodeConfigurationStartDebugEdgeNodeUnauthorized{}
}

/*
EdgeNodeConfigurationStartDebugEdgeNodeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationStartDebugEdgeNodeUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration start debug edge node unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration start debug edge node unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration start debug edge node unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration start debug edge node unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration start debug edge node unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationStartDebugEdgeNodeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/enable][%d] edgeNodeConfigurationStartDebugEdgeNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/enable][%d] edgeNodeConfigurationStartDebugEdgeNodeUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStartDebugEdgeNodeForbidden creates a EdgeNodeConfigurationStartDebugEdgeNodeForbidden with default headers values
func NewEdgeNodeConfigurationStartDebugEdgeNodeForbidden() *EdgeNodeConfigurationStartDebugEdgeNodeForbidden {
	return &EdgeNodeConfigurationStartDebugEdgeNodeForbidden{}
}

/*
EdgeNodeConfigurationStartDebugEdgeNodeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationStartDebugEdgeNodeForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration start debug edge node forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration start debug edge node forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration start debug edge node forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration start debug edge node forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration start debug edge node forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationStartDebugEdgeNodeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/enable][%d] edgeNodeConfigurationStartDebugEdgeNodeForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/enable][%d] edgeNodeConfigurationStartDebugEdgeNodeForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStartDebugEdgeNodeNotFound creates a EdgeNodeConfigurationStartDebugEdgeNodeNotFound with default headers values
func NewEdgeNodeConfigurationStartDebugEdgeNodeNotFound() *EdgeNodeConfigurationStartDebugEdgeNodeNotFound {
	return &EdgeNodeConfigurationStartDebugEdgeNodeNotFound{}
}

/*
EdgeNodeConfigurationStartDebugEdgeNodeNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationStartDebugEdgeNodeNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration start debug edge node not found response has a 2xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration start debug edge node not found response has a 3xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration start debug edge node not found response has a 4xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration start debug edge node not found response has a 5xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration start debug edge node not found response a status code equal to that given
func (o *EdgeNodeConfigurationStartDebugEdgeNodeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/enable][%d] edgeNodeConfigurationStartDebugEdgeNodeNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/enable][%d] edgeNodeConfigurationStartDebugEdgeNodeNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStartDebugEdgeNodeConflict creates a EdgeNodeConfigurationStartDebugEdgeNodeConflict with default headers values
func NewEdgeNodeConfigurationStartDebugEdgeNodeConflict() *EdgeNodeConfigurationStartDebugEdgeNodeConflict {
	return &EdgeNodeConfigurationStartDebugEdgeNodeConflict{}
}

/*
EdgeNodeConfigurationStartDebugEdgeNodeConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing edge node record.
*/
type EdgeNodeConfigurationStartDebugEdgeNodeConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration start debug edge node conflict response has a 2xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration start debug edge node conflict response has a 3xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration start debug edge node conflict response has a 4xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration start debug edge node conflict response has a 5xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration start debug edge node conflict response a status code equal to that given
func (o *EdgeNodeConfigurationStartDebugEdgeNodeConflict) IsCode(code int) bool {
	return code == 409
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/enable][%d] edgeNodeConfigurationStartDebugEdgeNodeConflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeConflict) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/enable][%d] edgeNodeConfigurationStartDebugEdgeNodeConflict  %+v", 409, o.Payload)
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStartDebugEdgeNodeInternalServerError creates a EdgeNodeConfigurationStartDebugEdgeNodeInternalServerError with default headers values
func NewEdgeNodeConfigurationStartDebugEdgeNodeInternalServerError() *EdgeNodeConfigurationStartDebugEdgeNodeInternalServerError {
	return &EdgeNodeConfigurationStartDebugEdgeNodeInternalServerError{}
}

/*
EdgeNodeConfigurationStartDebugEdgeNodeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationStartDebugEdgeNodeInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration start debug edge node internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration start debug edge node internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration start debug edge node internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration start debug edge node internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration start debug edge node internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationStartDebugEdgeNodeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/enable][%d] edgeNodeConfigurationStartDebugEdgeNodeInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/enable][%d] edgeNodeConfigurationStartDebugEdgeNodeInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout creates a EdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout with default headers values
func NewEdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout() *EdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout {
	return &EdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout{}
}

/*
EdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration start debug edge node gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration start debug edge node gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration start debug edge node gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration start debug edge node gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration start debug edge node gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/enable][%d] edgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/enable][%d] edgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationStartDebugEdgeNodeDefault creates a EdgeNodeConfigurationStartDebugEdgeNodeDefault with default headers values
func NewEdgeNodeConfigurationStartDebugEdgeNodeDefault(code int) *EdgeNodeConfigurationStartDebugEdgeNodeDefault {
	return &EdgeNodeConfigurationStartDebugEdgeNodeDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationStartDebugEdgeNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationStartDebugEdgeNodeDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration start debug edge node default response
func (o *EdgeNodeConfigurationStartDebugEdgeNodeDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node configuration start debug edge node default response has a 2xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration start debug edge node default response has a 3xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration start debug edge node default response has a 4xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration start debug edge node default response has a 5xx status code
func (o *EdgeNodeConfigurationStartDebugEdgeNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration start debug edge node default response a status code equal to that given
func (o *EdgeNodeConfigurationStartDebugEdgeNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/enable][%d] EdgeNodeConfiguration_StartDebugEdgeNode default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeDefault) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/debug/enable][%d] EdgeNodeConfiguration_StartDebugEdgeNode default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationStartDebugEdgeNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
EdgeNodeConfigurationStartDebugEdgeNodeBody Device debug knob configuration
//
// Device debug knob configuration request payload holds the device debug mode properties
swagger:model EdgeNodeConfigurationStartDebugEdgeNodeBody
*/
type EdgeNodeConfigurationStartDebugEdgeNodeBody struct {

	// debug knob flag
	DebugKnob bool `json:"debugKnob,omitempty"`

	// debug knob expiry status flag
	Expired bool `json:"expired,omitempty"`

	// debug expiry time in minutes
	Expiry string `json:"expiry,omitempty"`
}

// Validate validates this edge node configuration start debug edge node body
func (o *EdgeNodeConfigurationStartDebugEdgeNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this edge node configuration start debug edge node body based on context it is used
func (o *EdgeNodeConfigurationStartDebugEdgeNodeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *EdgeNodeConfigurationStartDebugEdgeNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *EdgeNodeConfigurationStartDebugEdgeNodeBody) UnmarshalBinary(b []byte) error {
	var res EdgeNodeConfigurationStartDebugEdgeNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
