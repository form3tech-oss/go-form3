// Code generated by go-swagger; DO NOT EDIT.

package mandates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v4/pkg/generated/models"
)

// GetMandateSubmissionReader is a Reader for the GetMandateSubmission structure.
type GetMandateSubmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMandateSubmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMandateSubmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetMandateSubmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMandateSubmissionOK creates a GetMandateSubmissionOK with default headers values
func NewGetMandateSubmissionOK() *GetMandateSubmissionOK {
	return &GetMandateSubmissionOK{}
}

/*GetMandateSubmissionOK handles this case with default header values.

Mandate Submission details
*/
type GetMandateSubmissionOK struct {

	//Payload

	// isStream: false
	*models.MandateSubmissionDetailsResponse
}

func (o *GetMandateSubmissionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/mandates/{id}/submissions/{submissionId}][%d] getMandateSubmissionOK", 200)
}

func (o *GetMandateSubmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.MandateSubmissionDetailsResponse = new(models.MandateSubmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.MandateSubmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMandateSubmissionBadRequest creates a GetMandateSubmissionBadRequest with default headers values
func NewGetMandateSubmissionBadRequest() *GetMandateSubmissionBadRequest {
	return &GetMandateSubmissionBadRequest{}
}

/*GetMandateSubmissionBadRequest handles this case with default header values.

Error
*/
type GetMandateSubmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetMandateSubmissionBadRequest) Error() string {
	return fmt.Sprintf("[GET /transaction/mandates/{id}/submissions/{submissionId}][%d] getMandateSubmissionBadRequest", 400)
}

func (o *GetMandateSubmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
