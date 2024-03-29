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

// PaymentSubmissionTaskUpdate payment submission task update
// swagger:model PaymentSubmissionTaskUpdate
type PaymentSubmissionTaskUpdate struct {

	// attributes
	Attributes *PaymentSubmissionTaskUpdateAttributes `json:"attributes,omitempty"`

	// Unique resource ID
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// type
	// Enum: [payment_submission_tasks]
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func PaymentSubmissionTaskUpdateWithDefaults(defaults client.Defaults) *PaymentSubmissionTaskUpdate {
	return &PaymentSubmissionTaskUpdate{

		Attributes: PaymentSubmissionTaskUpdateAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUID("PaymentSubmissionTaskUpdate", "id"),

		OrganisationID: defaults.GetStrfmtUUID("PaymentSubmissionTaskUpdate", "organisation_id"),

		Type: defaults.GetString("PaymentSubmissionTaskUpdate", "type"),

		Version: defaults.GetInt64Ptr("PaymentSubmissionTaskUpdate", "version"),
	}
}

func (m *PaymentSubmissionTaskUpdate) WithAttributes(attributes PaymentSubmissionTaskUpdateAttributes) *PaymentSubmissionTaskUpdate {

	m.Attributes = &attributes

	return m
}

func (m *PaymentSubmissionTaskUpdate) WithoutAttributes() *PaymentSubmissionTaskUpdate {
	m.Attributes = nil
	return m
}

func (m *PaymentSubmissionTaskUpdate) WithID(id strfmt.UUID) *PaymentSubmissionTaskUpdate {

	m.ID = id

	return m
}

func (m *PaymentSubmissionTaskUpdate) WithOrganisationID(organisationID strfmt.UUID) *PaymentSubmissionTaskUpdate {

	m.OrganisationID = organisationID

	return m
}

func (m *PaymentSubmissionTaskUpdate) WithType(typeVar string) *PaymentSubmissionTaskUpdate {

	m.Type = typeVar

	return m
}

func (m *PaymentSubmissionTaskUpdate) WithVersion(version int64) *PaymentSubmissionTaskUpdate {

	m.Version = &version

	return m
}

func (m *PaymentSubmissionTaskUpdate) WithoutVersion() *PaymentSubmissionTaskUpdate {
	m.Version = nil
	return m
}

// Validate validates this payment submission task update
func (m *PaymentSubmissionTaskUpdate) Validate(formats strfmt.Registry) error {
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

func (m *PaymentSubmissionTaskUpdate) validateAttributes(formats strfmt.Registry) error {

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

func (m *PaymentSubmissionTaskUpdate) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmissionTaskUpdate) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

var paymentSubmissionTaskUpdateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["payment_submission_tasks"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentSubmissionTaskUpdateTypeTypePropEnum = append(paymentSubmissionTaskUpdateTypeTypePropEnum, v)
	}
}

const (

	// PaymentSubmissionTaskUpdateTypePaymentSubmissionTasks captures enum value "payment_submission_tasks"
	PaymentSubmissionTaskUpdateTypePaymentSubmissionTasks string = "payment_submission_tasks"
)

// prop value enum
func (m *PaymentSubmissionTaskUpdate) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, paymentSubmissionTaskUpdateTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PaymentSubmissionTaskUpdate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *PaymentSubmissionTaskUpdate) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentSubmissionTaskUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentSubmissionTaskUpdate) UnmarshalBinary(b []byte) error {
	var res PaymentSubmissionTaskUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentSubmissionTaskUpdate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentSubmissionTaskUpdateAttributes payment submission task update attributes
// swagger:model PaymentSubmissionTaskUpdateAttributes
type PaymentSubmissionTaskUpdateAttributes struct {

	// Key Value map that contains the Task result.
	Output map[string]interface{} `json:"output,omitempty"`

	// status
	Status PaymentSubmissionTaskStatus `json:"status,omitempty"`
}

func PaymentSubmissionTaskUpdateAttributesWithDefaults(defaults client.Defaults) *PaymentSubmissionTaskUpdateAttributes {
	return &PaymentSubmissionTaskUpdateAttributes{

		Output: defaults.GetMapStringInterface("PaymentSubmissionTaskUpdateAttributes", "output"),

		// TODO Status: PaymentSubmissionTaskStatus,

	}
}

func (m *PaymentSubmissionTaskUpdateAttributes) WithOutput(output map[string]interface{}) *PaymentSubmissionTaskUpdateAttributes {

	m.Output = output

	return m
}

func (m *PaymentSubmissionTaskUpdateAttributes) WithStatus(status PaymentSubmissionTaskStatus) *PaymentSubmissionTaskUpdateAttributes {

	m.Status = status

	return m
}

// Validate validates this payment submission task update attributes
func (m *PaymentSubmissionTaskUpdateAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentSubmissionTaskUpdateAttributes) validateStatus(formats strfmt.Registry) error {

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
func (m *PaymentSubmissionTaskUpdateAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentSubmissionTaskUpdateAttributes) UnmarshalBinary(b []byte) error {
	var res PaymentSubmissionTaskUpdateAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentSubmissionTaskUpdateAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
