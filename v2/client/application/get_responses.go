package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xyuria-zededa/terraform-provider-zedcloud/models"
)

// EdgeApplicationConfigurationGetEdgeApplicationBundleReader is a Reader for the EdgeApplicationConfigurationGetEdgeApplicationBundle structure.
type EdgeApplicationConfigurationGetEdgeApplicationBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleOK creates a EdgeApplicationConfigurationGetEdgeApplicationBundleOK with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleOK() *EdgeApplicationConfigurationGetEdgeApplicationBundleOK {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleOK{}
}

/*
EdgeApplicationConfigurationGetEdgeApplicationBundleOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleOK struct {
	Payload *models.Application
}

// IsSuccess returns true when this edge application configuration get edge application bundle o k response has a 2xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application configuration get edge application bundle o k response has a 3xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get edge application bundle o k response has a 4xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration get edge application bundle o k response has a 5xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration get edge application bundle o k response a status code equal to that given
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge application configuration get edge application bundle o k response
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleOK) Code() int {
	return 200
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleOK) GetPayload() *models.Application {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized creates a EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized() *EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized{}
}

/*
EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get edge application bundle unauthorized response has a 2xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get edge application bundle unauthorized response has a 3xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get edge application bundle unauthorized response has a 4xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration get edge application bundle unauthorized response has a 5xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration get edge application bundle unauthorized response a status code equal to that given
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge application configuration get edge application bundle unauthorized response
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized) Code() int {
	return 401
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleForbidden creates a EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleForbidden() *EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden{}
}

/*
EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get edge application bundle forbidden response has a 2xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get edge application bundle forbidden response has a 3xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get edge application bundle forbidden response has a 4xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration get edge application bundle forbidden response has a 5xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration get edge application bundle forbidden response a status code equal to that given
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge application configuration get edge application bundle forbidden response
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden) Code() int {
	return 403
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleNotFound creates a EdgeApplicationConfigurationGetEdgeApplicationBundleNotFound with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleNotFound() *GetApplicationNotFound {
	return &GetApplicationNotFound{}
}

/*
GetApplicationNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type GetApplicationNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get edge application bundle not found response has a 2xx status code
func (o *GetApplicationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get edge application bundle not found response has a 3xx status code
func (o *GetApplicationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get edge application bundle not found response has a 4xx status code
func (o *GetApplicationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration get edge application bundle not found response has a 5xx status code
func (o *GetApplicationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration get edge application bundle not found response a status code equal to that given
func (o *GetApplicationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge application configuration get edge application bundle not found response
func (o *GetApplicationNotFound) Code() int {
	return 404
}

func (o *GetApplicationNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleNotFound  %+v", 404, o.Payload)
}

func (o *GetApplicationNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleNotFound  %+v", 404, o.Payload)
}

func (o *GetApplicationNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *GetApplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError creates a EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError() *EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError{}
}

/*
EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get edge application bundle internal server error response has a 2xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get edge application bundle internal server error response has a 3xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get edge application bundle internal server error response has a 4xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration get edge application bundle internal server error response has a 5xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application configuration get edge application bundle internal server error response a status code equal to that given
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge application configuration get edge application bundle internal server error response
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError) Code() int {
	return 500
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout creates a EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout() *EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout{}
}

/*
EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get edge application bundle gateway timeout response has a 2xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get edge application bundle gateway timeout response has a 3xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get edge application bundle gateway timeout response has a 4xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration get edge application bundle gateway timeout response has a 5xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application configuration get edge application bundle gateway timeout response a status code equal to that given
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge application configuration get edge application bundle gateway timeout response
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleDefault creates a EdgeApplicationConfigurationGetEdgeApplicationBundleDefault with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleDefault(code int) *EdgeApplicationConfigurationGetEdgeApplicationBundleDefault {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationConfigurationGetEdgeApplicationBundleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge application configuration get edge application bundle default response has a 2xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application configuration get edge application bundle default response has a 3xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application configuration get edge application bundle default response has a 4xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application configuration get edge application bundle default response has a 5xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application configuration get edge application bundle default response a status code equal to that given
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge application configuration get edge application bundle default response
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] EdgeApplicationConfiguration_GetEdgeApplicationBundle default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/id/{id}][%d] EdgeApplicationConfiguration_GetEdgeApplicationBundle default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
