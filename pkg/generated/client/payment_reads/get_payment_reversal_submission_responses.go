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

// GetPaymentReversalSubmissionReader is a Reader for the GetPaymentReversalSubmission structure.
type GetPaymentReversalSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentReversalSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentReversalSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentReversalSubmissionOK creates a GetPaymentReversalSubmissionOK with default headers values
func NewGetPaymentReversalSubmissionOK() *GetPaymentReversalSubmissionOK {
	return &GetPaymentReversalSubmissionOK{}
}

/*
GetPaymentReversalSubmissionOK handles this case with default header values.

Reversal Submission Details
*/
type GetPaymentReversalSubmissionOK struct {

	//Payload

	// isStream: false
	*models.ReversalSubmissionFetchResponse
}

// IsSuccess returns true when this get payment reversal submission o k response has a 2xx status code
func (o *GetPaymentReversalSubmissionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get payment reversal submission o k response has a 3xx status code
func (o *GetPaymentReversalSubmissionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payment reversal submission o k response has a 4xx status code
func (o *GetPaymentReversalSubmissionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get payment reversal submission o k response has a 5xx status code
func (o *GetPaymentReversalSubmissionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get payment reversal submission o k response a status code equal to that given
func (o *GetPaymentReversalSubmissionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get payment reversal submission o k response
func (o *GetPaymentReversalSubmissionOK) Code() int {
	return 200
}

func (o *GetPaymentReversalSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/reversals/{reversalId}/submissions/{submissionId}][%d] getPaymentReversalSubmissionOK", 200)
}

func (o *GetPaymentReversalSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReversalSubmissionFetchResponse = new(models.ReversalSubmissionFetchResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReversalSubmissionFetchResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
