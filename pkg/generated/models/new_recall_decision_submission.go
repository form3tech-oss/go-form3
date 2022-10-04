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

// NewRecallDecisionSubmission new recall decision submission
// swagger:model NewRecallDecisionSubmission
type NewRecallDecisionSubmission struct {

	// attributes
	Attributes *RecallDecisionSubmissionAttributes `json:"attributes,omitempty"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func NewRecallDecisionSubmissionWithDefaults(defaults client.Defaults) *NewRecallDecisionSubmission {
	return &NewRecallDecisionSubmission{

		Attributes: RecallDecisionSubmissionAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUIDPtr("NewRecallDecisionSubmission", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("NewRecallDecisionSubmission", "organisation_id"),

		Type: defaults.GetString("NewRecallDecisionSubmission", "type"),

		Version: defaults.GetInt64Ptr("NewRecallDecisionSubmission", "version"),
	}
}

func (m *NewRecallDecisionSubmission) WithAttributes(attributes RecallDecisionSubmissionAttributes) *NewRecallDecisionSubmission {

	m.Attributes = &attributes

	return m
}

func (m *NewRecallDecisionSubmission) WithoutAttributes() *NewRecallDecisionSubmission {
	m.Attributes = nil
	return m
}

func (m *NewRecallDecisionSubmission) WithID(id strfmt.UUID) *NewRecallDecisionSubmission {

	m.ID = &id

	return m
}

func (m *NewRecallDecisionSubmission) WithoutID() *NewRecallDecisionSubmission {
	m.ID = nil
	return m
}

func (m *NewRecallDecisionSubmission) WithOrganisationID(organisationID strfmt.UUID) *NewRecallDecisionSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *NewRecallDecisionSubmission) WithoutOrganisationID() *NewRecallDecisionSubmission {
	m.OrganisationID = nil
	return m
}

func (m *NewRecallDecisionSubmission) WithType(typeVar string) *NewRecallDecisionSubmission {

	m.Type = typeVar

	return m
}

func (m *NewRecallDecisionSubmission) WithVersion(version int64) *NewRecallDecisionSubmission {

	m.Version = &version

	return m
}

func (m *NewRecallDecisionSubmission) WithoutVersion() *NewRecallDecisionSubmission {
	m.Version = nil
	return m
}

// Validate validates this new recall decision submission
func (m *NewRecallDecisionSubmission) Validate(formats strfmt.Registry) error {
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

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewRecallDecisionSubmission) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
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

func (m *NewRecallDecisionSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewRecallDecisionSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewRecallDecisionSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *NewRecallDecisionSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewRecallDecisionSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewRecallDecisionSubmission) UnmarshalBinary(b []byte) error {
	var res NewRecallDecisionSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewRecallDecisionSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
