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

// ContactAccount contact account
// swagger:model ContactAccount
type ContactAccount struct {

	// attributes
	// Required: true
	Attributes *ContactAccountAttributes `json:"attributes"`

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
	Relationships *ContactAccountRelationships `json:"relationships,omitempty"`

	// type
	Type ContactAccountResourceType `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ContactAccountWithDefaults(defaults client.Defaults) *ContactAccount {
	return &ContactAccount{

		Attributes: ContactAccountAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTime("ContactAccount", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("ContactAccount", "id"),

		ModifiedOn: defaults.GetStrfmtDateTime("ContactAccount", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("ContactAccount", "organisation_id"),

		Relationships: ContactAccountRelationshipsWithDefaults(defaults),

		// TODO Type: ContactAccountResourceType,

		Version: defaults.GetInt64Ptr("ContactAccount", "version"),
	}
}

func (m *ContactAccount) WithAttributes(attributes ContactAccountAttributes) *ContactAccount {

	m.Attributes = &attributes

	return m
}

func (m *ContactAccount) WithoutAttributes() *ContactAccount {
	m.Attributes = nil
	return m
}

func (m *ContactAccount) WithCreatedOn(createdOn strfmt.DateTime) *ContactAccount {

	m.CreatedOn = createdOn

	return m
}

func (m *ContactAccount) WithID(id strfmt.UUID) *ContactAccount {

	m.ID = &id

	return m
}

func (m *ContactAccount) WithoutID() *ContactAccount {
	m.ID = nil
	return m
}

func (m *ContactAccount) WithModifiedOn(modifiedOn strfmt.DateTime) *ContactAccount {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *ContactAccount) WithOrganisationID(organisationID strfmt.UUID) *ContactAccount {

	m.OrganisationID = &organisationID

	return m
}

func (m *ContactAccount) WithoutOrganisationID() *ContactAccount {
	m.OrganisationID = nil
	return m
}

func (m *ContactAccount) WithRelationships(relationships ContactAccountRelationships) *ContactAccount {

	m.Relationships = &relationships

	return m
}

func (m *ContactAccount) WithoutRelationships() *ContactAccount {
	m.Relationships = nil
	return m
}

func (m *ContactAccount) WithType(typeVar ContactAccountResourceType) *ContactAccount {

	m.Type = typeVar

	return m
}

func (m *ContactAccount) WithVersion(version int64) *ContactAccount {

	m.Version = &version

	return m
}

func (m *ContactAccount) WithoutVersion() *ContactAccount {
	m.Version = nil
	return m
}

// Validate validates this contact account
func (m *ContactAccount) Validate(formats strfmt.Registry) error {
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

func (m *ContactAccount) validateAttributes(formats strfmt.Registry) error {

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

func (m *ContactAccount) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "datetime", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContactAccount) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContactAccount) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "datetime", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContactAccount) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContactAccount) validateRelationships(formats strfmt.Registry) error {

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

func (m *ContactAccount) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *ContactAccount) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactAccount) UnmarshalBinary(b []byte) error {
	var res ContactAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ContactAccount) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
