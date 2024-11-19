// Code generated by go-swagger; DO NOT EDIT.

package payment_reads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// GetReversalAdmissionTaskReader is a Reader for the GetReversalAdmissionTask structure.
type GetReversalAdmissionTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReversalAdmissionTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReversalAdmissionTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReversalAdmissionTaskOK creates a GetReversalAdmissionTaskOK with default headers values
func NewGetReversalAdmissionTaskOK() *GetReversalAdmissionTaskOK {
	return &GetReversalAdmissionTaskOK{}
}

/*
GetReversalAdmissionTaskOK handles this case with default header values.

Reversal Admission Task Details
*/
type GetReversalAdmissionTaskOK struct {

	//Payload

	// isStream: false
	*models.ReversalAdmissionTaskFetchResponse
}

func (o *GetReversalAdmissionTaskOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/reversals/{reversalId}/admissions/{admissionId}/tasks/{taskId}][%d] getReversalAdmissionTaskOK", 200)
}

func (o *GetReversalAdmissionTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReversalAdmissionTaskFetchResponse = new(models.ReversalAdmissionTaskFetchResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReversalAdmissionTaskFetchResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
