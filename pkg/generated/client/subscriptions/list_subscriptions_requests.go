// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

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

// Client.ListSubscriptions creates a new ListSubscriptionsRequest object
// with the default values initialized.
func (c *Client) ListSubscriptions() *ListSubscriptionsRequest {
	var ()
	return &ListSubscriptionsRequest{

		FilterCallbackTransport: c.Defaults.GetStringPtr("ListSubscriptions", "filter[callback_transport]"),

		FilterCallbackURISearchTerm: c.Defaults.GetStringPtr("ListSubscriptions", "filter[callback_uri_search_term]"),

		FilterDeactivated: c.Defaults.GetBoolPtr("ListSubscriptions", "filter[deactivated]"),

		FilterEventType: make([]string, 0),

		FilterNotificationFilter: c.Defaults.GetBoolPtr("ListSubscriptions", "filter[notification_filter]"),

		FilterOrganisationID: make([]strfmt.UUID, 0),

		FilterRecordType: make([]string, 0),

		PageNumber: c.Defaults.GetStringPtr("ListSubscriptions", "page[number]"),

		PageSize: c.Defaults.GetInt64Ptr("ListSubscriptions", "page[size]"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ListSubscriptionsRequest struct {

	/*FilterCallbackTransport      Filter by callback_transport      */

	FilterCallbackTransport *string

	/*FilterCallbackURISearchTerm      Filter on callback_uri containing a search term      */

	FilterCallbackURISearchTerm *string

	/*FilterDeactivated      Filter by deactivated      */

	FilterDeactivated *bool

	/*FilterEventType      Filter by event type      */

	FilterEventType []string

	/*FilterNotificationFilter      Filter by existence of notification filters      */

	FilterNotificationFilter *bool

	/*FilterOrganisationID      Filter by organisation id      */

	FilterOrganisationID []strfmt.UUID

	/*FilterRecordType      Filter by record type      */

	FilterRecordType []string

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

func (o *ListSubscriptionsRequest) FromJson(j string) *ListSubscriptionsRequest {

	return o
}

func (o *ListSubscriptionsRequest) WithFilterCallbackTransport(filterCallbackTransport string) *ListSubscriptionsRequest {

	o.FilterCallbackTransport = &filterCallbackTransport

	return o
}

func (o *ListSubscriptionsRequest) WithoutFilterCallbackTransport() *ListSubscriptionsRequest {

	o.FilterCallbackTransport = nil

	return o
}

func (o *ListSubscriptionsRequest) WithFilterCallbackURISearchTerm(filterCallbackURISearchTerm string) *ListSubscriptionsRequest {

	o.FilterCallbackURISearchTerm = &filterCallbackURISearchTerm

	return o
}

func (o *ListSubscriptionsRequest) WithoutFilterCallbackURISearchTerm() *ListSubscriptionsRequest {

	o.FilterCallbackURISearchTerm = nil

	return o
}

func (o *ListSubscriptionsRequest) WithFilterDeactivated(filterDeactivated bool) *ListSubscriptionsRequest {

	o.FilterDeactivated = &filterDeactivated

	return o
}

func (o *ListSubscriptionsRequest) WithoutFilterDeactivated() *ListSubscriptionsRequest {

	o.FilterDeactivated = nil

	return o
}

func (o *ListSubscriptionsRequest) WithFilterEventType(filterEventType []string) *ListSubscriptionsRequest {

	o.FilterEventType = filterEventType

	return o
}

func (o *ListSubscriptionsRequest) WithoutFilterEventType() *ListSubscriptionsRequest {

	o.FilterEventType = nil

	return o
}

func (o *ListSubscriptionsRequest) WithFilterNotificationFilter(filterNotificationFilter bool) *ListSubscriptionsRequest {

	o.FilterNotificationFilter = &filterNotificationFilter

	return o
}

func (o *ListSubscriptionsRequest) WithoutFilterNotificationFilter() *ListSubscriptionsRequest {

	o.FilterNotificationFilter = nil

	return o
}

func (o *ListSubscriptionsRequest) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *ListSubscriptionsRequest {

	o.FilterOrganisationID = filterOrganisationID

	return o
}

func (o *ListSubscriptionsRequest) WithoutFilterOrganisationID() *ListSubscriptionsRequest {

	o.FilterOrganisationID = nil

	return o
}

func (o *ListSubscriptionsRequest) WithFilterRecordType(filterRecordType []string) *ListSubscriptionsRequest {

	o.FilterRecordType = filterRecordType

	return o
}

func (o *ListSubscriptionsRequest) WithoutFilterRecordType() *ListSubscriptionsRequest {

	o.FilterRecordType = nil

	return o
}

func (o *ListSubscriptionsRequest) WithPageNumber(pageNumber string) *ListSubscriptionsRequest {

	o.PageNumber = &pageNumber

	return o
}

func (o *ListSubscriptionsRequest) WithoutPageNumber() *ListSubscriptionsRequest {

	o.PageNumber = nil

	return o
}

func (o *ListSubscriptionsRequest) WithPageSize(pageSize int64) *ListSubscriptionsRequest {

	o.PageSize = &pageSize

	return o
}

func (o *ListSubscriptionsRequest) WithoutPageSize() *ListSubscriptionsRequest {

	o.PageSize = nil

	return o
}

//////////////////
// WithContext adds the context to the list subscriptions Request
func (o *ListSubscriptionsRequest) WithContext(ctx context.Context) *ListSubscriptionsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the list subscriptions Request
func (o *ListSubscriptionsRequest) WithHTTPClient(client *http.Client) *ListSubscriptionsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ListSubscriptionsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FilterCallbackTransport != nil {

		// query param filter[callback_transport]
		var qrFilterCallbackTransport string
		if o.FilterCallbackTransport != nil {
			qrFilterCallbackTransport = *o.FilterCallbackTransport
		}
		qFilterCallbackTransport := qrFilterCallbackTransport
		if qFilterCallbackTransport != "" {
			if err := r.SetQueryParam("filter[callback_transport]", qFilterCallbackTransport); err != nil {
				return err
			}
		}

	}

	if o.FilterCallbackURISearchTerm != nil {

		// query param filter[callback_uri_search_term]
		var qrFilterCallbackURISearchTerm string
		if o.FilterCallbackURISearchTerm != nil {
			qrFilterCallbackURISearchTerm = *o.FilterCallbackURISearchTerm
		}
		qFilterCallbackURISearchTerm := qrFilterCallbackURISearchTerm
		if qFilterCallbackURISearchTerm != "" {
			if err := r.SetQueryParam("filter[callback_uri_search_term]", qFilterCallbackURISearchTerm); err != nil {
				return err
			}
		}

	}

	if o.FilterDeactivated != nil {

		// query param filter[deactivated]
		var qrFilterDeactivated bool
		if o.FilterDeactivated != nil {
			qrFilterDeactivated = *o.FilterDeactivated
		}
		qFilterDeactivated := swag.FormatBool(qrFilterDeactivated)
		if qFilterDeactivated != "" {
			if err := r.SetQueryParam("filter[deactivated]", qFilterDeactivated); err != nil {
				return err
			}
		}

	}

	valuesFilterEventType := o.FilterEventType

	joinedFilterEventType := swag.JoinByFormat(valuesFilterEventType, "csv")
	// query array param filter[event_type]
	if err := r.SetQueryParam("filter[event_type]", joinedFilterEventType...); err != nil {
		return err
	}

	if o.FilterNotificationFilter != nil {

		// query param filter[notification_filter]
		var qrFilterNotificationFilter bool
		if o.FilterNotificationFilter != nil {
			qrFilterNotificationFilter = *o.FilterNotificationFilter
		}
		qFilterNotificationFilter := swag.FormatBool(qrFilterNotificationFilter)
		if qFilterNotificationFilter != "" {
			if err := r.SetQueryParam("filter[notification_filter]", qFilterNotificationFilter); err != nil {
				return err
			}
		}

	}

	var valuesFilterOrganisationID []string
	for _, v := range o.FilterOrganisationID {
		valuesFilterOrganisationID = append(valuesFilterOrganisationID, v.String())
	}

	joinedFilterOrganisationID := swag.JoinByFormat(valuesFilterOrganisationID, "csv")
	// query array param filter[organisation_id]
	if err := r.SetQueryParam("filter[organisation_id]", joinedFilterOrganisationID...); err != nil {
		return err
	}

	valuesFilterRecordType := o.FilterRecordType

	joinedFilterRecordType := swag.JoinByFormat(valuesFilterRecordType, "csv")
	// query array param filter[record_type]
	if err := r.SetQueryParam("filter[record_type]", joinedFilterRecordType...); err != nil {
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
