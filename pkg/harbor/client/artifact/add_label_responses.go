// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// AddLabelReader is a Reader for the AddLabel structure.
type AddLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddLabelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddLabelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddLabelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddLabelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddLabelConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddLabelInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddLabelOK creates a AddLabelOK with default headers values
func NewAddLabelOK() *AddLabelOK {
	return &AddLabelOK{}
}

/* AddLabelOK describes a response with status code 200, with default header values.

Success
*/
type AddLabelOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *AddLabelOK) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelOK ", 200)
}

func (o *AddLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewAddLabelBadRequest creates a AddLabelBadRequest with default headers values
func NewAddLabelBadRequest() *AddLabelBadRequest {
	return &AddLabelBadRequest{}
}

/* AddLabelBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AddLabelBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *AddLabelBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelBadRequest  %+v", 400, o.Payload)
}
func (o *AddLabelBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddLabelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddLabelUnauthorized creates a AddLabelUnauthorized with default headers values
func NewAddLabelUnauthorized() *AddLabelUnauthorized {
	return &AddLabelUnauthorized{}
}

/* AddLabelUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AddLabelUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *AddLabelUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelUnauthorized  %+v", 401, o.Payload)
}
func (o *AddLabelUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddLabelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddLabelForbidden creates a AddLabelForbidden with default headers values
func NewAddLabelForbidden() *AddLabelForbidden {
	return &AddLabelForbidden{}
}

/* AddLabelForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AddLabelForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *AddLabelForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelForbidden  %+v", 403, o.Payload)
}
func (o *AddLabelForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddLabelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddLabelNotFound creates a AddLabelNotFound with default headers values
func NewAddLabelNotFound() *AddLabelNotFound {
	return &AddLabelNotFound{}
}

/* AddLabelNotFound describes a response with status code 404, with default header values.

Not found
*/
type AddLabelNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *AddLabelNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelNotFound  %+v", 404, o.Payload)
}
func (o *AddLabelNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddLabelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddLabelConflict creates a AddLabelConflict with default headers values
func NewAddLabelConflict() *AddLabelConflict {
	return &AddLabelConflict{}
}

/* AddLabelConflict describes a response with status code 409, with default header values.

Conflict
*/
type AddLabelConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *AddLabelConflict) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelConflict  %+v", 409, o.Payload)
}
func (o *AddLabelConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddLabelConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddLabelInternalServerError creates a AddLabelInternalServerError with default headers values
func NewAddLabelInternalServerError() *AddLabelInternalServerError {
	return &AddLabelInternalServerError{}
}

/* AddLabelInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type AddLabelInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *AddLabelInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/labels][%d] addLabelInternalServerError  %+v", 500, o.Payload)
}
func (o *AddLabelInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddLabelInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
