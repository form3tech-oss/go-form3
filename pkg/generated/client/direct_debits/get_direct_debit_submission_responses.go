// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v5/pkg/generated/models"
)

// GetDirectDebitSubmissionReader is a Reader for the GetDirectDebitSubmission structure.
type GetDirectDebitSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDirectDebitSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDirectDebitSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDirectDebitSubmissionOK creates a GetDirectDebitSubmissionOK with default headers values
func NewGetDirectDebitSubmissionOK() *GetDirectDebitSubmissionOK {
	return &GetDirectDebitSubmissionOK{}
}

/*GetDirectDebitSubmissionOK handles this case with default header values.

Direct debit submission details
*/
type GetDirectDebitSubmissionOK struct {

	//Payload

	// isStream: false
	*models.DirectDebitSubmissionDetailsResponse
}

func (o *GetDirectDebitSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/directdebits/{id}/submissions/{submissionId}][%d] getDirectDebitSubmissionOK", 200)
}

func (o *GetDirectDebitSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitSubmissionDetailsResponse = new(models.DirectDebitSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
