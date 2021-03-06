// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Party party
// swagger:model Party
type Party struct {

	// attributes
	// Required: true
	Attributes *PartyAttributes `json:"attributes"`

	// created on
	// Format: datetime
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// modified on
	// Format: datetime
	ModifiedOn strfmt.DateTime `json:"modified_on,omitempty"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *PartyRelationships `json:"relationships,omitempty"`

	// type
	// Enum: [parties]
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func PartyWithDefaults(defaults client.Defaults) *Party {
	return &Party{

		Attributes: PartyAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTime("Party", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("Party", "id"),

		ModifiedOn: defaults.GetStrfmtDateTime("Party", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("Party", "organisation_id"),

		Relationships: PartyRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("Party", "type"),

		Version: defaults.GetInt64Ptr("Party", "version"),
	}
}

func (m *Party) WithAttributes(attributes PartyAttributes) *Party {

	m.Attributes = &attributes

	return m
}

func (m *Party) WithoutAttributes() *Party {
	m.Attributes = nil
	return m
}

func (m *Party) WithCreatedOn(createdOn strfmt.DateTime) *Party {

	m.CreatedOn = createdOn

	return m
}

func (m *Party) WithID(id strfmt.UUID) *Party {

	m.ID = &id

	return m
}

func (m *Party) WithoutID() *Party {
	m.ID = nil
	return m
}

func (m *Party) WithModifiedOn(modifiedOn strfmt.DateTime) *Party {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *Party) WithOrganisationID(organisationID strfmt.UUID) *Party {

	m.OrganisationID = &organisationID

	return m
}

func (m *Party) WithoutOrganisationID() *Party {
	m.OrganisationID = nil
	return m
}

func (m *Party) WithRelationships(relationships PartyRelationships) *Party {

	m.Relationships = &relationships

	return m
}

func (m *Party) WithoutRelationships() *Party {
	m.Relationships = nil
	return m
}

func (m *Party) WithType(typeVar string) *Party {

	m.Type = typeVar

	return m
}

func (m *Party) WithVersion(version int64) *Party {

	m.Version = &version

	return m
}

func (m *Party) WithoutVersion() *Party {
	m.Version = nil
	return m
}

// Validate validates this party
func (m *Party) Validate(formats strfmt.Registry) error {
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

func (m *Party) validateAttributes(formats strfmt.Registry) error {

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

func (m *Party) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "datetime", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Party) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Party) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "datetime", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Party) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Party) validateRelationships(formats strfmt.Registry) error {

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

var partyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["parties"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		partyTypeTypePropEnum = append(partyTypeTypePropEnum, v)
	}
}

const (

	// PartyTypeParties captures enum value "parties"
	PartyTypeParties string = "parties"
)

// prop value enum
func (m *Party) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, partyTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Party) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Party) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Party) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Party) UnmarshalBinary(b []byte) error {
	var res Party
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *Party) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
