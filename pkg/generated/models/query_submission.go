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

// QuerySubmission query submission
// swagger:model QuerySubmission
type QuerySubmission struct {

	// type
	// Required: true
	Type *QuerySubmissionResourceType `json:"type"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// version
	// Required: true
	// Minimum: 0
	Version *int64 `json:"version"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// created on
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// modified on
	// Format: date-time
	ModifiedOn strfmt.DateTime `json:"modified_on,omitempty"`

	// attributes
	// Required: true
	Attributes *QuerySubmissionAttributes `json:"attributes"`

	// relationships
	Relationships *QuerySubmissionRelationships `json:"relationships,omitempty"`
}

func QuerySubmissionWithDefaults(defaults client.Defaults) *QuerySubmission {
	return &QuerySubmission{

		// TODO Type: QuerySubmissionResourceType,

		ID: defaults.GetStrfmtUUIDPtr("QuerySubmission", "id"),

		Version: defaults.GetInt64Ptr("QuerySubmission", "version"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("QuerySubmission", "organisation_id"),

		CreatedOn: defaults.GetStrfmtDateTime("QuerySubmission", "created_on"),

		ModifiedOn: defaults.GetStrfmtDateTime("QuerySubmission", "modified_on"),

		Attributes: QuerySubmissionAttributesWithDefaults(defaults),

		Relationships: QuerySubmissionRelationshipsWithDefaults(defaults),
	}
}

func (m *QuerySubmission) WithType(typeVar QuerySubmissionResourceType) *QuerySubmission {

	m.Type = &typeVar

	return m
}

func (m *QuerySubmission) WithoutType() *QuerySubmission {
	m.Type = nil
	return m
}

func (m *QuerySubmission) WithID(id strfmt.UUID) *QuerySubmission {

	m.ID = &id

	return m
}

func (m *QuerySubmission) WithoutID() *QuerySubmission {
	m.ID = nil
	return m
}

func (m *QuerySubmission) WithVersion(version int64) *QuerySubmission {

	m.Version = &version

	return m
}

func (m *QuerySubmission) WithoutVersion() *QuerySubmission {
	m.Version = nil
	return m
}

func (m *QuerySubmission) WithOrganisationID(organisationID strfmt.UUID) *QuerySubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *QuerySubmission) WithoutOrganisationID() *QuerySubmission {
	m.OrganisationID = nil
	return m
}

func (m *QuerySubmission) WithCreatedOn(createdOn strfmt.DateTime) *QuerySubmission {

	m.CreatedOn = createdOn

	return m
}

func (m *QuerySubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *QuerySubmission {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *QuerySubmission) WithAttributes(attributes QuerySubmissionAttributes) *QuerySubmission {

	m.Attributes = &attributes

	return m
}

func (m *QuerySubmission) WithoutAttributes() *QuerySubmission {
	m.Attributes = nil
	return m
}

func (m *QuerySubmission) WithRelationships(relationships QuerySubmissionRelationships) *QuerySubmission {

	m.Relationships = &relationships

	return m
}

func (m *QuerySubmission) WithoutRelationships() *QuerySubmission {
	m.Relationships = nil
	return m
}

// Validate validates this query submission
func (m *QuerySubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuerySubmission) validateType(formats strfmt.Registry) error {

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

func (m *QuerySubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QuerySubmission) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *QuerySubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QuerySubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QuerySubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QuerySubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *QuerySubmission) validateRelationships(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *QuerySubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuerySubmission) UnmarshalBinary(b []byte) error {
	var res QuerySubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QuerySubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
