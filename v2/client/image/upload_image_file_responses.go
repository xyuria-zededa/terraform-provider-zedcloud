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

// ImageConfigurationUploadImageFileReader is a Reader for the ImageConfigurationUploadImageFile structure.
type ImageConfigurationUploadImageFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageConfigurationUploadImageFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageConfigurationUploadImageFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewImageConfigurationUploadImageFileAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewImageConfigurationUploadImageFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewImageConfigurationUploadImageFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewImageConfigurationUploadImageFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewImageConfigurationUploadImageFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewImageConfigurationUploadImageFileConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageConfigurationUploadImageFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewImageConfigurationUploadImageFileGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewImageConfigurationUploadImageFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewImageConfigurationUploadImageFileOK creates a ImageConfigurationUploadImageFileOK with default headers values
func NewImageConfigurationUploadImageFileOK() *ImageConfigurationUploadImageFileOK {
	return &ImageConfigurationUploadImageFileOK{}
}

/*
ImageConfigurationUploadImageFileOK describes a response with status code 200, with default header values.

A successful response.
*/
type ImageConfigurationUploadImageFileOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image file o k response has a 2xx status code
func (o *ImageConfigurationUploadImageFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image configuration upload image file o k response has a 3xx status code
func (o *ImageConfigurationUploadImageFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image file o k response has a 4xx status code
func (o *ImageConfigurationUploadImageFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration upload image file o k response has a 5xx status code
func (o *ImageConfigurationUploadImageFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration upload image file o k response a status code equal to that given
func (o *ImageConfigurationUploadImageFileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the image configuration upload image file o k response
func (o *ImageConfigurationUploadImageFileOK) Code() int {
	return 200
}

func (o *ImageConfigurationUploadImageFileOK) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationUploadImageFileOK) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileOK  %+v", 200, o.Payload)
}

func (o *ImageConfigurationUploadImageFileOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageFileAccepted creates a ImageConfigurationUploadImageFileAccepted with default headers values
func NewImageConfigurationUploadImageFileAccepted() *ImageConfigurationUploadImageFileAccepted {
	return &ImageConfigurationUploadImageFileAccepted{}
}

/*
ImageConfigurationUploadImageFileAccepted describes a response with status code 202, with default header values.

Accepted. The API gateway accepted the request for uploading but the upload process has not been completed. Please check ImageStatus and ImageError fields to track the status of upload process and any error messages.
*/
type ImageConfigurationUploadImageFileAccepted struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image file accepted response has a 2xx status code
func (o *ImageConfigurationUploadImageFileAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this image configuration upload image file accepted response has a 3xx status code
func (o *ImageConfigurationUploadImageFileAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image file accepted response has a 4xx status code
func (o *ImageConfigurationUploadImageFileAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration upload image file accepted response has a 5xx status code
func (o *ImageConfigurationUploadImageFileAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration upload image file accepted response a status code equal to that given
func (o *ImageConfigurationUploadImageFileAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the image configuration upload image file accepted response
func (o *ImageConfigurationUploadImageFileAccepted) Code() int {
	return 202
}

func (o *ImageConfigurationUploadImageFileAccepted) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileAccepted  %+v", 202, o.Payload)
}

func (o *ImageConfigurationUploadImageFileAccepted) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileAccepted  %+v", 202, o.Payload)
}

func (o *ImageConfigurationUploadImageFileAccepted) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageFileAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageFileBadRequest creates a ImageConfigurationUploadImageFileBadRequest with default headers values
func NewImageConfigurationUploadImageFileBadRequest() *ImageConfigurationUploadImageFileBadRequest {
	return &ImageConfigurationUploadImageFileBadRequest{}
}

/*
ImageConfigurationUploadImageFileBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type ImageConfigurationUploadImageFileBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image file bad request response has a 2xx status code
func (o *ImageConfigurationUploadImageFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration upload image file bad request response has a 3xx status code
func (o *ImageConfigurationUploadImageFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image file bad request response has a 4xx status code
func (o *ImageConfigurationUploadImageFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration upload image file bad request response has a 5xx status code
func (o *ImageConfigurationUploadImageFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration upload image file bad request response a status code equal to that given
func (o *ImageConfigurationUploadImageFileBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the image configuration upload image file bad request response
func (o *ImageConfigurationUploadImageFileBadRequest) Code() int {
	return 400
}

func (o *ImageConfigurationUploadImageFileBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileBadRequest  %+v", 400, o.Payload)
}

func (o *ImageConfigurationUploadImageFileBadRequest) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileBadRequest  %+v", 400, o.Payload)
}

func (o *ImageConfigurationUploadImageFileBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageFileUnauthorized creates a ImageConfigurationUploadImageFileUnauthorized with default headers values
func NewImageConfigurationUploadImageFileUnauthorized() *ImageConfigurationUploadImageFileUnauthorized {
	return &ImageConfigurationUploadImageFileUnauthorized{}
}

/*
ImageConfigurationUploadImageFileUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ImageConfigurationUploadImageFileUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image file unauthorized response has a 2xx status code
func (o *ImageConfigurationUploadImageFileUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration upload image file unauthorized response has a 3xx status code
func (o *ImageConfigurationUploadImageFileUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image file unauthorized response has a 4xx status code
func (o *ImageConfigurationUploadImageFileUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration upload image file unauthorized response has a 5xx status code
func (o *ImageConfigurationUploadImageFileUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration upload image file unauthorized response a status code equal to that given
func (o *ImageConfigurationUploadImageFileUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the image configuration upload image file unauthorized response
func (o *ImageConfigurationUploadImageFileUnauthorized) Code() int {
	return 401
}

func (o *ImageConfigurationUploadImageFileUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationUploadImageFileUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileUnauthorized  %+v", 401, o.Payload)
}

func (o *ImageConfigurationUploadImageFileUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageFileForbidden creates a ImageConfigurationUploadImageFileForbidden with default headers values
func NewImageConfigurationUploadImageFileForbidden() *ImageConfigurationUploadImageFileForbidden {
	return &ImageConfigurationUploadImageFileForbidden{}
}

/*
ImageConfigurationUploadImageFileForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ImageConfigurationUploadImageFileForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image file forbidden response has a 2xx status code
func (o *ImageConfigurationUploadImageFileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration upload image file forbidden response has a 3xx status code
func (o *ImageConfigurationUploadImageFileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image file forbidden response has a 4xx status code
func (o *ImageConfigurationUploadImageFileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration upload image file forbidden response has a 5xx status code
func (o *ImageConfigurationUploadImageFileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration upload image file forbidden response a status code equal to that given
func (o *ImageConfigurationUploadImageFileForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the image configuration upload image file forbidden response
func (o *ImageConfigurationUploadImageFileForbidden) Code() int {
	return 403
}

func (o *ImageConfigurationUploadImageFileForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationUploadImageFileForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileForbidden  %+v", 403, o.Payload)
}

func (o *ImageConfigurationUploadImageFileForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageFileNotFound creates a ImageConfigurationUploadImageFileNotFound with default headers values
func NewImageConfigurationUploadImageFileNotFound() *ImageConfigurationUploadImageFileNotFound {
	return &ImageConfigurationUploadImageFileNotFound{}
}

/*
ImageConfigurationUploadImageFileNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ImageConfigurationUploadImageFileNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image file not found response has a 2xx status code
func (o *ImageConfigurationUploadImageFileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration upload image file not found response has a 3xx status code
func (o *ImageConfigurationUploadImageFileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image file not found response has a 4xx status code
func (o *ImageConfigurationUploadImageFileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration upload image file not found response has a 5xx status code
func (o *ImageConfigurationUploadImageFileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration upload image file not found response a status code equal to that given
func (o *ImageConfigurationUploadImageFileNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the image configuration upload image file not found response
func (o *ImageConfigurationUploadImageFileNotFound) Code() int {
	return 404
}

func (o *ImageConfigurationUploadImageFileNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileNotFound  %+v", 404, o.Payload)
}

func (o *ImageConfigurationUploadImageFileNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileNotFound  %+v", 404, o.Payload)
}

func (o *ImageConfigurationUploadImageFileNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageFileConflict creates a ImageConfigurationUploadImageFileConflict with default headers values
func NewImageConfigurationUploadImageFileConflict() *ImageConfigurationUploadImageFileConflict {
	return &ImageConfigurationUploadImageFileConflict{}
}

/*
ImageConfigurationUploadImageFileConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because
-- another request for uplink / upload is already in progress
-- image has been already uploaded, can't be modified again
*/
type ImageConfigurationUploadImageFileConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image file conflict response has a 2xx status code
func (o *ImageConfigurationUploadImageFileConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration upload image file conflict response has a 3xx status code
func (o *ImageConfigurationUploadImageFileConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image file conflict response has a 4xx status code
func (o *ImageConfigurationUploadImageFileConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this image configuration upload image file conflict response has a 5xx status code
func (o *ImageConfigurationUploadImageFileConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this image configuration upload image file conflict response a status code equal to that given
func (o *ImageConfigurationUploadImageFileConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the image configuration upload image file conflict response
func (o *ImageConfigurationUploadImageFileConflict) Code() int {
	return 409
}

func (o *ImageConfigurationUploadImageFileConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileConflict  %+v", 409, o.Payload)
}

func (o *ImageConfigurationUploadImageFileConflict) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileConflict  %+v", 409, o.Payload)
}

func (o *ImageConfigurationUploadImageFileConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageFileConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageFileInternalServerError creates a ImageConfigurationUploadImageFileInternalServerError with default headers values
func NewImageConfigurationUploadImageFileInternalServerError() *ImageConfigurationUploadImageFileInternalServerError {
	return &ImageConfigurationUploadImageFileInternalServerError{}
}

/*
ImageConfigurationUploadImageFileInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ImageConfigurationUploadImageFileInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image file internal server error response has a 2xx status code
func (o *ImageConfigurationUploadImageFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration upload image file internal server error response has a 3xx status code
func (o *ImageConfigurationUploadImageFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image file internal server error response has a 4xx status code
func (o *ImageConfigurationUploadImageFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration upload image file internal server error response has a 5xx status code
func (o *ImageConfigurationUploadImageFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration upload image file internal server error response a status code equal to that given
func (o *ImageConfigurationUploadImageFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the image configuration upload image file internal server error response
func (o *ImageConfigurationUploadImageFileInternalServerError) Code() int {
	return 500
}

func (o *ImageConfigurationUploadImageFileInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationUploadImageFileInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileInternalServerError  %+v", 500, o.Payload)
}

func (o *ImageConfigurationUploadImageFileInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageFileGatewayTimeout creates a ImageConfigurationUploadImageFileGatewayTimeout with default headers values
func NewImageConfigurationUploadImageFileGatewayTimeout() *ImageConfigurationUploadImageFileGatewayTimeout {
	return &ImageConfigurationUploadImageFileGatewayTimeout{}
}

/*
ImageConfigurationUploadImageFileGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ImageConfigurationUploadImageFileGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this image configuration upload image file gateway timeout response has a 2xx status code
func (o *ImageConfigurationUploadImageFileGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this image configuration upload image file gateway timeout response has a 3xx status code
func (o *ImageConfigurationUploadImageFileGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this image configuration upload image file gateway timeout response has a 4xx status code
func (o *ImageConfigurationUploadImageFileGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this image configuration upload image file gateway timeout response has a 5xx status code
func (o *ImageConfigurationUploadImageFileGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this image configuration upload image file gateway timeout response a status code equal to that given
func (o *ImageConfigurationUploadImageFileGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the image configuration upload image file gateway timeout response
func (o *ImageConfigurationUploadImageFileGatewayTimeout) Code() int {
	return 504
}

func (o *ImageConfigurationUploadImageFileGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationUploadImageFileGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] imageConfigurationUploadImageFileGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ImageConfigurationUploadImageFileGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ImageConfigurationUploadImageFileGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageConfigurationUploadImageFileDefault creates a ImageConfigurationUploadImageFileDefault with default headers values
func NewImageConfigurationUploadImageFileDefault(code int) *ImageConfigurationUploadImageFileDefault {
	return &ImageConfigurationUploadImageFileDefault{
		_statusCode: code,
	}
}

/*
ImageConfigurationUploadImageFileDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ImageConfigurationUploadImageFileDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this image configuration upload image file default response has a 2xx status code
func (o *ImageConfigurationUploadImageFileDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this image configuration upload image file default response has a 3xx status code
func (o *ImageConfigurationUploadImageFileDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this image configuration upload image file default response has a 4xx status code
func (o *ImageConfigurationUploadImageFileDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this image configuration upload image file default response has a 5xx status code
func (o *ImageConfigurationUploadImageFileDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this image configuration upload image file default response a status code equal to that given
func (o *ImageConfigurationUploadImageFileDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the image configuration upload image file default response
func (o *ImageConfigurationUploadImageFileDefault) Code() int {
	return o._statusCode
}

func (o *ImageConfigurationUploadImageFileDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] ImageConfiguration_UploadImageFile default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationUploadImageFileDefault) String() string {
	return fmt.Sprintf("[PUT /v1/apps/images/name/{name}/upload/file][%d] ImageConfiguration_UploadImageFile default  %+v", o._statusCode, o.Payload)
}

func (o *ImageConfigurationUploadImageFileDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ImageConfigurationUploadImageFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
