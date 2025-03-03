// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SubscriptionCreate subscription create
// swagger:model SubscriptionCreate
type SubscriptionCreate struct {

	// attributes
	// Required: true
	Attributes *SubscriptionAttributes `json:"attributes"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// Name of the resource type
	// Pattern: ^[A-Za-z]*$
	Type string `json:"type,omitempty"`
}

func SubscriptionCreateWithDefaults(defaults client.Defaults) *SubscriptionCreate {
	return &SubscriptionCreate{

		Attributes: SubscriptionAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUIDPtr("SubscriptionCreate", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("SubscriptionCreate", "organisation_id"),

		Type: defaults.GetString("SubscriptionCreate", "type"),
	}
}

func (m *SubscriptionCreate) WithAttributes(attributes SubscriptionAttributes) *SubscriptionCreate {

	m.Attributes = &attributes

	return m
}

func (m *SubscriptionCreate) WithoutAttributes() *SubscriptionCreate {
	m.Attributes = nil
	return m
}

func (m *SubscriptionCreate) WithID(id strfmt.UUID) *SubscriptionCreate {

	m.ID = &id

	return m
}

func (m *SubscriptionCreate) WithoutID() *SubscriptionCreate {
	m.ID = nil
	return m
}

func (m *SubscriptionCreate) WithOrganisationID(organisationID strfmt.UUID) *SubscriptionCreate {

	m.OrganisationID = &organisationID

	return m
}

func (m *SubscriptionCreate) WithoutOrganisationID() *SubscriptionCreate {
	m.OrganisationID = nil
	return m
}

func (m *SubscriptionCreate) WithType(typeVar string) *SubscriptionCreate {

	m.Type = typeVar

	return m
}

// Validate validates this subscription create
func (m *SubscriptionCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
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

func (m *SubscriptionCreate) validateAttributes(formats strfmt.Registry) error {

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

func (m *SubscriptionCreate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionCreate) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionCreate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", m.Type, `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionCreate) UnmarshalBinary(b []byte) error {
	var res SubscriptionCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SubscriptionCreate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
