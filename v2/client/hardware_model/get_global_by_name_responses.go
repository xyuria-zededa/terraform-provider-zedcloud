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

// HardwareModelGetGlobalHardwareModelByNameReader is a Reader for the HardwareModelGetGlobalHardwareModelByName structure.
type HardwareModelGetGlobalHardwareModelByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelGetGlobalHardwareModelByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelGetGlobalHardwareModelByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHardwareModelGetGlobalHardwareModelByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelGetGlobalHardwareModelByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHardwareModelGetGlobalHardwareModelByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelGetGlobalHardwareModelByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelGetGlobalHardwareModelByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelGetGlobalHardwareModelByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelGetGlobalHardwareModelByNameOK creates a HardwareModelGetGlobalHardwareModelByNameOK with default headers values
func NewHardwareModelGetGlobalHardwareModelByNameOK() *HardwareModelGetGlobalHardwareModelByNameOK {
	return &HardwareModelGetGlobalHardwareModelByNameOK{}
}

/*
HardwareModelGetGlobalHardwareModelByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelGetGlobalHardwareModelByNameOK struct {
	Payload *models.SysModel
}

// IsSuccess returns true when this hardware model get global hardware model by name o k response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hardware model get global hardware model by name o k response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware model by name o k response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get global hardware model by name o k response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get global hardware model by name o k response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareModelByNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *HardwareModelGetGlobalHardwareModelByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/name/{name}][%d] hardwareModelGetGlobalHardwareModelByNameOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/name/{name}][%d] hardwareModelGetGlobalHardwareModelByNameOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelByNameOK) GetPayload() *models.SysModel {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareModelByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SysModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareModelByNameUnauthorized creates a HardwareModelGetGlobalHardwareModelByNameUnauthorized with default headers values
func NewHardwareModelGetGlobalHardwareModelByNameUnauthorized() *HardwareModelGetGlobalHardwareModelByNameUnauthorized {
	return &HardwareModelGetGlobalHardwareModelByNameUnauthorized{}
}

/*
HardwareModelGetGlobalHardwareModelByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelGetGlobalHardwareModelByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware model by name unauthorized response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware model by name unauthorized response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware model by name unauthorized response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get global hardware model by name unauthorized response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get global hardware model by name unauthorized response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareModelByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *HardwareModelGetGlobalHardwareModelByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/name/{name}][%d] hardwareModelGetGlobalHardwareModelByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/name/{name}][%d] hardwareModelGetGlobalHardwareModelByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareModelByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareModelByNameForbidden creates a HardwareModelGetGlobalHardwareModelByNameForbidden with default headers values
func NewHardwareModelGetGlobalHardwareModelByNameForbidden() *HardwareModelGetGlobalHardwareModelByNameForbidden {
	return &HardwareModelGetGlobalHardwareModelByNameForbidden{}
}

/*
HardwareModelGetGlobalHardwareModelByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelGetGlobalHardwareModelByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware model by name forbidden response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware model by name forbidden response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware model by name forbidden response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get global hardware model by name forbidden response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get global hardware model by name forbidden response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareModelByNameForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *HardwareModelGetGlobalHardwareModelByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/name/{name}][%d] hardwareModelGetGlobalHardwareModelByNameForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/name/{name}][%d] hardwareModelGetGlobalHardwareModelByNameForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareModelByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareModelByNameNotFound creates a HardwareModelGetGlobalHardwareModelByNameNotFound with default headers values
func NewHardwareModelGetGlobalHardwareModelByNameNotFound() *HardwareModelGetGlobalHardwareModelByNameNotFound {
	return &HardwareModelGetGlobalHardwareModelByNameNotFound{}
}

/*
HardwareModelGetGlobalHardwareModelByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type HardwareModelGetGlobalHardwareModelByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware model by name not found response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware model by name not found response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware model by name not found response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get global hardware model by name not found response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get global hardware model by name not found response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareModelByNameNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *HardwareModelGetGlobalHardwareModelByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/name/{name}][%d] hardwareModelGetGlobalHardwareModelByNameNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/name/{name}][%d] hardwareModelGetGlobalHardwareModelByNameNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareModelByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareModelByNameInternalServerError creates a HardwareModelGetGlobalHardwareModelByNameInternalServerError with default headers values
func NewHardwareModelGetGlobalHardwareModelByNameInternalServerError() *HardwareModelGetGlobalHardwareModelByNameInternalServerError {
	return &HardwareModelGetGlobalHardwareModelByNameInternalServerError{}
}

/*
HardwareModelGetGlobalHardwareModelByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelGetGlobalHardwareModelByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware model by name internal server error response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware model by name internal server error response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware model by name internal server error response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get global hardware model by name internal server error response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get global hardware model by name internal server error response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareModelByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *HardwareModelGetGlobalHardwareModelByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/name/{name}][%d] hardwareModelGetGlobalHardwareModelByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/name/{name}][%d] hardwareModelGetGlobalHardwareModelByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareModelByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareModelByNameGatewayTimeout creates a HardwareModelGetGlobalHardwareModelByNameGatewayTimeout with default headers values
func NewHardwareModelGetGlobalHardwareModelByNameGatewayTimeout() *HardwareModelGetGlobalHardwareModelByNameGatewayTimeout {
	return &HardwareModelGetGlobalHardwareModelByNameGatewayTimeout{}
}

/*
HardwareModelGetGlobalHardwareModelByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelGetGlobalHardwareModelByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get global hardware model by name gateway timeout response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get global hardware model by name gateway timeout response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get global hardware model by name gateway timeout response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get global hardware model by name gateway timeout response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get global hardware model by name gateway timeout response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareModelByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *HardwareModelGetGlobalHardwareModelByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/name/{name}][%d] hardwareModelGetGlobalHardwareModelByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/name/{name}][%d] hardwareModelGetGlobalHardwareModelByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareModelByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetGlobalHardwareModelByNameDefault creates a HardwareModelGetGlobalHardwareModelByNameDefault with default headers values
func NewHardwareModelGetGlobalHardwareModelByNameDefault(code int) *HardwareModelGetGlobalHardwareModelByNameDefault {
	return &HardwareModelGetGlobalHardwareModelByNameDefault{
		_statusCode: code,
	}
}

/*
HardwareModelGetGlobalHardwareModelByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelGetGlobalHardwareModelByNameDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the hardware model get global hardware model by name default response
func (o *HardwareModelGetGlobalHardwareModelByNameDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this hardware model get global hardware model by name default response has a 2xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hardware model get global hardware model by name default response has a 3xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hardware model get global hardware model by name default response has a 4xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hardware model get global hardware model by name default response has a 5xx status code
func (o *HardwareModelGetGlobalHardwareModelByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hardware model get global hardware model by name default response a status code equal to that given
func (o *HardwareModelGetGlobalHardwareModelByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HardwareModelGetGlobalHardwareModelByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/name/{name}][%d] HardwareModel_GetGlobalHardwareModelByName default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/sysmodels/global/name/{name}][%d] HardwareModel_GetGlobalHardwareModelByName default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetGlobalHardwareModelByNameDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelGetGlobalHardwareModelByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
