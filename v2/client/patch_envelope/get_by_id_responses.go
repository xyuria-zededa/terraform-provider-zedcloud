// Code generated by go-swagger; DO NOT EDIT.

package patch_envelope

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xyuria-zededa/terraform-provider-zedcloud/v2/models"
)

// PatchEnvelopeConfigurationGetPatchEnvelopeByIDReader is a Reader for the PatchEnvelopeConfigurationGetPatchEnvelopeByID structure.
type PatchEnvelopeConfigurationGetPatchEnvelopeByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDOK creates a PatchEnvelopeConfigurationGetPatchEnvelopeByIDOK with default headers values
func NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDOK() *PatchEnvelopeConfigurationGetPatchEnvelopeByIDOK {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeByIDOK{}
}

/*
PatchEnvelopeConfigurationGetPatchEnvelopeByIDOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchEnvelopeConfigurationGetPatchEnvelopeByIDOK struct {
	Payload *models.PatchEnvelope
}

// IsSuccess returns true when this patch envelope configuration get patch envelope by Id o k response has a 2xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch envelope configuration get patch envelope by Id o k response has a 3xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration get patch envelope by Id o k response has a 4xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration get patch envelope by Id o k response has a 5xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration get patch envelope by Id o k response a status code equal to that given
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch envelope configuration get patch envelope by Id o k response
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDOK) Code() int {
	return 200
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByIdOK  %+v", 200, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDOK) String() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByIdOK  %+v", 200, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDOK) GetPayload() *models.PatchEnvelope {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PatchEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized creates a PatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized with default headers values
func NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized() *PatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized{}
}

/*
PatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type PatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration get patch envelope by Id unauthorized response has a 2xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration get patch envelope by Id unauthorized response has a 3xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration get patch envelope by Id unauthorized response has a 4xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch envelope configuration get patch envelope by Id unauthorized response has a 5xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration get patch envelope by Id unauthorized response a status code equal to that given
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch envelope configuration get patch envelope by Id unauthorized response
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized) Code() int {
	return 401
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden creates a PatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden with default headers values
func NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden() *PatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden{}
}

/*
PatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type PatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration get patch envelope by Id forbidden response has a 2xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration get patch envelope by Id forbidden response has a 3xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration get patch envelope by Id forbidden response has a 4xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch envelope configuration get patch envelope by Id forbidden response has a 5xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration get patch envelope by Id forbidden response a status code equal to that given
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch envelope configuration get patch envelope by Id forbidden response
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden) Code() int {
	return 403
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden) String() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound creates a PatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound with default headers values
func NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound() *PatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound{}
}

/*
PatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type PatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration get patch envelope by Id not found response has a 2xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration get patch envelope by Id not found response has a 3xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration get patch envelope by Id not found response has a 4xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch envelope configuration get patch envelope by Id not found response has a 5xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration get patch envelope by Id not found response a status code equal to that given
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the patch envelope configuration get patch envelope by Id not found response
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound) Code() int {
	return 404
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound) String() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError creates a PatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError with default headers values
func NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError() *PatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError{}
}

/*
PatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type PatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration get patch envelope by Id internal server error response has a 2xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration get patch envelope by Id internal server error response has a 3xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration get patch envelope by Id internal server error response has a 4xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration get patch envelope by Id internal server error response has a 5xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch envelope configuration get patch envelope by Id internal server error response a status code equal to that given
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the patch envelope configuration get patch envelope by Id internal server error response
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError) Code() int {
	return 500
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout creates a PatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout with default headers values
func NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout() *PatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout{}
}

/*
PatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type PatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration get patch envelope by Id gateway timeout response has a 2xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration get patch envelope by Id gateway timeout response has a 3xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration get patch envelope by Id gateway timeout response has a 4xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration get patch envelope by Id gateway timeout response has a 5xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch envelope configuration get patch envelope by Id gateway timeout response a status code equal to that given
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the patch envelope configuration get patch envelope by Id gateway timeout response
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout) Code() int {
	return 504
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationGetPatchEnvelopeByIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault creates a PatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault with default headers values
func NewPatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault(code int) *PatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault{
		_statusCode: code,
	}
}

/*
PatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this patch envelope configuration get patch envelope by Id default response has a 2xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch envelope configuration get patch envelope by Id default response has a 3xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch envelope configuration get patch envelope by Id default response has a 4xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch envelope configuration get patch envelope by Id default response has a 5xx status code
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch envelope configuration get patch envelope by Id default response a status code equal to that given
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch envelope configuration get patch envelope by Id default response
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault) Error() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/id/{id}][%d] PatchEnvelopeConfiguration_GetPatchEnvelopeById default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault) String() string {
	return fmt.Sprintf("[GET /v1/patch-envelope/id/{id}][%d] PatchEnvelopeConfiguration_GetPatchEnvelopeById default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationGetPatchEnvelopeByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
