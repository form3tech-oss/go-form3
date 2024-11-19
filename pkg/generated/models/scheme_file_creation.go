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

// SchemeFileCreation scheme file creation
// swagger:model SchemeFileCreation
type SchemeFileCreation struct {

	// data
	// Required: true
	Data *NewSchemeFile `json:"data"`
}

func SchemeFileCreationWithDefaults(defaults client.Defaults) *SchemeFileCreation {
	return &SchemeFileCreation{

		Data: NewSchemeFileWithDefaults(defaults),
	}
}

func (m *SchemeFileCreation) WithData(data NewSchemeFile) *SchemeFileCreation {

	m.Data = &data

	return m
}

func (m *SchemeFileCreation) WithoutData() *SchemeFileCreation {
	m.Data = nil
	return m
}

// Validate validates this scheme file creation
func (m *SchemeFileCreation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemeFileCreation) validateData(formats strfmt.Registry) error {

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
func (m *SchemeFileCreation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeFileCreation) UnmarshalBinary(b []byte) error {
	var res SchemeFileCreation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeFileCreation) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
