// Code generated by go-swagger; DO NOT EDIT.

package application_instance

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

// NewEdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams creates a new EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams() *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams {
	return &EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParamsWithTimeout creates a new EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams object
// with the ability to set a timeout on a request.
func NewEdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParamsWithTimeout(timeout time.Duration) *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams {
	return &EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams{
		timeout: timeout,
	}
}

// NewEdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParamsWithContext creates a new EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams object
// with the ability to set a context for a request.
func NewEdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParamsWithContext(ctx context.Context) *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams {
	return &EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams{
		Context: ctx,
	}
}

// NewEdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParamsWithHTTPClient creates a new EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParamsWithHTTPClient(client *http.Client) *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams {
	return &EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams{
		HTTPClient: client,
	}
}

/*
EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams contains all the parameters to send to the API endpoint

	for the edge application instance configuration update patch envelope referenceto app instance operation.

	Typically these are written to a http.Request.
*/
type EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.PatchReferenceUpdateConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge application instance configuration update patch envelope referenceto app instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams) WithDefaults() *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge application instance configuration update patch envelope referenceto app instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge application instance configuration update patch envelope referenceto app instance params
func (o *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams) WithTimeout(timeout time.Duration) *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge application instance configuration update patch envelope referenceto app instance params
func (o *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge application instance configuration update patch envelope referenceto app instance params
func (o *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams) WithContext(ctx context.Context) *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge application instance configuration update patch envelope referenceto app instance params
func (o *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge application instance configuration update patch envelope referenceto app instance params
func (o *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams) WithHTTPClient(client *http.Client) *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge application instance configuration update patch envelope referenceto app instance params
func (o *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge application instance configuration update patch envelope referenceto app instance params
func (o *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams) WithXRequestID(xRequestID *string) *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge application instance configuration update patch envelope referenceto app instance params
func (o *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the edge application instance configuration update patch envelope referenceto app instance params
func (o *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams) WithBody(body *models.PatchReferenceUpdateConfig) *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edge application instance configuration update patch envelope referenceto app instance params
func (o *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams) SetBody(body *models.PatchReferenceUpdateConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeApplicationInstanceConfigurationUpdatePatchEnvelopeReferencetoAppInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
