// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NewPaymentRelationships new payment relationships
// swagger:model NewPaymentRelationships
type NewPaymentRelationships struct {

	// beneficiary
	Beneficiary *NewPaymentRelationshipsBeneficiary `json:"beneficiary,omitempty"`

	// beneficiary account
	BeneficiaryAccount *NewPaymentRelationshipsBeneficiaryAccount `json:"beneficiary_account,omitempty"`

	// debtor
	Debtor *NewPaymentRelationshipsDebtor `json:"debtor,omitempty"`

	// debtor account
	DebtorAccount *NewPaymentRelationshipsDebtorAccount `json:"debtor_account,omitempty"`

	// forwarded payment
	ForwardedPayment *NewPaymentRelationshipsForwardedPayment `json:"forwarded_payment,omitempty"`

	// fx deal
	FxDeal *NewPaymentRelationshipsFxDeal `json:"fx_deal,omitempty"`

	// payment admission
	PaymentAdmission *NewPaymentRelationshipsPaymentAdmission `json:"payment_admission,omitempty"`

	// payment advice
	PaymentAdvice *NewPaymentRelationshipsPaymentAdvice `json:"payment_advice,omitempty"`

	// payment recall
	PaymentRecall *NewPaymentRelationshipsPaymentRecall `json:"payment_recall,omitempty"`

	// payment return
	PaymentReturn *NewPaymentRelationshipsPaymentReturn `json:"payment_return,omitempty"`

	// payment reversal
	PaymentReversal *NewPaymentRelationshipsPaymentReversal `json:"payment_reversal,omitempty"`

	// payment submission
	PaymentSubmission *NewPaymentRelationshipsPaymentSubmission `json:"payment_submission,omitempty"`
}

func NewPaymentRelationshipsWithDefaults(defaults client.Defaults) *NewPaymentRelationships {
	return &NewPaymentRelationships{

		Beneficiary: NewPaymentRelationshipsBeneficiaryWithDefaults(defaults),

		BeneficiaryAccount: NewPaymentRelationshipsBeneficiaryAccountWithDefaults(defaults),

		Debtor: NewPaymentRelationshipsDebtorWithDefaults(defaults),

		DebtorAccount: NewPaymentRelationshipsDebtorAccountWithDefaults(defaults),

		ForwardedPayment: NewPaymentRelationshipsForwardedPaymentWithDefaults(defaults),

		FxDeal: NewPaymentRelationshipsFxDealWithDefaults(defaults),

		PaymentAdmission: NewPaymentRelationshipsPaymentAdmissionWithDefaults(defaults),

		PaymentAdvice: NewPaymentRelationshipsPaymentAdviceWithDefaults(defaults),

		PaymentRecall: NewPaymentRelationshipsPaymentRecallWithDefaults(defaults),

		PaymentReturn: NewPaymentRelationshipsPaymentReturnWithDefaults(defaults),

		PaymentReversal: NewPaymentRelationshipsPaymentReversalWithDefaults(defaults),

		PaymentSubmission: NewPaymentRelationshipsPaymentSubmissionWithDefaults(defaults),
	}
}

func (m *NewPaymentRelationships) WithBeneficiary(beneficiary NewPaymentRelationshipsBeneficiary) *NewPaymentRelationships {

	m.Beneficiary = &beneficiary

	return m
}

func (m *NewPaymentRelationships) WithoutBeneficiary() *NewPaymentRelationships {
	m.Beneficiary = nil
	return m
}

func (m *NewPaymentRelationships) WithBeneficiaryAccount(beneficiaryAccount NewPaymentRelationshipsBeneficiaryAccount) *NewPaymentRelationships {

	m.BeneficiaryAccount = &beneficiaryAccount

	return m
}

func (m *NewPaymentRelationships) WithoutBeneficiaryAccount() *NewPaymentRelationships {
	m.BeneficiaryAccount = nil
	return m
}

func (m *NewPaymentRelationships) WithDebtor(debtor NewPaymentRelationshipsDebtor) *NewPaymentRelationships {

	m.Debtor = &debtor

	return m
}

func (m *NewPaymentRelationships) WithoutDebtor() *NewPaymentRelationships {
	m.Debtor = nil
	return m
}

