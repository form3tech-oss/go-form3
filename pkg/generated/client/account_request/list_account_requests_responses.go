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

// ListAccountRequestsReader is a Reader for the ListAccountRequests structure.
type ListAccountRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAccountRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAccountRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListAccountRequestsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListAccountRequestsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListAccountRequestsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListAccountRequestsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewListAccountRequestsBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewListAccountRequestsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewListAccountRequestsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAccountRequestsOK creates a ListAccountRequestsOK with default headers values
func NewListAccountRequestsOK() *ListAccountRequestsOK {
	return &ListAccountRequestsOK{}
}

/*ListAccountRequestsOK handles this case with default header values.

List Account Requests
*/
type ListAccountRequestsOK struct {

	//Payload

	// isStream: false
	*models.AccountRequestListResponse
}

func (o *ListAccountRequestsOK) Error() string {
	return fmt.Sprintf("[GET /organisation/accountrequests][%d] listAccountRequestsOK", 200)
}

func (o *ListAccountRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AccountRequestListResponse = new(models.AccountRequestListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AccountRequestListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAccountRequestsBadRequest creates a ListAccountRequestsBadRequest with default headers values
func NewListAccountRequestsBadRequest() *ListAccountRequestsBadRequest {
	return &ListAccountRequestsBadRequest{}
}

/*ListAccountRequestsBadRequest handles this case with default header values.

Account Request creation error
*/
type ListAccountRequestsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListAccountRequestsBadRequest) Error() string {
	return fmt.Sprintf("[GET /organisation/accountrequests][%d] listAccountRequestsBadRequest", 400)
}

func (o *ListAccountRequestsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAccountRequestsUnauthorized creates a ListAccountRequestsUnauthorized with default headers values
func NewListAccountRequestsUnauthorized() *ListAccountRequestsUnauthorized {
	return &ListAccountRequestsUnauthorized{}
}

/*ListAccountRequestsUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAccountRequestsUnauthorized struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListAccountRequestsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organisation/accountrequests][%d] listAccountRequestsUnauthorized", 401)
}

func (o *ListAccountRequestsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAccountRequestsForbidden creates a ListAccountRequestsForbidden with default headers values
func NewListAccountRequestsForbidden() *ListAccountRequestsForbidden {
	return &ListAccountRequestsForbidden{}
}

/*ListAccountRequestsForbidden handles this case with default header values.

Forbidden
*/
type ListAccountRequestsForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListAccountRequestsForbidden) Error() string {
	return fmt.Sprintf("[GET /organisation/accountrequests][%d] listAccountRequestsForbidden", 403)
}

func (o *ListAccountRequestsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAccountRequestsInternalServerError creates a ListAccountRequestsInternalServerError with default headers values
func NewListAccountRequestsInternalServerError() *ListAccountRequestsInternalServerError {
	return &ListAccountRequestsInternalServerError{}
}

/*ListAccountRequestsInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListAccountRequestsInternalServerError struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListAccountRequestsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organisation/accountrequests][%d] listAccountRequestsInternalServerError", 500)
}

func (o *ListAccountRequestsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAccountRequestsBadGateway creates a ListAccountRequestsBadGateway with default headers values
func NewListAccountRequestsBadGateway() *ListAccountRequestsBadGateway {
	return &ListAccountRequestsBadGateway{}
}

/*ListAccountRequestsBadGateway handles this case with default header values.

Bad Gateway
*/
type ListAccountRequestsBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListAccountRequestsBadGateway) Error() string {
	return fmt.Sprintf("[GET /organisation/accountrequests][%d] listAccountRequestsBadGateway", 502)
}

func (o *ListAccountRequestsBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAccountRequestsServiceUnavailable creates a ListAccountRequestsServiceUnavailable with default headers values
func NewListAccountRequestsServiceUnavailable() *ListAccountRequestsServiceUnavailable {
	return &ListAccountRequestsServiceUnavailable{}
}

/*ListAccountRequestsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type ListAccountRequestsServiceUnavailable struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListAccountRequestsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /organisation/accountrequests][%d] listAccountRequestsServiceUnavailable", 503)
}

func (o *ListAccountRequestsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAccountRequestsGatewayTimeout creates a ListAccountRequestsGatewayTimeout with default headers values
func NewListAccountRequestsGatewayTimeout() *ListAccountRequestsGatewayTimeout {
	return &ListAccountRequestsGatewayTimeout{}
}

/*ListAccountRequestsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type ListAccountRequestsGatewayTimeout struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListAccountRequestsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /organisation/accountrequests][%d] listAccountRequestsGatewayTimeout", 504)
}

func (o *ListAccountRequestsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
