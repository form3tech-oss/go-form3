// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetPaymentReader is a Reader for the GetPayment structure.
type GetPaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentOK creates a GetPaymentOK with default headers values
func NewGetPaymentOK() *GetPaymentOK {
	return &GetPaymentOK{}
}

/*GetPaymentOK handles this case with default header values.

Payment details
*/
type GetPaymentOK struct {

	//Payload
	*models.PaymentDetailsResponse
}

func (o *GetPaymentOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}][%d] getPaymentOK  %+v", 200, o)
}

func (o *GetPaymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.PaymentDetailsResponse = new(models.PaymentDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.PaymentDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
