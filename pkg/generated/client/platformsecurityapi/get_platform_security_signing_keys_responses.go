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

// GetPlatformSecuritySigningKeysReader is a Reader for the GetPlatformSecuritySigningKeys structure.
type GetPlatformSecuritySigningKeysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlatformSecuritySigningKeysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPlatformSecuritySigningKeysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPlatformSecuritySigningKeysBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetPlatformSecuritySigningKeysForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewGetPlatformSecuritySigningKeysBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetPlatformSecuritySigningKeysDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPlatformSecuritySigningKeysOK creates a GetPlatformSecuritySigningKeysOK with default headers values
func NewGetPlatformSecuritySigningKeysOK() *GetPlatformSecuritySigningKeysOK {
	return &GetPlatformSecuritySigningKeysOK{}
}

/*GetPlatformSecuritySigningKeysOK handles this case with default header values.

keys returned
*/
type GetPlatformSecuritySigningKeysOK struct {

	//Payload

	// isStream: false
	*models.SigningKeysListResponse
}

func (o *GetPlatformSecuritySigningKeysOK) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys][%d] getPlatformSecuritySigningKeysOK", 200)
}

func (o *GetPlatformSecuritySigningKeysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.SigningKeysListResponse = new(models.SigningKeysListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.SigningKeysListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysBadRequest creates a GetPlatformSecuritySigningKeysBadRequest with default headers values
func NewGetPlatformSecuritySigningKeysBadRequest() *GetPlatformSecuritySigningKeysBadRequest {
	return &GetPlatformSecuritySigningKeysBadRequest{}
}

/*GetPlatformSecuritySigningKeysBadRequest handles this case with default header values.

Bad Request
*/
type GetPlatformSecuritySigningKeysBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetPlatformSecuritySigningKeysBadRequest) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys][%d] getPlatformSecuritySigningKeysBadRequest", 400)
}

func (o *GetPlatformSecuritySigningKeysBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysForbidden creates a GetPlatformSecuritySigningKeysForbidden with default headers values
func NewGetPlatformSecuritySigningKeysForbidden() *GetPlatformSecuritySigningKeysForbidden {
	return &GetPlatformSecuritySigningKeysForbidden{}
}

/*GetPlatformSecuritySigningKeysForbidden handles this case with default header values.

Action Forbidden
*/
type GetPlatformSecuritySigningKeysForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetPlatformSecuritySigningKeysForbidden) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys][%d] getPlatformSecuritySigningKeysForbidden", 403)
}

func (o *GetPlatformSecuritySigningKeysForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysBadGateway creates a GetPlatformSecuritySigningKeysBadGateway with default headers values
func NewGetPlatformSecuritySigningKeysBadGateway() *GetPlatformSecuritySigningKeysBadGateway {
	return &GetPlatformSecuritySigningKeysBadGateway{}
}

/*GetPlatformSecuritySigningKeysBadGateway handles this case with default header values.

Bad Gateway
*/
type GetPlatformSecuritySigningKeysBadGateway struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetPlatformSecuritySigningKeysBadGateway) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys][%d] getPlatformSecuritySigningKeysBadGateway", 502)
}

func (o *GetPlatformSecuritySigningKeysBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPlatformSecuritySigningKeysDefault creates a GetPlatformSecuritySigningKeysDefault with default headers values
func NewGetPlatformSecuritySigningKeysDefault(code int) *GetPlatformSecuritySigningKeysDefault {
	return &GetPlatformSecuritySigningKeysDefault{
		_statusCode: code,
	}
}

/*GetPlatformSecuritySigningKeysDefault handles this case with default header values.

Unexpected Error
*/
type GetPlatformSecuritySigningKeysDefault struct {
	_statusCode int

	//Payload

	// isStream: false
	*models.APIError
}

// Code gets the status code for the get platform security signing keys default response
func (o *GetPlatformSecuritySigningKeysDefault) Code() int {
	return o._statusCode
}

func (o *GetPlatformSecuritySigningKeysDefault) Error() string {
	return fmt.Sprintf("[GET /platform/security/signing_keys][%d] GetPlatformSecuritySigningKeys default", o._statusCode)
}

func (o *GetPlatformSecuritySigningKeysDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
