// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v5/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountRequestSubmission account request submission
// swagger:model AccountRequestSubmission
type AccountRequestSubmission struct {

	// attributes
	// Required: true
	Attributes *AccountRequestSubmissionAttributes `json:"attributes"`

	// Date when the resource was created
	// Read Only: true
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID strfmt.UUID `json:"id"`

	// Date when the resource was last modified
	// Read Only: true
	// Format: date-time
	ModifiedOn strfmt.DateTime `json:"modified_on,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *AccountRequestSubmissionRelationships `json:"relationships,omitempty"`

	// type
	// Required: true
	Type ResourceType `json:"type"`

	// version
	// Required: true
	// Minimum: 0
	Version int64 `json:"version"`
}

func AccountRequestSubmissionWithDefaults(defaults client.Defaults) *AccountRequestSubmission {
	return &AccountRequestSubmission{

		Attributes: AccountRequestSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTime("AccountRequestSubmission", "created_on"),

		ID: defaults.GetStrfmtUUID("AccountRequestSubmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTime("AccountRequestSubmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("AccountRequestSubmission", "organisation_id"),

		Relationships: AccountRequestSubmissionRelationshipsWithDefaults(defaults),

		// TODO Type: ResourceType,

		Version: defaults.GetInt64("AccountRequestSubmission", "version"),
	}
}

func (m *AccountRequestSubmission) WithAttributes(attributes AccountRequestSubmissionAttributes) *AccountRequestSubmission {

	m.Attributes = &attributes

	return m
}

func (m *AccountRequestSubmission) WithoutAttributes() *AccountRequestSubmission {
	m.Attributes = nil
	return m
}

func (m *AccountRequestSubmission) WithCreatedOn(createdOn strfmt.DateTime) *AccountRequestSubmission {

	m.CreatedOn = createdOn

	return m
}

func (m *AccountRequestSubmission) WithID(id strfmt.UUID) *AccountRequestSubmission {

	m.ID = id

	return m
}

func (m *AccountRequestSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *AccountRequestSubmission {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *AccountRequestSubmission) WithOrganisationID(organisationID strfmt.UUID) *AccountRequestSubmission {

	m.OrganisationID = organisationID

	return m
}

func (m *AccountRequestSubmission) WithRelationships(relationships AccountRequestSubmissionRelationships) *AccountRequestSubmission {

	m.Relationships = &relationships

	return m
}

func (m *AccountRequestSubmission) WithoutRelationships() *AccountRequestSubmission {
	m.Relationships = nil
	return m
}

func (m *AccountRequestSubmission) WithType(typeVar ResourceType) *AccountRequestSubmission {

	m.Type = typeVar

	return m
}

func (m *AccountRequestSubmission) WithVersion(version int64) *AccountRequestSubmission {

	m.Version = version

	return m
}

// Validate validates this account request submission
func (m *AccountRequestSubmission) Validate(formats strfmt.Registry) error {
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

func (m *AccountRequestSubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *AccountRequestSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountRequestSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountRequestSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountRequestSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", strfmt.UUID(m.OrganisationID)); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountRequestSubmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *AccountRequestSubmission) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *AccountRequestSubmission) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", int64(m.Version)); err != nil {
		return err
	}

	if err := validate.MinimumInt("version", "body", int64(m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountRequestSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountRequestSubmission) UnmarshalBinary(b []byte) error {
	var res AccountRequestSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountRequestSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
