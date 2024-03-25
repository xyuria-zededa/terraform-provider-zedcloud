package network_instance

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

// CreateNetworkInstanceParams creates a new EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func CreateNetworkInstanceParams() *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams {
	return &EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParamsWithTimeout creates a new EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams object
// with the ability to set a timeout on a request.
func NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParamsWithTimeout(timeout time.Duration) *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams {
	return &EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams{
		timeout: timeout,
	}
}

// NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParamsWithContext creates a new EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams object
// with the ability to set a context for a request.
func NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParamsWithContext(ctx context.Context) *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams {
	return &EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams{
		Context: ctx,
	}
}

// NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParamsWithHTTPClient creates a new EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParamsWithHTTPClient(client *http.Client) *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams {
	return &EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams{
		HTTPClient: client,
	}
}

/*
EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams contains all the parameters to send to the API endpoint

	for the edge network instance configuration create edge network instance operation.

	Typically these are written to a http.Request.
*/
type EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.NetworkInstance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge network instance configuration create edge network instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams) WithDefaults() *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge network instance configuration create edge network instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge network instance configuration create edge network instance params
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams) WithTimeout(timeout time.Duration) *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge network instance configuration create edge network instance params
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge network instance configuration create edge network instance params
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams) WithContext(ctx context.Context) *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge network instance configuration create edge network instance params
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge network instance configuration create edge network instance params
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams) WithHTTPClient(client *http.Client) *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge network instance configuration create edge network instance params
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge network instance configuration create edge network instance params
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams) WithXRequestID(xRequestID *string) *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge network instance configuration create edge network instance params
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the edge network instance configuration create edge network instance params
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams) WithBody(body *models.NetworkInstance) *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edge network instance configuration create edge network instance params
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams) SetBody(body *models.NetworkInstance) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNetworkInstanceConfigurationCreateEdgeNetworkInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
