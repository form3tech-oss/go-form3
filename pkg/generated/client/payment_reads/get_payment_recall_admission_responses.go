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

// GetPaymentRecallAdmissionReader is a Reader for the GetPaymentRecallAdmission structure.
type GetPaymentRecallAdmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentRecallAdmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentRecallAdmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentRecallAdmissionOK creates a GetPaymentRecallAdmissionOK with default headers values
func NewGetPaymentRecallAdmissionOK() *GetPaymentRecallAdmissionOK {
	return &GetPaymentRecallAdmissionOK{}
}

/*
GetPaymentRecallAdmissionOK handles this case with default header values.

Recall Admission Details
*/
type GetPaymentRecallAdmissionOK struct {

	//Payload

	// isStream: false
	*models.RecallAdmissionFetchResponse
}

// IsSuccess returns true when this get payment recall admission o k response has a 2xx status code
func (o *GetPaymentRecallAdmissionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get payment recall admission o k response has a 3xx status code
func (o *GetPaymentRecallAdmissionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payment recall admission o k response has a 4xx status code
func (o *GetPaymentRecallAdmissionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get payment recall admission o k response has a 5xx status code
func (o *GetPaymentRecallAdmissionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get payment recall admission o k response a status code equal to that given
func (o *GetPaymentRecallAdmissionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get payment recall admission o k response
func (o *GetPaymentRecallAdmissionOK) Code() int {
	return 200
}

func (o *GetPaymentRecallAdmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/recalls/{recallId}/admissions/{admissionId}][%d] getPaymentRecallAdmissionOK", 200)
}

func (o *GetPaymentRecallAdmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.RecallAdmissionFetchResponse = new(models.RecallAdmissionFetchResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.RecallAdmissionFetchResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
