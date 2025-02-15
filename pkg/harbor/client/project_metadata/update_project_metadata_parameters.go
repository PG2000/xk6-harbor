// Code generated by go-swagger; DO NOT EDIT.

package project_metadata

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

// NewUpdateProjectMetadataParams creates a new UpdateProjectMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateProjectMetadataParams() *UpdateProjectMetadataParams {
	return &UpdateProjectMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProjectMetadataParamsWithTimeout creates a new UpdateProjectMetadataParams object
// with the ability to set a timeout on a request.
func NewUpdateProjectMetadataParamsWithTimeout(timeout time.Duration) *UpdateProjectMetadataParams {
	return &UpdateProjectMetadataParams{
		timeout: timeout,
	}
}

// NewUpdateProjectMetadataParamsWithContext creates a new UpdateProjectMetadataParams object
// with the ability to set a context for a request.
func NewUpdateProjectMetadataParamsWithContext(ctx context.Context) *UpdateProjectMetadataParams {
	return &UpdateProjectMetadataParams{
		Context: ctx,
	}
}

// NewUpdateProjectMetadataParamsWithHTTPClient creates a new UpdateProjectMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateProjectMetadataParamsWithHTTPClient(client *http.Client) *UpdateProjectMetadataParams {
	return &UpdateProjectMetadataParams{
		HTTPClient: client,
	}
}

/* UpdateProjectMetadataParams contains all the parameters to send to the API endpoint
   for the update project metadata operation.

   Typically these are written to a http.Request.
*/
type UpdateProjectMetadataParams struct {

	/* XIsResourceName.

	   The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.
	*/
	XIsResourceName *bool `js:"xIsResourceName"`

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	/* MetaName.

	   The name of metadata.
	*/
	MetaName string `js:"metaName"`

	// Metadata.
	Metadata map[string]string `js:"metadata"`

	/* ProjectNameOrID.

	   The name or id of the project
	*/
	ProjectNameOrID string `js:"projectNameOrID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the update project metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProjectMetadataParams) WithDefaults() *UpdateProjectMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update project metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProjectMetadataParams) SetDefaults() {
	var (
		xIsResourceNameDefault = bool(false)
	)

	val := UpdateProjectMetadataParams{
		XIsResourceName: &xIsResourceNameDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update project metadata params
func (o *UpdateProjectMetadataParams) WithTimeout(timeout time.Duration) *UpdateProjectMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update project metadata params
func (o *UpdateProjectMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update project metadata params
func (o *UpdateProjectMetadataParams) WithContext(ctx context.Context) *UpdateProjectMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update project metadata params
func (o *UpdateProjectMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update project metadata params
func (o *UpdateProjectMetadataParams) WithHTTPClient(client *http.Client) *UpdateProjectMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update project metadata params
func (o *UpdateProjectMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the update project metadata params
func (o *UpdateProjectMetadataParams) WithXIsResourceName(xIsResourceName *bool) *UpdateProjectMetadataParams {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the update project metadata params
func (o *UpdateProjectMetadataParams) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the update project metadata params
func (o *UpdateProjectMetadataParams) WithXRequestID(xRequestID *string) *UpdateProjectMetadataParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the update project metadata params
func (o *UpdateProjectMetadataParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithMetaName adds the metaName to the update project metadata params
func (o *UpdateProjectMetadataParams) WithMetaName(metaName string) *UpdateProjectMetadataParams {
	o.SetMetaName(metaName)
	return o
}

// SetMetaName adds the metaName to the update project metadata params
func (o *UpdateProjectMetadataParams) SetMetaName(metaName string) {
	o.MetaName = metaName
}

// WithMetadata adds the metadata to the update project metadata params
func (o *UpdateProjectMetadataParams) WithMetadata(metadata map[string]string) *UpdateProjectMetadataParams {
	o.SetMetadata(metadata)
	return o
}

// SetMetadata adds the metadata to the update project metadata params
func (o *UpdateProjectMetadataParams) SetMetadata(metadata map[string]string) {
	o.Metadata = metadata
}

// WithProjectNameOrID adds the projectNameOrID to the update project metadata params
func (o *UpdateProjectMetadataParams) WithProjectNameOrID(projectNameOrID string) *UpdateProjectMetadataParams {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the update project metadata params
func (o *UpdateProjectMetadataParams) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProjectMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param meta_name
	if err := r.SetPathParam("meta_name", o.MetaName); err != nil {
		return err
	}
	if o.Metadata != nil {
		if err := r.SetBodyParam(o.Metadata); err != nil {
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
