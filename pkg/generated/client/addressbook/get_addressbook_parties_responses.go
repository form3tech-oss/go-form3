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

// GetAddressbookPartiesReader is a Reader for the GetAddressbookParties structure.
type GetAddressbookPartiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAddressbookPartiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAddressbookPartiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAddressbookPartiesOK creates a GetAddressbookPartiesOK with default headers values
func NewGetAddressbookPartiesOK() *GetAddressbookPartiesOK {
	return &GetAddressbookPartiesOK{}
}

/*GetAddressbookPartiesOK handles this case with default header values.

list parties response
*/
type GetAddressbookPartiesOK struct {

	//Payload

	// isStream: false
	*models.ListPartiesResponse
}

func (o *GetAddressbookPartiesOK) Error() string {
	return fmt.Sprintf("[GET /addressbook/parties][%d] getAddressbookPartiesOK", 200)
}

func (o *GetAddressbookPartiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ListPartiesResponse = new(models.ListPartiesResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ListPartiesResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
