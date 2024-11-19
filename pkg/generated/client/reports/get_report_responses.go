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

// GetReportReader is a Reader for the GetReport structure.
type GetReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetReportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 406:
		result := NewGetReportNotAcceptable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReportOK creates a GetReportOK with default headers values
func NewGetReportOK() *GetReportOK {
	return &GetReportOK{}
}

/*
GetReportOK handles this case with default header values.

Report details
*/
type GetReportOK struct {

	//Payload

	// isStream: false
	*models.ReportDetailsResponse
}

func (o *GetReportOK) Error() string {
	return fmt.Sprintf("[GET /notification/reports/{id}][%d] getReportOK", 200)
}

func (o *GetReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReportDetailsResponse = new(models.ReportDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReportDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportBadRequest creates a GetReportBadRequest with default headers values
func NewGetReportBadRequest() *GetReportBadRequest {
	return &GetReportBadRequest{}
}

/*
GetReportBadRequest handles this case with default header values.

Bad Request
*/
type GetReportBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetReportBadRequest) Error() string {
	return fmt.Sprintf("[GET /notification/reports/{id}][%d] getReportBadRequest", 400)
}

func (o *GetReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportForbidden creates a GetReportForbidden with default headers values
func NewGetReportForbidden() *GetReportForbidden {
	return &GetReportForbidden{}
}

/*
GetReportForbidden handles this case with default header values.

Forbidden
*/
type GetReportForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetReportForbidden) Error() string {
	return fmt.Sprintf("[GET /notification/reports/{id}][%d] getReportForbidden", 403)
}

func (o *GetReportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportNotFound creates a GetReportNotFound with default headers values
func NewGetReportNotFound() *GetReportNotFound {
	return &GetReportNotFound{}
}

/*
GetReportNotFound handles this case with default header values.

Report Not Found
*/
type GetReportNotFound struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetReportNotFound) Error() string {
	return fmt.Sprintf("[GET /notification/reports/{id}][%d] getReportNotFound", 404)
}

func (o *GetReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportNotAcceptable creates a GetReportNotAcceptable with default headers values
func NewGetReportNotAcceptable() *GetReportNotAcceptable {
	return &GetReportNotAcceptable{}
}

/*
GetReportNotAcceptable handles this case with default header values.

Report not available in acceptable format
*/
type GetReportNotAcceptable struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetReportNotAcceptable) Error() string {
	return fmt.Sprintf("[GET /notification/reports/{id}][%d] getReportNotAcceptable", 406)
}

func (o *GetReportNotAcceptable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
