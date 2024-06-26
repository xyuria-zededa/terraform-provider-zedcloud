// Code generated by go-swagger; DO NOT EDIT.

package patch_envelope

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

// NewPatchEnvelopeConfigurationUpdatePatchEnvelopeParams creates a new PatchEnvelopeConfigurationUpdatePatchEnvelopeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchEnvelopeConfigurationUpdatePatchEnvelopeParams() *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams {
	return &PatchEnvelopeConfigurationUpdatePatchEnvelopeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchEnvelopeConfigurationUpdatePatchEnvelopeParamsWithTimeout creates a new PatchEnvelopeConfigurationUpdatePatchEnvelopeParams object
// with the ability to set a timeout on a request.
func NewPatchEnvelopeConfigurationUpdatePatchEnvelopeParamsWithTimeout(timeout time.Duration) *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams {
	return &PatchEnvelopeConfigurationUpdatePatchEnvelopeParams{
		timeout: timeout,
	}
}

// NewPatchEnvelopeConfigurationUpdatePatchEnvelopeParamsWithContext creates a new PatchEnvelopeConfigurationUpdatePatchEnvelopeParams object
// with the ability to set a context for a request.
func NewPatchEnvelopeConfigurationUpdatePatchEnvelopeParamsWithContext(ctx context.Context) *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams {
	return &PatchEnvelopeConfigurationUpdatePatchEnvelopeParams{
		Context: ctx,
	}
}

// NewPatchEnvelopeConfigurationUpdatePatchEnvelopeParamsWithHTTPClient creates a new PatchEnvelopeConfigurationUpdatePatchEnvelopeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchEnvelopeConfigurationUpdatePatchEnvelopeParamsWithHTTPClient(client *http.Client) *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams {
	return &PatchEnvelopeConfigurationUpdatePatchEnvelopeParams{
		HTTPClient: client,
	}
}

/*
PatchEnvelopeConfigurationUpdatePatchEnvelopeParams contains all the parameters to send to the API endpoint

	for the patch envelope configuration update patch envelope operation.

	Typically these are written to a http.Request.
*/
type PatchEnvelopeConfigurationUpdatePatchEnvelopeParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.PatchEnvelope

	/* ID.

	   System defined universally unique Id of the patch envelope.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch envelope configuration update patch envelope params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams) WithDefaults() *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch envelope configuration update patch envelope params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch envelope configuration update patch envelope params
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams) WithTimeout(timeout time.Duration) *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch envelope configuration update patch envelope params
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch envelope configuration update patch envelope params
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams) WithContext(ctx context.Context) *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch envelope configuration update patch envelope params
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch envelope configuration update patch envelope params
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams) WithHTTPClient(client *http.Client) *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch envelope configuration update patch envelope params
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the patch envelope configuration update patch envelope params
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams) WithXRequestID(xRequestID *string) *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the patch envelope configuration update patch envelope params
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the patch envelope configuration update patch envelope params
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams) WithBody(body *models.PatchEnvelope) *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch envelope configuration update patch envelope params
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams) SetBody(body *models.PatchEnvelope) {
	o.Body = body
}

// WithID adds the id to the patch envelope configuration update patch envelope params
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams) WithID(id string) *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch envelope configuration update patch envelope params
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchEnvelopeConfigurationUpdatePatchEnvelopeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
