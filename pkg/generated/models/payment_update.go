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

// PaymentUpdate payment update
// swagger:model PaymentUpdate
type PaymentUpdate struct {

	// attributes
	// Required: true
	Attributes *PaymentAttributes `json:"attributes"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// modified on
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *PaymentUpdateRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Required: true
	// Minimum: 0
	Version *int64 `json:"version"`
}

func PaymentUpdateWithDefaults(defaults client.Defaults) *PaymentUpdate {
	return &PaymentUpdate{

		Attributes: PaymentAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("PaymentUpdate", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("PaymentUpdate", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("PaymentUpdate", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("PaymentUpdate", "organisation_id"),

		Relationships: PaymentUpdateRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("PaymentUpdate", "type"),

		Version: defaults.GetInt64Ptr("PaymentUpdate", "version"),
	}
}

func (m *PaymentUpdate) WithAttributes(attributes PaymentAttributes) *PaymentUpdate {

	m.Attributes = &attributes

	return m
}

func (m *PaymentUpdate) WithoutAttributes() *PaymentUpdate {
	m.Attributes = nil
	return m
}

func (m *PaymentUpdate) WithCreatedOn(createdOn strfmt.DateTime) *PaymentUpdate {

	m.CreatedOn = &createdOn

	return m
}

func (m *PaymentUpdate) WithoutCreatedOn() *PaymentUpdate {
	m.CreatedOn = nil
	return m
}

func (m *PaymentUpdate) WithID(id strfmt.UUID) *PaymentUpdate {

	m.ID = &id

	return m
}

func (m *PaymentUpdate) WithoutID() *PaymentUpdate {
	m.ID = nil
	return m
}

func (m *PaymentUpdate) WithModifiedOn(modifiedOn strfmt.DateTime) *PaymentUpdate {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *PaymentUpdate) WithoutModifiedOn() *PaymentUpdate {
	m.ModifiedOn = nil
	return m
}

func (m *PaymentUpdate) WithOrganisationID(organisationID strfmt.UUID) *PaymentUpdate {

	m.OrganisationID = &organisationID

	return m
}

func (m *PaymentUpdate) WithoutOrganisationID() *PaymentUpdate {
	m.OrganisationID = nil
	return m
}

func (m *PaymentUpdate) WithRelationships(relationships PaymentUpdateRelationships) *PaymentUpdate {

	m.Relationships = &relationships

	return m
}

func (m *PaymentUpdate) WithoutRelationships() *PaymentUpdate {
	m.Relationships = nil
	return m
}

func (m *PaymentUpdate) WithType(typeVar string) *PaymentUpdate {

	m.Type = typeVar

	return m
}

func (m *PaymentUpdate) WithVersion(version int64) *PaymentUpdate {

	m.Version = &version

	return m
}

func (m *PaymentUpdate) WithoutVersion() *PaymentUpdate {
	m.Version = nil
	return m
}

// Validate validates this payment update
func (m *PaymentUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentUpdate) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentUpdate) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentUpdate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentUpdate) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentUpdate) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentUpdate) validateRelationships(formats strfmt.Registry) error {

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

func (m *PaymentUpdate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *PaymentUpdate) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentUpdate) UnmarshalBinary(b []byte) error {
	var res PaymentUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentUpdate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
