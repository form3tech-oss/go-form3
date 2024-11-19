// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DirectDebitRecall direct debit recall
// swagger:model DirectDebitRecall
type DirectDebitRecall struct {

	// attributes
	Attributes *DirectDebitRecallAttributes `json:"attributes,omitempty"`

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

	// organisation id
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// relationships
	Relationships *DirectDebitRecallRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func DirectDebitRecallWithDefaults(defaults client.Defaults) *DirectDebitRecall {
	return &DirectDebitRecall{

		Attributes: DirectDebitRecallAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("DirectDebitRecall", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("DirectDebitRecall", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("DirectDebitRecall", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("DirectDebitRecall", "organisation_id"),

		Relationships: DirectDebitRecallRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("DirectDebitRecall", "type"),

		Version: defaults.GetInt64Ptr("DirectDebitRecall", "version"),
	}
}

func (m *DirectDebitRecall) WithAttributes(attributes DirectDebitRecallAttributes) *DirectDebitRecall {

	m.Attributes = &attributes

	return m
}

func (m *DirectDebitRecall) WithoutAttributes() *DirectDebitRecall {
	m.Attributes = nil
	return m
}

func (m *DirectDebitRecall) WithCreatedOn(createdOn strfmt.DateTime) *DirectDebitRecall {

	m.CreatedOn = &createdOn

	return m
}

func (m *DirectDebitRecall) WithoutCreatedOn() *DirectDebitRecall {
	m.CreatedOn = nil
	return m
}

func (m *DirectDebitRecall) WithID(id strfmt.UUID) *DirectDebitRecall {

	m.ID = &id

	return m
}

func (m *DirectDebitRecall) WithoutID() *DirectDebitRecall {
	m.ID = nil
	return m
}

func (m *DirectDebitRecall) WithModifiedOn(modifiedOn strfmt.DateTime) *DirectDebitRecall {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *DirectDebitRecall) WithoutModifiedOn() *DirectDebitRecall {
	m.ModifiedOn = nil
	return m
}

func (m *DirectDebitRecall) WithOrganisationID(organisationID strfmt.UUID) *DirectDebitRecall {

	m.OrganisationID = organisationID

	return m
}

func (m *DirectDebitRecall) WithRelationships(relationships DirectDebitRecallRelationships) *DirectDebitRecall {

	m.Relationships = &relationships

	return m
}

func (m *DirectDebitRecall) WithoutRelationships() *DirectDebitRecall {
	m.Relationships = nil
	return m
}

func (m *DirectDebitRecall) WithType(typeVar string) *DirectDebitRecall {

	m.Type = typeVar

	return m
}

func (m *DirectDebitRecall) WithVersion(version int64) *DirectDebitRecall {

	m.Version = &version

	return m
}

func (m *DirectDebitRecall) WithoutVersion() *DirectDebitRecall {
	m.Version = nil
	return m
}

// Validate validates this direct debit recall
func (m *DirectDebitRecall) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitRecall) validateAttributes(formats strfmt.Registry) error {

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

func (m *DirectDebitRecall) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitRecall) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitRecall) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitRecall) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitRecall) validateRelationships(formats strfmt.Registry) error {

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

func (m *DirectDebitRecall) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitRecall) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitRecall) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRecall) UnmarshalBinary(b []byte) error {
	var res DirectDebitRecall
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRecall) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitRecallAttributes direct debit recall attributes
// swagger:model DirectDebitRecallAttributes
type DirectDebitRecallAttributes struct {

	// Further explanation of the reason given in reason_code. Only evaluated for certain reason codes.
	Reason string `json:"reason,omitempty"`

	// The reason for the recall. Has to be a valid [recall reason code](http://api-docs.form3.tech/api.html#enumerations-recall-reason-codes).
	ReasonCode string `json:"reason_code,omitempty"`

	// status
	Status RecallStatus `json:"status,omitempty"`
}

func DirectDebitRecallAttributesWithDefaults(defaults client.Defaults) *DirectDebitRecallAttributes {
	return &DirectDebitRecallAttributes{

		Reason: defaults.GetString("DirectDebitRecallAttributes", "reason"),

		ReasonCode: defaults.GetString("DirectDebitRecallAttributes", "reason_code"),

		// TODO Status: RecallStatus,

	}
}

