// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
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

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUserRoleCreated creates a CreateUserRoleCreated with default headers values
func NewCreateUserRoleCreated() *CreateUserRoleCreated {
	return &CreateUserRoleCreated{}
}

/*CreateUserRoleCreated handles this case with default header values.

Role set OK
*/
type CreateUserRoleCreated struct {
}

func (o *CreateUserRoleCreated) Error() string {
	return fmt.Sprintf("[POST /security/users/{user_id}/roles/{role_id}][%d] createUserRoleCreated", 201)
}

func (o *CreateUserRoleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
