// Code generated by go-swagger; DO NOT EDIT.

package identity_access_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewIdentityAccessManagementGetRoleByNameParams creates a new IdentityAccessManagementGetRoleByNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementGetRoleByNameParams() *IdentityAccessManagementGetRoleByNameParams {
	return &IdentityAccessManagementGetRoleByNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementGetRoleByNameParamsWithTimeout creates a new IdentityAccessManagementGetRoleByNameParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementGetRoleByNameParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementGetRoleByNameParams {
	return &IdentityAccessManagementGetRoleByNameParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementGetRoleByNameParamsWithContext creates a new IdentityAccessManagementGetRoleByNameParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementGetRoleByNameParamsWithContext(ctx context.Context) *IdentityAccessManagementGetRoleByNameParams {
	return &IdentityAccessManagementGetRoleByNameParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementGetRoleByNameParamsWithHTTPClient creates a new IdentityAccessManagementGetRoleByNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementGetRoleByNameParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementGetRoleByNameParams {
	return &IdentityAccessManagementGetRoleByNameParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementGetRoleByNameParams contains all the parameters to send to the API endpoint

	for the identity access management get role by name operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementGetRoleByNameParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the identity access management get role by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementGetRoleByNameParams) WithDefaults() *IdentityAccessManagementGetRoleByNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management get role by name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementGetRoleByNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management get role by name params
func (o *IdentityAccessManagementGetRoleByNameParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementGetRoleByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management get role by name params
func (o *IdentityAccessManagementGetRoleByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management get role by name params
func (o *IdentityAccessManagementGetRoleByNameParams) WithContext(ctx context.Context) *IdentityAccessManagementGetRoleByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management get role by name params
func (o *IdentityAccessManagementGetRoleByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management get role by name params
func (o *IdentityAccessManagementGetRoleByNameParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementGetRoleByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management get role by name params
func (o *IdentityAccessManagementGetRoleByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management get role by name params
func (o *IdentityAccessManagementGetRoleByNameParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementGetRoleByNameParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management get role by name params
func (o *IdentityAccessManagementGetRoleByNameParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithName adds the name to the identity access management get role by name params
func (o *IdentityAccessManagementGetRoleByNameParams) WithName(name string) *IdentityAccessManagementGetRoleByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the identity access management get role by name params
func (o *IdentityAccessManagementGetRoleByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementGetRoleByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
