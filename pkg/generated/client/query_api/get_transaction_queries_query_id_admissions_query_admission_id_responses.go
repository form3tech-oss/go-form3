// Code generated by go-swagger; DO NOT EDIT.

package query_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDReader is a Reader for the GetTransactionQueriesQueryIDAdmissionsQueryAdmissionID structure.
type GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDOK creates a GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDOK with default headers values
func NewGetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDOK() *GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDOK {
	return &GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDOK{}
}

/*GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDOK handles this case with default header values.

query admission response
*/
type GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDOK struct {

	//Payload

	// isStream: false
	*models.QueryAdmissionResponse
}

func (o *GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDOK) Error() string {
	return fmt.Sprintf("[GET /transaction/queries/{query_id}/admissions/{query_admission_id}][%d] getTransactionQueriesQueryIdAdmissionsQueryAdmissionIdOK", 200)
}

func (o *GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.QueryAdmissionResponse = new(models.QueryAdmissionResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.QueryAdmissionResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDBadRequest creates a GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDBadRequest with default headers values
func NewGetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDBadRequest() *GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDBadRequest {
	return &GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDBadRequest{}
}

/*GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDBadRequest handles this case with default header values.

bad request
*/
type GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /transaction/queries/{query_id}/admissions/{query_admission_id}][%d] getTransactionQueriesQueryIdAdmissionsQueryAdmissionIdBadRequest", 400)
}

func (o *GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDForbidden creates a GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDForbidden with default headers values
func NewGetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDForbidden() *GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDForbidden {
	return &GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDForbidden{}
}

/*GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDForbidden handles this case with default header values.

forbidden
*/
type GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDForbidden) Error() string {
	return fmt.Sprintf("[GET /transaction/queries/{query_id}/admissions/{query_admission_id}][%d] getTransactionQueriesQueryIdAdmissionsQueryAdmissionIdForbidden", 403)
}

func (o *GetTransactionQueriesQueryIDAdmissionsQueryAdmissionIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
