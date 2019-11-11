// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RecallReversal recall reversal
// swagger:model RecallReversal
type RecallReversal struct {

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
	Relationships *RecallReversalRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

// line 140

func RecallReversalWithDefaults(defaults client.Defaults) *RecallReversal {
	return &RecallReversal{

		// TODO Attributes: interface{},

		CreatedOn: defaults.GetStrfmtDateTimePtr("RecallReversal", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("RecallReversal", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("RecallReversal", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("RecallReversal", "organisation_id"),

		Relationships: RecallReversalRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("RecallReversal", "type"),

		Version: defaults.GetInt64Ptr("RecallReversal", "version"),
	}
}

func (m *RecallReversal) WithAttributes(attributes interface{}) *RecallReversal {

	m.Attributes = attributes

	return m
}

func (m *RecallReversal) WithCreatedOn(createdOn strfmt.DateTime) *RecallReversal {

	m.CreatedOn = &createdOn

	return m
}

func (m *RecallReversal) WithoutCreatedOn() *RecallReversal {
	m.CreatedOn = nil
	return m
}

func (m *RecallReversal) WithID(id strfmt.UUID) *RecallReversal {

	m.ID = &id

	return m
}

func (m *RecallReversal) WithoutID() *RecallReversal {
	m.ID = nil
	return m
}

func (m *RecallReversal) WithModifiedOn(modifiedOn strfmt.DateTime) *RecallReversal {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *RecallReversal) WithoutModifiedOn() *RecallReversal {
	m.ModifiedOn = nil
	return m
}

func (m *RecallReversal) WithOrganisationID(organisationID strfmt.UUID) *RecallReversal {

	m.OrganisationID = &organisationID

	return m
}

func (m *RecallReversal) WithoutOrganisationID() *RecallReversal {
	m.OrganisationID = nil
	return m
}

func (m *RecallReversal) WithRelationships(relationships RecallReversalRelationships) *RecallReversal {

	m.Relationships = &relationships

	return m
}

func (m *RecallReversal) WithoutRelationships() *RecallReversal {
	m.Relationships = nil
	return m
}

func (m *RecallReversal) WithType(typeVar string) *RecallReversal {

	m.Type = typeVar

	return m
}

func (m *RecallReversal) WithVersion(version int64) *RecallReversal {

	m.Version = &version

	return m
}

func (m *RecallReversal) WithoutVersion() *RecallReversal {
	m.Version = nil
	return m
}

// Validate validates this recall reversal
func (m *RecallReversal) Validate(formats strfmt.Registry) error {
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

func (m *RecallReversal) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallReversal) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallReversal) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallReversal) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallReversal) validateRelationships(formats strfmt.Registry) error {

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

func (m *RecallReversal) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *RecallReversal) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallReversal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallReversal) UnmarshalBinary(b []byte) error {
	var res RecallReversal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallReversal) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// RecallReversalRelationships recall reversal relationships
// swagger:model RecallReversalRelationships
type RecallReversalRelationships struct {

	// ID of the payment resource related to the recall reversal
	Payment *RelationshipLinks `json:"payment,omitempty"`

	// ID of the recall resource related to the recall reversal
	Recall *RelationshipLinks `json:"recall,omitempty"`

	// ID of the reversal admission resource related to the recall reversal
	ReversalAdmission *RelationshipLinks `json:"reversal_admission,omitempty"`
}

// line 140

func RecallReversalRelationshipsWithDefaults(defaults client.Defaults) *RecallReversalRelationships {
	return &RecallReversalRelationships{

		Payment: RelationshipLinksWithDefaults(defaults),

		Recall: RelationshipLinksWithDefaults(defaults),

		ReversalAdmission: RelationshipLinksWithDefaults(defaults),
	}
}

func (m *RecallReversalRelationships) WithPayment(payment RelationshipLinks) *RecallReversalRelationships {

	m.Payment = &payment

	return m
}

func (m *RecallReversalRelationships) WithoutPayment() *RecallReversalRelationships {
	m.Payment = nil
	return m
}

func (m *RecallReversalRelationships) WithRecall(recall RelationshipLinks) *RecallReversalRelationships {

	m.Recall = &recall

	return m
}

func (m *RecallReversalRelationships) WithoutRecall() *RecallReversalRelationships {
	m.Recall = nil
	return m
}

func (m *RecallReversalRelationships) WithReversalAdmission(reversalAdmission RelationshipLinks) *RecallReversalRelationships {

	m.ReversalAdmission = &reversalAdmission

	return m
}

func (m *RecallReversalRelationships) WithoutReversalAdmission() *RecallReversalRelationships {
	m.ReversalAdmission = nil
	return m
}

// Validate validates this recall reversal relationships
func (m *RecallReversalRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReversalAdmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecallReversalRelationships) validatePayment(formats strfmt.Registry) error {

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

func (m *RecallReversalRelationships) validateRecall(formats strfmt.Registry) error {

	if swag.IsZero(m.Recall) { // not required
		return nil
	}

	if m.Recall != nil {
		if err := m.Recall.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "recall")
			}
			return err
		}
	}

	return nil
}

func (m *RecallReversalRelationships) validateReversalAdmission(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *RecallReversalRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallReversalRelationships) UnmarshalBinary(b []byte) error {
	var res RecallReversalRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallReversalRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
