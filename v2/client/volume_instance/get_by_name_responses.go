package volume_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xyuria-zededa/terraform-provider-zedcloud/models"
)

// VolumeInstanceConfigurationGetVolumeInstanceByNameReader is a Reader for the VolumeInstanceConfigurationGetVolumeInstanceByName structure.
type VolumeInstanceConfigurationGetVolumeInstanceByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeInstanceConfigurationGetVolumeInstanceByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewVolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewVolumeInstanceConfigurationGetVolumeInstanceByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVolumeInstanceConfigurationGetVolumeInstanceByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewVolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewVolumeInstanceConfigurationGetVolumeInstanceByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeInstanceConfigurationGetVolumeInstanceByNameOK creates a VolumeInstanceConfigurationGetVolumeInstanceByNameOK with default headers values
func NewVolumeInstanceConfigurationGetVolumeInstanceByNameOK() *VolumeInstanceConfigurationGetVolumeInstanceByNameOK {
	return &VolumeInstanceConfigurationGetVolumeInstanceByNameOK{}
}

/*
VolumeInstanceConfigurationGetVolumeInstanceByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type VolumeInstanceConfigurationGetVolumeInstanceByNameOK struct {
	Payload *models.VolumeInstance
}

// IsSuccess returns true when this volume instance configuration get volume instance by name o k response has a 2xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume instance configuration get volume instance by name o k response has a 3xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration get volume instance by name o k response has a 4xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance configuration get volume instance by name o k response has a 5xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration get volume instance by name o k response a status code equal to that given
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the volume instance configuration get volume instance by name o k response
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameOK) Code() int {
	return 200
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}][%d] volumeInstanceConfigurationGetVolumeInstanceByNameOK  %+v", 200, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}][%d] volumeInstanceConfigurationGetVolumeInstanceByNameOK  %+v", 200, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameOK) GetPayload() *models.VolumeInstance {
	return o.Payload
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized creates a VolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized with default headers values
func NewVolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized() *VolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized {
	return &VolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized{}
}

/*
VolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type VolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration get volume instance by name unauthorized response has a 2xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration get volume instance by name unauthorized response has a 3xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration get volume instance by name unauthorized response has a 4xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance configuration get volume instance by name unauthorized response has a 5xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration get volume instance by name unauthorized response a status code equal to that given
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the volume instance configuration get volume instance by name unauthorized response
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized) Code() int {
	return 401
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}][%d] volumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}][%d] volumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationGetVolumeInstanceByNameForbidden creates a VolumeInstanceConfigurationGetVolumeInstanceByNameForbidden with default headers values
func NewVolumeInstanceConfigurationGetVolumeInstanceByNameForbidden() *VolumeInstanceConfigurationGetVolumeInstanceByNameForbidden {
	return &VolumeInstanceConfigurationGetVolumeInstanceByNameForbidden{}
}

/*
VolumeInstanceConfigurationGetVolumeInstanceByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type VolumeInstanceConfigurationGetVolumeInstanceByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration get volume instance by name forbidden response has a 2xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration get volume instance by name forbidden response has a 3xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration get volume instance by name forbidden response has a 4xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance configuration get volume instance by name forbidden response has a 5xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration get volume instance by name forbidden response a status code equal to that given
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the volume instance configuration get volume instance by name forbidden response
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameForbidden) Code() int {
	return 403
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}][%d] volumeInstanceConfigurationGetVolumeInstanceByNameForbidden  %+v", 403, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}][%d] volumeInstanceConfigurationGetVolumeInstanceByNameForbidden  %+v", 403, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationGetVolumeInstanceByNameNotFound creates a VolumeInstanceConfigurationGetVolumeInstanceByNameNotFound with default headers values
func NewVolumeInstanceConfigurationGetVolumeInstanceByNameNotFound() *VolumeInstanceConfigurationGetVolumeInstanceByNameNotFound {
	return &VolumeInstanceConfigurationGetVolumeInstanceByNameNotFound{}
}

/*
VolumeInstanceConfigurationGetVolumeInstanceByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type VolumeInstanceConfigurationGetVolumeInstanceByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration get volume instance by name not found response has a 2xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration get volume instance by name not found response has a 3xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration get volume instance by name not found response has a 4xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance configuration get volume instance by name not found response has a 5xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration get volume instance by name not found response a status code equal to that given
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the volume instance configuration get volume instance by name not found response
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameNotFound) Code() int {
	return 404
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}][%d] volumeInstanceConfigurationGetVolumeInstanceByNameNotFound  %+v", 404, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}][%d] volumeInstanceConfigurationGetVolumeInstanceByNameNotFound  %+v", 404, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError creates a VolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError with default headers values
func NewVolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError() *VolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError {
	return &VolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError{}
}

/*
VolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type VolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration get volume instance by name internal server error response has a 2xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration get volume instance by name internal server error response has a 3xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration get volume instance by name internal server error response has a 4xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance configuration get volume instance by name internal server error response has a 5xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this volume instance configuration get volume instance by name internal server error response a status code equal to that given
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the volume instance configuration get volume instance by name internal server error response
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError) Code() int {
	return 500
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}][%d] volumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}][%d] volumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout creates a VolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout with default headers values
func NewVolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout() *VolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout {
	return &VolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout{}
}

/*
VolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type VolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration get volume instance by name gateway timeout response has a 2xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration get volume instance by name gateway timeout response has a 3xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration get volume instance by name gateway timeout response has a 4xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance configuration get volume instance by name gateway timeout response has a 5xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this volume instance configuration get volume instance by name gateway timeout response a status code equal to that given
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the volume instance configuration get volume instance by name gateway timeout response
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout) Code() int {
	return 504
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}][%d] volumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}][%d] volumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationGetVolumeInstanceByNameDefault creates a VolumeInstanceConfigurationGetVolumeInstanceByNameDefault with default headers values
func NewVolumeInstanceConfigurationGetVolumeInstanceByNameDefault(code int) *VolumeInstanceConfigurationGetVolumeInstanceByNameDefault {
	return &VolumeInstanceConfigurationGetVolumeInstanceByNameDefault{
		_statusCode: code,
	}
}

/*
VolumeInstanceConfigurationGetVolumeInstanceByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type VolumeInstanceConfigurationGetVolumeInstanceByNameDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this volume instance configuration get volume instance by name default response has a 2xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this volume instance configuration get volume instance by name default response has a 3xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this volume instance configuration get volume instance by name default response has a 4xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this volume instance configuration get volume instance by name default response has a 5xx status code
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this volume instance configuration get volume instance by name default response a status code equal to that given
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the volume instance configuration get volume instance by name default response
func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameDefault) Code() int {
	return o._statusCode
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}][%d] VolumeInstanceConfiguration_GetVolumeInstanceByName default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances/name/{name}][%d] VolumeInstanceConfiguration_GetVolumeInstanceByName default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *VolumeInstanceConfigurationGetVolumeInstanceByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
