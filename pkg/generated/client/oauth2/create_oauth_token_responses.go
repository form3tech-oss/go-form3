// Code generated by go-swagger; DO NOT EDIT.

package oauth2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v5/pkg/generated/models"
)

// CreateOauthTokenReader is a Reader for the CreateOauthToken structure.
type CreateOauthTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOauthTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateOauthTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewCreateOauthTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateOauthTokenOK creates a CreateOauthTokenOK with default headers values
func NewCreateOauthTokenOK() *CreateOauthTokenOK {
	return &CreateOauthTokenOK{}
}

/*CreateOauthTokenOK handles this case with default header values.

Authorisation token (Bearer)
*/
type CreateOauthTokenOK struct {

	//Payload

	// isStream: false
	*models.Token
}

func (o *CreateOauthTokenOK) Error() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] createOauthTokenOK", 200)
}

func (o *CreateOauthTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Token = new(models.Token)

	// response payload

	if err := consumer.Consume(response.Body(), o.Token); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOauthTokenForbidden creates a CreateOauthTokenForbidden with default headers values
func NewCreateOauthTokenForbidden() *CreateOauthTokenForbidden {
	return &CreateOauthTokenForbidden{}
}

/*CreateOauthTokenForbidden handles this case with default header values.

Authentication failed
*/
type CreateOauthTokenForbidden struct {
}

func (o *CreateOauthTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /oauth2/token][%d] createOauthTokenForbidden", 403)
}

func (o *CreateOauthTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
