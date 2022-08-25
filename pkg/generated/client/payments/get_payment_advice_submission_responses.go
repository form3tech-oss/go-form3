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

/*GetPaymentAdviceSubmissionOK handles this case with default header values.

Advice submission details
*/
type GetPaymentAdviceSubmissionOK struct {

	//Payload

	// isStream: false
	*models.AdviceSubmissionDetailsResponse
}

func (o *GetPaymentAdviceSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/advices/{adviceId}/submissions/{submissionId}][%d] getPaymentAdviceSubmissionOK", 200)
}

func (o *GetPaymentAdviceSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AdviceSubmissionDetailsResponse = new(models.AdviceSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AdviceSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
