// Code generated by go-swagger; DO NOT EDIT.

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/form3tech-oss/go-form3/pkg/generated/models"
)

// GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDAdmissionsAdmissionIDReader is a Reader for the GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDAdmissionsAdmissionID structure.
type GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDAdmissionsAdmissionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.

func (o *GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDAdmissionsAdmissionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:

		result := NewGetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDAdmissionsAdmissionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDAdmissionsAdmissionIDOK creates a GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDAdmissionsAdmissionIDOK with default headers values
func NewGetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDAdmissionsAdmissionIDOK() *GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDAdmissionsAdmissionIDOK {
	return &GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDAdmissionsAdmissionIDOK{}
}

/*GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDAdmissionsAdmissionIDOK handles this case with default header values.

Reversal admission details
*/
type GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDAdmissionsAdmissionIDOK struct {

	//Payload

	// isStream: false
	*models.DirectDebitReturnReversalAdmissionDetailsResponse
}

func (o *GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDAdmissionsAdmissionIDOK) Error() string {
	return fmt.Sprintf("[GET /transaction/directdebits/{id}/returns/{returnId}/reversals/{reversalId}/admissions/{admissionId}][%d] getTransactionDirectdebitsIdReturnsReturnIdReversalsReversalIdAdmissionsAdmissionIdOK  %+v", 200, o)
}

func (o *GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDAdmissionsAdmissionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitReturnReversalAdmissionDetailsResponse = new(models.DirectDebitReturnReversalAdmissionDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitReturnReversalAdmissionDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
