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

// PostAddressbookPartiesIDContactsContactIDAccountsReader is a Reader for the PostAddressbookPartiesIDContactsContactIDAccounts structure.
type PostAddressbookPartiesIDContactsContactIDAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAddressbookPartiesIDContactsContactIDAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostAddressbookPartiesIDContactsContactIDAccountsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostAddressbookPartiesIDContactsContactIDAccountsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostAddressbookPartiesIDContactsContactIDAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostAddressbookPartiesIDContactsContactIDAccountsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostAddressbookPartiesIDContactsContactIDAccountsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAddressbookPartiesIDContactsContactIDAccountsCreated creates a PostAddressbookPartiesIDContactsContactIDAccountsCreated with default headers values
func NewPostAddressbookPartiesIDContactsContactIDAccountsCreated() *PostAddressbookPartiesIDContactsContactIDAccountsCreated {
	return &PostAddressbookPartiesIDContactsContactIDAccountsCreated{}
}

/*PostAddressbookPartiesIDContactsContactIDAccountsCreated handles this case with default header values.

creation response
*/
type PostAddressbookPartiesIDContactsContactIDAccountsCreated struct {

	//Payload

	// isStream: false
	*models.ContactAccountCreationResponse
}

func (o *PostAddressbookPartiesIDContactsContactIDAccountsCreated) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/contacts/{contact_id}/accounts][%d] postAddressbookPartiesIdContactsContactIdAccountsCreated", 201)
}

func (o *PostAddressbookPartiesIDContactsContactIDAccountsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ContactAccountCreationResponse = new(models.ContactAccountCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ContactAccountCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAddressbookPartiesIDContactsContactIDAccountsBadRequest creates a PostAddressbookPartiesIDContactsContactIDAccountsBadRequest with default headers values
func NewPostAddressbookPartiesIDContactsContactIDAccountsBadRequest() *PostAddressbookPartiesIDContactsContactIDAccountsBadRequest {
	return &PostAddressbookPartiesIDContactsContactIDAccountsBadRequest{}
}

/*PostAddressbookPartiesIDContactsContactIDAccountsBadRequest handles this case with default header values.

Bad Request
*/
type PostAddressbookPartiesIDContactsContactIDAccountsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostAddressbookPartiesIDContactsContactIDAccountsBadRequest) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/contacts/{contact_id}/accounts][%d] postAddressbookPartiesIdContactsContactIdAccountsBadRequest", 400)
}

func (o *PostAddressbookPartiesIDContactsContactIDAccountsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAddressbookPartiesIDContactsContactIDAccountsForbidden creates a PostAddressbookPartiesIDContactsContactIDAccountsForbidden with default headers values
func NewPostAddressbookPartiesIDContactsContactIDAccountsForbidden() *PostAddressbookPartiesIDContactsContactIDAccountsForbidden {
	return &PostAddressbookPartiesIDContactsContactIDAccountsForbidden{}
}

/*PostAddressbookPartiesIDContactsContactIDAccountsForbidden handles this case with default header values.

Forbidden
*/
type PostAddressbookPartiesIDContactsContactIDAccountsForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostAddressbookPartiesIDContactsContactIDAccountsForbidden) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/contacts/{contact_id}/accounts][%d] postAddressbookPartiesIdContactsContactIdAccountsForbidden", 403)
}

func (o *PostAddressbookPartiesIDContactsContactIDAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAddressbookPartiesIDContactsContactIDAccountsNotFound creates a PostAddressbookPartiesIDContactsContactIDAccountsNotFound with default headers values
func NewPostAddressbookPartiesIDContactsContactIDAccountsNotFound() *PostAddressbookPartiesIDContactsContactIDAccountsNotFound {
	return &PostAddressbookPartiesIDContactsContactIDAccountsNotFound{}
}

/*PostAddressbookPartiesIDContactsContactIDAccountsNotFound handles this case with default header values.

Not Found
*/
type PostAddressbookPartiesIDContactsContactIDAccountsNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostAddressbookPartiesIDContactsContactIDAccountsNotFound) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/contacts/{contact_id}/accounts][%d] postAddressbookPartiesIdContactsContactIdAccountsNotFound", 404)
}

func (o *PostAddressbookPartiesIDContactsContactIDAccountsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAddressbookPartiesIDContactsContactIDAccountsConflict creates a PostAddressbookPartiesIDContactsContactIDAccountsConflict with default headers values
func NewPostAddressbookPartiesIDContactsContactIDAccountsConflict() *PostAddressbookPartiesIDContactsContactIDAccountsConflict {
	return &PostAddressbookPartiesIDContactsContactIDAccountsConflict{}
}

/*PostAddressbookPartiesIDContactsContactIDAccountsConflict handles this case with default header values.

Conflict
*/
type PostAddressbookPartiesIDContactsContactIDAccountsConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostAddressbookPartiesIDContactsContactIDAccountsConflict) Error() string {
	return fmt.Sprintf("[POST /addressbook/parties/{id}/contacts/{contact_id}/accounts][%d] postAddressbookPartiesIdContactsContactIdAccountsConflict", 409)
}

func (o *PostAddressbookPartiesIDContactsContactIDAccountsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
