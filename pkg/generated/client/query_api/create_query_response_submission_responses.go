// Code generated by go-swagger; DO NOT EDIT.

package query_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// CreateQueryResponseSubmissionReader is a Reader for the CreateQueryResponseSubmission structure.
type CreateQueryResponseSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateQueryResponseSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateQueryResponseSubmissionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateQueryResponseSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateQueryResponseSubmissionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewCreateQueryResponseSubmissionBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateQueryResponseSubmissionCreated creates a CreateQueryResponseSubmissionCreated with default headers values
func NewCreateQueryResponseSubmissionCreated() *CreateQueryResponseSubmissionCreated {
	return &CreateQueryResponseSubmissionCreated{}
}

/*
CreateQueryResponseSubmissionCreated handles this case with default header values.

creation response
*/
type CreateQueryResponseSubmissionCreated struct {

	//Payload

	// isStream: false
	*models.QueryResponseSubmissionResponse
}

// IsSuccess returns true when this create query response submission created response has a 2xx status code
func (o *CreateQueryResponseSubmissionCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create query response submission created response has a 3xx status code
func (o *CreateQueryResponseSubmissionCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create query response submission created response has a 4xx status code
func (o *CreateQueryResponseSubmissionCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create query response submission created response has a 5xx status code
func (o *CreateQueryResponseSubmissionCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create query response submission created response a status code equal to that given
func (o *CreateQueryResponseSubmissionCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create query response submission created response
func (o *CreateQueryResponseSubmissionCreated) Code() int {
	return 201
}

func (o *CreateQueryResponseSubmissionCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/queries/{query_id}/responses/{query_response_id}/submissions][%d] createQueryResponseSubmissionCreated", 201)
}

func (o *CreateQueryResponseSubmissionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.QueryResponseSubmissionResponse = new(models.QueryResponseSubmissionResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.QueryResponseSubmissionResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryResponseSubmissionBadRequest creates a CreateQueryResponseSubmissionBadRequest with default headers values
func NewCreateQueryResponseSubmissionBadRequest() *CreateQueryResponseSubmissionBadRequest {
	return &CreateQueryResponseSubmissionBadRequest{}
}

/*
CreateQueryResponseSubmissionBadRequest handles this case with default header values.

bad request
*/
type CreateQueryResponseSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create query response submission bad request response has a 2xx status code
func (o *CreateQueryResponseSubmissionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create query response submission bad request response has a 3xx status code
func (o *CreateQueryResponseSubmissionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create query response submission bad request response has a 4xx status code
func (o *CreateQueryResponseSubmissionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create query response submission bad request response has a 5xx status code
func (o *CreateQueryResponseSubmissionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create query response submission bad request response a status code equal to that given
func (o *CreateQueryResponseSubmissionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create query response submission bad request response
func (o *CreateQueryResponseSubmissionBadRequest) Code() int {
	return 400
}

func (o *CreateQueryResponseSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/queries/{query_id}/responses/{query_response_id}/submissions][%d] createQueryResponseSubmissionBadRequest", 400)
}

func (o *CreateQueryResponseSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryResponseSubmissionForbidden creates a CreateQueryResponseSubmissionForbidden with default headers values
func NewCreateQueryResponseSubmissionForbidden() *CreateQueryResponseSubmissionForbidden {
	return &CreateQueryResponseSubmissionForbidden{}
}

/*
CreateQueryResponseSubmissionForbidden handles this case with default header values.

forbidden
*/
type CreateQueryResponseSubmissionForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create query response submission forbidden response has a 2xx status code
func (o *CreateQueryResponseSubmissionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create query response submission forbidden response has a 3xx status code
func (o *CreateQueryResponseSubmissionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create query response submission forbidden response has a 4xx status code
func (o *CreateQueryResponseSubmissionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create query response submission forbidden response has a 5xx status code
func (o *CreateQueryResponseSubmissionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create query response submission forbidden response a status code equal to that given
func (o *CreateQueryResponseSubmissionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create query response submission forbidden response
func (o *CreateQueryResponseSubmissionForbidden) Code() int {
	return 403
}

func (o *CreateQueryResponseSubmissionForbidden) Error() string {
	return fmt.Sprintf("[POST /transaction/queries/{query_id}/responses/{query_response_id}/submissions][%d] createQueryResponseSubmissionForbidden", 403)
}

func (o *CreateQueryResponseSubmissionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateQueryResponseSubmissionBadGateway creates a CreateQueryResponseSubmissionBadGateway with default headers values
func NewCreateQueryResponseSubmissionBadGateway() *CreateQueryResponseSubmissionBadGateway {
	return &CreateQueryResponseSubmissionBadGateway{}
}

/*
CreateQueryResponseSubmissionBadGateway handles this case with default header values.

gateway issue
*/
type CreateQueryResponseSubmissionBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this create query response submission bad gateway response has a 2xx status code
func (o *CreateQueryResponseSubmissionBadGateway) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create query response submission bad gateway response has a 3xx status code
func (o *CreateQueryResponseSubmissionBadGateway) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create query response submission bad gateway response has a 4xx status code
func (o *CreateQueryResponseSubmissionBadGateway) IsClientError() bool {
	return false
}

// IsServerError returns true when this create query response submission bad gateway response has a 5xx status code
func (o *CreateQueryResponseSubmissionBadGateway) IsServerError() bool {
	return true
}

// IsCode returns true when this create query response submission bad gateway response a status code equal to that given
func (o *CreateQueryResponseSubmissionBadGateway) IsCode(code int) bool {
	return code == 502
}

// Code gets the status code for the create query response submission bad gateway response
func (o *CreateQueryResponseSubmissionBadGateway) Code() int {
	return 502
}

func (o *CreateQueryResponseSubmissionBadGateway) Error() string {
	return fmt.Sprintf("[POST /transaction/queries/{query_id}/responses/{query_response_id}/submissions][%d] createQueryResponseSubmissionBadGateway", 502)
}

func (o *CreateQueryResponseSubmissionBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
