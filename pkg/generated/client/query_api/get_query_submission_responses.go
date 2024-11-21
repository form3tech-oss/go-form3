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

// GetQuerySubmissionReader is a Reader for the GetQuerySubmission structure.
type GetQuerySubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQuerySubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetQuerySubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetQuerySubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetQuerySubmissionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetQuerySubmissionOK creates a GetQuerySubmissionOK with default headers values
func NewGetQuerySubmissionOK() *GetQuerySubmissionOK {
	return &GetQuerySubmissionOK{}
}

/*
GetQuerySubmissionOK handles this case with default header values.

query submission response
*/
type GetQuerySubmissionOK struct {

	//Payload

	// isStream: false
	*models.QuerySubmissionResponse
}

// IsSuccess returns true when this get query submission o k response has a 2xx status code
func (o *GetQuerySubmissionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get query submission o k response has a 3xx status code
func (o *GetQuerySubmissionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get query submission o k response has a 4xx status code
func (o *GetQuerySubmissionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get query submission o k response has a 5xx status code
func (o *GetQuerySubmissionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get query submission o k response a status code equal to that given
func (o *GetQuerySubmissionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get query submission o k response
func (o *GetQuerySubmissionOK) Code() int {
	return 200
}

func (o *GetQuerySubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/queries/{query_id}/submissions/{query_submission_id}][%d] getQuerySubmissionOK", 200)
}

func (o *GetQuerySubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.QuerySubmissionResponse = new(models.QuerySubmissionResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.QuerySubmissionResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQuerySubmissionBadRequest creates a GetQuerySubmissionBadRequest with default headers values
func NewGetQuerySubmissionBadRequest() *GetQuerySubmissionBadRequest {
	return &GetQuerySubmissionBadRequest{}
}

/*
GetQuerySubmissionBadRequest handles this case with default header values.

bad request
*/
type GetQuerySubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this get query submission bad request response has a 2xx status code
func (o *GetQuerySubmissionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get query submission bad request response has a 3xx status code
func (o *GetQuerySubmissionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get query submission bad request response has a 4xx status code
func (o *GetQuerySubmissionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get query submission bad request response has a 5xx status code
func (o *GetQuerySubmissionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get query submission bad request response a status code equal to that given
func (o *GetQuerySubmissionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get query submission bad request response
func (o *GetQuerySubmissionBadRequest) Code() int {
	return 400
}

func (o *GetQuerySubmissionBadRequest) Error() string {
	return fmt.Sprintf("[GET /transaction/queries/{query_id}/submissions/{query_submission_id}][%d] getQuerySubmissionBadRequest", 400)
}

func (o *GetQuerySubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQuerySubmissionForbidden creates a GetQuerySubmissionForbidden with default headers values
func NewGetQuerySubmissionForbidden() *GetQuerySubmissionForbidden {
	return &GetQuerySubmissionForbidden{}
}

/*
GetQuerySubmissionForbidden handles this case with default header values.

forbidden
*/
type GetQuerySubmissionForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this get query submission forbidden response has a 2xx status code
func (o *GetQuerySubmissionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get query submission forbidden response has a 3xx status code
func (o *GetQuerySubmissionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get query submission forbidden response has a 4xx status code
func (o *GetQuerySubmissionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get query submission forbidden response has a 5xx status code
func (o *GetQuerySubmissionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get query submission forbidden response a status code equal to that given
func (o *GetQuerySubmissionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get query submission forbidden response
func (o *GetQuerySubmissionForbidden) Code() int {
	return 403
}

func (o *GetQuerySubmissionForbidden) Error() string {
	return fmt.Sprintf("[GET /transaction/queries/{query_id}/submissions/{query_submission_id}][%d] getQuerySubmissionForbidden", 403)
}

func (o *GetQuerySubmissionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}