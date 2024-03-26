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

	"github.com/xyuria-zededa/terraform-provider-zedcloud/v2/models"
)

// NewIdentityAccessManagementCreateCredentialParams creates a new IdentityAccessManagementCreateCredentialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementCreateCredentialParams() *IdentityAccessManagementCreateCredentialParams {
	return &IdentityAccessManagementCreateCredentialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementCreateCredentialParamsWithTimeout creates a new IdentityAccessManagementCreateCredentialParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementCreateCredentialParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementCreateCredentialParams {
	return &IdentityAccessManagementCreateCredentialParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementCreateCredentialParamsWithContext creates a new IdentityAccessManagementCreateCredentialParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementCreateCredentialParamsWithContext(ctx context.Context) *IdentityAccessManagementCreateCredentialParams {
	return &IdentityAccessManagementCreateCredentialParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementCreateCredentialParamsWithHTTPClient creates a new IdentityAccessManagementCreateCredentialParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementCreateCredentialParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementCreateCredentialParams {
	return &IdentityAccessManagementCreateCredentialParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementCreateCredentialParams contains all the parameters to send to the API endpoint

	for the identity access management create credential operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementCreateCredentialParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.Credential

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the identity access management create credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementCreateCredentialParams) WithDefaults() *IdentityAccessManagementCreateCredentialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management create credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementCreateCredentialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management create credential params
func (o *IdentityAccessManagementCreateCredentialParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementCreateCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management create credential params
func (o *IdentityAccessManagementCreateCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management create credential params
func (o *IdentityAccessManagementCreateCredentialParams) WithContext(ctx context.Context) *IdentityAccessManagementCreateCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management create credential params
func (o *IdentityAccessManagementCreateCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management create credential params
func (o *IdentityAccessManagementCreateCredentialParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementCreateCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management create credential params
func (o *IdentityAccessManagementCreateCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management create credential params
func (o *IdentityAccessManagementCreateCredentialParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementCreateCredentialParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management create credential params
func (o *IdentityAccessManagementCreateCredentialParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the identity access management create credential params
func (o *IdentityAccessManagementCreateCredentialParams) WithBody(body *models.Credential) *IdentityAccessManagementCreateCredentialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identity access management create credential params
func (o *IdentityAccessManagementCreateCredentialParams) SetBody(body *models.Credential) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementCreateCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
