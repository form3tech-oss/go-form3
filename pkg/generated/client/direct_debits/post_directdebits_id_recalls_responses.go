// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// PostDirectdebitsIDRecallsReader is a Reader for the PostDirectdebitsIDRecalls structure.
type PostDirectdebitsIDRecallsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDirectdebitsIDRecallsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostDirectdebitsIDRecallsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostDirectdebitsIDRecallsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostDirectdebitsIDRecallsCreated creates a PostDirectdebitsIDRecallsCreated with default headers values
func NewPostDirectdebitsIDRecallsCreated() *PostDirectdebitsIDRecallsCreated {
	return &PostDirectdebitsIDRecallsCreated{}
}

/*PostDirectdebitsIDRecallsCreated handles this case with default header values.

Recall creation response
*/
type PostDirectdebitsIDRecallsCreated struct {

	//Payload

	// isStream: false
	*models.DirectDebitRecallCreationResponse
}

func (o *PostDirectdebitsIDRecallsCreated) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/recalls][%d] postDirectdebitsIdRecallsCreated", 201)
}

func (o *PostDirectdebitsIDRecallsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitRecallCreationResponse = new(models.DirectDebitRecallCreationResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitRecallCreationResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDirectdebitsIDRecallsBadRequest creates a PostDirectdebitsIDRecallsBadRequest with default headers values
func NewPostDirectdebitsIDRecallsBadRequest() *PostDirectdebitsIDRecallsBadRequest {
	return &PostDirectdebitsIDRecallsBadRequest{}
}

/*PostDirectdebitsIDRecallsBadRequest handles this case with default header values.

Recall creation error
*/
type PostDirectdebitsIDRecallsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *PostDirectdebitsIDRecallsBadRequest) Error() string {
	return fmt.Sprintf("[POST /transaction/directdebits/{id}/recalls][%d] postDirectdebitsIdRecallsBadRequest", 400)
}

func (o *PostDirectdebitsIDRecallsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
