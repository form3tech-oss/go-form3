// Code generated by go-swagger; DO NOT EDIT.

package reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// ListReportsReader is a Reader for the ListReports structure.
type ListReportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.

func (o *ListReportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:

		result := NewListReportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:

		result := NewListReportsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:

		result := NewListReportsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListReportsOK creates a ListReportsOK with default headers values
func NewListReportsOK() *ListReportsOK {
	return &ListReportsOK{}
}

/*ListReportsOK handles this case with default header values.

List of reports
*/
type ListReportsOK struct {

	//Payload

	// isStream: false
	*models.ReportDetailsListResponse
}

func (o *ListReportsOK) Error() string {
	return fmt.Sprintf("[GET /notification/reports][%d] listReportsOK  %+v", 200, o)
}

func (o *ListReportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.ReportDetailsListResponse = new(models.ReportDetailsListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.ReportDetailsListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReportsBadRequest creates a ListReportsBadRequest with default headers values
func NewListReportsBadRequest() *ListReportsBadRequest {
	return &ListReportsBadRequest{}
}

/*ListReportsBadRequest handles this case with default header values.

Reports bad request
*/
type ListReportsBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListReportsBadRequest) Error() string {
	return fmt.Sprintf("[GET /notification/reports][%d] listReportsBadRequest  %+v", 400, o)
}

func (o *ListReportsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReportsForbidden creates a ListReportsForbidden with default headers values
func NewListReportsForbidden() *ListReportsForbidden {
	return &ListReportsForbidden{}
}

/*ListReportsForbidden handles this case with default header values.

Forbidden
*/
type ListReportsForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *ListReportsForbidden) Error() string {
	return fmt.Sprintf("[GET /notification/reports][%d] listReportsForbidden  %+v", 403, o)
}

func (o *ListReportsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
