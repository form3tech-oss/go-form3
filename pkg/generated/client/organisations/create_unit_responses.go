// Code generated by go-swagger; DO NOT EDIT.

package organisations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// CreateUnitReader is a Reader for the CreateUnit structure.
type CreateUnitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUnitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateUnitCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateUnitBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewCreateUnitConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUnitCreated creates a CreateUnitCreated with default headers values
func NewCreateUnitCreated() *CreateUnitCreated {
	return &CreateUnitCreated{}
}

/*
CreateUnitCreated handles this case with default header values.

Organisation creation response
*/
type CreateUnitCreated struct {

	//Payload

	// isStream: false
	*models.OrganisationCreationResponse
}

func (o *CreateUnitCreated) Error() string {
	return fmt.Sprintf("[POST /organisation/units][%d] createUnitCreated", 201)
}

func (o *CreateUnitCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.OrganisationCreationResponse = new(models.OrganisationCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.OrganisationCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUnitBadRequest creates a CreateUnitBadRequest with default headers values
func NewCreateUnitBadRequest() *CreateUnitBadRequest {
	return &CreateUnitBadRequest{}
}

/*
CreateUnitBadRequest handles this case with default header values.

Bad Request
*/
type CreateUnitBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateUnitBadRequest) Error() string {
	return fmt.Sprintf("[POST /organisation/units][%d] createUnitBadRequest", 400)
}

func (o *CreateUnitBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUnitConflict creates a CreateUnitConflict with default headers values
func NewCreateUnitConflict() *CreateUnitConflict {
	return &CreateUnitConflict{}
}

/*
CreateUnitConflict handles this case with default header values.

Conflict
*/
type CreateUnitConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *CreateUnitConflict) Error() string {
	return fmt.Sprintf("[POST /organisation/units][%d] createUnitConflict", 409)
}

func (o *CreateUnitConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
