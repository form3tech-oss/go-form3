// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v4/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Notification notification
// swagger:model Notification
type Notification struct {

	// The full resource itself (as you would see from a GET request)
	Data interface{} `json:"data,omitempty"`

	// Internal representation of the record type. Field values are subject to frequent change, evaluation of this field is discouraged.
	// Pattern: ^[A-Za-z]*$
	DataRecordType string `json:"data_record_type,omitempty"`

	// The type of event
	// Pattern: ^[a-z]*$
	EventType string `json:"event_type,omitempty"`

	// Unique resource ID
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// The type of resource contained in `data`
	// Pattern: ^[A-Za-z]*$
	RecordType string `json:"record_type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func NotificationWithDefaults(defaults client.Defaults) *Notification {
	return &Notification{

		// TODO Data: interface{},

		DataRecordType: defaults.GetString("Notification", "data_record_type"),

		EventType: defaults.GetString("Notification", "event_type"),

		ID: defaults.GetStrfmtUUID("Notification", "id"),

		OrganisationID: defaults.GetStrfmtUUID("Notification", "organisation_id"),

		RecordType: defaults.GetString("Notification", "record_type"),

		Version: defaults.GetInt64Ptr("Notification", "version"),
	}
}

func (m *Notification) WithData(data interface{}) *Notification {

	m.Data = data

	return m
}

func (m *Notification) WithDataRecordType(dataRecordType string) *Notification {

	m.DataRecordType = dataRecordType

	return m
}

func (m *Notification) WithEventType(eventType string) *Notification {

	m.EventType = eventType

	return m
}

func (m *Notification) WithID(id strfmt.UUID) *Notification {

	m.ID = id

	return m
}

func (m *Notification) WithOrganisationID(organisationID strfmt.UUID) *Notification {

	m.OrganisationID = organisationID

	return m
}

func (m *Notification) WithRecordType(recordType string) *Notification {

	m.RecordType = recordType

	return m
}

func (m *Notification) WithVersion(version int64) *Notification {

	m.Version = &version

	return m
}

func (m *Notification) WithoutVersion() *Notification {
	m.Version = nil
	return m
}

// Validate validates this notification
func (m *Notification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataRecordType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordType(formats); err != nil {
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

func (m *Notification) validateDataRecordType(formats strfmt.Registry) error {

	if swag.IsZero(m.DataRecordType) { // not required
		return nil
	}

	if err := validate.Pattern("data_record_type", "body", string(m.DataRecordType), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateEventType(formats strfmt.Registry) error {

	if swag.IsZero(m.EventType) { // not required
		return nil
	}

	if err := validate.Pattern("event_type", "body", string(m.EventType), `^[a-z]*$`); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateRecordType(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordType) { // not required
		return nil
	}

	if err := validate.Pattern("record_type", "body", string(m.RecordType), `^[A-Za-z]*$`); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Notification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Notification) UnmarshalBinary(b []byte) error {
	var res Notification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *Notification) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
