// Code generated by go-swagger; DO NOT EDIT.

package account_amendment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// GetAccountAmendmentReader is a Reader for the GetAccountAmendment structure.
type GetAccountAmendmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountAmendmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountAmendmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAccountAmendmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAccountAmendmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAccountAmendmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAccountAmendmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAccountAmendmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewGetAccountAmendmentBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetAccountAmendmentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetAccountAmendmentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountAmendmentOK creates a GetAccountAmendmentOK with default headers values
func NewGetAccountAmendmentOK() *GetAccountAmendmentOK {
	return &GetAccountAmendmentOK{}
}

/*GetAccountAmendmentOK handles this case with default header values.

Account Amendment Details
*/
type GetAccountAmendmentOK struct {

	//Payload

	// isStream: false
	*models.AccountAmendmentResponse
}

func (o *GetAccountAmendmentOK) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}][%d] getAccountAmendmentOK", 200)
}

func (o *GetAccountAmendmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AccountAmendmentResponse = new(models.AccountAmendmentResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AccountAmendmentResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAmendmentBadRequest creates a GetAccountAmendmentBadRequest with default headers values
func NewGetAccountAmendmentBadRequest() *GetAccountAmendmentBadRequest {
	return &GetAccountAmendmentBadRequest{}
}

/*GetAccountAmendmentBadRequest handles this case with default header values.

Account Request creation error
*/
type GetAccountAmendmentBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountAmendmentBadRequest) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}][%d] getAccountAmendmentBadRequest", 400)
}

func (o *GetAccountAmendmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAmendmentUnauthorized creates a GetAccountAmendmentUnauthorized with default headers values
func NewGetAccountAmendmentUnauthorized() *GetAccountAmendmentUnauthorized {
	return &GetAccountAmendmentUnauthorized{}
}

/*GetAccountAmendmentUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAccountAmendmentUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountAmendmentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}][%d] getAccountAmendmentUnauthorized", 401)
}

func (o *GetAccountAmendmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAmendmentForbidden creates a GetAccountAmendmentForbidden with default headers values
func NewGetAccountAmendmentForbidden() *GetAccountAmendmentForbidden {
	return &GetAccountAmendmentForbidden{}
}

/*GetAccountAmendmentForbidden handles this case with default header values.

Forbidden
*/
type GetAccountAmendmentForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountAmendmentForbidden) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}][%d] getAccountAmendmentForbidden", 403)
}

func (o *GetAccountAmendmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAmendmentNotFound creates a GetAccountAmendmentNotFound with default headers values
func NewGetAccountAmendmentNotFound() *GetAccountAmendmentNotFound {
	return &GetAccountAmendmentNotFound{}
}

/*GetAccountAmendmentNotFound handles this case with default header values.

Not Found
*/
type GetAccountAmendmentNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountAmendmentNotFound) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}][%d] getAccountAmendmentNotFound", 404)
}

func (o *GetAccountAmendmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAmendmentInternalServerError creates a GetAccountAmendmentInternalServerError with default headers values
func NewGetAccountAmendmentInternalServerError() *GetAccountAmendmentInternalServerError {
	return &GetAccountAmendmentInternalServerError{}
}

/*GetAccountAmendmentInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAccountAmendmentInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountAmendmentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}][%d] getAccountAmendmentInternalServerError", 500)
}

func (o *GetAccountAmendmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAmendmentBadGateway creates a GetAccountAmendmentBadGateway with default headers values
func NewGetAccountAmendmentBadGateway() *GetAccountAmendmentBadGateway {
	return &GetAccountAmendmentBadGateway{}
}

/*GetAccountAmendmentBadGateway handles this case with default header values.

Bad Gateway
*/
type GetAccountAmendmentBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountAmendmentBadGateway) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}][%d] getAccountAmendmentBadGateway", 502)
}

func (o *GetAccountAmendmentBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAmendmentServiceUnavailable creates a GetAccountAmendmentServiceUnavailable with default headers values
func NewGetAccountAmendmentServiceUnavailable() *GetAccountAmendmentServiceUnavailable {
	return &GetAccountAmendmentServiceUnavailable{}
}

/*GetAccountAmendmentServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetAccountAmendmentServiceUnavailable struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountAmendmentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}][%d] getAccountAmendmentServiceUnavailable", 503)
}

func (o *GetAccountAmendmentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAmendmentGatewayTimeout creates a GetAccountAmendmentGatewayTimeout with default headers values
func NewGetAccountAmendmentGatewayTimeout() *GetAccountAmendmentGatewayTimeout {
	return &GetAccountAmendmentGatewayTimeout{}
}

/*GetAccountAmendmentGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetAccountAmendmentGatewayTimeout struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountAmendmentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /organisation/accountamendments/{id}][%d] getAccountAmendmentGatewayTimeout", 504)
}

func (o *GetAccountAmendmentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}