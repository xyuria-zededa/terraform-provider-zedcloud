package image

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

// NewImageConfigurationUploadImageFileParams creates a new ImageConfigurationUploadImageFileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageConfigurationUploadImageFileParams() *ImageConfigurationUploadImageFileParams {
	return &ImageConfigurationUploadImageFileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageConfigurationUploadImageFileParamsWithTimeout creates a new ImageConfigurationUploadImageFileParams object
// with the ability to set a timeout on a request.
func NewImageConfigurationUploadImageFileParamsWithTimeout(timeout time.Duration) *ImageConfigurationUploadImageFileParams {
	return &ImageConfigurationUploadImageFileParams{
		timeout: timeout,
	}
}

// NewImageConfigurationUploadImageFileParamsWithContext creates a new ImageConfigurationUploadImageFileParams object
// with the ability to set a context for a request.
func NewImageConfigurationUploadImageFileParamsWithContext(ctx context.Context) *ImageConfigurationUploadImageFileParams {
	return &ImageConfigurationUploadImageFileParams{
		Context: ctx,
	}
}

// NewImageConfigurationUploadImageFileParamsWithHTTPClient creates a new ImageConfigurationUploadImageFileParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageConfigurationUploadImageFileParamsWithHTTPClient(client *http.Client) *ImageConfigurationUploadImageFileParams {
	return &ImageConfigurationUploadImageFileParams{
		HTTPClient: client,
	}
}

/*
ImageConfigurationUploadImageFileParams contains all the parameters to send to the API endpoint

	for the image configuration upload image file operation.

	Typically these are written to a http.Request.
*/
type ImageConfigurationUploadImageFileParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ImageFile.

	   Image binary file
	*/
	ImageFile runtime.NamedReadCloser

	/* Name.

	   User defined name of the image, unique across the enterprise. Once image is created, name can’t be changed.
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image configuration upload image file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationUploadImageFileParams) WithDefaults() *ImageConfigurationUploadImageFileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image configuration upload image file params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageConfigurationUploadImageFileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image configuration upload image file params
func (o *ImageConfigurationUploadImageFileParams) WithTimeout(timeout time.Duration) *ImageConfigurationUploadImageFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image configuration upload image file params
func (o *ImageConfigurationUploadImageFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image configuration upload image file params
func (o *ImageConfigurationUploadImageFileParams) WithContext(ctx context.Context) *ImageConfigurationUploadImageFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image configuration upload image file params
func (o *ImageConfigurationUploadImageFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image configuration upload image file params
func (o *ImageConfigurationUploadImageFileParams) WithHTTPClient(client *http.Client) *ImageConfigurationUploadImageFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image configuration upload image file params
func (o *ImageConfigurationUploadImageFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the image configuration upload image file params
func (o *ImageConfigurationUploadImageFileParams) WithXRequestID(xRequestID *string) *ImageConfigurationUploadImageFileParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the image configuration upload image file params
func (o *ImageConfigurationUploadImageFileParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithImageFile adds the imageFile to the image configuration upload image file params
func (o *ImageConfigurationUploadImageFileParams) WithImageFile(imageFile runtime.NamedReadCloser) *ImageConfigurationUploadImageFileParams {
	o.SetImageFile(imageFile)
	return o
}

// SetImageFile adds the imageFile to the image configuration upload image file params
func (o *ImageConfigurationUploadImageFileParams) SetImageFile(imageFile runtime.NamedReadCloser) {
	o.ImageFile = imageFile
}

// WithName adds the name to the image configuration upload image file params
func (o *ImageConfigurationUploadImageFileParams) WithName(name string) *ImageConfigurationUploadImageFileParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the image configuration upload image file params
func (o *ImageConfigurationUploadImageFileParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ImageConfigurationUploadImageFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	// form file param imageFile
	if err := r.SetFileParam("imageFile", o.ImageFile); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}