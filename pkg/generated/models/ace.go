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

// Ace ace
// swagger:model Ace
type Ace struct {

	// attributes
	// Required: true
	Attributes *AceAttributes `json:"attributes"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// Name of the resource type
	// Pattern: ^[A-Za-z]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func AceWithDefaults(defaults client.Defaults) *Ace {
	return &Ace{

		Attributes: AceAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUIDPtr("Ace", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("Ace", "organisation_id"),

		Type: defaults.GetString("Ace", "type"),

		Version: defaults.GetInt64Ptr("Ace", "version"),
	}
}

func (m *Ace) WithAttributes(attributes AceAttributes) *Ace {

	m.Attributes = &attributes

	return m
}

func (m *Ace) WithoutAttributes() *Ace {
	m.Attributes = nil
	return m
}

func (m *Ace) WithID(id strfmt.UUID) *Ace {

	m.ID = &id

	return m
}

func (m *Ace) WithoutID() *Ace {
	m.ID = nil
	return m
}

func (m *Ace) WithOrganisationID(organisationID strfmt.UUID) *Ace {

	m.OrganisationID = &organisationID

	return m
}

func (m *Ace) WithoutOrganisationID() *Ace {
	m.OrganisationID = nil
	return m
}

func (m *Ace) WithType(typeVar string) *Ace {

	m.Type = typeVar

	return m
}

func (m *Ace) WithVersion(version int64) *Ace {

	m.Version = &version

	return m
}

func (m *Ace) WithoutVersion() *Ace {
	m.Version = nil
	return m
}

// Validate validates this ace
func (m *Ace) Validate(formats strfmt.Registry) error {
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

func (m *Ace) validateAttributes(formats strfmt.Registry) error {

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

func (m *Ace) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Ace) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Ace) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

func (m *Ace) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Ace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Ace) UnmarshalBinary(b []byte) error {
	var res Ace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *Ace) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// AceAttributes ace attributes
// swagger:model AceAttributes
type AceAttributes struct {

	// Action that this ACE controls
	Action string `json:"action,omitempty"`

	// filter
	Filter string `json:"filter,omitempty"`

	// Type of record that this ACE gives access to
	// Pattern: ^[A-Za-z]*$
	RecordType string `json:"record_type,omitempty"`

	// Role ID of the role that this ACE belongs to
	// Format: uuid
	RoleID strfmt.UUID `json:"role_id,omitempty"`
}

func AceAttributesWithDefaults(defaults client.Defaults) *AceAttributes {
	return &AceAttributes{

		Action: defaults.GetString("AceAttributes", "action"),

		Filter: defaults.GetString("AceAttributes", "filter"),

		RecordType: defaults.GetString("AceAttributes", "record_type"),

		RoleID: defaults.GetStrfmtUUID("AceAttributes", "role_id"),
	}
}

func (m *AceAttributes) WithAction(action string) *AceAttributes {

	m.Action = action

	return m
}

func (m *AceAttributes) WithFilter(filter string) *AceAttributes {

	m.Filter = filter

	return m
}

func (m *AceAttributes) WithRecordType(recordType string) *AceAttributes {

	m.RecordType = recordType

	return m
}

func (m *AceAttributes) WithRoleID(roleID strfmt.UUID) *AceAttributes {

	m.RoleID = roleID

	return m
}

// Validate validates this ace attributes
func (m *AceAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecordType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoleID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AceAttributes) validateRecordType(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordType) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"record_type", "body", string(m.RecordType), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AceAttributes) validateRoleID(formats strfmt.Registry) error {

	if swag.IsZero(m.RoleID) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"role_id", "body", "uuid", m.RoleID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AceAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AceAttributes) UnmarshalBinary(b []byte) error {
	var res AceAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AceAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
