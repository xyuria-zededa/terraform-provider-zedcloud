package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xyuria-zededa/terraform-provider-zedcloud/v2/models"
)

// HardwareModelGetHardwareBrandByNameReader is a Reader for the HardwareModelGetHardwareBrandByName structure.
type HardwareModelGetHardwareBrandByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelGetHardwareBrandByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelGetHardwareBrandByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHardwareModelGetHardwareBrandByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelGetHardwareBrandByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHardwareModelGetHardwareBrandByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelGetHardwareBrandByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelGetHardwareBrandByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelGetHardwareBrandByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelGetHardwareBrandByNameOK creates a HardwareModelGetHardwareBrandByNameOK with default headers values
func NewHardwareModelGetHardwareBrandByNameOK() *HardwareModelGetHardwareBrandByNameOK {
	return &HardwareModelGetHardwareBrandByNameOK{}
}

/*
HardwareModelGetHardwareBrandByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelGetHardwareBrandByNameOK struct {
	Payload *models.SysBrand
}

// IsSuccess returns true when this hardware model get hardware brand by name o k response has a 2xx status code
func (o *HardwareModelGetHardwareBrandByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hardware model get hardware brand by name o k response has a 3xx status code
func (o *HardwareModelGetHardwareBrandByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware brand by name o k response has a 4xx status code
func (o *HardwareModelGetHardwareBrandByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get hardware brand by name o k response has a 5xx status code
func (o *HardwareModelGetHardwareBrandByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get hardware brand by name o k response a status code equal to that given
func (o *HardwareModelGetHardwareBrandByNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *HardwareModelGetHardwareBrandByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] hardwareModelGetHardwareBrandByNameOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetHardwareBrandByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] hardwareModelGetHardwareBrandByNameOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetHardwareBrandByNameOK) GetPayload() *models.SysBrand {
	return o.Payload
}

func (o *HardwareModelGetHardwareBrandByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SysBrand)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareBrandByNameUnauthorized creates a HardwareModelGetHardwareBrandByNameUnauthorized with default headers values
func NewHardwareModelGetHardwareBrandByNameUnauthorized() *HardwareModelGetHardwareBrandByNameUnauthorized {
	return &HardwareModelGetHardwareBrandByNameUnauthorized{}
}

/*
HardwareModelGetHardwareBrandByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelGetHardwareBrandByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware brand by name unauthorized response has a 2xx status code
func (o *HardwareModelGetHardwareBrandByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware brand by name unauthorized response has a 3xx status code
func (o *HardwareModelGetHardwareBrandByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware brand by name unauthorized response has a 4xx status code
func (o *HardwareModelGetHardwareBrandByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get hardware brand by name unauthorized response has a 5xx status code
func (o *HardwareModelGetHardwareBrandByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get hardware brand by name unauthorized response a status code equal to that given
func (o *HardwareModelGetHardwareBrandByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *HardwareModelGetHardwareBrandByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] hardwareModelGetHardwareBrandByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetHardwareBrandByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] hardwareModelGetHardwareBrandByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetHardwareBrandByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareBrandByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareBrandByNameForbidden creates a HardwareModelGetHardwareBrandByNameForbidden with default headers values
func NewHardwareModelGetHardwareBrandByNameForbidden() *HardwareModelGetHardwareBrandByNameForbidden {
	return &HardwareModelGetHardwareBrandByNameForbidden{}
}

/*
HardwareModelGetHardwareBrandByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelGetHardwareBrandByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware brand by name forbidden response has a 2xx status code
func (o *HardwareModelGetHardwareBrandByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware brand by name forbidden response has a 3xx status code
func (o *HardwareModelGetHardwareBrandByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware brand by name forbidden response has a 4xx status code
func (o *HardwareModelGetHardwareBrandByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get hardware brand by name forbidden response has a 5xx status code
func (o *HardwareModelGetHardwareBrandByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get hardware brand by name forbidden response a status code equal to that given
func (o *HardwareModelGetHardwareBrandByNameForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *HardwareModelGetHardwareBrandByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] hardwareModelGetHardwareBrandByNameForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetHardwareBrandByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] hardwareModelGetHardwareBrandByNameForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetHardwareBrandByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareBrandByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareBrandByNameNotFound creates a HardwareModelGetHardwareBrandByNameNotFound with default headers values
func NewHardwareModelGetHardwareBrandByNameNotFound() *HardwareModelGetHardwareBrandByNameNotFound {
	return &HardwareModelGetHardwareBrandByNameNotFound{}
}

/*
HardwareModelGetHardwareBrandByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type HardwareModelGetHardwareBrandByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware brand by name not found response has a 2xx status code
func (o *HardwareModelGetHardwareBrandByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware brand by name not found response has a 3xx status code
func (o *HardwareModelGetHardwareBrandByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware brand by name not found response has a 4xx status code
func (o *HardwareModelGetHardwareBrandByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get hardware brand by name not found response has a 5xx status code
func (o *HardwareModelGetHardwareBrandByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get hardware brand by name not found response a status code equal to that given
func (o *HardwareModelGetHardwareBrandByNameNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *HardwareModelGetHardwareBrandByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] hardwareModelGetHardwareBrandByNameNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelGetHardwareBrandByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] hardwareModelGetHardwareBrandByNameNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelGetHardwareBrandByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareBrandByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareBrandByNameInternalServerError creates a HardwareModelGetHardwareBrandByNameInternalServerError with default headers values
func NewHardwareModelGetHardwareBrandByNameInternalServerError() *HardwareModelGetHardwareBrandByNameInternalServerError {
	return &HardwareModelGetHardwareBrandByNameInternalServerError{}
}

/*
HardwareModelGetHardwareBrandByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelGetHardwareBrandByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware brand by name internal server error response has a 2xx status code
func (o *HardwareModelGetHardwareBrandByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware brand by name internal server error response has a 3xx status code
func (o *HardwareModelGetHardwareBrandByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware brand by name internal server error response has a 4xx status code
func (o *HardwareModelGetHardwareBrandByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get hardware brand by name internal server error response has a 5xx status code
func (o *HardwareModelGetHardwareBrandByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get hardware brand by name internal server error response a status code equal to that given
func (o *HardwareModelGetHardwareBrandByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *HardwareModelGetHardwareBrandByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] hardwareModelGetHardwareBrandByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetHardwareBrandByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] hardwareModelGetHardwareBrandByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetHardwareBrandByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareBrandByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareBrandByNameGatewayTimeout creates a HardwareModelGetHardwareBrandByNameGatewayTimeout with default headers values
func NewHardwareModelGetHardwareBrandByNameGatewayTimeout() *HardwareModelGetHardwareBrandByNameGatewayTimeout {
	return &HardwareModelGetHardwareBrandByNameGatewayTimeout{}
}

/*
HardwareModelGetHardwareBrandByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelGetHardwareBrandByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get hardware brand by name gateway timeout response has a 2xx status code
func (o *HardwareModelGetHardwareBrandByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get hardware brand by name gateway timeout response has a 3xx status code
func (o *HardwareModelGetHardwareBrandByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get hardware brand by name gateway timeout response has a 4xx status code
func (o *HardwareModelGetHardwareBrandByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get hardware brand by name gateway timeout response has a 5xx status code
func (o *HardwareModelGetHardwareBrandByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get hardware brand by name gateway timeout response a status code equal to that given
func (o *HardwareModelGetHardwareBrandByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *HardwareModelGetHardwareBrandByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] hardwareModelGetHardwareBrandByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetHardwareBrandByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] hardwareModelGetHardwareBrandByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetHardwareBrandByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetHardwareBrandByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetHardwareBrandByNameDefault creates a HardwareModelGetHardwareBrandByNameDefault with default headers values
func NewHardwareModelGetHardwareBrandByNameDefault(code int) *HardwareModelGetHardwareBrandByNameDefault {
	return &HardwareModelGetHardwareBrandByNameDefault{
		_statusCode: code,
	}
}

/*
HardwareModelGetHardwareBrandByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelGetHardwareBrandByNameDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the hardware model get hardware brand by name default response
func (o *HardwareModelGetHardwareBrandByNameDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this hardware model get hardware brand by name default response has a 2xx status code
func (o *HardwareModelGetHardwareBrandByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hardware model get hardware brand by name default response has a 3xx status code
func (o *HardwareModelGetHardwareBrandByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hardware model get hardware brand by name default response has a 4xx status code
func (o *HardwareModelGetHardwareBrandByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hardware model get hardware brand by name default response has a 5xx status code
func (o *HardwareModelGetHardwareBrandByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hardware model get hardware brand by name default response a status code equal to that given
func (o *HardwareModelGetHardwareBrandByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HardwareModelGetHardwareBrandByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] HardwareModel_GetHardwareBrandByName default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetHardwareBrandByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/brands/name/{name}][%d] HardwareModel_GetHardwareBrandByName default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetHardwareBrandByNameDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelGetHardwareBrandByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
