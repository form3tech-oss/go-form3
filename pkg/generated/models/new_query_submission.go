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

// NewQuerySubmission new query submission
// swagger:model NewQuerySubmission
type NewQuerySubmission struct {

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// type
	// Required: true
	Type *QuerySubmissionResourceType `json:"type"`
}

func NewQuerySubmissionWithDefaults(defaults client.Defaults) *NewQuerySubmission {
	return &NewQuerySubmission{

		ID: defaults.GetStrfmtUUIDPtr("NewQuerySubmission", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("NewQuerySubmission", "organisation_id"),

		// TODO Type: QuerySubmissionResourceType,

	}
}

func (m *NewQuerySubmission) WithID(id strfmt.UUID) *NewQuerySubmission {

	m.ID = &id

	return m
}

func (m *NewQuerySubmission) WithoutID() *NewQuerySubmission {
	m.ID = nil
	return m
}

func (m *NewQuerySubmission) WithOrganisationID(organisationID strfmt.UUID) *NewQuerySubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *NewQuerySubmission) WithoutOrganisationID() *NewQuerySubmission {
	m.OrganisationID = nil
	return m
}

func (m *NewQuerySubmission) WithType(typeVar QuerySubmissionResourceType) *NewQuerySubmission {

	m.Type = &typeVar

	return m
}

func (m *NewQuerySubmission) WithoutType() *NewQuerySubmission {
	m.Type = nil
	return m
}

// Validate validates this new query submission
func (m *NewQuerySubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewQuerySubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewQuerySubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewQuerySubmission) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewQuerySubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewQuerySubmission) UnmarshalBinary(b []byte) error {
	var res NewQuerySubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewQuerySubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
