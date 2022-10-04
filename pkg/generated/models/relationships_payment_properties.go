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

// RelationshipsPaymentProperties relationships payment properties
// swagger:model RelationshipsPaymentProperties
type RelationshipsPaymentProperties struct {

	// The payment associated with this query
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// type
	// Required: true
	Type *PaymentResourceType `json:"type"`
}

func RelationshipsPaymentPropertiesWithDefaults(defaults client.Defaults) *RelationshipsPaymentProperties {
	return &RelationshipsPaymentProperties{

		ID: defaults.GetStrfmtUUIDPtr("RelationshipsPaymentProperties", "id"),

		// TODO Type: PaymentResourceType,

	}
}

func (m *RelationshipsPaymentProperties) WithID(id strfmt.UUID) *RelationshipsPaymentProperties {

	m.ID = &id

	return m
}

func (m *RelationshipsPaymentProperties) WithoutID() *RelationshipsPaymentProperties {
	m.ID = nil
	return m
}

func (m *RelationshipsPaymentProperties) WithType(typeVar PaymentResourceType) *RelationshipsPaymentProperties {

	m.Type = &typeVar

	return m
}

func (m *RelationshipsPaymentProperties) WithoutType() *RelationshipsPaymentProperties {
	m.Type = nil
	return m
}

// Validate validates this relationships payment properties
func (m *RelationshipsPaymentProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelationshipsPaymentProperties) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RelationshipsPaymentProperties) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelationshipsPaymentProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationshipsPaymentProperties) UnmarshalBinary(b []byte) error {
	var res RelationshipsPaymentProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RelationshipsPaymentProperties) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
