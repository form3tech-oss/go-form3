// Code generated by go-swagger; DO NOT EDIT.

package claims

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

// Client.ListClaims creates a new ListClaimsRequest object
// with the default values initialized.
func (c *Client) ListClaims() *ListClaimsRequest {
	var ()
	return &ListClaimsRequest{

		FilterBeneficiaryPartyAccountNumber: c.Defaults.GetStringPtr("ListClaims", "filter[beneficiary_party.account_number]"),

		FilterBeneficiaryPartyBankID: c.Defaults.GetStringPtr("ListClaims", "filter[beneficiary_party.bank_id]"),

		FilterClearingID: c.Defaults.GetStringPtr("ListClaims", "filter[clearing_id]"),

		FilterContactName: c.Defaults.GetStringPtr("ListClaims", "filter[contact_name]"),

		FilterDebtorPartyAccountNumber: c.Defaults.GetStringPtr("ListClaims", "filter[debtor_party.account_number]"),

		FilterDebtorPartyBankID: c.Defaults.GetStringPtr("ListClaims", "filter[debtor_party.bank_id]"),

		FilterOrganisationID: make([]strfmt.UUID, 0),

		FilterOriginalInstructionReference: c.Defaults.GetStringPtr("ListClaims", "filter[original_instruction.reference]"),

		FilterPaymentScheme: c.Defaults.GetStringPtr("ListClaims", "filter[payment_scheme]"),

		FilterReasonCode: c.Defaults.GetStringPtr("ListClaims", "filter[reason_code]"),

		FilterReference: c.Defaults.GetStringPtr("ListClaims", "filter[reference]"),

		FilterReversalStatus: c.Defaults.GetStringPtr("ListClaims", "filter[reversal.status]"),

		FilterReversalSubmissionDateFrom: c.Defaults.GetStrfmtDateTimePtr("ListClaims", "filter[reversal.submission_date_from]"),

		FilterReversalSubmissionDateTo: c.Defaults.GetStrfmtDateTimePtr("ListClaims", "filter[reversal.submission_date_to]"),

		FilterSubmissionStatus: c.Defaults.GetStringPtr("ListClaims", "filter[submission.status]"),

		FilterSubmissionSubmissionDateFrom: c.Defaults.GetStrfmtDateTimePtr("ListClaims", "filter[submission.submission_date_from]"),

		FilterSubmissionSubmissionDateTo: c.Defaults.GetStrfmtDateTimePtr("ListClaims", "filter[submission.submission_date_to]"),

		PageNumber: c.Defaults.GetStringPtr("ListClaims", "page[number]"),

		PageSize: c.Defaults.GetInt64Ptr("ListClaims", "page[size]"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type ListClaimsRequest struct {

	/*FilterBeneficiaryPartyAccountNumber*/

	FilterBeneficiaryPartyAccountNumber *string

	/*FilterBeneficiaryPartyBankID*/

	FilterBeneficiaryPartyBankID *string

	/*FilterClearingID*/

	FilterClearingID *string

	/*FilterContactName*/

	FilterContactName *string

	/*FilterDebtorPartyAccountNumber*/

	FilterDebtorPartyAccountNumber *string

	/*FilterDebtorPartyBankID*/

	FilterDebtorPartyBankID *string

	/*FilterOrganisationID      Filter by organisation id      */

	FilterOrganisationID []strfmt.UUID

	/*FilterOriginalInstructionReference*/

	FilterOriginalInstructionReference *string

	/*FilterPaymentScheme*/

	FilterPaymentScheme *string

	/*FilterReasonCode*/

	FilterReasonCode *string

	/*FilterReference*/

	FilterReference *string

	/*FilterReversalStatus*/

	FilterReversalStatus *string

	/*FilterReversalSubmissionDateFrom*/

	FilterReversalSubmissionDateFrom *strfmt.DateTime

	/*FilterReversalSubmissionDateTo*/

	FilterReversalSubmissionDateTo *strfmt.DateTime

	/*FilterSubmissionStatus*/

	FilterSubmissionStatus *string

	/*FilterSubmissionSubmissionDateFrom*/

	FilterSubmissionSubmissionDateFrom *strfmt.DateTime

	/*FilterSubmissionSubmissionDateTo*/

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

func (o *ListClaimsRequest) FromJson(j string) *ListClaimsRequest {

	return o
}

func (o *ListClaimsRequest) WithFilterBeneficiaryPartyAccountNumber(filterBeneficiaryPartyAccountNumber string) *ListClaimsRequest {

	o.FilterBeneficiaryPartyAccountNumber = &filterBeneficiaryPartyAccountNumber

	return o
}

func (o *ListClaimsRequest) WithoutFilterBeneficiaryPartyAccountNumber() *ListClaimsRequest {

	o.FilterBeneficiaryPartyAccountNumber = nil

	return o
}

func (o *ListClaimsRequest) WithFilterBeneficiaryPartyBankID(filterBeneficiaryPartyBankID string) *ListClaimsRequest {

	o.FilterBeneficiaryPartyBankID = &filterBeneficiaryPartyBankID

	return o
}

func (o *ListClaimsRequest) WithoutFilterBeneficiaryPartyBankID() *ListClaimsRequest {

	o.FilterBeneficiaryPartyBankID = nil

	return o
}

func (o *ListClaimsRequest) WithFilterClearingID(filterClearingID string) *ListClaimsRequest {

	o.FilterClearingID = &filterClearingID

	return o
}

func (o *ListClaimsRequest) WithoutFilterClearingID() *ListClaimsRequest {

	o.FilterClearingID = nil

	return o
}

func (o *ListClaimsRequest) WithFilterContactName(filterContactName string) *ListClaimsRequest {

	o.FilterContactName = &filterContactName

	return o
}

func (o *ListClaimsRequest) WithoutFilterContactName() *ListClaimsRequest {

	o.FilterContactName = nil

	return o
}

func (o *ListClaimsRequest) WithFilterDebtorPartyAccountNumber(filterDebtorPartyAccountNumber string) *ListClaimsRequest {

	o.FilterDebtorPartyAccountNumber = &filterDebtorPartyAccountNumber

	return o
}

func (o *ListClaimsRequest) WithoutFilterDebtorPartyAccountNumber() *ListClaimsRequest {

	o.FilterDebtorPartyAccountNumber = nil

	return o
}

func (o *ListClaimsRequest) WithFilterDebtorPartyBankID(filterDebtorPartyBankID string) *ListClaimsRequest {

	o.FilterDebtorPartyBankID = &filterDebtorPartyBankID

	return o
}

func (o *ListClaimsRequest) WithoutFilterDebtorPartyBankID() *ListClaimsRequest {

	o.FilterDebtorPartyBankID = nil

	return o
}

func (o *ListClaimsRequest) WithFilterOrganisationID(filterOrganisationID []strfmt.UUID) *ListClaimsRequest {

	o.FilterOrganisationID = filterOrganisationID

	return o
}

func (o *ListClaimsRequest) WithoutFilterOrganisationID() *ListClaimsRequest {

	o.FilterOrganisationID = nil

	return o
}

func (o *ListClaimsRequest) WithFilterOriginalInstructionReference(filterOriginalInstructionReference string) *ListClaimsRequest {

	o.FilterOriginalInstructionReference = &filterOriginalInstructionReference

	return o
}

func (o *ListClaimsRequest) WithoutFilterOriginalInstructionReference() *ListClaimsRequest {

	o.FilterOriginalInstructionReference = nil

	return o
}

func (o *ListClaimsRequest) WithFilterPaymentScheme(filterPaymentScheme string) *ListClaimsRequest {

	o.FilterPaymentScheme = &filterPaymentScheme

	return o
}

func (o *ListClaimsRequest) WithoutFilterPaymentScheme() *ListClaimsRequest {

	o.FilterPaymentScheme = nil

	return o
}

func (o *ListClaimsRequest) WithFilterReasonCode(filterReasonCode string) *ListClaimsRequest {

	o.FilterReasonCode = &filterReasonCode

	return o
}

func (o *ListClaimsRequest) WithoutFilterReasonCode() *ListClaimsRequest {

	o.FilterReasonCode = nil

	return o
}

func (o *ListClaimsRequest) WithFilterReference(filterReference string) *ListClaimsRequest {

	o.FilterReference = &filterReference

	return o
}

func (o *ListClaimsRequest) WithoutFilterReference() *ListClaimsRequest {

	o.FilterReference = nil

	return o
}

func (o *ListClaimsRequest) WithFilterReversalStatus(filterReversalStatus string) *ListClaimsRequest {

	o.FilterReversalStatus = &filterReversalStatus

	return o
}

func (o *ListClaimsRequest) WithoutFilterReversalStatus() *ListClaimsRequest {

	o.FilterReversalStatus = nil

	return o
}

func (o *ListClaimsRequest) WithFilterReversalSubmissionDateFrom(filterReversalSubmissionDateFrom strfmt.DateTime) *ListClaimsRequest {

	o.FilterReversalSubmissionDateFrom = &filterReversalSubmissionDateFrom

	return o
}

func (o *ListClaimsRequest) WithoutFilterReversalSubmissionDateFrom() *ListClaimsRequest {

	o.FilterReversalSubmissionDateFrom = nil

	return o
}

func (o *ListClaimsRequest) WithFilterReversalSubmissionDateTo(filterReversalSubmissionDateTo strfmt.DateTime) *ListClaimsRequest {

	o.FilterReversalSubmissionDateTo = &filterReversalSubmissionDateTo

	return o
}

func (o *ListClaimsRequest) WithoutFilterReversalSubmissionDateTo() *ListClaimsRequest {

	o.FilterReversalSubmissionDateTo = nil

	return o
}

func (o *ListClaimsRequest) WithFilterSubmissionStatus(filterSubmissionStatus string) *ListClaimsRequest {

	o.FilterSubmissionStatus = &filterSubmissionStatus

	return o
}

func (o *ListClaimsRequest) WithoutFilterSubmissionStatus() *ListClaimsRequest {

	o.FilterSubmissionStatus = nil

	return o
}

func (o *ListClaimsRequest) WithFilterSubmissionSubmissionDateFrom(filterSubmissionSubmissionDateFrom strfmt.DateTime) *ListClaimsRequest {

	o.FilterSubmissionSubmissionDateFrom = &filterSubmissionSubmissionDateFrom

	return o
}

func (o *ListClaimsRequest) WithoutFilterSubmissionSubmissionDateFrom() *ListClaimsRequest {

	o.FilterSubmissionSubmissionDateFrom = nil

	return o
}

func (o *ListClaimsRequest) WithFilterSubmissionSubmissionDateTo(filterSubmissionSubmissionDateTo strfmt.DateTime) *ListClaimsRequest {

	o.FilterSubmissionSubmissionDateTo = &filterSubmissionSubmissionDateTo

	return o
}

func (o *ListClaimsRequest) WithoutFilterSubmissionSubmissionDateTo() *ListClaimsRequest {

	o.FilterSubmissionSubmissionDateTo = nil

	return o
}

func (o *ListClaimsRequest) WithPageNumber(pageNumber string) *ListClaimsRequest {

	o.PageNumber = &pageNumber

	return o
}

func (o *ListClaimsRequest) WithoutPageNumber() *ListClaimsRequest {

	o.PageNumber = nil

	return o
}

func (o *ListClaimsRequest) WithPageSize(pageSize int64) *ListClaimsRequest {

	o.PageSize = &pageSize

	return o
}

func (o *ListClaimsRequest) WithoutPageSize() *ListClaimsRequest {

	o.PageSize = nil

	return o
}

//////////////////
// WithContext adds the context to the list claims Request
func (o *ListClaimsRequest) WithContext(ctx context.Context) *ListClaimsRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the list claims Request
func (o *ListClaimsRequest) WithHTTPClient(client *http.Client) *ListClaimsRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *ListClaimsRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.FilterClearingID != nil {

		// query param filter[clearing_id]
		var qrFilterClearingID string
		if o.FilterClearingID != nil {
			qrFilterClearingID = *o.FilterClearingID
		}
		qFilterClearingID := qrFilterClearingID
		if qFilterClearingID != "" {
			if err := r.SetQueryParam("filter[clearing_id]", qFilterClearingID); err != nil {
				return err
			}
		}

	}

	if o.FilterContactName != nil {

		// query param filter[contact_name]
		var qrFilterContactName string
		if o.FilterContactName != nil {
			qrFilterContactName = *o.FilterContactName
		}
		qFilterContactName := qrFilterContactName
		if qFilterContactName != "" {
			if err := r.SetQueryParam("filter[contact_name]", qFilterContactName); err != nil {
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

	var valuesFilterOrganisationID []string
	for _, v := range o.FilterOrganisationID {
		valuesFilterOrganisationID = append(valuesFilterOrganisationID, v.String())
	}

	joinedFilterOrganisationID := swag.JoinByFormat(valuesFilterOrganisationID, "")
	// query array param filter[organisation_id]
	if err := r.SetQueryParam("filter[organisation_id]", joinedFilterOrganisationID...); err != nil {
		return err
	}

	if o.FilterOriginalInstructionReference != nil {

		// query param filter[original_instruction.reference]
		var qrFilterOriginalInstructionReference string
		if o.FilterOriginalInstructionReference != nil {
			qrFilterOriginalInstructionReference = *o.FilterOriginalInstructionReference
		}
		qFilterOriginalInstructionReference := qrFilterOriginalInstructionReference
		if qFilterOriginalInstructionReference != "" {
			if err := r.SetQueryParam("filter[original_instruction.reference]", qFilterOriginalInstructionReference); err != nil {
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

	if o.FilterReasonCode != nil {

		// query param filter[reason_code]
		var qrFilterReasonCode string
		if o.FilterReasonCode != nil {
			qrFilterReasonCode = *o.FilterReasonCode
		}
		qFilterReasonCode := qrFilterReasonCode
		if qFilterReasonCode != "" {
			if err := r.SetQueryParam("filter[reason_code]", qFilterReasonCode); err != nil {
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

	if o.FilterReversalStatus != nil {

		// query param filter[reversal.status]
		var qrFilterReversalStatus string
		if o.FilterReversalStatus != nil {
			qrFilterReversalStatus = *o.FilterReversalStatus
		}
		qFilterReversalStatus := qrFilterReversalStatus
		if qFilterReversalStatus != "" {
			if err := r.SetQueryParam("filter[reversal.status]", qFilterReversalStatus); err != nil {
				return err
			}
		}

	}

	if o.FilterReversalSubmissionDateFrom != nil {

		// query param filter[reversal.submission_date_from]
		var qrFilterReversalSubmissionDateFrom strfmt.DateTime
		if o.FilterReversalSubmissionDateFrom != nil {
			qrFilterReversalSubmissionDateFrom = *o.FilterReversalSubmissionDateFrom
		}
		qFilterReversalSubmissionDateFrom := qrFilterReversalSubmissionDateFrom.String()
		if qFilterReversalSubmissionDateFrom != "" {
			if err := r.SetQueryParam("filter[reversal.submission_date_from]", qFilterReversalSubmissionDateFrom); err != nil {
				return err
			}
		}

	}

	if o.FilterReversalSubmissionDateTo != nil {

		// query param filter[reversal.submission_date_to]
		var qrFilterReversalSubmissionDateTo strfmt.DateTime
		if o.FilterReversalSubmissionDateTo != nil {
			qrFilterReversalSubmissionDateTo = *o.FilterReversalSubmissionDateTo
		}
		qFilterReversalSubmissionDateTo := qrFilterReversalSubmissionDateTo.String()
		if qFilterReversalSubmissionDateTo != "" {
			if err := r.SetQueryParam("filter[reversal.submission_date_to]", qFilterReversalSubmissionDateTo); err != nil {
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
