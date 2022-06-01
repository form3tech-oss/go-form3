// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v4/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReversalPayment reversal payment
// swagger:model ReversalPayment
type ReversalPayment struct {

	// attributes
	Attributes interface{} `json:"attributes,omitempty"`

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

	// relationships
	Relationships *ReversalPaymentRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ReversalPaymentWithDefaults(defaults client.Defaults) *ReversalPayment {
	return &ReversalPayment{

		// TODO Attributes: interface{},

		CreatedOn: defaults.GetStrfmtDateTimePtr("ReversalPayment", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("ReversalPayment", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("ReversalPayment", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("ReversalPayment", "organisation_id"),

		Relationships: ReversalPaymentRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReversalPayment", "type"),

		Version: defaults.GetInt64Ptr("ReversalPayment", "version"),
	}
}

func (m *ReversalPayment) WithAttributes(attributes interface{}) *ReversalPayment {

	m.Attributes = attributes

	return m
}

func (m *ReversalPayment) WithCreatedOn(createdOn strfmt.DateTime) *ReversalPayment {

	m.CreatedOn = &createdOn

	return m
}

func (m *ReversalPayment) WithoutCreatedOn() *ReversalPayment {
	m.CreatedOn = nil
	return m
}

func (m *ReversalPayment) WithID(id strfmt.UUID) *ReversalPayment {

	m.ID = &id

	return m
}

func (m *ReversalPayment) WithoutID() *ReversalPayment {
	m.ID = nil
	return m
}

func (m *ReversalPayment) WithModifiedOn(modifiedOn strfmt.DateTime) *ReversalPayment {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *ReversalPayment) WithoutModifiedOn() *ReversalPayment {
	m.ModifiedOn = nil
	return m
}

func (m *ReversalPayment) WithOrganisationID(organisationID strfmt.UUID) *ReversalPayment {

	m.OrganisationID = &organisationID

	return m
}

func (m *ReversalPayment) WithoutOrganisationID() *ReversalPayment {
	m.OrganisationID = nil
	return m
}

func (m *ReversalPayment) WithRelationships(relationships ReversalPaymentRelationships) *ReversalPayment {

	m.Relationships = &relationships

	return m
}

func (m *ReversalPayment) WithoutRelationships() *ReversalPayment {
	m.Relationships = nil
	return m
}

func (m *ReversalPayment) WithType(typeVar string) *ReversalPayment {

	m.Type = typeVar

	return m
}

func (m *ReversalPayment) WithVersion(version int64) *ReversalPayment {

	m.Version = &version

	return m
}

func (m *ReversalPayment) WithoutVersion() *ReversalPayment {
	m.Version = nil
	return m
}

// Validate validates this reversal payment
func (m *ReversalPayment) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *ReversalPayment) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalPayment) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalPayment) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalPayment) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalPayment) validateRelationships(formats strfmt.Registry) error {

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

func (m *ReversalPayment) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ReversalPayment) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalPayment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalPayment) UnmarshalBinary(b []byte) error {
	var res ReversalPayment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalPayment) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReversalPaymentRelationships reversal payment relationships
// swagger:model ReversalPaymentRelationships
type ReversalPaymentRelationships struct {

	// ID of the payment resource related to the reversal
	Payment *RelationshipLinks `json:"payment,omitempty"`

	// ID of the reversal admission resource related to the reversal
	ReversalAdmission *RelationshipLinks `json:"reversal_admission,omitempty"`

	// ID of the reversal submission resource related to the reversal
	ReversalSubmission *RelationshipLinks `json:"reversal_submission,omitempty"`
}

func ReversalPaymentRelationshipsWithDefaults(defaults client.Defaults) *ReversalPaymentRelationships {
	return &ReversalPaymentRelationships{

		Payment: RelationshipLinksWithDefaults(defaults),

		ReversalAdmission: RelationshipLinksWithDefaults(defaults),

		ReversalSubmission: RelationshipLinksWithDefaults(defaults),
	}
}

func (m *ReversalPaymentRelationships) WithPayment(payment RelationshipLinks) *ReversalPaymentRelationships {

	m.Payment = &payment

	return m
}

func (m *ReversalPaymentRelationships) WithoutPayment() *ReversalPaymentRelationships {
	m.Payment = nil
	return m
}

func (m *ReversalPaymentRelationships) WithReversalAdmission(reversalAdmission RelationshipLinks) *ReversalPaymentRelationships {

	m.ReversalAdmission = &reversalAdmission

	return m
}

func (m *ReversalPaymentRelationships) WithoutReversalAdmission() *ReversalPaymentRelationships {
	m.ReversalAdmission = nil
	return m
}

func (m *ReversalPaymentRelationships) WithReversalSubmission(reversalSubmission RelationshipLinks) *ReversalPaymentRelationships {

	m.ReversalSubmission = &reversalSubmission

	return m
}

func (m *ReversalPaymentRelationships) WithoutReversalSubmission() *ReversalPaymentRelationships {
	m.ReversalSubmission = nil
	return m
}

// Validate validates this reversal payment relationships
func (m *ReversalPaymentRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReversalAdmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReversalSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReversalPaymentRelationships) validatePayment(formats strfmt.Registry) error {

	if swag.IsZero(m.Payment) { // not required
		return nil
	}

	if m.Payment != nil {
		if err := m.Payment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "payment")
			}
			return err
		}
	}

	return nil
}

func (m *ReversalPaymentRelationships) validateReversalAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.ReversalAdmission) { // not required
		return nil
	}

	if m.ReversalAdmission != nil {
		if err := m.ReversalAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "reversal_admission")
			}
			return err
		}
	}

	return nil
}

func (m *ReversalPaymentRelationships) validateReversalSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.ReversalSubmission) { // not required
		return nil
	}

	if m.ReversalSubmission != nil {
		if err := m.ReversalSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "reversal_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalPaymentRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalPaymentRelationships) UnmarshalBinary(b []byte) error {
	var res ReversalPaymentRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalPaymentRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
