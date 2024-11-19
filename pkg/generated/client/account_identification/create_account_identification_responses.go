// Code generated by go-swagger; DO NOT EDIT.

package account_identification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// CreateAccountIdentificationReader is a Reader for the CreateAccountIdentification structure.
type CreateAccountIdentificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAccountIdentificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateAccountIdentificationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewCreateAccountIdentificationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAccountIdentificationCreated creates a CreateAccountIdentificationCreated with default headers values
func NewCreateAccountIdentificationCreated() *CreateAccountIdentificationCreated {
	return &CreateAccountIdentificationCreated{}
}

/*
CreateAccountIdentificationCreated handles this case with default header values.

Account Identification creation response
*/
type CreateAccountIdentificationCreated struct {

	//Payload

	// isStream: false
	*models.AccountIdentificationResponse
}

func (o *CreateAccountIdentificationCreated) Error() string {
	return fmt.Sprintf("[POST /organisation/accounts/{account_id}/identifications][%d] createAccountIdentificationCreated", 201)
}

func (o *CreateAccountIdentificationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AccountIdentificationResponse = new(models.AccountIdentificationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AccountIdentificationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountIdentificationConflict creates a CreateAccountIdentificationConflict with default headers values
func NewCreateAccountIdentificationConflict() *CreateAccountIdentificationConflict {
	return &CreateAccountIdentificationConflict{}
}

/*
CreateAccountIdentificationConflict handles this case with default header values.

Account Identification creation error, constraint violation of secondary identification
*/
type CreateAccountIdentificationConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateAccountIdentificationConflict) Error() string {
	return fmt.Sprintf("[POST /organisation/accounts/{account_id}/identifications][%d] createAccountIdentificationConflict", 409)
}

func (o *CreateAccountIdentificationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
