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

// AccountAmendment account amendment
// swagger:model AccountAmendment
type AccountAmendment struct {

	// data
	Data *AccountUpdate `json:"data,omitempty"`
}

func AccountAmendmentWithDefaults(defaults client.Defaults) *AccountAmendment {
	return &AccountAmendment{

		Data: AccountUpdateWithDefaults(defaults),
	}
}

func (m *AccountAmendment) WithData(data AccountUpdate) *AccountAmendment {

	m.Data = &data

	return m
}

func (m *AccountAmendment) WithoutData() *AccountAmendment {
	m.Data = nil
	return m
}

// Validate validates this account amendment
func (m *AccountAmendment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountAmendment) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAmendment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAmendment) UnmarshalBinary(b []byte) error {
	var res AccountAmendment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountAmendment) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
