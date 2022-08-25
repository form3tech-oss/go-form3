// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v5/pkg/generated/models"
)

// GetPaymentReturnSubmissionReader is a Reader for the GetPaymentReturnSubmission structure.
type GetPaymentReturnSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentReturnSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentReturnSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentReturnSubmissionOK creates a GetPaymentReturnSubmissionOK with default headers values
func NewGetPaymentReturnSubmissionOK() *GetPaymentReturnSubmissionOK {
	return &GetPaymentReturnSubmissionOK{}
}

/*GetPaymentReturnSubmissionOK handles this case with default header values.

Return submission details
*/
type GetPaymentReturnSubmissionOK struct {

	//Payload

	// isStream: false
	*models.ReturnSubmissionDetailsResponse
}

func (o *GetPaymentReturnSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/returns/{returnId}/submissions/{submissionId}][%d] getPaymentReturnSubmissionOK", 200)
}

func (o *GetPaymentReturnSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReturnSubmissionDetailsResponse = new(models.ReturnSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReturnSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
