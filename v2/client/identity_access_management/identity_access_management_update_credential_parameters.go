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

	"github.com/xyuria-zededa/terraform-provider-zedcloud/models"
)

// NewIdentityAccessManagementUpdateCredentialParams creates a new IdentityAccessManagementUpdateCredentialParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIdentityAccessManagementUpdateCredentialParams() *IdentityAccessManagementUpdateCredentialParams {
	return &IdentityAccessManagementUpdateCredentialParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIdentityAccessManagementUpdateCredentialParamsWithTimeout creates a new IdentityAccessManagementUpdateCredentialParams object
// with the ability to set a timeout on a request.
func NewIdentityAccessManagementUpdateCredentialParamsWithTimeout(timeout time.Duration) *IdentityAccessManagementUpdateCredentialParams {
	return &IdentityAccessManagementUpdateCredentialParams{
		timeout: timeout,
	}
}

// NewIdentityAccessManagementUpdateCredentialParamsWithContext creates a new IdentityAccessManagementUpdateCredentialParams object
// with the ability to set a context for a request.
func NewIdentityAccessManagementUpdateCredentialParamsWithContext(ctx context.Context) *IdentityAccessManagementUpdateCredentialParams {
	return &IdentityAccessManagementUpdateCredentialParams{
		Context: ctx,
	}
}

// NewIdentityAccessManagementUpdateCredentialParamsWithHTTPClient creates a new IdentityAccessManagementUpdateCredentialParams object
// with the ability to set a custom HTTPClient for a request.
func NewIdentityAccessManagementUpdateCredentialParamsWithHTTPClient(client *http.Client) *IdentityAccessManagementUpdateCredentialParams {
	return &IdentityAccessManagementUpdateCredentialParams{
		HTTPClient: client,
	}
}

/*
IdentityAccessManagementUpdateCredentialParams contains all the parameters to send to the API endpoint

	for the identity access management update credential operation.

	Typically these are written to a http.Request.
*/
type IdentityAccessManagementUpdateCredentialParams struct {

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

// WithDefaults hydrates default values in the identity access management update credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementUpdateCredentialParams) WithDefaults() *IdentityAccessManagementUpdateCredentialParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the identity access management update credential params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IdentityAccessManagementUpdateCredentialParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the identity access management update credential params
func (o *IdentityAccessManagementUpdateCredentialParams) WithTimeout(timeout time.Duration) *IdentityAccessManagementUpdateCredentialParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the identity access management update credential params
func (o *IdentityAccessManagementUpdateCredentialParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the identity access management update credential params
func (o *IdentityAccessManagementUpdateCredentialParams) WithContext(ctx context.Context) *IdentityAccessManagementUpdateCredentialParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the identity access management update credential params
func (o *IdentityAccessManagementUpdateCredentialParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the identity access management update credential params
func (o *IdentityAccessManagementUpdateCredentialParams) WithHTTPClient(client *http.Client) *IdentityAccessManagementUpdateCredentialParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the identity access management update credential params
func (o *IdentityAccessManagementUpdateCredentialParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the identity access management update credential params
func (o *IdentityAccessManagementUpdateCredentialParams) WithXRequestID(xRequestID *string) *IdentityAccessManagementUpdateCredentialParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the identity access management update credential params
func (o *IdentityAccessManagementUpdateCredentialParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the identity access management update credential params
func (o *IdentityAccessManagementUpdateCredentialParams) WithBody(body *models.Credential) *IdentityAccessManagementUpdateCredentialParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the identity access management update credential params
func (o *IdentityAccessManagementUpdateCredentialParams) SetBody(body *models.Credential) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IdentityAccessManagementUpdateCredentialParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
