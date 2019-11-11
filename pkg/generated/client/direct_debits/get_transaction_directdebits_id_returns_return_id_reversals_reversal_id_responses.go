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

// GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDReader is a Reader for the GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalID structure.
type GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.

func (o *GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:

		result := NewGetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDOK creates a GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDOK with default headers values
func NewGetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDOK() *GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDOK {
	return &GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDOK{}
}

/*GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDOK handles this case with default header values.

Reversal reversal details
*/
type GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDOK struct {

	//Payload

	// isStream: false
	*models.DirectDebitReturnReversalDetailsResponse
}

func (o *GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDOK) Error() string {
	return fmt.Sprintf("[GET /transaction/directdebits/{id}/returns/{returnId}/reversals/{reversalId}][%d] getTransactionDirectdebitsIdReturnsReturnIdReversalsReversalIdOK  %+v", 200, o)
}

func (o *GetTransactionDirectdebitsIDReturnsReturnIDReversalsReversalIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.DirectDebitReturnReversalDetailsResponse = new(models.DirectDebitReturnReversalDetailsResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.DirectDebitReturnReversalDetailsResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}
