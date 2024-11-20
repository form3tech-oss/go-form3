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

// GetPaymentAdviceSubmissionReader is a Reader for the GetPaymentAdviceSubmission structure.
type GetPaymentAdviceSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentAdviceSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentAdviceSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentAdviceSubmissionOK creates a GetPaymentAdviceSubmissionOK with default headers values
func NewGetPaymentAdviceSubmissionOK() *GetPaymentAdviceSubmissionOK {
	return &GetPaymentAdviceSubmissionOK{}
}

/*
GetPaymentAdviceSubmissionOK handles this case with default header values.

Advice Submission Details
*/
type GetPaymentAdviceSubmissionOK struct {

	//Payload

	// isStream: false
	*models.AdviceSubmissionFetchResponse
}

// IsSuccess returns true when this get payment advice submission o k response has a 2xx status code
func (o *GetPaymentAdviceSubmissionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get payment advice submission o k response has a 3xx status code
func (o *GetPaymentAdviceSubmissionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payment advice submission o k response has a 4xx status code
func (o *GetPaymentAdviceSubmissionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get payment advice submission o k response has a 5xx status code
func (o *GetPaymentAdviceSubmissionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get payment advice submission o k response a status code equal to that given
func (o *GetPaymentAdviceSubmissionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get payment advice submission o k response
func (o *GetPaymentAdviceSubmissionOK) Code() int {
	return 200
}

func (o *GetPaymentAdviceSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/advices/{adviceId}/submissions/{submissionId}][%d] getPaymentAdviceSubmissionOK", 200)
}

func (o *GetPaymentAdviceSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AdviceSubmissionFetchResponse = new(models.AdviceSubmissionFetchResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AdviceSubmissionFetchResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
