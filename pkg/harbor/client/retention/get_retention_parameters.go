// Code generated by go-swagger; DO NOT EDIT.

package retention

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
	"github.com/go-openapi/swag"
)

// NewGetRetentionParams creates a new GetRetentionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRetentionParams() *GetRetentionParams {
	return &GetRetentionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRetentionParamsWithTimeout creates a new GetRetentionParams object
// with the ability to set a timeout on a request.
func NewGetRetentionParamsWithTimeout(timeout time.Duration) *GetRetentionParams {
	return &GetRetentionParams{
		timeout: timeout,
	}
}

// NewGetRetentionParamsWithContext creates a new GetRetentionParams object
// with the ability to set a context for a request.
func NewGetRetentionParamsWithContext(ctx context.Context) *GetRetentionParams {
	return &GetRetentionParams{
		Context: ctx,
	}
}

// NewGetRetentionParamsWithHTTPClient creates a new GetRetentionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRetentionParamsWithHTTPClient(client *http.Client) *GetRetentionParams {
	return &GetRetentionParams{
		HTTPClient: client,
	}
}

/* GetRetentionParams contains all the parameters to send to the API endpoint
   for the get retention operation.

   Typically these are written to a http.Request.
*/
type GetRetentionParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* ID.

	   Retention ID.

	   Format: int64
	*/
	ID int64 `js:"id"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the get retention params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRetentionParams) WithDefaults() *GetRetentionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get retention params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRetentionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get retention params
func (o *GetRetentionParams) WithTimeout(timeout time.Duration) *GetRetentionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get retention params
func (o *GetRetentionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get retention params
func (o *GetRetentionParams) WithContext(ctx context.Context) *GetRetentionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get retention params
func (o *GetRetentionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get retention params
func (o *GetRetentionParams) WithHTTPClient(client *http.Client) *GetRetentionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get retention params
func (o *GetRetentionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the get retention params
func (o *GetRetentionParams) WithXRequestID(xRequestID *string) *GetRetentionParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get retention params
func (o *GetRetentionParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the get retention params
func (o *GetRetentionParams) WithID(id int64) *GetRetentionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get retention params
func (o *GetRetentionParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRetentionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
