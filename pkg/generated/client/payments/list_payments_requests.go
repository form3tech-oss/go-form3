// Code generated by go-swagger; DO NOT EDIT.

package payments

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

// Client.ListPayments creates a new ListPaymentsRequest object
// with the default values initialized.
func (c *Client) ListPayments() *ListPaymentsRequest {
	var ()
	return &ListPaymentsRequest{

		FilterAdmissionAdmissionDateFrom: c.Defaults.GetStrfmtDateTimePtr("ListPayments", "filter[admission.admission_date_from]"),

		FilterAdmissionAdmissionDateTo: c.Defaults.GetStrfmtDateTimePtr("ListPayments", "filter[admission.admission_date_to]"),

		FilterAdmissionSchemeStatusCode: c.Defaults.GetStringPtr("ListPayments", "filter[admission.scheme_status_code]"),

		FilterAdmissionStatus: c.Defaults.GetStringPtr("ListPayments", "filter[admission.status]"),

		FilterAmount: c.Defaults.GetStringPtr("ListPayments", "filter[amount]"),

		FilterBeneficiaryPartyAccountName: c.Defaults.GetStringPtr("ListPayments", "filter[beneficiary_party.account_name]"),

		FilterBeneficiaryPartyAccountNumber: c.Defaults.GetStringPtr("ListPayments", "filter[beneficiary_party.account_number]"),

		FilterBeneficiaryPartyBankID: c.Defaults.GetStringPtr("ListPayments", "filter[beneficiary_party.bank_id]"),

		FilterCurrency: c.Defaults.GetStringPtr("ListPayments", "filter[currency]"),

		FilterDebtorPartyAccountName: c.Defaults.GetStringPtr("ListPayments", "filter[debtor_party.account_name]"),

		FilterDebtorPartyAccountNumber: c.Defaults.GetStringPtr("ListPayments", "filter[debtor_party.account_number]"),

		FilterDebtorPartyBankID: c.Defaults.GetStringPtr("ListPayments", "filter[debtor_party.bank_id]"),

		FilterEndToEndReference: c.Defaults.GetStringPtr("ListPayments", "filter[end_to_end_reference]"),

		FilterNotRelationships: make([]string, 0),

		FilterOrganisationID: make([]strfmt.UUID, 0),

		FilterPaymentScheme: c.Defaults.GetStringPtr("ListPayments", "filter[payment_scheme]"),

		FilterPaymentType: c.Defaults.GetStringPtr("ListPayments", "filter[payment_type]"),

		FilterProcessingDateFrom: c.Defaults.GetStrfmtDatePtr("ListPayments", "filter[processing_date_from]"),

		FilterProcessingDateTo: c.Defaults.GetStrfmtDatePtr("ListPayments", "filter[processing_date_to]"),

		FilterReference: c.Defaults.GetStringPtr("ListPayments", "filter[reference]"),

		FilterRelationships: make([]string, 0),

		FilterReturnUniqueSchemeID: c.Defaults.GetStringPtr("ListPayments", "filter[return.unique_scheme_id]"),

		FilterReturnSubmissionStatus: c.Defaults.GetStringPtr("ListPayments", "filter[return_submission.status]"),

		FilterReturnSubmissionSubmissionDateFrom: c.Defaults.GetStrfmtDateTimePtr("ListPayments", "filter[return_submission.submission_date_from]"),

		FilterReturnSubmissionSubmissionDateTo: c.Defaults.GetStrfmtDateTimePtr("ListPayments", "filter[return_submission.submission_date_to]"),

		FilterSchemeTransactionID: c.Defaults.GetStringPtr("ListPayments", "filter[scheme_transaction_id]"),

		FilterSubmissionSchemeStatusCode: c.Defaults.GetStringPtr("ListPayments", "filter[submission.scheme_status_code]"),

		FilterSubmissionStatus: c.Defaults.GetStringPtr("ListPayments", "filter[submission.status]"),

		FilterSubmissionSubmissionDateFrom: c.Defaults.GetStrfmtDateTimePtr("ListPayments", "filter[submission.submission_date_from]"),

		FilterSubmissionSubmissionDateTo: c.Defaults.GetStrfmtDateTimePtr("ListPayments", "filter[submission.submission_date_to]"),

		FilterUniqueSchemeID: c.Defaults.GetStringPtr("ListPayments", "filter[unique_scheme_id]"),

		PageNumber: c.Defaults.GetStringPtr("ListPayments", "page[number]"),

		PageSize: c.Defaults.GetInt64Ptr("ListPayments", "page[size]"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ListPaymentsRequest struct {

	/*FilterAdmissionAdmissionDateFrom*/

	FilterAdmissionAdmissionDateFrom *strfmt.DateTime

	/*FilterAdmissionAdmissionDateTo*/

	FilterAdmissionAdmissionDateTo *strfmt.DateTime

	/*FilterAdmissionSchemeStatusCode      Filter by admission scheme status code      */

	FilterAdmissionSchemeStatusCode *string

	/*FilterAdmissionStatus      Filter by admission status      */

	FilterAdmissionStatus *string

	/*FilterAmount*/

	FilterAmount *string

	/*FilterBeneficiaryPartyAccountName*/

	FilterBeneficiaryPartyAccountName *string

	/*FilterBeneficiaryPartyAccountNumber*/

	FilterBeneficiaryPartyAccountNumber *string

	/*FilterBeneficiaryPartyBankID*/

	FilterBeneficiaryPartyBankID *string

	/*FilterCurrency*/

	FilterCurrency *string

	/*FilterDebtorPartyAccountName*/

	FilterDebtorPartyAccountName *string

	/*FilterDebtorPartyAccountNumber*/

	FilterDebtorPartyAccountNumber *string

	/*FilterDebtorPartyBankID*/

	FilterDebtorPartyBankID *string

	/*FilterEndToEndReference*/

	FilterEndToEndReference *string

	/*FilterNotRelationships      Filter for payments containing none of the requested relationships      */

	FilterNotRelationships []string

	/*FilterOrganisationID      Filter by organisation id      */

	FilterOrganisationID []strfmt.UUID

	/*FilterPaymentScheme*/

	FilterPaymentScheme *string

	/*FilterPaymentType*/

	FilterPaymentType *string

	/*FilterProcessingDateFrom*/

	FilterProcessingDateFrom *strfmt.Date

	/*FilterProcessingDateTo*/

	FilterProcessingDateTo *strfmt.Date

	/*FilterReference*/

	FilterReference *string

	/*FilterRelationships      Filter for payments containing all of the requested relationships      */

	FilterRelationships []string

	/*FilterReturnUniqueSchemeID*/

	FilterReturnUniqueSchemeID *string

	/*FilterReturnSubmissionStatus      Filter by return submission status      */

	FilterReturnSubmissionStatus *string

	/*FilterReturnSubmissionSubmissionDateFrom*/

	FilterReturnSubmissionSubmissionDateFrom *strfmt.DateTime

	/*FilterReturnSubmissionSubmissionDateTo*/

	FilterReturnSubmissionSubmissionDateTo *strfmt.DateTime

	/*FilterSchemeTransactionID*/

	FilterSchemeTransactionID *string

	/*FilterSubmissionSchemeStatusCode      Filter by submission scheme status code      */

	FilterSubmissionSchemeStatusCode *string

	/*FilterSubmissionStatus      Filter by submission status      */

	FilterSubmissionStatus *string

	/*FilterSubmissionSubmissionDateFrom*/

	FilterSubmissionSubmissionDateFrom *strfmt.DateTime

	/*FilterSubmissionSubmissionDateTo*/

	FilterSubmissionSubmissionDateTo *strfmt.DateTime

	/*FilterUniqueSchemeID*/

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

func (o *ListPaymentsRequest) FromJson(j string) (*ListPaymentsRequest, error) {

	return o, nil
}

func (o *ListPaymentsRequest) WithFilterAdmissionAdmissionDateFrom(filterAdmissionAdmissionDateFrom strfmt.DateTime) *ListPaymentsRequest {

	o.FilterAdmissionAdmissionDateFrom = &filterAdmissionAdmissionDateFrom

	return o
}

func (o *ListPaymentsRequest) WithoutFilterAdmissionAdmissionDateFrom() *ListPaymentsRequest {

	o.FilterAdmissionAdmissionDateFrom = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterAdmissionAdmissionDateTo(filterAdmissionAdmissionDateTo strfmt.DateTime) *ListPaymentsRequest {

	o.FilterAdmissionAdmissionDateTo = &filterAdmissionAdmissionDateTo

	return o
}

func (o *ListPaymentsRequest) WithoutFilterAdmissionAdmissionDateTo() *ListPaymentsRequest {

	o.FilterAdmissionAdmissionDateTo = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterAdmissionSchemeStatusCode(filterAdmissionSchemeStatusCode string) *ListPaymentsRequest {

	o.FilterAdmissionSchemeStatusCode = &filterAdmissionSchemeStatusCode

	return o
}

func (o *ListPaymentsRequest) WithoutFilterAdmissionSchemeStatusCode() *ListPaymentsRequest {

	o.FilterAdmissionSchemeStatusCode = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterAdmissionStatus(filterAdmissionStatus string) *ListPaymentsRequest {

	o.FilterAdmissionStatus = &filterAdmissionStatus

	return o
}

func (o *ListPaymentsRequest) WithoutFilterAdmissionStatus() *ListPaymentsRequest {

	o.FilterAdmissionStatus = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterAmount(filterAmount string) *ListPaymentsRequest {

	o.FilterAmount = &filterAmount

	return o
}

func (o *ListPaymentsRequest) WithoutFilterAmount() *ListPaymentsRequest {

	o.FilterAmount = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterBeneficiaryPartyAccountName(filterBeneficiaryPartyAccountName string) *ListPaymentsRequest {

	o.FilterBeneficiaryPartyAccountName = &filterBeneficiaryPartyAccountName

	return o
}

func (o *ListPaymentsRequest) WithoutFilterBeneficiaryPartyAccountName() *ListPaymentsRequest {

	o.FilterBeneficiaryPartyAccountName = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterBeneficiaryPartyAccountNumber(filterBeneficiaryPartyAccountNumber string) *ListPaymentsRequest {

	o.FilterBeneficiaryPartyAccountNumber = &filterBeneficiaryPartyAccountNumber

	return o
}

func (o *ListPaymentsRequest) WithoutFilterBeneficiaryPartyAccountNumber() *ListPaymentsRequest {

	o.FilterBeneficiaryPartyAccountNumber = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterBeneficiaryPartyBankID(filterBeneficiaryPartyBankID string) *ListPaymentsRequest {

	o.FilterBeneficiaryPartyBankID = &filterBeneficiaryPartyBankID

	return o
}

func (o *ListPaymentsRequest) WithoutFilterBeneficiaryPartyBankID() *ListPaymentsRequest {

	o.FilterBeneficiaryPartyBankID = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterCurrency(filterCurrency string) *ListPaymentsRequest {

	o.FilterCurrency = &filterCurrency

	return o
}

func (o *ListPaymentsRequest) WithoutFilterCurrency() *ListPaymentsRequest {

	o.FilterCurrency = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterDebtorPartyAccountName(filterDebtorPartyAccountName string) *ListPaymentsRequest {

	o.FilterDebtorPartyAccountName = &filterDebtorPartyAccountName

	return o
}

func (o *ListPaymentsRequest) WithoutFilterDebtorPartyAccountName() *ListPaymentsRequest {

	o.FilterDebtorPartyAccountName = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterDebtorPartyAccountNumber(filterDebtorPartyAccountNumber string) *ListPaymentsRequest {

	o.FilterDebtorPartyAccountNumber = &filterDebtorPartyAccountNumber

	return o
}

func (o *ListPaymentsRequest) WithoutFilterDebtorPartyAccountNumber() *ListPaymentsRequest {

	o.FilterDebtorPartyAccountNumber = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterDebtorPartyBankID(filterDebtorPartyBankID string) *ListPaymentsRequest {

	o.FilterDebtorPartyBankID = &filterDebtorPartyBankID

	return o
}

func (o *ListPaymentsRequest) WithoutFilterDebtorPartyBankID() *ListPaymentsRequest {

	o.FilterDebtorPartyBankID = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterEndToEndReference(filterEndToEndReference string) *ListPaymentsRequest {

	o.FilterEndToEndReference = &filterEndToEndReference

	return o
}

func (o *ListPaymentsRequest) WithoutFilterEndToEndReference() *ListPaymentsRequest {

	o.FilterEndToEndReference = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterNotRelationships(filterNotRelationships []string) *ListPaymentsRequest {

	o.FilterNotRelationships = filterNotRelationships

	return o
}

func (o *ListPaymentsRequest) WithoutFilterNotRelationships() *ListPaymentsRequest {

	o.FilterNotRelationships = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *ListPaymentsRequest {

	o.FilterOrganisationID = filterOrganisationID

	return o
}

func (o *ListPaymentsRequest) WithoutFilterOrganisationID() *ListPaymentsRequest {

	o.FilterOrganisationID = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterPaymentScheme(filterPaymentScheme string) *ListPaymentsRequest {

	o.FilterPaymentScheme = &filterPaymentScheme

	return o
}

func (o *ListPaymentsRequest) WithoutFilterPaymentScheme() *ListPaymentsRequest {

	o.FilterPaymentScheme = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterPaymentType(filterPaymentType string) *ListPaymentsRequest {

	o.FilterPaymentType = &filterPaymentType

	return o
}

func (o *ListPaymentsRequest) WithoutFilterPaymentType() *ListPaymentsRequest {

	o.FilterPaymentType = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterProcessingDateFrom(filterProcessingDateFrom strfmt.Date) *ListPaymentsRequest {

	o.FilterProcessingDateFrom = &filterProcessingDateFrom

	return o
}

func (o *ListPaymentsRequest) WithoutFilterProcessingDateFrom() *ListPaymentsRequest {

	o.FilterProcessingDateFrom = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterProcessingDateTo(filterProcessingDateTo strfmt.Date) *ListPaymentsRequest {

	o.FilterProcessingDateTo = &filterProcessingDateTo

	return o
}

func (o *ListPaymentsRequest) WithoutFilterProcessingDateTo() *ListPaymentsRequest {

	o.FilterProcessingDateTo = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterReference(filterReference string) *ListPaymentsRequest {

	o.FilterReference = &filterReference

	return o
}

func (o *ListPaymentsRequest) WithoutFilterReference() *ListPaymentsRequest {

	o.FilterReference = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterRelationships(filterRelationships []string) *ListPaymentsRequest {

	o.FilterRelationships = filterRelationships

	return o
}

func (o *ListPaymentsRequest) WithoutFilterRelationships() *ListPaymentsRequest {

	o.FilterRelationships = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterReturnUniqueSchemeID(filterReturnUniqueSchemeID string) *ListPaymentsRequest {

	o.FilterReturnUniqueSchemeID = &filterReturnUniqueSchemeID

	return o
}

func (o *ListPaymentsRequest) WithoutFilterReturnUniqueSchemeID() *ListPaymentsRequest {

	o.FilterReturnUniqueSchemeID = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterReturnSubmissionStatus(filterReturnSubmissionStatus string) *ListPaymentsRequest {

	o.FilterReturnSubmissionStatus = &filterReturnSubmissionStatus

	return o
}

func (o *ListPaymentsRequest) WithoutFilterReturnSubmissionStatus() *ListPaymentsRequest {

	o.FilterReturnSubmissionStatus = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterReturnSubmissionSubmissionDateFrom(filterReturnSubmissionSubmissionDateFrom strfmt.DateTime) *ListPaymentsRequest {

	o.FilterReturnSubmissionSubmissionDateFrom = &filterReturnSubmissionSubmissionDateFrom

	return o
}

func (o *ListPaymentsRequest) WithoutFilterReturnSubmissionSubmissionDateFrom() *ListPaymentsRequest {

	o.FilterReturnSubmissionSubmissionDateFrom = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterReturnSubmissionSubmissionDateTo(filterReturnSubmissionSubmissionDateTo strfmt.DateTime) *ListPaymentsRequest {

	o.FilterReturnSubmissionSubmissionDateTo = &filterReturnSubmissionSubmissionDateTo

	return o
}

func (o *ListPaymentsRequest) WithoutFilterReturnSubmissionSubmissionDateTo() *ListPaymentsRequest {

	o.FilterReturnSubmissionSubmissionDateTo = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterSchemeTransactionID(filterSchemeTransactionID string) *ListPaymentsRequest {

	o.FilterSchemeTransactionID = &filterSchemeTransactionID

	return o
}

func (o *ListPaymentsRequest) WithoutFilterSchemeTransactionID() *ListPaymentsRequest {

	o.FilterSchemeTransactionID = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterSubmissionSchemeStatusCode(filterSubmissionSchemeStatusCode string) *ListPaymentsRequest {

	o.FilterSubmissionSchemeStatusCode = &filterSubmissionSchemeStatusCode

	return o
}

func (o *ListPaymentsRequest) WithoutFilterSubmissionSchemeStatusCode() *ListPaymentsRequest {

	o.FilterSubmissionSchemeStatusCode = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterSubmissionStatus(filterSubmissionStatus string) *ListPaymentsRequest {

	o.FilterSubmissionStatus = &filterSubmissionStatus

	return o
}

func (o *ListPaymentsRequest) WithoutFilterSubmissionStatus() *ListPaymentsRequest {

	o.FilterSubmissionStatus = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterSubmissionSubmissionDateFrom(filterSubmissionSubmissionDateFrom strfmt.DateTime) *ListPaymentsRequest {

	o.FilterSubmissionSubmissionDateFrom = &filterSubmissionSubmissionDateFrom

	return o
}

func (o *ListPaymentsRequest) WithoutFilterSubmissionSubmissionDateFrom() *ListPaymentsRequest {

	o.FilterSubmissionSubmissionDateFrom = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterSubmissionSubmissionDateTo(filterSubmissionSubmissionDateTo strfmt.DateTime) *ListPaymentsRequest {

	o.FilterSubmissionSubmissionDateTo = &filterSubmissionSubmissionDateTo

	return o
}

func (o *ListPaymentsRequest) WithoutFilterSubmissionSubmissionDateTo() *ListPaymentsRequest {

	o.FilterSubmissionSubmissionDateTo = nil

	return o
}

func (o *ListPaymentsRequest) WithFilterUniqueSchemeID(filterUniqueSchemeID string) *ListPaymentsRequest {

	o.FilterUniqueSchemeID = &filterUniqueSchemeID

	return o
}

func (o *ListPaymentsRequest) WithoutFilterUniqueSchemeID() *ListPaymentsRequest {

	o.FilterUniqueSchemeID = nil

	return o
}

func (o *ListPaymentsRequest) WithPageNumber(pageNumber string) *ListPaymentsRequest {

	o.PageNumber = &pageNumber

	return o
}

func (o *ListPaymentsRequest) WithoutPageNumber() *ListPaymentsRequest {

	o.PageNumber = nil

	return o
}

func (o *ListPaymentsRequest) WithPageSize(pageSize int64) *ListPaymentsRequest {

	o.PageSize = &pageSize

	return o
}

func (o *ListPaymentsRequest) WithoutPageSize() *ListPaymentsRequest {

	o.PageSize = nil

	return o
}

//////////////////
// WithContext adds the context to the list payments Request
func (o *ListPaymentsRequest) WithContext(ctx context.Context) *ListPaymentsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the list payments Request
func (o *ListPaymentsRequest) WithHTTPClient(client *http.Client) *ListPaymentsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ListPaymentsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FilterAdmissionSchemeStatusCode != nil {

		// query param filter[admission.scheme_status_code]
		var qrFilterAdmissionSchemeStatusCode string
		if o.FilterAdmissionSchemeStatusCode != nil {
			qrFilterAdmissionSchemeStatusCode = *o.FilterAdmissionSchemeStatusCode
		}
		qFilterAdmissionSchemeStatusCode := qrFilterAdmissionSchemeStatusCode
		if qFilterAdmissionSchemeStatusCode != "" {
			if err := r.SetQueryParam("filter[admission.scheme_status_code]", qFilterAdmissionSchemeStatusCode); err != nil {
				return err
			}
		}

	}

	if o.FilterAdmissionStatus != nil {

		// query param filter[admission.status]
		var qrFilterAdmissionStatus string
		if o.FilterAdmissionStatus != nil {
			qrFilterAdmissionStatus = *o.FilterAdmissionStatus
		}
		qFilterAdmissionStatus := qrFilterAdmissionStatus
		if qFilterAdmissionStatus != "" {
			if err := r.SetQueryParam("filter[admission.status]", qFilterAdmissionStatus); err != nil {
				return err
			}
		}

	}

	if o.FilterAmount != nil {

		// query param filter[amount]
		var qrFilterAmount string
		if o.FilterAmount != nil {
			qrFilterAmount = *o.FilterAmount
		}
		qFilterAmount := qrFilterAmount
		if qFilterAmount != "" {
			if err := r.SetQueryParam("filter[amount]", qFilterAmount); err != nil {
				return err
			}
		}

	}

	if o.FilterBeneficiaryPartyAccountName != nil {

		// query param filter[beneficiary_party.account_name]
		var qrFilterBeneficiaryPartyAccountName string
		if o.FilterBeneficiaryPartyAccountName != nil {
			qrFilterBeneficiaryPartyAccountName = *o.FilterBeneficiaryPartyAccountName
		}
		qFilterBeneficiaryPartyAccountName := qrFilterBeneficiaryPartyAccountName
		if qFilterBeneficiaryPartyAccountName != "" {
			if err := r.SetQueryParam("filter[beneficiary_party.account_name]", qFilterBeneficiaryPartyAccountName); err != nil {
				return err
			}
		}

	}

	if o.FilterBeneficiaryPartyAccountNumber != nil {

		// query param filter[beneficiary_party.account_number]
		var qrFilterBeneficiaryPartyAccountNumber string
		if o.FilterBeneficiaryPartyAccountNumber != nil {
			qrFilterBeneficiaryPartyAccountNumber = *o.FilterBeneficiaryPartyAccountNumber
		}
		qFilterBeneficiaryPartyAccountNumber := qrFilterBeneficiaryPartyAccountNumber
		if qFilterBeneficiaryPartyAccountNumber != "" {
			if err := r.SetQueryParam("filter[beneficiary_party.account_number]", qFilterBeneficiaryPartyAccountNumber); err != nil {
				return err
			}
		}

	}

	if o.FilterBeneficiaryPartyBankID != nil {

		// query param filter[beneficiary_party.bank_id]
		var qrFilterBeneficiaryPartyBankID string
		if o.FilterBeneficiaryPartyBankID != nil {
			qrFilterBeneficiaryPartyBankID = *o.FilterBeneficiaryPartyBankID
		}
		qFilterBeneficiaryPartyBankID := qrFilterBeneficiaryPartyBankID
		if qFilterBeneficiaryPartyBankID != "" {
			if err := r.SetQueryParam("filter[beneficiary_party.bank_id]", qFilterBeneficiaryPartyBankID); err != nil {
				return err
			}
		}

	}

	if o.FilterCurrency != nil {

		// query param filter[currency]
		var qrFilterCurrency string
		if o.FilterCurrency != nil {
			qrFilterCurrency = *o.FilterCurrency
		}
		qFilterCurrency := qrFilterCurrency
		if qFilterCurrency != "" {
			if err := r.SetQueryParam("filter[currency]", qFilterCurrency); err != nil {
				return err
			}
		}

	}

	if o.FilterDebtorPartyAccountName != nil {

		// query param filter[debtor_party.account_name]
		var qrFilterDebtorPartyAccountName string
		if o.FilterDebtorPartyAccountName != nil {
			qrFilterDebtorPartyAccountName = *o.FilterDebtorPartyAccountName
		}
		qFilterDebtorPartyAccountName := qrFilterDebtorPartyAccountName
		if qFilterDebtorPartyAccountName != "" {
			if err := r.SetQueryParam("filter[debtor_party.account_name]", qFilterDebtorPartyAccountName); err != nil {
				return err
			}
		}

	}

	if o.FilterDebtorPartyAccountNumber != nil {

		// query param filter[debtor_party.account_number]
		var qrFilterDebtorPartyAccountNumber string
		if o.FilterDebtorPartyAccountNumber != nil {
			qrFilterDebtorPartyAccountNumber = *o.FilterDebtorPartyAccountNumber
		}
		qFilterDebtorPartyAccountNumber := qrFilterDebtorPartyAccountNumber
		if qFilterDebtorPartyAccountNumber != "" {
			if err := r.SetQueryParam("filter[debtor_party.account_number]", qFilterDebtorPartyAccountNumber); err != nil {
				return err
			}
		}

	}

	if o.FilterDebtorPartyBankID != nil {

		// query param filter[debtor_party.bank_id]
		var qrFilterDebtorPartyBankID string
		if o.FilterDebtorPartyBankID != nil {
			qrFilterDebtorPartyBankID = *o.FilterDebtorPartyBankID
		}
		qFilterDebtorPartyBankID := qrFilterDebtorPartyBankID
		if qFilterDebtorPartyBankID != "" {
			if err := r.SetQueryParam("filter[debtor_party.bank_id]", qFilterDebtorPartyBankID); err != nil {
				return err
			}
		}

	}

	if o.FilterEndToEndReference != nil {

		// query param filter[end_to_end_reference]
		var qrFilterEndToEndReference string
		if o.FilterEndToEndReference != nil {
			qrFilterEndToEndReference = *o.FilterEndToEndReference
		}
		qFilterEndToEndReference := qrFilterEndToEndReference
		if qFilterEndToEndReference != "" {
			if err := r.SetQueryParam("filter[end_to_end_reference]", qFilterEndToEndReference); err != nil {
				return err
			}
		}

	}

	valuesFilterNotRelationships := o.FilterNotRelationships

	joinedFilterNotRelationships := swag.JoinByFormat(valuesFilterNotRelationships, "csv")
	// query array param filter[not_relationships]
	if err := r.SetQueryParam("filter[not_relationships]", joinedFilterNotRelationships...); err != nil {
		return err
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

	if o.FilterPaymentType != nil {

		// query param filter[payment_type]
		var qrFilterPaymentType string
		if o.FilterPaymentType != nil {
			qrFilterPaymentType = *o.FilterPaymentType
		}
		qFilterPaymentType := qrFilterPaymentType
		if qFilterPaymentType != "" {
			if err := r.SetQueryParam("filter[payment_type]", qFilterPaymentType); err != nil {
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

	if o.FilterReference != nil {

		// query param filter[reference]
		var qrFilterReference string
		if o.FilterReference != nil {
			qrFilterReference = *o.FilterReference
		}
		qFilterReference := qrFilterReference
		if qFilterReference != "" {
			if err := r.SetQueryParam("filter[reference]", qFilterReference); err != nil {
				return err
			}
		}

	}

	valuesFilterRelationships := o.FilterRelationships

	joinedFilterRelationships := swag.JoinByFormat(valuesFilterRelationships, "")
	// query array param filter[relationships]
	if err := r.SetQueryParam("filter[relationships]", joinedFilterRelationships...); err != nil {
		return err
	}

	if o.FilterReturnUniqueSchemeID != nil {

		// query param filter[return.unique_scheme_id]
		var qrFilterReturnUniqueSchemeID string
		if o.FilterReturnUniqueSchemeID != nil {
			qrFilterReturnUniqueSchemeID = *o.FilterReturnUniqueSchemeID
		}
		qFilterReturnUniqueSchemeID := qrFilterReturnUniqueSchemeID
		if qFilterReturnUniqueSchemeID != "" {
			if err := r.SetQueryParam("filter[return.unique_scheme_id]", qFilterReturnUniqueSchemeID); err != nil {
				return err
			}
		}

	}

	if o.FilterReturnSubmissionStatus != nil {

		// query param filter[return_submission.status]
		var qrFilterReturnSubmissionStatus string
		if o.FilterReturnSubmissionStatus != nil {
			qrFilterReturnSubmissionStatus = *o.FilterReturnSubmissionStatus
		}
		qFilterReturnSubmissionStatus := qrFilterReturnSubmissionStatus
		if qFilterReturnSubmissionStatus != "" {
			if err := r.SetQueryParam("filter[return_submission.status]", qFilterReturnSubmissionStatus); err != nil {
				return err
			}
		}

	}

	if o.FilterReturnSubmissionSubmissionDateFrom != nil {

		// query param filter[return_submission.submission_date_from]
		var qrFilterReturnSubmissionSubmissionDateFrom strfmt.DateTime
		if o.FilterReturnSubmissionSubmissionDateFrom != nil {
			qrFilterReturnSubmissionSubmissionDateFrom = *o.FilterReturnSubmissionSubmissionDateFrom
		}
		qFilterReturnSubmissionSubmissionDateFrom := qrFilterReturnSubmissionSubmissionDateFrom.String()
		if qFilterReturnSubmissionSubmissionDateFrom != "" {
			if err := r.SetQueryParam("filter[return_submission.submission_date_from]", qFilterReturnSubmissionSubmissionDateFrom); err != nil {
				return err
			}
		}

	}

	if o.FilterReturnSubmissionSubmissionDateTo != nil {

		// query param filter[return_submission.submission_date_to]
		var qrFilterReturnSubmissionSubmissionDateTo strfmt.DateTime
		if o.FilterReturnSubmissionSubmissionDateTo != nil {
			qrFilterReturnSubmissionSubmissionDateTo = *o.FilterReturnSubmissionSubmissionDateTo
		}
		qFilterReturnSubmissionSubmissionDateTo := qrFilterReturnSubmissionSubmissionDateTo.String()
		if qFilterReturnSubmissionSubmissionDateTo != "" {
			if err := r.SetQueryParam("filter[return_submission.submission_date_to]", qFilterReturnSubmissionSubmissionDateTo); err != nil {
				return err
			}
		}

	}

	if o.FilterSchemeTransactionID != nil {

		// query param filter[scheme_transaction_id]
		var qrFilterSchemeTransactionID string
		if o.FilterSchemeTransactionID != nil {
			qrFilterSchemeTransactionID = *o.FilterSchemeTransactionID
		}
		qFilterSchemeTransactionID := qrFilterSchemeTransactionID
		if qFilterSchemeTransactionID != "" {
			if err := r.SetQueryParam("filter[scheme_transaction_id]", qFilterSchemeTransactionID); err != nil {
				return err
			}
		}

	}

	if o.FilterSubmissionSchemeStatusCode != nil {

		// query param filter[submission.scheme_status_code]
		var qrFilterSubmissionSchemeStatusCode string
		if o.FilterSubmissionSchemeStatusCode != nil {
			qrFilterSubmissionSchemeStatusCode = *o.FilterSubmissionSchemeStatusCode
		}
		qFilterSubmissionSchemeStatusCode := qrFilterSubmissionSchemeStatusCode
		if qFilterSubmissionSchemeStatusCode != "" {
			if err := r.SetQueryParam("filter[submission.scheme_status_code]", qFilterSubmissionSchemeStatusCode); err != nil {
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
