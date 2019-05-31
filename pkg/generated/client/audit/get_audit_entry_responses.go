// Code generated by go-swagger; DO NOT EDIT.

package audit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetAuditEntryReader is a Reader for the GetAuditEntry structure.
type GetAuditEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAuditEntryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAuditEntryOK creates a GetAuditEntryOK with default headers values
func NewGetAuditEntryOK() *GetAuditEntryOK {
	return &GetAuditEntryOK{}
}

/*GetAuditEntryOK handles this case with default header values.

Audit details
*/
type GetAuditEntryOK struct {

	//Payload
	*models.AuditEntryListResponse
}

func (o *GetAuditEntryOK) Error() string {
	return fmt.Sprintf("[GET /audit/entries/{record_type}/{id}][%d] getAuditEntryOK  %+v", 200, o)
}

func (o *GetAuditEntryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.AuditEntryListResponse = new(models.AuditEntryListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.AuditEntryListResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
