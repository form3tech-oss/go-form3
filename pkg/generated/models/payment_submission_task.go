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

// PaymentSubmissionTask payment submission task
// swagger:model PaymentSubmissionTask
type PaymentSubmissionTask struct {

	// attributes
	Attributes *PaymentSubmissionTaskAttributes `json:"attributes,omitempty"`

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
	Relationships *PaymentSubmissionTaskRelationships `json:"relationships,omitempty"`

	// type
	// Enum: [payment_submission_tasks]
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func PaymentSubmissionTaskWithDefaults(defaults client.Defaults) *PaymentSubmissionTask {
	return &PaymentSubmissionTask{

		Attributes: PaymentSubmissionTaskAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTime("PaymentSubmissionTask", "created_on"),

		ID: defaults.GetStrfmtUUID("PaymentSubmissionTask", "id"),

		ModifiedOn: defaults.GetStrfmtDateTime("PaymentSubmissionTask", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("PaymentSubmissionTask", "organisation_id"),

		Relationships: PaymentSubmissionTaskRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("PaymentSubmissionTask", "type"),

		Version: defaults.GetInt64Ptr("PaymentSubmissionTask", "version"),
	}
}

func (m *PaymentSubmissionTask) WithAttributes(attributes PaymentSubmissionTaskAttributes) *PaymentSubmissionTask {

	m.Attributes = &attributes

	return m
}

func (m *PaymentSubmissionTask) WithoutAttributes() *PaymentSubmissionTask {
	m.Attributes = nil
	return m
}

func (m *PaymentSubmissionTask) WithCreatedOn(createdOn strfmt.DateTime) *PaymentSubmissionTask {

	m.CreatedOn = createdOn

	return m
}

func (m *PaymentSubmissionTask) WithID(id strfmt.UUID) *PaymentSubmissionTask {

	m.ID = id

	return m
}

func (m *PaymentSubmissionTask) WithModifiedOn(modifiedOn strfmt.DateTime) *PaymentSubmissionTask {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *PaymentSubmissionTask) WithOrganisationID(organisationID strfmt.UUID) *PaymentSubmissionTask {

	m.OrganisationID = organisationID

	return m
}

func (m *PaymentSubmissionTask) WithRelationships(relationships PaymentSubmissionTaskRelationships) *PaymentSubmissionTask {

	m.Relationships = &relationships

	return m
}

func (m *PaymentSubmissionTask) WithoutRelationships() *PaymentSubmissionTask {
	m.Relationships = nil
	return m
}

func (m *PaymentSubmissionTask) WithType(typeVar string) *PaymentSubmissionTask {

	m.Type = typeVar

	return m
}

func (m *PaymentSubmissionTask) WithVersion(version int64) *PaymentSubmissionTask {

	m.Version = &version

	return m
}

func (m *PaymentSubmissionTask) WithoutVersion() *PaymentSubmissionTask {
	m.Version = nil
	return m
}

// Validate validates this payment submission task
func (m *PaymentSubmissionTask) Validate(formats strfmt.Registry) error {
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

func (m *PaymentSubmissionTask) validateAttributes(formats strfmt.Registry) error {

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

func (m *PaymentSubmissionTask) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmissionTask) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmissionTask) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmissionTask) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmissionTask) validateRelationships(formats strfmt.Registry) error {

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

var paymentSubmissionTaskTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["payment_submission_tasks"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentSubmissionTaskTypeTypePropEnum = append(paymentSubmissionTaskTypeTypePropEnum, v)
	}
}

const (

	// PaymentSubmissionTaskTypePaymentSubmissionTasks captures enum value "payment_submission_tasks"
	PaymentSubmissionTaskTypePaymentSubmissionTasks string = "payment_submission_tasks"
)

// prop value enum
func (m *PaymentSubmissionTask) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentSubmissionTaskTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentSubmissionTask) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmissionTask) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentSubmissionTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentSubmissionTask) UnmarshalBinary(b []byte) error {
	var res PaymentSubmissionTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentSubmissionTask) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentSubmissionTaskAttributes payment submission task attributes
// swagger:model PaymentSubmissionTaskAttributes
type PaymentSubmissionTaskAttributes struct {

	// assignee
	Assignee PaymentSubmissionTaskAssignee `json:"assignee,omitempty"`

	// group
	Group PaymentSubmissionTaskGroup `json:"group,omitempty"`

	// Key Value map that contains additional data for executing the task.
	Input map[string]interface{} `json:"input,omitempty"`

	// name
	Name PaymentSubmissionTaskName `json:"name,omitempty"`

	// Key Value map that contains the Task result.
	Output map[string]interface{} `json:"output,omitempty"`

	// status
	Status PaymentSubmissionTaskStatus `json:"status,omitempty"`
}

func PaymentSubmissionTaskAttributesWithDefaults(defaults client.Defaults) *PaymentSubmissionTaskAttributes {
	return &PaymentSubmissionTaskAttributes{

		// TODO Assignee: PaymentSubmissionTaskAssignee,

		// TODO Group: PaymentSubmissionTaskGroup,

		Input: defaults.GetMapStringInterface("PaymentSubmissionTaskAttributes", "input"),

		// TODO Name: PaymentSubmissionTaskName,

		Output: defaults.GetMapStringInterface("PaymentSubmissionTaskAttributes", "output"),

		// TODO Status: PaymentSubmissionTaskStatus,

	}
}

func (m *PaymentSubmissionTaskAttributes) WithAssignee(assignee PaymentSubmissionTaskAssignee) *PaymentSubmissionTaskAttributes {

	m.Assignee = assignee

	return m
}

func (m *PaymentSubmissionTaskAttributes) WithGroup(group PaymentSubmissionTaskGroup) *PaymentSubmissionTaskAttributes {

	m.Group = group

	return m
}

func (m *PaymentSubmissionTaskAttributes) WithInput(input map[string]interface{}) *PaymentSubmissionTaskAttributes {

	m.Input = input

	return m
}

func (m *PaymentSubmissionTaskAttributes) WithName(name PaymentSubmissionTaskName) *PaymentSubmissionTaskAttributes {

	m.Name = name

	return m
}

func (m *PaymentSubmissionTaskAttributes) WithOutput(output map[string]interface{}) *PaymentSubmissionTaskAttributes {

	m.Output = output

	return m
}

func (m *PaymentSubmissionTaskAttributes) WithStatus(status PaymentSubmissionTaskStatus) *PaymentSubmissionTaskAttributes {

	m.Status = status

	return m
}

// Validate validates this payment submission task attributes
func (m *PaymentSubmissionTaskAttributes) Validate(formats strfmt.Registry) error {
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

func (m *PaymentSubmissionTaskAttributes) validateAssignee(formats strfmt.Registry) error {

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

func (m *PaymentSubmissionTaskAttributes) validateGroup(formats strfmt.Registry) error {

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

func (m *PaymentSubmissionTaskAttributes) validateName(formats strfmt.Registry) error {

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

func (m *PaymentSubmissionTaskAttributes) validateStatus(formats strfmt.Registry) error {

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
func (m *PaymentSubmissionTaskAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentSubmissionTaskAttributes) UnmarshalBinary(b []byte) error {
	var res PaymentSubmissionTaskAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentSubmissionTaskAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
