// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v5/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AccountAmendmentRelationships account amendment relationships
// swagger:model AccountAmendmentRelationships
type AccountAmendmentRelationships struct {

	// account
	Account *AccountAmendmentRelationshipsAccount `json:"account,omitempty"`

	// account amendment submission
	AccountAmendmentSubmission *AccountAmendmentRelationshipsAccountAmendmentSubmission `json:"account_amendment_submission,omitempty"`
}

func AccountAmendmentRelationshipsWithDefaults(defaults client.Defaults) *AccountAmendmentRelationships {
	return &AccountAmendmentRelationships{

		Account: AccountAmendmentRelationshipsAccountWithDefaults(defaults),

		AccountAmendmentSubmission: AccountAmendmentRelationshipsAccountAmendmentSubmissionWithDefaults(defaults),
	}
}

func (m *AccountAmendmentRelationships) WithAccount(account AccountAmendmentRelationshipsAccount) *AccountAmendmentRelationships {

	m.Account = &account

	return m
}

func (m *AccountAmendmentRelationships) WithoutAccount() *AccountAmendmentRelationships {
	m.Account = nil
	return m
}

func (m *AccountAmendmentRelationships) WithAccountAmendmentSubmission(accountAmendmentSubmission AccountAmendmentRelationshipsAccountAmendmentSubmission) *AccountAmendmentRelationships {

	m.AccountAmendmentSubmission = &accountAmendmentSubmission

	return m
}

func (m *AccountAmendmentRelationships) WithoutAccountAmendmentSubmission() *AccountAmendmentRelationships {
	m.AccountAmendmentSubmission = nil
	return m
}

// Validate validates this account amendment relationships
func (m *AccountAmendmentRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountAmendmentSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountAmendmentRelationships) validateAccount(formats strfmt.Registry) error {

	if swag.IsZero(m.Account) { // not required
		return nil
	}

	if m.Account != nil {
		if err := m.Account.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account")
			}
			return err
		}
	}

	return nil
}

func (m *AccountAmendmentRelationships) validateAccountAmendmentSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountAmendmentSubmission) { // not required
		return nil
	}

	if m.AccountAmendmentSubmission != nil {
		if err := m.AccountAmendmentSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("account_amendment_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAmendmentRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAmendmentRelationships) UnmarshalBinary(b []byte) error {
	var res AccountAmendmentRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountAmendmentRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// AccountAmendmentRelationshipsAccount account amendment relationships account
// swagger:model AccountAmendmentRelationshipsAccount
type AccountAmendmentRelationshipsAccount struct {

	// data
	Data []*AmendmentAccountReference `json:"data"`
}

func AccountAmendmentRelationshipsAccountWithDefaults(defaults client.Defaults) *AccountAmendmentRelationshipsAccount {
	return &AccountAmendmentRelationshipsAccount{

		Data: make([]*AmendmentAccountReference, 0),
	}
}

func (m *AccountAmendmentRelationshipsAccount) WithData(data []*AmendmentAccountReference) *AccountAmendmentRelationshipsAccount {

	m.Data = data

	return m
}

// Validate validates this account amendment relationships account
func (m *AccountAmendmentRelationshipsAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountAmendmentRelationshipsAccount) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("account" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAmendmentRelationshipsAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAmendmentRelationshipsAccount) UnmarshalBinary(b []byte) error {
	var res AccountAmendmentRelationshipsAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountAmendmentRelationshipsAccount) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// AccountAmendmentRelationshipsAccountAmendmentSubmission account amendment relationships account amendment submission
// swagger:model AccountAmendmentRelationshipsAccountAmendmentSubmission
type AccountAmendmentRelationshipsAccountAmendmentSubmission struct {

	// Array containing the amendment submission
	// Read Only: true
	Data []*AccountAmendmentSubmission `json:"data,omitempty"`
}

func AccountAmendmentRelationshipsAccountAmendmentSubmissionWithDefaults(defaults client.Defaults) *AccountAmendmentRelationshipsAccountAmendmentSubmission {
	return &AccountAmendmentRelationshipsAccountAmendmentSubmission{

		Data: make([]*AccountAmendmentSubmission, 0),
	}
}

func (m *AccountAmendmentRelationshipsAccountAmendmentSubmission) WithData(data []*AccountAmendmentSubmission) *AccountAmendmentRelationshipsAccountAmendmentSubmission {

	m.Data = data

	return m
}

// Validate validates this account amendment relationships account amendment submission
func (m *AccountAmendmentRelationshipsAccountAmendmentSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountAmendmentRelationshipsAccountAmendmentSubmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("account_amendment_submission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAmendmentRelationshipsAccountAmendmentSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAmendmentRelationshipsAccountAmendmentSubmission) UnmarshalBinary(b []byte) error {
	var res AccountAmendmentRelationshipsAccountAmendmentSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountAmendmentRelationshipsAccountAmendmentSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
