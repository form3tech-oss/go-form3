// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewRecall new recall
// swagger:model NewRecall
type NewRecall struct {

	// attributes
	Attributes *NewRecallAttributes `json:"attributes,omitempty"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *RecallRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

// line 140

func NewRecallWithDefaults(defaults client.Defaults) *NewRecall {
	return &NewRecall{

		Attributes: NewRecallAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUIDPtr("NewRecall", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("NewRecall", "organisation_id"),

		Relationships: RecallRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("NewRecall", "type"),

		Version: defaults.GetInt64Ptr("NewRecall", "version"),
	}
}

func (m *NewRecall) WithAttributes(attributes NewRecallAttributes) *NewRecall {

	m.Attributes = &attributes

	return m
}

func (m *NewRecall) WithoutAttributes() *NewRecall {
	m.Attributes = nil
	return m
}

func (m *NewRecall) WithID(id strfmt.UUID) *NewRecall {

	m.ID = &id

	return m
}

func (m *NewRecall) WithoutID() *NewRecall {
	m.ID = nil
	return m
}

func (m *NewRecall) WithOrganisationID(organisationID strfmt.UUID) *NewRecall {

	m.OrganisationID = &organisationID

	return m
}

func (m *NewRecall) WithoutOrganisationID() *NewRecall {
	m.OrganisationID = nil
	return m
}

func (m *NewRecall) WithRelationships(relationships RecallRelationships) *NewRecall {

	m.Relationships = &relationships

	return m
}

func (m *NewRecall) WithoutRelationships() *NewRecall {
	m.Relationships = nil
	return m
}

func (m *NewRecall) WithType(typeVar string) *NewRecall {

	m.Type = typeVar

	return m
}

func (m *NewRecall) WithVersion(version int64) *NewRecall {

	m.Version = &version

	return m
}

func (m *NewRecall) WithoutVersion() *NewRecall {
	m.Version = nil
	return m
}

// Validate validates this new recall
func (m *NewRecall) Validate(formats strfmt.Registry) error {
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

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewRecall) validateAttributes(formats strfmt.Registry) error {

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

func (m *NewRecall) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewRecall) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewRecall) validateRelationships(formats strfmt.Registry) error {

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

func (m *NewRecall) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *NewRecall) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewRecall) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewRecall) UnmarshalBinary(b []byte) error {
	var res NewRecall
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewRecall) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// NewRecallAttributes new recall attributes
// swagger:model NewRecallAttributes
type NewRecallAttributes struct {

	// Further explanation of the reason given in reason_code. Only evaluated for certain reason codes.
	Reason string `json:"reason,omitempty"`

	// The reason for the recall. Has to be a valid [recall reason code](http://api-docs.form3.tech/api.html#enumerations-recall-reason-codes).
	ReasonCode string `json:"reason_code,omitempty"`
}

// line 140

func NewRecallAttributesWithDefaults(defaults client.Defaults) *NewRecallAttributes {
	return &NewRecallAttributes{

		Reason: defaults.GetString("NewRecallAttributes", "reason"),

		ReasonCode: defaults.GetString("NewRecallAttributes", "reason_code"),
	}
}

func (m *NewRecallAttributes) WithReason(reason string) *NewRecallAttributes {

	m.Reason = reason

	return m
}

func (m *NewRecallAttributes) WithReasonCode(reasonCode string) *NewRecallAttributes {

	m.ReasonCode = reasonCode

	return m
}

// Validate validates this new recall attributes
func (m *NewRecallAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewRecallAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewRecallAttributes) UnmarshalBinary(b []byte) error {
	var res NewRecallAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewRecallAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
