// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MandateAttributesBeneficiaryParty mandate attributes beneficiary party
// swagger:model MandateAttributesBeneficiaryParty
type MandateAttributesBeneficiaryParty struct {

	// account name
	AccountName string `json:"account_name"`

	// account number
	// Pattern: ^[A-Z0-9]{6,34}$
	AccountNumber string `json:"account_number,omitempty"`

	// account number code
	AccountNumberCode AccountNumberCode `json:"account_number_code,omitempty"`

	// account type
	AccountType int64 `json:"account_type,omitempty"`

	// account with
	AccountWith *AccountHoldingEntity `json:"account_with,omitempty"`

	// address
	Address []string `json:"address"`

	// country
	Country string `json:"country,omitempty"`

	// private identification
	PrivateIdentification *PrivateIdentification `json:"private_identification,omitempty"`
}

func MandateAttributesBeneficiaryPartyWithDefaults(defaults client.Defaults) *MandateAttributesBeneficiaryParty {
	return &MandateAttributesBeneficiaryParty{

		AccountName: defaults.GetString("MandateAttributesBeneficiaryParty", "account_name"),

		AccountNumber: defaults.GetString("MandateAttributesBeneficiaryParty", "account_number"),

		// TODO AccountNumberCode: AccountNumberCode,

		AccountType: defaults.GetInt64("MandateAttributesBeneficiaryParty", "account_type"),

		AccountWith: AccountHoldingEntityWithDefaults(defaults),

		Address: make([]string, 0),

		Country: defaults.GetString("MandateAttributesBeneficiaryParty", "country"),

		PrivateIdentification: PrivateIdentificationWithDefaults(defaults),
	}
}

func (m *MandateAttributesBeneficiaryParty) WithAccountName(accountName string) *MandateAttributesBeneficiaryParty {

	m.AccountName = accountName

	return m
}

func (m *MandateAttributesBeneficiaryParty) WithAccountNumber(accountNumber string) *MandateAttributesBeneficiaryParty {

	m.AccountNumber = accountNumber

	return m
}

func (m *MandateAttributesBeneficiaryParty) WithAccountNumberCode(accountNumberCode AccountNumberCode) *MandateAttributesBeneficiaryParty {

	m.AccountNumberCode = accountNumberCode

	return m
}

func (m *MandateAttributesBeneficiaryParty) WithAccountType(accountType int64) *MandateAttributesBeneficiaryParty {

	m.AccountType = accountType

	return m
}

func (m *MandateAttributesBeneficiaryParty) WithAccountWith(accountWith AccountHoldingEntity) *MandateAttributesBeneficiaryParty {

	m.AccountWith = &accountWith

	return m
}

func (m *MandateAttributesBeneficiaryParty) WithoutAccountWith() *MandateAttributesBeneficiaryParty {
	m.AccountWith = nil
	return m
}

func (m *MandateAttributesBeneficiaryParty) WithAddress(address []string) *MandateAttributesBeneficiaryParty {

	m.Address = address

	return m
}

func (m *MandateAttributesBeneficiaryParty) WithCountry(country string) *MandateAttributesBeneficiaryParty {

	m.Country = country

	return m
}

func (m *MandateAttributesBeneficiaryParty) WithPrivateIdentification(privateIdentification PrivateIdentification) *MandateAttributesBeneficiaryParty {

	m.PrivateIdentification = &privateIdentification

	return m
}

func (m *MandateAttributesBeneficiaryParty) WithoutPrivateIdentification() *MandateAttributesBeneficiaryParty {
	m.PrivateIdentification = nil
	return m
}

// Validate validates this mandate attributes beneficiary party
func (m *MandateAttributesBeneficiaryParty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountNumberCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountWith(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateIdentification(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateAttributesBeneficiaryParty) validateAccountNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumber) { // not required
		return nil
	}

	if err := validate.Pattern("account_number", "body", string(m.AccountNumber), `^[A-Z0-9]{6,34}$`); err != nil {
		return err
	}

	return nil
}

func (m *MandateAttributesBeneficiaryParty) validateAccountNumberCode(formats strfmt.Registry) error {

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

func (m *MandateAttributesBeneficiaryParty) validateAccountWith(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountWith) { // not required
		return nil
	}

	if m.AccountWith != nil {
		if err := m.AccountWith.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account_with")
			}
			return err
		}
	}

	return nil
}

func (m *MandateAttributesBeneficiaryParty) validatePrivateIdentification(formats strfmt.Registry) error {

	if swag.IsZero(m.PrivateIdentification) { // not required
		return nil
	}

	if m.PrivateIdentification != nil {
		if err := m.PrivateIdentification.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("private_identification")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateAttributesBeneficiaryParty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateAttributesBeneficiaryParty) UnmarshalBinary(b []byte) error {
	var res MandateAttributesBeneficiaryParty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateAttributesBeneficiaryParty) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
