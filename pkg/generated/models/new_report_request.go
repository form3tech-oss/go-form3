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

// NewReportRequest new report request
// swagger:model NewReportRequest
type NewReportRequest struct {

	// attributes
	// Required: true
	Attributes *NewReportRequestAttributes `json:"attributes"`

	// id
	// Required: true
	// Format: uuid
	ID strfmt.UUID `json:"id"`

	// organisation id
	// Required: true
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id"`

	// type
	// Enum: ["report_requests"]
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func NewReportRequestWithDefaults(defaults client.Defaults) *NewReportRequest {
	return &NewReportRequest{

		Attributes: NewReportRequestAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUID("NewReportRequest", "id"),

		OrganisationID: defaults.GetStrfmtUUID("NewReportRequest", "organisation_id"),

		Type: defaults.GetString("NewReportRequest", "type"),

		Version: defaults.GetInt64Ptr("NewReportRequest", "version"),
	}
}

func (m *NewReportRequest) WithAttributes(attributes NewReportRequestAttributes) *NewReportRequest {

	m.Attributes = &attributes

	return m
}

func (m *NewReportRequest) WithoutAttributes() *NewReportRequest {
	m.Attributes = nil
	return m
}

func (m *NewReportRequest) WithID(id strfmt.UUID) *NewReportRequest {

	m.ID = id

	return m
}

func (m *NewReportRequest) WithOrganisationID(organisationID strfmt.UUID) *NewReportRequest {

	m.OrganisationID = organisationID

	return m
}

func (m *NewReportRequest) WithType(typeVar string) *NewReportRequest {

	m.Type = typeVar

	return m
}

func (m *NewReportRequest) WithVersion(version int64) *NewReportRequest {

	m.Version = &version

	return m
}

func (m *NewReportRequest) WithoutVersion() *NewReportRequest {
	m.Version = nil
	return m
}

// Validate validates this new report request
func (m *NewReportRequest) Validate(formats strfmt.Registry) error {
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

func (m *NewReportRequest) validateAttributes(formats strfmt.Registry) error {

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

func (m *NewReportRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", strfmt.UUID(m.ID)); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *NewReportRequest) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", strfmt.UUID(m.OrganisationID)); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

var newReportRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["report_requests"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		newReportRequestTypeTypePropEnum = append(newReportRequestTypeTypePropEnum, v)
	}
}

const (

	// NewReportRequestTypeReportRequests captures enum value "report_requests"
	NewReportRequestTypeReportRequests string = "report_requests"
)

// prop value enum
func (m *NewReportRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, newReportRequestTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NewReportRequest) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *NewReportRequest) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewReportRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewReportRequest) UnmarshalBinary(b []byte) error {
	var res NewReportRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NewReportRequest) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
