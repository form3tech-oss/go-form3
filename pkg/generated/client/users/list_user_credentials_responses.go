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

// ListUserCredentialsReader is a Reader for the ListUserCredentials structure.
type ListUserCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUserCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListUserCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewListUserCredentialsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListUserCredentialsOK creates a ListUserCredentialsOK with default headers values
func NewListUserCredentialsOK() *ListUserCredentialsOK {
	return &ListUserCredentialsOK{}
}

/*
ListUserCredentialsOK handles this case with default header values.

List of credentials for user
*/
type ListUserCredentialsOK struct {

	//Payload

	// isStream: false
	*models.UserCredentialListResponse
}

func (o *ListUserCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /security/users/{user_id}/credentials][%d] listUserCredentialsOK", 200)
}

func (o *ListUserCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.UserCredentialListResponse = new(models.UserCredentialListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.UserCredentialListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserCredentialsNotFound creates a ListUserCredentialsNotFound with default headers values
func NewListUserCredentialsNotFound() *ListUserCredentialsNotFound {
	return &ListUserCredentialsNotFound{}
}

/*
ListUserCredentialsNotFound handles this case with default header values.

Not Found
*/
type ListUserCredentialsNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListUserCredentialsNotFound) Error() string {
	return fmt.Sprintf("[GET /security/users/{user_id}/credentials][%d] listUserCredentialsNotFound", 404)
}

func (o *ListUserCredentialsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
