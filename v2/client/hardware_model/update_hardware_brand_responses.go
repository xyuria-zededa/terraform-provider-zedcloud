package hardware_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/xyuria-zededa/terraform-provider-zedcloud/v2/models"
)

// HardwareModelUpdateHardwareBrandReader is a Reader for the HardwareModelUpdateHardwareBrand structure.
type HardwareModelUpdateHardwareBrandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HardwareModelUpdateHardwareBrandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHardwareModelUpdateHardwareBrandOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHardwareModelUpdateHardwareBrandUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHardwareModelUpdateHardwareBrandForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHardwareModelUpdateHardwareBrandNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewHardwareModelUpdateHardwareBrandConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHardwareModelUpdateHardwareBrandInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewHardwareModelUpdateHardwareBrandGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHardwareModelUpdateHardwareBrandDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHardwareModelUpdateHardwareBrandOK creates a HardwareModelUpdateHardwareBrandOK with default headers values
func NewHardwareModelUpdateHardwareBrandOK() *HardwareModelUpdateHardwareBrandOK {
	return &HardwareModelUpdateHardwareBrandOK{}
}

/*
HardwareModelUpdateHardwareBrandOK describes a response with status code 200, with default header values.

A successful response.
*/
type HardwareModelUpdateHardwareBrandOK struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model update hardware brand o k response has a 2xx status code
func (o *HardwareModelUpdateHardwareBrandOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hardware model update hardware brand o k response has a 3xx status code
func (o *HardwareModelUpdateHardwareBrandOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model update hardware brand o k response has a 4xx status code
func (o *HardwareModelUpdateHardwareBrandOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model update hardware brand o k response has a 5xx status code
func (o *HardwareModelUpdateHardwareBrandOK) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model update hardware brand o k response a status code equal to that given
func (o *HardwareModelUpdateHardwareBrandOK) IsCode(code int) bool {
	return code == 200
}

func (o *HardwareModelUpdateHardwareBrandOK) Error() string {
	return fmt.Sprintf("[PUT /v1/brands/id/{id}][%d] hardwareModelUpdateHardwareBrandOK  %+v", 200, o.Payload)
}

func (o *HardwareModelUpdateHardwareBrandOK) String() string {
	return fmt.Sprintf("[PUT /v1/brands/id/{id}][%d] hardwareModelUpdateHardwareBrandOK  %+v", 200, o.Payload)
}

func (o *HardwareModelUpdateHardwareBrandOK) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelUpdateHardwareBrandOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelUpdateHardwareBrandUnauthorized creates a HardwareModelUpdateHardwareBrandUnauthorized with default headers values
func NewHardwareModelUpdateHardwareBrandUnauthorized() *HardwareModelUpdateHardwareBrandUnauthorized {
	return &HardwareModelUpdateHardwareBrandUnauthorized{}
}

/*
HardwareModelUpdateHardwareBrandUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type HardwareModelUpdateHardwareBrandUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model update hardware brand unauthorized response has a 2xx status code
func (o *HardwareModelUpdateHardwareBrandUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model update hardware brand unauthorized response has a 3xx status code
func (o *HardwareModelUpdateHardwareBrandUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model update hardware brand unauthorized response has a 4xx status code
func (o *HardwareModelUpdateHardwareBrandUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model update hardware brand unauthorized response has a 5xx status code
func (o *HardwareModelUpdateHardwareBrandUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model update hardware brand unauthorized response a status code equal to that given
func (o *HardwareModelUpdateHardwareBrandUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *HardwareModelUpdateHardwareBrandUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v1/brands/id/{id}][%d] hardwareModelUpdateHardwareBrandUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelUpdateHardwareBrandUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v1/brands/id/{id}][%d] hardwareModelUpdateHardwareBrandUnauthorized  %+v", 401, o.Payload)
}

func (o *HardwareModelUpdateHardwareBrandUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelUpdateHardwareBrandUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelUpdateHardwareBrandForbidden creates a HardwareModelUpdateHardwareBrandForbidden with default headers values
func NewHardwareModelUpdateHardwareBrandForbidden() *HardwareModelUpdateHardwareBrandForbidden {
	return &HardwareModelUpdateHardwareBrandForbidden{}
}

/*
HardwareModelUpdateHardwareBrandForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have edge-node level access permission for the operation or does not have access scope to the project.
*/
type HardwareModelUpdateHardwareBrandForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model update hardware brand forbidden response has a 2xx status code
func (o *HardwareModelUpdateHardwareBrandForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model update hardware brand forbidden response has a 3xx status code
func (o *HardwareModelUpdateHardwareBrandForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model update hardware brand forbidden response has a 4xx status code
func (o *HardwareModelUpdateHardwareBrandForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model update hardware brand forbidden response has a 5xx status code
func (o *HardwareModelUpdateHardwareBrandForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model update hardware brand forbidden response a status code equal to that given
func (o *HardwareModelUpdateHardwareBrandForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *HardwareModelUpdateHardwareBrandForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/brands/id/{id}][%d] hardwareModelUpdateHardwareBrandForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelUpdateHardwareBrandForbidden) String() string {
	return fmt.Sprintf("[PUT /v1/brands/id/{id}][%d] hardwareModelUpdateHardwareBrandForbidden  %+v", 403, o.Payload)
}

func (o *HardwareModelUpdateHardwareBrandForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelUpdateHardwareBrandForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelUpdateHardwareBrandNotFound creates a HardwareModelUpdateHardwareBrandNotFound with default headers values
func NewHardwareModelUpdateHardwareBrandNotFound() *HardwareModelUpdateHardwareBrandNotFound {
	return &HardwareModelUpdateHardwareBrandNotFound{}
}

/*
HardwareModelUpdateHardwareBrandNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type HardwareModelUpdateHardwareBrandNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model update hardware brand not found response has a 2xx status code
func (o *HardwareModelUpdateHardwareBrandNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model update hardware brand not found response has a 3xx status code
func (o *HardwareModelUpdateHardwareBrandNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model update hardware brand not found response has a 4xx status code
func (o *HardwareModelUpdateHardwareBrandNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model update hardware brand not found response has a 5xx status code
func (o *HardwareModelUpdateHardwareBrandNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model update hardware brand not found response a status code equal to that given
func (o *HardwareModelUpdateHardwareBrandNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *HardwareModelUpdateHardwareBrandNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/brands/id/{id}][%d] hardwareModelUpdateHardwareBrandNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelUpdateHardwareBrandNotFound) String() string {
	return fmt.Sprintf("[PUT /v1/brands/id/{id}][%d] hardwareModelUpdateHardwareBrandNotFound  %+v", 404, o.Payload)
}

func (o *HardwareModelUpdateHardwareBrandNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelUpdateHardwareBrandNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelUpdateHardwareBrandConflict creates a HardwareModelUpdateHardwareBrandConflict with default headers values
func NewHardwareModelUpdateHardwareBrandConflict() *HardwareModelUpdateHardwareBrandConflict {
	return &HardwareModelUpdateHardwareBrandConflict{}
}

/*
HardwareModelUpdateHardwareBrandConflict describes a response with status code 409, with default header values.

Conflict. The API gateway did not process the request because this operation will conflict with an already existing hardware brand record.
*/
type HardwareModelUpdateHardwareBrandConflict struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model update hardware brand conflict response has a 2xx status code
func (o *HardwareModelUpdateHardwareBrandConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model update hardware brand conflict response has a 3xx status code
func (o *HardwareModelUpdateHardwareBrandConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model update hardware brand conflict response has a 4xx status code
func (o *HardwareModelUpdateHardwareBrandConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this hardware model update hardware brand conflict response has a 5xx status code
func (o *HardwareModelUpdateHardwareBrandConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this hardware model update hardware brand conflict response a status code equal to that given
func (o *HardwareModelUpdateHardwareBrandConflict) IsCode(code int) bool {
	return code == 409
}

func (o *HardwareModelUpdateHardwareBrandConflict) Error() string {
	return fmt.Sprintf("[PUT /v1/brands/id/{id}][%d] hardwareModelUpdateHardwareBrandConflict  %+v", 409, o.Payload)
}

func (o *HardwareModelUpdateHardwareBrandConflict) String() string {
	return fmt.Sprintf("[PUT /v1/brands/id/{id}][%d] hardwareModelUpdateHardwareBrandConflict  %+v", 409, o.Payload)
}

func (o *HardwareModelUpdateHardwareBrandConflict) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelUpdateHardwareBrandConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelUpdateHardwareBrandInternalServerError creates a HardwareModelUpdateHardwareBrandInternalServerError with default headers values
func NewHardwareModelUpdateHardwareBrandInternalServerError() *HardwareModelUpdateHardwareBrandInternalServerError {
	return &HardwareModelUpdateHardwareBrandInternalServerError{}
}

/*
HardwareModelUpdateHardwareBrandInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type HardwareModelUpdateHardwareBrandInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model update hardware brand internal server error response has a 2xx status code
func (o *HardwareModelUpdateHardwareBrandInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model update hardware brand internal server error response has a 3xx status code
func (o *HardwareModelUpdateHardwareBrandInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model update hardware brand internal server error response has a 4xx status code
func (o *HardwareModelUpdateHardwareBrandInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model update hardware brand internal server error response has a 5xx status code
func (o *HardwareModelUpdateHardwareBrandInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model update hardware brand internal server error response a status code equal to that given
func (o *HardwareModelUpdateHardwareBrandInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *HardwareModelUpdateHardwareBrandInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/brands/id/{id}][%d] hardwareModelUpdateHardwareBrandInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelUpdateHardwareBrandInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v1/brands/id/{id}][%d] hardwareModelUpdateHardwareBrandInternalServerError  %+v", 500, o.Payload)
}

func (o *HardwareModelUpdateHardwareBrandInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelUpdateHardwareBrandInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelUpdateHardwareBrandGatewayTimeout creates a HardwareModelUpdateHardwareBrandGatewayTimeout with default headers values
func NewHardwareModelUpdateHardwareBrandGatewayTimeout() *HardwareModelUpdateHardwareBrandGatewayTimeout {
	return &HardwareModelUpdateHardwareBrandGatewayTimeout{}
}

/*
HardwareModelUpdateHardwareBrandGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type HardwareModelUpdateHardwareBrandGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this hardware model update hardware brand gateway timeout response has a 2xx status code
func (o *HardwareModelUpdateHardwareBrandGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this hardware model update hardware brand gateway timeout response has a 3xx status code
func (o *HardwareModelUpdateHardwareBrandGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hardware model update hardware brand gateway timeout response has a 4xx status code
func (o *HardwareModelUpdateHardwareBrandGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this hardware model update hardware brand gateway timeout response has a 5xx status code
func (o *HardwareModelUpdateHardwareBrandGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this hardware model update hardware brand gateway timeout response a status code equal to that given
func (o *HardwareModelUpdateHardwareBrandGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *HardwareModelUpdateHardwareBrandGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /v1/brands/id/{id}][%d] hardwareModelUpdateHardwareBrandGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelUpdateHardwareBrandGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /v1/brands/id/{id}][%d] hardwareModelUpdateHardwareBrandGatewayTimeout  %+v", 504, o.Payload)
}

func (o *HardwareModelUpdateHardwareBrandGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *HardwareModelUpdateHardwareBrandGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHardwareModelUpdateHardwareBrandDefault creates a HardwareModelUpdateHardwareBrandDefault with default headers values
func NewHardwareModelUpdateHardwareBrandDefault(code int) *HardwareModelUpdateHardwareBrandDefault {
	return &HardwareModelUpdateHardwareBrandDefault{
		_statusCode: code,
	}
}

/*
HardwareModelUpdateHardwareBrandDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type HardwareModelUpdateHardwareBrandDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// Code gets the status code for the hardware model update hardware brand default response
func (o *HardwareModelUpdateHardwareBrandDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this hardware model update hardware brand default response has a 2xx status code
func (o *HardwareModelUpdateHardwareBrandDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this hardware model update hardware brand default response has a 3xx status code
func (o *HardwareModelUpdateHardwareBrandDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this hardware model update hardware brand default response has a 4xx status code
func (o *HardwareModelUpdateHardwareBrandDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this hardware model update hardware brand default response has a 5xx status code
func (o *HardwareModelUpdateHardwareBrandDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this hardware model update hardware brand default response a status code equal to that given
func (o *HardwareModelUpdateHardwareBrandDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *HardwareModelUpdateHardwareBrandDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/brands/id/{id}][%d] HardwareModel_UpdateHardwareBrand default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelUpdateHardwareBrandDefault) String() string {
	return fmt.Sprintf("[PUT /v1/brands/id/{id}][%d] HardwareModel_UpdateHardwareBrand default  %+v", o._statusCode, o.Payload)
}

func (o *HardwareModelUpdateHardwareBrandDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *HardwareModelUpdateHardwareBrandDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
HardwareModelUpdateHardwareBrandBody SysBrand  payload details
//
// SysModel consists of various brand attributes like id, name, title, svg, state, systemMfgName
// Example: {"id":"d85a545f-6510-4327-b03d-c02eef119e99","name":"zed-brand","title":"sample brand"}
swagger:model HardwareModelUpdateHardwareBrandBody
*/
type HardwareModelUpdateHardwareBrandBody struct {

	// Map of <string, string>
	Attr map[string]string `json:"attr,omitempty"`

	// Detailed description of the image.
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// Map of <string, string> which holds the key:url for the logo artifact of the the brand
	Logo map[string]string `json:"logo,omitempty"`

	// user defined sys brand name
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name *string `json:"name"`

	// origin of object
	// Required: true
	OriginType *models.Origin `json:"originType"`

	// Object Revision  of the sys brand
	Revision *models.ObjectRevision `json:"revision,omitempty"`

	// Sys Model Status
	State *models.SysModelState `json:"state,omitempty"`

	// Deprecated: base64 encoded string of svg file
	Svg string `json:"svg,omitempty"`

	// System Manufacturer name
	SystemMfgName string `json:"systemMfgName,omitempty"`

	// user defined title for sys brand
	// Required: true
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Title *string `json:"title"`
}

// Validate validates this hardware model update hardware brand body
func (o *HardwareModelUpdateHardwareBrandBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOriginType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *HardwareModelUpdateHardwareBrandBody) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("body"+"."+"description", "body", o.Description, 256); err != nil {
		return err
	}

	return nil
}

func (o *HardwareModelUpdateHardwareBrandBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	if err := validate.MinLength("body"+"."+"name", "body", *o.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("body"+"."+"name", "body", *o.Name, 256); err != nil {
		return err
	}

	if err := validate.Pattern("body"+"."+"name", "body", *o.Name, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (o *HardwareModelUpdateHardwareBrandBody) validateOriginType(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"originType", "body", o.OriginType); err != nil {
		return err
	}

	if err := validate.Required("body"+"."+"originType", "body", o.OriginType); err != nil {
		return err
	}

	if o.OriginType != nil {
		if err := o.OriginType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "originType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "originType")
			}
			return err
		}
	}

	return nil
}

func (o *HardwareModelUpdateHardwareBrandBody) validateRevision(formats strfmt.Registry) error {
	if swag.IsZero(o.Revision) { // not required
		return nil
	}

	if o.Revision != nil {
		if err := o.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "revision")
			}
			return err
		}
	}

	return nil
}

func (o *HardwareModelUpdateHardwareBrandBody) validateState(formats strfmt.Registry) error {
	if swag.IsZero(o.State) { // not required
		return nil
	}

	if o.State != nil {
		if err := o.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "state")
			}
			return err
		}
	}

	return nil
}

func (o *HardwareModelUpdateHardwareBrandBody) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"title", "body", o.Title); err != nil {
		return err
	}

	if err := validate.Pattern("body"+"."+"title", "body", *o.Title, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this hardware model update hardware brand body based on the context it is used
func (o *HardwareModelUpdateHardwareBrandBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateOriginType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *HardwareModelUpdateHardwareBrandBody) contextValidateOriginType(ctx context.Context, formats strfmt.Registry) error {

	if o.OriginType != nil {
		if err := o.OriginType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "originType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "originType")
			}
			return err
		}
	}

	return nil
}

func (o *HardwareModelUpdateHardwareBrandBody) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if o.Revision != nil {
		if err := o.Revision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "revision")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "revision")
			}
			return err
		}
	}

	return nil
}

func (o *HardwareModelUpdateHardwareBrandBody) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if o.State != nil {
		if err := o.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *HardwareModelUpdateHardwareBrandBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *HardwareModelUpdateHardwareBrandBody) UnmarshalBinary(b []byte) error {
	var res HardwareModelUpdateHardwareBrandBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
