package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xyuria-zededa/terraform-provider-zedcloud/v2/models"
)

// ProjectsCreateReader is a Reader for the ProjectsCreate structure.
type ProjectsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewProjectsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewProjectsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewProjectsCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewProjectsCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewProjectsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewProjectsCreateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewProjectsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectsCreateOK creates a ProjectsCreateOK with default headers values
func NewProjectsCreateOK() *ProjectsCreateOK {
	return &ProjectsCreateOK{}
}

/*
ProjectsCreateOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProjectsCreateOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this projects create o k response has a 2xx status code
func (o *ProjectsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this projects create o k response has a 3xx status code
func (o *ProjectsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects create o k response has a 4xx status code
func (o *ProjectsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects create o k response has a 5xx status code
func (o *ProjectsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this projects create o k response a status code equal to that given
func (o *ProjectsCreateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ProjectsCreateOK) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] projectsCreateOK  %+v", 200, o.Payload)
}

func (o *ProjectsCreateOK) String() string {
	return fmt.Sprintf("[POST /v1/projects][%d] projectsCreateOK  %+v", 200, o.Payload)
}

func (o *ProjectsCreateOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ProjectsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsCreateBadRequest creates a ProjectsCreateBadRequest with default headers values
func NewProjectsCreateBadRequest() *ProjectsCreateBadRequest {
	return &ProjectsCreateBadRequest{}
}

/*
ProjectsCreateBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of missing parameter or invalid value of parameters.
*/
type ProjectsCreateBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this projects create bad request response has a 2xx status code
func (o *ProjectsCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects create bad request response has a 3xx status code
func (o *ProjectsCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects create bad request response has a 4xx status code
func (o *ProjectsCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects create bad request response has a 5xx status code
func (o *ProjectsCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this projects create bad request response a status code equal to that given
func (o *ProjectsCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ProjectsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/projects] create project: 400 bad request\n%s", spew.Sdump(o.Payload.Error))
}

func (o *ProjectsCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/projects] create project: 400 bad request\n%s", spew.Sdump(o.Payload.Error))
}

func (o *ProjectsCreateBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ProjectsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsCreateUnauthorized creates a ProjectsCreateUnauthorized with default headers values
func NewProjectsCreateUnauthorized() *ProjectsCreateUnauthorized {
	return &ProjectsCreateUnauthorized{}
}

/*
ProjectsCreateUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type ProjectsCreateUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this projects create unauthorized response has a 2xx status code
func (o *ProjectsCreateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects create unauthorized response has a 3xx status code
func (o *ProjectsCreateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects create unauthorized response has a 4xx status code
func (o *ProjectsCreateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects create unauthorized response has a 5xx status code
func (o *ProjectsCreateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this projects create unauthorized response a status code equal to that given
func (o *ProjectsCreateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ProjectsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] projectsCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsCreateUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/projects][%d] projectsCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *ProjectsCreateUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ProjectsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsCreateForbidden creates a ProjectsCreateForbidden with default headers values
func NewProjectsCreateForbidden() *ProjectsCreateForbidden {
	return &ProjectsCreateForbidden{}
}

/*
ProjectsCreateForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the request or does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type ProjectsCreateForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this projects create forbidden response has a 2xx status code
func (o *ProjectsCreateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects create forbidden response has a 3xx status code
func (o *ProjectsCreateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects create forbidden response has a 4xx status code
func (o *ProjectsCreateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects create forbidden response has a 5xx status code
func (o *ProjectsCreateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this projects create forbidden response a status code equal to that given
func (o *ProjectsCreateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ProjectsCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] projectsCreateForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsCreateForbidden) String() string {
	return fmt.Sprintf("[POST /v1/projects][%d] projectsCreateForbidden  %+v", 403, o.Payload)
}

func (o *ProjectsCreateForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ProjectsCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsCreateConflict creates a ProjectsCreateConflict with default headers values
func NewProjectsCreateConflict() *ProjectsCreateConflict {
	return &ProjectsCreateConflict{}
}

/*
ProjectsCreateConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this edge network record will conflict with an already existing edge network record.
*/
type ProjectsCreateConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this projects create conflict response has a 2xx status code
func (o *ProjectsCreateConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects create conflict response has a 3xx status code
func (o *ProjectsCreateConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects create conflict response has a 4xx status code
func (o *ProjectsCreateConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this projects create conflict response has a 5xx status code
func (o *ProjectsCreateConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this projects create conflict response a status code equal to that given
func (o *ProjectsCreateConflict) IsCode(code int) bool {
	return code == 409
}

func (o *ProjectsCreateConflict) Error() string {
	return fmt.Sprintf("[POST /v1/projects] create project: 409 conflict\n%s", spew.Sdump(o.Payload.Error))
}

func (o *ProjectsCreateConflict) String() string {
	return fmt.Sprintf("[POST /v1/projects] create project: 409 conflict\n%s", spew.Sdump(o.Payload.Error))
}

func (o *ProjectsCreateConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ProjectsCreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsCreateInternalServerError creates a ProjectsCreateInternalServerError with default headers values
func NewProjectsCreateInternalServerError() *ProjectsCreateInternalServerError {
	return &ProjectsCreateInternalServerError{}
}

/*
ProjectsCreateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type ProjectsCreateInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this projects create internal server error response has a 2xx status code
func (o *ProjectsCreateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects create internal server error response has a 3xx status code
func (o *ProjectsCreateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects create internal server error response has a 4xx status code
func (o *ProjectsCreateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects create internal server error response has a 5xx status code
func (o *ProjectsCreateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this projects create internal server error response a status code equal to that given
func (o *ProjectsCreateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ProjectsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] projectsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *ProjectsCreateInternalServerError) String() string {
	return fmt.Sprintf("[POST /v1/projects][%d] projectsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *ProjectsCreateInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ProjectsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsCreateGatewayTimeout creates a ProjectsCreateGatewayTimeout with default headers values
func NewProjectsCreateGatewayTimeout() *ProjectsCreateGatewayTimeout {
	return &ProjectsCreateGatewayTimeout{}
}

/*
ProjectsCreateGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type ProjectsCreateGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this projects create gateway timeout response has a 2xx status code
func (o *ProjectsCreateGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this projects create gateway timeout response has a 3xx status code
func (o *ProjectsCreateGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this projects create gateway timeout response has a 4xx status code
func (o *ProjectsCreateGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this projects create gateway timeout response has a 5xx status code
func (o *ProjectsCreateGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this projects create gateway timeout response a status code equal to that given
func (o *ProjectsCreateGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *ProjectsCreateGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] projectsCreateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ProjectsCreateGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /v1/projects][%d] projectsCreateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *ProjectsCreateGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *ProjectsCreateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsCreateDefault creates a ProjectsCreateDefault with default headers values
func NewProjectsCreateDefault(code int) *ProjectsCreateDefault {
	return &ProjectsCreateDefault{
		_statusCode: code,
	}
}

/*
ProjectsCreateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProjectsCreateDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the projects create default response
func (o *ProjectsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this projects create default response has a 2xx status code
func (o *ProjectsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this projects create default response has a 3xx status code
func (o *ProjectsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this projects create default response has a 4xx status code
func (o *ProjectsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this projects create default response has a 5xx status code
func (o *ProjectsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this projects create default response a status code equal to that given
func (o *ProjectsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ProjectsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] Projects_Create default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectsCreateDefault) String() string {
	return fmt.Sprintf("[POST /v1/projects][%d] Projects_Create default  %+v", o._statusCode, o.Payload)
}

func (o *ProjectsCreateDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ProjectsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
