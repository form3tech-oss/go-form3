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
)

// DebtorParty debtor party
// swagger:model DebtorParty
type DebtorParty struct {

	// account with
	AccountWith *DebtorPartyAccountWith `json:"account_with,omitempty"`
}

func DebtorPartyWithDefaults(defaults client.Defaults) *DebtorParty {
	return &DebtorParty{

		AccountWith: DebtorPartyAccountWithWithDefaults(defaults),
	}
}

func (m *DebtorParty) WithAccountWith(accountWith DebtorPartyAccountWith) *DebtorParty {

	m.AccountWith = &accountWith

	return m
}

func (m *DebtorParty) WithoutAccountWith() *DebtorParty {
	m.AccountWith = nil
	return m
}

// Validate validates this debtor party
func (m *DebtorParty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountWith(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DebtorParty) validateAccountWith(formats strfmt.Registry) error {

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
func (m *DebtorParty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DebtorParty) UnmarshalBinary(b []byte) error {
	var res DebtorParty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DebtorParty) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DebtorPartyAccountWith debtor party account with
// swagger:model DebtorPartyAccountWith
type DebtorPartyAccountWith struct {

	// Financial institution name
	BankName string `json:"bank_name,omitempty"`
}

func DebtorPartyAccountWithWithDefaults(defaults client.Defaults) *DebtorPartyAccountWith {
	return &DebtorPartyAccountWith{

		BankName: defaults.GetString("DebtorPartyAccountWith", "bank_name"),
	}
}

func (m *DebtorPartyAccountWith) WithBankName(bankName string) *DebtorPartyAccountWith {

	m.BankName = bankName

	return m
}

// Validate validates this debtor party account with
func (m *DebtorPartyAccountWith) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DebtorPartyAccountWith) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DebtorPartyAccountWith) UnmarshalBinary(b []byte) error {
	var res DebtorPartyAccountWith
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DebtorPartyAccountWith) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
