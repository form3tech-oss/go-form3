// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RelationshipsPaymentAdmissionProperties relationships payment admission properties
// swagger:model RelationshipsPaymentAdmissionProperties
type RelationshipsPaymentAdmissionProperties struct {

	// The payment admission associated with this query
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// type
	// Required: true
	Type *PaymentAdmissionResourceType `json:"type"`
}

func RelationshipsPaymentAdmissionPropertiesWithDefaults(defaults client.Defaults) *RelationshipsPaymentAdmissionProperties {
	return &RelationshipsPaymentAdmissionProperties{

		ID: defaults.GetStrfmtUUIDPtr("RelationshipsPaymentAdmissionProperties", "id"),

		// TODO Type: PaymentAdmissionResourceType,

	}
}

func (m *RelationshipsPaymentAdmissionProperties) WithID(id strfmt.UUID) *RelationshipsPaymentAdmissionProperties {

	m.ID = &id

	return m
}

func (m *RelationshipsPaymentAdmissionProperties) WithoutID() *RelationshipsPaymentAdmissionProperties {
	m.ID = nil
	return m
}

func (m *RelationshipsPaymentAdmissionProperties) WithType(typeVar PaymentAdmissionResourceType) *RelationshipsPaymentAdmissionProperties {

	m.Type = &typeVar

	return m
}

func (m *RelationshipsPaymentAdmissionProperties) WithoutType() *RelationshipsPaymentAdmissionProperties {
	m.Type = nil
	return m
}

// Validate validates this relationships payment admission properties
func (m *RelationshipsPaymentAdmissionProperties) Validate(formats strfmt.Registry) error {
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

func (m *RelationshipsPaymentAdmissionProperties) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RelationshipsPaymentAdmissionProperties) validateType(formats strfmt.Registry) error {

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
func (m *RelationshipsPaymentAdmissionProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationshipsPaymentAdmissionProperties) UnmarshalBinary(b []byte) error {
	var res RelationshipsPaymentAdmissionProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RelationshipsPaymentAdmissionProperties) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
