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

// AccountGenerationAttributes account generation attributes
// swagger:model AccountGenerationAttributes
type AccountGenerationAttributes struct {

	// Account number of the account. A unique number will automatically be generated if not provided.
	// Pattern: ^[A-Z0-9]{0,64}$
	AccountNumber string `json:"account_number,omitempty"`

	// Local country bank identifier. In the UK this is the sort code.
	// Pattern: ^[A-Z0-9]{0,16}$
	BankID string `json:"bank_id,omitempty"`

	// ISO 4217 code used to identify the base currency of the account
	// Pattern: ^[A-Z]{3}$
	BaseCurrency string `json:"base_currency,omitempty"`

	// SWIFT BIC in either 8 or 11 character format
	// Pattern: ^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$
	Bic string `json:"bic,omitempty"`

	// ISO 3166-1 code used to identify the domicile of the account
	// Required: true
	// Pattern: ^[A-Z]{2}$
	Country *string `json:"country"`
}

func AccountGenerationAttributesWithDefaults(defaults client.Defaults) *AccountGenerationAttributes {
	return &AccountGenerationAttributes{

		AccountNumber: defaults.GetString("AccountGenerationAttributes", "account_number"),

		BankID: defaults.GetString("AccountGenerationAttributes", "bank_id"),

		BaseCurrency: defaults.GetString("AccountGenerationAttributes", "base_currency"),

		Bic: defaults.GetString("AccountGenerationAttributes", "bic"),

		Country: defaults.GetStringPtr("AccountGenerationAttributes", "country"),
	}
}

func (m *AccountGenerationAttributes) WithAccountNumber(accountNumber string) *AccountGenerationAttributes {

	m.AccountNumber = accountNumber

	return m
}

func (m *AccountGenerationAttributes) WithBankID(bankID string) *AccountGenerationAttributes {

	m.BankID = bankID

	return m
}

func (m *AccountGenerationAttributes) WithBaseCurrency(baseCurrency string) *AccountGenerationAttributes {

	m.BaseCurrency = baseCurrency

	return m
}

func (m *AccountGenerationAttributes) WithBic(bic string) *AccountGenerationAttributes {

	m.Bic = bic

	return m
}

func (m *AccountGenerationAttributes) WithCountry(country string) *AccountGenerationAttributes {

	m.Country = &country

	return m
}

func (m *AccountGenerationAttributes) WithoutCountry() *AccountGenerationAttributes {
	m.Country = nil
	return m
}

// Validate validates this account generation attributes
func (m *AccountGenerationAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBaseCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountGenerationAttributes) validateAccountNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumber) { // not required
		return nil
	}

	if err := validate.Pattern("account_number", "body", string(m.AccountNumber), `^[A-Z0-9]{0,64}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountGenerationAttributes) validateBankID(formats strfmt.Registry) error {

	if swag.IsZero(m.BankID) { // not required
		return nil
	}

	if err := validate.Pattern("bank_id", "body", string(m.BankID), `^[A-Z0-9]{0,16}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountGenerationAttributes) validateBaseCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.BaseCurrency) { // not required
		return nil
	}

	if err := validate.Pattern("base_currency", "body", string(m.BaseCurrency), `^[A-Z]{3}$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountGenerationAttributes) validateBic(formats strfmt.Registry) error {

	if swag.IsZero(m.Bic) { // not required
		return nil
	}

	if err := validate.Pattern("bic", "body", string(m.Bic), `^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$`); err != nil {
		return err
	}

	return nil
}

func (m *AccountGenerationAttributes) validateCountry(formats strfmt.Registry) error {

	if err := validate.Required("country", "body", m.Country); err != nil {
		return err
	}

	if err := validate.Pattern("country", "body", string(*m.Country), `^[A-Z]{2}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountGenerationAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountGenerationAttributes) UnmarshalBinary(b []byte) error {
	var res AccountGenerationAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountGenerationAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
