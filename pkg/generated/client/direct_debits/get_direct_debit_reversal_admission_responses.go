// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// GetDirectDebitReversalAdmissionReader is a Reader for the GetDirectDebitReversalAdmission structure.
type GetDirectDebitReversalAdmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDirectDebitReversalAdmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDirectDebitReversalAdmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDirectDebitReversalAdmissionOK creates a GetDirectDebitReversalAdmissionOK with default headers values
func NewGetDirectDebitReversalAdmissionOK() *GetDirectDebitReversalAdmissionOK {
	return &GetDirectDebitReversalAdmissionOK{}
}

/*
GetDirectDebitReversalAdmissionOK handles this case with default header values.

Reversal admission details
*/
type GetDirectDebitReversalAdmissionOK struct {

	//Payload

	// isStream: false
	*models.DirectDebitReversalAdmissionDetailsResponse
}

func (o *GetDirectDebitReversalAdmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/directdebits/{id}/reversals/{reversalId}/admissions/{admissionId}][%d] getDirectDebitReversalAdmissionOK", 200)
}

func (o *GetDirectDebitReversalAdmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitReversalAdmissionDetailsResponse = new(models.DirectDebitReversalAdmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitReversalAdmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
