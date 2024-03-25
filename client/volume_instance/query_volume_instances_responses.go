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

// VolumeInstanceConfigurationQueryVolumeInstancesReader is a Reader for the VolumeInstanceConfigurationQueryVolumeInstances structure.
type VolumeInstanceConfigurationQueryVolumeInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeInstanceConfigurationQueryVolumeInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeInstanceConfigurationQueryVolumeInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVolumeInstanceConfigurationQueryVolumeInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewVolumeInstanceConfigurationQueryVolumeInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewVolumeInstanceConfigurationQueryVolumeInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVolumeInstanceConfigurationQueryVolumeInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewVolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewVolumeInstanceConfigurationQueryVolumeInstancesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeInstanceConfigurationQueryVolumeInstancesOK creates a VolumeInstanceConfigurationQueryVolumeInstancesOK with default headers values
func NewVolumeInstanceConfigurationQueryVolumeInstancesOK() *VolumeInstanceConfigurationQueryVolumeInstancesOK {
	return &VolumeInstanceConfigurationQueryVolumeInstancesOK{}
}

/*
VolumeInstanceConfigurationQueryVolumeInstancesOK describes a response with status code 200, with default header values.

A successful response.
*/
type VolumeInstanceConfigurationQueryVolumeInstancesOK struct {
	Payload *models.VolInstList
}

