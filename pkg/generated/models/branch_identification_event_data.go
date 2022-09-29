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
)

// BranchIdentificationEventData branch identification event data
// swagger:model BranchIdentificationEventData
type BranchIdentificationEventData struct {
	BranchIdentification

	// relationships
	Relationships *BranchIdentificationEventRelationships `json:"relationships,omitempty"`
}

func BranchIdentificationEventDataWithDefaults(defaults client.Defaults) *BranchIdentificationEventData {
	return &BranchIdentificationEventData{}
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BranchIdentificationEventData) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BranchIdentification
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BranchIdentification = aO0

	// AO1
	var dataAO1 struct {
		Relationships *BranchIdentificationEventRelationships `json:"relationships,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Relationships = dataAO1.Relationships

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BranchIdentificationEventData) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BranchIdentification)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Relationships *BranchIdentificationEventRelationships `json:"relationships,omitempty"`
	}

	dataAO1.Relationships = m.Relationships

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this branch identification event data
func (m *BranchIdentificationEventData) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BranchIdentification
	if err := m.BranchIdentification.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BranchIdentificationEventData) validateRelationships(formats strfmt.Registry) error {

	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	if m.Relationships != nil {
		if err := m.Relationships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BranchIdentificationEventData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BranchIdentificationEventData) UnmarshalBinary(b []byte) error {
	var res BranchIdentificationEventData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *BranchIdentificationEventData) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
