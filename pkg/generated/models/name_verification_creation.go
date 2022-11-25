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

// NameVerificationCreation name verification creation
// swagger:model NameVerificationCreation
type NameVerificationCreation struct {

	// data
	// Required: true
	Data *NameVerification `json:"data"`
}

func NameVerificationCreationWithDefaults(defaults client.Defaults) *NameVerificationCreation {
	return &NameVerificationCreation{

		Data: NameVerificationWithDefaults(defaults),
	}
}

func (m *NameVerificationCreation) WithData(data NameVerification) *NameVerificationCreation {

	m.Data = &data

	return m
}

func (m *NameVerificationCreation) WithoutData() *NameVerificationCreation {
	m.Data = nil
	return m
}

// Validate validates this name verification creation
func (m *NameVerificationCreation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameVerificationCreation) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
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
func (m *NameVerificationCreation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameVerificationCreation) UnmarshalBinary(b []byte) error {
	var res NameVerificationCreation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NameVerificationCreation) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
