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

/*GetPaymentReversalSubmissionOK handles this case with default header values.

Reversal submission details
*/
type GetPaymentReversalSubmissionOK struct {

	//Payload

	// isStream: false
	*models.ReversalSubmissionDetailsResponse
}

func (o *GetPaymentReversalSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/reversals/{reversalId}/submissions/{submissionId}][%d] getPaymentReversalSubmissionOK", 200)
}

func (o *GetPaymentReversalSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReversalSubmissionDetailsResponse = new(models.ReversalSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReversalSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
