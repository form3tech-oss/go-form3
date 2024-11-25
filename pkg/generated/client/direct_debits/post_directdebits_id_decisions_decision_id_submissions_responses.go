// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// PostDirectdebitsIDDecisionsDecisionIDSubmissionsReader is a Reader for the PostDirectdebitsIDDecisionsDecisionIDSubmissions structure.
type PostDirectdebitsIDDecisionsDecisionIDSubmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostDirectdebitsIDDecisionsDecisionIDSubmissionsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostDirectdebitsIDDecisionsDecisionIDSubmissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPostDirectdebitsIDDecisionsDecisionIDSubmissionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDirectdebitsIDDecisionsDecisionIDSubmissionsCreated creates a PostDirectdebitsIDDecisionsDecisionIDSubmissionsCreated with default headers values
func NewPostDirectdebitsIDDecisionsDecisionIDSubmissionsCreated() *PostDirectdebitsIDDecisionsDecisionIDSubmissionsCreated {
	return &PostDirectdebitsIDDecisionsDecisionIDSubmissionsCreated{}
}

/*
PostDirectdebitsIDDecisionsDecisionIDSubmissionsCreated handles this case with default header values.

Direct Debit decision submission creation response
*/
type PostDirectdebitsIDDecisionsDecisionIDSubmissionsCreated struct {

	//Payload

	// isStream: false
	*models.DirectDebitDecisionSubmissionCreationResponse
}

// IsSuccess returns true when this post directdebits Id decisions decision Id submissions created response has a 2xx status code
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post directdebits Id decisions decision Id submissions created response has a 3xx status code
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post directdebits Id decisions decision Id submissions created response has a 4xx status code
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post directdebits Id decisions decision Id submissions created response has a 5xx status code
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post directdebits Id decisions decision Id submissions created response a status code equal to that given
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the post directdebits Id decisions decision Id submissions created response
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsCreated) Code() int {
	return 201
}

func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/decisions/{decisionId}/submissions][%d] postDirectdebitsIdDecisionsDecisionIdSubmissionsCreated", 201)
}

func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitDecisionSubmissionCreationResponse = new(models.DirectDebitDecisionSubmissionCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitDecisionSubmissionCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDirectdebitsIDDecisionsDecisionIDSubmissionsBadRequest creates a PostDirectdebitsIDDecisionsDecisionIDSubmissionsBadRequest with default headers values
func NewPostDirectdebitsIDDecisionsDecisionIDSubmissionsBadRequest() *PostDirectdebitsIDDecisionsDecisionIDSubmissionsBadRequest {
	return &PostDirectdebitsIDDecisionsDecisionIDSubmissionsBadRequest{}
}

/*
PostDirectdebitsIDDecisionsDecisionIDSubmissionsBadRequest handles this case with default header values.

Direct Debit decision submission creation error
*/
type PostDirectdebitsIDDecisionsDecisionIDSubmissionsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this post directdebits Id decisions decision Id submissions bad request response has a 2xx status code
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post directdebits Id decisions decision Id submissions bad request response has a 3xx status code
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post directdebits Id decisions decision Id submissions bad request response has a 4xx status code
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post directdebits Id decisions decision Id submissions bad request response has a 5xx status code
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post directdebits Id decisions decision Id submissions bad request response a status code equal to that given
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post directdebits Id decisions decision Id submissions bad request response
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsBadRequest) Code() int {
	return 400
}

func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/decisions/{decisionId}/submissions][%d] postDirectdebitsIdDecisionsDecisionIdSubmissionsBadRequest", 400)
}

func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDirectdebitsIDDecisionsDecisionIDSubmissionsConflict creates a PostDirectdebitsIDDecisionsDecisionIDSubmissionsConflict with default headers values
func NewPostDirectdebitsIDDecisionsDecisionIDSubmissionsConflict() *PostDirectdebitsIDDecisionsDecisionIDSubmissionsConflict {
	return &PostDirectdebitsIDDecisionsDecisionIDSubmissionsConflict{}
}

/*
PostDirectdebitsIDDecisionsDecisionIDSubmissionsConflict handles this case with default header values.

Direct Debit decision submission creation conflict error
*/
type PostDirectdebitsIDDecisionsDecisionIDSubmissionsConflict struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this post directdebits Id decisions decision Id submissions conflict response has a 2xx status code
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post directdebits Id decisions decision Id submissions conflict response has a 3xx status code
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post directdebits Id decisions decision Id submissions conflict response has a 4xx status code
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this post directdebits Id decisions decision Id submissions conflict response has a 5xx status code
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this post directdebits Id decisions decision Id submissions conflict response a status code equal to that given
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the post directdebits Id decisions decision Id submissions conflict response
func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsConflict) Code() int {
	return 409
}

func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsConflict) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/decisions/{decisionId}/submissions][%d] postDirectdebitsIdDecisionsDecisionIdSubmissionsConflict", 409)
}

func (o *PostDirectdebitsIDDecisionsDecisionIDSubmissionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
