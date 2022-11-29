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

// PaymentUpdateRelationships payment update relationships
// swagger:model PaymentUpdateRelationships
type PaymentUpdateRelationships struct {

	// beneficiary
	Beneficiary *PaymentUpdateRelationshipsBeneficiary `json:"beneficiary,omitempty"`

	// beneficiary account
	BeneficiaryAccount *PaymentUpdateRelationshipsBeneficiaryAccount `json:"beneficiary_account,omitempty"`

	// debtor
	Debtor *PaymentUpdateRelationshipsDebtor `json:"debtor,omitempty"`

	// debtor account
	DebtorAccount *PaymentUpdateRelationshipsDebtorAccount `json:"debtor_account,omitempty"`

	// forwarding payment
	ForwardingPayment *PaymentUpdateRelationshipsForwardingPayment `json:"forwarding_payment,omitempty"`

	// fx deal
	FxDeal *PaymentUpdateRelationshipsFxDeal `json:"fx_deal,omitempty"`

	// payment admission
	PaymentAdmission *PaymentUpdateRelationshipsPaymentAdmission `json:"payment_admission,omitempty"`

	// payment advice
	PaymentAdvice *PaymentUpdateRelationshipsPaymentAdvice `json:"payment_advice,omitempty"`

	// payment recall
	PaymentRecall *PaymentUpdateRelationshipsPaymentRecall `json:"payment_recall,omitempty"`

	// payment return
	PaymentReturn *PaymentUpdateRelationshipsPaymentReturn `json:"payment_return,omitempty"`

	// payment reversal
	PaymentReversal *PaymentUpdateRelationshipsPaymentReversal `json:"payment_reversal,omitempty"`

	// payment submission
	PaymentSubmission *PaymentUpdateRelationshipsPaymentSubmission `json:"payment_submission,omitempty"`
}

func PaymentUpdateRelationshipsWithDefaults(defaults client.Defaults) *PaymentUpdateRelationships {
	return &PaymentUpdateRelationships{

		Beneficiary: PaymentUpdateRelationshipsBeneficiaryWithDefaults(defaults),

		BeneficiaryAccount: PaymentUpdateRelationshipsBeneficiaryAccountWithDefaults(defaults),

		Debtor: PaymentUpdateRelationshipsDebtorWithDefaults(defaults),

		DebtorAccount: PaymentUpdateRelationshipsDebtorAccountWithDefaults(defaults),

		ForwardingPayment: PaymentUpdateRelationshipsForwardingPaymentWithDefaults(defaults),

		FxDeal: PaymentUpdateRelationshipsFxDealWithDefaults(defaults),

		PaymentAdmission: PaymentUpdateRelationshipsPaymentAdmissionWithDefaults(defaults),

		PaymentAdvice: PaymentUpdateRelationshipsPaymentAdviceWithDefaults(defaults),

		PaymentRecall: PaymentUpdateRelationshipsPaymentRecallWithDefaults(defaults),

		PaymentReturn: PaymentUpdateRelationshipsPaymentReturnWithDefaults(defaults),

		PaymentReversal: PaymentUpdateRelationshipsPaymentReversalWithDefaults(defaults),

		PaymentSubmission: PaymentUpdateRelationshipsPaymentSubmissionWithDefaults(defaults),
	}
}

func (m *PaymentUpdateRelationships) WithBeneficiary(beneficiary PaymentUpdateRelationshipsBeneficiary) *PaymentUpdateRelationships {

	m.Beneficiary = &beneficiary

	return m
}

func (m *PaymentUpdateRelationships) WithoutBeneficiary() *PaymentUpdateRelationships {
	m.Beneficiary = nil
	return m
}

func (m *PaymentUpdateRelationships) WithBeneficiaryAccount(beneficiaryAccount PaymentUpdateRelationshipsBeneficiaryAccount) *PaymentUpdateRelationships {

	m.BeneficiaryAccount = &beneficiaryAccount

	return m
}

func (m *PaymentUpdateRelationships) WithoutBeneficiaryAccount() *PaymentUpdateRelationships {
	m.BeneficiaryAccount = nil
	return m
}

