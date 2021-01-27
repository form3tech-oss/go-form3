// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsReader is a Reader for the PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissions structure.
type PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsCreated creates a PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsCreated with default headers values
func NewPostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsCreated() *PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsCreated {
	return &PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsCreated{}
}

/*PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsCreated handles this case with default header values.

Direct Debit decision admission creation response
*/
type PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsCreated struct {

	//Payload

	// isStream: false
	*models.DirectDebitDecisionAdmissionCreationResponse
}

func (o *PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/decisions/{decisionId}/admissions][%d] postTransactionDirectdebitsIdDecisionsDecisionIdAdmissionsCreated", 201)
}

func (o *PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitDecisionAdmissionCreationResponse = new(models.DirectDebitDecisionAdmissionCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitDecisionAdmissionCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsBadRequest creates a PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsBadRequest with default headers values
func NewPostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsBadRequest() *PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsBadRequest {
	return &PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsBadRequest{}
}

/*PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsBadRequest handles this case with default header values.

Direct Debit decision admission creation error
*/
type PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/decisions/{decisionId}/admissions][%d] postTransactionDirectdebitsIdDecisionsDecisionIdAdmissionsBadRequest", 400)
}

func (o *PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsConflict creates a PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsConflict with default headers values
func NewPostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsConflict() *PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsConflict {
	return &PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsConflict{}
}

/*PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsConflict handles this case with default header values.

Direct Debit decision submission creation conflict error
*/
type PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsConflict) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/decisions/{decisionId}/admissions][%d] postTransactionDirectdebitsIdDecisionsDecisionIdAdmissionsConflict", 409)
}

func (o *PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
