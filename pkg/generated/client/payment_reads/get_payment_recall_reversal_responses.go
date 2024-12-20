// Code generated by go-swagger; DO NOT EDIT.

package payment_reads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// GetPaymentRecallReversalReader is a Reader for the GetPaymentRecallReversal structure.
type GetPaymentRecallReversalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentRecallReversalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentRecallReversalOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentRecallReversalOK creates a GetPaymentRecallReversalOK with default headers values
func NewGetPaymentRecallReversalOK() *GetPaymentRecallReversalOK {
	return &GetPaymentRecallReversalOK{}
}

/*
GetPaymentRecallReversalOK handles this case with default header values.

Recall Reversal details
*/
type GetPaymentRecallReversalOK struct {

	//Payload

	// isStream: false
	*models.RecallReversalDetailsResponse
}

// IsSuccess returns true when this get payment recall reversal o k response has a 2xx status code
func (o *GetPaymentRecallReversalOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get payment recall reversal o k response has a 3xx status code
func (o *GetPaymentRecallReversalOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payment recall reversal o k response has a 4xx status code
func (o *GetPaymentRecallReversalOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get payment recall reversal o k response has a 5xx status code
func (o *GetPaymentRecallReversalOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get payment recall reversal o k response a status code equal to that given
func (o *GetPaymentRecallReversalOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get payment recall reversal o k response
func (o *GetPaymentRecallReversalOK) Code() int {
	return 200
}

func (o *GetPaymentRecallReversalOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/recalls/{recallId}/reversals/{reversalId}][%d] getPaymentRecallReversalOK", 200)
}

func (o *GetPaymentRecallReversalOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.RecallReversalDetailsResponse = new(models.RecallReversalDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.RecallReversalDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
