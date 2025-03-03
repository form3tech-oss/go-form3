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

// NewPaymentSubmission new payment submission
// swagger:model NewPaymentSubmission
type NewPaymentSubmission struct {

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *NewPaymentSubmissionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func NewPaymentSubmissionWithDefaults(defaults client.Defaults) *NewPaymentSubmission {
	return &NewPaymentSubmission{

		ID: defaults.GetStrfmtUUIDPtr("NewPaymentSubmission", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("NewPaymentSubmission", "organisation_id"),

		Relationships: NewPaymentSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("NewPaymentSubmission", "type"),

		Version: defaults.GetInt64Ptr("NewPaymentSubmission", "version"),
	}
}

func (m *NewPaymentSubmission) WithID(id strfmt.UUID) *NewPaymentSubmission {

	m.ID = &id

	return m
}

func (m *NewPaymentSubmission) WithoutID() *NewPaymentSubmission {
	m.ID = nil
	return m
}

func (m *NewPaymentSubmission) WithOrganisationID(organisationID strfmt.UUID) *NewPaymentSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *NewPaymentSubmission) WithoutOrganisationID() *NewPaymentSubmission {
	m.OrganisationID = nil
	return m
}

func (m *NewPaymentSubmission) WithRelationships(relationships NewPaymentSubmissionRelationships) *NewPaymentSubmission {

	m.Relationships = &relationships

	return m
}

func (m *NewPaymentSubmission) WithoutRelationships() *NewPaymentSubmission {
	m.Relationships = nil
	return m
}

func (m *NewPaymentSubmission) WithType(typeVar string) *NewPaymentSubmission {

	m.Type = typeVar

	return m
}

func (m *NewPaymentSubmission) WithVersion(version int64) *NewPaymentSubmission {

	m.Version = &version

	return m
}

func (m *NewPaymentSubmission) WithoutVersion() *NewPaymentSubmission {
	m.Version = nil
	return m
}

// Validate validates this new payment submission
func (m *NewPaymentSubmission) Validate(formats strfmt.Registry) error {
	var res []error

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

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPaymentSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewPaymentSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewPaymentSubmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *NewPaymentSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", m.Type, `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *NewPaymentSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPaymentSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPaymentSubmission) UnmarshalBinary(b []byte) error {
	var res NewPaymentSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewPaymentSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// NewPaymentSubmissionRelationships new payment submission relationships
// swagger:model NewPaymentSubmissionRelationships
type NewPaymentSubmissionRelationships struct {

	// validations
	Validations *RelationshipLinks `json:"validations,omitempty"`
}

func NewPaymentSubmissionRelationshipsWithDefaults(defaults client.Defaults) *NewPaymentSubmissionRelationships {
	return &NewPaymentSubmissionRelationships{

		Validations: RelationshipLinksWithDefaults(defaults),
	}
}

func (m *NewPaymentSubmissionRelationships) WithValidations(validations RelationshipLinks) *NewPaymentSubmissionRelationships {

	m.Validations = &validations

	return m
}

func (m *NewPaymentSubmissionRelationships) WithoutValidations() *NewPaymentSubmissionRelationships {
	m.Validations = nil
	return m
}

// Validate validates this new payment submission relationships
func (m *NewPaymentSubmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValidations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewPaymentSubmissionRelationships) validateValidations(formats strfmt.Registry) error {

	if swag.IsZero(m.Validations) { // not required
		return nil
	}

	if m.Validations != nil {
		if err := m.Validations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "validations")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewPaymentSubmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPaymentSubmissionRelationships) UnmarshalBinary(b []byte) error {
	var res NewPaymentSubmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewPaymentSubmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
