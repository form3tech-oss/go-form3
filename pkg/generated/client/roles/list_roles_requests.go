// Code generated by go-swagger; DO NOT EDIT.
// HELLO :form3: !

package roles

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

// Client.ListRoles creates a new ListRolesRequest object
// with the default values initialized.
func (c *Client) ListRoles() *ListRolesRequest {
	var ()
	return &ListRolesRequest{

		PageNumber: c.Defaults.GetInt64Ptr("ListRoles", "page[number]"),

		PageSize: c.Defaults.GetInt64Ptr("ListRoles", "page[size]"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ListRolesRequest struct {

	/*PageNumber      Which page to select      */

	PageNumber *int64

	/*PageSize      Number of items to select      */

	PageSize *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *ListRolesRequest) FromJson(j string) *ListRolesRequest {

	return o
}

func (o *ListRolesRequest) WithPageNumber(pageNumber int64) *ListRolesRequest {

	o.PageNumber = &pageNumber

	return o
}

func (o *ListRolesRequest) WithoutPageNumber() *ListRolesRequest {

	o.PageNumber = nil

	return o
}

func (o *ListRolesRequest) WithPageSize(pageSize int64) *ListRolesRequest {

	o.PageSize = &pageSize

	return o
}

func (o *ListRolesRequest) WithoutPageSize() *ListRolesRequest {

	o.PageSize = nil

	return o
}

//////////////////
// WithContext adds the context to the list roles Request
func (o *ListRolesRequest) WithContext(ctx context.Context) *ListRolesRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the list roles Request
func (o *ListRolesRequest) WithHTTPClient(client *http.Client) *ListRolesRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ListRolesRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PageNumber != nil {

		// query param page[number]
		var qrPageNumber int64
		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt64(qrPageNumber)
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
