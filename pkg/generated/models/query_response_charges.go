// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QueryResponseCharges query response charges
// swagger:model QueryResponseCharges
type QueryResponseCharges struct {

	// account number
	AccountNumber string `json:"account_number,omitempty"`

	// account number code
	AccountNumberCode string `json:"account_number_code,omitempty"`

	// amount
	Amount string `json:"amount,omitempty"`

	// currency
	Currency string `json:"currency,omitempty"`
}

func QueryResponseChargesWithDefaults(defaults client.Defaults) *QueryResponseCharges {
	return &QueryResponseCharges{

		AccountNumber: defaults.GetString("QueryResponseCharges", "account_number"),

		AccountNumberCode: defaults.GetString("QueryResponseCharges", "account_number_code"),

		Amount: defaults.GetString("QueryResponseCharges", "amount"),

		Currency: defaults.GetString("QueryResponseCharges", "currency"),
	}
}

func (m *QueryResponseCharges) WithAccountNumber(accountNumber string) *QueryResponseCharges {

	m.AccountNumber = accountNumber

	return m
}

func (m *QueryResponseCharges) WithAccountNumberCode(accountNumberCode string) *QueryResponseCharges {

	m.AccountNumberCode = accountNumberCode

	return m
}

func (m *QueryResponseCharges) WithAmount(amount string) *QueryResponseCharges {

	m.Amount = amount

	return m
}

func (m *QueryResponseCharges) WithCurrency(currency string) *QueryResponseCharges {

	m.Currency = currency

	return m
}

// Validate validates this query response charges
func (m *QueryResponseCharges) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QueryResponseCharges) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResponseCharges) UnmarshalBinary(b []byte) error {
	var res QueryResponseCharges
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryResponseCharges) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
