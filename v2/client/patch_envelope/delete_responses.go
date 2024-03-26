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

// PatchEnvelopeConfigurationDeletePatchEnvelopeReader is a Reader for the PatchEnvelopeConfigurationDeletePatchEnvelope structure.
type PatchEnvelopeConfigurationDeletePatchEnvelopeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchEnvelopeConfigurationDeletePatchEnvelopeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchEnvelopeConfigurationDeletePatchEnvelopeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPatchEnvelopeConfigurationDeletePatchEnvelopeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchEnvelopeConfigurationDeletePatchEnvelopeOK creates a PatchEnvelopeConfigurationDeletePatchEnvelopeOK with default headers values
func NewPatchEnvelopeConfigurationDeletePatchEnvelopeOK() *PatchEnvelopeConfigurationDeletePatchEnvelopeOK {
	return &PatchEnvelopeConfigurationDeletePatchEnvelopeOK{}
}

/*
PatchEnvelopeConfigurationDeletePatchEnvelopeOK describes a response with status code 200, with default header values.

A successful response.
*/
type PatchEnvelopeConfigurationDeletePatchEnvelopeOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration delete patch envelope o k response has a 2xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch envelope configuration delete patch envelope o k response has a 3xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration delete patch envelope o k response has a 4xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration delete patch envelope o k response has a 5xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration delete patch envelope o k response a status code equal to that given
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch envelope configuration delete patch envelope o k response
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeOK) Code() int {
	return 200
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeletePatchEnvelopeOK  %+v", 200, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeOK) String() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeletePatchEnvelopeOK  %+v", 200, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized creates a PatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized with default headers values
func NewPatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized() *PatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized {
	return &PatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized{}
}

/*
PatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type PatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration delete patch envelope unauthorized response has a 2xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration delete patch envelope unauthorized response has a 3xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration delete patch envelope unauthorized response has a 4xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch envelope configuration delete patch envelope unauthorized response has a 5xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration delete patch envelope unauthorized response a status code equal to that given
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch envelope configuration delete patch envelope unauthorized response
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized) Code() int {
	return 401
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationDeletePatchEnvelopeForbidden creates a PatchEnvelopeConfigurationDeletePatchEnvelopeForbidden with default headers values
func NewPatchEnvelopeConfigurationDeletePatchEnvelopeForbidden() *PatchEnvelopeConfigurationDeletePatchEnvelopeForbidden {
	return &PatchEnvelopeConfigurationDeletePatchEnvelopeForbidden{}
}

/*
PatchEnvelopeConfigurationDeletePatchEnvelopeForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type PatchEnvelopeConfigurationDeletePatchEnvelopeForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration delete patch envelope forbidden response has a 2xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration delete patch envelope forbidden response has a 3xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration delete patch envelope forbidden response has a 4xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch envelope configuration delete patch envelope forbidden response has a 5xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch envelope configuration delete patch envelope forbidden response a status code equal to that given
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch envelope configuration delete patch envelope forbidden response
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeForbidden) Code() int {
	return 403
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeletePatchEnvelopeForbidden  %+v", 403, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeletePatchEnvelopeForbidden  %+v", 403, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError creates a PatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError with default headers values
func NewPatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError() *PatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError {
	return &PatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError{}
}

/*
PatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type PatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration delete patch envelope internal server error response has a 2xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration delete patch envelope internal server error response has a 3xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration delete patch envelope internal server error response has a 4xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration delete patch envelope internal server error response has a 5xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch envelope configuration delete patch envelope internal server error response a status code equal to that given
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the patch envelope configuration delete patch envelope internal server error response
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError) Code() int {
	return 500
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout creates a PatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout with default headers values
func NewPatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout() *PatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout {
	return &PatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout{}
}

/*
PatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type PatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this patch envelope configuration delete patch envelope gateway timeout response has a 2xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch envelope configuration delete patch envelope gateway timeout response has a 3xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch envelope configuration delete patch envelope gateway timeout response has a 4xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch envelope configuration delete patch envelope gateway timeout response has a 5xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch envelope configuration delete patch envelope gateway timeout response a status code equal to that given
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the patch envelope configuration delete patch envelope gateway timeout response
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout) Code() int {
	return 504
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] patchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchEnvelopeConfigurationDeletePatchEnvelopeDefault creates a PatchEnvelopeConfigurationDeletePatchEnvelopeDefault with default headers values
func NewPatchEnvelopeConfigurationDeletePatchEnvelopeDefault(code int) *PatchEnvelopeConfigurationDeletePatchEnvelopeDefault {
	return &PatchEnvelopeConfigurationDeletePatchEnvelopeDefault{
		_statusCode: code,
	}
}

/*
PatchEnvelopeConfigurationDeletePatchEnvelopeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type PatchEnvelopeConfigurationDeletePatchEnvelopeDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this patch envelope configuration delete patch envelope default response has a 2xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this patch envelope configuration delete patch envelope default response has a 3xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this patch envelope configuration delete patch envelope default response has a 4xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this patch envelope configuration delete patch envelope default response has a 5xx status code
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this patch envelope configuration delete patch envelope default response a status code equal to that given
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the patch envelope configuration delete patch envelope default response
func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeDefault) Code() int {
	return o._statusCode
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] PatchEnvelopeConfiguration_DeletePatchEnvelope default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/patch-envelope/id/{id}][%d] PatchEnvelopeConfiguration_DeletePatchEnvelope default  %+v", o._statusCode, o.Payload)
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *PatchEnvelopeConfigurationDeletePatchEnvelopeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
