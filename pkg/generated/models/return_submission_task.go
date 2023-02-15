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

// ReturnSubmissionTask return submission task
// swagger:model ReturnSubmissionTask
type ReturnSubmissionTask struct {

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
	Relationships *ReturnSubmissionTaskRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ReturnSubmissionTaskWithDefaults(defaults client.Defaults) *ReturnSubmissionTask {
	return &ReturnSubmissionTask{

		Attributes: ReturnSubmissionTaskAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTime("ReturnSubmissionTask", "created_on"),

		ID: defaults.GetStrfmtUUID("ReturnSubmissionTask", "id"),

		ModifiedOn: defaults.GetStrfmtDateTime("ReturnSubmissionTask", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("ReturnSubmissionTask", "organisation_id"),

		Relationships: ReturnSubmissionTaskRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReturnSubmissionTask", "type"),

		Version: defaults.GetInt64Ptr("ReturnSubmissionTask", "version"),
	}
}

func (m *ReturnSubmissionTask) WithAttributes(attributes ReturnSubmissionTaskAttributes) *ReturnSubmissionTask {

	m.Attributes = &attributes

	return m
}

func (m *ReturnSubmissionTask) WithoutAttributes() *ReturnSubmissionTask {
	m.Attributes = nil
	return m
}

func (m *ReturnSubmissionTask) WithCreatedOn(createdOn strfmt.DateTime) *ReturnSubmissionTask {

	m.CreatedOn = createdOn

	return m
}

func (m *ReturnSubmissionTask) WithID(id strfmt.UUID) *ReturnSubmissionTask {

	m.ID = id

	return m
}

func (m *ReturnSubmissionTask) WithModifiedOn(modifiedOn strfmt.DateTime) *ReturnSubmissionTask {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *ReturnSubmissionTask) WithOrganisationID(organisationID strfmt.UUID) *ReturnSubmissionTask {

	m.OrganisationID = organisationID

	return m
}

func (m *ReturnSubmissionTask) WithRelationships(relationships ReturnSubmissionTaskRelationships) *ReturnSubmissionTask {

	m.Relationships = &relationships

	return m
}

func (m *ReturnSubmissionTask) WithoutRelationships() *ReturnSubmissionTask {
	m.Relationships = nil
	return m
}

func (m *ReturnSubmissionTask) WithType(typeVar string) *ReturnSubmissionTask {

	m.Type = typeVar

	return m
}

func (m *ReturnSubmissionTask) WithVersion(version int64) *ReturnSubmissionTask {

	m.Version = &version

	return m
}

func (m *ReturnSubmissionTask) WithoutVersion() *ReturnSubmissionTask {
	m.Version = nil
	return m
}

// Validate validates this return submission task
func (m *ReturnSubmissionTask) Validate(formats strfmt.Registry) error {
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

func (m *ReturnSubmissionTask) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReturnSubmissionTask) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionTask) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionTask) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionTask) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionTask) validateRelationships(formats strfmt.Registry) error {

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

func (m *ReturnSubmissionTask) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionTask) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnSubmissionTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnSubmissionTask) UnmarshalBinary(b []byte) error {
	var res ReturnSubmissionTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnSubmissionTask) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReturnSubmissionTaskAttributes return submission task attributes
// swagger:model ReturnSubmissionTaskAttributes
type ReturnSubmissionTaskAttributes struct {

	// assignee
	Assignee ReturnSubmissionTaskAssignee `json:"assignee,omitempty"`

	// Key Value map that contains additional data for executing the task.
	Input map[string]interface{} `json:"input,omitempty"`

	// name
	Name ReturnSubmissionTaskName `json:"name,omitempty"`

	// Key Value map that contains the Task result.
	Output map[string]interface{} `json:"output,omitempty"`

	// status
	Status ReturnSubmissionTaskStatus `json:"status,omitempty"`
}

func ReturnSubmissionTaskAttributesWithDefaults(defaults client.Defaults) *ReturnSubmissionTaskAttributes {
	return &ReturnSubmissionTaskAttributes{

		// TODO Assignee: ReturnSubmissionTaskAssignee,

		Input: defaults.GetMapStringInterface("ReturnSubmissionTaskAttributes", "input"),

		// TODO Name: ReturnSubmissionTaskName,

		Output: defaults.GetMapStringInterface("ReturnSubmissionTaskAttributes", "output"),

		// TODO Status: ReturnSubmissionTaskStatus,

	}
}

func (m *ReturnSubmissionTaskAttributes) WithAssignee(assignee ReturnSubmissionTaskAssignee) *ReturnSubmissionTaskAttributes {

	m.Assignee = assignee

	return m
}

func (m *ReturnSubmissionTaskAttributes) WithInput(input map[string]interface{}) *ReturnSubmissionTaskAttributes {

	m.Input = input

	return m
}

func (m *ReturnSubmissionTaskAttributes) WithName(name ReturnSubmissionTaskName) *ReturnSubmissionTaskAttributes {

	m.Name = name

	return m
}

func (m *ReturnSubmissionTaskAttributes) WithOutput(output map[string]interface{}) *ReturnSubmissionTaskAttributes {

	m.Output = output

	return m
}

func (m *ReturnSubmissionTaskAttributes) WithStatus(status ReturnSubmissionTaskStatus) *ReturnSubmissionTaskAttributes {

	m.Status = status

	return m
}

// Validate validates this return submission task attributes
func (m *ReturnSubmissionTaskAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignee(formats); err != nil {
		res = append(res, err)
	}

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

func (m *ReturnSubmissionTaskAttributes) validateAssignee(formats strfmt.Registry) error {

	if swag.IsZero(m.Assignee) { // not required
		return nil
	}

	if err := m.Assignee.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "assignee")
		}
		return err
	}

	return nil
}

func (m *ReturnSubmissionTaskAttributes) validateName(formats strfmt.Registry) error {

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

func (m *ReturnSubmissionTaskAttributes) validateStatus(formats strfmt.Registry) error {

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
func (m *ReturnSubmissionTaskAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnSubmissionTaskAttributes) UnmarshalBinary(b []byte) error {
	var res ReturnSubmissionTaskAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnSubmissionTaskAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
