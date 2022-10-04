// Code generated by go-swagger; DO NOT EDIT.

package account_amendment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// GetAccountAmendmentSubmissionReader is a Reader for the GetAccountAmendmentSubmission structure.
type GetAccountAmendmentSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountAmendmentSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountAmendmentSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAccountAmendmentSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAccountAmendmentSubmissionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAccountAmendmentSubmissionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAccountAmendmentSubmissionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAccountAmendmentSubmissionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewGetAccountAmendmentSubmissionBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetAccountAmendmentSubmissionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetAccountAmendmentSubmissionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountAmendmentSubmissionOK creates a GetAccountAmendmentSubmissionOK with default headers values
func NewGetAccountAmendmentSubmissionOK() *GetAccountAmendmentSubmissionOK {
	return &GetAccountAmendmentSubmissionOK{}
}

/*GetAccountAmendmentSubmissionOK handles this case with default header values.

Account Amendment Submission details
*/
type GetAccountAmendmentSubmissionOK struct {

	//Payload

	// isStream: false
	*models.AccountAmendmentSubmissionResponse
}

func (o *GetAccountAmendmentSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}/submissions/{submissionId}][%d] getAccountAmendmentSubmissionOK", 200)
}

func (o *GetAccountAmendmentSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AccountAmendmentSubmissionResponse = new(models.AccountAmendmentSubmissionResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AccountAmendmentSubmissionResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAmendmentSubmissionBadRequest creates a GetAccountAmendmentSubmissionBadRequest with default headers values
func NewGetAccountAmendmentSubmissionBadRequest() *GetAccountAmendmentSubmissionBadRequest {
	return &GetAccountAmendmentSubmissionBadRequest{}
}

/*GetAccountAmendmentSubmissionBadRequest handles this case with default header values.

Account Amendment creation error
*/
type GetAccountAmendmentSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountAmendmentSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}/submissions/{submissionId}][%d] getAccountAmendmentSubmissionBadRequest", 400)
}

func (o *GetAccountAmendmentSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAmendmentSubmissionUnauthorized creates a GetAccountAmendmentSubmissionUnauthorized with default headers values
func NewGetAccountAmendmentSubmissionUnauthorized() *GetAccountAmendmentSubmissionUnauthorized {
	return &GetAccountAmendmentSubmissionUnauthorized{}
}

/*GetAccountAmendmentSubmissionUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAccountAmendmentSubmissionUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountAmendmentSubmissionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}/submissions/{submissionId}][%d] getAccountAmendmentSubmissionUnauthorized", 401)
}

func (o *GetAccountAmendmentSubmissionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAmendmentSubmissionForbidden creates a GetAccountAmendmentSubmissionForbidden with default headers values
func NewGetAccountAmendmentSubmissionForbidden() *GetAccountAmendmentSubmissionForbidden {
	return &GetAccountAmendmentSubmissionForbidden{}
}

/*GetAccountAmendmentSubmissionForbidden handles this case with default header values.

Forbidden
*/
type GetAccountAmendmentSubmissionForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountAmendmentSubmissionForbidden) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}/submissions/{submissionId}][%d] getAccountAmendmentSubmissionForbidden", 403)
}

func (o *GetAccountAmendmentSubmissionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAmendmentSubmissionNotFound creates a GetAccountAmendmentSubmissionNotFound with default headers values
func NewGetAccountAmendmentSubmissionNotFound() *GetAccountAmendmentSubmissionNotFound {
	return &GetAccountAmendmentSubmissionNotFound{}
}

/*GetAccountAmendmentSubmissionNotFound handles this case with default header values.

Not Found
*/
type GetAccountAmendmentSubmissionNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountAmendmentSubmissionNotFound) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}/submissions/{submissionId}][%d] getAccountAmendmentSubmissionNotFound", 404)
}

func (o *GetAccountAmendmentSubmissionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAmendmentSubmissionInternalServerError creates a GetAccountAmendmentSubmissionInternalServerError with default headers values
func NewGetAccountAmendmentSubmissionInternalServerError() *GetAccountAmendmentSubmissionInternalServerError {
	return &GetAccountAmendmentSubmissionInternalServerError{}
}

/*GetAccountAmendmentSubmissionInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAccountAmendmentSubmissionInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountAmendmentSubmissionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}/submissions/{submissionId}][%d] getAccountAmendmentSubmissionInternalServerError", 500)
}

func (o *GetAccountAmendmentSubmissionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAmendmentSubmissionBadGateway creates a GetAccountAmendmentSubmissionBadGateway with default headers values
func NewGetAccountAmendmentSubmissionBadGateway() *GetAccountAmendmentSubmissionBadGateway {
	return &GetAccountAmendmentSubmissionBadGateway{}
}

/*GetAccountAmendmentSubmissionBadGateway handles this case with default header values.

Bad Gateway
*/
type GetAccountAmendmentSubmissionBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountAmendmentSubmissionBadGateway) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}/submissions/{submissionId}][%d] getAccountAmendmentSubmissionBadGateway", 502)
}

func (o *GetAccountAmendmentSubmissionBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAmendmentSubmissionServiceUnavailable creates a GetAccountAmendmentSubmissionServiceUnavailable with default headers values
func NewGetAccountAmendmentSubmissionServiceUnavailable() *GetAccountAmendmentSubmissionServiceUnavailable {
	return &GetAccountAmendmentSubmissionServiceUnavailable{}
}

/*GetAccountAmendmentSubmissionServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetAccountAmendmentSubmissionServiceUnavailable struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountAmendmentSubmissionServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}/submissions/{submissionId}][%d] getAccountAmendmentSubmissionServiceUnavailable", 503)
}

func (o *GetAccountAmendmentSubmissionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAmendmentSubmissionGatewayTimeout creates a GetAccountAmendmentSubmissionGatewayTimeout with default headers values
func NewGetAccountAmendmentSubmissionGatewayTimeout() *GetAccountAmendmentSubmissionGatewayTimeout {
	return &GetAccountAmendmentSubmissionGatewayTimeout{}
}

/*GetAccountAmendmentSubmissionGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetAccountAmendmentSubmissionGatewayTimeout struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountAmendmentSubmissionGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}/submissions/{submissionId}][%d] getAccountAmendmentSubmissionGatewayTimeout", 504)
}

func (o *GetAccountAmendmentSubmissionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
