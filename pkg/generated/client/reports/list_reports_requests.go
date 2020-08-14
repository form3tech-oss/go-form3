// Code generated by go-swagger; DO NOT EDIT.

package reports

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

// Client.ListReports creates a new ListReportsRequest object
// with the default values initialized.
func (c *Client) ListReports() *ListReportsRequest {
	var (
		pageSizeDefault = int64(100)
	)
	return &ListReportsRequest{

		FilterCreatedOnAfter: c.Defaults.GetStrfmtDateTimePtr("ListReports", "filter[created_on_after]"),

		FilterCreatedOnBefore: c.Defaults.GetStrfmtDateTimePtr("ListReports", "filter[created_on_before]"),

		FilterModifiedOnAfter: c.Defaults.GetStrfmtDateTimePtr("ListReports", "filter[modified_on_after]"),

		FilterModifiedOnBefore: c.Defaults.GetStrfmtDateTimePtr("ListReports", "filter[modified_on_before]"),

		FilterOrganisationID: make([]strfmt.UUID, 0),

		FilterReportSource: c.Defaults.GetStringPtr("ListReports", "filter[report_source]"),

		FilterReportType: c.Defaults.GetStringPtr("ListReports", "filter[report_type]"),

		FilterReportTypeDescription: c.Defaults.GetStringPtr("ListReports", "filter[report_type_description]"),

		PageNumber: c.Defaults.GetStringPtr("ListReports", "page[number]"),

		PageSize: &pageSizeDefault,

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ListReportsRequest struct {

	/*FilterCreatedOnAfter      Request reports created after specific date time      */

	FilterCreatedOnAfter *strfmt.DateTime

	/*FilterCreatedOnBefore      Request reports created after specific date time      */

	FilterCreatedOnBefore *strfmt.DateTime

	/*FilterModifiedOnAfter      Request reports modified after specific date time      */

	FilterModifiedOnAfter *strfmt.DateTime

	/*FilterModifiedOnBefore      Request reports modified before specific date time      */

	FilterModifiedOnBefore *strfmt.DateTime

	/*FilterOrganisationID      Filter by organisation Ids      */

	FilterOrganisationID []strfmt.UUID

	/*FilterReportSource      Filter by Report Source      */

	FilterReportSource *string

	/*FilterReportType      Filter bu ReportType      */

	FilterReportType *string

	/*FilterReportTypeDescription      Filter by Report Type Description      */

	FilterReportTypeDescription *string

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

func (o *ListReportsRequest) FromJson(j string) *ListReportsRequest {

	return o
}

func (o *ListReportsRequest) WithFilterCreatedOnAfter(filterCreatedOnAfter strfmt.DateTime) *ListReportsRequest {

	o.FilterCreatedOnAfter = &filterCreatedOnAfter

	return o
}

func (o *ListReportsRequest) WithoutFilterCreatedOnAfter() *ListReportsRequest {

	o.FilterCreatedOnAfter = nil

	return o
}

func (o *ListReportsRequest) WithFilterCreatedOnBefore(filterCreatedOnBefore strfmt.DateTime) *ListReportsRequest {

	o.FilterCreatedOnBefore = &filterCreatedOnBefore

	return o
}

func (o *ListReportsRequest) WithoutFilterCreatedOnBefore() *ListReportsRequest {

	o.FilterCreatedOnBefore = nil

	return o
}

func (o *ListReportsRequest) WithFilterModifiedOnAfter(filterModifiedOnAfter strfmt.DateTime) *ListReportsRequest {

	o.FilterModifiedOnAfter = &filterModifiedOnAfter

	return o
}

func (o *ListReportsRequest) WithoutFilterModifiedOnAfter() *ListReportsRequest {

	o.FilterModifiedOnAfter = nil

	return o
}

func (o *ListReportsRequest) WithFilterModifiedOnBefore(filterModifiedOnBefore strfmt.DateTime) *ListReportsRequest {

	o.FilterModifiedOnBefore = &filterModifiedOnBefore

	return o
}

func (o *ListReportsRequest) WithoutFilterModifiedOnBefore() *ListReportsRequest {

	o.FilterModifiedOnBefore = nil

	return o
}

func (o *ListReportsRequest) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *ListReportsRequest {

	o.FilterOrganisationID = filterOrganisationID

	return o
}

func (o *ListReportsRequest) WithoutFilterOrganisationID() *ListReportsRequest {

	o.FilterOrganisationID = nil

	return o
}

func (o *ListReportsRequest) WithFilterReportSource(filterReportSource string) *ListReportsRequest {

	o.FilterReportSource = &filterReportSource

	return o
}

func (o *ListReportsRequest) WithoutFilterReportSource() *ListReportsRequest {

	o.FilterReportSource = nil

	return o
}

func (o *ListReportsRequest) WithFilterReportType(filterReportType string) *ListReportsRequest {

	o.FilterReportType = &filterReportType

	return o
}

func (o *ListReportsRequest) WithoutFilterReportType() *ListReportsRequest {

	o.FilterReportType = nil

	return o
}

func (o *ListReportsRequest) WithFilterReportTypeDescription(filterReportTypeDescription string) *ListReportsRequest {

	o.FilterReportTypeDescription = &filterReportTypeDescription

	return o
}

func (o *ListReportsRequest) WithoutFilterReportTypeDescription() *ListReportsRequest {

	o.FilterReportTypeDescription = nil

	return o
}

func (o *ListReportsRequest) WithPageNumber(pageNumber string) *ListReportsRequest {

	o.PageNumber = &pageNumber

	return o
}

func (o *ListReportsRequest) WithoutPageNumber() *ListReportsRequest {

	o.PageNumber = nil

	return o
}

func (o *ListReportsRequest) WithPageSize(pageSize int64) *ListReportsRequest {

	o.PageSize = &pageSize

	return o
}

func (o *ListReportsRequest) WithoutPageSize() *ListReportsRequest {

	o.PageSize = nil

	return o
}

//////////////////
// WithContext adds the context to the list reports Request
func (o *ListReportsRequest) WithContext(ctx context.Context) *ListReportsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the list reports Request
func (o *ListReportsRequest) WithHTTPClient(client *http.Client) *ListReportsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ListReportsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterCreatedOnAfter != nil {

		// query param filter[created_on_after]
		var qrFilterCreatedOnAfter strfmt.DateTime
		if o.FilterCreatedOnAfter != nil {
			qrFilterCreatedOnAfter = *o.FilterCreatedOnAfter
		}
		qFilterCreatedOnAfter := qrFilterCreatedOnAfter.String()
		if qFilterCreatedOnAfter != "" {
			if err := r.SetQueryParam("filter[created_on_after]", qFilterCreatedOnAfter); err != nil {
				return err
			}
		}

	}

	if o.FilterCreatedOnBefore != nil {

		// query param filter[created_on_before]
		var qrFilterCreatedOnBefore strfmt.DateTime
		if o.FilterCreatedOnBefore != nil {
			qrFilterCreatedOnBefore = *o.FilterCreatedOnBefore
		}
		qFilterCreatedOnBefore := qrFilterCreatedOnBefore.String()
		if qFilterCreatedOnBefore != "" {
			if err := r.SetQueryParam("filter[created_on_before]", qFilterCreatedOnBefore); err != nil {
				return err
			}
		}

	}

	if o.FilterModifiedOnAfter != nil {

		// query param filter[modified_on_after]
		var qrFilterModifiedOnAfter strfmt.DateTime
		if o.FilterModifiedOnAfter != nil {
			qrFilterModifiedOnAfter = *o.FilterModifiedOnAfter
		}
		qFilterModifiedOnAfter := qrFilterModifiedOnAfter.String()
		if qFilterModifiedOnAfter != "" {
			if err := r.SetQueryParam("filter[modified_on_after]", qFilterModifiedOnAfter); err != nil {
				return err
			}
		}

	}

	if o.FilterModifiedOnBefore != nil {

		// query param filter[modified_on_before]
		var qrFilterModifiedOnBefore strfmt.DateTime
		if o.FilterModifiedOnBefore != nil {
			qrFilterModifiedOnBefore = *o.FilterModifiedOnBefore
		}
		qFilterModifiedOnBefore := qrFilterModifiedOnBefore.String()
		if qFilterModifiedOnBefore != "" {
			if err := r.SetQueryParam("filter[modified_on_before]", qFilterModifiedOnBefore); err != nil {
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

	if o.FilterReportSource != nil {

		// query param filter[report_source]
		var qrFilterReportSource string
		if o.FilterReportSource != nil {
			qrFilterReportSource = *o.FilterReportSource
		}
		qFilterReportSource := qrFilterReportSource
		if qFilterReportSource != "" {
			if err := r.SetQueryParam("filter[report_source]", qFilterReportSource); err != nil {
				return err
			}
		}

	}

	if o.FilterReportType != nil {

		// query param filter[report_type]
		var qrFilterReportType string
		if o.FilterReportType != nil {
			qrFilterReportType = *o.FilterReportType
		}
		qFilterReportType := qrFilterReportType
		if qFilterReportType != "" {
			if err := r.SetQueryParam("filter[report_type]", qFilterReportType); err != nil {
				return err
			}
		}

	}

	if o.FilterReportTypeDescription != nil {

		// query param filter[report_type_description]
		var qrFilterReportTypeDescription string
		if o.FilterReportTypeDescription != nil {
			qrFilterReportTypeDescription = *o.FilterReportTypeDescription
		}
		qFilterReportTypeDescription := qrFilterReportTypeDescription
		if qFilterReportTypeDescription != "" {
			if err := r.SetQueryParam("filter[report_type_description]", qFilterReportTypeDescription); err != nil {
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
