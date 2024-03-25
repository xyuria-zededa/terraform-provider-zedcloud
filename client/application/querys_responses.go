package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/xyuria-zededa/terraform-provider-zedcloud/models"
)

// EdgeApplicationConfigurationQueryEdgeApplicationBundlesReader is a Reader for the EdgeApplicationConfigurationQueryEdgeApplicationBundles structure.
type EdgeApplicationConfigurationQueryEdgeApplicationBundlesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesOK creates a EdgeApplicationConfigurationQueryEdgeApplicationBundlesOK with default headers values
func NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesOK() *EdgeApplicationConfigurationQueryEdgeApplicationBundlesOK {
	return &EdgeApplicationConfigurationQueryEdgeApplicationBundlesOK{}
}

/*
EdgeApplicationConfigurationQueryEdgeApplicationBundlesOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationConfigurationQueryEdgeApplicationBundlesOK struct {
	Payload *models.Apps
}

// IsSuccess returns true when this edge application configuration query edge application bundles o k response has a 2xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application configuration query edge application bundles o k response has a 3xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration query edge application bundles o k response has a 4xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration query edge application bundles o k response has a 5xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration query edge application bundles o k response a status code equal to that given
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge application configuration query edge application bundles o k response
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesOK) Code() int {
	return 200
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps][%d] edgeApplicationConfigurationQueryEdgeApplicationBundlesOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesOK) String() string {
	return fmt.Sprintf("[GET /v1/apps][%d] edgeApplicationConfigurationQueryEdgeApplicationBundlesOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesOK) GetPayload() *models.Apps {
	return o.Payload
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Apps)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest creates a EdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest with default headers values
func NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest() *EdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest {
	return &EdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest{}
}

/*
EdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest describes a response with status code 400, with default header values.

Bad Request. The API gateway did not process the request because of invalid value of filter parameters.
*/
type EdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration query edge application bundles bad request response has a 2xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration query edge application bundles bad request response has a 3xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration query edge application bundles bad request response has a 4xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration query edge application bundles bad request response has a 5xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration query edge application bundles bad request response a status code equal to that given
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the edge application configuration query edge application bundles bad request response
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest) Code() int {
	return 400
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/apps][%d] edgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/apps][%d] edgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest  %+v", 400, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized creates a EdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized with default headers values
func NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized() *EdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized {
	return &EdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized{}
}

/*
EdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration query edge application bundles unauthorized response has a 2xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration query edge application bundles unauthorized response has a 3xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration query edge application bundles unauthorized response has a 4xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration query edge application bundles unauthorized response has a 5xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration query edge application bundles unauthorized response a status code equal to that given
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge application configuration query edge application bundles unauthorized response
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized) Code() int {
	return 401
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps][%d] edgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps][%d] edgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden creates a EdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden with default headers values
func NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden() *EdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden {
	return &EdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden{}
}

/*
EdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration query edge application bundles forbidden response has a 2xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration query edge application bundles forbidden response has a 3xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration query edge application bundles forbidden response has a 4xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration query edge application bundles forbidden response has a 5xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration query edge application bundles forbidden response a status code equal to that given
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge application configuration query edge application bundles forbidden response
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden) Code() int {
	return 403
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps][%d] edgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps][%d] edgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError creates a EdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError with default headers values
func NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError() *EdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError {
	return &EdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError{}
}

/*
EdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration query edge application bundles internal server error response has a 2xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration query edge application bundles internal server error response has a 3xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration query edge application bundles internal server error response has a 4xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration query edge application bundles internal server error response has a 5xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application configuration query edge application bundles internal server error response a status code equal to that given
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge application configuration query edge application bundles internal server error response
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError) Code() int {
	return 500
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps][%d] edgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps][%d] edgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout creates a EdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout with default headers values
func NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout() *EdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout {
	return &EdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout{}
}

/*
EdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration query edge application bundles gateway timeout response has a 2xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration query edge application bundles gateway timeout response has a 3xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration query edge application bundles gateway timeout response has a 4xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration query edge application bundles gateway timeout response has a 5xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application configuration query edge application bundles gateway timeout response a status code equal to that given
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge application configuration query edge application bundles gateway timeout response
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps][%d] edgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps][%d] edgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault creates a EdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault with default headers values
func NewEdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault(code int) *EdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault {
	return &EdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge application configuration query edge application bundles default response has a 2xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application configuration query edge application bundles default response has a 3xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application configuration query edge application bundles default response has a 4xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application configuration query edge application bundles default response has a 5xx status code
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application configuration query edge application bundles default response a status code equal to that given
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge application configuration query edge application bundles default response
func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps][%d] EdgeApplicationConfiguration_QueryEdgeApplicationBundles default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps][%d] EdgeApplicationConfiguration_QueryEdgeApplicationBundles default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationConfigurationQueryEdgeApplicationBundlesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
