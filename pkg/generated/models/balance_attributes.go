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

// BalanceAttributes balance attributes
// swagger:model BalanceAttributes
type BalanceAttributes struct {

	// account number
	AccountNumber string `json:"account_number,omitempty"`

	// account number code
	// Enum: [IBAN]
	AccountNumberCode string `json:"account_number_code,omitempty"`

	// Amount of funds
	Amount string `json:"amount,omitempty"`

	// Currency of funds
	// Pattern: ^[A-Z]{3}$
	Currency string `json:"currency,omitempty"`

	// Balance description
	Description string `json:"description,omitempty"`

	// Institution that holds the funds
	HoldingInstitution string `json:"holding_institution,omitempty"`
}

func BalanceAttributesWithDefaults(defaults client.Defaults) *BalanceAttributes {
	return &BalanceAttributes{

		AccountNumber: defaults.GetString("BalanceAttributes", "account_number"),

		AccountNumberCode: defaults.GetString("BalanceAttributes", "account_number_code"),

		Amount: defaults.GetString("BalanceAttributes", "amount"),

		Currency: defaults.GetString("BalanceAttributes", "currency"),

		Description: defaults.GetString("BalanceAttributes", "description"),

		HoldingInstitution: defaults.GetString("BalanceAttributes", "holding_institution"),
	}
}

func (m *BalanceAttributes) WithAccountNumber(accountNumber string) *BalanceAttributes {

	m.AccountNumber = accountNumber

	return m
}

func (m *BalanceAttributes) WithAccountNumberCode(accountNumberCode string) *BalanceAttributes {

	m.AccountNumberCode = accountNumberCode

	return m
}

func (m *BalanceAttributes) WithAmount(amount string) *BalanceAttributes {

	m.Amount = amount

	return m
}

func (m *BalanceAttributes) WithCurrency(currency string) *BalanceAttributes {

	m.Currency = currency

	return m
}

func (m *BalanceAttributes) WithDescription(description string) *BalanceAttributes {

	m.Description = description

	return m
}

func (m *BalanceAttributes) WithHoldingInstitution(holdingInstitution string) *BalanceAttributes {

	m.HoldingInstitution = holdingInstitution

	return m
}

// Validate validates this balance attributes
func (m *BalanceAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumberCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var balanceAttributesTypeAccountNumberCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IBAN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		balanceAttributesTypeAccountNumberCodePropEnum = append(balanceAttributesTypeAccountNumberCodePropEnum, v)
	}
}

const (

	// BalanceAttributesAccountNumberCodeIBAN captures enum value "IBAN"
	BalanceAttributesAccountNumberCodeIBAN string = "IBAN"
)

// prop value enum
func (m *BalanceAttributes) validateAccountNumberCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, balanceAttributesTypeAccountNumberCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BalanceAttributes) validateAccountNumberCode(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountNumberCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccountNumberCodeEnum("account_number_code", "body", m.AccountNumberCode); err != nil {
		return err
	}

	return nil
}

func (m *BalanceAttributes) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	if err := validate.Pattern("currency", "body", string(m.Currency), `^[A-Z]{3}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BalanceAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BalanceAttributes) UnmarshalBinary(b []byte) error {
	var res BalanceAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *BalanceAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
