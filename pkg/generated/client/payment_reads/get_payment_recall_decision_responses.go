// Code generated by go-swagger; DO NOT EDIT.

package payment_reads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// GetPaymentRecallDecisionReader is a Reader for the GetPaymentRecallDecision structure.
type GetPaymentRecallDecisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentRecallDecisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentRecallDecisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentRecallDecisionOK creates a GetPaymentRecallDecisionOK with default headers values
func NewGetPaymentRecallDecisionOK() *GetPaymentRecallDecisionOK {
	return &GetPaymentRecallDecisionOK{}
}

/*
GetPaymentRecallDecisionOK handles this case with default header values.

Recall Decision Details
*/
type GetPaymentRecallDecisionOK struct {

	//Payload

	// isStream: false
	*models.RecallDecisionFetchResponse
}

// IsSuccess returns true when this get payment recall decision o k response has a 2xx status code
func (o *GetPaymentRecallDecisionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get payment recall decision o k response has a 3xx status code
func (o *GetPaymentRecallDecisionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payment recall decision o k response has a 4xx status code
func (o *GetPaymentRecallDecisionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get payment recall decision o k response has a 5xx status code
func (o *GetPaymentRecallDecisionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get payment recall decision o k response a status code equal to that given
func (o *GetPaymentRecallDecisionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get payment recall decision o k response
func (o *GetPaymentRecallDecisionOK) Code() int {
	return 200
}

func (o *GetPaymentRecallDecisionOK) Error() string {
	return fmt.Sprintf("[GET /transaction/payments/{id}/recalls/{recallId}/decisions/{decisionId}][%d] getPaymentRecallDecisionOK", 200)
}

func (o *GetPaymentRecallDecisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.RecallDecisionFetchResponse = new(models.RecallDecisionFetchResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.RecallDecisionFetchResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
