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

/*GetPaymentRecallAdmissionOK handles this case with default header values.

Recall admission details
*/
type GetPaymentRecallAdmissionOK struct {

	//Payload

	// isStream: false
	*models.RecallAdmissionDetailsResponse
}

func (o *GetPaymentRecallAdmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/recalls/{recallId}/admissions/{admissionId}][%d] getPaymentRecallAdmissionOK  %+v", 200, o)
}

func (o *GetPaymentRecallAdmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.RecallAdmissionDetailsResponse = new(models.RecallAdmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.RecallAdmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
