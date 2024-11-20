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

// ClaimReversalSubmission claim reversal submission
// swagger:model ClaimReversalSubmission
type ClaimReversalSubmission struct {

	// attributes
	Attributes *ClaimReversalSubmissionAttributes `json:"attributes,omitempty"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// modified on
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *ClaimReversalSubmissionRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ClaimReversalSubmissionWithDefaults(defaults client.Defaults) *ClaimReversalSubmission {
	return &ClaimReversalSubmission{

		Attributes: ClaimReversalSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("ClaimReversalSubmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("ClaimReversalSubmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("ClaimReversalSubmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("ClaimReversalSubmission", "organisation_id"),

		Relationships: ClaimReversalSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ClaimReversalSubmission", "type"),

		Version: defaults.GetInt64Ptr("ClaimReversalSubmission", "version"),
	}
}

func (m *ClaimReversalSubmission) WithAttributes(attributes ClaimReversalSubmissionAttributes) *ClaimReversalSubmission {

	m.Attributes = &attributes

	return m
}

func (m *ClaimReversalSubmission) WithoutAttributes() *ClaimReversalSubmission {
	m.Attributes = nil
	return m
}

func (m *ClaimReversalSubmission) WithCreatedOn(createdOn strfmt.DateTime) *ClaimReversalSubmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *ClaimReversalSubmission) WithoutCreatedOn() *ClaimReversalSubmission {
	m.CreatedOn = nil
	return m
}

func (m *ClaimReversalSubmission) WithID(id strfmt.UUID) *ClaimReversalSubmission {

	m.ID = &id

	return m
}

func (m *ClaimReversalSubmission) WithoutID() *ClaimReversalSubmission {
	m.ID = nil
	return m
}

func (m *ClaimReversalSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *ClaimReversalSubmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *ClaimReversalSubmission) WithoutModifiedOn() *ClaimReversalSubmission {
	m.ModifiedOn = nil
	return m
}

func (m *ClaimReversalSubmission) WithOrganisationID(organisationID strfmt.UUID) *ClaimReversalSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *ClaimReversalSubmission) WithoutOrganisationID() *ClaimReversalSubmission {
	m.OrganisationID = nil
	return m
}

func (m *ClaimReversalSubmission) WithRelationships(relationships ClaimReversalSubmissionRelationships) *ClaimReversalSubmission {

	m.Relationships = &relationships

	return m
}

func (m *ClaimReversalSubmission) WithoutRelationships() *ClaimReversalSubmission {
	m.Relationships = nil
	return m
}

func (m *ClaimReversalSubmission) WithType(typeVar string) *ClaimReversalSubmission {

	m.Type = typeVar

	return m
}

func (m *ClaimReversalSubmission) WithVersion(version int64) *ClaimReversalSubmission {

	m.Version = &version

	return m
}

func (m *ClaimReversalSubmission) WithoutVersion() *ClaimReversalSubmission {
	m.Version = nil
	return m
}

// Validate validates this claim reversal submission
func (m *ClaimReversalSubmission) Validate(formats strfmt.Registry) error {
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

func (m *ClaimReversalSubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *ClaimReversalSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClaimReversalSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClaimReversalSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClaimReversalSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClaimReversalSubmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *ClaimReversalSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ClaimReversalSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimReversalSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimReversalSubmission) UnmarshalBinary(b []byte) error {
	var res ClaimReversalSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimReversalSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ClaimReversalSubmissionAttributes claim reversal submission attributes
// swagger:model ClaimReversalSubmissionAttributes
type ClaimReversalSubmissionAttributes struct {

	// status
	Status ClaimSubmissionStatus `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"status_reason,omitempty"`

	// submission datetime
	// Read Only: true
	// Format: date-time
	SubmissionDatetime strfmt.DateTime `json:"submission_datetime,omitempty"`
}

func ClaimReversalSubmissionAttributesWithDefaults(defaults client.Defaults) *ClaimReversalSubmissionAttributes {
	return &ClaimReversalSubmissionAttributes{

		// TODO Status: ClaimSubmissionStatus,

		StatusReason: defaults.GetString("ClaimReversalSubmissionAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTime("ClaimReversalSubmissionAttributes", "submission_datetime"),
	}
}

func (m *ClaimReversalSubmissionAttributes) WithStatus(status ClaimSubmissionStatus) *ClaimReversalSubmissionAttributes {

	m.Status = status

	return m
}

func (m *ClaimReversalSubmissionAttributes) WithStatusReason(statusReason string) *ClaimReversalSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *ClaimReversalSubmissionAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *ClaimReversalSubmissionAttributes {

	m.SubmissionDatetime = submissionDatetime

	return m
}

// Validate validates this claim reversal submission attributes
func (m *ClaimReversalSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimReversalSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

func (m *ClaimReversalSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClaimReversalSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimReversalSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res ClaimReversalSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimReversalSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
