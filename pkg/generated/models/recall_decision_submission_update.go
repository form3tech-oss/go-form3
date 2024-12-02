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

// RecallDecisionSubmissionUpdate recall decision submission update
// swagger:model RecallDecisionSubmissionUpdate
type RecallDecisionSubmissionUpdate struct {

	// attributes
	Attributes *RecallDecisionSubmissionUpdateAttributes `json:"attributes,omitempty"`

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

func RecallDecisionSubmissionUpdateWithDefaults(defaults client.Defaults) *RecallDecisionSubmissionUpdate {
	return &RecallDecisionSubmissionUpdate{

		Attributes: RecallDecisionSubmissionUpdateAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUIDPtr("RecallDecisionSubmissionUpdate", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("RecallDecisionSubmissionUpdate", "organisation_id"),

		Type: defaults.GetString("RecallDecisionSubmissionUpdate", "type"),

		Version: defaults.GetInt64Ptr("RecallDecisionSubmissionUpdate", "version"),
	}
}

func (m *RecallDecisionSubmissionUpdate) WithAttributes(attributes RecallDecisionSubmissionUpdateAttributes) *RecallDecisionSubmissionUpdate {

	m.Attributes = &attributes

	return m
}

func (m *RecallDecisionSubmissionUpdate) WithoutAttributes() *RecallDecisionSubmissionUpdate {
	m.Attributes = nil
	return m
}

func (m *RecallDecisionSubmissionUpdate) WithID(id strfmt.UUID) *RecallDecisionSubmissionUpdate {

	m.ID = &id

	return m
}

func (m *RecallDecisionSubmissionUpdate) WithoutID() *RecallDecisionSubmissionUpdate {
	m.ID = nil
	return m
}

func (m *RecallDecisionSubmissionUpdate) WithOrganisationID(organisationID strfmt.UUID) *RecallDecisionSubmissionUpdate {

	m.OrganisationID = &organisationID

	return m
}

func (m *RecallDecisionSubmissionUpdate) WithoutOrganisationID() *RecallDecisionSubmissionUpdate {
	m.OrganisationID = nil
	return m
}

func (m *RecallDecisionSubmissionUpdate) WithType(typeVar string) *RecallDecisionSubmissionUpdate {

	m.Type = typeVar

	return m
}

func (m *RecallDecisionSubmissionUpdate) WithVersion(version int64) *RecallDecisionSubmissionUpdate {

	m.Version = &version

	return m
}

func (m *RecallDecisionSubmissionUpdate) WithoutVersion() *RecallDecisionSubmissionUpdate {
	m.Version = nil
	return m
}

// Validate validates this recall decision submission update
func (m *RecallDecisionSubmissionUpdate) Validate(formats strfmt.Registry) error {
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

func (m *RecallDecisionSubmissionUpdate) validateAttributes(formats strfmt.Registry) error {

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

func (m *RecallDecisionSubmissionUpdate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionSubmissionUpdate) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionSubmissionUpdate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", m.Type, `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *RecallDecisionSubmissionUpdate) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallDecisionSubmissionUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecisionSubmissionUpdate) UnmarshalBinary(b []byte) error {
	var res RecallDecisionSubmissionUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallDecisionSubmissionUpdate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// RecallDecisionSubmissionUpdateAttributes recall decision submission update attributes
// swagger:model RecallDecisionSubmissionUpdateAttributes
type RecallDecisionSubmissionUpdateAttributes struct {

	// Additional payment reference assigned by the scheme
	ReferenceID string `json:"reference_id,omitempty"`

	// status
	Status RecallDecisionSubmissionStatus `json:"status,omitempty"`

	// Reason for submission failure if status is `delivery_failed`
	StatusReason string `json:"status_reason,omitempty"`
}

func RecallDecisionSubmissionUpdateAttributesWithDefaults(defaults client.Defaults) *RecallDecisionSubmissionUpdateAttributes {
	return &RecallDecisionSubmissionUpdateAttributes{

		ReferenceID: defaults.GetString("RecallDecisionSubmissionUpdateAttributes", "reference_id"),

		// TODO Status: RecallDecisionSubmissionStatus,

		StatusReason: defaults.GetString("RecallDecisionSubmissionUpdateAttributes", "status_reason"),
	}
}

func (m *RecallDecisionSubmissionUpdateAttributes) WithReferenceID(referenceID string) *RecallDecisionSubmissionUpdateAttributes {

	m.ReferenceID = referenceID

	return m
}

func (m *RecallDecisionSubmissionUpdateAttributes) WithStatus(status RecallDecisionSubmissionStatus) *RecallDecisionSubmissionUpdateAttributes {

	m.Status = status

	return m
}

func (m *RecallDecisionSubmissionUpdateAttributes) WithStatusReason(statusReason string) *RecallDecisionSubmissionUpdateAttributes {

	m.StatusReason = statusReason

	return m
}

// Validate validates this recall decision submission update attributes
func (m *RecallDecisionSubmissionUpdateAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecallDecisionSubmissionUpdateAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallDecisionSubmissionUpdateAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecisionSubmissionUpdateAttributes) UnmarshalBinary(b []byte) error {
	var res RecallDecisionSubmissionUpdateAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallDecisionSubmissionUpdateAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
