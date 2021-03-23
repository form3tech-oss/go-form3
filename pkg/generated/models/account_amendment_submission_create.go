// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v2/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountAmendmentSubmissionCreate account amendment submission create
// swagger:model AccountAmendmentSubmissionCreate
type AccountAmendmentSubmissionCreate struct {

	// id
	// Required: true
	// Format: uuid
	ID strfmt.UUID `json:"id"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id"`

	// type
	// Required: true
	// Enum: [account_amendment_submissions]
	Type string `json:"type"`
}

func AccountAmendmentSubmissionCreateWithDefaults(defaults client.Defaults) *AccountAmendmentSubmissionCreate {
	return &AccountAmendmentSubmissionCreate{

		ID: defaults.GetStrfmtUUID("AccountAmendmentSubmissionCreate", "id"),

		OrganisationID: defaults.GetStrfmtUUID("AccountAmendmentSubmissionCreate", "organisation_id"),

		Type: defaults.GetString("AccountAmendmentSubmissionCreate", "type"),
	}
}

func (m *AccountAmendmentSubmissionCreate) WithID(id strfmt.UUID) *AccountAmendmentSubmissionCreate {

	m.ID = id

	return m
}

func (m *AccountAmendmentSubmissionCreate) WithOrganisationID(organisationID strfmt.UUID) *AccountAmendmentSubmissionCreate {

	m.OrganisationID = organisationID

	return m
}

func (m *AccountAmendmentSubmissionCreate) WithType(typeVar string) *AccountAmendmentSubmissionCreate {

	m.Type = typeVar

	return m
}

// Validate validates this account amendment submission create
func (m *AccountAmendmentSubmissionCreate) Validate(formats strfmt.Registry) error {
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

func (m *AccountAmendmentSubmissionCreate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountAmendmentSubmissionCreate) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", strfmt.UUID(m.OrganisationID)); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

var accountAmendmentSubmissionCreateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["account_amendment_submissions"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountAmendmentSubmissionCreateTypeTypePropEnum = append(accountAmendmentSubmissionCreateTypeTypePropEnum, v)
	}
}

const (

	// AccountAmendmentSubmissionCreateTypeAccountAmendmentSubmissions captures enum value "account_amendment_submissions"
	AccountAmendmentSubmissionCreateTypeAccountAmendmentSubmissions string = "account_amendment_submissions"
)

// prop value enum
func (m *AccountAmendmentSubmissionCreate) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, accountAmendmentSubmissionCreateTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AccountAmendmentSubmissionCreate) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAmendmentSubmissionCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAmendmentSubmissionCreate) UnmarshalBinary(b []byte) error {
	var res AccountAmendmentSubmissionCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountAmendmentSubmissionCreate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}