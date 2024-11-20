// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DirectDebitDecisionSubmission direct debit decision submission
// swagger:model DirectDebitDecisionSubmission
type DirectDebitDecisionSubmission struct {

	// attributes
	Attributes *DirectDebitDecisionSubmissionAttributes `json:"attributes,omitempty"`

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
	Relationships *DirectDebitDecisionSubmissionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func DirectDebitDecisionSubmissionWithDefaults(defaults client.Defaults) *DirectDebitDecisionSubmission {
	return &DirectDebitDecisionSubmission{

		Attributes: DirectDebitDecisionSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("DirectDebitDecisionSubmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("DirectDebitDecisionSubmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("DirectDebitDecisionSubmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("DirectDebitDecisionSubmission", "organisation_id"),

		Relationships: DirectDebitDecisionSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("DirectDebitDecisionSubmission", "type"),

		Version: defaults.GetInt64Ptr("DirectDebitDecisionSubmission", "version"),
	}
}

func (m *DirectDebitDecisionSubmission) WithAttributes(attributes DirectDebitDecisionSubmissionAttributes) *DirectDebitDecisionSubmission {

	m.Attributes = &attributes

	return m
}

func (m *DirectDebitDecisionSubmission) WithoutAttributes() *DirectDebitDecisionSubmission {
	m.Attributes = nil
	return m
}

func (m *DirectDebitDecisionSubmission) WithCreatedOn(createdOn strfmt.DateTime) *DirectDebitDecisionSubmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *DirectDebitDecisionSubmission) WithoutCreatedOn() *DirectDebitDecisionSubmission {
	m.CreatedOn = nil
	return m
}

func (m *DirectDebitDecisionSubmission) WithID(id strfmt.UUID) *DirectDebitDecisionSubmission {

	m.ID = &id

	return m
}

func (m *DirectDebitDecisionSubmission) WithoutID() *DirectDebitDecisionSubmission {
	m.ID = nil
	return m
}

func (m *DirectDebitDecisionSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *DirectDebitDecisionSubmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *DirectDebitDecisionSubmission) WithoutModifiedOn() *DirectDebitDecisionSubmission {
	m.ModifiedOn = nil
	return m
}

func (m *DirectDebitDecisionSubmission) WithOrganisationID(organisationID strfmt.UUID) *DirectDebitDecisionSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *DirectDebitDecisionSubmission) WithoutOrganisationID() *DirectDebitDecisionSubmission {
	m.OrganisationID = nil
	return m
}

func (m *DirectDebitDecisionSubmission) WithRelationships(relationships DirectDebitDecisionSubmissionRelationships) *DirectDebitDecisionSubmission {

	m.Relationships = &relationships

	return m
}

func (m *DirectDebitDecisionSubmission) WithoutRelationships() *DirectDebitDecisionSubmission {
	m.Relationships = nil
	return m
}

func (m *DirectDebitDecisionSubmission) WithType(typeVar string) *DirectDebitDecisionSubmission {

	m.Type = typeVar

	return m
}

func (m *DirectDebitDecisionSubmission) WithVersion(version int64) *DirectDebitDecisionSubmission {

	m.Version = &version

	return m
}

func (m *DirectDebitDecisionSubmission) WithoutVersion() *DirectDebitDecisionSubmission {
	m.Version = nil
	return m
}

// Validate validates this direct debit decision submission
func (m *DirectDebitDecisionSubmission) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitDecisionSubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *DirectDebitDecisionSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitDecisionSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitDecisionSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitDecisionSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitDecisionSubmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *DirectDebitDecisionSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitDecisionSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitDecisionSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitDecisionSubmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitDecisionSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitDecisionSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitDecisionSubmissionAttributes direct debit decision submission attributes
// swagger:model DirectDebitDecisionSubmissionAttributes
type DirectDebitDecisionSubmissionAttributes struct {

	// Indicates if the submission was created automatically by the system (`true`) or manually (`false`).
	Auto bool `json:"auto,omitempty"`

	// Scheme-specific status (if submission has been submitted to a scheme)
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// status
	Status DirectDebitDecisionSubmissionStatus `json:"status,omitempty"`

	// Reason for submission failure if status is `delivery_failed`
	StatusReason string `json:"status_reason,omitempty"`

	// Date and time of the submission
	// Read Only: true
	// Format: date-time
	SubmissionDatetime *strfmt.DateTime `json:"submission_datetime,omitempty"`
}

func DirectDebitDecisionSubmissionAttributesWithDefaults(defaults client.Defaults) *DirectDebitDecisionSubmissionAttributes {
	return &DirectDebitDecisionSubmissionAttributes{

		Auto: defaults.GetBool("DirectDebitDecisionSubmissionAttributes", "auto"),

		SchemeStatusCode: defaults.GetString("DirectDebitDecisionSubmissionAttributes", "scheme_status_code"),

		// TODO Status: DirectDebitDecisionSubmissionStatus,

		StatusReason: defaults.GetString("DirectDebitDecisionSubmissionAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTimePtr("DirectDebitDecisionSubmissionAttributes", "submission_datetime"),
	}
}

func (m *DirectDebitDecisionSubmissionAttributes) WithAuto(auto bool) *DirectDebitDecisionSubmissionAttributes {

	m.Auto = auto

	return m
}

func (m *DirectDebitDecisionSubmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *DirectDebitDecisionSubmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *DirectDebitDecisionSubmissionAttributes) WithStatus(status DirectDebitDecisionSubmissionStatus) *DirectDebitDecisionSubmissionAttributes {

	m.Status = status

	return m
}

func (m *DirectDebitDecisionSubmissionAttributes) WithStatusReason(statusReason string) *DirectDebitDecisionSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *DirectDebitDecisionSubmissionAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *DirectDebitDecisionSubmissionAttributes {

	m.SubmissionDatetime = &submissionDatetime

	return m
}

func (m *DirectDebitDecisionSubmissionAttributes) WithoutSubmissionDatetime() *DirectDebitDecisionSubmissionAttributes {
	m.SubmissionDatetime = nil
	return m
}

// Validate validates this direct debit decision submission attributes
func (m *DirectDebitDecisionSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitDecisionSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

func (m *DirectDebitDecisionSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitDecisionSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitDecisionSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res DirectDebitDecisionSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitDecisionSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitDecisionSubmissionRelationships direct debit decision submission relationships
// swagger:model DirectDebitDecisionSubmissionRelationships
type DirectDebitDecisionSubmissionRelationships struct {

	// direct debit
	DirectDebit *DirectDebitDecisionSubmissionRelationshipsDirectDebit `json:"direct_debit,omitempty"`

	// direct debit decision
	DirectDebitDecision *DirectDebitDecisionSubmissionRelationshipsDirectDebitDecision `json:"direct_debit_decision,omitempty"`
}

func DirectDebitDecisionSubmissionRelationshipsWithDefaults(defaults client.Defaults) *DirectDebitDecisionSubmissionRelationships {
	return &DirectDebitDecisionSubmissionRelationships{

		DirectDebit: DirectDebitDecisionSubmissionRelationshipsDirectDebitWithDefaults(defaults),

		DirectDebitDecision: DirectDebitDecisionSubmissionRelationshipsDirectDebitDecisionWithDefaults(defaults),
	}
}

func (m *DirectDebitDecisionSubmissionRelationships) WithDirectDebit(directDebit DirectDebitDecisionSubmissionRelationshipsDirectDebit) *DirectDebitDecisionSubmissionRelationships {

	m.DirectDebit = &directDebit

	return m
}

func (m *DirectDebitDecisionSubmissionRelationships) WithoutDirectDebit() *DirectDebitDecisionSubmissionRelationships {
	m.DirectDebit = nil
	return m
}

func (m *DirectDebitDecisionSubmissionRelationships) WithDirectDebitDecision(directDebitDecision DirectDebitDecisionSubmissionRelationshipsDirectDebitDecision) *DirectDebitDecisionSubmissionRelationships {

	m.DirectDebitDecision = &directDebitDecision

	return m
}

func (m *DirectDebitDecisionSubmissionRelationships) WithoutDirectDebitDecision() *DirectDebitDecisionSubmissionRelationships {
	m.DirectDebitDecision = nil
	return m
}

// Validate validates this direct debit decision submission relationships
func (m *DirectDebitDecisionSubmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectDebit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitDecision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitDecisionSubmissionRelationships) validateDirectDebit(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebit) { // not required
		return nil
	}

	if m.DirectDebit != nil {
		if err := m.DirectDebit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitDecisionSubmissionRelationships) validateDirectDebitDecision(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitDecision) { // not required
		return nil
	}

	if m.DirectDebitDecision != nil {
		if err := m.DirectDebitDecision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit_decision")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitDecisionSubmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitDecisionSubmissionRelationships) UnmarshalBinary(b []byte) error {
	var res DirectDebitDecisionSubmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitDecisionSubmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitDecisionSubmissionRelationshipsDirectDebit direct debit decision submission relationships direct debit
// swagger:model DirectDebitDecisionSubmissionRelationshipsDirectDebit
type DirectDebitDecisionSubmissionRelationshipsDirectDebit struct {

	// data
	Data []*DirectDebit `json:"data"`
}

func DirectDebitDecisionSubmissionRelationshipsDirectDebitWithDefaults(defaults client.Defaults) *DirectDebitDecisionSubmissionRelationshipsDirectDebit {
	return &DirectDebitDecisionSubmissionRelationshipsDirectDebit{

		Data: make([]*DirectDebit, 0),
	}
}

func (m *DirectDebitDecisionSubmissionRelationshipsDirectDebit) WithData(data []*DirectDebit) *DirectDebitDecisionSubmissionRelationshipsDirectDebit {

	m.Data = data

	return m
}

// Validate validates this direct debit decision submission relationships direct debit
func (m *DirectDebitDecisionSubmissionRelationshipsDirectDebit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitDecisionSubmissionRelationshipsDirectDebit) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relationships" + "." + "direct_debit" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitDecisionSubmissionRelationshipsDirectDebit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitDecisionSubmissionRelationshipsDirectDebit) UnmarshalBinary(b []byte) error {
	var res DirectDebitDecisionSubmissionRelationshipsDirectDebit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitDecisionSubmissionRelationshipsDirectDebit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitDecisionSubmissionRelationshipsDirectDebitDecision direct debit decision submission relationships direct debit decision
// swagger:model DirectDebitDecisionSubmissionRelationshipsDirectDebitDecision
type DirectDebitDecisionSubmissionRelationshipsDirectDebitDecision struct {

	// data
	Data []*DirectDebitDecision `json:"data"`
}

func DirectDebitDecisionSubmissionRelationshipsDirectDebitDecisionWithDefaults(defaults client.Defaults) *DirectDebitDecisionSubmissionRelationshipsDirectDebitDecision {
	return &DirectDebitDecisionSubmissionRelationshipsDirectDebitDecision{

		Data: make([]*DirectDebitDecision, 0),
	}
}

func (m *DirectDebitDecisionSubmissionRelationshipsDirectDebitDecision) WithData(data []*DirectDebitDecision) *DirectDebitDecisionSubmissionRelationshipsDirectDebitDecision {

	m.Data = data

	return m
}

// Validate validates this direct debit decision submission relationships direct debit decision
func (m *DirectDebitDecisionSubmissionRelationshipsDirectDebitDecision) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitDecisionSubmissionRelationshipsDirectDebitDecision) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relationships" + "." + "direct_debit_decision" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitDecisionSubmissionRelationshipsDirectDebitDecision) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitDecisionSubmissionRelationshipsDirectDebitDecision) UnmarshalBinary(b []byte) error {
	var res DirectDebitDecisionSubmissionRelationshipsDirectDebitDecision
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitDecisionSubmissionRelationshipsDirectDebitDecision) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
