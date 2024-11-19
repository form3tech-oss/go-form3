// Code generated by go-swagger; DO NOT EDIT.

package a_c_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// CreateAceReader is a Reader for the CreateAce structure.
type CreateAceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateAceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateAceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateAceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateAceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAceCreated creates a CreateAceCreated with default headers values
func NewCreateAceCreated() *CreateAceCreated {
	return &CreateAceCreated{}
}

/*
CreateAceCreated handles this case with default header values.

ACE creation response
*/
type CreateAceCreated struct {

	//Payload

	// isStream: false
	*models.AceCreationResponse
}

func (o *CreateAceCreated) Error() string {
	return fmt.Sprintf("[POST /security/roles/{role_id}/aces][%d] createAceCreated", 201)
}

func (o *CreateAceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AceCreationResponse = new(models.AceCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AceCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAceBadRequest creates a CreateAceBadRequest with default headers values
func NewCreateAceBadRequest() *CreateAceBadRequest {
	return &CreateAceBadRequest{}
}

/*
CreateAceBadRequest handles this case with default header values.

Bad request
*/
type CreateAceBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAceBadRequest) Error() string {
	return fmt.Sprintf("[POST /security/roles/{role_id}/aces][%d] createAceBadRequest", 400)
}

func (o *CreateAceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAceNotFound creates a CreateAceNotFound with default headers values
func NewCreateAceNotFound() *CreateAceNotFound {
	return &CreateAceNotFound{}
}

/*
CreateAceNotFound handles this case with default header values.

Not Found
*/
type CreateAceNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAceNotFound) Error() string {
	return fmt.Sprintf("[POST /security/roles/{role_id}/aces][%d] createAceNotFound", 404)
}

func (o *CreateAceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAceConflict creates a CreateAceConflict with default headers values
func NewCreateAceConflict() *CreateAceConflict {
	return &CreateAceConflict{}
}

/*
CreateAceConflict handles this case with default header values.

Conflict
*/
type CreateAceConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAceConflict) Error() string {
	return fmt.Sprintf("[POST /security/roles/{role_id}/aces][%d] createAceConflict", 409)
}

func (o *CreateAceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