func (m *NewPaymentRelationships) WithDebtorAccount(debtorAccount NewPaymentRelationshipsDebtorAccount) *NewPaymentRelationships {

	m.DebtorAccount = &debtorAccount

	return m
}

func (m *NewPaymentRelationships) WithoutDebtorAccount() *NewPaymentRelationships {
	m.DebtorAccount = nil
	return m
}

func (m *NewPaymentRelationships) WithForwardedPayment(forwardedPayment NewPaymentRelationshipsForwardedPayment) *NewPaymentRelationships {

	m.ForwardedPayment = &forwardedPayment

	return m
}

func (m *NewPaymentRelationships) WithoutForwardedPayment() *NewPaymentRelationships {
	m.ForwardedPayment = nil
	return m
}

func (m *NewPaymentRelationships) WithFxDeal(fxDeal NewPaymentRelationshipsFxDeal) *NewPaymentRelationships {

	m.FxDeal = &fxDeal

	return m
}

func (m *NewPaymentRelationships) WithoutFxDeal() *NewPaymentRelationships {
	m.FxDeal = nil
	return m
}

func (m *NewPaymentRelationships) WithPaymentAdmission(paymentAdmission NewPaymentRelationshipsPaymentAdmission) *NewPaymentRelationships {

	m.PaymentAdmission = &paymentAdmission

	return m
}

func (m *NewPaymentRelationships) WithoutPaymentAdmission() *NewPaymentRelationships {
	m.PaymentAdmission = nil
	return m
}

func (m *NewPaymentRelationships) WithPaymentAdvice(paymentAdvice NewPaymentRelationshipsPaymentAdvice) *NewPaymentRelationships {

	m.PaymentAdvice = &paymentAdvice

	return m
}

func (m *NewPaymentRelationships) WithoutPaymentAdvice() *NewPaymentRelationships {
	m.PaymentAdvice = nil
	return m
}

func (m *NewPaymentRelationships) WithPaymentRecall(paymentRecall NewPaymentRelationshipsPaymentRecall) *NewPaymentRelationships {

	m.PaymentRecall = &paymentRecall

	return m
}

func (m *NewPaymentRelationships) WithoutPaymentRecall() *NewPaymentRelationships {
	m.PaymentRecall = nil
	return m
}

func (m *NewPaymentRelationships) WithPaymentReturn(paymentReturn NewPaymentRelationshipsPaymentReturn) *NewPaymentRelationships {

	m.PaymentReturn = &paymentReturn

	return m
}

func (m *NewPaymentRelationships) WithoutPaymentReturn() *NewPaymentRelationships {
	m.PaymentReturn = nil
	return m
}

func (m *NewPaymentRelationships) WithPaymentReversal(paymentReversal NewPaymentRelationshipsPaymentReversal) *NewPaymentRelationships {

	m.PaymentReversal = &paymentReversal

	return m
}

func (m *NewPaymentRelationships) WithoutPaymentReversal() *NewPaymentRelationships {
	m.PaymentReversal = nil
	return m
}

func (m *NewPaymentRelationships) WithPaymentSubmission(paymentSubmission NewPaymentRelationshipsPaymentSubmission) *NewPaymentRelationships {

	m.PaymentSubmission = &paymentSubmission

	return m
}

func (m *NewPaymentRelationships) WithoutPaymentSubmission() *NewPaymentRelationships {
	m.PaymentSubmission = nil
	return m
}

// Validate validates this new payment relationships
func (m *NewPaymentRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBeneficiary(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBeneficiaryAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDebtorAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForwardedPayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFxDeal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentAdmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentAdvice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentRecall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentReturn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentReversal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPaymentRelationships) validateBeneficiary(formats strfmt.Registry) error {

	if swag.IsZero(m.Beneficiary) { // not required
		return nil
	}

	if m.Beneficiary != nil {
		if err := m.Beneficiary.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("beneficiary")
			}
			return err
		}
	}

	return nil
}

func (m *NewPaymentRelationships) validateBeneficiaryAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.BeneficiaryAccount) { // not required
		return nil
	}

	if m.BeneficiaryAccount != nil {
		if err := m.BeneficiaryAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("beneficiary_account")
			}
			return err
		}
	}

	return nil
}

