// Code generated by go-swagger; DO NOT EDIT.

package platformsecurityapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v5/pkg/generated/models"
)

// GetPlatformSecuritySigningKeysSigningkeyIDReader is a Reader for the GetPlatformSecuritySigningKeysSigningkeyID structure.
type GetPlatformSecuritySigningKeysSigningkeyIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlatformSecuritySigningKeysSigningkeyIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPlatformSecuritySigningKeysSigningkeyIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPlatformSecuritySigningKeysSigningkeyIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetPlatformSecuritySigningKeysSigningkeyIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetPlatformSecuritySigningKeysSigningkeyIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewGetPlatformSecuritySigningKeysSigningkeyIDBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetPlatformSecuritySigningKeysSigningkeyIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPlatformSecuritySigningKeysSigningkeyIDOK creates a GetPlatformSecuritySigningKeysSigningkeyIDOK with default headers values
func NewGetPlatformSecuritySigningKeysSigningkeyIDOK() *GetPlatformSecuritySigningKeysSigningkeyIDOK {
	return &GetPlatformSecuritySigningKeysSigningkeyIDOK{}
}

/*GetPlatformSecuritySigningKeysSigningkeyIDOK handles this case with default header values.

signing key response
*/
type GetPlatformSecuritySigningKeysSigningkeyIDOK struct {

	//Payload

	// isStream: false
	*models.SigningKeysResponse
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDOK) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys/{signingkey_id}][%d] getPlatformSecuritySigningKeysSigningkeyIdOK", 200)
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.SigningKeysResponse = new(models.SigningKeysResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.SigningKeysResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysSigningkeyIDBadRequest creates a GetPlatformSecuritySigningKeysSigningkeyIDBadRequest with default headers values
func NewGetPlatformSecuritySigningKeysSigningkeyIDBadRequest() *GetPlatformSecuritySigningKeysSigningkeyIDBadRequest {
	return &GetPlatformSecuritySigningKeysSigningkeyIDBadRequest{}
}

/*GetPlatformSecuritySigningKeysSigningkeyIDBadRequest handles this case with default header values.

Bad Request
*/
type GetPlatformSecuritySigningKeysSigningkeyIDBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys/{signingkey_id}][%d] getPlatformSecuritySigningKeysSigningkeyIdBadRequest", 400)
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysSigningkeyIDForbidden creates a GetPlatformSecuritySigningKeysSigningkeyIDForbidden with default headers values
func NewGetPlatformSecuritySigningKeysSigningkeyIDForbidden() *GetPlatformSecuritySigningKeysSigningkeyIDForbidden {
	return &GetPlatformSecuritySigningKeysSigningkeyIDForbidden{}
}

/*GetPlatformSecuritySigningKeysSigningkeyIDForbidden handles this case with default header values.

Action Forbidden
*/
type GetPlatformSecuritySigningKeysSigningkeyIDForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDForbidden) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys/{signingkey_id}][%d] getPlatformSecuritySigningKeysSigningkeyIdForbidden", 403)
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysSigningkeyIDNotFound creates a GetPlatformSecuritySigningKeysSigningkeyIDNotFound with default headers values
func NewGetPlatformSecuritySigningKeysSigningkeyIDNotFound() *GetPlatformSecuritySigningKeysSigningkeyIDNotFound {
	return &GetPlatformSecuritySigningKeysSigningkeyIDNotFound{}
}

/*GetPlatformSecuritySigningKeysSigningkeyIDNotFound handles this case with default header values.

Not Found
*/
type GetPlatformSecuritySigningKeysSigningkeyIDNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys/{signingkey_id}][%d] getPlatformSecuritySigningKeysSigningkeyIdNotFound", 404)
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysSigningkeyIDBadGateway creates a GetPlatformSecuritySigningKeysSigningkeyIDBadGateway with default headers values
func NewGetPlatformSecuritySigningKeysSigningkeyIDBadGateway() *GetPlatformSecuritySigningKeysSigningkeyIDBadGateway {
	return &GetPlatformSecuritySigningKeysSigningkeyIDBadGateway{}
}

/*GetPlatformSecuritySigningKeysSigningkeyIDBadGateway handles this case with default header values.

Bad Gateway
*/
type GetPlatformSecuritySigningKeysSigningkeyIDBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDBadGateway) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys/{signingkey_id}][%d] getPlatformSecuritySigningKeysSigningkeyIdBadGateway", 502)
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysSigningkeyIDDefault creates a GetPlatformSecuritySigningKeysSigningkeyIDDefault with default headers values
func NewGetPlatformSecuritySigningKeysSigningkeyIDDefault(code int) *GetPlatformSecuritySigningKeysSigningkeyIDDefault {
	return &GetPlatformSecuritySigningKeysSigningkeyIDDefault{
		_statusCode: code,
	}
}

/*GetPlatformSecuritySigningKeysSigningkeyIDDefault handles this case with default header values.

Unexpected Error
*/
type GetPlatformSecuritySigningKeysSigningkeyIDDefault struct {
	_statusCode int

	//Payload

	// isStream: false
	*models.APIError
}

// Code gets the status code for the get platform security signing keys signingkey ID default response
func (o *GetPlatformSecuritySigningKeysSigningkeyIDDefault) Code() int {
	return o._statusCode
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDDefault) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys/{signingkey_id}][%d] GetPlatformSecuritySigningKeysSigningkeyID default", o._statusCode)
}

func (o *GetPlatformSecuritySigningKeysSigningkeyIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
