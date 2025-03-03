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

// ReversalAdmissionFetch reversal admission fetch
// swagger:model ReversalAdmissionFetch
type ReversalAdmissionFetch struct {

	// attributes
	Attributes *ReversalAdmissionFetchAttributes `json:"attributes,omitempty"`

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
	Relationships *ReversalAdmissionFetchRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ReversalAdmissionFetchWithDefaults(defaults client.Defaults) *ReversalAdmissionFetch {
	return &ReversalAdmissionFetch{

		Attributes: ReversalAdmissionFetchAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("ReversalAdmissionFetch", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("ReversalAdmissionFetch", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("ReversalAdmissionFetch", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("ReversalAdmissionFetch", "organisation_id"),

		Relationships: ReversalAdmissionFetchRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReversalAdmissionFetch", "type"),

		Version: defaults.GetInt64Ptr("ReversalAdmissionFetch", "version"),
	}
}

func (m *ReversalAdmissionFetch) WithAttributes(attributes ReversalAdmissionFetchAttributes) *ReversalAdmissionFetch {

	m.Attributes = &attributes

	return m
}

func (m *ReversalAdmissionFetch) WithoutAttributes() *ReversalAdmissionFetch {
	m.Attributes = nil
	return m
}

func (m *ReversalAdmissionFetch) WithCreatedOn(createdOn strfmt.DateTime) *ReversalAdmissionFetch {

	m.CreatedOn = &createdOn

	return m
}

func (m *ReversalAdmissionFetch) WithoutCreatedOn() *ReversalAdmissionFetch {
	m.CreatedOn = nil
	return m
}

func (m *ReversalAdmissionFetch) WithID(id strfmt.UUID) *ReversalAdmissionFetch {

	m.ID = &id

	return m
}

func (m *ReversalAdmissionFetch) WithoutID() *ReversalAdmissionFetch {
	m.ID = nil
	return m
}

func (m *ReversalAdmissionFetch) WithModifiedOn(modifiedOn strfmt.DateTime) *ReversalAdmissionFetch {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *ReversalAdmissionFetch) WithoutModifiedOn() *ReversalAdmissionFetch {
	m.ModifiedOn = nil
	return m
}

func (m *ReversalAdmissionFetch) WithOrganisationID(organisationID strfmt.UUID) *ReversalAdmissionFetch {

	m.OrganisationID = &organisationID

	return m
}

func (m *ReversalAdmissionFetch) WithoutOrganisationID() *ReversalAdmissionFetch {
	m.OrganisationID = nil
	return m
}

func (m *ReversalAdmissionFetch) WithRelationships(relationships ReversalAdmissionFetchRelationships) *ReversalAdmissionFetch {

	m.Relationships = &relationships

	return m
}

func (m *ReversalAdmissionFetch) WithoutRelationships() *ReversalAdmissionFetch {
	m.Relationships = nil
	return m
}

func (m *ReversalAdmissionFetch) WithType(typeVar string) *ReversalAdmissionFetch {

	m.Type = typeVar

	return m
}

func (m *ReversalAdmissionFetch) WithVersion(version int64) *ReversalAdmissionFetch {

	m.Version = &version

	return m
}

func (m *ReversalAdmissionFetch) WithoutVersion() *ReversalAdmissionFetch {
	m.Version = nil
	return m
}

// Validate validates this reversal admission fetch
func (m *ReversalAdmissionFetch) Validate(formats strfmt.Registry) error {
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

func (m *ReversalAdmissionFetch) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReversalAdmissionFetch) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmissionFetch) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmissionFetch) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmissionFetch) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmissionFetch) validateRelationships(formats strfmt.Registry) error {

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

func (m *ReversalAdmissionFetch) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", m.Type, `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ReversalAdmissionFetch) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalAdmissionFetch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalAdmissionFetch) UnmarshalBinary(b []byte) error {
	var res ReversalAdmissionFetch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalAdmissionFetch) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReversalAdmissionFetchAttributes reversal admission fetch attributes
// swagger:model ReversalAdmissionFetchAttributes
type ReversalAdmissionFetchAttributes struct {

	// Scheme-specific status code. Refer to scheme documentation where available.
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// Description of `scheme_status_code`
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// source gateway
	SourceGateway string `json:"source_gateway,omitempty"`

	// status
	Status ReversalAdmissionStatus `json:"status,omitempty"`
}

