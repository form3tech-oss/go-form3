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

// GetPaymentReversalAdmissionReader is a Reader for the GetPaymentReversalAdmission structure.
type GetPaymentReversalAdmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentReversalAdmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentReversalAdmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentReversalAdmissionOK creates a GetPaymentReversalAdmissionOK with default headers values
func NewGetPaymentReversalAdmissionOK() *GetPaymentReversalAdmissionOK {
	return &GetPaymentReversalAdmissionOK{}
}

/*
GetPaymentReversalAdmissionOK handles this case with default header values.

Reversal admission details
*/
type GetPaymentReversalAdmissionOK struct {

	//Payload

	// isStream: false
	*models.ReversalAdmissionFetchResponse
}

// IsSuccess returns true when this get payment reversal admission o k response has a 2xx status code
func (o *GetPaymentReversalAdmissionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get payment reversal admission o k response has a 3xx status code
func (o *GetPaymentReversalAdmissionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payment reversal admission o k response has a 4xx status code
func (o *GetPaymentReversalAdmissionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get payment reversal admission o k response has a 5xx status code
func (o *GetPaymentReversalAdmissionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get payment reversal admission o k response a status code equal to that given
func (o *GetPaymentReversalAdmissionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get payment reversal admission o k response
func (o *GetPaymentReversalAdmissionOK) Code() int {
	return 200
}

func (o *GetPaymentReversalAdmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/reversals/{reversalId}/admissions/{admissionId}][%d] getPaymentReversalAdmissionOK", 200)
}

func (o *GetPaymentReversalAdmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReversalAdmissionFetchResponse = new(models.ReversalAdmissionFetchResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReversalAdmissionFetchResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}