func (m *PaymentUpdateRelationships) WithDebtor(debtor PaymentUpdateRelationshipsDebtor) *PaymentUpdateRelationships {

	m.Debtor = &debtor

	return m
}

func (m *PaymentUpdateRelationships) WithoutDebtor() *PaymentUpdateRelationships {
	m.Debtor = nil
	return m
}

func (m *PaymentUpdateRelationships) WithDebtorAccount(debtorAccount PaymentUpdateRelationshipsDebtorAccount) *PaymentUpdateRelationships {

	m.DebtorAccount = &debtorAccount

	return m
}

func (m *PaymentUpdateRelationships) WithoutDebtorAccount() *PaymentUpdateRelationships {
	m.DebtorAccount = nil
	return m
}

func (m *PaymentUpdateRelationships) WithForwardingPayment(forwardingPayment PaymentUpdateRelationshipsForwardingPayment) *PaymentUpdateRelationships {

	m.ForwardingPayment = &forwardingPayment

	return m
}

func (m *PaymentUpdateRelationships) WithoutForwardingPayment() *PaymentUpdateRelationships {
	m.ForwardingPayment = nil
	return m
}

func (m *PaymentUpdateRelationships) WithFxDeal(fxDeal PaymentUpdateRelationshipsFxDeal) *PaymentUpdateRelationships {

	m.FxDeal = &fxDeal

	return m
}

func (m *PaymentUpdateRelationships) WithoutFxDeal() *PaymentUpdateRelationships {
	m.FxDeal = nil
	return m
}

func (m *PaymentUpdateRelationships) WithPaymentAdmission(paymentAdmission PaymentUpdateRelationshipsPaymentAdmission) *PaymentUpdateRelationships {

	m.PaymentAdmission = &paymentAdmission

	return m
}

func (m *PaymentUpdateRelationships) WithoutPaymentAdmission() *PaymentUpdateRelationships {
	m.PaymentAdmission = nil
	return m
}

func (m *PaymentUpdateRelationships) WithPaymentAdvice(paymentAdvice PaymentUpdateRelationshipsPaymentAdvice) *PaymentUpdateRelationships {

	m.PaymentAdvice = &paymentAdvice

	return m
}

func (m *PaymentUpdateRelationships) WithoutPaymentAdvice() *PaymentUpdateRelationships {
	m.PaymentAdvice = nil
	return m
}

func (m *PaymentUpdateRelationships) WithPaymentRecall(paymentRecall PaymentUpdateRelationshipsPaymentRecall) *PaymentUpdateRelationships {

	m.PaymentRecall = &paymentRecall

	return m
}

func (m *PaymentUpdateRelationships) WithoutPaymentRecall() *PaymentUpdateRelationships {
	m.PaymentRecall = nil
	return m
}

func (m *PaymentUpdateRelationships) WithPaymentReturn(paymentReturn PaymentUpdateRelationshipsPaymentReturn) *PaymentUpdateRelationships {

	m.PaymentReturn = &paymentReturn

	return m
}

func (m *PaymentUpdateRelationships) WithoutPaymentReturn() *PaymentUpdateRelationships {
	m.PaymentReturn = nil
	return m
}

func (m *PaymentUpdateRelationships) WithPaymentReversal(paymentReversal PaymentUpdateRelationshipsPaymentReversal) *PaymentUpdateRelationships {

	m.PaymentReversal = &paymentReversal

	return m
}

func (m *PaymentUpdateRelationships) WithoutPaymentReversal() *PaymentUpdateRelationships {
	m.PaymentReversal = nil
	return m
}

func (m *PaymentUpdateRelationships) WithPaymentSubmission(paymentSubmission PaymentUpdateRelationshipsPaymentSubmission) *PaymentUpdateRelationships {

	m.PaymentSubmission = &paymentSubmission

	return m
}

func (m *PaymentUpdateRelationships) WithoutPaymentSubmission() *PaymentUpdateRelationships {
	m.PaymentSubmission = nil
	return m
}

