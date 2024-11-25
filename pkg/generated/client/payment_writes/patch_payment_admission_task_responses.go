// Code generated by go-swagger; DO NOT EDIT.

package payment_writes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// PatchPaymentAdmissionTaskReader is a Reader for the PatchPaymentAdmissionTask structure.
type PatchPaymentAdmissionTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchPaymentAdmissionTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchPaymentAdmissionTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchPaymentAdmissionTaskBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewPatchPaymentAdmissionTaskConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchPaymentAdmissionTaskOK creates a PatchPaymentAdmissionTaskOK with default headers values
func NewPatchPaymentAdmissionTaskOK() *PatchPaymentAdmissionTaskOK {
	return &PatchPaymentAdmissionTaskOK{}
}

/*
PatchPaymentAdmissionTaskOK handles this case with default header values.

Payment Admission Task update response
*/
type PatchPaymentAdmissionTaskOK struct {

	//Payload

	// isStream: false
	*models.PaymentAdmissionTaskDetailsResponse
}

// IsSuccess returns true when this patch payment admission task o k response has a 2xx status code
func (o *PatchPaymentAdmissionTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch payment admission task o k response has a 3xx status code
func (o *PatchPaymentAdmissionTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch payment admission task o k response has a 4xx status code
func (o *PatchPaymentAdmissionTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch payment admission task o k response has a 5xx status code
func (o *PatchPaymentAdmissionTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch payment admission task o k response a status code equal to that given
func (o *PatchPaymentAdmissionTaskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch payment admission task o k response
func (o *PatchPaymentAdmissionTaskOK) Code() int {
	return 200
}

func (o *PatchPaymentAdmissionTaskOK) Error() string {
	return fmt.Sprintf("[PATCH /transaction/payments/{id}/admissions/{admissionId}/tasks/{taskId}][%d] patchPaymentAdmissionTaskOK", 200)
}

func (o *PatchPaymentAdmissionTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.PaymentAdmissionTaskDetailsResponse = new(models.PaymentAdmissionTaskDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.PaymentAdmissionTaskDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPaymentAdmissionTaskBadRequest creates a PatchPaymentAdmissionTaskBadRequest with default headers values
func NewPatchPaymentAdmissionTaskBadRequest() *PatchPaymentAdmissionTaskBadRequest {
	return &PatchPaymentAdmissionTaskBadRequest{}
}

/*
PatchPaymentAdmissionTaskBadRequest handles this case with default header values.

Error
*/
type PatchPaymentAdmissionTaskBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

// IsSuccess returns true when this patch payment admission task bad request response has a 2xx status code
func (o *PatchPaymentAdmissionTaskBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch payment admission task bad request response has a 3xx status code
func (o *PatchPaymentAdmissionTaskBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch payment admission task bad request response has a 4xx status code
func (o *PatchPaymentAdmissionTaskBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch payment admission task bad request response has a 5xx status code
func (o *PatchPaymentAdmissionTaskBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch payment admission task bad request response a status code equal to that given
func (o *PatchPaymentAdmissionTaskBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the patch payment admission task bad request response
func (o *PatchPaymentAdmissionTaskBadRequest) Code() int {
	return 400
}

func (o *PatchPaymentAdmissionTaskBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /transaction/payments/{id}/admissions/{admissionId}/tasks/{taskId}][%d] patchPaymentAdmissionTaskBadRequest", 400)
}

func (o *PatchPaymentAdmissionTaskBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchPaymentAdmissionTaskConflict creates a PatchPaymentAdmissionTaskConflict with default headers values
func NewPatchPaymentAdmissionTaskConflict() *PatchPaymentAdmissionTaskConflict {
	return &PatchPaymentAdmissionTaskConflict{}
}

/*
PatchPaymentAdmissionTaskConflict handles this case with default header values.

Conflict
*/
type PatchPaymentAdmissionTaskConflict struct {

	//Payload

	// isStream: false
	*models.APIErrorWithActualResource
}

// IsSuccess returns true when this patch payment admission task conflict response has a 2xx status code
func (o *PatchPaymentAdmissionTaskConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch payment admission task conflict response has a 3xx status code
func (o *PatchPaymentAdmissionTaskConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch payment admission task conflict response has a 4xx status code
func (o *PatchPaymentAdmissionTaskConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch payment admission task conflict response has a 5xx status code
func (o *PatchPaymentAdmissionTaskConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this patch payment admission task conflict response a status code equal to that given
func (o *PatchPaymentAdmissionTaskConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the patch payment admission task conflict response
func (o *PatchPaymentAdmissionTaskConflict) Code() int {
	return 409
}

func (o *PatchPaymentAdmissionTaskConflict) Error() string {
	return fmt.Sprintf("[PATCH /transaction/payments/{id}/admissions/{admissionId}/tasks/{taskId}][%d] patchPaymentAdmissionTaskConflict", 409)
}

func (o *PatchPaymentAdmissionTaskConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIErrorWithActualResource = new(models.APIErrorWithActualResource)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIErrorWithActualResource); err != nil && err != io.EOF {
		return err
	}

	return nil
}
