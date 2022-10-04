// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// GetReportAdmissionReader is a Reader for the GetReportAdmission structure.
type GetReportAdmissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportAdmissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReportAdmissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetReportAdmissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetReportAdmissionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetReportAdmissionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReportAdmissionOK creates a GetReportAdmissionOK with default headers values
func NewGetReportAdmissionOK() *GetReportAdmissionOK {
	return &GetReportAdmissionOK{}
}

/*GetReportAdmissionOK handles this case with default header values.

Get report admission by Id
*/
type GetReportAdmissionOK struct {

	//Payload

	// isStream: false
	*models.ReportAdmissionDetailsResponse
}

func (o *GetReportAdmissionOK) Error() string {
	return fmt.Sprintf("[GET /notification/reports/{id}/admissions/{admissionId}][%d] getReportAdmissionOK", 200)
}

func (o *GetReportAdmissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReportAdmissionDetailsResponse = new(models.ReportAdmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReportAdmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportAdmissionBadRequest creates a GetReportAdmissionBadRequest with default headers values
func NewGetReportAdmissionBadRequest() *GetReportAdmissionBadRequest {
	return &GetReportAdmissionBadRequest{}
}

/*GetReportAdmissionBadRequest handles this case with default header values.

Bad Request
*/
type GetReportAdmissionBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetReportAdmissionBadRequest) Error() string {
	return fmt.Sprintf("[GET /notification/reports/{id}/admissions/{admissionId}][%d] getReportAdmissionBadRequest", 400)
}

func (o *GetReportAdmissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportAdmissionForbidden creates a GetReportAdmissionForbidden with default headers values
func NewGetReportAdmissionForbidden() *GetReportAdmissionForbidden {
	return &GetReportAdmissionForbidden{}
}

/*GetReportAdmissionForbidden handles this case with default header values.

Forbidden
*/
type GetReportAdmissionForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetReportAdmissionForbidden) Error() string {
	return fmt.Sprintf("[GET /notification/reports/{id}/admissions/{admissionId}][%d] getReportAdmissionForbidden", 403)
}

func (o *GetReportAdmissionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportAdmissionNotFound creates a GetReportAdmissionNotFound with default headers values
func NewGetReportAdmissionNotFound() *GetReportAdmissionNotFound {
	return &GetReportAdmissionNotFound{}
}

/*GetReportAdmissionNotFound handles this case with default header values.

Not Found
*/
type GetReportAdmissionNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetReportAdmissionNotFound) Error() string {
	return fmt.Sprintf("[GET /notification/reports/{id}/admissions/{admissionId}][%d] getReportAdmissionNotFound", 404)
}

func (o *GetReportAdmissionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
