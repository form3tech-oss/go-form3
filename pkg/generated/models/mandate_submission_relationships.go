// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MandateSubmissionRelationships mandate submission relationships
// swagger:model MandateSubmissionRelationships
type MandateSubmissionRelationships struct {

	// mandate
	Mandate *MandateSubmissionRelationshipsMandate `json:"mandate,omitempty"`
}

func MandateSubmissionRelationshipsWithDefaults(defaults client.Defaults) *MandateSubmissionRelationships {
	return &MandateSubmissionRelationships{

		Mandate: MandateSubmissionRelationshipsMandateWithDefaults(defaults),
	}
}

func (m *MandateSubmissionRelationships) WithMandate(mandate MandateSubmissionRelationshipsMandate) *MandateSubmissionRelationships {

	m.Mandate = &mandate

	return m
}

func (m *MandateSubmissionRelationships) WithoutMandate() *MandateSubmissionRelationships {
	m.Mandate = nil
	return m
}

// Validate validates this mandate submission relationships
func (m *MandateSubmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMandate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateSubmissionRelationships) validateMandate(formats strfmt.Registry) error {

	if swag.IsZero(m.Mandate) { // not required
		return nil
	}

	if m.Mandate != nil {
		if err := m.Mandate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mandate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateSubmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateSubmissionRelationships) UnmarshalBinary(b []byte) error {
	var res MandateSubmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateSubmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// MandateSubmissionRelationshipsMandate mandate submission relationships mandate
// swagger:model MandateSubmissionRelationshipsMandate
type MandateSubmissionRelationshipsMandate struct {

	// data
	Data []*Mandate `json:"data"`
}

func MandateSubmissionRelationshipsMandateWithDefaults(defaults client.Defaults) *MandateSubmissionRelationshipsMandate {
	return &MandateSubmissionRelationshipsMandate{

		Data: make([]*Mandate, 0),
	}
}

func (m *MandateSubmissionRelationshipsMandate) WithData(data []*Mandate) *MandateSubmissionRelationshipsMandate {

	m.Data = data

	return m
}

// Validate validates this mandate submission relationships mandate
func (m *MandateSubmissionRelationshipsMandate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateSubmissionRelationshipsMandate) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("mandate" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MandateSubmissionRelationshipsMandate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MandateSubmissionRelationshipsMandate) UnmarshalBinary(b []byte) error {
	var res MandateSubmissionRelationshipsMandate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *MandateSubmissionRelationshipsMandate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
