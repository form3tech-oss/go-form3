// Code generated by go-swagger; DO NOT EDIT.

package addressbook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// GetAddressbookPartiesIDProductsProductIDReader is a Reader for the GetAddressbookPartiesIDProductsProductID structure.
type GetAddressbookPartiesIDProductsProductIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAddressbookPartiesIDProductsProductIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAddressbookPartiesIDProductsProductIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetAddressbookPartiesIDProductsProductIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAddressbookPartiesIDProductsProductIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAddressbookPartiesIDProductsProductIDOK creates a GetAddressbookPartiesIDProductsProductIDOK with default headers values
func NewGetAddressbookPartiesIDProductsProductIDOK() *GetAddressbookPartiesIDProductsProductIDOK {
	return &GetAddressbookPartiesIDProductsProductIDOK{}
}

/*GetAddressbookPartiesIDProductsProductIDOK handles this case with default header values.

get reponse
*/
type GetAddressbookPartiesIDProductsProductIDOK struct {

	//Payload

	// isStream: false
	*models.PartyProductGetResponse
}

func (o *GetAddressbookPartiesIDProductsProductIDOK) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties/{id}/products/{product_id}][%d] getAddressbookPartiesIdProductsProductIdOK", 200)
}

func (o *GetAddressbookPartiesIDProductsProductIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.PartyProductGetResponse = new(models.PartyProductGetResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.PartyProductGetResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAddressbookPartiesIDProductsProductIDForbidden creates a GetAddressbookPartiesIDProductsProductIDForbidden with default headers values
func NewGetAddressbookPartiesIDProductsProductIDForbidden() *GetAddressbookPartiesIDProductsProductIDForbidden {
	return &GetAddressbookPartiesIDProductsProductIDForbidden{}
}

/*GetAddressbookPartiesIDProductsProductIDForbidden handles this case with default header values.

Forbidden
*/
type GetAddressbookPartiesIDProductsProductIDForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAddressbookPartiesIDProductsProductIDForbidden) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties/{id}/products/{product_id}][%d] getAddressbookPartiesIdProductsProductIdForbidden", 403)
}

func (o *GetAddressbookPartiesIDProductsProductIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAddressbookPartiesIDProductsProductIDNotFound creates a GetAddressbookPartiesIDProductsProductIDNotFound with default headers values
func NewGetAddressbookPartiesIDProductsProductIDNotFound() *GetAddressbookPartiesIDProductsProductIDNotFound {
	return &GetAddressbookPartiesIDProductsProductIDNotFound{}
}

/*GetAddressbookPartiesIDProductsProductIDNotFound handles this case with default header values.

Record not found
*/
type GetAddressbookPartiesIDProductsProductIDNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAddressbookPartiesIDProductsProductIDNotFound) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties/{id}/products/{product_id}][%d] getAddressbookPartiesIdProductsProductIdNotFound", 404)
}

func (o *GetAddressbookPartiesIDProductsProductIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
