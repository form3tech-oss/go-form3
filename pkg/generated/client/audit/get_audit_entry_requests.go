// Code generated by go-swagger; DO NOT EDIT.

package audit

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// Client.GetAuditEntry creates a new GetAuditEntryRequest object
// with the default values initialized.
func (c *Client) GetAuditEntry() *GetAuditEntryRequest {
	var ()
	return &GetAuditEntryRequest{

		ID: c.Defaults.GetStrfmtUUID("GetAuditEntry", "id"),

		RecordType: c.Defaults.GetString("GetAuditEntry", "record_type"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetAuditEntryRequest struct {

	/*ID      Record Id      */

	ID strfmt.UUID

	/*RecordType      Record Type      */

	RecordType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetAuditEntryRequest) FromJson(j string) (*GetAuditEntryRequest, error) {

	return o, nil
}

func (o *GetAuditEntryRequest) WithID(id strfmt.UUID) *GetAuditEntryRequest {

	o.ID = id

	return o
}

func (o *GetAuditEntryRequest) WithRecordType(recordType string) *GetAuditEntryRequest {

	o.RecordType = recordType

	return o
}

// ////////////////
// WithContext adds the context to the get audit entry Request
func (o *GetAuditEntryRequest) WithContext(ctx context.Context) *GetAuditEntryRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get audit entry Request
func (o *GetAuditEntryRequest) WithHTTPClient(client *http.Client) *GetAuditEntryRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetAuditEntryRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param record_type
	if err := r.SetPathParam("record_type", o.RecordType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
