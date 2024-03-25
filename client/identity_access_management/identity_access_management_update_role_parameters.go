// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"github.com/xyuria-zededa/terraform-provider-zedcloud/models"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewIdentityAccessManagementUpdateRoleParams creates a new IdentityAccessManagementUpdateRoleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementUpdateRoleParams() *IdentityAccessManagementUpdateRoleParams {
	return &IdentityAccessManagementUpdateRoleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementUpdateRoleParamsWithTimeout creates a new IdentityAccessManagementUpdateRoleParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementUpdateRoleParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementUpdateRoleParams {
	return &IdentityAccessManagementUpdateRoleParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementUpdateRoleParamsWithContext creates a new IdentityAccessManagementUpdateRoleParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementUpdateRoleParamsWithContext(ctx context.Context) *IdentityAccessManagementUpdateRoleParams {
	return &IdentityAccessManagementUpdateRoleParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementUpdateRoleParamsWithHTTPClient creates a new IdentityAccessManagementUpdateRoleParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementUpdateRoleParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementUpdateRoleParams {
	return &IdentityAccessManagementUpdateRoleParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementUpdateRoleParams contains all the parameters to send to the API endpoint

	for the identity access management update role operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementUpdateRoleParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.Role

	/* ID.

	   Unique system defined role ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the identity access management update role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementUpdateRoleParams) WithDefaults() *IdentityAccessManagementUpdateRoleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management update role params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementUpdateRoleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management update role params
func (o *IdentityAccessManagementUpdateRoleParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementUpdateRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management update role params
func (o *IdentityAccessManagementUpdateRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management update role params
func (o *IdentityAccessManagementUpdateRoleParams) WithContext(ctx context.Context) *IdentityAccessManagementUpdateRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management update role params
func (o *IdentityAccessManagementUpdateRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management update role params
func (o *IdentityAccessManagementUpdateRoleParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementUpdateRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management update role params
func (o *IdentityAccessManagementUpdateRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management update role params
func (o *IdentityAccessManagementUpdateRoleParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementUpdateRoleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management update role params
func (o *IdentityAccessManagementUpdateRoleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the identity access management update role params
func (o *IdentityAccessManagementUpdateRoleParams) WithBody(body *models.Role) *IdentityAccessManagementUpdateRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identity access management update role params
func (o *IdentityAccessManagementUpdateRoleParams) SetBody(body *models.Role) {
	o.Body = body
}

// WithID adds the id to the identity access management update role params
func (o *IdentityAccessManagementUpdateRoleParams) WithID(id string) *IdentityAccessManagementUpdateRoleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the identity access management update role params
func (o *IdentityAccessManagementUpdateRoleParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementUpdateRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
