// Code generated by go-swagger; DO NOT EDIT.

package query_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v5/pkg/generated/models"
)

// GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDReader is a Reader for the GetTransactionQueriesQueryIDSubmissionsQuerySubmissionID structure.
type GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDOK creates a GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDOK with default headers values
func NewGetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDOK() *GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDOK {
	return &GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDOK{}
}

/*GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDOK handles this case with default header values.

query submission response
*/
type GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDOK struct {

	//Payload

	// isStream: false
	*models.QuerySubmissionResponse
}

func (o *GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDOK) Error() string {
	return fmt.Sprintf("[GET /transaction/queries/{query_id}/submissions/{query_submission_id}][%d] getTransactionQueriesQueryIdSubmissionsQuerySubmissionIdOK", 200)
}

func (o *GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.QuerySubmissionResponse = new(models.QuerySubmissionResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.QuerySubmissionResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDBadRequest creates a GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDBadRequest with default headers values
func NewGetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDBadRequest() *GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDBadRequest {
	return &GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDBadRequest{}
}

/*GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDBadRequest handles this case with default header values.

bad request
*/
type GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDBadRequest struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /transaction/queries/{query_id}/submissions/{query_submission_id}][%d] getTransactionQueriesQueryIdSubmissionsQuerySubmissionIdBadRequest", 400)
}

func (o *GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDForbidden creates a GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDForbidden with default headers values
func NewGetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDForbidden() *GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDForbidden {
	return &GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDForbidden{}
}

/*GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDForbidden handles this case with default header values.

forbidden
*/
type GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDForbidden struct {

	//Payload

	// isStream: false
	*models.APIError
}

func (o *GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDForbidden) Error() string {
	return fmt.Sprintf("[GET /transaction/queries/{query_id}/submissions/{query_submission_id}][%d] getTransactionQueriesQueryIdSubmissionsQuerySubmissionIdForbidden", 403)
}

func (o *GetTransactionQueriesQueryIDSubmissionsQuerySubmissionIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.APIError = new(models.APIError)

	// response payload

	if err := consumer.Consume(response.Body(), o.APIError); err != nil && err != io.EOF {
		return err
	}

	return nil
}
