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

// NewQuery new query
// swagger:model NewQuery
type NewQuery struct {

	// attributes
	// Required: true
	Attributes *QueryAttributes `json:"attributes"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *NewQueryRelationships `json:"relationships,omitempty"`

	// type
	// Required: true
	Type *QueryResourceType `json:"type"`
}

func NewQueryWithDefaults(defaults client.Defaults) *NewQuery {
	return &NewQuery{

		Attributes: QueryAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUIDPtr("NewQuery", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("NewQuery", "organisation_id"),

		Relationships: NewQueryRelationshipsWithDefaults(defaults),

		// TODO Type: QueryResourceType,

	}
}

func (m *NewQuery) WithAttributes(attributes QueryAttributes) *NewQuery {

	m.Attributes = &attributes

	return m
}

func (m *NewQuery) WithoutAttributes() *NewQuery {
	m.Attributes = nil
	return m
}

func (m *NewQuery) WithID(id strfmt.UUID) *NewQuery {

	m.ID = &id

	return m
}

func (m *NewQuery) WithoutID() *NewQuery {
	m.ID = nil
	return m
}

func (m *NewQuery) WithOrganisationID(organisationID strfmt.UUID) *NewQuery {

	m.OrganisationID = &organisationID

	return m
}

func (m *NewQuery) WithoutOrganisationID() *NewQuery {
	m.OrganisationID = nil
	return m
}

func (m *NewQuery) WithRelationships(relationships NewQueryRelationships) *NewQuery {

	m.Relationships = &relationships

	return m
}

func (m *NewQuery) WithoutRelationships() *NewQuery {
	m.Relationships = nil
	return m
}

func (m *NewQuery) WithType(typeVar QueryResourceType) *NewQuery {

	m.Type = &typeVar

	return m
}

func (m *NewQuery) WithoutType() *NewQuery {
	m.Type = nil
	return m
}

// Validate validates this new query
func (m *NewQuery) Validate(formats strfmt.Registry) error {
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

	if err := m.validateRelationships(formats); err != nil {
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

func (m *NewQuery) validateAttributes(formats strfmt.Registry) error {

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

func (m *NewQuery) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewQuery) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewQuery) validateRelationships(formats strfmt.Registry) error {

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

func (m *NewQuery) validateType(formats strfmt.Registry) error {

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
func (m *NewQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewQuery) UnmarshalBinary(b []byte) error {
	var res NewQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewQuery) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
