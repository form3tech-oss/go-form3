// Code generated by go-swagger; DO NOT EDIT.

package addressbook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetAddressbookPartiesIDReader is a Reader for the GetAddressbookPartiesID structure.
type GetAddressbookPartiesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAddressbookPartiesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAddressbookPartiesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAddressbookPartiesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAddressbookPartiesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAddressbookPartiesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAddressbookPartiesIDOK creates a GetAddressbookPartiesIDOK with default headers values
func NewGetAddressbookPartiesIDOK() *GetAddressbookPartiesIDOK {
	return &GetAddressbookPartiesIDOK{}
}

/*GetAddressbookPartiesIDOK handles this case with default header values.

party returned
*/
type GetAddressbookPartiesIDOK struct {

	//Payload

	// isStream: false
	*models.PartyGetResponse
}

func (o *GetAddressbookPartiesIDOK) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties/{id}][%d] getAddressbookPartiesIdOK", 200)
}

func (o *GetAddressbookPartiesIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.PartyGetResponse = new(models.PartyGetResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.PartyGetResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAddressbookPartiesIDBadRequest creates a GetAddressbookPartiesIDBadRequest with default headers values
func NewGetAddressbookPartiesIDBadRequest() *GetAddressbookPartiesIDBadRequest {
	return &GetAddressbookPartiesIDBadRequest{}
}

/*GetAddressbookPartiesIDBadRequest handles this case with default header values.

Bad Request
*/
type GetAddressbookPartiesIDBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAddressbookPartiesIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties/{id}][%d] getAddressbookPartiesIdBadRequest", 400)
}

func (o *GetAddressbookPartiesIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAddressbookPartiesIDForbidden creates a GetAddressbookPartiesIDForbidden with default headers values
func NewGetAddressbookPartiesIDForbidden() *GetAddressbookPartiesIDForbidden {
	return &GetAddressbookPartiesIDForbidden{}
}

/*GetAddressbookPartiesIDForbidden handles this case with default header values.

Forbidden
*/
type GetAddressbookPartiesIDForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAddressbookPartiesIDForbidden) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties/{id}][%d] getAddressbookPartiesIdForbidden", 403)
}

func (o *GetAddressbookPartiesIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAddressbookPartiesIDNotFound creates a GetAddressbookPartiesIDNotFound with default headers values
func NewGetAddressbookPartiesIDNotFound() *GetAddressbookPartiesIDNotFound {
	return &GetAddressbookPartiesIDNotFound{}
}

/*GetAddressbookPartiesIDNotFound handles this case with default header values.

Record not found
*/
type GetAddressbookPartiesIDNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAddressbookPartiesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties/{id}][%d] getAddressbookPartiesIdNotFound", 404)
}

func (o *GetAddressbookPartiesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}