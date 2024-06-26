package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xyuria-zededa/terraform-provider-zedcloud/v2/models"
)

// EdgeApplicationConfigurationDeleteEdgeApplicationBundleReader is a Reader for the EdgeApplicationConfigurationDeleteEdgeApplicationBundle structure.
type EdgeApplicationConfigurationDeleteEdgeApplicationBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleOK creates a EdgeApplicationConfigurationDeleteEdgeApplicationBundleOK with default headers values
func NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleOK() *EdgeApplicationConfigurationDeleteEdgeApplicationBundleOK {
	return &EdgeApplicationConfigurationDeleteEdgeApplicationBundleOK{}
}

/*
EdgeApplicationConfigurationDeleteEdgeApplicationBundleOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationConfigurationDeleteEdgeApplicationBundleOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration delete edge application bundle o k response has a 2xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application configuration delete edge application bundle o k response has a 3xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration delete edge application bundle o k response has a 4xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration delete edge application bundle o k response has a 5xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration delete edge application bundle o k response a status code equal to that given
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge application configuration delete edge application bundle o k response
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleOK) Code() int {
	return 200
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/id/{id}][%d] edgeApplicationConfigurationDeleteEdgeApplicationBundleOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleOK) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/id/{id}][%d] edgeApplicationConfigurationDeleteEdgeApplicationBundleOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized creates a EdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized with default headers values
func NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized() *EdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized {
	return &EdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized{}
}

/*
EdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration delete edge application bundle unauthorized response has a 2xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration delete edge application bundle unauthorized response has a 3xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration delete edge application bundle unauthorized response has a 4xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration delete edge application bundle unauthorized response has a 5xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration delete edge application bundle unauthorized response a status code equal to that given
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge application configuration delete edge application bundle unauthorized response
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized) Code() int {
	return 401
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/id/{id}][%d] edgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/id/{id}][%d] edgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden creates a EdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden with default headers values
func NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden() *EdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden {
	return &EdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden{}
}

/*
EdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration delete edge application bundle forbidden response has a 2xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration delete edge application bundle forbidden response has a 3xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration delete edge application bundle forbidden response has a 4xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration delete edge application bundle forbidden response has a 5xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration delete edge application bundle forbidden response a status code equal to that given
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge application configuration delete edge application bundle forbidden response
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden) Code() int {
	return 403
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/id/{id}][%d] edgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/id/{id}][%d] edgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleNotFound creates a EdgeApplicationConfigurationDeleteEdgeApplicationBundleNotFound with default headers values
func NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleNotFound() *ApplicationNotFound {
	return &ApplicationNotFound{}
}

/*
ApplicationNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type ApplicationNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration delete edge application bundle not found response has a 2xx status code
func (o *ApplicationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration delete edge application bundle not found response has a 3xx status code
func (o *ApplicationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration delete edge application bundle not found response has a 4xx status code
func (o *ApplicationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration delete edge application bundle not found response has a 5xx status code
func (o *ApplicationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration delete edge application bundle not found response a status code equal to that given
func (o *ApplicationNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge application configuration delete edge application bundle not found response
func (o *ApplicationNotFound) Code() int {
	return 404
}

func (o *ApplicationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/id/{id}][%d] edgeApplicationConfigurationDeleteEdgeApplicationBundleNotFound  %+v", 404, o.Payload)
}

func (o *ApplicationNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/id/{id}][%d] edgeApplicationConfigurationDeleteEdgeApplicationBundleNotFound  %+v", 404, o.Payload)
}

func (o *ApplicationNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ApplicationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict creates a EdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict with default headers values
func NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict() *EdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict {
	return &EdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict{}
}

/*
EdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because there are instances of this edge application deployed in edge node(s)
*/
type EdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration delete edge application bundle conflict response has a 2xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration delete edge application bundle conflict response has a 3xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration delete edge application bundle conflict response has a 4xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration delete edge application bundle conflict response has a 5xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration delete edge application bundle conflict response a status code equal to that given
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the edge application configuration delete edge application bundle conflict response
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict) Code() int {
	return 409
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/id/{id}][%d] edgeApplicationConfigurationDeleteEdgeApplicationBundleConflict  %+v", 409, o.Payload)
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/id/{id}][%d] edgeApplicationConfigurationDeleteEdgeApplicationBundleConflict  %+v", 409, o.Payload)
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError creates a EdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError with default headers values
func NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError() *EdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError {
	return &EdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError{}
}

/*
EdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration delete edge application bundle internal server error response has a 2xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration delete edge application bundle internal server error response has a 3xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration delete edge application bundle internal server error response has a 4xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration delete edge application bundle internal server error response has a 5xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application configuration delete edge application bundle internal server error response a status code equal to that given
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge application configuration delete edge application bundle internal server error response
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError) Code() int {
	return 500
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/id/{id}][%d] edgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/id/{id}][%d] edgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout creates a EdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout with default headers values
func NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout() *EdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout {
	return &EdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout{}
}

/*
EdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration delete edge application bundle gateway timeout response has a 2xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration delete edge application bundle gateway timeout response has a 3xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration delete edge application bundle gateway timeout response has a 4xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration delete edge application bundle gateway timeout response has a 5xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application configuration delete edge application bundle gateway timeout response a status code equal to that given
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge application configuration delete edge application bundle gateway timeout response
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/id/{id}][%d] edgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/id/{id}][%d] edgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault creates a EdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault with default headers values
func NewEdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault(code int) *EdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault {
	return &EdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge application configuration delete edge application bundle default response has a 2xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application configuration delete edge application bundle default response has a 3xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application configuration delete edge application bundle default response has a 4xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application configuration delete edge application bundle default response has a 5xx status code
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application configuration delete edge application bundle default response a status code equal to that given
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge application configuration delete edge application bundle default response
func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/apps/id/{id}][%d] EdgeApplicationConfiguration_DeleteEdgeApplicationBundle default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/apps/id/{id}][%d] EdgeApplicationConfiguration_DeleteEdgeApplicationBundle default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationConfigurationDeleteEdgeApplicationBundleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
