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

// Client.GetAddressbookPartiesIDActors creates a new GetAddressbookPartiesIDActorsRequest object
// with the default values initialized.
func (c *Client) GetAddressbookPartiesIDActors() *GetAddressbookPartiesIDActorsRequest {
	var (
		pageSizeDefault = int64(100)
	)
	return &GetAddressbookPartiesIDActorsRequest{

		FilterCustomerID: c.Defaults.GetStringPtr("GetAddressbookPartiesIDActors", "filter[customer_id]"),

		ID: c.Defaults.GetStrfmtUUID("GetAddressbookPartiesIDActors", "id"),

		PageNumber: c.Defaults.GetStringPtr("GetAddressbookPartiesIDActors", "page[number]"),

		PageSize: &pageSizeDefault,

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetAddressbookPartiesIDActorsRequest struct {

	/*FilterCustomerID      Filter parties by customer_id value      */

	FilterCustomerID *string

	/*ID      Id of party      */

	ID strfmt.UUID

	/*PageNumber      Which page to select      */

	PageNumber *string

	/*PageSize      Number of items to select      */

	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetAddressbookPartiesIDActorsRequest) FromJson(j string) (*GetAddressbookPartiesIDActorsRequest, error) {

	return o, nil
}

func (o *GetAddressbookPartiesIDActorsRequest) WithFilterCustomerID(filterCustomerID string) *GetAddressbookPartiesIDActorsRequest {

	o.FilterCustomerID = &filterCustomerID

	return o
}

func (o *GetAddressbookPartiesIDActorsRequest) WithoutFilterCustomerID() *GetAddressbookPartiesIDActorsRequest {

	o.FilterCustomerID = nil

	return o
}

func (o *GetAddressbookPartiesIDActorsRequest) WithID(id strfmt.UUID) *GetAddressbookPartiesIDActorsRequest {

	o.ID = id

	return o
}

func (o *GetAddressbookPartiesIDActorsRequest) WithPageNumber(pageNumber string) *GetAddressbookPartiesIDActorsRequest {

	o.PageNumber = &pageNumber

	return o
}

func (o *GetAddressbookPartiesIDActorsRequest) WithoutPageNumber() *GetAddressbookPartiesIDActorsRequest {

	o.PageNumber = nil

	return o
}

func (o *GetAddressbookPartiesIDActorsRequest) WithPageSize(pageSize int64) *GetAddressbookPartiesIDActorsRequest {

	o.PageSize = &pageSize

	return o
}

func (o *GetAddressbookPartiesIDActorsRequest) WithoutPageSize() *GetAddressbookPartiesIDActorsRequest {

	o.PageSize = nil

	return o
}

//////////////////
// WithContext adds the context to the get addressbook parties ID actors Request
func (o *GetAddressbookPartiesIDActorsRequest) WithContext(ctx context.Context) *GetAddressbookPartiesIDActorsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get addressbook parties ID actors Request
func (o *GetAddressbookPartiesIDActorsRequest) WithHTTPClient(client *http.Client) *GetAddressbookPartiesIDActorsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetAddressbookPartiesIDActorsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterCustomerID != nil {

		// query param filter[customer_id]
		var qrFilterCustomerID string
		if o.FilterCustomerID != nil {
			qrFilterCustomerID = *o.FilterCustomerID
		}
		qFilterCustomerID := qrFilterCustomerID
		if qFilterCustomerID != "" {
			if err := r.SetQueryParam("filter[customer_id]", qFilterCustomerID); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.PageNumber != nil {

		// query param page[number]
		var qrPageNumber string
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := qrPageNumber
		if qPageNumber != "" {
			if err := r.SetQueryParam("page[number]", qPageNumber); err != nil {
				return err
			}
		}

	}

	if o.PageSize != nil {

		// query param page[size]
		var qrPageSize int64
		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {
			if err := r.SetQueryParam("page[size]", qPageSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