// Validate validates this payment update relationships
func (m *PaymentUpdateRelationships) Validate(formats strfmt.Registry) error {
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

	if err := m.validateForwardingPayment(formats); err != nil {
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

func (m *PaymentUpdateRelationships) validateBeneficiary(formats strfmt.Registry) error {

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

func (m *PaymentUpdateRelationships) validateBeneficiaryAccount(formats strfmt.Registry) error {

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

func (m *PaymentUpdateRelationships) validateDebtor(formats strfmt.Registry) error {

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

func (m *PaymentUpdateRelationships) validateDebtorAccount(formats strfmt.Registry) error {

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

func (m *PaymentUpdateRelationships) validateForwardingPayment(formats strfmt.Registry) error {

	if swag.IsZero(m.ForwardingPayment) { // not required
		return nil
	}

	if m.ForwardingPayment != nil {
		if err := m.ForwardingPayment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("forwarding_payment")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentUpdateRelationships) validateFxDeal(formats strfmt.Registry) error {

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

func (m *PaymentUpdateRelationships) validatePaymentAdmission(formats strfmt.Registry) error {

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

func (m *PaymentUpdateRelationships) validatePaymentAdvice(formats strfmt.Registry) error {

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

func (m *PaymentUpdateRelationships) validatePaymentRecall(formats strfmt.Registry) error {

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

func (m *PaymentUpdateRelationships) validatePaymentReturn(formats strfmt.Registry) error {

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

func (m *PaymentUpdateRelationships) validatePaymentReversal(formats strfmt.Registry) error {

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

func (m *PaymentUpdateRelationships) validatePaymentSubmission(formats strfmt.Registry) error {

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
func (m *PaymentUpdateRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentUpdateRelationships) UnmarshalBinary(b []byte) error {
	var res PaymentUpdateRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentUpdateRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentUpdateRelationshipsBeneficiary payment update relationships beneficiary
// swagger:model PaymentUpdateRelationshipsBeneficiary
type PaymentUpdateRelationshipsBeneficiary struct {

	// Array of beneficiary resources related to the payment
	Data []*RelationshipData `json:"data"`
}

func PaymentUpdateRelationshipsBeneficiaryWithDefaults(defaults client.Defaults) *PaymentUpdateRelationshipsBeneficiary {
	return &PaymentUpdateRelationshipsBeneficiary{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *PaymentUpdateRelationshipsBeneficiary) WithData(data []*RelationshipData) *PaymentUpdateRelationshipsBeneficiary {

	m.Data = data

	return m
}

// Validate validates this payment update relationships beneficiary
func (m *PaymentUpdateRelationshipsBeneficiary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentUpdateRelationshipsBeneficiary) validateData(formats strfmt.Registry) error {

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
func (m *PaymentUpdateRelationshipsBeneficiary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentUpdateRelationshipsBeneficiary) UnmarshalBinary(b []byte) error {
	var res PaymentUpdateRelationshipsBeneficiary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentUpdateRelationshipsBeneficiary) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentUpdateRelationshipsBeneficiaryAccount payment update relationships beneficiary account
// swagger:model PaymentUpdateRelationshipsBeneficiaryAccount
type PaymentUpdateRelationshipsBeneficiaryAccount struct {

	// Array of beneficiary account resources related to the payment
	Data []*RelationshipData `json:"data"`
}

func PaymentUpdateRelationshipsBeneficiaryAccountWithDefaults(defaults client.Defaults) *PaymentUpdateRelationshipsBeneficiaryAccount {
	return &PaymentUpdateRelationshipsBeneficiaryAccount{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *PaymentUpdateRelationshipsBeneficiaryAccount) WithData(data []*RelationshipData) *PaymentUpdateRelationshipsBeneficiaryAccount {

	m.Data = data

	return m
}

// Validate validates this payment update relationships beneficiary account
func (m *PaymentUpdateRelationshipsBeneficiaryAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentUpdateRelationshipsBeneficiaryAccount) validateData(formats strfmt.Registry) error {

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
func (m *PaymentUpdateRelationshipsBeneficiaryAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentUpdateRelationshipsBeneficiaryAccount) UnmarshalBinary(b []byte) error {
	var res PaymentUpdateRelationshipsBeneficiaryAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentUpdateRelationshipsBeneficiaryAccount) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentUpdateRelationshipsDebtor payment update relationships debtor
// swagger:model PaymentUpdateRelationshipsDebtor
type PaymentUpdateRelationshipsDebtor struct {

	// Array of debtor resources related to the payment
	Data []*RelationshipData `json:"data"`
}

func PaymentUpdateRelationshipsDebtorWithDefaults(defaults client.Defaults) *PaymentUpdateRelationshipsDebtor {
	return &PaymentUpdateRelationshipsDebtor{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *PaymentUpdateRelationshipsDebtor) WithData(data []*RelationshipData) *PaymentUpdateRelationshipsDebtor {

	m.Data = data

	return m
}

// Validate validates this payment update relationships debtor
func (m *PaymentUpdateRelationshipsDebtor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentUpdateRelationshipsDebtor) validateData(formats strfmt.Registry) error {

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
func (m *PaymentUpdateRelationshipsDebtor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentUpdateRelationshipsDebtor) UnmarshalBinary(b []byte) error {
	var res PaymentUpdateRelationshipsDebtor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentUpdateRelationshipsDebtor) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentUpdateRelationshipsDebtorAccount payment update relationships debtor account
// swagger:model PaymentUpdateRelationshipsDebtorAccount
type PaymentUpdateRelationshipsDebtorAccount struct {

	// Array of debtor account resources related to the payment
	Data []*RelationshipData `json:"data"`
}

func PaymentUpdateRelationshipsDebtorAccountWithDefaults(defaults client.Defaults) *PaymentUpdateRelationshipsDebtorAccount {
	return &PaymentUpdateRelationshipsDebtorAccount{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *PaymentUpdateRelationshipsDebtorAccount) WithData(data []*RelationshipData) *PaymentUpdateRelationshipsDebtorAccount {

	m.Data = data

	return m
}

// Validate validates this payment update relationships debtor account
func (m *PaymentUpdateRelationshipsDebtorAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentUpdateRelationshipsDebtorAccount) validateData(formats strfmt.Registry) error {

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
func (m *PaymentUpdateRelationshipsDebtorAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentUpdateRelationshipsDebtorAccount) UnmarshalBinary(b []byte) error {
	var res PaymentUpdateRelationshipsDebtorAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentUpdateRelationshipsDebtorAccount) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentUpdateRelationshipsForwardingPayment payment update relationships forwarding payment
// swagger:model PaymentUpdateRelationshipsForwardingPayment
type PaymentUpdateRelationshipsForwardingPayment struct {

	// ID of the Outbound Payment created to forward the payment.
	Data []*RelationshipData `json:"data"`
}

func PaymentUpdateRelationshipsForwardingPaymentWithDefaults(defaults client.Defaults) *PaymentUpdateRelationshipsForwardingPayment {
	return &PaymentUpdateRelationshipsForwardingPayment{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *PaymentUpdateRelationshipsForwardingPayment) WithData(data []*RelationshipData) *PaymentUpdateRelationshipsForwardingPayment {

	m.Data = data

	return m
}

// Validate validates this payment update relationships forwarding payment
func (m *PaymentUpdateRelationshipsForwardingPayment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentUpdateRelationshipsForwardingPayment) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("forwarding_payment" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentUpdateRelationshipsForwardingPayment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentUpdateRelationshipsForwardingPayment) UnmarshalBinary(b []byte) error {
	var res PaymentUpdateRelationshipsForwardingPayment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentUpdateRelationshipsForwardingPayment) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentUpdateRelationshipsFxDeal payment update relationships fx deal
// swagger:model PaymentUpdateRelationshipsFxDeal
type PaymentUpdateRelationshipsFxDeal struct {

	// Array of FXDeal resources related to the payment
	Data []*RelationshipData `json:"data"`
}

func PaymentUpdateRelationshipsFxDealWithDefaults(defaults client.Defaults) *PaymentUpdateRelationshipsFxDeal {
	return &PaymentUpdateRelationshipsFxDeal{

		Data: make([]*RelationshipData, 0),
	}
}

func (m *PaymentUpdateRelationshipsFxDeal) WithData(data []*RelationshipData) *PaymentUpdateRelationshipsFxDeal {

	m.Data = data

	return m
}

// Validate validates this payment update relationships fx deal
func (m *PaymentUpdateRelationshipsFxDeal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentUpdateRelationshipsFxDeal) validateData(formats strfmt.Registry) error {

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
func (m *PaymentUpdateRelationshipsFxDeal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentUpdateRelationshipsFxDeal) UnmarshalBinary(b []byte) error {
	var res PaymentUpdateRelationshipsFxDeal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentUpdateRelationshipsFxDeal) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentUpdateRelationshipsPaymentAdmission payment update relationships payment admission
// swagger:model PaymentUpdateRelationshipsPaymentAdmission
type PaymentUpdateRelationshipsPaymentAdmission struct {

	// Array of Payment Admission resources related to the payment
	Data []*PaymentAdmission `json:"data"`
}

func PaymentUpdateRelationshipsPaymentAdmissionWithDefaults(defaults client.Defaults) *PaymentUpdateRelationshipsPaymentAdmission {
	return &PaymentUpdateRelationshipsPaymentAdmission{

		Data: make([]*PaymentAdmission, 0),
	}
}

func (m *PaymentUpdateRelationshipsPaymentAdmission) WithData(data []*PaymentAdmission) *PaymentUpdateRelationshipsPaymentAdmission {

	m.Data = data

	return m
}

// Validate validates this payment update relationships payment admission
func (m *PaymentUpdateRelationshipsPaymentAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentUpdateRelationshipsPaymentAdmission) validateData(formats strfmt.Registry) error {

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
func (m *PaymentUpdateRelationshipsPaymentAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentUpdateRelationshipsPaymentAdmission) UnmarshalBinary(b []byte) error {
	var res PaymentUpdateRelationshipsPaymentAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentUpdateRelationshipsPaymentAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentUpdateRelationshipsPaymentAdvice payment update relationships payment advice
// swagger:model PaymentUpdateRelationshipsPaymentAdvice
type PaymentUpdateRelationshipsPaymentAdvice struct {

	// Array of Payment Advice resources related to the payment
	Data []*PaymentAdvice `json:"data"`
}

func PaymentUpdateRelationshipsPaymentAdviceWithDefaults(defaults client.Defaults) *PaymentUpdateRelationshipsPaymentAdvice {
	return &PaymentUpdateRelationshipsPaymentAdvice{

		Data: make([]*PaymentAdvice, 0),
	}
}

func (m *PaymentUpdateRelationshipsPaymentAdvice) WithData(data []*PaymentAdvice) *PaymentUpdateRelationshipsPaymentAdvice {

	m.Data = data

	return m
}

// Validate validates this payment update relationships payment advice
func (m *PaymentUpdateRelationshipsPaymentAdvice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentUpdateRelationshipsPaymentAdvice) validateData(formats strfmt.Registry) error {

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
func (m *PaymentUpdateRelationshipsPaymentAdvice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentUpdateRelationshipsPaymentAdvice) UnmarshalBinary(b []byte) error {
	var res PaymentUpdateRelationshipsPaymentAdvice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentUpdateRelationshipsPaymentAdvice) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentUpdateRelationshipsPaymentRecall payment update relationships payment recall
// swagger:model PaymentUpdateRelationshipsPaymentRecall
type PaymentUpdateRelationshipsPaymentRecall struct {

	// Array of Payment Recall resources related to the payment
	Data []*Recall `json:"data"`
}

func PaymentUpdateRelationshipsPaymentRecallWithDefaults(defaults client.Defaults) *PaymentUpdateRelationshipsPaymentRecall {
	return &PaymentUpdateRelationshipsPaymentRecall{

		Data: make([]*Recall, 0),
	}
}

func (m *PaymentUpdateRelationshipsPaymentRecall) WithData(data []*Recall) *PaymentUpdateRelationshipsPaymentRecall {

	m.Data = data

	return m
}

// Validate validates this payment update relationships payment recall
func (m *PaymentUpdateRelationshipsPaymentRecall) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentUpdateRelationshipsPaymentRecall) validateData(formats strfmt.Registry) error {

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
func (m *PaymentUpdateRelationshipsPaymentRecall) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentUpdateRelationshipsPaymentRecall) UnmarshalBinary(b []byte) error {
	var res PaymentUpdateRelationshipsPaymentRecall
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentUpdateRelationshipsPaymentRecall) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentUpdateRelationshipsPaymentReturn payment update relationships payment return
// swagger:model PaymentUpdateRelationshipsPaymentReturn
type PaymentUpdateRelationshipsPaymentReturn struct {

	// Array of Payment Return resources related to the payment
	Data []*ReturnPayment `json:"data"`
}

func PaymentUpdateRelationshipsPaymentReturnWithDefaults(defaults client.Defaults) *PaymentUpdateRelationshipsPaymentReturn {
	return &PaymentUpdateRelationshipsPaymentReturn{

		Data: make([]*ReturnPayment, 0),
	}
}

func (m *PaymentUpdateRelationshipsPaymentReturn) WithData(data []*ReturnPayment) *PaymentUpdateRelationshipsPaymentReturn {

	m.Data = data

	return m
}

// Validate validates this payment update relationships payment return
func (m *PaymentUpdateRelationshipsPaymentReturn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentUpdateRelationshipsPaymentReturn) validateData(formats strfmt.Registry) error {

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
func (m *PaymentUpdateRelationshipsPaymentReturn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentUpdateRelationshipsPaymentReturn) UnmarshalBinary(b []byte) error {
	var res PaymentUpdateRelationshipsPaymentReturn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentUpdateRelationshipsPaymentReturn) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentUpdateRelationshipsPaymentReversal payment update relationships payment reversal
// swagger:model PaymentUpdateRelationshipsPaymentReversal
type PaymentUpdateRelationshipsPaymentReversal struct {

	// Array of Payment Reversal resources related to the payment
	Data []*ReversalPayment `json:"data"`
}

func PaymentUpdateRelationshipsPaymentReversalWithDefaults(defaults client.Defaults) *PaymentUpdateRelationshipsPaymentReversal {
	return &PaymentUpdateRelationshipsPaymentReversal{

		Data: make([]*ReversalPayment, 0),
	}
}

func (m *PaymentUpdateRelationshipsPaymentReversal) WithData(data []*ReversalPayment) *PaymentUpdateRelationshipsPaymentReversal {

	m.Data = data

	return m
}

// Validate validates this payment update relationships payment reversal
func (m *PaymentUpdateRelationshipsPaymentReversal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentUpdateRelationshipsPaymentReversal) validateData(formats strfmt.Registry) error {

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
func (m *PaymentUpdateRelationshipsPaymentReversal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentUpdateRelationshipsPaymentReversal) UnmarshalBinary(b []byte) error {
	var res PaymentUpdateRelationshipsPaymentReversal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentUpdateRelationshipsPaymentReversal) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentUpdateRelationshipsPaymentSubmission payment update relationships payment submission
// swagger:model PaymentUpdateRelationshipsPaymentSubmission
type PaymentUpdateRelationshipsPaymentSubmission struct {

	// Array of Payment Submission resources related to the payment
	Data []*PaymentSubmission `json:"data"`
}

func PaymentUpdateRelationshipsPaymentSubmissionWithDefaults(defaults client.Defaults) *PaymentUpdateRelationshipsPaymentSubmission {
	return &PaymentUpdateRelationshipsPaymentSubmission{

		Data: make([]*PaymentSubmission, 0),
	}
}

func (m *PaymentUpdateRelationshipsPaymentSubmission) WithData(data []*PaymentSubmission) *PaymentUpdateRelationshipsPaymentSubmission {

	m.Data = data

	return m
}

// Validate validates this payment update relationships payment submission
func (m *PaymentUpdateRelationshipsPaymentSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentUpdateRelationshipsPaymentSubmission) validateData(formats strfmt.Registry) error {

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
func (m *PaymentUpdateRelationshipsPaymentSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentUpdateRelationshipsPaymentSubmission) UnmarshalBinary(b []byte) error {
	var res PaymentUpdateRelationshipsPaymentSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentUpdateRelationshipsPaymentSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
