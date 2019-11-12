// Code generated by go-swagger; DO NOT EDIT.

package scheme_messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetSchemeMessageReader is a Reader for the GetSchemeMessage structure.
type GetSchemeMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchemeMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSchemeMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetSchemeMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSchemeMessageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSchemeMessageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSchemeMessageOK creates a GetSchemeMessageOK with default headers values
func NewGetSchemeMessageOK() *GetSchemeMessageOK {
	return &GetSchemeMessageOK{}
}

/*GetSchemeMessageOK handles this case with default header values.

Scheme Message details
*/
type GetSchemeMessageOK struct {

	//Payload

	// isStream: false
	*models.SchemeMessageDetailsResponse
}

func (o *GetSchemeMessageOK) Error() string {
	return fmt.Sprintf("[GET /notification/schememessages/{id}][%d] getSchemeMessageOK", 200)
}

func (o *GetSchemeMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.SchemeMessageDetailsResponse = new(models.SchemeMessageDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.SchemeMessageDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemeMessageBadRequest creates a GetSchemeMessageBadRequest with default headers values
func NewGetSchemeMessageBadRequest() *GetSchemeMessageBadRequest {
	return &GetSchemeMessageBadRequest{}
}

/*GetSchemeMessageBadRequest handles this case with default header values.

Scheme Message bad request
*/
type GetSchemeMessageBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetSchemeMessageBadRequest) Error() string {
	return fmt.Sprintf("[GET /notification/schememessages/{id}][%d] getSchemeMessageBadRequest", 400)
}

func (o *GetSchemeMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemeMessageForbidden creates a GetSchemeMessageForbidden with default headers values
func NewGetSchemeMessageForbidden() *GetSchemeMessageForbidden {
	return &GetSchemeMessageForbidden{}
}

/*GetSchemeMessageForbidden handles this case with default header values.

Forbidden
*/
type GetSchemeMessageForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetSchemeMessageForbidden) Error() string {
	return fmt.Sprintf("[GET /notification/schememessages/{id}][%d] getSchemeMessageForbidden", 403)
}

func (o *GetSchemeMessageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchemeMessageNotFound creates a GetSchemeMessageNotFound with default headers values
func NewGetSchemeMessageNotFound() *GetSchemeMessageNotFound {
	return &GetSchemeMessageNotFound{}
}

/*GetSchemeMessageNotFound handles this case with default header values.

Scheme Message Not found
*/
type GetSchemeMessageNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetSchemeMessageNotFound) Error() string {
	return fmt.Sprintf("[GET /notification/schememessages/{id}][%d] getSchemeMessageNotFound", 404)
}

func (o *GetSchemeMessageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
