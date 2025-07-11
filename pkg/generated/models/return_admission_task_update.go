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

// ReturnAdmissionTaskUpdate return admission task update
// swagger:model ReturnAdmissionTaskUpdate
type ReturnAdmissionTaskUpdate struct {

	// attributes
	Attributes *ReturnAdmissionTaskUpdateAttributes `json:"attributes,omitempty"`

	// Unique resource ID
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// type
	// Enum: ["return_admission_tasks"]
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ReturnAdmissionTaskUpdateWithDefaults(defaults client.Defaults) *ReturnAdmissionTaskUpdate {
	return &ReturnAdmissionTaskUpdate{

		Attributes: ReturnAdmissionTaskUpdateAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUID("ReturnAdmissionTaskUpdate", "id"),

		OrganisationID: defaults.GetStrfmtUUID("ReturnAdmissionTaskUpdate", "organisation_id"),

		Type: defaults.GetString("ReturnAdmissionTaskUpdate", "type"),

		Version: defaults.GetInt64Ptr("ReturnAdmissionTaskUpdate", "version"),
	}
}

func (m *ReturnAdmissionTaskUpdate) WithAttributes(attributes ReturnAdmissionTaskUpdateAttributes) *ReturnAdmissionTaskUpdate {

	m.Attributes = &attributes

	return m
}

func (m *ReturnAdmissionTaskUpdate) WithoutAttributes() *ReturnAdmissionTaskUpdate {
	m.Attributes = nil
	return m
}

func (m *ReturnAdmissionTaskUpdate) WithID(id strfmt.UUID) *ReturnAdmissionTaskUpdate {

	m.ID = id

	return m
}

func (m *ReturnAdmissionTaskUpdate) WithOrganisationID(organisationID strfmt.UUID) *ReturnAdmissionTaskUpdate {

	m.OrganisationID = organisationID

	return m
}

func (m *ReturnAdmissionTaskUpdate) WithType(typeVar string) *ReturnAdmissionTaskUpdate {

	m.Type = typeVar

	return m
}

func (m *ReturnAdmissionTaskUpdate) WithVersion(version int64) *ReturnAdmissionTaskUpdate {

	m.Version = &version

	return m
}

func (m *ReturnAdmissionTaskUpdate) WithoutVersion() *ReturnAdmissionTaskUpdate {
	m.Version = nil
	return m
}

// Validate validates this return admission task update
func (m *ReturnAdmissionTaskUpdate) Validate(formats strfmt.Registry) error {
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

func (m *ReturnAdmissionTaskUpdate) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReturnAdmissionTaskUpdate) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionTaskUpdate) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

var returnAdmissionTaskUpdateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["return_admission_tasks"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		returnAdmissionTaskUpdateTypeTypePropEnum = append(returnAdmissionTaskUpdateTypeTypePropEnum, v)
	}
}

const (

	// ReturnAdmissionTaskUpdateTypeReturnAdmissionTasks captures enum value "return_admission_tasks"
	ReturnAdmissionTaskUpdateTypeReturnAdmissionTasks string = "return_admission_tasks"
)

// prop value enum
func (m *ReturnAdmissionTaskUpdate) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, returnAdmissionTaskUpdateTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ReturnAdmissionTaskUpdate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionTaskUpdate) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnAdmissionTaskUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnAdmissionTaskUpdate) UnmarshalBinary(b []byte) error {
	var res ReturnAdmissionTaskUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnAdmissionTaskUpdate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReturnAdmissionTaskUpdateAttributes return admission task update attributes
// swagger:model ReturnAdmissionTaskUpdateAttributes
type ReturnAdmissionTaskUpdateAttributes struct {

	// name
	Name ReturnAdmissionTaskName `json:"name,omitempty"`

	// Key Value map that contains the Task result.
	Output map[string]interface{} `json:"output,omitempty"`

	// status
	Status ReturnAdmissionTaskStatus `json:"status,omitempty"`
}

func ReturnAdmissionTaskUpdateAttributesWithDefaults(defaults client.Defaults) *ReturnAdmissionTaskUpdateAttributes {
	return &ReturnAdmissionTaskUpdateAttributes{

		// TODO Name: ReturnAdmissionTaskName,

		Output: defaults.GetMapStringInterface("ReturnAdmissionTaskUpdateAttributes", "output"),

		// TODO Status: ReturnAdmissionTaskStatus,

	}
}

func (m *ReturnAdmissionTaskUpdateAttributes) WithName(name ReturnAdmissionTaskName) *ReturnAdmissionTaskUpdateAttributes {

	m.Name = name

	return m
}

func (m *ReturnAdmissionTaskUpdateAttributes) WithOutput(output map[string]interface{}) *ReturnAdmissionTaskUpdateAttributes {

	m.Output = output

	return m
}

func (m *ReturnAdmissionTaskUpdateAttributes) WithStatus(status ReturnAdmissionTaskStatus) *ReturnAdmissionTaskUpdateAttributes {

	m.Status = status

	return m
}

// Validate validates this return admission task update attributes
func (m *ReturnAdmissionTaskUpdateAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnAdmissionTaskUpdateAttributes) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := m.Name.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "name")
		}
		return err
	}

	return nil
}

func (m *ReturnAdmissionTaskUpdateAttributes) validateStatus(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ReturnAdmissionTaskUpdateAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnAdmissionTaskUpdateAttributes) UnmarshalBinary(b []byte) error {
	var res ReturnAdmissionTaskUpdateAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnAdmissionTaskUpdateAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
