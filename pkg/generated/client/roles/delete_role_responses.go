// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// DeleteRoleReader is a Reader for the DeleteRole structure.
type DeleteRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteRoleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteRoleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRoleNoContent creates a DeleteRoleNoContent with default headers values
func NewDeleteRoleNoContent() *DeleteRoleNoContent {
	return &DeleteRoleNoContent{}
}

/*
DeleteRoleNoContent handles this case with default header values.

Role deleted
*/
type DeleteRoleNoContent struct {
}

// IsSuccess returns true when this delete role no content response has a 2xx status code
func (o *DeleteRoleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete role no content response has a 3xx status code
func (o *DeleteRoleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete role no content response has a 4xx status code
func (o *DeleteRoleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete role no content response has a 5xx status code
func (o *DeleteRoleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete role no content response a status code equal to that given
func (o *DeleteRoleNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete role no content response
func (o *DeleteRoleNoContent) Code() int {
	return 204
}

func (o *DeleteRoleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /security/roles/{role_id}][%d] deleteRoleNoContent", 204)
}

func (o *DeleteRoleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoleNotFound creates a DeleteRoleNotFound with default headers values
func NewDeleteRoleNotFound() *DeleteRoleNotFound {
	return &DeleteRoleNotFound{}
}

/*
DeleteRoleNotFound handles this case with default header values.

Not Found
*/
type DeleteRoleNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this delete role not found response has a 2xx status code
func (o *DeleteRoleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete role not found response has a 3xx status code
func (o *DeleteRoleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete role not found response has a 4xx status code
func (o *DeleteRoleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete role not found response has a 5xx status code
func (o *DeleteRoleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete role not found response a status code equal to that given
func (o *DeleteRoleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete role not found response
func (o *DeleteRoleNotFound) Code() int {
	return 404
}

func (o *DeleteRoleNotFound) Error() string {
	return fmt.Sprintf("[DELETE /security/roles/{role_id}][%d] deleteRoleNotFound", 404)
}

func (o *DeleteRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoleConflict creates a DeleteRoleConflict with default headers values
func NewDeleteRoleConflict() *DeleteRoleConflict {
	return &DeleteRoleConflict{}
}

/*
DeleteRoleConflict handles this case with default header values.

Conflict
*/
type DeleteRoleConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this delete role conflict response has a 2xx status code
func (o *DeleteRoleConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete role conflict response has a 3xx status code
func (o *DeleteRoleConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete role conflict response has a 4xx status code
func (o *DeleteRoleConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete role conflict response has a 5xx status code
func (o *DeleteRoleConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this delete role conflict response a status code equal to that given
func (o *DeleteRoleConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the delete role conflict response
func (o *DeleteRoleConflict) Code() int {
	return 409
}

func (o *DeleteRoleConflict) Error() string {
	return fmt.Sprintf("[DELETE /security/roles/{role_id}][%d] deleteRoleConflict", 409)
}

func (o *DeleteRoleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
