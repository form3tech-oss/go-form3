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

// GetDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDReader is a Reader for the GetDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionID structure.
type GetDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDOK creates a GetDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDOK with default headers values
func NewGetDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDOK() *GetDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDOK {
	return &GetDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDOK{}
}

/*
GetDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDOK handles this case with default header values.

Direct debit decision details
*/
type GetDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDOK struct {

	//Payload

	// isStream: false
	*models.DirectDebitDecisionSubmissionDetailsResponse
}

func (o *GetDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDOK) Error() string {
	return fmt.Sprintf("[GET /transaction/directdebits/{id}/decisions/{decisionId}/submissions/{submissionId}][%d] getDirectdebitsIdDecisionsDecisionIdSubmissionsSubmissionIdOK", 200)
}

func (o *GetDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitDecisionSubmissionDetailsResponse = new(models.DirectDebitDecisionSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitDecisionSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