func (m *NewPaymentRelationships) validateDebtor(formats strfmt.Registry) error {

	if swag.IsZero(m.Debtor) { // not required
		return nil
	}

	if m.Debtor != nil {
		if err := m.Debtor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("debtor")
			}
			return err
		}
	}

	return nil
}

func (m *NewPaymentRelationships) validateDebtorAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.DebtorAccount) { // not required
		return nil
	}

	if m.DebtorAccount != nil {
		if err := m.DebtorAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("debtor_account")
			}
			return err
		}
	}

	return nil
}

func (m *NewPaymentRelationships) validateForwardedPayment(formats strfmt.Registry) error {

	if swag.IsZero(m.ForwardedPayment) { // not required
		return nil
	}

	if m.ForwardedPayment != nil {
		if err := m.ForwardedPayment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forwarded_payment")
			}
			return err
		}
	}

	return nil
}

func (m *NewPaymentRelationships) validateFxDeal(formats strfmt.Registry) error {

	if swag.IsZero(m.FxDeal) { // not required
		return nil
	}

	if m.FxDeal != nil {
		if err := m.FxDeal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fx_deal")
			}
			return err
		}
	}

	return nil
}

func (m *NewPaymentRelationships) validatePaymentAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentAdmission) { // not required
		return nil
	}

	if m.PaymentAdmission != nil {
		if err := m.PaymentAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_admission")
			}
			return err
		}
	}

	return nil
}

func (m *NewPaymentRelationships) validatePaymentAdvice(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentAdvice) { // not required
		return nil
	}

	if m.PaymentAdvice != nil {
		if err := m.PaymentAdvice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_advice")
			}
			return err
		}
	}

	return nil
}

func (m *NewPaymentRelationships) validatePaymentRecall(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentRecall) { // not required
		return nil
	}

	if m.PaymentRecall != nil {
		if err := m.PaymentRecall.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_recall")
			}
			return err
		}
	}

	return nil
}

func (m *NewPaymentRelationships) validatePaymentReturn(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentReturn) { // not required
		return nil
	}

	if m.PaymentReturn != nil {
		if err := m.PaymentReturn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_return")
			}
			return err
		}
	}

	return nil
}

func (m *NewPaymentRelationships) validatePaymentReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentReversal) { // not required
		return nil
	}

	if m.PaymentReversal != nil {
		if err := m.PaymentReversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_reversal")
			}
			return err
		}
	}

	return nil
}

