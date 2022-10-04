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

// AccountHoldingEntity Information about the financial institution servicing the account.
// swagger:model AccountHoldingEntity
type AccountHoldingEntity struct {

	// Financial institution address
	BankAddress []string `json:"bank_address,omitempty"`

	// bank id
	// Required: true
	// Pattern: ^[A-Z0-9]{0,16}$
	BankID *string `json:"bank_id"`

	// ISO 20022 code used to identify the type of bank ID being used
	// Required: true
	// Enum: [GBDSC]
	BankIDCode *BankIDCode `json:"bank_id_code"`

	// Financial institution name
	BankName string `json:"bank_name,omitempty"`

	// Identifier of the financial institution which services the account
	BankPartyID string `json:"bank_party_id,omitempty"`
}

func AccountHoldingEntityWithDefaults(defaults client.Defaults) *AccountHoldingEntity {
	return &AccountHoldingEntity{

		BankAddress: make([]string, 0),

		BankID: defaults.GetStringPtr("AccountHoldingEntity", "bank_id"),

		// TODO BankIDCode: BankIDCode,

		BankName: defaults.GetString("AccountHoldingEntity", "bank_name"),

		BankPartyID: defaults.GetString("AccountHoldingEntity", "bank_party_id"),
	}
}

func (m *AccountHoldingEntity) WithBankAddress(bankAddress []string) *AccountHoldingEntity {

	m.BankAddress = bankAddress

	return m
}

func (m *AccountHoldingEntity) WithBankID(bankID string) *AccountHoldingEntity {

	m.BankID = &bankID

	return m
}

func (m *AccountHoldingEntity) WithoutBankID() *AccountHoldingEntity {
	m.BankID = nil
	return m
}

func (m *AccountHoldingEntity) WithBankIDCode(bankIDCode BankIDCode) *AccountHoldingEntity {

	m.BankIDCode = &bankIDCode

	return m
}

func (m *AccountHoldingEntity) WithoutBankIDCode() *AccountHoldingEntity {
	m.BankIDCode = nil
	return m
}

func (m *AccountHoldingEntity) WithBankName(bankName string) *AccountHoldingEntity {

	m.BankName = bankName

	return m
}

func (m *AccountHoldingEntity) WithBankPartyID(bankPartyID string) *AccountHoldingEntity {

	m.BankPartyID = bankPartyID

	return m
}

// Validate validates this account holding entity
func (m *AccountHoldingEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBankID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankIDCode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountHoldingEntity) validateBankID(formats strfmt.Registry) error {

	if err := validate.Required("bank_id", "body", m.BankID); err != nil {
		return err
	}

	if err := validate.Pattern("bank_id", "body", string(*m.BankID), `^[A-Z0-9]{0,16}$`); err != nil {
		return err
	}

	return nil
}

var accountHoldingEntityTypeBankIDCodePropEnum []interface{}

func init() {
	var res []BankIDCode
	if err := json.Unmarshal([]byte(`["GBDSC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountHoldingEntityTypeBankIDCodePropEnum = append(accountHoldingEntityTypeBankIDCodePropEnum, v)
	}
}

const (

	// AccountHoldingEntityBankIDCodeGBDSC captures enum value "GBDSC"
	AccountHoldingEntityBankIDCodeGBDSC BankIDCode = "GBDSC"
)

// prop value enum
func (m *AccountHoldingEntity) validateBankIDCodeEnum(path, location string, value BankIDCode) error {
	if err := validate.Enum(path, location, value, accountHoldingEntityTypeBankIDCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountHoldingEntity) validateBankIDCode(formats strfmt.Registry) error {

	if err := validate.Required("bank_id_code", "body", m.BankIDCode); err != nil {
		return err
	}

	if m.BankIDCode != nil {
		if err := m.BankIDCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bank_id_code")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountHoldingEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountHoldingEntity) UnmarshalBinary(b []byte) error {
	var res AccountHoldingEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountHoldingEntity) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
