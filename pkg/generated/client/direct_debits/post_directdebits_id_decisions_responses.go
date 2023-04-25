// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// PostDirectdebitsIDDecisionsReader is a Reader for the PostDirectdebitsIDDecisions structure.
type PostDirectdebitsIDDecisionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDirectdebitsIDDecisionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostDirectdebitsIDDecisionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostDirectdebitsIDDecisionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostDirectdebitsIDDecisionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDirectdebitsIDDecisionsCreated creates a PostDirectdebitsIDDecisionsCreated with default headers values
func NewPostDirectdebitsIDDecisionsCreated() *PostDirectdebitsIDDecisionsCreated {
	return &PostDirectdebitsIDDecisionsCreated{}
}

/*PostDirectdebitsIDDecisionsCreated handles this case with default header values.

Direct Debit decision creation response
*/
type PostDirectdebitsIDDecisionsCreated struct {

	//Payload

	// isStream: false
	*models.DirectDebitDecisionCreationResponse
}

func (o *PostDirectdebitsIDDecisionsCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/decisions][%d] postDirectdebitsIdDecisionsCreated", 201)
}

func (o *PostDirectdebitsIDDecisionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitDecisionCreationResponse = new(models.DirectDebitDecisionCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitDecisionCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDirectdebitsIDDecisionsBadRequest creates a PostDirectdebitsIDDecisionsBadRequest with default headers values
func NewPostDirectdebitsIDDecisionsBadRequest() *PostDirectdebitsIDDecisionsBadRequest {
	return &PostDirectdebitsIDDecisionsBadRequest{}
}

/*PostDirectdebitsIDDecisionsBadRequest handles this case with default header values.

Direct Debit decision creation error
*/
type PostDirectdebitsIDDecisionsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostDirectdebitsIDDecisionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/decisions][%d] postDirectdebitsIdDecisionsBadRequest", 400)
}

func (o *PostDirectdebitsIDDecisionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDirectdebitsIDDecisionsConflict creates a PostDirectdebitsIDDecisionsConflict with default headers values
func NewPostDirectdebitsIDDecisionsConflict() *PostDirectdebitsIDDecisionsConflict {
	return &PostDirectdebitsIDDecisionsConflict{}
}

/*PostDirectdebitsIDDecisionsConflict handles this case with default header values.

Direct Debit decision creation conflict error
*/
type PostDirectdebitsIDDecisionsConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostDirectdebitsIDDecisionsConflict) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/decisions][%d] postDirectdebitsIdDecisionsConflict", 409)
}

func (o *PostDirectdebitsIDDecisionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