func ReversalAdmissionFetchAttributesWithDefaults(defaults client.Defaults) *ReversalAdmissionFetchAttributes {
	return &ReversalAdmissionFetchAttributes{

		SchemeStatusCode: defaults.GetString("ReversalAdmissionFetchAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("ReversalAdmissionFetchAttributes", "scheme_status_code_description"),

		SourceGateway: defaults.GetString("ReversalAdmissionFetchAttributes", "source_gateway"),

		// TODO Status: ReversalAdmissionStatus,

	}
}

func (m *ReversalAdmissionFetchAttributes) WithSchemeStatusCode(schemeStatusCode string) *ReversalAdmissionFetchAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *ReversalAdmissionFetchAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *ReversalAdmissionFetchAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *ReversalAdmissionFetchAttributes) WithSourceGateway(sourceGateway string) *ReversalAdmissionFetchAttributes {

	m.SourceGateway = sourceGateway

	return m
}

func (m *ReversalAdmissionFetchAttributes) WithStatus(status ReversalAdmissionStatus) *ReversalAdmissionFetchAttributes {

	m.Status = status

	return m
}

// Validate validates this reversal admission fetch attributes
func (m *ReversalAdmissionFetchAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReversalAdmissionFetchAttributes) validateStatus(formats strfmt.Registry) error {

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
func (m *ReversalAdmissionFetchAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalAdmissionFetchAttributes) UnmarshalBinary(b []byte) error {
	var res ReversalAdmissionFetchAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalAdmissionFetchAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReversalAdmissionFetchRelationships reversal admission fetch relationships
// swagger:model ReversalAdmissionFetchRelationships
type ReversalAdmissionFetchRelationships struct {

	// payment
	Payment *RelationshipLinks `json:"payment,omitempty"`

	// reversal
	Reversal *RelationshipLinks `json:"reversal,omitempty"`

	// reversal admission task
	ReversalAdmissionTask *RelationshipReversalAdmissionTasks `json:"reversal_admission_task,omitempty"`
}

func ReversalAdmissionFetchRelationshipsWithDefaults(defaults client.Defaults) *ReversalAdmissionFetchRelationships {
	return &ReversalAdmissionFetchRelationships{

		Payment: RelationshipLinksWithDefaults(defaults),

		Reversal: RelationshipLinksWithDefaults(defaults),

		ReversalAdmissionTask: RelationshipReversalAdmissionTasksWithDefaults(defaults),
	}
}

func (m *ReversalAdmissionFetchRelationships) WithPayment(payment RelationshipLinks) *ReversalAdmissionFetchRelationships {

	m.Payment = &payment

	return m
}

func (m *ReversalAdmissionFetchRelationships) WithoutPayment() *ReversalAdmissionFetchRelationships {
	m.Payment = nil
	return m
}

func (m *ReversalAdmissionFetchRelationships) WithReversal(reversal RelationshipLinks) *ReversalAdmissionFetchRelationships {

	m.Reversal = &reversal

	return m
}

func (m *ReversalAdmissionFetchRelationships) WithoutReversal() *ReversalAdmissionFetchRelationships {
	m.Reversal = nil
	return m
}

func (m *ReversalAdmissionFetchRelationships) WithReversalAdmissionTask(reversalAdmissionTask RelationshipReversalAdmissionTasks) *ReversalAdmissionFetchRelationships {

	m.ReversalAdmissionTask = &reversalAdmissionTask

	return m
}

func (m *ReversalAdmissionFetchRelationships) WithoutReversalAdmissionTask() *ReversalAdmissionFetchRelationships {
	m.ReversalAdmissionTask = nil
	return m
}

// Validate validates this reversal admission fetch relationships
func (m *ReversalAdmissionFetchRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReversal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReversalAdmissionTask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReversalAdmissionFetchRelationships) validatePayment(formats strfmt.Registry) error {

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

func (m *ReversalAdmissionFetchRelationships) validateReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.Reversal) { // not required
		return nil
	}

	if m.Reversal != nil {
		if err := m.Reversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "reversal")
			}
			return err
		}
	}

	return nil
}

func (m *ReversalAdmissionFetchRelationships) validateReversalAdmissionTask(formats strfmt.Registry) error {

	if swag.IsZero(m.ReversalAdmissionTask) { // not required
		return nil
	}

	if m.ReversalAdmissionTask != nil {
		if err := m.ReversalAdmissionTask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "reversal_admission_task")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalAdmissionFetchRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalAdmissionFetchRelationships) UnmarshalBinary(b []byte) error {
	var res ReversalAdmissionFetchRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalAdmissionFetchRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
