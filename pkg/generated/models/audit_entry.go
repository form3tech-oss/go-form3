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

// AuditEntry audit entry
// swagger:model AuditEntry
type AuditEntry struct {

	// attributes
	Attributes *AuditEntryAttributes `json:"attributes,omitempty"`

	// Unique resource ID
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func AuditEntryWithDefaults(defaults client.Defaults) *AuditEntry {
	return &AuditEntry{

		Attributes: AuditEntryAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUID("AuditEntry", "id"),

		OrganisationID: defaults.GetStrfmtUUID("AuditEntry", "organisation_id"),

		Type: defaults.GetString("AuditEntry", "type"),

		Version: defaults.GetInt64Ptr("AuditEntry", "version"),
	}
}

func (m *AuditEntry) WithAttributes(attributes AuditEntryAttributes) *AuditEntry {

	m.Attributes = &attributes

	return m
}

func (m *AuditEntry) WithoutAttributes() *AuditEntry {
	m.Attributes = nil
	return m
}

func (m *AuditEntry) WithID(id strfmt.UUID) *AuditEntry {

	m.ID = id

	return m
}

func (m *AuditEntry) WithOrganisationID(organisationID strfmt.UUID) *AuditEntry {

	m.OrganisationID = organisationID

	return m
}

func (m *AuditEntry) WithType(typeVar string) *AuditEntry {

	m.Type = typeVar

	return m
}

func (m *AuditEntry) WithVersion(version int64) *AuditEntry {

	m.Version = &version

	return m
}

func (m *AuditEntry) WithoutVersion() *AuditEntry {
	m.Version = nil
	return m
}

// Validate validates this audit entry
func (m *AuditEntry) Validate(formats strfmt.Registry) error {
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

func (m *AuditEntry) validateAttributes(formats strfmt.Registry) error {

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

func (m *AuditEntry) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuditEntry) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuditEntry) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AuditEntry) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditEntry) UnmarshalBinary(b []byte) error {
	var res AuditEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AuditEntry) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// AuditEntryAttributes audit entry attributes
// swagger:model AuditEntryAttributes
type AuditEntryAttributes struct {

	// Timestamp when the change was requested
	// Format: date-time
	ActionTime strfmt.DateTime `json:"action_time,omitempty"`

	// User ID of the user who requested the change
	// Format: uuid
	ActionedBy strfmt.UUID `json:"actioned_by,omitempty"`

	// Snapshot of what the data would be after the change (empty if the request was to DELETE a record)
	AfterData interface{} `json:"after_data,omitempty"`

	// Snapshot of the data before the change (empty if the request was to CREATE a new record)
	BeforeData interface{} `json:"before_data,omitempty"`

	// Textual description of the change being made
	// Pattern: ^[A-Za-z0-9 .,@:]*$
	Description string `json:"description,omitempty"`

	// ID of the resource that is being changed
	// Format: uuid
	RecordID strfmt.UUID `json:"record_id,omitempty"`

	// Type of the resource that is being changed
	// Pattern: ^[A-Za-z]*$
	RecordType string `json:"record_type,omitempty"`
}

func AuditEntryAttributesWithDefaults(defaults client.Defaults) *AuditEntryAttributes {
	return &AuditEntryAttributes{

		ActionTime: defaults.GetStrfmtDateTime("AuditEntryAttributes", "action_time"),

		ActionedBy: defaults.GetStrfmtUUID("AuditEntryAttributes", "actioned_by"),

		// TODO AfterData: interface{},

		// TODO BeforeData: interface{},

		Description: defaults.GetString("AuditEntryAttributes", "description"),

		RecordID: defaults.GetStrfmtUUID("AuditEntryAttributes", "record_id"),

		RecordType: defaults.GetString("AuditEntryAttributes", "record_type"),
	}
}

func (m *AuditEntryAttributes) WithActionTime(actionTime strfmt.DateTime) *AuditEntryAttributes {

	m.ActionTime = actionTime

	return m
}

func (m *AuditEntryAttributes) WithActionedBy(actionedBy strfmt.UUID) *AuditEntryAttributes {

	m.ActionedBy = actionedBy

	return m
}

func (m *AuditEntryAttributes) WithAfterData(afterData interface{}) *AuditEntryAttributes {

	m.AfterData = afterData

	return m
}

func (m *AuditEntryAttributes) WithBeforeData(beforeData interface{}) *AuditEntryAttributes {

	m.BeforeData = beforeData

	return m
}

func (m *AuditEntryAttributes) WithDescription(description string) *AuditEntryAttributes {

	m.Description = description

	return m
}

func (m *AuditEntryAttributes) WithRecordID(recordID strfmt.UUID) *AuditEntryAttributes {

	m.RecordID = recordID

	return m
}

func (m *AuditEntryAttributes) WithRecordType(recordType string) *AuditEntryAttributes {

	m.RecordType = recordType

	return m
}

// Validate validates this audit entry attributes
func (m *AuditEntryAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActionedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditEntryAttributes) validateActionTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionTime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"action_time", "body", "date-time", m.ActionTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuditEntryAttributes) validateActionedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionedBy) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"actioned_by", "body", "uuid", m.ActionedBy.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuditEntryAttributes) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"description", "body", string(m.Description), `^[A-Za-z0-9 .,@:]*$`); err != nil {
		return err
	}

	return nil
}

func (m *AuditEntryAttributes) validateRecordID(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordID) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"record_id", "body", "uuid", m.RecordID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuditEntryAttributes) validateRecordType(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordType) { // not required
		return nil
	}

	if err := validate.Pattern("attributes"+"."+"record_type", "body", string(m.RecordType), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditEntryAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditEntryAttributes) UnmarshalBinary(b []byte) error {
	var res AuditEntryAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AuditEntryAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
