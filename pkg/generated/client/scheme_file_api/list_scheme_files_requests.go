// Code generated by go-swagger; DO NOT EDIT.

package scheme_file_api

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

// Client.ListSchemeFiles creates a new ListSchemeFilesRequest object
// with the default values initialized.
func (c *Client) ListSchemeFiles() *ListSchemeFilesRequest {
	var (
		pageSizeDefault = int64(100)
	)
	return &ListSchemeFilesRequest{

		FilterCreatedOnFrom: c.Defaults.GetStrfmtDatePtr("ListSchemeFiles", "filter[created_on_from]"),

		FilterCreatedOnTo: c.Defaults.GetStrfmtDatePtr("ListSchemeFiles", "filter[created_on_to]"),

		FilterFileFormat: c.Defaults.GetStringPtr("ListSchemeFiles", "filter[file_format]"),

		FilterFileType: c.Defaults.GetStringPtr("ListSchemeFiles", "filter[file_type]"),

		FilterOrganisationID: make([]strfmt.UUID, 0),

		FilterPaymentScheme: c.Defaults.GetStringPtr("ListSchemeFiles", "filter[payment_scheme]"),

		FilterSubmissionStatus: c.Defaults.GetStringPtr("ListSchemeFiles", "filter[submission.status]"),

		FilterSubmissionSubmissionDateFrom: c.Defaults.GetStrfmtDateTimePtr("ListSchemeFiles", "filter[submission.submission_date_from]"),

		FilterSubmissionSubmissionDateTo: c.Defaults.GetStrfmtDateTimePtr("ListSchemeFiles", "filter[submission.submission_date_to]"),

		PageNumber: c.Defaults.GetStringPtr("ListSchemeFiles", "page[number]"),

		PageSize: &pageSizeDefault,

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ListSchemeFilesRequest struct {

	/*FilterCreatedOnFrom      Find all Scheme File resources created from this date, in format YYYY-MM-DD      */

	FilterCreatedOnFrom *strfmt.Date

	/*FilterCreatedOnTo      Find all Scheme File resources created up to this date, in format YYYY-MM-DD      */

	FilterCreatedOnTo *strfmt.Date

	/*FilterFileFormat      Find Scheme File resources by the format of the file      */

	FilterFileFormat *string

	/*FilterFileType      Find Scheme File resources by the type of the file      */

	FilterFileType *string

	/*FilterOrganisationID      Find all Scheme File resources with a given organisation ID      */

	FilterOrganisationID []strfmt.UUID

	/*FilterPaymentScheme      Find Scheme File resources by a certain scheme      */

	FilterPaymentScheme *string

	/*FilterSubmissionStatus      Find all Scheme File resources with a certain submission status      */

	FilterSubmissionStatus *string

	/*FilterSubmissionSubmissionDateFrom      Find all Scheme File resources submitted from and including this date/time      */

	FilterSubmissionSubmissionDateFrom *strfmt.DateTime

	/*FilterSubmissionSubmissionDateTo      Find all Scheme File resources submitted up to and included this date/time      */

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

func (o *ListSchemeFilesRequest) FromJson(j string) (*ListSchemeFilesRequest, error) {

	return o, nil
}

func (o *ListSchemeFilesRequest) WithFilterCreatedOnFrom(filterCreatedOnFrom strfmt.Date) *ListSchemeFilesRequest {

	o.FilterCreatedOnFrom = &filterCreatedOnFrom

	return o
}

func (o *ListSchemeFilesRequest) WithoutFilterCreatedOnFrom() *ListSchemeFilesRequest {

	o.FilterCreatedOnFrom = nil

	return o
}

func (o *ListSchemeFilesRequest) WithFilterCreatedOnTo(filterCreatedOnTo strfmt.Date) *ListSchemeFilesRequest {

	o.FilterCreatedOnTo = &filterCreatedOnTo

	return o
}

func (o *ListSchemeFilesRequest) WithoutFilterCreatedOnTo() *ListSchemeFilesRequest {

	o.FilterCreatedOnTo = nil

	return o
}

func (o *ListSchemeFilesRequest) WithFilterFileFormat(filterFileFormat string) *ListSchemeFilesRequest {

	o.FilterFileFormat = &filterFileFormat

	return o
}

func (o *ListSchemeFilesRequest) WithoutFilterFileFormat() *ListSchemeFilesRequest {

	o.FilterFileFormat = nil

	return o
}

func (o *ListSchemeFilesRequest) WithFilterFileType(filterFileType string) *ListSchemeFilesRequest {

	o.FilterFileType = &filterFileType

	return o
}

func (o *ListSchemeFilesRequest) WithoutFilterFileType() *ListSchemeFilesRequest {

	o.FilterFileType = nil

	return o
}

func (o *ListSchemeFilesRequest) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *ListSchemeFilesRequest {

	o.FilterOrganisationID = filterOrganisationID

	return o
}

func (o *ListSchemeFilesRequest) WithoutFilterOrganisationID() *ListSchemeFilesRequest {

	o.FilterOrganisationID = nil

	return o
}

func (o *ListSchemeFilesRequest) WithFilterPaymentScheme(filterPaymentScheme string) *ListSchemeFilesRequest {

	o.FilterPaymentScheme = &filterPaymentScheme

	return o
}

func (o *ListSchemeFilesRequest) WithoutFilterPaymentScheme() *ListSchemeFilesRequest {

	o.FilterPaymentScheme = nil

	return o
}

func (o *ListSchemeFilesRequest) WithFilterSubmissionStatus(filterSubmissionStatus string) *ListSchemeFilesRequest {

	o.FilterSubmissionStatus = &filterSubmissionStatus

	return o
}

func (o *ListSchemeFilesRequest) WithoutFilterSubmissionStatus() *ListSchemeFilesRequest {

	o.FilterSubmissionStatus = nil

	return o
}

func (o *ListSchemeFilesRequest) WithFilterSubmissionSubmissionDateFrom(filterSubmissionSubmissionDateFrom strfmt.DateTime) *ListSchemeFilesRequest {

	o.FilterSubmissionSubmissionDateFrom = &filterSubmissionSubmissionDateFrom

	return o
}

func (o *ListSchemeFilesRequest) WithoutFilterSubmissionSubmissionDateFrom() *ListSchemeFilesRequest {

	o.FilterSubmissionSubmissionDateFrom = nil

	return o
}

func (o *ListSchemeFilesRequest) WithFilterSubmissionSubmissionDateTo(filterSubmissionSubmissionDateTo strfmt.DateTime) *ListSchemeFilesRequest {

	o.FilterSubmissionSubmissionDateTo = &filterSubmissionSubmissionDateTo

	return o
}

func (o *ListSchemeFilesRequest) WithoutFilterSubmissionSubmissionDateTo() *ListSchemeFilesRequest {

	o.FilterSubmissionSubmissionDateTo = nil

	return o
}

func (o *ListSchemeFilesRequest) WithPageNumber(pageNumber string) *ListSchemeFilesRequest {

	o.PageNumber = &pageNumber

	return o
}

func (o *ListSchemeFilesRequest) WithoutPageNumber() *ListSchemeFilesRequest {

	o.PageNumber = nil

	return o
}

func (o *ListSchemeFilesRequest) WithPageSize(pageSize int64) *ListSchemeFilesRequest {

	o.PageSize = &pageSize

	return o
}

func (o *ListSchemeFilesRequest) WithoutPageSize() *ListSchemeFilesRequest {

	o.PageSize = nil

	return o
}

//////////////////
// WithContext adds the context to the list scheme files Request
func (o *ListSchemeFilesRequest) WithContext(ctx context.Context) *ListSchemeFilesRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the list scheme files Request
func (o *ListSchemeFilesRequest) WithHTTPClient(client *http.Client) *ListSchemeFilesRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ListSchemeFilesRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterCreatedOnFrom != nil {

		// query param filter[created_on_from]
		var qrFilterCreatedOnFrom strfmt.Date
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
		var qrFilterCreatedOnTo strfmt.Date
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

	if o.FilterFileFormat != nil {

		// query param filter[file_format]
		var qrFilterFileFormat string
		if o.FilterFileFormat != nil {
			qrFilterFileFormat = *o.FilterFileFormat
		}
		qFilterFileFormat := qrFilterFileFormat
		if qFilterFileFormat != "" {
			if err := r.SetQueryParam("filter[file_format]", qFilterFileFormat); err != nil {
				return err
			}
		}

	}

	if o.FilterFileType != nil {

		// query param filter[file_type]
		var qrFilterFileType string
		if o.FilterFileType != nil {
			qrFilterFileType = *o.FilterFileType
		}
		qFilterFileType := qrFilterFileType
		if qFilterFileType != "" {
			if err := r.SetQueryParam("filter[file_type]", qFilterFileType); err != nil {
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
