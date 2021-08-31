// Code generated by go-swagger; DO NOT EDIT.

package account_amendment

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

// Client.ListAccountAmendments creates a new ListAccountAmendmentsRequest object
// with the default values initialized.
func (c *Client) ListAccountAmendments() *ListAccountAmendmentsRequest {
	var (
		pageNumberDefault = string("0")
		pageSizeDefault   = int64(100)
	)
	return &ListAccountAmendmentsRequest{

		FilterAccountID: c.Defaults.GetStrfmtUUIDPtr("ListAccountAmendments", "filter[account_id]"),

		FilterOrganisationID: c.Defaults.GetStrfmtUUIDPtr("ListAccountAmendments", "filter[organisation_id]"),

		FilterSubmissionStatus: c.Defaults.GetStringPtr("ListAccountAmendments", "filter[submission.status]"),

		FilterSubmissionSubmissionDateFrom: c.Defaults.GetStrfmtDateTimePtr("ListAccountAmendments", "filter[submission.submission_date_from]"),

		FilterSubmissionSubmissionDateTo: c.Defaults.GetStrfmtDateTimePtr("ListAccountAmendments", "filter[submission.submission_date_to]"),

		PageNumber: &pageNumberDefault,

		PageSize: &pageSizeDefault,

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ListAccountAmendmentsRequest struct {

	/*FilterAccountID      Filter by account ID      */

	FilterAccountID *strfmt.UUID

	/*FilterOrganisationID      Filter by organisationID      */

	FilterOrganisationID *strfmt.UUID

	/*FilterSubmissionStatus      Filter account request submission status      */

	FilterSubmissionStatus *string

	/*FilterSubmissionSubmissionDateFrom      Filter account amendments submission by date from      */

	FilterSubmissionSubmissionDateFrom *strfmt.DateTime

	/*FilterSubmissionSubmissionDateTo      Filter account amendments submission by date to      */

	FilterSubmissionSubmissionDateTo *strfmt.DateTime

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

func (o *ListAccountAmendmentsRequest) FromJson(j string) (*ListAccountAmendmentsRequest, error) {

	return o, nil
}

func (o *ListAccountAmendmentsRequest) WithFilterAccountID(filterAccountID strfmt.UUID) *ListAccountAmendmentsRequest {

	o.FilterAccountID = &filterAccountID

	return o
}

func (o *ListAccountAmendmentsRequest) WithoutFilterAccountID() *ListAccountAmendmentsRequest {

	o.FilterAccountID = nil

	return o
}

func (o *ListAccountAmendmentsRequest) WithFilterOrganisationID(filterOrganisationID strfmt.UUID) *ListAccountAmendmentsRequest {

	o.FilterOrganisationID = &filterOrganisationID

	return o
}

func (o *ListAccountAmendmentsRequest) WithoutFilterOrganisationID() *ListAccountAmendmentsRequest {

	o.FilterOrganisationID = nil

	return o
}

func (o *ListAccountAmendmentsRequest) WithFilterSubmissionStatus(filterSubmissionStatus string) *ListAccountAmendmentsRequest {

	o.FilterSubmissionStatus = &filterSubmissionStatus

	return o
}

func (o *ListAccountAmendmentsRequest) WithoutFilterSubmissionStatus() *ListAccountAmendmentsRequest {

	o.FilterSubmissionStatus = nil

	return o
}

func (o *ListAccountAmendmentsRequest) WithFilterSubmissionSubmissionDateFrom(filterSubmissionSubmissionDateFrom strfmt.DateTime) *ListAccountAmendmentsRequest {

	o.FilterSubmissionSubmissionDateFrom = &filterSubmissionSubmissionDateFrom

	return o
}

func (o *ListAccountAmendmentsRequest) WithoutFilterSubmissionSubmissionDateFrom() *ListAccountAmendmentsRequest {

	o.FilterSubmissionSubmissionDateFrom = nil

	return o
}

func (o *ListAccountAmendmentsRequest) WithFilterSubmissionSubmissionDateTo(filterSubmissionSubmissionDateTo strfmt.DateTime) *ListAccountAmendmentsRequest {

	o.FilterSubmissionSubmissionDateTo = &filterSubmissionSubmissionDateTo

	return o
}

func (o *ListAccountAmendmentsRequest) WithoutFilterSubmissionSubmissionDateTo() *ListAccountAmendmentsRequest {

	o.FilterSubmissionSubmissionDateTo = nil

	return o
}

func (o *ListAccountAmendmentsRequest) WithPageNumber(pageNumber string) *ListAccountAmendmentsRequest {

	o.PageNumber = &pageNumber

	return o
}

func (o *ListAccountAmendmentsRequest) WithoutPageNumber() *ListAccountAmendmentsRequest {

	o.PageNumber = nil

	return o
}

func (o *ListAccountAmendmentsRequest) WithPageSize(pageSize int64) *ListAccountAmendmentsRequest {

	o.PageSize = &pageSize

	return o
}

func (o *ListAccountAmendmentsRequest) WithoutPageSize() *ListAccountAmendmentsRequest {

	o.PageSize = nil

	return o
}

//////////////////
// WithContext adds the context to the list account amendments Request
func (o *ListAccountAmendmentsRequest) WithContext(ctx context.Context) *ListAccountAmendmentsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the list account amendments Request
func (o *ListAccountAmendmentsRequest) WithHTTPClient(client *http.Client) *ListAccountAmendmentsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ListAccountAmendmentsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterAccountID != nil {

		// query param filter[account_id]
		var qrFilterAccountID strfmt.UUID
		if o.FilterAccountID != nil {
			qrFilterAccountID = *o.FilterAccountID
		}
		qFilterAccountID := qrFilterAccountID.String()
		if qFilterAccountID != "" {
			if err := r.SetQueryParam("filter[account_id]", qFilterAccountID); err != nil {
				return err
			}
		}

	}

	if o.FilterOrganisationID != nil {

		// query param filter[organisation_id]
		var qrFilterOrganisationID strfmt.UUID
		if o.FilterOrganisationID != nil {
			qrFilterOrganisationID = *o.FilterOrganisationID
		}
		qFilterOrganisationID := qrFilterOrganisationID.String()
		if qFilterOrganisationID != "" {
			if err := r.SetQueryParam("filter[organisation_id]", qFilterOrganisationID); err != nil {
				return err
			}
		}

	}

	if o.FilterSubmissionStatus != nil {

		// query param filter[submission.status]
		var qrFilterSubmissionStatus string
		if o.FilterSubmissionStatus != nil {
			qrFilterSubmissionStatus = *o.FilterSubmissionStatus
		}
		qFilterSubmissionStatus := qrFilterSubmissionStatus
		if qFilterSubmissionStatus != "" {
			if err := r.SetQueryParam("filter[submission.status]", qFilterSubmissionStatus); err != nil {
				return err
			}
		}

	}

	if o.FilterSubmissionSubmissionDateFrom != nil {

		// query param filter[submission.submission_date_from]
		var qrFilterSubmissionSubmissionDateFrom strfmt.DateTime
		if o.FilterSubmissionSubmissionDateFrom != nil {
			qrFilterSubmissionSubmissionDateFrom = *o.FilterSubmissionSubmissionDateFrom
		}
		qFilterSubmissionSubmissionDateFrom := qrFilterSubmissionSubmissionDateFrom.String()
		if qFilterSubmissionSubmissionDateFrom != "" {
			if err := r.SetQueryParam("filter[submission.submission_date_from]", qFilterSubmissionSubmissionDateFrom); err != nil {
				return err
			}
		}

	}

	if o.FilterSubmissionSubmissionDateTo != nil {

		// query param filter[submission.submission_date_to]
		var qrFilterSubmissionSubmissionDateTo strfmt.DateTime
		if o.FilterSubmissionSubmissionDateTo != nil {
			qrFilterSubmissionSubmissionDateTo = *o.FilterSubmissionSubmissionDateTo
		}
		qFilterSubmissionSubmissionDateTo := qrFilterSubmissionSubmissionDateTo.String()
		if qFilterSubmissionSubmissionDateTo != "" {
			if err := r.SetQueryParam("filter[submission.submission_date_to]", qFilterSubmissionSubmissionDateTo); err != nil {
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
