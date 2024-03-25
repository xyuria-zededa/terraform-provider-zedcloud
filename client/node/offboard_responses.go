package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xyuria-zededa/terraform-provider-zedcloud/models"
)

// EdgeNodeConfigurationOffboardReader is a Reader for the EdgeNodeConfigurationOffboard structure.
type EdgeNodeConfigurationOffboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationOffboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationOffboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeNodeConfigurationOffboardUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationOffboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeNodeConfigurationOffboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationOffboardInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationOffboardGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationOffboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationOffboardOK creates a EdgeNodeConfigurationOffboardOK with default headers values
func NewEdgeNodeConfigurationOffboardOK() *EdgeNodeConfigurationOffboardOK {
	return &EdgeNodeConfigurationOffboardOK{}
}

/*
EdgeNodeConfigurationOffboardOK describes a response with status code 200, with default header values.

Success. The API gateway offboarded the edge-node successfully..
*/
type EdgeNodeConfigurationOffboardOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration offboard o k response has a 2xx status code
func (o *EdgeNodeConfigurationOffboardOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration offboard o k response has a 3xx status code
func (o *EdgeNodeConfigurationOffboardOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration offboard o k response has a 4xx status code
func (o *EdgeNodeConfigurationOffboardOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration offboard o k response has a 5xx status code
func (o *EdgeNodeConfigurationOffboardOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration offboard o k response a status code equal to that given
func (o *EdgeNodeConfigurationOffboardOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeConfigurationOffboardOK) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationOffboardOK) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationOffboardOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationOffboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationOffboardUnauthorized creates a EdgeNodeConfigurationOffboardUnauthorized with default headers values
func NewEdgeNodeConfigurationOffboardUnauthorized() *EdgeNodeConfigurationOffboardUnauthorized {
	return &EdgeNodeConfigurationOffboardUnauthorized{}
}

/*
EdgeNodeConfigurationOffboardUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationOffboardUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration offboard unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationOffboardUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration offboard unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationOffboardUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration offboard unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationOffboardUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration offboard unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationOffboardUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration offboard unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationOffboardUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeConfigurationOffboardUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationOffboardUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationOffboardUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationOffboardUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationOffboardForbidden creates a EdgeNodeConfigurationOffboardForbidden with default headers values
func NewEdgeNodeConfigurationOffboardForbidden() *EdgeNodeConfigurationOffboardForbidden {
	return &EdgeNodeConfigurationOffboardForbidden{}
}

/*
EdgeNodeConfigurationOffboardForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationOffboardForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration offboard forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationOffboardForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration offboard forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationOffboardForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration offboard forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationOffboardForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration offboard forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationOffboardForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration offboard forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationOffboardForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeConfigurationOffboardForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationOffboardForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationOffboardForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationOffboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationOffboardNotFound creates a EdgeNodeConfigurationOffboardNotFound with default headers values
func NewEdgeNodeConfigurationOffboardNotFound() *EdgeNodeConfigurationOffboardNotFound {
	return &EdgeNodeConfigurationOffboardNotFound{}
}

/*
EdgeNodeConfigurationOffboardNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeNodeConfigurationOffboardNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration offboard not found response has a 2xx status code
func (o *EdgeNodeConfigurationOffboardNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration offboard not found response has a 3xx status code
func (o *EdgeNodeConfigurationOffboardNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration offboard not found response has a 4xx status code
func (o *EdgeNodeConfigurationOffboardNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration offboard not found response has a 5xx status code
func (o *EdgeNodeConfigurationOffboardNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration offboard not found response a status code equal to that given
func (o *EdgeNodeConfigurationOffboardNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *EdgeNodeConfigurationOffboardNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationOffboardNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardNotFound  %+v", 404, o.Payload)
}

func (o *EdgeNodeConfigurationOffboardNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationOffboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationOffboardInternalServerError creates a EdgeNodeConfigurationOffboardInternalServerError with default headers values
func NewEdgeNodeConfigurationOffboardInternalServerError() *EdgeNodeConfigurationOffboardInternalServerError {
	return &EdgeNodeConfigurationOffboardInternalServerError{}
}

/*
EdgeNodeConfigurationOffboardInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationOffboardInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration offboard internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationOffboardInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration offboard internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationOffboardInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration offboard internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationOffboardInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration offboard internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationOffboardInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration offboard internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationOffboardInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeConfigurationOffboardInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationOffboardInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationOffboardInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationOffboardInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationOffboardGatewayTimeout creates a EdgeNodeConfigurationOffboardGatewayTimeout with default headers values
func NewEdgeNodeConfigurationOffboardGatewayTimeout() *EdgeNodeConfigurationOffboardGatewayTimeout {
	return &EdgeNodeConfigurationOffboardGatewayTimeout{}
}

/*
EdgeNodeConfigurationOffboardGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationOffboardGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration offboard gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationOffboardGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration offboard gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationOffboardGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration offboard gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationOffboardGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration offboard gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationOffboardGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration offboard gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationOffboardGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeConfigurationOffboardGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationOffboardGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] edgeNodeConfigurationOffboardGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationOffboardGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationOffboardGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationOffboardDefault creates a EdgeNodeConfigurationOffboardDefault with default headers values
func NewEdgeNodeConfigurationOffboardDefault(code int) *EdgeNodeConfigurationOffboardDefault {
	return &EdgeNodeConfigurationOffboardDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationOffboardDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationOffboardDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration offboard default response
func (o *EdgeNodeConfigurationOffboardDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node configuration offboard default response has a 2xx status code
func (o *EdgeNodeConfigurationOffboardDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration offboard default response has a 3xx status code
func (o *EdgeNodeConfigurationOffboardDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration offboard default response has a 4xx status code
func (o *EdgeNodeConfigurationOffboardDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration offboard default response has a 5xx status code
func (o *EdgeNodeConfigurationOffboardDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration offboard default response a status code equal to that given
func (o *EdgeNodeConfigurationOffboardDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeConfigurationOffboardDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] EdgeNodeConfiguration_Offboard default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationOffboardDefault) String() string {
	return fmt.Sprintf("[PUT /v1/devices/id/{id}/Offboard][%d] EdgeNodeConfiguration_Offboard default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationOffboardDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationOffboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
