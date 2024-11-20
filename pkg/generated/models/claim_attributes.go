// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClaimAttributes claim attributes
// swagger:model ClaimAttributes
type ClaimAttributes struct {

	// beneficiary party
	// Required: true
	BeneficiaryParty *ClaimAttributesBeneficiaryParty `json:"beneficiary_party"`

	// clearing id
	// Required: true
	// Pattern: ^[0-9]{6}$
	ClearingID *string `json:"clearing_id"`

	// contact name
	ContactName string `json:"contact_name,omitempty"`

	// debtor party
	// Required: true
	DebtorParty *ClaimAttributesDebtorParty `json:"debtor_party"`

	// disputed transactions
	// Required: true
	DisputedTransactions []*DisputedTransaction `json:"disputed_transactions"`

	// number of claims
	// Required: true
	NumberOfClaims *int64 `json:"number_of_claims"`

	// original instruction
	// Required: true
	OriginalInstruction *ClaimAttributesOriginalInstruction `json:"original_instruction"`

	// payment scheme
	// Required: true
	PaymentScheme *string `json:"payment_scheme"`

	// processing date
	// Format: date
	ProcessingDate strfmt.Date `json:"processing_date,omitempty"`

	// reason code
	// Required: true
	// Pattern: ^[1-9]$
	ReasonCode *string `json:"reason_code"`

	// reference
	// Required: true
	Reference *string `json:"reference"`

	// request date
	// Format: date
	RequestDate strfmt.Date `json:"request_date,omitempty"`
}

func ClaimAttributesWithDefaults(defaults client.Defaults) *ClaimAttributes {
	return &ClaimAttributes{

		BeneficiaryParty: ClaimAttributesBeneficiaryPartyWithDefaults(defaults),

		ClearingID: defaults.GetStringPtr("ClaimAttributes", "clearing_id"),

		ContactName: defaults.GetString("ClaimAttributes", "contact_name"),

		DebtorParty: ClaimAttributesDebtorPartyWithDefaults(defaults),

		DisputedTransactions: make([]*DisputedTransaction, 0),

		NumberOfClaims: defaults.GetInt64Ptr("ClaimAttributes", "number_of_claims"),

		OriginalInstruction: ClaimAttributesOriginalInstructionWithDefaults(defaults),

		PaymentScheme: defaults.GetStringPtr("ClaimAttributes", "payment_scheme"),

		ProcessingDate: defaults.GetStrfmtDate("ClaimAttributes", "processing_date"),

		ReasonCode: defaults.GetStringPtr("ClaimAttributes", "reason_code"),

		Reference: defaults.GetStringPtr("ClaimAttributes", "reference"),

		RequestDate: defaults.GetStrfmtDate("ClaimAttributes", "request_date"),
	}
}

func (m *ClaimAttributes) WithBeneficiaryParty(beneficiaryParty ClaimAttributesBeneficiaryParty) *ClaimAttributes {

	m.BeneficiaryParty = &beneficiaryParty

	return m
}

func (m *ClaimAttributes) WithoutBeneficiaryParty() *ClaimAttributes {
	m.BeneficiaryParty = nil
	return m
}

func (m *ClaimAttributes) WithClearingID(clearingID string) *ClaimAttributes {

	m.ClearingID = &clearingID

	return m
}

func (m *ClaimAttributes) WithoutClearingID() *ClaimAttributes {
	m.ClearingID = nil
	return m
}

func (m *ClaimAttributes) WithContactName(contactName string) *ClaimAttributes {

	m.ContactName = contactName

	return m
}

func (m *ClaimAttributes) WithDebtorParty(debtorParty ClaimAttributesDebtorParty) *ClaimAttributes {

	m.DebtorParty = &debtorParty

	return m
}

func (m *ClaimAttributes) WithoutDebtorParty() *ClaimAttributes {
	m.DebtorParty = nil
	return m
}

func (m *ClaimAttributes) WithDisputedTransactions(disputedTransactions []*DisputedTransaction) *ClaimAttributes {

	m.DisputedTransactions = disputedTransactions

	return m
}

func (m *ClaimAttributes) WithNumberOfClaims(numberOfClaims int64) *ClaimAttributes {

	m.NumberOfClaims = &numberOfClaims

	return m
}

func (m *ClaimAttributes) WithoutNumberOfClaims() *ClaimAttributes {
	m.NumberOfClaims = nil
	return m
}

