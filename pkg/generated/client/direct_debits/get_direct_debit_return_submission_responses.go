// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v4/pkg/generated/models"
)

// GetDirectDebitReturnSubmissionReader is a Reader for the GetDirectDebitReturnSubmission structure.
type GetDirectDebitReturnSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDirectDebitReturnSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDirectDebitReturnSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDirectDebitReturnSubmissionOK creates a GetDirectDebitReturnSubmissionOK with default headers values
func NewGetDirectDebitReturnSubmissionOK() *GetDirectDebitReturnSubmissionOK {
	return &GetDirectDebitReturnSubmissionOK{}
}

/*GetDirectDebitReturnSubmissionOK handles this case with default header values.

Return submission details
*/
type GetDirectDebitReturnSubmissionOK struct {

	//Payload

	// isStream: false
	*models.DirectDebitReturnSubmissionDetailsResponse
}

func (o *GetDirectDebitReturnSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/directdebits/{id}/returns/{returnId}/submissions/{submissionId}][%d] getDirectDebitReturnSubmissionOK", 200)
}

func (o *GetDirectDebitReturnSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitReturnSubmissionDetailsResponse = new(models.DirectDebitReturnSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitReturnSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
