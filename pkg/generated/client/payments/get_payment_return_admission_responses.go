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

// GetPaymentReturnAdmissionReader is a Reader for the GetPaymentReturnAdmission structure.
type GetPaymentReturnAdmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentReturnAdmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentReturnAdmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentReturnAdmissionOK creates a GetPaymentReturnAdmissionOK with default headers values
func NewGetPaymentReturnAdmissionOK() *GetPaymentReturnAdmissionOK {
	return &GetPaymentReturnAdmissionOK{}
}

/*GetPaymentReturnAdmissionOK handles this case with default header values.

Return admission details
*/
type GetPaymentReturnAdmissionOK struct {

	//Payload
	*models.ReturnAdmissionDetailsResponse
}

func (o *GetPaymentReturnAdmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/returns/{returnId}/admissions/{admissionId}][%d] getPaymentReturnAdmissionOK  %+v", 200, o)
}

func (o *GetPaymentReturnAdmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReturnAdmissionDetailsResponse = new(models.ReturnAdmissionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.ReturnAdmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
