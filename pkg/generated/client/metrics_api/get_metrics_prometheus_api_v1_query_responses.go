// Code generated by go-swagger; DO NOT EDIT.

package metrics_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/models"
)

// GetMetricsPrometheusAPIV1QueryReader is a Reader for the GetMetricsPrometheusAPIV1Query structure.
type GetMetricsPrometheusAPIV1QueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMetricsPrometheusAPIV1QueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMetricsPrometheusAPIV1QueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetMetricsPrometheusAPIV1QueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetMetricsPrometheusAPIV1QueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetMetricsPrometheusAPIV1QueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMetricsPrometheusAPIV1QueryOK creates a GetMetricsPrometheusAPIV1QueryOK with default headers values
func NewGetMetricsPrometheusAPIV1QueryOK() *GetMetricsPrometheusAPIV1QueryOK {
	return &GetMetricsPrometheusAPIV1QueryOK{}
}

/*
GetMetricsPrometheusAPIV1QueryOK handles this case with default header values.

See Prometheus api https://prometheus.io/docs/prometheus/latest/querying/api/#instant-queries
*/
type GetMetricsPrometheusAPIV1QueryOK struct {

	//Payload

	// isStream: false
	*models.MetricsQueryResponse
}

// IsSuccess returns true when this get metrics prometheus Api v1 query o k response has a 2xx status code
func (o *GetMetricsPrometheusAPIV1QueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get metrics prometheus Api v1 query o k response has a 3xx status code
func (o *GetMetricsPrometheusAPIV1QueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get metrics prometheus Api v1 query o k response has a 4xx status code
func (o *GetMetricsPrometheusAPIV1QueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get metrics prometheus Api v1 query o k response has a 5xx status code
func (o *GetMetricsPrometheusAPIV1QueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get metrics prometheus Api v1 query o k response a status code equal to that given
func (o *GetMetricsPrometheusAPIV1QueryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get metrics prometheus Api v1 query o k response
func (o *GetMetricsPrometheusAPIV1QueryOK) Code() int {
	return 200
}

func (o *GetMetricsPrometheusAPIV1QueryOK) Error() string {
	return fmt.Sprintf("[GET /metrics/prometheus/api/v1/query][%d] getMetricsPrometheusApiV1QueryOK", 200)
}

func (o *GetMetricsPrometheusAPIV1QueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.MetricsQueryResponse = new(models.MetricsQueryResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.MetricsQueryResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetricsPrometheusAPIV1QueryBadRequest creates a GetMetricsPrometheusAPIV1QueryBadRequest with default headers values
func NewGetMetricsPrometheusAPIV1QueryBadRequest() *GetMetricsPrometheusAPIV1QueryBadRequest {
	return &GetMetricsPrometheusAPIV1QueryBadRequest{}
}

/*
GetMetricsPrometheusAPIV1QueryBadRequest handles this case with default header values.

Bad Request
*/
type GetMetricsPrometheusAPIV1QueryBadRequest struct {

	//Payload

	// isStream: false
	*models.MetricsQueryResponse
}

// IsSuccess returns true when this get metrics prometheus Api v1 query bad request response has a 2xx status code
func (o *GetMetricsPrometheusAPIV1QueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get metrics prometheus Api v1 query bad request response has a 3xx status code
func (o *GetMetricsPrometheusAPIV1QueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get metrics prometheus Api v1 query bad request response has a 4xx status code
func (o *GetMetricsPrometheusAPIV1QueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get metrics prometheus Api v1 query bad request response has a 5xx status code
func (o *GetMetricsPrometheusAPIV1QueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get metrics prometheus Api v1 query bad request response a status code equal to that given
func (o *GetMetricsPrometheusAPIV1QueryBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get metrics prometheus Api v1 query bad request response
func (o *GetMetricsPrometheusAPIV1QueryBadRequest) Code() int {
	return 400
}

func (o *GetMetricsPrometheusAPIV1QueryBadRequest) Error() string {
	return fmt.Sprintf("[GET /metrics/prometheus/api/v1/query][%d] getMetricsPrometheusApiV1QueryBadRequest", 400)
}

func (o *GetMetricsPrometheusAPIV1QueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.MetricsQueryResponse = new(models.MetricsQueryResponse)

	// response payload

	if err := consumer.Consume(response.Body(), o.MetricsQueryResponse); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMetricsPrometheusAPIV1QueryForbidden creates a GetMetricsPrometheusAPIV1QueryForbidden with default headers values
func NewGetMetricsPrometheusAPIV1QueryForbidden() *GetMetricsPrometheusAPIV1QueryForbidden {
	return &GetMetricsPrometheusAPIV1QueryForbidden{}
}

/*
GetMetricsPrometheusAPIV1QueryForbidden handles this case with default header values.

Forbidden
*/
type GetMetricsPrometheusAPIV1QueryForbidden struct {
}

// IsSuccess returns true when this get metrics prometheus Api v1 query forbidden response has a 2xx status code
func (o *GetMetricsPrometheusAPIV1QueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get metrics prometheus Api v1 query forbidden response has a 3xx status code
func (o *GetMetricsPrometheusAPIV1QueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get metrics prometheus Api v1 query forbidden response has a 4xx status code
func (o *GetMetricsPrometheusAPIV1QueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get metrics prometheus Api v1 query forbidden response has a 5xx status code
func (o *GetMetricsPrometheusAPIV1QueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get metrics prometheus Api v1 query forbidden response a status code equal to that given
func (o *GetMetricsPrometheusAPIV1QueryForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get metrics prometheus Api v1 query forbidden response
func (o *GetMetricsPrometheusAPIV1QueryForbidden) Code() int {
	return 403
}

func (o *GetMetricsPrometheusAPIV1QueryForbidden) Error() string {
	return fmt.Sprintf("[GET /metrics/prometheus/api/v1/query][%d] getMetricsPrometheusApiV1QueryForbidden", 403)
}

func (o *GetMetricsPrometheusAPIV1QueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMetricsPrometheusAPIV1QueryInternalServerError creates a GetMetricsPrometheusAPIV1QueryInternalServerError with default headers values
func NewGetMetricsPrometheusAPIV1QueryInternalServerError() *GetMetricsPrometheusAPIV1QueryInternalServerError {
	return &GetMetricsPrometheusAPIV1QueryInternalServerError{}
}

/*
GetMetricsPrometheusAPIV1QueryInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetMetricsPrometheusAPIV1QueryInternalServerError struct {
}

// IsSuccess returns true when this get metrics prometheus Api v1 query internal server error response has a 2xx status code
func (o *GetMetricsPrometheusAPIV1QueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get metrics prometheus Api v1 query internal server error response has a 3xx status code
func (o *GetMetricsPrometheusAPIV1QueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get metrics prometheus Api v1 query internal server error response has a 4xx status code
func (o *GetMetricsPrometheusAPIV1QueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get metrics prometheus Api v1 query internal server error response has a 5xx status code
func (o *GetMetricsPrometheusAPIV1QueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get metrics prometheus Api v1 query internal server error response a status code equal to that given
func (o *GetMetricsPrometheusAPIV1QueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get metrics prometheus Api v1 query internal server error response
func (o *GetMetricsPrometheusAPIV1QueryInternalServerError) Code() int {
	return 500
}

func (o *GetMetricsPrometheusAPIV1QueryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /metrics/prometheus/api/v1/query][%d] getMetricsPrometheusApiV1QueryInternalServerError", 500)
}

func (o *GetMetricsPrometheusAPIV1QueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
