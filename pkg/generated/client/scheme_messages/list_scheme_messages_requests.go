// Code generated by go-swagger; DO NOT EDIT.

package scheme_messages

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

// Client.ListSchemeMessages creates a new ListSchemeMessagesRequest object
// with the default values initialized.
func (c *Client) ListSchemeMessages() *ListSchemeMessagesRequest {
	var (
		pageSizeDefault = int64(100)
	)
	return &ListSchemeMessagesRequest{

		FilterAdmissionAdmissionDateFrom: c.Defaults.GetStrfmtDateTimePtr("ListSchemeMessages", "filter[admission.admission_date_from]"),

		FilterAdmissionAdmissionDateTo: c.Defaults.GetStrfmtDateTimePtr("ListSchemeMessages", "filter[admission.admission_date_to]"),

		FilterPaymentScheme: c.Defaults.GetStringPtr("ListSchemeMessages", "filter[payment_scheme]"),

		FilterSchemeMessageType: c.Defaults.GetStringPtr("ListSchemeMessages", "filter[scheme_message_type]"),

		FilterUniqueSchemeID: c.Defaults.GetStringPtr("ListSchemeMessages", "filter[unique_scheme_id]"),

		PageNumber: c.Defaults.GetStringPtr("ListSchemeMessages", "page[number]"),

		PageSize: &pageSizeDefault,

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ListSchemeMessagesRequest struct {

	/*FilterAdmissionAdmissionDateFrom      Filter by Admission DateTime      */

	FilterAdmissionAdmissionDateFrom *strfmt.DateTime

	/*FilterAdmissionAdmissionDateTo      Filter by Admission DateTime      */

	FilterAdmissionAdmissionDateTo *strfmt.DateTime

	/*FilterPaymentScheme      Filter by Payment Scheme      */

	FilterPaymentScheme *string

	/*FilterSchemeMessageType      Filter by Scheme Message Type      */

	FilterSchemeMessageType *string

	/*FilterUniqueSchemeID      Filter by Unique SchemeId      */

	FilterUniqueSchemeID *string

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

func (o *ListSchemeMessagesRequest) FromJson(j string) (*ListSchemeMessagesRequest, error) {

	return o, nil
}

func (o *ListSchemeMessagesRequest) WithFilterAdmissionAdmissionDateFrom(filterAdmissionAdmissionDateFrom strfmt.DateTime) *ListSchemeMessagesRequest {

	o.FilterAdmissionAdmissionDateFrom = &filterAdmissionAdmissionDateFrom

	return o
}

func (o *ListSchemeMessagesRequest) WithoutFilterAdmissionAdmissionDateFrom() *ListSchemeMessagesRequest {

	o.FilterAdmissionAdmissionDateFrom = nil

	return o
}

func (o *ListSchemeMessagesRequest) WithFilterAdmissionAdmissionDateTo(filterAdmissionAdmissionDateTo strfmt.DateTime) *ListSchemeMessagesRequest {

	o.FilterAdmissionAdmissionDateTo = &filterAdmissionAdmissionDateTo

	return o
}

func (o *ListSchemeMessagesRequest) WithoutFilterAdmissionAdmissionDateTo() *ListSchemeMessagesRequest {

	o.FilterAdmissionAdmissionDateTo = nil

	return o
}

func (o *ListSchemeMessagesRequest) WithFilterPaymentScheme(filterPaymentScheme string) *ListSchemeMessagesRequest {

	o.FilterPaymentScheme = &filterPaymentScheme

	return o
}

func (o *ListSchemeMessagesRequest) WithoutFilterPaymentScheme() *ListSchemeMessagesRequest {

	o.FilterPaymentScheme = nil

	return o
}

func (o *ListSchemeMessagesRequest) WithFilterSchemeMessageType(filterSchemeMessageType string) *ListSchemeMessagesRequest {

	o.FilterSchemeMessageType = &filterSchemeMessageType

	return o
}

func (o *ListSchemeMessagesRequest) WithoutFilterSchemeMessageType() *ListSchemeMessagesRequest {

	o.FilterSchemeMessageType = nil

	return o
}

func (o *ListSchemeMessagesRequest) WithFilterUniqueSchemeID(filterUniqueSchemeID string) *ListSchemeMessagesRequest {

	o.FilterUniqueSchemeID = &filterUniqueSchemeID

	return o
}

func (o *ListSchemeMessagesRequest) WithoutFilterUniqueSchemeID() *ListSchemeMessagesRequest {

	o.FilterUniqueSchemeID = nil

	return o
}

func (o *ListSchemeMessagesRequest) WithPageNumber(pageNumber string) *ListSchemeMessagesRequest {

	o.PageNumber = &pageNumber

	return o
}

func (o *ListSchemeMessagesRequest) WithoutPageNumber() *ListSchemeMessagesRequest {

	o.PageNumber = nil

	return o
}

func (o *ListSchemeMessagesRequest) WithPageSize(pageSize int64) *ListSchemeMessagesRequest {

	o.PageSize = &pageSize

	return o
}

func (o *ListSchemeMessagesRequest) WithoutPageSize() *ListSchemeMessagesRequest {

	o.PageSize = nil

	return o
}

// ////////////////
// WithContext adds the context to the list scheme messages Request
func (o *ListSchemeMessagesRequest) WithContext(ctx context.Context) *ListSchemeMessagesRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the list scheme messages Request
func (o *ListSchemeMessagesRequest) WithHTTPClient(client *http.Client) *ListSchemeMessagesRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ListSchemeMessagesRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterAdmissionAdmissionDateFrom != nil {

		// query param filter[admission.admission_date_from]
		var qrFilterAdmissionAdmissionDateFrom strfmt.DateTime
		if o.FilterAdmissionAdmissionDateFrom != nil {
			qrFilterAdmissionAdmissionDateFrom = *o.FilterAdmissionAdmissionDateFrom
		}
		qFilterAdmissionAdmissionDateFrom := qrFilterAdmissionAdmissionDateFrom.String()
		if qFilterAdmissionAdmissionDateFrom != "" {
			if err := r.SetQueryParam("filter[admission.admission_date_from]", qFilterAdmissionAdmissionDateFrom); err != nil {
				return err
			}
		}

	}

	if o.FilterAdmissionAdmissionDateTo != nil {

		// query param filter[admission.admission_date_to]
		var qrFilterAdmissionAdmissionDateTo strfmt.DateTime
		if o.FilterAdmissionAdmissionDateTo != nil {
			qrFilterAdmissionAdmissionDateTo = *o.FilterAdmissionAdmissionDateTo
		}
		qFilterAdmissionAdmissionDateTo := qrFilterAdmissionAdmissionDateTo.String()
		if qFilterAdmissionAdmissionDateTo != "" {
			if err := r.SetQueryParam("filter[admission.admission_date_to]", qFilterAdmissionAdmissionDateTo); err != nil {
				return err
			}
		}

	}

	if o.FilterPaymentScheme != nil {

		// query param filter[payment_scheme]
		var qrFilterPaymentScheme string
		if o.FilterPaymentScheme != nil {
			qrFilterPaymentScheme = *o.FilterPaymentScheme
		}
		qFilterPaymentScheme := qrFilterPaymentScheme
		if qFilterPaymentScheme != "" {
			if err := r.SetQueryParam("filter[payment_scheme]", qFilterPaymentScheme); err != nil {
				return err
			}
		}

	}

	if o.FilterSchemeMessageType != nil {

		// query param filter[scheme_message_type]
		var qrFilterSchemeMessageType string
		if o.FilterSchemeMessageType != nil {
			qrFilterSchemeMessageType = *o.FilterSchemeMessageType
		}
		qFilterSchemeMessageType := qrFilterSchemeMessageType
		if qFilterSchemeMessageType != "" {
			if err := r.SetQueryParam("filter[scheme_message_type]", qFilterSchemeMessageType); err != nil {
				return err
			}
		}

	}

	if o.FilterUniqueSchemeID != nil {

		// query param filter[unique_scheme_id]
		var qrFilterUniqueSchemeID string
		if o.FilterUniqueSchemeID != nil {
			qrFilterUniqueSchemeID = *o.FilterUniqueSchemeID
		}
		qFilterUniqueSchemeID := qrFilterUniqueSchemeID
		if qFilterUniqueSchemeID != "" {
			if err := r.SetQueryParam("filter[unique_scheme_id]", qFilterUniqueSchemeID); err != nil {
				return err
			}
		}

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