func (m *NewPaymentRelationships) validatePaymentSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentSubmission) { // not required
		return nil
	}

	if m.PaymentSubmission != nil {
		if err := m.PaymentSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPaymentRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPaymentRelationships) UnmarshalBinary(b []byte) error {
	var res NewPaymentRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewPaymentRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// NewPaymentRelationshipsBeneficiary new payment relationships beneficiary
// swagger:model NewPaymentRelationshipsBeneficiary
type NewPaymentRelationshipsBeneficiary struct {

	// Array of beneficiary resources related to the payment
	Data []*RelationshipData `json:"data"`
}

func NewPaymentRelationshipsBeneficiaryWithDefaults(defaults client.Defaults) *NewPaymentRelationshipsBeneficiary {
	return &NewPaymentRelationshipsBeneficiary{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *NewPaymentRelationshipsBeneficiary) WithData(data []*RelationshipData) *NewPaymentRelationshipsBeneficiary {

	m.Data = data

	return m
}

// Validate validates this new payment relationships beneficiary
func (m *NewPaymentRelationshipsBeneficiary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPaymentRelationshipsBeneficiary) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("beneficiary" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPaymentRelationshipsBeneficiary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPaymentRelationshipsBeneficiary) UnmarshalBinary(b []byte) error {
	var res NewPaymentRelationshipsBeneficiary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewPaymentRelationshipsBeneficiary) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// NewPaymentRelationshipsBeneficiaryAccount new payment relationships beneficiary account
// swagger:model NewPaymentRelationshipsBeneficiaryAccount
type NewPaymentRelationshipsBeneficiaryAccount struct {

	// Array of beneficiary account resources related to the payment
	Data []*RelationshipData `json:"data"`
}

func NewPaymentRelationshipsBeneficiaryAccountWithDefaults(defaults client.Defaults) *NewPaymentRelationshipsBeneficiaryAccount {
	return &NewPaymentRelationshipsBeneficiaryAccount{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *NewPaymentRelationshipsBeneficiaryAccount) WithData(data []*RelationshipData) *NewPaymentRelationshipsBeneficiaryAccount {

	m.Data = data

	return m
}

// Validate validates this new payment relationships beneficiary account
func (m *NewPaymentRelationshipsBeneficiaryAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPaymentRelationshipsBeneficiaryAccount) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("beneficiary_account" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPaymentRelationshipsBeneficiaryAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPaymentRelationshipsBeneficiaryAccount) UnmarshalBinary(b []byte) error {
	var res NewPaymentRelationshipsBeneficiaryAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewPaymentRelationshipsBeneficiaryAccount) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// NewPaymentRelationshipsDebtor new payment relationships debtor
// swagger:model NewPaymentRelationshipsDebtor
type NewPaymentRelationshipsDebtor struct {

	// Array of debtor resources related to the payment
	Data []*RelationshipData `json:"data"`
}

func NewPaymentRelationshipsDebtorWithDefaults(defaults client.Defaults) *NewPaymentRelationshipsDebtor {
	return &NewPaymentRelationshipsDebtor{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *NewPaymentRelationshipsDebtor) WithData(data []*RelationshipData) *NewPaymentRelationshipsDebtor {

	m.Data = data

	return m
}

// Validate validates this new payment relationships debtor
func (m *NewPaymentRelationshipsDebtor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPaymentRelationshipsDebtor) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("debtor" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPaymentRelationshipsDebtor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPaymentRelationshipsDebtor) UnmarshalBinary(b []byte) error {
	var res NewPaymentRelationshipsDebtor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewPaymentRelationshipsDebtor) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// NewPaymentRelationshipsDebtorAccount new payment relationships debtor account
// swagger:model NewPaymentRelationshipsDebtorAccount
type NewPaymentRelationshipsDebtorAccount struct {

	// Array of debtor account resources related to the payment
	Data []*RelationshipData `json:"data"`
}

func NewPaymentRelationshipsDebtorAccountWithDefaults(defaults client.Defaults) *NewPaymentRelationshipsDebtorAccount {
	return &NewPaymentRelationshipsDebtorAccount{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *NewPaymentRelationshipsDebtorAccount) WithData(data []*RelationshipData) *NewPaymentRelationshipsDebtorAccount {

	m.Data = data

	return m
}

// Validate validates this new payment relationships debtor account
func (m *NewPaymentRelationshipsDebtorAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPaymentRelationshipsDebtorAccount) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("debtor_account" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPaymentRelationshipsDebtorAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPaymentRelationshipsDebtorAccount) UnmarshalBinary(b []byte) error {
	var res NewPaymentRelationshipsDebtorAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewPaymentRelationshipsDebtorAccount) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// NewPaymentRelationshipsForwardedPayment new payment relationships forwarded payment
// swagger:model NewPaymentRelationshipsForwardedPayment
type NewPaymentRelationshipsForwardedPayment struct {

	//  Array of Inbound Payments that triggered the creation of this Forwarding Payment.
	Data []*RelationshipData `json:"data"`
}

func NewPaymentRelationshipsForwardedPaymentWithDefaults(defaults client.Defaults) *NewPaymentRelationshipsForwardedPayment {
	return &NewPaymentRelationshipsForwardedPayment{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *NewPaymentRelationshipsForwardedPayment) WithData(data []*RelationshipData) *NewPaymentRelationshipsForwardedPayment {

	m.Data = data

	return m
}

// Validate validates this new payment relationships forwarded payment
func (m *NewPaymentRelationshipsForwardedPayment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPaymentRelationshipsForwardedPayment) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("forwarded_payment" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPaymentRelationshipsForwardedPayment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPaymentRelationshipsForwardedPayment) UnmarshalBinary(b []byte) error {
	var res NewPaymentRelationshipsForwardedPayment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewPaymentRelationshipsForwardedPayment) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// NewPaymentRelationshipsFxDeal new payment relationships fx deal
// swagger:model NewPaymentRelationshipsFxDeal
type NewPaymentRelationshipsFxDeal struct {

	// Array of FXDeal resources related to the payment
	Data []*RelationshipData `json:"data"`
}

func NewPaymentRelationshipsFxDealWithDefaults(defaults client.Defaults) *NewPaymentRelationshipsFxDeal {
	return &NewPaymentRelationshipsFxDeal{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *NewPaymentRelationshipsFxDeal) WithData(data []*RelationshipData) *NewPaymentRelationshipsFxDeal {

	m.Data = data

	return m
}

// Validate validates this new payment relationships fx deal
func (m *NewPaymentRelationshipsFxDeal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPaymentRelationshipsFxDeal) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fx_deal" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPaymentRelationshipsFxDeal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPaymentRelationshipsFxDeal) UnmarshalBinary(b []byte) error {
	var res NewPaymentRelationshipsFxDeal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewPaymentRelationshipsFxDeal) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// NewPaymentRelationshipsPaymentAdmission new payment relationships payment admission
// swagger:model NewPaymentRelationshipsPaymentAdmission
type NewPaymentRelationshipsPaymentAdmission struct {

	// Array of Payment Admission resources related to the payment
	Data []*PaymentAdmission `json:"data"`
}

func NewPaymentRelationshipsPaymentAdmissionWithDefaults(defaults client.Defaults) *NewPaymentRelationshipsPaymentAdmission {
	return &NewPaymentRelationshipsPaymentAdmission{

		Data: make([]*PaymentAdmission, 0),
	}
}

func (m *NewPaymentRelationshipsPaymentAdmission) WithData(data []*PaymentAdmission) *NewPaymentRelationshipsPaymentAdmission {

	m.Data = data

	return m
}

// Validate validates this new payment relationships payment admission
func (m *NewPaymentRelationshipsPaymentAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPaymentRelationshipsPaymentAdmission) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payment_admission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPaymentRelationshipsPaymentAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPaymentRelationshipsPaymentAdmission) UnmarshalBinary(b []byte) error {
	var res NewPaymentRelationshipsPaymentAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewPaymentRelationshipsPaymentAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// NewPaymentRelationshipsPaymentAdvice new payment relationships payment advice
// swagger:model NewPaymentRelationshipsPaymentAdvice
type NewPaymentRelationshipsPaymentAdvice struct {

	// Array of Payment Advice resources related to the payment
	Data []*PaymentAdvice `json:"data"`
}

func NewPaymentRelationshipsPaymentAdviceWithDefaults(defaults client.Defaults) *NewPaymentRelationshipsPaymentAdvice {
	return &NewPaymentRelationshipsPaymentAdvice{

		Data: make([]*PaymentAdvice, 0),
	}
}

func (m *NewPaymentRelationshipsPaymentAdvice) WithData(data []*PaymentAdvice) *NewPaymentRelationshipsPaymentAdvice {

	m.Data = data

	return m
}

// Validate validates this new payment relationships payment advice
func (m *NewPaymentRelationshipsPaymentAdvice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPaymentRelationshipsPaymentAdvice) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payment_advice" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPaymentRelationshipsPaymentAdvice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPaymentRelationshipsPaymentAdvice) UnmarshalBinary(b []byte) error {
	var res NewPaymentRelationshipsPaymentAdvice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewPaymentRelationshipsPaymentAdvice) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// NewPaymentRelationshipsPaymentRecall new payment relationships payment recall
// swagger:model NewPaymentRelationshipsPaymentRecall
type NewPaymentRelationshipsPaymentRecall struct {

	// Array of Payment Recall resources related to the payment
	Data []*Recall `json:"data"`
}

func NewPaymentRelationshipsPaymentRecallWithDefaults(defaults client.Defaults) *NewPaymentRelationshipsPaymentRecall {
	return &NewPaymentRelationshipsPaymentRecall{

		Data: make([]*Recall, 0),
	}
}

func (m *NewPaymentRelationshipsPaymentRecall) WithData(data []*Recall) *NewPaymentRelationshipsPaymentRecall {

	m.Data = data

	return m
}

// Validate validates this new payment relationships payment recall
func (m *NewPaymentRelationshipsPaymentRecall) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPaymentRelationshipsPaymentRecall) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payment_recall" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPaymentRelationshipsPaymentRecall) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPaymentRelationshipsPaymentRecall) UnmarshalBinary(b []byte) error {
	var res NewPaymentRelationshipsPaymentRecall
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewPaymentRelationshipsPaymentRecall) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// NewPaymentRelationshipsPaymentReturn new payment relationships payment return
// swagger:model NewPaymentRelationshipsPaymentReturn
type NewPaymentRelationshipsPaymentReturn struct {

	// Array of Payment Return resources related to the payment
	Data []*ReturnPayment `json:"data"`
}

func NewPaymentRelationshipsPaymentReturnWithDefaults(defaults client.Defaults) *NewPaymentRelationshipsPaymentReturn {
	return &NewPaymentRelationshipsPaymentReturn{

		Data: make([]*ReturnPayment, 0),
	}
}

func (m *NewPaymentRelationshipsPaymentReturn) WithData(data []*ReturnPayment) *NewPaymentRelationshipsPaymentReturn {

	m.Data = data

	return m
}

// Validate validates this new payment relationships payment return
func (m *NewPaymentRelationshipsPaymentReturn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPaymentRelationshipsPaymentReturn) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payment_return" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPaymentRelationshipsPaymentReturn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPaymentRelationshipsPaymentReturn) UnmarshalBinary(b []byte) error {
	var res NewPaymentRelationshipsPaymentReturn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewPaymentRelationshipsPaymentReturn) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// NewPaymentRelationshipsPaymentReversal new payment relationships payment reversal
// swagger:model NewPaymentRelationshipsPaymentReversal
type NewPaymentRelationshipsPaymentReversal struct {

	// Array of Payment Reversal resources related to the payment
	Data []*ReversalPayment `json:"data"`
}

func NewPaymentRelationshipsPaymentReversalWithDefaults(defaults client.Defaults) *NewPaymentRelationshipsPaymentReversal {
	return &NewPaymentRelationshipsPaymentReversal{

		Data: make([]*ReversalPayment, 0),
	}
}

func (m *NewPaymentRelationshipsPaymentReversal) WithData(data []*ReversalPayment) *NewPaymentRelationshipsPaymentReversal {

	m.Data = data

	return m
}

// Validate validates this new payment relationships payment reversal
func (m *NewPaymentRelationshipsPaymentReversal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPaymentRelationshipsPaymentReversal) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payment_reversal" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPaymentRelationshipsPaymentReversal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPaymentRelationshipsPaymentReversal) UnmarshalBinary(b []byte) error {
	var res NewPaymentRelationshipsPaymentReversal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewPaymentRelationshipsPaymentReversal) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// NewPaymentRelationshipsPaymentSubmission new payment relationships payment submission
// swagger:model NewPaymentRelationshipsPaymentSubmission
type NewPaymentRelationshipsPaymentSubmission struct {

	// Array of Payment Submission resources related to the payment
	Data []*PaymentSubmission `json:"data"`
}

func NewPaymentRelationshipsPaymentSubmissionWithDefaults(defaults client.Defaults) *NewPaymentRelationshipsPaymentSubmission {
	return &NewPaymentRelationshipsPaymentSubmission{

		Data: make([]*PaymentSubmission, 0),
	}
}

func (m *NewPaymentRelationshipsPaymentSubmission) WithData(data []*PaymentSubmission) *NewPaymentRelationshipsPaymentSubmission {

	m.Data = data

	return m
}

// Validate validates this new payment relationships payment submission
func (m *NewPaymentRelationshipsPaymentSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPaymentRelationshipsPaymentSubmission) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("payment_submission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPaymentRelationshipsPaymentSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPaymentRelationshipsPaymentSubmission) UnmarshalBinary(b []byte) error {
	var res NewPaymentRelationshipsPaymentSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewPaymentRelationshipsPaymentSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