func (m *DirectDebitRecallAttributes) WithReason(reason string) *DirectDebitRecallAttributes {

	m.Reason = reason

	return m
}

func (m *DirectDebitRecallAttributes) WithReasonCode(reasonCode string) *DirectDebitRecallAttributes {

	m.ReasonCode = reasonCode

	return m
}

func (m *DirectDebitRecallAttributes) WithStatus(status RecallStatus) *DirectDebitRecallAttributes {

	m.Status = status

	return m
}

// Validate validates this direct debit recall attributes
func (m *DirectDebitRecallAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitRecallAttributes) validateStatus(formats strfmt.Registry) error {

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
func (m *DirectDebitRecallAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRecallAttributes) UnmarshalBinary(b []byte) error {
	var res DirectDebitRecallAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRecallAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitRecallRelationships direct debit recall relationships
// swagger:model DirectDebitRecallRelationships
type DirectDebitRecallRelationships struct {

	// direct debit
	DirectDebit *DirectDebitRecallRelationshipsDirectDebit `json:"direct_debit,omitempty"`

	// direct debit recall admission
	DirectDebitRecallAdmission *DirectDebitRecallRelationshipsDirectDebitRecallAdmission `json:"direct_debit_recall_admission,omitempty"`

	// direct debit recall submission
	DirectDebitRecallSubmission *DirectDebitRecallRelationshipsDirectDebitRecallSubmission `json:"direct_debit_recall_submission,omitempty"`
}

func DirectDebitRecallRelationshipsWithDefaults(defaults client.Defaults) *DirectDebitRecallRelationships {
	return &DirectDebitRecallRelationships{

		DirectDebit: DirectDebitRecallRelationshipsDirectDebitWithDefaults(defaults),

		DirectDebitRecallAdmission: DirectDebitRecallRelationshipsDirectDebitRecallAdmissionWithDefaults(defaults),

		DirectDebitRecallSubmission: DirectDebitRecallRelationshipsDirectDebitRecallSubmissionWithDefaults(defaults),
	}
}

func (m *DirectDebitRecallRelationships) WithDirectDebit(directDebit DirectDebitRecallRelationshipsDirectDebit) *DirectDebitRecallRelationships {

	m.DirectDebit = &directDebit

	return m
}

func (m *DirectDebitRecallRelationships) WithoutDirectDebit() *DirectDebitRecallRelationships {
	m.DirectDebit = nil
	return m
}

func (m *DirectDebitRecallRelationships) WithDirectDebitRecallAdmission(directDebitRecallAdmission DirectDebitRecallRelationshipsDirectDebitRecallAdmission) *DirectDebitRecallRelationships {

	m.DirectDebitRecallAdmission = &directDebitRecallAdmission

	return m
}

func (m *DirectDebitRecallRelationships) WithoutDirectDebitRecallAdmission() *DirectDebitRecallRelationships {
	m.DirectDebitRecallAdmission = nil
	return m
}

func (m *DirectDebitRecallRelationships) WithDirectDebitRecallSubmission(directDebitRecallSubmission DirectDebitRecallRelationshipsDirectDebitRecallSubmission) *DirectDebitRecallRelationships {

	m.DirectDebitRecallSubmission = &directDebitRecallSubmission

	return m
}

func (m *DirectDebitRecallRelationships) WithoutDirectDebitRecallSubmission() *DirectDebitRecallRelationships {
	m.DirectDebitRecallSubmission = nil
	return m
}

// Validate validates this direct debit recall relationships
func (m *DirectDebitRecallRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectDebit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitRecallAdmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitRecallSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitRecallRelationships) validateDirectDebit(formats strfmt.Registry) error {

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

func (m *DirectDebitRecallRelationships) validateDirectDebitRecallAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitRecallAdmission) { // not required
		return nil
	}

	if m.DirectDebitRecallAdmission != nil {
		if err := m.DirectDebitRecallAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit_recall_admission")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitRecallRelationships) validateDirectDebitRecallSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitRecallSubmission) { // not required
		return nil
	}

	if m.DirectDebitRecallSubmission != nil {
		if err := m.DirectDebitRecallSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit_recall_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitRecallRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRecallRelationships) UnmarshalBinary(b []byte) error {
	var res DirectDebitRecallRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRecallRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitRecallRelationshipsDirectDebit direct debit recall relationships direct debit
// swagger:model DirectDebitRecallRelationshipsDirectDebit
type DirectDebitRecallRelationshipsDirectDebit struct {

	// data
	Data []*DirectDebit `json:"data"`
}

func DirectDebitRecallRelationshipsDirectDebitWithDefaults(defaults client.Defaults) *DirectDebitRecallRelationshipsDirectDebit {
	return &DirectDebitRecallRelationshipsDirectDebit{

		Data: make([]*DirectDebit, 0),
	}
}

func (m *DirectDebitRecallRelationshipsDirectDebit) WithData(data []*DirectDebit) *DirectDebitRecallRelationshipsDirectDebit {

	m.Data = data

	return m
}

// Validate validates this direct debit recall relationships direct debit
func (m *DirectDebitRecallRelationshipsDirectDebit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitRecallRelationshipsDirectDebit) validateData(formats strfmt.Registry) error {

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
func (m *DirectDebitRecallRelationshipsDirectDebit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRecallRelationshipsDirectDebit) UnmarshalBinary(b []byte) error {
	var res DirectDebitRecallRelationshipsDirectDebit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRecallRelationshipsDirectDebit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitRecallRelationshipsDirectDebitRecallAdmission direct debit recall relationships direct debit recall admission
// swagger:model DirectDebitRecallRelationshipsDirectDebitRecallAdmission
type DirectDebitRecallRelationshipsDirectDebitRecallAdmission struct {

	// data
	Data []*DirectDebitRecallAdmission `json:"data"`
}

func DirectDebitRecallRelationshipsDirectDebitRecallAdmissionWithDefaults(defaults client.Defaults) *DirectDebitRecallRelationshipsDirectDebitRecallAdmission {
	return &DirectDebitRecallRelationshipsDirectDebitRecallAdmission{

		Data: make([]*DirectDebitRecallAdmission, 0),
	}
}

func (m *DirectDebitRecallRelationshipsDirectDebitRecallAdmission) WithData(data []*DirectDebitRecallAdmission) *DirectDebitRecallRelationshipsDirectDebitRecallAdmission {

	m.Data = data

	return m
}

// Validate validates this direct debit recall relationships direct debit recall admission
func (m *DirectDebitRecallRelationshipsDirectDebitRecallAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitRecallRelationshipsDirectDebitRecallAdmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "direct_debit_recall_admission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitRecallRelationshipsDirectDebitRecallAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRecallRelationshipsDirectDebitRecallAdmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitRecallRelationshipsDirectDebitRecallAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRecallRelationshipsDirectDebitRecallAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitRecallRelationshipsDirectDebitRecallSubmission direct debit recall relationships direct debit recall submission
// swagger:model DirectDebitRecallRelationshipsDirectDebitRecallSubmission
type DirectDebitRecallRelationshipsDirectDebitRecallSubmission struct {

	// data
	Data []*DirectDebitRecallSubmission `json:"data"`
}

func DirectDebitRecallRelationshipsDirectDebitRecallSubmissionWithDefaults(defaults client.Defaults) *DirectDebitRecallRelationshipsDirectDebitRecallSubmission {
	return &DirectDebitRecallRelationshipsDirectDebitRecallSubmission{

		Data: make([]*DirectDebitRecallSubmission, 0),
	}
}

func (m *DirectDebitRecallRelationshipsDirectDebitRecallSubmission) WithData(data []*DirectDebitRecallSubmission) *DirectDebitRecallRelationshipsDirectDebitRecallSubmission {

	m.Data = data

	return m
}

// Validate validates this direct debit recall relationships direct debit recall submission
func (m *DirectDebitRecallRelationshipsDirectDebitRecallSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitRecallRelationshipsDirectDebitRecallSubmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "direct_debit_recall_submission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitRecallRelationshipsDirectDebitRecallSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRecallRelationshipsDirectDebitRecallSubmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitRecallRelationshipsDirectDebitRecallSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRecallRelationshipsDirectDebitRecallSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
