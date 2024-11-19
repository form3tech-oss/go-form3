// Code generated by go-swagger; DO NOT EDIT.

package query_api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Client.ListQueries creates a new ListQueriesRequest object
// with the default values initialized.
func (c *Client) ListQueries() *ListQueriesRequest {
	var ()
	return &ListQueriesRequest{

		FilterAutoHandled: c.Defaults.GetBoolPtr("ListQueries", "filter[auto_handled]"),

		FilterCreatedOnFrom: c.Defaults.GetStrfmtDateTimePtr("ListQueries", "filter[created_on_from]"),

		FilterCreatedOnTo: c.Defaults.GetStrfmtDateTimePtr("ListQueries", "filter[created_on_to]"),

		FilterOrganisationID: make([]strfmt.UUID, 0),

		FilterPaymentID: c.Defaults.GetStrfmtUUIDPtr("ListQueries", "filter[payment.id]"),

		FilterPaymentAdmissionID: c.Defaults.GetStrfmtUUIDPtr("ListQueries", "filter[payment_admission.id]"),

		FilterPaymentSubmissionID: c.Defaults.GetStrfmtUUIDPtr("ListQueries", "filter[payment_submission.id]"),

		FilterProcessingDateFrom: c.Defaults.GetStrfmtDatePtr("ListQueries", "filter[processing_date_from]"),

		FilterProcessingDateTo: c.Defaults.GetStrfmtDatePtr("ListQueries", "filter[processing_date_to]"),

		FilterQueryID: c.Defaults.GetStrfmtUUIDPtr("ListQueries", "filter[query.id]"),

		FilterQueryType: c.Defaults.GetStringPtr("ListQueries", "filter[query_type]"),

		FilterRecallID: c.Defaults.GetStrfmtUUIDPtr("ListQueries", "filter[recall.id]"),

		FilterRecallSubmissionID: c.Defaults.GetStrfmtUUIDPtr("ListQueries", "filter[recall_submission.id]"),

		FilterStatus: c.Defaults.GetStringPtr("ListQueries", "filter[status]"),

		PageNumber: c.Defaults.GetStringPtr("ListQueries", "page[number]"),

		PageSize: c.Defaults.GetInt64Ptr("ListQueries", "page[size]"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ListQueriesRequest struct {

	/*FilterAutoHandled      Find all queries for given auto handled flag      */

	FilterAutoHandled *bool

	/*FilterCreatedOnFrom      Find all queries from a certain created date.      */

	FilterCreatedOnFrom *strfmt.DateTime

	/*FilterCreatedOnTo      Find all queries up to a certain created date.      */

	FilterCreatedOnTo *strfmt.DateTime

	/*FilterOrganisationID      The organisations to filter on      */

	FilterOrganisationID []strfmt.UUID

	/*FilterPaymentID      Find all queries with a certain payment id.      */

	FilterPaymentID *strfmt.UUID

	/*FilterPaymentAdmissionID      Find all queries with a certain payment admission id.      */

	FilterPaymentAdmissionID *strfmt.UUID

	/*FilterPaymentSubmissionID      Find all queries with a certain payment submission id.      */

	FilterPaymentSubmissionID *strfmt.UUID

	/*FilterProcessingDateFrom      Find all queries from a certain value date.      */

	FilterProcessingDateFrom *strfmt.Date

	/*FilterProcessingDateTo      Find all queries up to a certain value date.      */

	FilterProcessingDateTo *strfmt.Date

	/*FilterQueryID      Find all queries with a certain query id.      */

	FilterQueryID *strfmt.UUID

	/*FilterQueryType      Find all queries for a given query type      */

	FilterQueryType *string

	/*FilterRecallID      Find all queries with a certain recall id.      */

	FilterRecallID *strfmt.UUID

	/*FilterRecallSubmissionID      Find all queries with a certain recall submission id.      */

	FilterRecallSubmissionID *strfmt.UUID

	/*FilterStatus      Find all queries for a given status      */

	FilterStatus *string

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

func (o *ListQueriesRequest) FromJson(j string) (*ListQueriesRequest, error) {

	return o, nil
}

func (o *ListQueriesRequest) WithFilterAutoHandled(filterAutoHandled bool) *ListQueriesRequest {

	o.FilterAutoHandled = &filterAutoHandled

	return o
}

func (o *ListQueriesRequest) WithoutFilterAutoHandled() *ListQueriesRequest {

	o.FilterAutoHandled = nil

	return o
}

func (o *ListQueriesRequest) WithFilterCreatedOnFrom(filterCreatedOnFrom strfmt.DateTime) *ListQueriesRequest {

	o.FilterCreatedOnFrom = &filterCreatedOnFrom

	return o
}

func (o *ListQueriesRequest) WithoutFilterCreatedOnFrom() *ListQueriesRequest {

	o.FilterCreatedOnFrom = nil

	return o
}

func (o *ListQueriesRequest) WithFilterCreatedOnTo(filterCreatedOnTo strfmt.DateTime) *ListQueriesRequest {

	o.FilterCreatedOnTo = &filterCreatedOnTo

	return o
}

func (o *ListQueriesRequest) WithoutFilterCreatedOnTo() *ListQueriesRequest {

	o.FilterCreatedOnTo = nil

	return o
}

func (o *ListQueriesRequest) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *ListQueriesRequest {

	o.FilterOrganisationID = filterOrganisationID

	return o
}

func (o *ListQueriesRequest) WithoutFilterOrganisationID() *ListQueriesRequest {

	o.FilterOrganisationID = nil

	return o
}

func (o *ListQueriesRequest) WithFilterPaymentID(filterPaymentID strfmt.UUID) *ListQueriesRequest {

	o.FilterPaymentID = &filterPaymentID

	return o
}

func (o *ListQueriesRequest) WithoutFilterPaymentID() *ListQueriesRequest {

	o.FilterPaymentID = nil

	return o
}

func (o *ListQueriesRequest) WithFilterPaymentAdmissionID(filterPaymentAdmissionID strfmt.UUID) *ListQueriesRequest {

	o.FilterPaymentAdmissionID = &filterPaymentAdmissionID

	return o
}

func (o *ListQueriesRequest) WithoutFilterPaymentAdmissionID() *ListQueriesRequest {

	o.FilterPaymentAdmissionID = nil

	return o
}

func (o *ListQueriesRequest) WithFilterPaymentSubmissionID(filterPaymentSubmissionID strfmt.UUID) *ListQueriesRequest {

	o.FilterPaymentSubmissionID = &filterPaymentSubmissionID

	return o
}

func (o *ListQueriesRequest) WithoutFilterPaymentSubmissionID() *ListQueriesRequest {

	o.FilterPaymentSubmissionID = nil

	return o
}

func (o *ListQueriesRequest) WithFilterProcessingDateFrom(filterProcessingDateFrom strfmt.Date) *ListQueriesRequest {

	o.FilterProcessingDateFrom = &filterProcessingDateFrom

	return o
}

func (o *ListQueriesRequest) WithoutFilterProcessingDateFrom() *ListQueriesRequest {

	o.FilterProcessingDateFrom = nil

	return o
}

func (o *ListQueriesRequest) WithFilterProcessingDateTo(filterProcessingDateTo strfmt.Date) *ListQueriesRequest {

	o.FilterProcessingDateTo = &filterProcessingDateTo

	return o
}

func (o *ListQueriesRequest) WithoutFilterProcessingDateTo() *ListQueriesRequest {

	o.FilterProcessingDateTo = nil

	return o
}

func (o *ListQueriesRequest) WithFilterQueryID(filterQueryID strfmt.UUID) *ListQueriesRequest {

	o.FilterQueryID = &filterQueryID

	return o
}

func (o *ListQueriesRequest) WithoutFilterQueryID() *ListQueriesRequest {

	o.FilterQueryID = nil

	return o
}

func (o *ListQueriesRequest) WithFilterQueryType(filterQueryType string) *ListQueriesRequest {

	o.FilterQueryType = &filterQueryType

	return o
}

func (o *ListQueriesRequest) WithoutFilterQueryType() *ListQueriesRequest {

	o.FilterQueryType = nil

	return o
}

func (o *ListQueriesRequest) WithFilterRecallID(filterRecallID strfmt.UUID) *ListQueriesRequest {

	o.FilterRecallID = &filterRecallID

	return o
}

func (o *ListQueriesRequest) WithoutFilterRecallID() *ListQueriesRequest {

	o.FilterRecallID = nil

	return o
}

func (o *ListQueriesRequest) WithFilterRecallSubmissionID(filterRecallSubmissionID strfmt.UUID) *ListQueriesRequest {

	o.FilterRecallSubmissionID = &filterRecallSubmissionID

	return o
}

func (o *ListQueriesRequest) WithoutFilterRecallSubmissionID() *ListQueriesRequest {

	o.FilterRecallSubmissionID = nil

	return o
}

func (o *ListQueriesRequest) WithFilterStatus(filterStatus string) *ListQueriesRequest {

	o.FilterStatus = &filterStatus

	return o
}

func (o *ListQueriesRequest) WithoutFilterStatus() *ListQueriesRequest {

	o.FilterStatus = nil

	return o
}

func (o *ListQueriesRequest) WithPageNumber(pageNumber string) *ListQueriesRequest {

	o.PageNumber = &pageNumber

	return o
}

func (o *ListQueriesRequest) WithoutPageNumber() *ListQueriesRequest {

	o.PageNumber = nil

	return o
}

func (o *ListQueriesRequest) WithPageSize(pageSize int64) *ListQueriesRequest {

	o.PageSize = &pageSize

	return o
}

func (o *ListQueriesRequest) WithoutPageSize() *ListQueriesRequest {

	o.PageSize = nil

	return o
}

// ////////////////
// WithContext adds the context to the list queries Request
func (o *ListQueriesRequest) WithContext(ctx context.Context) *ListQueriesRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the list queries Request
func (o *ListQueriesRequest) WithHTTPClient(client *http.Client) *ListQueriesRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ListQueriesRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterAutoHandled != nil {

		// query param filter[auto_handled]
		var qrFilterAutoHandled bool
		if o.FilterAutoHandled != nil {
			qrFilterAutoHandled = *o.FilterAutoHandled
		}
		qFilterAutoHandled := swag.FormatBool(qrFilterAutoHandled)
		if qFilterAutoHandled != "" {
			if err := r.SetQueryParam("filter[auto_handled]", qFilterAutoHandled); err != nil {
				return err
			}
		}

	}

	if o.FilterCreatedOnFrom != nil {

		// query param filter[created_on_from]
		var qrFilterCreatedOnFrom strfmt.DateTime
		if o.FilterCreatedOnFrom != nil {
			qrFilterCreatedOnFrom = *o.FilterCreatedOnFrom
		}
		qFilterCreatedOnFrom := qrFilterCreatedOnFrom.String()
		if qFilterCreatedOnFrom != "" {
			if err := r.SetQueryParam("filter[created_on_from]", qFilterCreatedOnFrom); err != nil {
				return err
			}
		}

	}

	if o.FilterCreatedOnTo != nil {

		// query param filter[created_on_to]
		var qrFilterCreatedOnTo strfmt.DateTime
		if o.FilterCreatedOnTo != nil {
			qrFilterCreatedOnTo = *o.FilterCreatedOnTo
		}
		qFilterCreatedOnTo := qrFilterCreatedOnTo.String()
		if qFilterCreatedOnTo != "" {
			if err := r.SetQueryParam("filter[created_on_to]", qFilterCreatedOnTo); err != nil {
				return err
			}
		}

	}

	var valuesFilterOrganisationID []string
	for _, v := range o.FilterOrganisationID {
		valuesFilterOrganisationID = append(valuesFilterOrganisationID, v.String())
	}

	joinedFilterOrganisationID := swag.JoinByFormat(valuesFilterOrganisationID, "")
	// query array param filter[organisation_id]
	if err := r.SetQueryParam("filter[organisation_id]", joinedFilterOrganisationID...); err != nil {
		return err
	}

	if o.FilterPaymentID != nil {

		// query param filter[payment.id]
		var qrFilterPaymentID strfmt.UUID
		if o.FilterPaymentID != nil {
			qrFilterPaymentID = *o.FilterPaymentID
		}
		qFilterPaymentID := qrFilterPaymentID.String()
		if qFilterPaymentID != "" {
			if err := r.SetQueryParam("filter[payment.id]", qFilterPaymentID); err != nil {
				return err
			}
		}

	}

	if o.FilterPaymentAdmissionID != nil {

		// query param filter[payment_admission.id]
		var qrFilterPaymentAdmissionID strfmt.UUID
		if o.FilterPaymentAdmissionID != nil {
			qrFilterPaymentAdmissionID = *o.FilterPaymentAdmissionID
		}
		qFilterPaymentAdmissionID := qrFilterPaymentAdmissionID.String()
		if qFilterPaymentAdmissionID != "" {
			if err := r.SetQueryParam("filter[payment_admission.id]", qFilterPaymentAdmissionID); err != nil {
				return err
			}
		}

	}

	if o.FilterPaymentSubmissionID != nil {

		// query param filter[payment_submission.id]
		var qrFilterPaymentSubmissionID strfmt.UUID
		if o.FilterPaymentSubmissionID != nil {
			qrFilterPaymentSubmissionID = *o.FilterPaymentSubmissionID
		}
		qFilterPaymentSubmissionID := qrFilterPaymentSubmissionID.String()
		if qFilterPaymentSubmissionID != "" {
			if err := r.SetQueryParam("filter[payment_submission.id]", qFilterPaymentSubmissionID); err != nil {
				return err
			}
		}

	}

	if o.FilterProcessingDateFrom != nil {

		// query param filter[processing_date_from]
		var qrFilterProcessingDateFrom strfmt.Date
		if o.FilterProcessingDateFrom != nil {
			qrFilterProcessingDateFrom = *o.FilterProcessingDateFrom
		}
		qFilterProcessingDateFrom := qrFilterProcessingDateFrom.String()
		if qFilterProcessingDateFrom != "" {
			if err := r.SetQueryParam("filter[processing_date_from]", qFilterProcessingDateFrom); err != nil {
				return err
			}
		}

	}

	if o.FilterProcessingDateTo != nil {

		// query param filter[processing_date_to]
		var qrFilterProcessingDateTo strfmt.Date
		if o.FilterProcessingDateTo != nil {
			qrFilterProcessingDateTo = *o.FilterProcessingDateTo
		}
		qFilterProcessingDateTo := qrFilterProcessingDateTo.String()
		if qFilterProcessingDateTo != "" {
			if err := r.SetQueryParam("filter[processing_date_to]", qFilterProcessingDateTo); err != nil {
				return err
			}
		}

	}

	if o.FilterQueryID != nil {

		// query param filter[query.id]
		var qrFilterQueryID strfmt.UUID
		if o.FilterQueryID != nil {
			qrFilterQueryID = *o.FilterQueryID
		}
		qFilterQueryID := qrFilterQueryID.String()
		if qFilterQueryID != "" {
			if err := r.SetQueryParam("filter[query.id]", qFilterQueryID); err != nil {
				return err
			}
		}

	}

	if o.FilterQueryType != nil {

		// query param filter[query_type]
		var qrFilterQueryType string
		if o.FilterQueryType != nil {
			qrFilterQueryType = *o.FilterQueryType
		}
		qFilterQueryType := qrFilterQueryType
		if qFilterQueryType != "" {
			if err := r.SetQueryParam("filter[query_type]", qFilterQueryType); err != nil {
				return err
			}
		}

	}

	if o.FilterRecallID != nil {

		// query param filter[recall.id]
		var qrFilterRecallID strfmt.UUID
		if o.FilterRecallID != nil {
			qrFilterRecallID = *o.FilterRecallID
		}
		qFilterRecallID := qrFilterRecallID.String()
		if qFilterRecallID != "" {
			if err := r.SetQueryParam("filter[recall.id]", qFilterRecallID); err != nil {
				return err
			}
		}

	}

	if o.FilterRecallSubmissionID != nil {

		// query param filter[recall_submission.id]
		var qrFilterRecallSubmissionID strfmt.UUID
		if o.FilterRecallSubmissionID != nil {
			qrFilterRecallSubmissionID = *o.FilterRecallSubmissionID
		}
		qFilterRecallSubmissionID := qrFilterRecallSubmissionID.String()
		if qFilterRecallSubmissionID != "" {
			if err := r.SetQueryParam("filter[recall_submission.id]", qFilterRecallSubmissionID); err != nil {
				return err
			}
		}

	}

	if o.FilterStatus != nil {

		// query param filter[status]
		var qrFilterStatus string
		if o.FilterStatus != nil {
			qrFilterStatus = *o.FilterStatus
		}
		qFilterStatus := qrFilterStatus
		if qFilterStatus != "" {
			if err := r.SetQueryParam("filter[status]", qFilterStatus); err != nil {
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
