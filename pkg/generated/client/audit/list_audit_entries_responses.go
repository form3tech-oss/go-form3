// Code generated by go-swagger; DO NOT EDIT.

package audit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v5/pkg/generated/models"
)

// ListAuditEntriesReader is a Reader for the ListAuditEntries structure.
type ListAuditEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAuditEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListAuditEntriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAuditEntriesOK creates a ListAuditEntriesOK with default headers values
func NewListAuditEntriesOK() *ListAuditEntriesOK {
	return &ListAuditEntriesOK{}
}

/*ListAuditEntriesOK handles this case with default header values.

Audit details
*/
type ListAuditEntriesOK struct {

	//Payload

	// isStream: false
	*models.AuditEntryListResponse
}

func (o *ListAuditEntriesOK) Error() string {
	return fmt.Sprintf("[GET /audit/entries/{record_type}][%d] listAuditEntriesOK", 200)
}

func (o *ListAuditEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AuditEntryListResponse = new(models.AuditEntryListResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.AuditEntryListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
