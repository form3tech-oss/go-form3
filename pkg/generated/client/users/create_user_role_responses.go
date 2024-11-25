// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// CreateUserRoleReader is a Reader for the CreateUserRole structure.
type CreateUserRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateUserRoleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewCreateUserRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUserRoleCreated creates a CreateUserRoleCreated with default headers values
func NewCreateUserRoleCreated() *CreateUserRoleCreated {
	return &CreateUserRoleCreated{}
}

/*
CreateUserRoleCreated handles this case with default header values.

Role set OK
*/
type CreateUserRoleCreated struct {
}

// IsSuccess returns true when this create user role created response has a 2xx status code
func (o *CreateUserRoleCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user role created response has a 3xx status code
func (o *CreateUserRoleCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user role created response has a 4xx status code
func (o *CreateUserRoleCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user role created response has a 5xx status code
func (o *CreateUserRoleCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create user role created response a status code equal to that given
func (o *CreateUserRoleCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create user role created response
func (o *CreateUserRoleCreated) Code() int {
	return 201
}

func (o *CreateUserRoleCreated) Error() string {
	return fmt.Sprintf("[POST /security/users/{user_id}/roles/{role_id}][%d] createUserRoleCreated", 201)
}

func (o *CreateUserRoleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserRoleNotFound creates a CreateUserRoleNotFound with default headers values
func NewCreateUserRoleNotFound() *CreateUserRoleNotFound {
	return &CreateUserRoleNotFound{}
}

/*
CreateUserRoleNotFound handles this case with default header values.

Not Found
*/
type CreateUserRoleNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create user role not found response has a 2xx status code
func (o *CreateUserRoleNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user role not found response has a 3xx status code
func (o *CreateUserRoleNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user role not found response has a 4xx status code
func (o *CreateUserRoleNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user role not found response has a 5xx status code
func (o *CreateUserRoleNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create user role not found response a status code equal to that given
func (o *CreateUserRoleNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the create user role not found response
func (o *CreateUserRoleNotFound) Code() int {
	return 404
}

func (o *CreateUserRoleNotFound) Error() string {
	return fmt.Sprintf("[POST /security/users/{user_id}/roles/{role_id}][%d] createUserRoleNotFound", 404)
}

func (o *CreateUserRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
