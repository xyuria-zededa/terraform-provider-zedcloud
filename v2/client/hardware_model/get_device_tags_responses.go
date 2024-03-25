package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xyuria-zededa/terraform-provider-zedcloud/models"
)

// HardwareModelGetDeviceTagsReader is a Reader for the HardwareModelGetDeviceTags structure.
type HardwareModelGetDeviceTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelGetDeviceTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelGetDeviceTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewHardwareModelGetDeviceTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewHardwareModelGetDeviceTagsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelGetDeviceTagsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelGetDeviceTagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelGetDeviceTagsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelGetDeviceTagsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelGetDeviceTagsOK creates a HardwareModelGetDeviceTagsOK with default headers values
func NewHardwareModelGetDeviceTagsOK() *HardwareModelGetDeviceTagsOK {
	return &HardwareModelGetDeviceTagsOK{}
}

/*
HardwareModelGetDeviceTagsOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelGetDeviceTagsOK struct {
	Payload *models.ObjectTagsList
}

// IsSuccess returns true when this hardware model get device tags o k response has a 2xx status code
func (o *HardwareModelGetDeviceTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hardware model get device tags o k response has a 3xx status code
func (o *HardwareModelGetDeviceTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get device tags o k response has a 4xx status code
func (o *HardwareModelGetDeviceTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get device tags o k response has a 5xx status code
func (o *HardwareModelGetDeviceTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get device tags o k response a status code equal to that given
func (o *HardwareModelGetDeviceTagsOK) IsCode(code int) bool {
	return code == 200
}

func (o *HardwareModelGetDeviceTagsOK) Error() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetDeviceTagsOK) String() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsOK  %+v", 200, o.Payload)
}

func (o *HardwareModelGetDeviceTagsOK) GetPayload() *models.ObjectTagsList {
	return o.Payload
}

func (o *HardwareModelGetDeviceTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectTagsList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceTagsBadRequest creates a HardwareModelGetDeviceTagsBadRequest with default headers values
func NewHardwareModelGetDeviceTagsBadRequest() *HardwareModelGetDeviceTagsBadRequest {
	return &HardwareModelGetDeviceTagsBadRequest{}
}

/*
HardwareModelGetDeviceTagsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type HardwareModelGetDeviceTagsBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get device tags bad request response has a 2xx status code
func (o *HardwareModelGetDeviceTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get device tags bad request response has a 3xx status code
func (o *HardwareModelGetDeviceTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get device tags bad request response has a 4xx status code
func (o *HardwareModelGetDeviceTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get device tags bad request response has a 5xx status code
func (o *HardwareModelGetDeviceTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get device tags bad request response a status code equal to that given
func (o *HardwareModelGetDeviceTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *HardwareModelGetDeviceTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsBadRequest  %+v", 400, o.Payload)
}

func (o *HardwareModelGetDeviceTagsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsBadRequest  %+v", 400, o.Payload)
}

func (o *HardwareModelGetDeviceTagsBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetDeviceTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceTagsUnauthorized creates a HardwareModelGetDeviceTagsUnauthorized with default headers values
func NewHardwareModelGetDeviceTagsUnauthorized() *HardwareModelGetDeviceTagsUnauthorized {
	return &HardwareModelGetDeviceTagsUnauthorized{}
}

/*
HardwareModelGetDeviceTagsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelGetDeviceTagsUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get device tags unauthorized response has a 2xx status code
func (o *HardwareModelGetDeviceTagsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get device tags unauthorized response has a 3xx status code
func (o *HardwareModelGetDeviceTagsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get device tags unauthorized response has a 4xx status code
func (o *HardwareModelGetDeviceTagsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get device tags unauthorized response has a 5xx status code
func (o *HardwareModelGetDeviceTagsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get device tags unauthorized response a status code equal to that given
func (o *HardwareModelGetDeviceTagsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *HardwareModelGetDeviceTagsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetDeviceTagsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelGetDeviceTagsUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetDeviceTagsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceTagsForbidden creates a HardwareModelGetDeviceTagsForbidden with default headers values
func NewHardwareModelGetDeviceTagsForbidden() *HardwareModelGetDeviceTagsForbidden {
	return &HardwareModelGetDeviceTagsForbidden{}
}

/*
HardwareModelGetDeviceTagsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelGetDeviceTagsForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get device tags forbidden response has a 2xx status code
func (o *HardwareModelGetDeviceTagsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get device tags forbidden response has a 3xx status code
func (o *HardwareModelGetDeviceTagsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get device tags forbidden response has a 4xx status code
func (o *HardwareModelGetDeviceTagsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model get device tags forbidden response has a 5xx status code
func (o *HardwareModelGetDeviceTagsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model get device tags forbidden response a status code equal to that given
func (o *HardwareModelGetDeviceTagsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *HardwareModelGetDeviceTagsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetDeviceTagsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelGetDeviceTagsForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetDeviceTagsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceTagsInternalServerError creates a HardwareModelGetDeviceTagsInternalServerError with default headers values
func NewHardwareModelGetDeviceTagsInternalServerError() *HardwareModelGetDeviceTagsInternalServerError {
	return &HardwareModelGetDeviceTagsInternalServerError{}
}

/*
HardwareModelGetDeviceTagsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelGetDeviceTagsInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get device tags internal server error response has a 2xx status code
func (o *HardwareModelGetDeviceTagsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get device tags internal server error response has a 3xx status code
func (o *HardwareModelGetDeviceTagsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get device tags internal server error response has a 4xx status code
func (o *HardwareModelGetDeviceTagsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get device tags internal server error response has a 5xx status code
func (o *HardwareModelGetDeviceTagsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get device tags internal server error response a status code equal to that given
func (o *HardwareModelGetDeviceTagsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *HardwareModelGetDeviceTagsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetDeviceTagsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelGetDeviceTagsInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetDeviceTagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceTagsGatewayTimeout creates a HardwareModelGetDeviceTagsGatewayTimeout with default headers values
func NewHardwareModelGetDeviceTagsGatewayTimeout() *HardwareModelGetDeviceTagsGatewayTimeout {
	return &HardwareModelGetDeviceTagsGatewayTimeout{}
}

/*
HardwareModelGetDeviceTagsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelGetDeviceTagsGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model get device tags gateway timeout response has a 2xx status code
func (o *HardwareModelGetDeviceTagsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model get device tags gateway timeout response has a 3xx status code
func (o *HardwareModelGetDeviceTagsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model get device tags gateway timeout response has a 4xx status code
func (o *HardwareModelGetDeviceTagsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model get device tags gateway timeout response has a 5xx status code
func (o *HardwareModelGetDeviceTagsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model get device tags gateway timeout response a status code equal to that given
func (o *HardwareModelGetDeviceTagsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *HardwareModelGetDeviceTagsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetDeviceTagsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] hardwareModelGetDeviceTagsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelGetDeviceTagsGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelGetDeviceTagsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelGetDeviceTagsDefault creates a HardwareModelGetDeviceTagsDefault with default headers values
func NewHardwareModelGetDeviceTagsDefault(code int) *HardwareModelGetDeviceTagsDefault {
	return &HardwareModelGetDeviceTagsDefault{
		_statusCode: code,
	}
}

/*
HardwareModelGetDeviceTagsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelGetDeviceTagsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the hardware model get device tags default response
func (o *HardwareModelGetDeviceTagsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this hardware model get device tags default response has a 2xx status code
func (o *HardwareModelGetDeviceTagsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hardware model get device tags default response has a 3xx status code
func (o *HardwareModelGetDeviceTagsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hardware model get device tags default response has a 4xx status code
func (o *HardwareModelGetDeviceTagsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hardware model get device tags default response has a 5xx status code
func (o *HardwareModelGetDeviceTagsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hardware model get device tags default response a status code equal to that given
func (o *HardwareModelGetDeviceTagsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HardwareModelGetDeviceTagsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] HardwareModel_GetDeviceTags default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetDeviceTagsDefault) String() string {
	return fmt.Sprintf("[GET /v1/devices/tags][%d] HardwareModel_GetDeviceTags default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelGetDeviceTagsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelGetDeviceTagsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
