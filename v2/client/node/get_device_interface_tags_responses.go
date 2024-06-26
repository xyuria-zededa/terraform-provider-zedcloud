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

// EdgeNodeConfigurationGetDeviceInterfaceTagsReader is a Reader for the EdgeNodeConfigurationGetDeviceInterfaceTags structure.
type EdgeNodeConfigurationGetDeviceInterfaceTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeNodeConfigurationGetDeviceInterfaceTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeNodeConfigurationGetDeviceInterfaceTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeNodeConfigurationGetDeviceInterfaceTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeNodeConfigurationGetDeviceInterfaceTagsOK creates a EdgeNodeConfigurationGetDeviceInterfaceTagsOK with default headers values
func NewEdgeNodeConfigurationGetDeviceInterfaceTagsOK() *EdgeNodeConfigurationGetDeviceInterfaceTagsOK {
	return &EdgeNodeConfigurationGetDeviceInterfaceTagsOK{}
}

/*
EdgeNodeConfigurationGetDeviceInterfaceTagsOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeNodeConfigurationGetDeviceInterfaceTagsOK struct {
	Payload *models.ObjectTagsList
}

// IsSuccess returns true when this edge node configuration get device interface tags o k response has a 2xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge node configuration get device interface tags o k response has a 3xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get device interface tags o k response has a 4xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration get device interface tags o k response has a 5xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get device interface tags o k response a status code equal to that given
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsOK) IsCode(code int) bool {
	return code == 200
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/interfaces/tags][%d] edgeNodeConfigurationGetDeviceInterfaceTagsOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/interfaces/tags][%d] edgeNodeConfigurationGetDeviceInterfaceTagsOK  %+v", 200, o.Payload)
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsOK) GetPayload() *models.ObjectTagsList {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectTagsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest creates a EdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest with default headers values
func NewEdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest() *EdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest {
	return &EdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest{}
}

/*
EdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type EdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get device interface tags bad request response has a 2xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get device interface tags bad request response has a 3xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get device interface tags bad request response has a 4xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration get device interface tags bad request response has a 5xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get device interface tags bad request response a status code equal to that given
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/devices/interfaces/tags][%d] edgeNodeConfigurationGetDeviceInterfaceTagsBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/devices/interfaces/tags][%d] edgeNodeConfigurationGetDeviceInterfaceTagsBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized creates a EdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized with default headers values
func NewEdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized() *EdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized {
	return &EdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized{}
}

/*
EdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get device interface tags unauthorized response has a 2xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get device interface tags unauthorized response has a 3xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get device interface tags unauthorized response has a 4xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration get device interface tags unauthorized response has a 5xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get device interface tags unauthorized response a status code equal to that given
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/interfaces/tags][%d] edgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/interfaces/tags][%d] edgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetDeviceInterfaceTagsForbidden creates a EdgeNodeConfigurationGetDeviceInterfaceTagsForbidden with default headers values
func NewEdgeNodeConfigurationGetDeviceInterfaceTagsForbidden() *EdgeNodeConfigurationGetDeviceInterfaceTagsForbidden {
	return &EdgeNodeConfigurationGetDeviceInterfaceTagsForbidden{}
}

/*
EdgeNodeConfigurationGetDeviceInterfaceTagsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type EdgeNodeConfigurationGetDeviceInterfaceTagsForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get device interface tags forbidden response has a 2xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get device interface tags forbidden response has a 3xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get device interface tags forbidden response has a 4xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge node configuration get device interface tags forbidden response has a 5xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge node configuration get device interface tags forbidden response a status code equal to that given
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/interfaces/tags][%d] edgeNodeConfigurationGetDeviceInterfaceTagsForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/interfaces/tags][%d] edgeNodeConfigurationGetDeviceInterfaceTagsForbidden  %+v", 403, o.Payload)
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError creates a EdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError with default headers values
func NewEdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError() *EdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError {
	return &EdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError{}
}

/*
EdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get device interface tags internal server error response has a 2xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get device interface tags internal server error response has a 3xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get device interface tags internal server error response has a 4xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration get device interface tags internal server error response has a 5xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration get device interface tags internal server error response a status code equal to that given
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/interfaces/tags][%d] edgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/interfaces/tags][%d] edgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout creates a EdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout with default headers values
func NewEdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout() *EdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout {
	return &EdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout{}
}

/*
EdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge node configuration get device interface tags gateway timeout response has a 2xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge node configuration get device interface tags gateway timeout response has a 3xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge node configuration get device interface tags gateway timeout response has a 4xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge node configuration get device interface tags gateway timeout response has a 5xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge node configuration get device interface tags gateway timeout response a status code equal to that given
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/interfaces/tags][%d] edgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/interfaces/tags][%d] edgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeNodeConfigurationGetDeviceInterfaceTagsDefault creates a EdgeNodeConfigurationGetDeviceInterfaceTagsDefault with default headers values
func NewEdgeNodeConfigurationGetDeviceInterfaceTagsDefault(code int) *EdgeNodeConfigurationGetDeviceInterfaceTagsDefault {
	return &EdgeNodeConfigurationGetDeviceInterfaceTagsDefault{
		_statusCode: code,
	}
}

/*
EdgeNodeConfigurationGetDeviceInterfaceTagsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeNodeConfigurationGetDeviceInterfaceTagsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the edge node configuration get device interface tags default response
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this edge node configuration get device interface tags default response has a 2xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge node configuration get device interface tags default response has a 3xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge node configuration get device interface tags default response has a 4xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge node configuration get device interface tags default response has a 5xx status code
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge node configuration get device interface tags default response a status code equal to that given
func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/interfaces/tags][%d] EdgeNodeConfiguration_GetDeviceInterfaceTags default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsDefault) String() string {
	return fmt.Sprintf("[GET /v1/devices/interfaces/tags][%d] EdgeNodeConfiguration_GetDeviceInterfaceTags default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeNodeConfigurationGetDeviceInterfaceTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
