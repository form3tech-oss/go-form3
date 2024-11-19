// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SchemeMessageRelationships scheme message relationships
// swagger:model SchemeMessageRelationships
type SchemeMessageRelationships struct {

	// scheme message admission
	SchemeMessageAdmission *SchemeMessageRelationshipsSchemeMessageAdmission `json:"scheme_message_admission,omitempty"`
}

func SchemeMessageRelationshipsWithDefaults(defaults client.Defaults) *SchemeMessageRelationships {
	return &SchemeMessageRelationships{

		SchemeMessageAdmission: SchemeMessageRelationshipsSchemeMessageAdmissionWithDefaults(defaults),
	}
}

func (m *SchemeMessageRelationships) WithSchemeMessageAdmission(schemeMessageAdmission SchemeMessageRelationshipsSchemeMessageAdmission) *SchemeMessageRelationships {

	m.SchemeMessageAdmission = &schemeMessageAdmission

	return m
}

func (m *SchemeMessageRelationships) WithoutSchemeMessageAdmission() *SchemeMessageRelationships {
	m.SchemeMessageAdmission = nil
	return m
}

// Validate validates this scheme message relationships
func (m *SchemeMessageRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchemeMessageAdmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemeMessageRelationships) validateSchemeMessageAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.SchemeMessageAdmission) { // not required
		return nil
	}

	if m.SchemeMessageAdmission != nil {
		if err := m.SchemeMessageAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheme_message_admission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemeMessageRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeMessageRelationships) UnmarshalBinary(b []byte) error {
	var res SchemeMessageRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeMessageRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// SchemeMessageRelationshipsSchemeMessageAdmission scheme message relationships scheme message admission
// swagger:model SchemeMessageRelationshipsSchemeMessageAdmission
type SchemeMessageRelationshipsSchemeMessageAdmission struct {

	// data
	Data []*SchemeMessageAdmission `json:"data"`
}

func SchemeMessageRelationshipsSchemeMessageAdmissionWithDefaults(defaults client.Defaults) *SchemeMessageRelationshipsSchemeMessageAdmission {
	return &SchemeMessageRelationshipsSchemeMessageAdmission{

		Data: make([]*SchemeMessageAdmission, 0),
	}
}

func (m *SchemeMessageRelationshipsSchemeMessageAdmission) WithData(data []*SchemeMessageAdmission) *SchemeMessageRelationshipsSchemeMessageAdmission {

	m.Data = data

	return m
}

// Validate validates this scheme message relationships scheme message admission
func (m *SchemeMessageRelationshipsSchemeMessageAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemeMessageRelationshipsSchemeMessageAdmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("scheme_message_admission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemeMessageRelationshipsSchemeMessageAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeMessageRelationshipsSchemeMessageAdmission) UnmarshalBinary(b []byte) error {
	var res SchemeMessageRelationshipsSchemeMessageAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeMessageRelationshipsSchemeMessageAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
