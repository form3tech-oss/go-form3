// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ContactAccountAttributes contact account attributes
// swagger:model ContactAccountAttributes
type ContactAccountAttributes struct {

	// account name
	AccountName string `json:"account_name,omitempty"`

	// account number
	AccountNumber string `json:"account_number,omitempty"`

	// account number code
	AccountNumberCode string `json:"account_number_code,omitempty"`

	// account type
	AccountType string `json:"account_type,omitempty"`

	// account with
	AccountWith *ContactAccountAttributesAccountWith `json:"account_with,omitempty"`

	// country
	Country string `json:"country,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`
}

func ContactAccountAttributesWithDefaults(defaults client.Defaults) *ContactAccountAttributes {
	return &ContactAccountAttributes{

		AccountName: defaults.GetString("ContactAccountAttributes", "account_name"),

		AccountNumber: defaults.GetString("ContactAccountAttributes", "account_number"),

		AccountNumberCode: defaults.GetString("ContactAccountAttributes", "account_number_code"),

		AccountType: defaults.GetString("ContactAccountAttributes", "account_type"),

		AccountWith: ContactAccountAttributesAccountWithWithDefaults(defaults),

		Country: defaults.GetString("ContactAccountAttributes", "country"),

		Currency: defaults.GetString("ContactAccountAttributes", "currency"),
	}
}

func (m *ContactAccountAttributes) WithAccountName(accountName string) *ContactAccountAttributes {

	m.AccountName = accountName

	return m
}

func (m *ContactAccountAttributes) WithAccountNumber(accountNumber string) *ContactAccountAttributes {

	m.AccountNumber = accountNumber

	return m
}

func (m *ContactAccountAttributes) WithAccountNumberCode(accountNumberCode string) *ContactAccountAttributes {

	m.AccountNumberCode = accountNumberCode

	return m
}

func (m *ContactAccountAttributes) WithAccountType(accountType string) *ContactAccountAttributes {

	m.AccountType = accountType

	return m
}

func (m *ContactAccountAttributes) WithAccountWith(accountWith ContactAccountAttributesAccountWith) *ContactAccountAttributes {

	m.AccountWith = &accountWith

	return m
}

func (m *ContactAccountAttributes) WithoutAccountWith() *ContactAccountAttributes {
	m.AccountWith = nil
	return m
}

func (m *ContactAccountAttributes) WithCountry(country string) *ContactAccountAttributes {

	m.Country = country

	return m
}

func (m *ContactAccountAttributes) WithCurrency(currency string) *ContactAccountAttributes {

	m.Currency = currency

	return m
}

// Validate validates this contact account attributes
func (m *ContactAccountAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountWith(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactAccountAttributes) validateAccountWith(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ContactAccountAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactAccountAttributes) UnmarshalBinary(b []byte) error {
	var res ContactAccountAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ContactAccountAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
