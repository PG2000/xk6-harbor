// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewGetProjectParams creates a new GetProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectParams() *GetProjectParams {
	return &GetProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectParamsWithTimeout creates a new GetProjectParams object
// with the ability to set a timeout on a request.
func NewGetProjectParamsWithTimeout(timeout time.Duration) *GetProjectParams {
	return &GetProjectParams{
		timeout: timeout,
	}
}

// NewGetProjectParamsWithContext creates a new GetProjectParams object
// with the ability to set a context for a request.
func NewGetProjectParamsWithContext(ctx context.Context) *GetProjectParams {
	return &GetProjectParams{
		Context: ctx,
	}
}

// NewGetProjectParamsWithHTTPClient creates a new GetProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectParamsWithHTTPClient(client *http.Client) *GetProjectParams {
	return &GetProjectParams{
		HTTPClient: client,
	}
}

/* GetProjectParams contains all the parameters to send to the API endpoint
   for the get project operation.

   Typically these are written to a http.Request.
*/
type GetProjectParams struct {

	/* XIsResourceName.

	   The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	*/
	XIsResourceName *bool `js:"xIsResourceName"`

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* ProjectNameOrID.

	   The name or id of the project
	*/
	ProjectNameOrID string `js:"projectNameOrID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the get project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectParams) WithDefaults() *GetProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectParams) SetDefaults() {
	var (
		xIsResourceNameDefault = bool(false)
	)

	val := GetProjectParams{
		XIsResourceName: &xIsResourceNameDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get project params
func (o *GetProjectParams) WithTimeout(timeout time.Duration) *GetProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project params
func (o *GetProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project params
func (o *GetProjectParams) WithContext(ctx context.Context) *GetProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project params
func (o *GetProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project params
func (o *GetProjectParams) WithHTTPClient(client *http.Client) *GetProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project params
func (o *GetProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the get project params
func (o *GetProjectParams) WithXIsResourceName(xIsResourceName *bool) *GetProjectParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the get project params
func (o *GetProjectParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the get project params
func (o *GetProjectParams) WithXRequestID(xRequestID *string) *GetProjectParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get project params
func (o *GetProjectParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectNameOrID adds the projectNameOrID to the get project params
func (o *GetProjectParams) WithProjectNameOrID(projectNameOrID string) *GetProjectParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the get project params
func (o *GetProjectParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XIsResourceName != nil {

		// header param X-Is-Resource-Name
		if err := r.SetHeaderParam("X-Is-Resource-Name", swag.FormatBool(*o.XIsResourceName)); err != nil {
			return err
		}
	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
