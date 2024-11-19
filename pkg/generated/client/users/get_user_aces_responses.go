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

// GetUserAcesReader is a Reader for the GetUserAces structure.
type GetUserAcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserAcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUserAcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetUserAcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserAcesOK creates a GetUserAcesOK with default headers values
func NewGetUserAcesOK() *GetUserAcesOK {
	return &GetUserAcesOK{}
}

/*
GetUserAcesOK handles this case with default header values.

List of access control entries for this user
*/
type GetUserAcesOK struct {

	//Payload

	// isStream: false
	*models.AceDetailsListResponse
}

func (o *GetUserAcesOK) Error() string {
	return fmt.Sprintf("[GET /security/users/{user_id}/aces][%d] getUserAcesOK", 200)
}

func (o *GetUserAcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AceDetailsListResponse = new(models.AceDetailsListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AceDetailsListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserAcesNotFound creates a GetUserAcesNotFound with default headers values
func NewGetUserAcesNotFound() *GetUserAcesNotFound {
	return &GetUserAcesNotFound{}
}

/*
GetUserAcesNotFound handles this case with default header values.

Not Found
*/
type GetUserAcesNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetUserAcesNotFound) Error() string {
	return fmt.Sprintf("[GET /security/users/{user_id}/aces][%d] getUserAcesNotFound", 404)
}

func (o *GetUserAcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
