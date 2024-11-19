// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// ListRolesReader is a Reader for the ListRoles structure.
type ListRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRolesOK creates a ListRolesOK with default headers values
func NewListRolesOK() *ListRolesOK {
	return &ListRolesOK{}
}

/*
ListRolesOK handles this case with default header values.

List of role details
*/
type ListRolesOK struct {

	//Payload

	// isStream: false
	*models.RoleDetailsListResponse
}

func (o *ListRolesOK) Error() string {
	return fmt.Sprintf("[GET /security/roles][%d] listRolesOK", 200)
}

func (o *ListRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.RoleDetailsListResponse = new(models.RoleDetailsListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.RoleDetailsListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
