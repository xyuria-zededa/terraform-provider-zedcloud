// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xyuria-zededa/terraform-provider-zedcloud/v2/models"
)

// IdentityAccessManagementDeleteUserReader is a Reader for the IdentityAccessManagementDeleteUser structure.
type IdentityAccessManagementDeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IdentityAccessManagementDeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIdentityAccessManagementDeleteUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewIdentityAccessManagementDeleteUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewIdentityAccessManagementDeleteUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewIdentityAccessManagementDeleteUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIdentityAccessManagementDeleteUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewIdentityAccessManagementDeleteUserGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIdentityAccessManagementDeleteUserDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIdentityAccessManagementDeleteUserOK creates a IdentityAccessManagementDeleteUserOK with default headers values
func NewIdentityAccessManagementDeleteUserOK() *IdentityAccessManagementDeleteUserOK {
	return &IdentityAccessManagementDeleteUserOK{}
}

/*
IdentityAccessManagementDeleteUserOK describes a response with status code 200, with default header values.

A successful response.
*/
type IdentityAccessManagementDeleteUserOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete user o k response has a 2xx status code
func (o *IdentityAccessManagementDeleteUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this identity access management delete user o k response has a 3xx status code
func (o *IdentityAccessManagementDeleteUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete user o k response has a 4xx status code
func (o *IdentityAccessManagementDeleteUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management delete user o k response has a 5xx status code
func (o *IdentityAccessManagementDeleteUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete user o k response a status code equal to that given
func (o *IdentityAccessManagementDeleteUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the identity access management delete user o k response
func (o *IdentityAccessManagementDeleteUserOK) Code() int {
	return 200
}

func (o *IdentityAccessManagementDeleteUserOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/users/id/{id}][%d] identityAccessManagementDeleteUserOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementDeleteUserOK) String() string {
	return fmt.Sprintf("[DELETE /v1/users/id/{id}][%d] identityAccessManagementDeleteUserOK  %+v", 200, o.Payload)
}

func (o *IdentityAccessManagementDeleteUserOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteUserUnauthorized creates a IdentityAccessManagementDeleteUserUnauthorized with default headers values
func NewIdentityAccessManagementDeleteUserUnauthorized() *IdentityAccessManagementDeleteUserUnauthorized {
	return &IdentityAccessManagementDeleteUserUnauthorized{}
}

/*
IdentityAccessManagementDeleteUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type IdentityAccessManagementDeleteUserUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete user unauthorized response has a 2xx status code
func (o *IdentityAccessManagementDeleteUserUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete user unauthorized response has a 3xx status code
func (o *IdentityAccessManagementDeleteUserUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete user unauthorized response has a 4xx status code
func (o *IdentityAccessManagementDeleteUserUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management delete user unauthorized response has a 5xx status code
func (o *IdentityAccessManagementDeleteUserUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete user unauthorized response a status code equal to that given
func (o *IdentityAccessManagementDeleteUserUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the identity access management delete user unauthorized response
func (o *IdentityAccessManagementDeleteUserUnauthorized) Code() int {
	return 401
}

func (o *IdentityAccessManagementDeleteUserUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/users/id/{id}][%d] identityAccessManagementDeleteUserUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementDeleteUserUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /v1/users/id/{id}][%d] identityAccessManagementDeleteUserUnauthorized  %+v", 401, o.Payload)
}

func (o *IdentityAccessManagementDeleteUserUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteUserForbidden creates a IdentityAccessManagementDeleteUserForbidden with default headers values
func NewIdentityAccessManagementDeleteUserForbidden() *IdentityAccessManagementDeleteUserForbidden {
	return &IdentityAccessManagementDeleteUserForbidden{}
}

/*
IdentityAccessManagementDeleteUserForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type IdentityAccessManagementDeleteUserForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete user forbidden response has a 2xx status code
func (o *IdentityAccessManagementDeleteUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete user forbidden response has a 3xx status code
func (o *IdentityAccessManagementDeleteUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete user forbidden response has a 4xx status code
func (o *IdentityAccessManagementDeleteUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management delete user forbidden response has a 5xx status code
func (o *IdentityAccessManagementDeleteUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete user forbidden response a status code equal to that given
func (o *IdentityAccessManagementDeleteUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the identity access management delete user forbidden response
func (o *IdentityAccessManagementDeleteUserForbidden) Code() int {
	return 403
}

func (o *IdentityAccessManagementDeleteUserForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/users/id/{id}][%d] identityAccessManagementDeleteUserForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementDeleteUserForbidden) String() string {
	return fmt.Sprintf("[DELETE /v1/users/id/{id}][%d] identityAccessManagementDeleteUserForbidden  %+v", 403, o.Payload)
}

func (o *IdentityAccessManagementDeleteUserForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteUserNotFound creates a IdentityAccessManagementDeleteUserNotFound with default headers values
func NewIdentityAccessManagementDeleteUserNotFound() *IdentityAccessManagementDeleteUserNotFound {
	return &IdentityAccessManagementDeleteUserNotFound{}
}

/*
IdentityAccessManagementDeleteUserNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type IdentityAccessManagementDeleteUserNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete user not found response has a 2xx status code
func (o *IdentityAccessManagementDeleteUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete user not found response has a 3xx status code
func (o *IdentityAccessManagementDeleteUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete user not found response has a 4xx status code
func (o *IdentityAccessManagementDeleteUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this identity access management delete user not found response has a 5xx status code
func (o *IdentityAccessManagementDeleteUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this identity access management delete user not found response a status code equal to that given
func (o *IdentityAccessManagementDeleteUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the identity access management delete user not found response
func (o *IdentityAccessManagementDeleteUserNotFound) Code() int {
	return 404
}

func (o *IdentityAccessManagementDeleteUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/users/id/{id}][%d] identityAccessManagementDeleteUserNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementDeleteUserNotFound) String() string {
	return fmt.Sprintf("[DELETE /v1/users/id/{id}][%d] identityAccessManagementDeleteUserNotFound  %+v", 404, o.Payload)
}

func (o *IdentityAccessManagementDeleteUserNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteUserInternalServerError creates a IdentityAccessManagementDeleteUserInternalServerError with default headers values
func NewIdentityAccessManagementDeleteUserInternalServerError() *IdentityAccessManagementDeleteUserInternalServerError {
	return &IdentityAccessManagementDeleteUserInternalServerError{}
}

/*
IdentityAccessManagementDeleteUserInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type IdentityAccessManagementDeleteUserInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete user internal server error response has a 2xx status code
func (o *IdentityAccessManagementDeleteUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete user internal server error response has a 3xx status code
func (o *IdentityAccessManagementDeleteUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete user internal server error response has a 4xx status code
func (o *IdentityAccessManagementDeleteUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management delete user internal server error response has a 5xx status code
func (o *IdentityAccessManagementDeleteUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management delete user internal server error response a status code equal to that given
func (o *IdentityAccessManagementDeleteUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the identity access management delete user internal server error response
func (o *IdentityAccessManagementDeleteUserInternalServerError) Code() int {
	return 500
}

func (o *IdentityAccessManagementDeleteUserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/users/id/{id}][%d] identityAccessManagementDeleteUserInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementDeleteUserInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /v1/users/id/{id}][%d] identityAccessManagementDeleteUserInternalServerError  %+v", 500, o.Payload)
}

func (o *IdentityAccessManagementDeleteUserInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteUserGatewayTimeout creates a IdentityAccessManagementDeleteUserGatewayTimeout with default headers values
func NewIdentityAccessManagementDeleteUserGatewayTimeout() *IdentityAccessManagementDeleteUserGatewayTimeout {
	return &IdentityAccessManagementDeleteUserGatewayTimeout{}
}

/*
IdentityAccessManagementDeleteUserGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type IdentityAccessManagementDeleteUserGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this identity access management delete user gateway timeout response has a 2xx status code
func (o *IdentityAccessManagementDeleteUserGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this identity access management delete user gateway timeout response has a 3xx status code
func (o *IdentityAccessManagementDeleteUserGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this identity access management delete user gateway timeout response has a 4xx status code
func (o *IdentityAccessManagementDeleteUserGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this identity access management delete user gateway timeout response has a 5xx status code
func (o *IdentityAccessManagementDeleteUserGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this identity access management delete user gateway timeout response a status code equal to that given
func (o *IdentityAccessManagementDeleteUserGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the identity access management delete user gateway timeout response
func (o *IdentityAccessManagementDeleteUserGatewayTimeout) Code() int {
	return 504
}

func (o *IdentityAccessManagementDeleteUserGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /v1/users/id/{id}][%d] identityAccessManagementDeleteUserGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementDeleteUserGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /v1/users/id/{id}][%d] identityAccessManagementDeleteUserGatewayTimeout  %+v", 504, o.Payload)
}

func (o *IdentityAccessManagementDeleteUserGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteUserGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIdentityAccessManagementDeleteUserDefault creates a IdentityAccessManagementDeleteUserDefault with default headers values
func NewIdentityAccessManagementDeleteUserDefault(code int) *IdentityAccessManagementDeleteUserDefault {
	return &IdentityAccessManagementDeleteUserDefault{
		_statusCode: code,
	}
}

/*
IdentityAccessManagementDeleteUserDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type IdentityAccessManagementDeleteUserDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this identity access management delete user default response has a 2xx status code
func (o *IdentityAccessManagementDeleteUserDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this identity access management delete user default response has a 3xx status code
func (o *IdentityAccessManagementDeleteUserDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this identity access management delete user default response has a 4xx status code
func (o *IdentityAccessManagementDeleteUserDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this identity access management delete user default response has a 5xx status code
func (o *IdentityAccessManagementDeleteUserDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this identity access management delete user default response a status code equal to that given
func (o *IdentityAccessManagementDeleteUserDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the identity access management delete user default response
func (o *IdentityAccessManagementDeleteUserDefault) Code() int {
	return o._statusCode
}

func (o *IdentityAccessManagementDeleteUserDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/users/id/{id}][%d] IdentityAccessManagement_DeleteUser default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementDeleteUserDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/users/id/{id}][%d] IdentityAccessManagement_DeleteUser default  %+v", o._statusCode, o.Payload)
}

func (o *IdentityAccessManagementDeleteUserDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *IdentityAccessManagementDeleteUserDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
