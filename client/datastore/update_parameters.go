package datastore

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
	"github.com/zededa/terraform-provider-zedcloud/models"
)

// UpdateParams creates a new DatastoreConfigurationUpdateDatastoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func UpdateParams() *DatastoreConfigurationUpdateDatastoreParams {
	return &DatastoreConfigurationUpdateDatastoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDatastoreConfigurationUpdateDatastoreParamsWithTimeout creates a new DatastoreConfigurationUpdateDatastoreParams object
// with the ability to set a timeout on a request.
func NewDatastoreConfigurationUpdateDatastoreParamsWithTimeout(timeout time.Duration) *DatastoreConfigurationUpdateDatastoreParams {
	return &DatastoreConfigurationUpdateDatastoreParams{
		timeout: timeout,
	}
}

// NewDatastoreConfigurationUpdateDatastoreParamsWithContext creates a new DatastoreConfigurationUpdateDatastoreParams object
// with the ability to set a context for a request.
func NewDatastoreConfigurationUpdateDatastoreParamsWithContext(ctx context.Context) *DatastoreConfigurationUpdateDatastoreParams {
	return &DatastoreConfigurationUpdateDatastoreParams{
		Context: ctx,
	}
}

// NewDatastoreConfigurationUpdateDatastoreParamsWithHTTPClient creates a new DatastoreConfigurationUpdateDatastoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewDatastoreConfigurationUpdateDatastoreParamsWithHTTPClient(client *http.Client) *DatastoreConfigurationUpdateDatastoreParams {
	return &DatastoreConfigurationUpdateDatastoreParams{
		HTTPClient: client,
	}
}

/*
DatastoreConfigurationUpdateDatastoreParams contains all the parameters to send to the API endpoint

	for the datastore configuration update datastore operation.

	Typically these are written to a http.Request.
*/
type DatastoreConfigurationUpdateDatastoreParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	// Body.
	Body *models.Datastore

	/* ID.

	   System defined universally unique Id of the datastore.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the datastore configuration update datastore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DatastoreConfigurationUpdateDatastoreParams) WithDefaults() *DatastoreConfigurationUpdateDatastoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the datastore configuration update datastore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DatastoreConfigurationUpdateDatastoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the datastore configuration update datastore params
func (o *DatastoreConfigurationUpdateDatastoreParams) WithTimeout(timeout time.Duration) *DatastoreConfigurationUpdateDatastoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the datastore configuration update datastore params
func (o *DatastoreConfigurationUpdateDatastoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the datastore configuration update datastore params
func (o *DatastoreConfigurationUpdateDatastoreParams) WithContext(ctx context.Context) *DatastoreConfigurationUpdateDatastoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the datastore configuration update datastore params
func (o *DatastoreConfigurationUpdateDatastoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the datastore configuration update datastore params
func (o *DatastoreConfigurationUpdateDatastoreParams) WithHTTPClient(client *http.Client) *DatastoreConfigurationUpdateDatastoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the datastore configuration update datastore params
func (o *DatastoreConfigurationUpdateDatastoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the datastore configuration update datastore params
func (o *DatastoreConfigurationUpdateDatastoreParams) WithXRequestID(xRequestID *string) *DatastoreConfigurationUpdateDatastoreParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the datastore configuration update datastore params
func (o *DatastoreConfigurationUpdateDatastoreParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the datastore configuration update datastore params
func (o *DatastoreConfigurationUpdateDatastoreParams) WithBody(body *models.Datastore) *DatastoreConfigurationUpdateDatastoreParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the datastore configuration update datastore params
func (o *DatastoreConfigurationUpdateDatastoreParams) SetBody(body *models.Datastore) {
	o.Body = body
}

// WithID adds the id to the datastore configuration update datastore params
func (o *DatastoreConfigurationUpdateDatastoreParams) WithID(id string) *DatastoreConfigurationUpdateDatastoreParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the datastore configuration update datastore params
func (o *DatastoreConfigurationUpdateDatastoreParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DatastoreConfigurationUpdateDatastoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