func (m *ClaimAttributes) WithOriginalInstruction(originalInstruction ClaimAttributesOriginalInstruction) *ClaimAttributes {

	m.OriginalInstruction = &originalInstruction

	return m
}

func (m *ClaimAttributes) WithoutOriginalInstruction() *ClaimAttributes {
	m.OriginalInstruction = nil
	return m
}

func (m *ClaimAttributes) WithPaymentScheme(paymentScheme string) *ClaimAttributes {

	m.PaymentScheme = &paymentScheme

	return m
}

func (m *ClaimAttributes) WithoutPaymentScheme() *ClaimAttributes {
	m.PaymentScheme = nil
	return m
}

func (m *ClaimAttributes) WithProcessingDate(processingDate strfmt.Date) *ClaimAttributes {

	m.ProcessingDate = processingDate

	return m
}

func (m *ClaimAttributes) WithReasonCode(reasonCode string) *ClaimAttributes {

	m.ReasonCode = &reasonCode

	return m
}

func (m *ClaimAttributes) WithoutReasonCode() *ClaimAttributes {
	m.ReasonCode = nil
	return m
}

func (m *ClaimAttributes) WithReference(reference string) *ClaimAttributes {

	m.Reference = &reference

	return m
}

func (m *ClaimAttributes) WithoutReference() *ClaimAttributes {
	m.Reference = nil
	return m
}

func (m *ClaimAttributes) WithRequestDate(requestDate strfmt.Date) *ClaimAttributes {

	m.RequestDate = requestDate

	return m
}

// Validate validates this claim attributes
func (m *ClaimAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBeneficiaryParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClearingID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtorParty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisputedTransactions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberOfClaims(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginalInstruction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentScheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcessingDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReasonCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimAttributes) validateBeneficiaryParty(formats strfmt.Registry) error {

	if err := validate.Required("beneficiary_party", "body", m.BeneficiaryParty); err != nil {
		return err
	}

	if m.BeneficiaryParty != nil {
		if err := m.BeneficiaryParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("beneficiary_party")
			}
			return err
		}
	}

	return nil
}

func (m *ClaimAttributes) validateClearingID(formats strfmt.Registry) error {

	if err := validate.Required("clearing_id", "body", m.ClearingID); err != nil {
		return err
	}

	if err := validate.Pattern("clearing_id", "body", string(*m.ClearingID), `^[0-9]{6}$`); err != nil {
		return err
	}

	return nil
}

func (m *ClaimAttributes) validateDebtorParty(formats strfmt.Registry) error {

	if err := validate.Required("debtor_party", "body", m.DebtorParty); err != nil {
		return err
	}

	if m.DebtorParty != nil {
		if err := m.DebtorParty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("debtor_party")
			}
			return err
		}
	}

	return nil
}

func (m *ClaimAttributes) validateDisputedTransactions(formats strfmt.Registry) error {

	if err := validate.Required("disputed_transactions", "body", m.DisputedTransactions); err != nil {
		return err
	}

	for i := 0; i < len(m.DisputedTransactions); i++ {
		if swag.IsZero(m.DisputedTransactions[i]) { // not required
			continue
		}

		if m.DisputedTransactions[i] != nil {
			if err := m.DisputedTransactions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disputed_transactions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClaimAttributes) validateNumberOfClaims(formats strfmt.Registry) error {

	if err := validate.Required("number_of_claims", "body", m.NumberOfClaims); err != nil {
		return err
	}

	return nil
}

func (m *ClaimAttributes) validateOriginalInstruction(formats strfmt.Registry) error {

	if err := validate.Required("original_instruction", "body", m.OriginalInstruction); err != nil {
		return err
	}

	if m.OriginalInstruction != nil {
		if err := m.OriginalInstruction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("original_instruction")
			}
			return err
		}
	}

	return nil
}

func (m *ClaimAttributes) validatePaymentScheme(formats strfmt.Registry) error {

	if err := validate.Required("payment_scheme", "body", m.PaymentScheme); err != nil {
		return err
	}

	return nil
}

