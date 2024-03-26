package volume_instance

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

// CreateParams creates a new VolumeInstanceConfigurationCreateVolumeInstanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func CreateParams() *VolumeInstanceConfigurationCreateVolumeInstanceParams {
	return &VolumeInstanceConfigurationCreateVolumeInstanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceParamsWithTimeout creates a new VolumeInstanceConfigurationCreateVolumeInstanceParams object
// with the ability to set a timeout on a request.
func NewVolumeInstanceConfigurationCreateVolumeInstanceParamsWithTimeout(timeout time.Duration) *VolumeInstanceConfigurationCreateVolumeInstanceParams {
	return &VolumeInstanceConfigurationCreateVolumeInstanceParams{
		timeout: timeout,
	}
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceParamsWithContext creates a new VolumeInstanceConfigurationCreateVolumeInstanceParams object
// with the ability to set a context for a request.
func NewVolumeInstanceConfigurationCreateVolumeInstanceParamsWithContext(ctx context.Context) *VolumeInstanceConfigurationCreateVolumeInstanceParams {
	return &VolumeInstanceConfigurationCreateVolumeInstanceParams{
		Context: ctx,
	}
}

// NewVolumeInstanceConfigurationCreateVolumeInstanceParamsWithHTTPClient creates a new VolumeInstanceConfigurationCreateVolumeInstanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeInstanceConfigurationCreateVolumeInstanceParamsWithHTTPClient(client *http.Client) *VolumeInstanceConfigurationCreateVolumeInstanceParams {
	return &VolumeInstanceConfigurationCreateVolumeInstanceParams{
		HTTPClient: client,
	}
}

/*
VolumeInstanceConfigurationCreateVolumeInstanceParams contains all the parameters to send to the API endpoint

	for the volume instance configuration create volume instance operation.

	Typically these are written to a http.Request.
*/
type VolumeInstanceConfigurationCreateVolumeInstanceParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.VolumeInstance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volume instance configuration create volume instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeInstanceConfigurationCreateVolumeInstanceParams) WithDefaults() *VolumeInstanceConfigurationCreateVolumeInstanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume instance configuration create volume instance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeInstanceConfigurationCreateVolumeInstanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the volume instance configuration create volume instance params
func (o *VolumeInstanceConfigurationCreateVolumeInstanceParams) WithTimeout(timeout time.Duration) *VolumeInstanceConfigurationCreateVolumeInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume instance configuration create volume instance params
func (o *VolumeInstanceConfigurationCreateVolumeInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume instance configuration create volume instance params
func (o *VolumeInstanceConfigurationCreateVolumeInstanceParams) WithContext(ctx context.Context) *VolumeInstanceConfigurationCreateVolumeInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume instance configuration create volume instance params
func (o *VolumeInstanceConfigurationCreateVolumeInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume instance configuration create volume instance params
func (o *VolumeInstanceConfigurationCreateVolumeInstanceParams) WithHTTPClient(client *http.Client) *VolumeInstanceConfigurationCreateVolumeInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume instance configuration create volume instance params
func (o *VolumeInstanceConfigurationCreateVolumeInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the volume instance configuration create volume instance params
func (o *VolumeInstanceConfigurationCreateVolumeInstanceParams) WithXRequestID(xRequestID *string) *VolumeInstanceConfigurationCreateVolumeInstanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the volume instance configuration create volume instance params
func (o *VolumeInstanceConfigurationCreateVolumeInstanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the volume instance configuration create volume instance params
func (o *VolumeInstanceConfigurationCreateVolumeInstanceParams) WithBody(body *models.VolumeInstance) *VolumeInstanceConfigurationCreateVolumeInstanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the volume instance configuration create volume instance params
func (o *VolumeInstanceConfigurationCreateVolumeInstanceParams) SetBody(body *models.VolumeInstance) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeInstanceConfigurationCreateVolumeInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
