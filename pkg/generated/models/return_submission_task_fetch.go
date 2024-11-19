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

// ReturnSubmissionTaskFetch return submission task fetch
// swagger:model ReturnSubmissionTaskFetch
type ReturnSubmissionTaskFetch struct {

	// attributes
	Attributes *ReturnSubmissionTaskAttributes `json:"attributes,omitempty"`

	// created on
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// modified on
	// Format: date-time
	ModifiedOn strfmt.DateTime `json:"modified_on,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// relationships
	Relationships *ReturnSubmissionTaskFetchRelationships `json:"relationships,omitempty"`

	// type
	// Enum: [return_submission_tasks]
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ReturnSubmissionTaskFetchWithDefaults(defaults client.Defaults) *ReturnSubmissionTaskFetch {
	return &ReturnSubmissionTaskFetch{

		Attributes: ReturnSubmissionTaskAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTime("ReturnSubmissionTaskFetch", "created_on"),

		ID: defaults.GetStrfmtUUID("ReturnSubmissionTaskFetch", "id"),

		ModifiedOn: defaults.GetStrfmtDateTime("ReturnSubmissionTaskFetch", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("ReturnSubmissionTaskFetch", "organisation_id"),

		Relationships: ReturnSubmissionTaskFetchRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReturnSubmissionTaskFetch", "type"),

		Version: defaults.GetInt64Ptr("ReturnSubmissionTaskFetch", "version"),
	}
}

func (m *ReturnSubmissionTaskFetch) WithAttributes(attributes ReturnSubmissionTaskAttributes) *ReturnSubmissionTaskFetch {

	m.Attributes = &attributes

	return m
}

func (m *ReturnSubmissionTaskFetch) WithoutAttributes() *ReturnSubmissionTaskFetch {
	m.Attributes = nil
	return m
}

func (m *ReturnSubmissionTaskFetch) WithCreatedOn(createdOn strfmt.DateTime) *ReturnSubmissionTaskFetch {

	m.CreatedOn = createdOn

	return m
}

func (m *ReturnSubmissionTaskFetch) WithID(id strfmt.UUID) *ReturnSubmissionTaskFetch {

	m.ID = id

	return m
}

func (m *ReturnSubmissionTaskFetch) WithModifiedOn(modifiedOn strfmt.DateTime) *ReturnSubmissionTaskFetch {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *ReturnSubmissionTaskFetch) WithOrganisationID(organisationID strfmt.UUID) *ReturnSubmissionTaskFetch {

	m.OrganisationID = organisationID

	return m
}

func (m *ReturnSubmissionTaskFetch) WithRelationships(relationships ReturnSubmissionTaskFetchRelationships) *ReturnSubmissionTaskFetch {

	m.Relationships = &relationships

	return m
}

func (m *ReturnSubmissionTaskFetch) WithoutRelationships() *ReturnSubmissionTaskFetch {
	m.Relationships = nil
	return m
}

func (m *ReturnSubmissionTaskFetch) WithType(typeVar string) *ReturnSubmissionTaskFetch {

	m.Type = typeVar

	return m
}

func (m *ReturnSubmissionTaskFetch) WithVersion(version int64) *ReturnSubmissionTaskFetch {

	m.Version = &version

	return m
}

func (m *ReturnSubmissionTaskFetch) WithoutVersion() *ReturnSubmissionTaskFetch {
	m.Version = nil
	return m
}

// Validate validates this return submission task fetch
func (m *ReturnSubmissionTaskFetch) Validate(formats strfmt.Registry) error {
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

func (m *ReturnSubmissionTaskFetch) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReturnSubmissionTaskFetch) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionTaskFetch) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionTaskFetch) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionTaskFetch) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionTaskFetch) validateRelationships(formats strfmt.Registry) error {

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

var returnSubmissionTaskFetchTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["return_submission_tasks"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		returnSubmissionTaskFetchTypeTypePropEnum = append(returnSubmissionTaskFetchTypeTypePropEnum, v)
	}
}

const (

	// ReturnSubmissionTaskFetchTypeReturnSubmissionTasks captures enum value "return_submission_tasks"
	ReturnSubmissionTaskFetchTypeReturnSubmissionTasks string = "return_submission_tasks"
)

// prop value enum
func (m *ReturnSubmissionTaskFetch) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, returnSubmissionTaskFetchTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ReturnSubmissionTaskFetch) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionTaskFetch) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnSubmissionTaskFetch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnSubmissionTaskFetch) UnmarshalBinary(b []byte) error {
	var res ReturnSubmissionTaskFetch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnSubmissionTaskFetch) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
