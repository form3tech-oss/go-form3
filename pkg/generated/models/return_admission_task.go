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

// ReturnAdmissionTask return admission task
// swagger:model ReturnAdmissionTask
type ReturnAdmissionTask struct {

	// attributes
	Attributes *ReturnAdmissionTaskAttributes `json:"attributes,omitempty"`

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
	Relationships *ReturnAdmissionTaskRelationships `json:"relationships,omitempty"`

	// type
	// Enum: [return_admission_tasks]
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ReturnAdmissionTaskWithDefaults(defaults client.Defaults) *ReturnAdmissionTask {
	return &ReturnAdmissionTask{

		Attributes: ReturnAdmissionTaskAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTime("ReturnAdmissionTask", "created_on"),

		ID: defaults.GetStrfmtUUID("ReturnAdmissionTask", "id"),

		ModifiedOn: defaults.GetStrfmtDateTime("ReturnAdmissionTask", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("ReturnAdmissionTask", "organisation_id"),

		Relationships: ReturnAdmissionTaskRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReturnAdmissionTask", "type"),

		Version: defaults.GetInt64Ptr("ReturnAdmissionTask", "version"),
	}
}

func (m *ReturnAdmissionTask) WithAttributes(attributes ReturnAdmissionTaskAttributes) *ReturnAdmissionTask {

	m.Attributes = &attributes

	return m
}

func (m *ReturnAdmissionTask) WithoutAttributes() *ReturnAdmissionTask {
	m.Attributes = nil
	return m
}

func (m *ReturnAdmissionTask) WithCreatedOn(createdOn strfmt.DateTime) *ReturnAdmissionTask {

	m.CreatedOn = createdOn

	return m
}

func (m *ReturnAdmissionTask) WithID(id strfmt.UUID) *ReturnAdmissionTask {

	m.ID = id

	return m
}

func (m *ReturnAdmissionTask) WithModifiedOn(modifiedOn strfmt.DateTime) *ReturnAdmissionTask {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *ReturnAdmissionTask) WithOrganisationID(organisationID strfmt.UUID) *ReturnAdmissionTask {

	m.OrganisationID = organisationID

	return m
}

func (m *ReturnAdmissionTask) WithRelationships(relationships ReturnAdmissionTaskRelationships) *ReturnAdmissionTask {

	m.Relationships = &relationships

	return m
}

func (m *ReturnAdmissionTask) WithoutRelationships() *ReturnAdmissionTask {
	m.Relationships = nil
	return m
}

func (m *ReturnAdmissionTask) WithType(typeVar string) *ReturnAdmissionTask {

	m.Type = typeVar

	return m
}

func (m *ReturnAdmissionTask) WithVersion(version int64) *ReturnAdmissionTask {

	m.Version = &version

	return m
}

func (m *ReturnAdmissionTask) WithoutVersion() *ReturnAdmissionTask {
	m.Version = nil
	return m
}

// Validate validates this return admission task
func (m *ReturnAdmissionTask) Validate(formats strfmt.Registry) error {
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

func (m *ReturnAdmissionTask) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReturnAdmissionTask) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionTask) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionTask) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionTask) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionTask) validateRelationships(formats strfmt.Registry) error {

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

var returnAdmissionTaskTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["return_admission_tasks"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		returnAdmissionTaskTypeTypePropEnum = append(returnAdmissionTaskTypeTypePropEnum, v)
	}
}

const (

	// ReturnAdmissionTaskTypeReturnAdmissionTasks captures enum value "return_admission_tasks"
	ReturnAdmissionTaskTypeReturnAdmissionTasks string = "return_admission_tasks"
)

// prop value enum
func (m *ReturnAdmissionTask) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, returnAdmissionTaskTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ReturnAdmissionTask) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionTask) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnAdmissionTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnAdmissionTask) UnmarshalBinary(b []byte) error {
	var res ReturnAdmissionTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnAdmissionTask) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReturnAdmissionTaskAttributes return admission task attributes
// swagger:model ReturnAdmissionTaskAttributes
type ReturnAdmissionTaskAttributes struct {

	// assignee
	Assignee ReturnAdmissionTaskAssignee `json:"assignee,omitempty"`

	// group
	Group ReturnAdmissionTaskGroup `json:"group,omitempty"`

	// Key Value map that contains additional data for executing the task.
	Input map[string]interface{} `json:"input,omitempty"`

	// name
	Name ReturnAdmissionTaskName `json:"name,omitempty"`

	// Key Value map that contains the Task result.
	Output map[string]interface{} `json:"output,omitempty"`

	// status
	Status ReturnAdmissionTaskStatus `json:"status,omitempty"`
}

func ReturnAdmissionTaskAttributesWithDefaults(defaults client.Defaults) *ReturnAdmissionTaskAttributes {
	return &ReturnAdmissionTaskAttributes{

		// TODO Assignee: ReturnAdmissionTaskAssignee,

		// TODO Group: ReturnAdmissionTaskGroup,

		Input: defaults.GetMapStringInterface("ReturnAdmissionTaskAttributes", "input"),

		// TODO Name: ReturnAdmissionTaskName,

		Output: defaults.GetMapStringInterface("ReturnAdmissionTaskAttributes", "output"),

		// TODO Status: ReturnAdmissionTaskStatus,

	}
}

func (m *ReturnAdmissionTaskAttributes) WithAssignee(assignee ReturnAdmissionTaskAssignee) *ReturnAdmissionTaskAttributes {

	m.Assignee = assignee

	return m
}

func (m *ReturnAdmissionTaskAttributes) WithGroup(group ReturnAdmissionTaskGroup) *ReturnAdmissionTaskAttributes {

	m.Group = group

	return m
}

func (m *ReturnAdmissionTaskAttributes) WithInput(input map[string]interface{}) *ReturnAdmissionTaskAttributes {

	m.Input = input

	return m
}

func (m *ReturnAdmissionTaskAttributes) WithName(name ReturnAdmissionTaskName) *ReturnAdmissionTaskAttributes {

	m.Name = name

	return m
}

func (m *ReturnAdmissionTaskAttributes) WithOutput(output map[string]interface{}) *ReturnAdmissionTaskAttributes {

	m.Output = output

	return m
}

func (m *ReturnAdmissionTaskAttributes) WithStatus(status ReturnAdmissionTaskStatus) *ReturnAdmissionTaskAttributes {

	m.Status = status

	return m
}

// Validate validates this return admission task attributes
func (m *ReturnAdmissionTaskAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignee(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
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

func (m *ReturnAdmissionTaskAttributes) validateAssignee(formats strfmt.Registry) error {

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

func (m *ReturnAdmissionTaskAttributes) validateGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if err := m.Group.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "group")
		}
		return err
	}

	return nil
}

func (m *ReturnAdmissionTaskAttributes) validateName(formats strfmt.Registry) error {

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

func (m *ReturnAdmissionTaskAttributes) validateStatus(formats strfmt.Registry) error {

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
func (m *ReturnAdmissionTaskAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnAdmissionTaskAttributes) UnmarshalBinary(b []byte) error {
	var res ReturnAdmissionTaskAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnAdmissionTaskAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
