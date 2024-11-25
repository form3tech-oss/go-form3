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

// SwitchedAccountDetails Alternate Account details to use in case the account has been switched away from this organisation.
// swagger:model SwitchedAccountDetails
type SwitchedAccountDetails struct {

	// Switched account number. Must be a UK account number, maximum length 8 characters.
	// Required: true
	// Pattern: ^[0-9]{8}$
	AccountNumber *string `json:"account_number"`

	// ISO 20022 code used to identify the type of account number being used
	// Required: true
	// Enum: ["BBAN"]
	AccountNumberCode *string `json:"account_number_code"`

	// The type of the account provided in account_number. Only required if requested by the beneficiary party.
	// Maximum: 9
	// Minimum: 0
	AccountType *int64 `json:"account_type,omitempty"`

	// account with
	// Required: true
	AccountWith *AccountHoldingEntity `json:"account_with"`

	// Starting date for the account to be effectively switched
	// Required: true
	// Format: date
	SwitchedEffectiveDate *strfmt.Date `json:"switched_effective_date"`
}

func SwitchedAccountDetailsWithDefaults(defaults client.Defaults) *SwitchedAccountDetails {
	return &SwitchedAccountDetails{

		AccountNumber: defaults.GetStringPtr("SwitchedAccountDetails", "account_number"),

		AccountNumberCode: defaults.GetStringPtr("SwitchedAccountDetails", "account_number_code"),

		AccountType: defaults.GetInt64Ptr("SwitchedAccountDetails", "account_type"),

		AccountWith: AccountHoldingEntityWithDefaults(defaults),

		SwitchedEffectiveDate: defaults.GetStrfmtDatePtr("SwitchedAccountDetails", "switched_effective_date"),
	}
}

func (m *SwitchedAccountDetails) WithAccountNumber(accountNumber string) *SwitchedAccountDetails {

	m.AccountNumber = &accountNumber

	return m
}

func (m *SwitchedAccountDetails) WithoutAccountNumber() *SwitchedAccountDetails {
	m.AccountNumber = nil
	return m
}

func (m *SwitchedAccountDetails) WithAccountNumberCode(accountNumberCode string) *SwitchedAccountDetails {

	m.AccountNumberCode = &accountNumberCode

	return m
}

func (m *SwitchedAccountDetails) WithoutAccountNumberCode() *SwitchedAccountDetails {
	m.AccountNumberCode = nil
	return m
}

func (m *SwitchedAccountDetails) WithAccountType(accountType int64) *SwitchedAccountDetails {

	m.AccountType = &accountType

	return m
}

func (m *SwitchedAccountDetails) WithoutAccountType() *SwitchedAccountDetails {
	m.AccountType = nil
	return m
}

func (m *SwitchedAccountDetails) WithAccountWith(accountWith AccountHoldingEntity) *SwitchedAccountDetails {

	m.AccountWith = &accountWith

	return m
}

func (m *SwitchedAccountDetails) WithoutAccountWith() *SwitchedAccountDetails {
	m.AccountWith = nil
	return m
}

func (m *SwitchedAccountDetails) WithSwitchedEffectiveDate(switchedEffectiveDate strfmt.Date) *SwitchedAccountDetails {

	m.SwitchedEffectiveDate = &switchedEffectiveDate

	return m
}

func (m *SwitchedAccountDetails) WithoutSwitchedEffectiveDate() *SwitchedAccountDetails {
	m.SwitchedEffectiveDate = nil
	return m
}

// Validate validates this switched account details
func (m *SwitchedAccountDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountNumberCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountWith(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwitchedEffectiveDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SwitchedAccountDetails) validateAccountNumber(formats strfmt.Registry) error {

	if err := validate.Required("account_number", "body", m.AccountNumber); err != nil {
		return err
	}

	if err := validate.Pattern("account_number", "body", string(*m.AccountNumber), `^[0-9]{8}$`); err != nil {
		return err
	}

	return nil
}

var switchedAccountDetailsTypeAccountNumberCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BBAN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		switchedAccountDetailsTypeAccountNumberCodePropEnum = append(switchedAccountDetailsTypeAccountNumberCodePropEnum, v)
	}
}

const (

	// SwitchedAccountDetailsAccountNumberCodeBBAN captures enum value "BBAN"
	SwitchedAccountDetailsAccountNumberCodeBBAN string = "BBAN"
)

// prop value enum
func (m *SwitchedAccountDetails) validateAccountNumberCodeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, switchedAccountDetailsTypeAccountNumberCodePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SwitchedAccountDetails) validateAccountNumberCode(formats strfmt.Registry) error {

	if err := validate.Required("account_number_code", "body", m.AccountNumberCode); err != nil {
		return err
	}

	// value enum
	if err := m.validateAccountNumberCodeEnum("account_number_code", "body", *m.AccountNumberCode); err != nil {
		return err
	}

	return nil
}

func (m *SwitchedAccountDetails) validateAccountType(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountType) { // not required
		return nil
	}

	if err := validate.MinimumInt("account_type", "body", int64(*m.AccountType), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("account_type", "body", int64(*m.AccountType), 9, false); err != nil {
		return err
	}

	return nil
}

func (m *SwitchedAccountDetails) validateAccountWith(formats strfmt.Registry) error {

	if err := validate.Required("account_with", "body", m.AccountWith); err != nil {
		return err
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

func (m *SwitchedAccountDetails) validateSwitchedEffectiveDate(formats strfmt.Registry) error {

	if err := validate.Required("switched_effective_date", "body", m.SwitchedEffectiveDate); err != nil {
		return err
	}

	if err := validate.FormatOf("switched_effective_date", "body", "date", m.SwitchedEffectiveDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SwitchedAccountDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwitchedAccountDetails) UnmarshalBinary(b []byte) error {
	var res SwitchedAccountDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SwitchedAccountDetails) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
