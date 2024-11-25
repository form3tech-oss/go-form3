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

// Branch branch
// swagger:model Branch
type Branch struct {

	// attributes
	// Required: true
	Attributes *BranchAttributes `json:"attributes"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// modified on
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

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

func BranchWithDefaults(defaults client.Defaults) *Branch {
	return &Branch{

		Attributes: BranchAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("Branch", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("Branch", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("Branch", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("Branch", "organisation_id"),

		Type: defaults.GetString("Branch", "type"),

		Version: defaults.GetInt64Ptr("Branch", "version"),
	}
}

func (m *Branch) WithAttributes(attributes BranchAttributes) *Branch {

	m.Attributes = &attributes

	return m
}

func (m *Branch) WithoutAttributes() *Branch {
	m.Attributes = nil
	return m
}

func (m *Branch) WithCreatedOn(createdOn strfmt.DateTime) *Branch {

	m.CreatedOn = &createdOn

	return m
}

func (m *Branch) WithoutCreatedOn() *Branch {
	m.CreatedOn = nil
	return m
}

func (m *Branch) WithID(id strfmt.UUID) *Branch {

	m.ID = &id

	return m
}

func (m *Branch) WithoutID() *Branch {
	m.ID = nil
	return m
}

func (m *Branch) WithModifiedOn(modifiedOn strfmt.DateTime) *Branch {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *Branch) WithoutModifiedOn() *Branch {
	m.ModifiedOn = nil
	return m
}

func (m *Branch) WithOrganisationID(organisationID strfmt.UUID) *Branch {

	m.OrganisationID = &organisationID

	return m
}

func (m *Branch) WithoutOrganisationID() *Branch {
	m.OrganisationID = nil
	return m
}

func (m *Branch) WithType(typeVar string) *Branch {

	m.Type = typeVar

	return m
}

func (m *Branch) WithVersion(version int64) *Branch {

	m.Version = &version

	return m
}

func (m *Branch) WithoutVersion() *Branch {
	m.Version = nil
	return m
}

// Validate validates this branch
func (m *Branch) Validate(formats strfmt.Registry) error {
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

func (m *Branch) validateAttributes(formats strfmt.Registry) error {

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

func (m *Branch) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Branch) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Branch) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Branch) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Branch) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *Branch) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Branch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Branch) UnmarshalBinary(b []byte) error {
	var res Branch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *Branch) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
