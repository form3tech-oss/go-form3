// Code generated by go-swagger; DO NOT EDIT.

package metrics_api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Client.GetMetricsPrometheusAPIV1Query creates a new GetMetricsPrometheusAPIV1QueryRequest object
// with the default values initialized.
func (c *Client) GetMetricsPrometheusAPIV1Query() *GetMetricsPrometheusAPIV1QueryRequest {
	var ()
	return &GetMetricsPrometheusAPIV1QueryRequest{

		Query: c.Defaults.GetString("GetMetricsPrometheusAPIV1Query", "query"),

		Time: c.Defaults.GetStringPtr("GetMetricsPrometheusAPIV1Query", "time"),

		Timeout: c.Defaults.GetStringPtr("GetMetricsPrometheusAPIV1Query", "timeout"),

		requestTimeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetMetricsPrometheusAPIV1QueryRequest struct {

	/*Query      Query to Execute      */

	Query string

	/*Time      RFC3339 or unix_timestamp      */

	Time *string

	/*Timeout      See https://prometheus.io/docs/prometheus/latest/querying/basics/#time-durations      */

	Timeout *string

	requestTimeout time.Duration
	Context        context.Context
	HTTPClient     *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetMetricsPrometheusAPIV1QueryRequest) FromJson(j string) (*GetMetricsPrometheusAPIV1QueryRequest, error) {

	return o, nil
}

func (o *GetMetricsPrometheusAPIV1QueryRequest) WithQuery(query string) *GetMetricsPrometheusAPIV1QueryRequest {

	o.Query = query

	return o
}

func (o *GetMetricsPrometheusAPIV1QueryRequest) WithTime(time string) *GetMetricsPrometheusAPIV1QueryRequest {

	o.Time = &time

	return o
}

func (o *GetMetricsPrometheusAPIV1QueryRequest) WithoutTime() *GetMetricsPrometheusAPIV1QueryRequest {

	o.Time = nil

	return o
}

func (o *GetMetricsPrometheusAPIV1QueryRequest) WithTimeout(timeout string) *GetMetricsPrometheusAPIV1QueryRequest {

	o.Timeout = &timeout

	return o
}

func (o *GetMetricsPrometheusAPIV1QueryRequest) WithoutTimeout() *GetMetricsPrometheusAPIV1QueryRequest {

	o.Timeout = nil

	return o
}

// ////////////////
// WithContext adds the context to the get metrics prometheus API v1 query Request
func (o *GetMetricsPrometheusAPIV1QueryRequest) WithContext(ctx context.Context) *GetMetricsPrometheusAPIV1QueryRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get metrics prometheus API v1 query Request
func (o *GetMetricsPrometheusAPIV1QueryRequest) WithHTTPClient(client *http.Client) *GetMetricsPrometheusAPIV1QueryRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetMetricsPrometheusAPIV1QueryRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.requestTimeout); err != nil {
		return err
	}
	var res []error

	// query param query
	qrQuery := o.Query
	qQuery := qrQuery
	if qQuery != "" {
		if err := r.SetQueryParam("query", qQuery); err != nil {
			return err
		}
	}

	if o.Time != nil {

		// query param time
		var qrTime string
		if o.Time != nil {
			qrTime = *o.Time
		}
		qTime := qrTime
		if qTime != "" {
			if err := r.SetQueryParam("time", qTime); err != nil {
				return err
			}
		}

	}

	if o.Timeout != nil {

		// query param timeout
		var qrTimeout string
		if o.Timeout != nil {
			qrTimeout = *o.Timeout
		}
		qTimeout := qrTimeout
		if qTimeout != "" {
			if err := r.SetQueryParam("timeout", qTimeout); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
