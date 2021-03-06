// Code generated by go-swagger; DO NOT EDIT.

package addressbook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v3/pkg/generated/models"
)

// PostAddressbookPartiesIDActorsReader is a Reader for the PostAddressbookPartiesIDActors structure.
type PostAddressbookPartiesIDActorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAddressbookPartiesIDActorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostAddressbookPartiesIDActorsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostAddressbookPartiesIDActorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostAddressbookPartiesIDActorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostAddressbookPartiesIDActorsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAddressbookPartiesIDActorsCreated creates a PostAddressbookPartiesIDActorsCreated with default headers values
func NewPostAddressbookPartiesIDActorsCreated() *PostAddressbookPartiesIDActorsCreated {
	return &PostAddressbookPartiesIDActorsCreated{}
}

/*PostAddressbookPartiesIDActorsCreated handles this case with default header values.

creation response
*/
type PostAddressbookPartiesIDActorsCreated struct {

	//Payload

	// isStream: false
	*models.PartyActorCreationResponse
}

func (o *PostAddressbookPartiesIDActorsCreated) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/actors][%d] postAddressbookPartiesIdActorsCreated", 201)
}

func (o *PostAddressbookPartiesIDActorsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.PartyActorCreationResponse = new(models.PartyActorCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.PartyActorCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAddressbookPartiesIDActorsBadRequest creates a PostAddressbookPartiesIDActorsBadRequest with default headers values
func NewPostAddressbookPartiesIDActorsBadRequest() *PostAddressbookPartiesIDActorsBadRequest {
	return &PostAddressbookPartiesIDActorsBadRequest{}
}

/*PostAddressbookPartiesIDActorsBadRequest handles this case with default header values.

Bad Request
*/
type PostAddressbookPartiesIDActorsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostAddressbookPartiesIDActorsBadRequest) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/actors][%d] postAddressbookPartiesIdActorsBadRequest", 400)
}

func (o *PostAddressbookPartiesIDActorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAddressbookPartiesIDActorsForbidden creates a PostAddressbookPartiesIDActorsForbidden with default headers values
func NewPostAddressbookPartiesIDActorsForbidden() *PostAddressbookPartiesIDActorsForbidden {
	return &PostAddressbookPartiesIDActorsForbidden{}
}

/*PostAddressbookPartiesIDActorsForbidden handles this case with default header values.

Forbidden
*/
type PostAddressbookPartiesIDActorsForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostAddressbookPartiesIDActorsForbidden) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/actors][%d] postAddressbookPartiesIdActorsForbidden", 403)
}

func (o *PostAddressbookPartiesIDActorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAddressbookPartiesIDActorsConflict creates a PostAddressbookPartiesIDActorsConflict with default headers values
func NewPostAddressbookPartiesIDActorsConflict() *PostAddressbookPartiesIDActorsConflict {
	return &PostAddressbookPartiesIDActorsConflict{}
}

/*PostAddressbookPartiesIDActorsConflict handles this case with default header values.

Conflict
*/
type PostAddressbookPartiesIDActorsConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostAddressbookPartiesIDActorsConflict) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/actors][%d] postAddressbookPartiesIdActorsConflict", 409)
}

func (o *PostAddressbookPartiesIDActorsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
