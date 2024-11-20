// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// FetchPublicKeyForUserReader is a Reader for the FetchPublicKeyForUser structure.
type FetchPublicKeyForUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FetchPublicKeyForUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFetchPublicKeyForUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewFetchPublicKeyForUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFetchPublicKeyForUserOK creates a FetchPublicKeyForUserOK with default headers values
func NewFetchPublicKeyForUserOK() *FetchPublicKeyForUserOK {
	return &FetchPublicKeyForUserOK{}
}

/*
FetchPublicKeyForUserOK handles this case with default header values.

Public key
*/
type FetchPublicKeyForUserOK struct {

	//Payload

	// isStream: false
	*models.PublicKeyResponse
}

// IsSuccess returns true when this fetch public key for user o k response has a 2xx status code
func (o *FetchPublicKeyForUserOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fetch public key for user o k response has a 3xx status code
func (o *FetchPublicKeyForUserOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fetch public key for user o k response has a 4xx status code
func (o *FetchPublicKeyForUserOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fetch public key for user o k response has a 5xx status code
func (o *FetchPublicKeyForUserOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fetch public key for user o k response a status code equal to that given
func (o *FetchPublicKeyForUserOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the fetch public key for user o k response
func (o *FetchPublicKeyForUserOK) Code() int {
	return 200
}

func (o *FetchPublicKeyForUserOK) Error() string {
	return fmt.Sprintf("[GET /security/users/{user_id}/authn/public_keys/{public_key_id}][%d] fetchPublicKeyForUserOK", 200)
}

func (o *FetchPublicKeyForUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.PublicKeyResponse = new(models.PublicKeyResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.PublicKeyResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFetchPublicKeyForUserNotFound creates a FetchPublicKeyForUserNotFound with default headers values
func NewFetchPublicKeyForUserNotFound() *FetchPublicKeyForUserNotFound {
	return &FetchPublicKeyForUserNotFound{}
}

/*
FetchPublicKeyForUserNotFound handles this case with default header values.

Not Found
*/
type FetchPublicKeyForUserNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this fetch public key for user not found response has a 2xx status code
func (o *FetchPublicKeyForUserNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this fetch public key for user not found response has a 3xx status code
func (o *FetchPublicKeyForUserNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fetch public key for user not found response has a 4xx status code
func (o *FetchPublicKeyForUserNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this fetch public key for user not found response has a 5xx status code
func (o *FetchPublicKeyForUserNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this fetch public key for user not found response a status code equal to that given
func (o *FetchPublicKeyForUserNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the fetch public key for user not found response
func (o *FetchPublicKeyForUserNotFound) Code() int {
	return 404
}

func (o *FetchPublicKeyForUserNotFound) Error() string {
	return fmt.Sprintf("[GET /security/users/{user_id}/authn/public_keys/{public_key_id}][%d] fetchPublicKeyForUserNotFound", 404)
}

func (o *FetchPublicKeyForUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
