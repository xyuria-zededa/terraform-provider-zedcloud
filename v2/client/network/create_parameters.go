package network

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

// CreateParams creates a new EdgeNetworkConfigurationCreateEdgeNetworkParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func CreateParams() *EdgeNetworkConfigurationCreateEdgeNetworkParams {
	return &EdgeNetworkConfigurationCreateEdgeNetworkParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkParamsWithTimeout creates a new EdgeNetworkConfigurationCreateEdgeNetworkParams object
// with the ability to set a timeout on a request.
func NewEdgeNetworkConfigurationCreateEdgeNetworkParamsWithTimeout(timeout time.Duration) *EdgeNetworkConfigurationCreateEdgeNetworkParams {
	return &EdgeNetworkConfigurationCreateEdgeNetworkParams{
		timeout: timeout,
	}
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkParamsWithContext creates a new EdgeNetworkConfigurationCreateEdgeNetworkParams object
// with the ability to set a context for a request.
func NewEdgeNetworkConfigurationCreateEdgeNetworkParamsWithContext(ctx context.Context) *EdgeNetworkConfigurationCreateEdgeNetworkParams {
	return &EdgeNetworkConfigurationCreateEdgeNetworkParams{
		Context: ctx,
	}
}

// NewEdgeNetworkConfigurationCreateEdgeNetworkParamsWithHTTPClient creates a new EdgeNetworkConfigurationCreateEdgeNetworkParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNetworkConfigurationCreateEdgeNetworkParamsWithHTTPClient(client *http.Client) *EdgeNetworkConfigurationCreateEdgeNetworkParams {
	return &EdgeNetworkConfigurationCreateEdgeNetworkParams{
		HTTPClient: client,
	}
}

/*
EdgeNetworkConfigurationCreateEdgeNetworkParams contains all the parameters to send to the API endpoint

	for the edge network configuration create edge network operation.

	Typically these are written to a http.Request.
*/
type EdgeNetworkConfigurationCreateEdgeNetworkParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.Network

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge network configuration create edge network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkConfigurationCreateEdgeNetworkParams) WithDefaults() *EdgeNetworkConfigurationCreateEdgeNetworkParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge network configuration create edge network params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNetworkConfigurationCreateEdgeNetworkParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge network configuration create edge network params
func (o *EdgeNetworkConfigurationCreateEdgeNetworkParams) WithTimeout(timeout time.Duration) *EdgeNetworkConfigurationCreateEdgeNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge network configuration create edge network params
func (o *EdgeNetworkConfigurationCreateEdgeNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge network configuration create edge network params
func (o *EdgeNetworkConfigurationCreateEdgeNetworkParams) WithContext(ctx context.Context) *EdgeNetworkConfigurationCreateEdgeNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge network configuration create edge network params
func (o *EdgeNetworkConfigurationCreateEdgeNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge network configuration create edge network params
func (o *EdgeNetworkConfigurationCreateEdgeNetworkParams) WithHTTPClient(client *http.Client) *EdgeNetworkConfigurationCreateEdgeNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge network configuration create edge network params
func (o *EdgeNetworkConfigurationCreateEdgeNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge network configuration create edge network params
func (o *EdgeNetworkConfigurationCreateEdgeNetworkParams) WithXRequestID(xRequestID *string) *EdgeNetworkConfigurationCreateEdgeNetworkParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge network configuration create edge network params
func (o *EdgeNetworkConfigurationCreateEdgeNetworkParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the edge network configuration create edge network params
func (o *EdgeNetworkConfigurationCreateEdgeNetworkParams) WithBody(body *models.Network) *EdgeNetworkConfigurationCreateEdgeNetworkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edge network configuration create edge network params
func (o *EdgeNetworkConfigurationCreateEdgeNetworkParams) SetBody(body *models.Network) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNetworkConfigurationCreateEdgeNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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