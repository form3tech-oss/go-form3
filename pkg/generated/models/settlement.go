// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Settlement Specifies the details on how the settlement of the transaction between the instructing agent and the instructed agent is completed
// swagger:model Settlement
type Settlement struct {

	// account number
	AccountNumber *string `json:"account_number,omitempty"`

	// account number code
	// Enum: [IBAN BBAN]
	AccountNumberCode AccountNumberCode `json:"account_number_code,omitempty"`

	// bank id
	BankID *string `json:"bank_id,omitempty"`

	// bank id code
	// Enum: [GBDSC]
	BankIDCode BankIDCode `json:"bank_id_code,omitempty"`

	// Method used to settle the payment instruction. Acceptable Values for SEPA: CLRG. Acceptable Values for SWIFT: INDA (settled by Instructed Agent), INGA (Settled by Instructing Agent), COVE (Cover Payment)
	// Enum: [CLRG COVE INGA INDA]
	Method *string `json:"method,omitempty"`
}

func SettlementWithDefaults(defaults client.Defaults) *Settlement {
	return &Settlement{

		AccountNumber: defaults.GetStringPtr("Settlement", "account_number"),

		// TODO AccountNumberCode: AccountNumberCode,

		BankID: defaults.GetStringPtr("Settlement", "bank_id"),

		// TODO BankIDCode: BankIDCode,

		Method: defaults.GetStringPtr("Settlement", "method"),
	}
}

func (m *Settlement) WithAccountNumber(accountNumber string) *Settlement {

	m.AccountNumber = &accountNumber

	return m
}

func (m *Settlement) WithoutAccountNumber() *Settlement {
	m.AccountNumber = nil
	return m
}

func (m *Settlement) WithAccountNumberCode(accountNumberCode AccountNumberCode) *Settlement {

	m.AccountNumberCode = accountNumberCode

	return m
}

func (m *Settlement) WithBankID(bankID string) *Settlement {

	m.BankID = &bankID

	return m
}

func (m *Settlement) WithoutBankID() *Settlement {
	m.BankID = nil
	return m
}

func (m *Settlement) WithBankIDCode(bankIDCode BankIDCode) *Settlement {

	m.BankIDCode = bankIDCode

	return m
}

func (m *Settlement) WithMethod(method string) *Settlement {

	m.Method = &method

	return m
}

func (m *Settlement) WithoutMethod() *Settlement {
	m.Method = nil
	return m
}

// Validate validates this settlement
func (m *Settlement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumberCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankIDCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var settlementTypeAccountNumberCodePropEnum []interface{}

func init() {
	var res []AccountNumberCode
	if err := json.Unmarshal([]byte(`["IBAN","BBAN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		settlementTypeAccountNumberCodePropEnum = append(settlementTypeAccountNumberCodePropEnum, v)
	}
}

const (

	// SettlementAccountNumberCodeIBAN captures enum value "IBAN"
	SettlementAccountNumberCodeIBAN AccountNumberCode = "IBAN"

	// SettlementAccountNumberCodeBBAN captures enum value "BBAN"
	SettlementAccountNumberCodeBBAN AccountNumberCode = "BBAN"
)

// prop value enum
func (m *Settlement) validateAccountNumberCodeEnum(path, location string, value AccountNumberCode) error {
	if err := validate.Enum(path, location, value, settlementTypeAccountNumberCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Settlement) validateAccountNumberCode(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumberCode) { // not required
		return nil
	}

	if err := m.AccountNumberCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("account_number_code")
		}
		return err
	}

	return nil
}

var settlementTypeBankIDCodePropEnum []interface{}

func init() {
	var res []BankIDCode
	if err := json.Unmarshal([]byte(`["GBDSC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		settlementTypeBankIDCodePropEnum = append(settlementTypeBankIDCodePropEnum, v)
	}
}

const (

	// SettlementBankIDCodeGBDSC captures enum value "GBDSC"
	SettlementBankIDCodeGBDSC BankIDCode = "GBDSC"
)

// prop value enum
func (m *Settlement) validateBankIDCodeEnum(path, location string, value BankIDCode) error {
	if err := validate.Enum(path, location, value, settlementTypeBankIDCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Settlement) validateBankIDCode(formats strfmt.Registry) error {

	if swag.IsZero(m.BankIDCode) { // not required
		return nil
	}

	if err := m.BankIDCode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("bank_id_code")
		}
		return err
	}

	return nil
}

var settlementTypeMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CLRG","COVE","INGA","INDA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		settlementTypeMethodPropEnum = append(settlementTypeMethodPropEnum, v)
	}
}

const (

	// SettlementMethodCLRG captures enum value "CLRG"
	SettlementMethodCLRG string = "CLRG"

	// SettlementMethodCOVE captures enum value "COVE"
	SettlementMethodCOVE string = "COVE"

	// SettlementMethodINGA captures enum value "INGA"
	SettlementMethodINGA string = "INGA"

	// SettlementMethodINDA captures enum value "INDA"
	SettlementMethodINDA string = "INDA"
)

// prop value enum
func (m *Settlement) validateMethodEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, settlementTypeMethodPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Settlement) validateMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.Method) { // not required
		return nil
	}

	// value enum
	if err := m.validateMethodEnum("method", "body", *m.Method); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Settlement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Settlement) UnmarshalBinary(b []byte) error {
	var res Settlement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *Settlement) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
