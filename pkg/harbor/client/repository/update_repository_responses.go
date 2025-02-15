// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

// UpdateRepositoryReader is a Reader for the UpdateRepository structure.
type UpdateRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRepositoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRepositoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateRepositoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRepositoryOK creates a UpdateRepositoryOK with default headers values
func NewUpdateRepositoryOK() *UpdateRepositoryOK {
	return &UpdateRepositoryOK{}
}

/* UpdateRepositoryOK describes a response with status code 200, with default header values.

Success
*/
type UpdateRepositoryOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateRepositoryOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name}/repositories/{repository_name}][%d] updateRepositoryOK ", 200)
}

func (o *UpdateRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewUpdateRepositoryBadRequest creates a UpdateRepositoryBadRequest with default headers values
func NewUpdateRepositoryBadRequest() *UpdateRepositoryBadRequest {
	return &UpdateRepositoryBadRequest{}
}

/* UpdateRepositoryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateRepositoryBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateRepositoryBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name}/repositories/{repository_name}][%d] updateRepositoryBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateRepositoryBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRepositoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRepositoryUnauthorized creates a UpdateRepositoryUnauthorized with default headers values
func NewUpdateRepositoryUnauthorized() *UpdateRepositoryUnauthorized {
	return &UpdateRepositoryUnauthorized{}
}

/* UpdateRepositoryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateRepositoryUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name}/repositories/{repository_name}][%d] updateRepositoryUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateRepositoryUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRepositoryForbidden creates a UpdateRepositoryForbidden with default headers values
func NewUpdateRepositoryForbidden() *UpdateRepositoryForbidden {
	return &UpdateRepositoryForbidden{}
}

/* UpdateRepositoryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateRepositoryForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateRepositoryForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name}/repositories/{repository_name}][%d] updateRepositoryForbidden  %+v", 403, o.Payload)
}
func (o *UpdateRepositoryForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRepositoryNotFound creates a UpdateRepositoryNotFound with default headers values
func NewUpdateRepositoryNotFound() *UpdateRepositoryNotFound {
	return &UpdateRepositoryNotFound{}
}

/* UpdateRepositoryNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateRepositoryNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateRepositoryNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name}/repositories/{repository_name}][%d] updateRepositoryNotFound  %+v", 404, o.Payload)
}
func (o *UpdateRepositoryNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRepositoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRepositoryInternalServerError creates a UpdateRepositoryInternalServerError with default headers values
func NewUpdateRepositoryInternalServerError() *UpdateRepositoryInternalServerError {
	return &UpdateRepositoryInternalServerError{}
}

/* UpdateRepositoryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateRepositoryInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateRepositoryInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name}/repositories/{repository_name}][%d] updateRepositoryInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateRepositoryInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRepositoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
