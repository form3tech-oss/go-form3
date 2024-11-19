// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// GetUserReader is a Reader for the GetUser structure.
type GetUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserOK creates a GetUserOK with default headers values
func NewGetUserOK() *GetUserOK {
	return &GetUserOK{}
}

/*
GetUserOK handles this case with default header values.

User details
*/
type GetUserOK struct {

	//Payload

	// isStream: false
	*models.UserDetailsResponse
}

func (o *GetUserOK) Error() string {
	return fmt.Sprintf("[GET /security/users/{user_id}][%d] getUserOK", 200)
}

func (o *GetUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.UserDetailsResponse = new(models.UserDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.UserDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserNotFound creates a GetUserNotFound with default headers values
func NewGetUserNotFound() *GetUserNotFound {
	return &GetUserNotFound{}
}

/*
GetUserNotFound handles this case with default header values.

Not Found
*/
type GetUserNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetUserNotFound) Error() string {
	return fmt.Sprintf("[GET /security/users/{user_id}][%d] getUserNotFound", 404)
}

func (o *GetUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
