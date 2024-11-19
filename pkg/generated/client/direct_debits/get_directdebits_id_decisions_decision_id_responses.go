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

// GetDirectdebitsIDDecisionsDecisionIDReader is a Reader for the GetDirectdebitsIDDecisionsDecisionID structure.
type GetDirectdebitsIDDecisionsDecisionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDirectdebitsIDDecisionsDecisionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDirectdebitsIDDecisionsDecisionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDirectdebitsIDDecisionsDecisionIDOK creates a GetDirectdebitsIDDecisionsDecisionIDOK with default headers values
func NewGetDirectdebitsIDDecisionsDecisionIDOK() *GetDirectdebitsIDDecisionsDecisionIDOK {
	return &GetDirectdebitsIDDecisionsDecisionIDOK{}
}

/*
GetDirectdebitsIDDecisionsDecisionIDOK handles this case with default header values.

Direct debit decision details
*/
type GetDirectdebitsIDDecisionsDecisionIDOK struct {

	//Payload

	// isStream: false
	*models.DirectDebitDecisionDetailsResponse
}

func (o *GetDirectdebitsIDDecisionsDecisionIDOK) Error() string {
	return fmt.Sprintf("[GET /transaction/directdebits/{id}/decisions/{decisionId}][%d] getDirectdebitsIdDecisionsDecisionIdOK", 200)
}

func (o *GetDirectdebitsIDDecisionsDecisionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitDecisionDetailsResponse = new(models.DirectDebitDecisionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitDecisionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
