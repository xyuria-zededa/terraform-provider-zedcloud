// Code generated by go-swagger; DO NOT EDIT.

package projects

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

	"github.com/zededa/terraform-provider/models"
)

// UpdateParams creates a new ProjectsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func UpdateParams() *ProjectsUpdateParams {
	return &ProjectsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProjectsUpdateParamsWithTimeout creates a new ProjectsUpdateParams object
// with the ability to set a timeout on a request.
func NewProjectsUpdateParamsWithTimeout(timeout time.Duration) *ProjectsUpdateParams {
	return &ProjectsUpdateParams{
		timeout: timeout,
	}
}

// NewProjectsUpdateParamsWithContext creates a new ProjectsUpdateParams object
// with the ability to set a context for a request.
func NewProjectsUpdateParamsWithContext(ctx context.Context) *ProjectsUpdateParams {
	return &ProjectsUpdateParams{
		Context: ctx,
	}
}

// NewProjectsUpdateParamsWithHTTPClient creates a new ProjectsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewProjectsUpdateParamsWithHTTPClient(client *http.Client) *ProjectsUpdateParams {
	return &ProjectsUpdateParams{
		HTTPClient: client,
	}
}

/*
ProjectsUpdateParams contains all the parameters to send to the API endpoint

	for the projects update operation.

	Typically these are written to a http.Request.
*/
type ProjectsUpdateParams struct {

	/* XRequestID.

	   correlation-id
	*/
	XRequestID *string

	// Body.
	Body *models.Tag

	/* ID.

	   System defined universally unique Id of the resource group.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the projects update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsUpdateParams) WithDefaults() *ProjectsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the projects update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProjectsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the projects update params
func (o *ProjectsUpdateParams) WithTimeout(timeout time.Duration) *ProjectsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the projects update params
func (o *ProjectsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the projects update params
func (o *ProjectsUpdateParams) WithContext(ctx context.Context) *ProjectsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the projects update params
func (o *ProjectsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the projects update params
func (o *ProjectsUpdateParams) WithHTTPClient(client *http.Client) *ProjectsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the projects update params
func (o *ProjectsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the projects update params
func (o *ProjectsUpdateParams) WithXRequestID(xRequestID *string) *ProjectsUpdateParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the projects update params
func (o *ProjectsUpdateParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithBody adds the body to the projects update params
func (o *ProjectsUpdateParams) WithBody(body *models.Tag) *ProjectsUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the projects update params
func (o *ProjectsUpdateParams) SetBody(body *models.Tag) {
	o.Body = body
}

// WithID adds the id to the projects update params
func (o *ProjectsUpdateParams) WithID(id string) *ProjectsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the projects update params
func (o *ProjectsUpdateParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ProjectsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
