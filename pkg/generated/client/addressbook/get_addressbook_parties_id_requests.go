// Code generated by go-swagger; DO NOT EDIT.

package addressbook

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetAddressbookPartiesID creates a new GetAddressbookPartiesIDRequest object
// with the default values initialized.
func (c *Client) GetAddressbookPartiesID() *GetAddressbookPartiesIDRequest {
	var ()
	return &GetAddressbookPartiesIDRequest{

		ID: c.Defaults.GetStrfmtUUID("GetAddressbookPartiesID", "id"),

		Version: c.Defaults.GetInt64Ptr("GetAddressbookPartiesID", "version"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetAddressbookPartiesIDRequest struct {

	/*ID      Id of party      */

	ID strfmt.UUID

	/*Version*/

	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetAddressbookPartiesIDRequest) FromJson(j string) *GetAddressbookPartiesIDRequest {

	return o
}

func (o *GetAddressbookPartiesIDRequest) WithID(id strfmt.UUID) *GetAddressbookPartiesIDRequest {

	o.ID = id

	return o
}

func (o *GetAddressbookPartiesIDRequest) WithVersion(version int64) *GetAddressbookPartiesIDRequest {

	o.Version = &version

	return o
}

func (o *GetAddressbookPartiesIDRequest) WithoutVersion() *GetAddressbookPartiesIDRequest {

	o.Version = nil

	return o
}

//////////////////
// WithContext adds the context to the get addressbook parties ID Request
func (o *GetAddressbookPartiesIDRequest) WithContext(ctx context.Context) *GetAddressbookPartiesIDRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get addressbook parties ID Request
func (o *GetAddressbookPartiesIDRequest) WithHTTPClient(client *http.Client) *GetAddressbookPartiesIDRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetAddressbookPartiesIDRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion int64
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}