func (m *ClaimAttributes) validateProcessingDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ProcessingDate) { // not required
		return nil
	}

	if err := validate.FormatOf("processing_date", "body", "date", m.ProcessingDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClaimAttributes) validateReasonCode(formats strfmt.Registry) error {

	if err := validate.Required("reason_code", "body", m.ReasonCode); err != nil {
		return err
	}

	if err := validate.Pattern("reason_code", "body", string(*m.ReasonCode), `^[1-9]$`); err != nil {
		return err
	}

	return nil
}

func (m *ClaimAttributes) validateReference(formats strfmt.Registry) error {

	if err := validate.Required("reference", "body", m.Reference); err != nil {
		return err
	}

	return nil
}

func (m *ClaimAttributes) validateRequestDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestDate) { // not required
		return nil
	}

	if err := validate.FormatOf("request_date", "body", "date", m.RequestDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimAttributes) UnmarshalBinary(b []byte) error {
	var res ClaimAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ClaimAttributesBeneficiaryParty claim attributes beneficiary party
// swagger:model ClaimAttributesBeneficiaryParty
type ClaimAttributesBeneficiaryParty struct {

	// account number
	// Required: true
	// Pattern: ^[A-Z0-9]{6,34}$
	AccountNumber *string `json:"account_number"`

	// bank id
	// Required: true
	BankID *string `json:"bank_id"`
}

func ClaimAttributesBeneficiaryPartyWithDefaults(defaults client.Defaults) *ClaimAttributesBeneficiaryParty {
	return &ClaimAttributesBeneficiaryParty{

		AccountNumber: defaults.GetStringPtr("ClaimAttributesBeneficiaryParty", "account_number"),

		BankID: defaults.GetStringPtr("ClaimAttributesBeneficiaryParty", "bank_id"),
	}
}

func (m *ClaimAttributesBeneficiaryParty) WithAccountNumber(accountNumber string) *ClaimAttributesBeneficiaryParty {

	m.AccountNumber = &accountNumber

	return m
}

func (m *ClaimAttributesBeneficiaryParty) WithoutAccountNumber() *ClaimAttributesBeneficiaryParty {
	m.AccountNumber = nil
	return m
}

func (m *ClaimAttributesBeneficiaryParty) WithBankID(bankID string) *ClaimAttributesBeneficiaryParty {

	m.BankID = &bankID

	return m
}

func (m *ClaimAttributesBeneficiaryParty) WithoutBankID() *ClaimAttributesBeneficiaryParty {
	m.BankID = nil
	return m
}

// Validate validates this claim attributes beneficiary party
func (m *ClaimAttributesBeneficiaryParty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimAttributesBeneficiaryParty) validateAccountNumber(formats strfmt.Registry) error {

	if err := validate.Required("beneficiary_party"+"."+"account_number", "body", m.AccountNumber); err != nil {
		return err
	}

	if err := validate.Pattern("beneficiary_party"+"."+"account_number", "body", string(*m.AccountNumber), `^[A-Z0-9]{6,34}$`); err != nil {
		return err
	}

	return nil
}

func (m *ClaimAttributesBeneficiaryParty) validateBankID(formats strfmt.Registry) error {

	if err := validate.Required("beneficiary_party"+"."+"bank_id", "body", m.BankID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimAttributesBeneficiaryParty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimAttributesBeneficiaryParty) UnmarshalBinary(b []byte) error {
	var res ClaimAttributesBeneficiaryParty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimAttributesBeneficiaryParty) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ClaimAttributesDebtorParty claim attributes debtor party
// swagger:model ClaimAttributesDebtorParty
type ClaimAttributesDebtorParty struct {

	// account name
	// Required: true
	AccountName *string `json:"account_name"`

	// account number
	// Required: true
	// Pattern: ^[A-Z0-9]{6,34}$
	AccountNumber *string `json:"account_number"`

	// bank id
	// Required: true
	BankID *string `json:"bank_id"`
}

func ClaimAttributesDebtorPartyWithDefaults(defaults client.Defaults) *ClaimAttributesDebtorParty {
	return &ClaimAttributesDebtorParty{

		AccountName: defaults.GetStringPtr("ClaimAttributesDebtorParty", "account_name"),

		AccountNumber: defaults.GetStringPtr("ClaimAttributesDebtorParty", "account_number"),

		BankID: defaults.GetStringPtr("ClaimAttributesDebtorParty", "bank_id"),
	}
}

func (m *ClaimAttributesDebtorParty) WithAccountName(accountName string) *ClaimAttributesDebtorParty {

	m.AccountName = &accountName

	return m
}

func (m *ClaimAttributesDebtorParty) WithoutAccountName() *ClaimAttributesDebtorParty {
	m.AccountName = nil
	return m
}

func (m *ClaimAttributesDebtorParty) WithAccountNumber(accountNumber string) *ClaimAttributesDebtorParty {

	m.AccountNumber = &accountNumber

	return m
}

func (m *ClaimAttributesDebtorParty) WithoutAccountNumber() *ClaimAttributesDebtorParty {
	m.AccountNumber = nil
	return m
}

func (m *ClaimAttributesDebtorParty) WithBankID(bankID string) *ClaimAttributesDebtorParty {

	m.BankID = &bankID

	return m
}

func (m *ClaimAttributesDebtorParty) WithoutBankID() *ClaimAttributesDebtorParty {
	m.BankID = nil
	return m
}

// Validate validates this claim attributes debtor party
func (m *ClaimAttributesDebtorParty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimAttributesDebtorParty) validateAccountName(formats strfmt.Registry) error {

	if err := validate.Required("debtor_party"+"."+"account_name", "body", m.AccountName); err != nil {
		return err
	}

	return nil
}

func (m *ClaimAttributesDebtorParty) validateAccountNumber(formats strfmt.Registry) error {

	if err := validate.Required("debtor_party"+"."+"account_number", "body", m.AccountNumber); err != nil {
		return err
	}

	if err := validate.Pattern("debtor_party"+"."+"account_number", "body", string(*m.AccountNumber), `^[A-Z0-9]{6,34}$`); err != nil {
		return err
	}

	return nil
}

func (m *ClaimAttributesDebtorParty) validateBankID(formats strfmt.Registry) error {

	if err := validate.Required("debtor_party"+"."+"bank_id", "body", m.BankID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimAttributesDebtorParty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimAttributesDebtorParty) UnmarshalBinary(b []byte) error {
	var res ClaimAttributesDebtorParty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimAttributesDebtorParty) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ClaimAttributesOriginalInstruction claim attributes original instruction
// swagger:model ClaimAttributesOriginalInstruction
type ClaimAttributesOriginalInstruction struct {

	// clearing id
	// Required: true
	// Pattern: ^[0-9]{6}$
	ClearingID *string `json:"clearing_id"`

	// reference
	// Required: true
	Reference *string `json:"reference"`
}

func ClaimAttributesOriginalInstructionWithDefaults(defaults client.Defaults) *ClaimAttributesOriginalInstruction {
	return &ClaimAttributesOriginalInstruction{

		ClearingID: defaults.GetStringPtr("ClaimAttributesOriginalInstruction", "clearing_id"),

		Reference: defaults.GetStringPtr("ClaimAttributesOriginalInstruction", "reference"),
	}
}

func (m *ClaimAttributesOriginalInstruction) WithClearingID(clearingID string) *ClaimAttributesOriginalInstruction {

	m.ClearingID = &clearingID

	return m
}

func (m *ClaimAttributesOriginalInstruction) WithoutClearingID() *ClaimAttributesOriginalInstruction {
	m.ClearingID = nil
	return m
}

func (m *ClaimAttributesOriginalInstruction) WithReference(reference string) *ClaimAttributesOriginalInstruction {

	m.Reference = &reference

	return m
}

func (m *ClaimAttributesOriginalInstruction) WithoutReference() *ClaimAttributesOriginalInstruction {
	m.Reference = nil
	return m
}

// Validate validates this claim attributes original instruction
func (m *ClaimAttributesOriginalInstruction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClearingID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReference(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimAttributesOriginalInstruction) validateClearingID(formats strfmt.Registry) error {

	if err := validate.Required("original_instruction"+"."+"clearing_id", "body", m.ClearingID); err != nil {
		return err
	}

	if err := validate.Pattern("original_instruction"+"."+"clearing_id", "body", string(*m.ClearingID), `^[0-9]{6}$`); err != nil {
		return err
	}

	return nil
}

func (m *ClaimAttributesOriginalInstruction) validateReference(formats strfmt.Registry) error {

	if err := validate.Required("original_instruction"+"."+"reference", "body", m.Reference); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimAttributesOriginalInstruction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimAttributesOriginalInstruction) UnmarshalBinary(b []byte) error {
	var res ClaimAttributesOriginalInstruction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimAttributesOriginalInstruction) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
