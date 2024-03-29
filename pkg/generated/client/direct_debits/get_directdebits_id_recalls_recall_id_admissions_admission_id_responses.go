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

// GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDReader is a Reader for the GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionID structure.
type GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDOK creates a GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDOK with default headers values
func NewGetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDOK() *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDOK {
	return &GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDOK{}
}

/*GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDOK handles this case with default header values.

Recall admission details
*/
type GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDOK struct {

	//Payload

	// isStream: false
	*models.DirectDebitRecallAdmissionDetailsResponse
}

func (o *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDOK) Error() string {
	return fmt.Sprintf("[GET /transaction/directdebits/{id}/recalls/{recallId}/admissions/{admissionId}][%d] getDirectdebitsIdRecallsRecallIdAdmissionsAdmissionIdOK", 200)
}

func (o *GetDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitRecallAdmissionDetailsResponse = new(models.DirectDebitRecallAdmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitRecallAdmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
