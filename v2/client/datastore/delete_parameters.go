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
)

// DeleteParams creates a new DatastoreConfigurationDeleteDatastoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func DeleteParams() *DatastoreConfigurationDeleteDatastoreParams {
	return &DatastoreConfigurationDeleteDatastoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDatastoreConfigurationDeleteDatastoreParamsWithTimeout creates a new DatastoreConfigurationDeleteDatastoreParams object
// with the ability to set a timeout on a request.
func NewDatastoreConfigurationDeleteDatastoreParamsWithTimeout(timeout time.Duration) *DatastoreConfigurationDeleteDatastoreParams {
	return &DatastoreConfigurationDeleteDatastoreParams{
		timeout: timeout,
	}
}

// NewDatastoreConfigurationDeleteDatastoreParamsWithContext creates a new DatastoreConfigurationDeleteDatastoreParams object
// with the ability to set a context for a request.
func NewDatastoreConfigurationDeleteDatastoreParamsWithContext(ctx context.Context) *DatastoreConfigurationDeleteDatastoreParams {
	return &DatastoreConfigurationDeleteDatastoreParams{
		Context: ctx,
	}
}

// NewDatastoreConfigurationDeleteDatastoreParamsWithHTTPClient creates a new DatastoreConfigurationDeleteDatastoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewDatastoreConfigurationDeleteDatastoreParamsWithHTTPClient(client *http.Client) *DatastoreConfigurationDeleteDatastoreParams {
	return &DatastoreConfigurationDeleteDatastoreParams{
		HTTPClient: client,
	}
}

/*
DatastoreConfigurationDeleteDatastoreParams contains all the parameters to send to the API endpoint

	for the datastore configuration delete datastore operation.

	Typically these are written to a http.Request.
*/
type DatastoreConfigurationDeleteDatastoreParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   System defined universally unique Id of the datastore
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the datastore configuration delete datastore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DatastoreConfigurationDeleteDatastoreParams) WithDefaults() *DatastoreConfigurationDeleteDatastoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the datastore configuration delete datastore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DatastoreConfigurationDeleteDatastoreParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the datastore configuration delete datastore params
func (o *DatastoreConfigurationDeleteDatastoreParams) WithTimeout(timeout time.Duration) *DatastoreConfigurationDeleteDatastoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the datastore configuration delete datastore params
func (o *DatastoreConfigurationDeleteDatastoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the datastore configuration delete datastore params
func (o *DatastoreConfigurationDeleteDatastoreParams) WithContext(ctx context.Context) *DatastoreConfigurationDeleteDatastoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the datastore configuration delete datastore params
func (o *DatastoreConfigurationDeleteDatastoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the datastore configuration delete datastore params
func (o *DatastoreConfigurationDeleteDatastoreParams) WithHTTPClient(client *http.Client) *DatastoreConfigurationDeleteDatastoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the datastore configuration delete datastore params
func (o *DatastoreConfigurationDeleteDatastoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the datastore configuration delete datastore params
func (o *DatastoreConfigurationDeleteDatastoreParams) WithXRequestID(xRequestID *string) *DatastoreConfigurationDeleteDatastoreParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the datastore configuration delete datastore params
func (o *DatastoreConfigurationDeleteDatastoreParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the datastore configuration delete datastore params
func (o *DatastoreConfigurationDeleteDatastoreParams) WithID(id string) *DatastoreConfigurationDeleteDatastoreParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the datastore configuration delete datastore params
func (o *DatastoreConfigurationDeleteDatastoreParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DatastoreConfigurationDeleteDatastoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
