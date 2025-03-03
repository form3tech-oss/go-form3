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

// GetDirectDebitReturnReader is a Reader for the GetDirectDebitReturn structure.
type GetDirectDebitReturnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDirectDebitReturnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDirectDebitReturnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDirectDebitReturnOK creates a GetDirectDebitReturnOK with default headers values
func NewGetDirectDebitReturnOK() *GetDirectDebitReturnOK {
	return &GetDirectDebitReturnOK{}
}

/*
GetDirectDebitReturnOK handles this case with default header values.

Return details
*/
type GetDirectDebitReturnOK struct {

	//Payload

	// isStream: false
	*models.DirectDebitReturnDetailsResponse
}

// IsSuccess returns true when this get direct debit return o k response has a 2xx status code
func (o *GetDirectDebitReturnOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get direct debit return o k response has a 3xx status code
func (o *GetDirectDebitReturnOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get direct debit return o k response has a 4xx status code
func (o *GetDirectDebitReturnOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get direct debit return o k response has a 5xx status code
func (o *GetDirectDebitReturnOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get direct debit return o k response a status code equal to that given
func (o *GetDirectDebitReturnOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get direct debit return o k response
func (o *GetDirectDebitReturnOK) Code() int {
	return 200
}

func (o *GetDirectDebitReturnOK) Error() string {
	return fmt.Sprintf("[GET /transaction/directdebits/{id}/returns/{returnId}][%d] getDirectDebitReturnOK", 200)
}

func (o *GetDirectDebitReturnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitReturnDetailsResponse = new(models.DirectDebitReturnDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitReturnDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
