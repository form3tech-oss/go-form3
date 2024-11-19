// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// DeleteUserReader is a Reader for the DeleteUser structure.
type DeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewDeleteUserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserNoContent creates a DeleteUserNoContent with default headers values
func NewDeleteUserNoContent() *DeleteUserNoContent {
	return &DeleteUserNoContent{}
}

/*
DeleteUserNoContent handles this case with default header values.

User deleted
*/
type DeleteUserNoContent struct {
}

func (o *DeleteUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /security/users/{user_id}][%d] deleteUserNoContent", 204)
}

func (o *DeleteUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserNotFound creates a DeleteUserNotFound with default headers values
func NewDeleteUserNotFound() *DeleteUserNotFound {
	return &DeleteUserNotFound{}
}

/*
DeleteUserNotFound handles this case with default header values.

Not Found
*/
type DeleteUserNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *DeleteUserNotFound) Error() string {
	return fmt.Sprintf("[DELETE /security/users/{user_id}][%d] deleteUserNotFound", 404)
}

func (o *DeleteUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserConflict creates a DeleteUserConflict with default headers values
func NewDeleteUserConflict() *DeleteUserConflict {
	return &DeleteUserConflict{}
}

/*
DeleteUserConflict handles this case with default header values.

Conflict
*/
type DeleteUserConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *DeleteUserConflict) Error() string {
	return fmt.Sprintf("[DELETE /security/users/{user_id}][%d] deleteUserConflict", 409)
}

func (o *DeleteUserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
