package application

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

// UpdateParams creates a new EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func UpdateParams() *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams {
	return &EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleParamsWithTimeout creates a new EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams object
// with the ability to set a timeout on a request.
func NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleParamsWithTimeout(timeout time.Duration) *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams {
	return &EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams{
		timeout: timeout,
	}
}

// NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleParamsWithContext creates a new EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams object
// with the ability to set a context for a request.
func NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleParamsWithContext(ctx context.Context) *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams {
	return &EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams{
		Context: ctx,
	}
}

// NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleParamsWithHTTPClient creates a new EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeApplicationConfigurationUpdateEdgeApplicationBundleParamsWithHTTPClient(client *http.Client) *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams {
	return &EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams{
		HTTPClient: client,
	}
}

/*
EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams contains all the parameters to send to the API endpoint

	for the edge application configuration update edge application bundle operation.

	Typically these are written to a http.Request.
*/
type EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.Application

	/* ID.

	   System defined universally unique Id of the edge application
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge application configuration update edge application bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams) WithDefaults() *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge application configuration update edge application bundle params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge application configuration update edge application bundle params
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams) WithTimeout(timeout time.Duration) *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge application configuration update edge application bundle params
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge application configuration update edge application bundle params
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams) WithContext(ctx context.Context) *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge application configuration update edge application bundle params
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge application configuration update edge application bundle params
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams) WithHTTPClient(client *http.Client) *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge application configuration update edge application bundle params
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge application configuration update edge application bundle params
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams) WithXRequestID(xRequestID *string) *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge application configuration update edge application bundle params
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the edge application configuration update edge application bundle params
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams) WithBody(body *models.Application) *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edge application configuration update edge application bundle params
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams) SetBody(body *models.Application) {
	o.Body = body
}

// WithID adds the id to the edge application configuration update edge application bundle params
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams) WithID(id string) *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge application configuration update edge application bundle params
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeApplicationConfigurationUpdateEdgeApplicationBundleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
