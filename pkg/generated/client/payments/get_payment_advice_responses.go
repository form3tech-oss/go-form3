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

// GetPaymentAdviceReader is a Reader for the GetPaymentAdvice structure.
type GetPaymentAdviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentAdviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentAdviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentAdviceOK creates a GetPaymentAdviceOK with default headers values
func NewGetPaymentAdviceOK() *GetPaymentAdviceOK {
	return &GetPaymentAdviceOK{}
}

/*GetPaymentAdviceOK handles this case with default header values.

Advice details
*/
type GetPaymentAdviceOK struct {

	//Payload

	// isStream: false
	*models.AdviceDetailsResponse
}

func (o *GetPaymentAdviceOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/advices/{adviceId}][%d] getPaymentAdviceOK", 200)
}

func (o *GetPaymentAdviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AdviceDetailsResponse = new(models.AdviceDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AdviceDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
