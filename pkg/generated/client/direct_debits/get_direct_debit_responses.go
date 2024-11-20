// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
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

/*
GetDirectDebitOK handles this case with default header values.

Direct Debit details
*/
type GetDirectDebitOK struct {

	//Payload

	// isStream: false
	*models.DirectDebitDetailsResponse
}

// IsSuccess returns true when this get direct debit o k response has a 2xx status code
func (o *GetDirectDebitOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get direct debit o k response has a 3xx status code
func (o *GetDirectDebitOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get direct debit o k response has a 4xx status code
func (o *GetDirectDebitOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get direct debit o k response has a 5xx status code
func (o *GetDirectDebitOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get direct debit o k response a status code equal to that given
func (o *GetDirectDebitOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get direct debit o k response
func (o *GetDirectDebitOK) Code() int {
	return 200
}

func (o *GetDirectDebitOK) Error() string {
	return fmt.Sprintf("[GET /transaction/directdebits/{id}][%d] getDirectDebitOK", 200)
}

func (o *GetDirectDebitOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitDetailsResponse = new(models.DirectDebitDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
