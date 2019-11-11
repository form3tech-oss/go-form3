// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetDirectDebitReader is a Reader for the GetDirectDebit structure.
type GetDirectDebitReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.

func (o *GetDirectDebitReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:

		result := NewGetDirectDebitOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDirectDebitOK creates a GetDirectDebitOK with default headers values
func NewGetDirectDebitOK() *GetDirectDebitOK {
	return &GetDirectDebitOK{}
}

/*GetDirectDebitOK handles this case with default header values.

Direct Debit details
*/
type GetDirectDebitOK struct {

	//Payload

	// isStream: false
	*models.DirectDebitDetailsResponse
}

func (o *GetDirectDebitOK) Error() string {
	return fmt.Sprintf("[GET /transaction/directdebits/{id}][%d] getDirectDebitOK  %+v", 200, o)
}

func (o *GetDirectDebitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitDetailsResponse = new(models.DirectDebitDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
