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

// SchemeMessageAdmissionRelationships scheme message admission relationships
// swagger:model SchemeMessageAdmissionRelationships
type SchemeMessageAdmissionRelationships struct {

	// scheme message
	SchemeMessage *SchemeMessageAdmissionRelationshipsSchemeMessage `json:"scheme_message,omitempty"`
}

func SchemeMessageAdmissionRelationshipsWithDefaults(defaults client.Defaults) *SchemeMessageAdmissionRelationships {
	return &SchemeMessageAdmissionRelationships{

		SchemeMessage: SchemeMessageAdmissionRelationshipsSchemeMessageWithDefaults(defaults),
	}
}

func (m *SchemeMessageAdmissionRelationships) WithSchemeMessage(schemeMessage SchemeMessageAdmissionRelationshipsSchemeMessage) *SchemeMessageAdmissionRelationships {

	m.SchemeMessage = &schemeMessage

	return m
}

func (m *SchemeMessageAdmissionRelationships) WithoutSchemeMessage() *SchemeMessageAdmissionRelationships {
	m.SchemeMessage = nil
	return m
}

// Validate validates this scheme message admission relationships
func (m *SchemeMessageAdmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchemeMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemeMessageAdmissionRelationships) validateSchemeMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.SchemeMessage) { // not required
		return nil
	}

	if m.SchemeMessage != nil {
		if err := m.SchemeMessage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheme_message")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemeMessageAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeMessageAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res SchemeMessageAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeMessageAdmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// SchemeMessageAdmissionRelationshipsSchemeMessage scheme message admission relationships scheme message
// swagger:model SchemeMessageAdmissionRelationshipsSchemeMessage
type SchemeMessageAdmissionRelationshipsSchemeMessage struct {

	// data
	Data []*SchemeMessage `json:"data"`
}

func SchemeMessageAdmissionRelationshipsSchemeMessageWithDefaults(defaults client.Defaults) *SchemeMessageAdmissionRelationshipsSchemeMessage {
	return &SchemeMessageAdmissionRelationshipsSchemeMessage{

		Data: make([]*SchemeMessage, 0),
	}
}

func (m *SchemeMessageAdmissionRelationshipsSchemeMessage) WithData(data []*SchemeMessage) *SchemeMessageAdmissionRelationshipsSchemeMessage {

	m.Data = data

	return m
}

// Validate validates this scheme message admission relationships scheme message
func (m *SchemeMessageAdmissionRelationshipsSchemeMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemeMessageAdmissionRelationshipsSchemeMessage) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("scheme_message" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemeMessageAdmissionRelationshipsSchemeMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeMessageAdmissionRelationshipsSchemeMessage) UnmarshalBinary(b []byte) error {
	var res SchemeMessageAdmissionRelationshipsSchemeMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeMessageAdmissionRelationshipsSchemeMessage) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