// IsSuccess returns true when this volume instance configuration query volume instances o k response has a 2xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume instance configuration query volume instances o k response has a 3xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration query volume instances o k response has a 4xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance configuration query volume instances o k response has a 5xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration query volume instances o k response a status code equal to that given
func (o *VolumeInstanceConfigurationQueryVolumeInstancesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the volume instance configuration query volume instances o k response
func (o *VolumeInstanceConfigurationQueryVolumeInstancesOK) Code() int {
	return 200
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesOK) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances][%d] volumeInstanceConfigurationQueryVolumeInstancesOK  %+v", 200, o.Payload)
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesOK) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances][%d] volumeInstanceConfigurationQueryVolumeInstancesOK  %+v", 200, o.Payload)
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesOK) GetPayload() *models.VolInstList {
	return o.Payload
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolInstList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationQueryVolumeInstancesBadRequest creates a VolumeInstanceConfigurationQueryVolumeInstancesBadRequest with default headers values
func NewVolumeInstanceConfigurationQueryVolumeInstancesBadRequest() *VolumeInstanceConfigurationQueryVolumeInstancesBadRequest {
	return &VolumeInstanceConfigurationQueryVolumeInstancesBadRequest{}
}

/*
VolumeInstanceConfigurationQueryVolumeInstancesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type VolumeInstanceConfigurationQueryVolumeInstancesBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration query volume instances bad request response has a 2xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration query volume instances bad request response has a 3xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration query volume instances bad request response has a 4xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance configuration query volume instances bad request response has a 5xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration query volume instances bad request response a status code equal to that given
func (o *VolumeInstanceConfigurationQueryVolumeInstancesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the volume instance configuration query volume instances bad request response
func (o *VolumeInstanceConfigurationQueryVolumeInstancesBadRequest) Code() int {
	return 400
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances][%d] volumeInstanceConfigurationQueryVolumeInstancesBadRequest  %+v", 400, o.Payload)
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances][%d] volumeInstanceConfigurationQueryVolumeInstancesBadRequest  %+v", 400, o.Payload)
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationQueryVolumeInstancesUnauthorized creates a VolumeInstanceConfigurationQueryVolumeInstancesUnauthorized with default headers values
func NewVolumeInstanceConfigurationQueryVolumeInstancesUnauthorized() *VolumeInstanceConfigurationQueryVolumeInstancesUnauthorized {
	return &VolumeInstanceConfigurationQueryVolumeInstancesUnauthorized{}
}

/*
VolumeInstanceConfigurationQueryVolumeInstancesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type VolumeInstanceConfigurationQueryVolumeInstancesUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration query volume instances unauthorized response has a 2xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration query volume instances unauthorized response has a 3xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration query volume instances unauthorized response has a 4xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance configuration query volume instances unauthorized response has a 5xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration query volume instances unauthorized response a status code equal to that given
func (o *VolumeInstanceConfigurationQueryVolumeInstancesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the volume instance configuration query volume instances unauthorized response
func (o *VolumeInstanceConfigurationQueryVolumeInstancesUnauthorized) Code() int {
	return 401
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances][%d] volumeInstanceConfigurationQueryVolumeInstancesUnauthorized  %+v", 401, o.Payload)
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances][%d] volumeInstanceConfigurationQueryVolumeInstancesUnauthorized  %+v", 401, o.Payload)
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationQueryVolumeInstancesForbidden creates a VolumeInstanceConfigurationQueryVolumeInstancesForbidden with default headers values
func NewVolumeInstanceConfigurationQueryVolumeInstancesForbidden() *VolumeInstanceConfigurationQueryVolumeInstancesForbidden {
	return &VolumeInstanceConfigurationQueryVolumeInstancesForbidden{}
}

/*
VolumeInstanceConfigurationQueryVolumeInstancesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type VolumeInstanceConfigurationQueryVolumeInstancesForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration query volume instances forbidden response has a 2xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration query volume instances forbidden response has a 3xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration query volume instances forbidden response has a 4xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this volume instance configuration query volume instances forbidden response has a 5xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this volume instance configuration query volume instances forbidden response a status code equal to that given
func (o *VolumeInstanceConfigurationQueryVolumeInstancesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the volume instance configuration query volume instances forbidden response
func (o *VolumeInstanceConfigurationQueryVolumeInstancesForbidden) Code() int {
	return 403
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances][%d] volumeInstanceConfigurationQueryVolumeInstancesForbidden  %+v", 403, o.Payload)
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesForbidden) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances][%d] volumeInstanceConfigurationQueryVolumeInstancesForbidden  %+v", 403, o.Payload)
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationQueryVolumeInstancesInternalServerError creates a VolumeInstanceConfigurationQueryVolumeInstancesInternalServerError with default headers values
func NewVolumeInstanceConfigurationQueryVolumeInstancesInternalServerError() *VolumeInstanceConfigurationQueryVolumeInstancesInternalServerError {
	return &VolumeInstanceConfigurationQueryVolumeInstancesInternalServerError{}
}

/*
VolumeInstanceConfigurationQueryVolumeInstancesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type VolumeInstanceConfigurationQueryVolumeInstancesInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration query volume instances internal server error response has a 2xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration query volume instances internal server error response has a 3xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration query volume instances internal server error response has a 4xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance configuration query volume instances internal server error response has a 5xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this volume instance configuration query volume instances internal server error response a status code equal to that given
func (o *VolumeInstanceConfigurationQueryVolumeInstancesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the volume instance configuration query volume instances internal server error response
func (o *VolumeInstanceConfigurationQueryVolumeInstancesInternalServerError) Code() int {
	return 500
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances][%d] volumeInstanceConfigurationQueryVolumeInstancesInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances][%d] volumeInstanceConfigurationQueryVolumeInstancesInternalServerError  %+v", 500, o.Payload)
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout creates a VolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout with default headers values
func NewVolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout() *VolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout {
	return &VolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout{}
}

/*
VolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type VolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this volume instance configuration query volume instances gateway timeout response has a 2xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this volume instance configuration query volume instances gateway timeout response has a 3xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume instance configuration query volume instances gateway timeout response has a 4xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume instance configuration query volume instances gateway timeout response has a 5xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this volume instance configuration query volume instances gateway timeout response a status code equal to that given
func (o *VolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the volume instance configuration query volume instances gateway timeout response
func (o *VolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout) Code() int {
	return 504
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances][%d] volumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances][%d] volumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInstanceConfigurationQueryVolumeInstancesDefault creates a VolumeInstanceConfigurationQueryVolumeInstancesDefault with default headers values
func NewVolumeInstanceConfigurationQueryVolumeInstancesDefault(code int) *VolumeInstanceConfigurationQueryVolumeInstancesDefault {
	return &VolumeInstanceConfigurationQueryVolumeInstancesDefault{
		_statusCode: code,
	}
}

/*
VolumeInstanceConfigurationQueryVolumeInstancesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type VolumeInstanceConfigurationQueryVolumeInstancesDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this volume instance configuration query volume instances default response has a 2xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this volume instance configuration query volume instances default response has a 3xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this volume instance configuration query volume instances default response has a 4xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this volume instance configuration query volume instances default response has a 5xx status code
func (o *VolumeInstanceConfigurationQueryVolumeInstancesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this volume instance configuration query volume instances default response a status code equal to that given
func (o *VolumeInstanceConfigurationQueryVolumeInstancesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the volume instance configuration query volume instances default response
func (o *VolumeInstanceConfigurationQueryVolumeInstancesDefault) Code() int {
	return o._statusCode
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/volumes/instances][%d] VolumeInstanceConfiguration_QueryVolumeInstances default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesDefault) String() string {
	return fmt.Sprintf("[GET /v1/volumes/instances][%d] VolumeInstanceConfiguration_QueryVolumeInstances default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *VolumeInstanceConfigurationQueryVolumeInstancesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
