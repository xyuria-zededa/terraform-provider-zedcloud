package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xyuria-zededa/terraform-provider-zedcloud/v2/models"
)

// ImageConfigurationQueryLatestImageVersionsReader is a Reader for the ImageConfigurationQueryLatestImageVersions structure.
type ImageConfigurationQueryLatestImageVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationQueryLatestImageVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationQueryLatestImageVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImageConfigurationQueryLatestImageVersionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImageConfigurationQueryLatestImageVersionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationQueryLatestImageVersionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationQueryLatestImageVersionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationQueryLatestImageVersionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationQueryLatestImageVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationQueryLatestImageVersionsOK creates a ImageConfigurationQueryLatestImageVersionsOK with default headers values
func NewImageConfigurationQueryLatestImageVersionsOK() *ImageConfigurationQueryLatestImageVersionsOK {
	return &ImageConfigurationQueryLatestImageVersionsOK{}
}

/*
ImageConfigurationQueryLatestImageVersionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationQueryLatestImageVersionsOK struct {
	Payload *models.Images
}

// IsSuccess returns true when this image configuration query latest image versions o k response has a 2xx status code
func (o *ImageConfigurationQueryLatestImageVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image configuration query latest image versions o k response has a 3xx status code
func (o *ImageConfigurationQueryLatestImageVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query latest image versions o k response has a 4xx status code
func (o *ImageConfigurationQueryLatestImageVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration query latest image versions o k response has a 5xx status code
func (o *ImageConfigurationQueryLatestImageVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration query latest image versions o k response a status code equal to that given
func (o *ImageConfigurationQueryLatestImageVersionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the image configuration query latest image versions o k response
func (o *ImageConfigurationQueryLatestImageVersionsOK) Code() int {
	return 200
}

func (o *ImageConfigurationQueryLatestImageVersionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] imageConfigurationQueryLatestImageVersionsOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationQueryLatestImageVersionsOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] imageConfigurationQueryLatestImageVersionsOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationQueryLatestImageVersionsOK) GetPayload() *models.Images {
	return o.Payload
}

func (o *ImageConfigurationQueryLatestImageVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Images)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryLatestImageVersionsBadRequest creates a ImageConfigurationQueryLatestImageVersionsBadRequest with default headers values
func NewImageConfigurationQueryLatestImageVersionsBadRequest() *ImageConfigurationQueryLatestImageVersionsBadRequest {
	return &ImageConfigurationQueryLatestImageVersionsBadRequest{}
}

/*
ImageConfigurationQueryLatestImageVersionsBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type ImageConfigurationQueryLatestImageVersionsBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration query latest image versions bad request response has a 2xx status code
func (o *ImageConfigurationQueryLatestImageVersionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration query latest image versions bad request response has a 3xx status code
func (o *ImageConfigurationQueryLatestImageVersionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query latest image versions bad request response has a 4xx status code
func (o *ImageConfigurationQueryLatestImageVersionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration query latest image versions bad request response has a 5xx status code
func (o *ImageConfigurationQueryLatestImageVersionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration query latest image versions bad request response a status code equal to that given
func (o *ImageConfigurationQueryLatestImageVersionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the image configuration query latest image versions bad request response
func (o *ImageConfigurationQueryLatestImageVersionsBadRequest) Code() int {
	return 400
}

func (o *ImageConfigurationQueryLatestImageVersionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] imageConfigurationQueryLatestImageVersionsBadRequest  %+v", 400, o.Payload)
}

func (o *ImageConfigurationQueryLatestImageVersionsBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] imageConfigurationQueryLatestImageVersionsBadRequest  %+v", 400, o.Payload)
}

func (o *ImageConfigurationQueryLatestImageVersionsBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryLatestImageVersionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryLatestImageVersionsUnauthorized creates a ImageConfigurationQueryLatestImageVersionsUnauthorized with default headers values
func NewImageConfigurationQueryLatestImageVersionsUnauthorized() *ImageConfigurationQueryLatestImageVersionsUnauthorized {
	return &ImageConfigurationQueryLatestImageVersionsUnauthorized{}
}

/*
ImageConfigurationQueryLatestImageVersionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationQueryLatestImageVersionsUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration query latest image versions unauthorized response has a 2xx status code
func (o *ImageConfigurationQueryLatestImageVersionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration query latest image versions unauthorized response has a 3xx status code
func (o *ImageConfigurationQueryLatestImageVersionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query latest image versions unauthorized response has a 4xx status code
func (o *ImageConfigurationQueryLatestImageVersionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration query latest image versions unauthorized response has a 5xx status code
func (o *ImageConfigurationQueryLatestImageVersionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration query latest image versions unauthorized response a status code equal to that given
func (o *ImageConfigurationQueryLatestImageVersionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the image configuration query latest image versions unauthorized response
func (o *ImageConfigurationQueryLatestImageVersionsUnauthorized) Code() int {
	return 401
}

func (o *ImageConfigurationQueryLatestImageVersionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] imageConfigurationQueryLatestImageVersionsUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationQueryLatestImageVersionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] imageConfigurationQueryLatestImageVersionsUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationQueryLatestImageVersionsUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryLatestImageVersionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryLatestImageVersionsForbidden creates a ImageConfigurationQueryLatestImageVersionsForbidden with default headers values
func NewImageConfigurationQueryLatestImageVersionsForbidden() *ImageConfigurationQueryLatestImageVersionsForbidden {
	return &ImageConfigurationQueryLatestImageVersionsForbidden{}
}

/*
ImageConfigurationQueryLatestImageVersionsForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationQueryLatestImageVersionsForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration query latest image versions forbidden response has a 2xx status code
func (o *ImageConfigurationQueryLatestImageVersionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration query latest image versions forbidden response has a 3xx status code
func (o *ImageConfigurationQueryLatestImageVersionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query latest image versions forbidden response has a 4xx status code
func (o *ImageConfigurationQueryLatestImageVersionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration query latest image versions forbidden response has a 5xx status code
func (o *ImageConfigurationQueryLatestImageVersionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration query latest image versions forbidden response a status code equal to that given
func (o *ImageConfigurationQueryLatestImageVersionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the image configuration query latest image versions forbidden response
func (o *ImageConfigurationQueryLatestImageVersionsForbidden) Code() int {
	return 403
}

func (o *ImageConfigurationQueryLatestImageVersionsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] imageConfigurationQueryLatestImageVersionsForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationQueryLatestImageVersionsForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] imageConfigurationQueryLatestImageVersionsForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationQueryLatestImageVersionsForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryLatestImageVersionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryLatestImageVersionsInternalServerError creates a ImageConfigurationQueryLatestImageVersionsInternalServerError with default headers values
func NewImageConfigurationQueryLatestImageVersionsInternalServerError() *ImageConfigurationQueryLatestImageVersionsInternalServerError {
	return &ImageConfigurationQueryLatestImageVersionsInternalServerError{}
}

/*
ImageConfigurationQueryLatestImageVersionsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationQueryLatestImageVersionsInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration query latest image versions internal server error response has a 2xx status code
func (o *ImageConfigurationQueryLatestImageVersionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration query latest image versions internal server error response has a 3xx status code
func (o *ImageConfigurationQueryLatestImageVersionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query latest image versions internal server error response has a 4xx status code
func (o *ImageConfigurationQueryLatestImageVersionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration query latest image versions internal server error response has a 5xx status code
func (o *ImageConfigurationQueryLatestImageVersionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration query latest image versions internal server error response a status code equal to that given
func (o *ImageConfigurationQueryLatestImageVersionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the image configuration query latest image versions internal server error response
func (o *ImageConfigurationQueryLatestImageVersionsInternalServerError) Code() int {
	return 500
}

func (o *ImageConfigurationQueryLatestImageVersionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] imageConfigurationQueryLatestImageVersionsInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationQueryLatestImageVersionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] imageConfigurationQueryLatestImageVersionsInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationQueryLatestImageVersionsInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryLatestImageVersionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryLatestImageVersionsGatewayTimeout creates a ImageConfigurationQueryLatestImageVersionsGatewayTimeout with default headers values
func NewImageConfigurationQueryLatestImageVersionsGatewayTimeout() *ImageConfigurationQueryLatestImageVersionsGatewayTimeout {
	return &ImageConfigurationQueryLatestImageVersionsGatewayTimeout{}
}

/*
ImageConfigurationQueryLatestImageVersionsGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationQueryLatestImageVersionsGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration query latest image versions gateway timeout response has a 2xx status code
func (o *ImageConfigurationQueryLatestImageVersionsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration query latest image versions gateway timeout response has a 3xx status code
func (o *ImageConfigurationQueryLatestImageVersionsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration query latest image versions gateway timeout response has a 4xx status code
func (o *ImageConfigurationQueryLatestImageVersionsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration query latest image versions gateway timeout response has a 5xx status code
func (o *ImageConfigurationQueryLatestImageVersionsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration query latest image versions gateway timeout response a status code equal to that given
func (o *ImageConfigurationQueryLatestImageVersionsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the image configuration query latest image versions gateway timeout response
func (o *ImageConfigurationQueryLatestImageVersionsGatewayTimeout) Code() int {
	return 504
}

func (o *ImageConfigurationQueryLatestImageVersionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] imageConfigurationQueryLatestImageVersionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationQueryLatestImageVersionsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] imageConfigurationQueryLatestImageVersionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationQueryLatestImageVersionsGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationQueryLatestImageVersionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationQueryLatestImageVersionsDefault creates a ImageConfigurationQueryLatestImageVersionsDefault with default headers values
func NewImageConfigurationQueryLatestImageVersionsDefault(code int) *ImageConfigurationQueryLatestImageVersionsDefault {
	return &ImageConfigurationQueryLatestImageVersionsDefault{
		_statusCode: code,
	}
}

/*
ImageConfigurationQueryLatestImageVersionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationQueryLatestImageVersionsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this image configuration query latest image versions default response has a 2xx status code
func (o *ImageConfigurationQueryLatestImageVersionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this image configuration query latest image versions default response has a 3xx status code
func (o *ImageConfigurationQueryLatestImageVersionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this image configuration query latest image versions default response has a 4xx status code
func (o *ImageConfigurationQueryLatestImageVersionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this image configuration query latest image versions default response has a 5xx status code
func (o *ImageConfigurationQueryLatestImageVersionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this image configuration query latest image versions default response a status code equal to that given
func (o *ImageConfigurationQueryLatestImageVersionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the image configuration query latest image versions default response
func (o *ImageConfigurationQueryLatestImageVersionsDefault) Code() int {
	return o._statusCode
}

func (o *ImageConfigurationQueryLatestImageVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] ImageConfiguration_QueryLatestImageVersions default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationQueryLatestImageVersionsDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/images/baseos/latest][%d] ImageConfiguration_QueryLatestImageVersions default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationQueryLatestImageVersionsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationQueryLatestImageVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}