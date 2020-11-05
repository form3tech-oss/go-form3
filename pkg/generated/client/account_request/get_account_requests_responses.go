// Code generated by go-swagger; DO NOT EDIT.

package account_request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/generated/models"
)

// GetAccountRequestsReader is a Reader for the GetAccountRequests structure.
type GetAccountRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAccountRequestsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAccountRequestsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAccountRequestsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAccountRequestsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewGetAccountRequestsBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetAccountRequestsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetAccountRequestsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAccountRequestsOK creates a GetAccountRequestsOK with default headers values
func NewGetAccountRequestsOK() *GetAccountRequestsOK {
	return &GetAccountRequestsOK{}
}

/*GetAccountRequestsOK handles this case with default header values.

List Account Requests
*/
type GetAccountRequestsOK struct {

	//Payload

	// isStream: false
	*models.AccountRequestListResponse
}

func (o *GetAccountRequestsOK) Error() string {
	return fmt.Sprintf("[GET /organisation/accountrequests][%d] getAccountRequestsOK", 200)
}

func (o *GetAccountRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AccountRequestListResponse = new(models.AccountRequestListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AccountRequestListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountRequestsBadRequest creates a GetAccountRequestsBadRequest with default headers values
func NewGetAccountRequestsBadRequest() *GetAccountRequestsBadRequest {
	return &GetAccountRequestsBadRequest{}
}

/*GetAccountRequestsBadRequest handles this case with default header values.

Account Request creation error
*/
type GetAccountRequestsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountRequestsBadRequest) Error() string {
	return fmt.Sprintf("[GET /organisation/accountrequests][%d] getAccountRequestsBadRequest", 400)
}

func (o *GetAccountRequestsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountRequestsUnauthorized creates a GetAccountRequestsUnauthorized with default headers values
func NewGetAccountRequestsUnauthorized() *GetAccountRequestsUnauthorized {
	return &GetAccountRequestsUnauthorized{}
}

/*GetAccountRequestsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAccountRequestsUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountRequestsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organisation/accountrequests][%d] getAccountRequestsUnauthorized", 401)
}

func (o *GetAccountRequestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountRequestsForbidden creates a GetAccountRequestsForbidden with default headers values
func NewGetAccountRequestsForbidden() *GetAccountRequestsForbidden {
	return &GetAccountRequestsForbidden{}
}

/*GetAccountRequestsForbidden handles this case with default header values.

Forbidden
*/
type GetAccountRequestsForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountRequestsForbidden) Error() string {
	return fmt.Sprintf("[GET /organisation/accountrequests][%d] getAccountRequestsForbidden", 403)
}

func (o *GetAccountRequestsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountRequestsInternalServerError creates a GetAccountRequestsInternalServerError with default headers values
func NewGetAccountRequestsInternalServerError() *GetAccountRequestsInternalServerError {
	return &GetAccountRequestsInternalServerError{}
}

/*GetAccountRequestsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAccountRequestsInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountRequestsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organisation/accountrequests][%d] getAccountRequestsInternalServerError", 500)
}

func (o *GetAccountRequestsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountRequestsBadGateway creates a GetAccountRequestsBadGateway with default headers values
func NewGetAccountRequestsBadGateway() *GetAccountRequestsBadGateway {
	return &GetAccountRequestsBadGateway{}
}

/*GetAccountRequestsBadGateway handles this case with default header values.

Bad Gateway
*/
type GetAccountRequestsBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountRequestsBadGateway) Error() string {
	return fmt.Sprintf("[GET /organisation/accountrequests][%d] getAccountRequestsBadGateway", 502)
}

func (o *GetAccountRequestsBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountRequestsServiceUnavailable creates a GetAccountRequestsServiceUnavailable with default headers values
func NewGetAccountRequestsServiceUnavailable() *GetAccountRequestsServiceUnavailable {
	return &GetAccountRequestsServiceUnavailable{}
}

/*GetAccountRequestsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetAccountRequestsServiceUnavailable struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountRequestsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /organisation/accountrequests][%d] getAccountRequestsServiceUnavailable", 503)
}

func (o *GetAccountRequestsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountRequestsGatewayTimeout creates a GetAccountRequestsGatewayTimeout with default headers values
func NewGetAccountRequestsGatewayTimeout() *GetAccountRequestsGatewayTimeout {
	return &GetAccountRequestsGatewayTimeout{}
}

/*GetAccountRequestsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetAccountRequestsGatewayTimeout struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetAccountRequestsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /organisation/accountrequests][%d] getAccountRequestsGatewayTimeout", 504)
}

func (o *GetAccountRequestsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}