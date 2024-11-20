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
	"github.com/go-openapi/validate"
)

// ThinRelationship thin relationship
// swagger:model ThinRelationship
type ThinRelationship struct {

	// A relationship which just contains id and type of the related resource
	Data []*ThinRelationshipDataItems0 `json:"data"`
}

func ThinRelationshipWithDefaults(defaults client.Defaults) *ThinRelationship {
	return &ThinRelationship{

		Data: make([]*ThinRelationshipDataItems0, 0),
	}
}

func (m *ThinRelationship) WithData(data []*ThinRelationshipDataItems0) *ThinRelationship {

	m.Data = data

	return m
}

// Validate validates this thin relationship
func (m *ThinRelationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinRelationship) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThinRelationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinRelationship) UnmarshalBinary(b []byte) error {
	var res ThinRelationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ThinRelationship) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ThinRelationshipDataItems0 thin relationship data items0
// swagger:model ThinRelationshipDataItems0
type ThinRelationshipDataItems0 struct {

	// ID of the referenced resource
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Name of the referenced resource type
	Type string `json:"type,omitempty"`
}

func ThinRelationshipDataItems0WithDefaults(defaults client.Defaults) *ThinRelationshipDataItems0 {
	return &ThinRelationshipDataItems0{

		ID: defaults.GetStrfmtUUID("ThinRelationshipDataItems0", "id"),

		Type: defaults.GetString("ThinRelationshipDataItems0", "type"),
	}
}

func (m *ThinRelationshipDataItems0) WithID(id strfmt.UUID) *ThinRelationshipDataItems0 {

	m.ID = id

	return m
}

func (m *ThinRelationshipDataItems0) WithType(typeVar string) *ThinRelationshipDataItems0 {

	m.Type = typeVar

	return m
}

// Validate validates this thin relationship data items0
func (m *ThinRelationshipDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThinRelationshipDataItems0) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThinRelationshipDataItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThinRelationshipDataItems0) UnmarshalBinary(b []byte) error {
	var res ThinRelationshipDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ThinRelationshipDataItems0) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
