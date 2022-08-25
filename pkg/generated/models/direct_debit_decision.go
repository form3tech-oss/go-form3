// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v5/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DirectDebitDecision direct debit decision
// swagger:model DirectDebitDecision
type DirectDebitDecision struct {

	// attributes
	// Required: true
	Attributes *DirectDebitDecisionAttributes `json:"attributes"`

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
	Relationships *DirectDebitDecisionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func DirectDebitDecisionWithDefaults(defaults client.Defaults) *DirectDebitDecision {
	return &DirectDebitDecision{

		Attributes: DirectDebitDecisionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("DirectDebitDecision", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("DirectDebitDecision", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("DirectDebitDecision", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("DirectDebitDecision", "organisation_id"),

		Relationships: DirectDebitDecisionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("DirectDebitDecision", "type"),

		Version: defaults.GetInt64Ptr("DirectDebitDecision", "version"),
	}
}

func (m *DirectDebitDecision) WithAttributes(attributes DirectDebitDecisionAttributes) *DirectDebitDecision {

	m.Attributes = &attributes

	return m
}

func (m *DirectDebitDecision) WithoutAttributes() *DirectDebitDecision {
	m.Attributes = nil
	return m
}

func (m *DirectDebitDecision) WithCreatedOn(createdOn strfmt.DateTime) *DirectDebitDecision {

	m.CreatedOn = &createdOn

	return m
}

func (m *DirectDebitDecision) WithoutCreatedOn() *DirectDebitDecision {
	m.CreatedOn = nil
	return m
}

func (m *DirectDebitDecision) WithID(id strfmt.UUID) *DirectDebitDecision {

	m.ID = &id

	return m
}

func (m *DirectDebitDecision) WithoutID() *DirectDebitDecision {
	m.ID = nil
	return m
}

func (m *DirectDebitDecision) WithModifiedOn(modifiedOn strfmt.DateTime) *DirectDebitDecision {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *DirectDebitDecision) WithoutModifiedOn() *DirectDebitDecision {
	m.ModifiedOn = nil
	return m
}

func (m *DirectDebitDecision) WithOrganisationID(organisationID strfmt.UUID) *DirectDebitDecision {

	m.OrganisationID = &organisationID

	return m
}

func (m *DirectDebitDecision) WithoutOrganisationID() *DirectDebitDecision {
	m.OrganisationID = nil
	return m
}

func (m *DirectDebitDecision) WithRelationships(relationships DirectDebitDecisionRelationships) *DirectDebitDecision {

	m.Relationships = &relationships

	return m
}

func (m *DirectDebitDecision) WithoutRelationships() *DirectDebitDecision {
	m.Relationships = nil
	return m
}

func (m *DirectDebitDecision) WithType(typeVar string) *DirectDebitDecision {

	m.Type = typeVar

	return m
}

func (m *DirectDebitDecision) WithVersion(version int64) *DirectDebitDecision {

	m.Version = &version

	return m
}

func (m *DirectDebitDecision) WithoutVersion() *DirectDebitDecision {
	m.Version = nil
	return m
}

// Validate validates this direct debit decision
func (m *DirectDebitDecision) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitDecision) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
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

func (m *DirectDebitDecision) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitDecision) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitDecision) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitDecision) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitDecision) validateRelationships(formats strfmt.Registry) error {

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

func (m *DirectDebitDecision) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitDecision) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitDecision) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitDecision) UnmarshalBinary(b []byte) error {
	var res DirectDebitDecision
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitDecision) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitDecisionAttributes direct debit decision attributes
// swagger:model DirectDebitDecisionAttributes
type DirectDebitDecisionAttributes struct {

	// Answer to direct debit request. Only `rejected` can be used.
	// Enum: [rejected]
	Answer string `json:"answer,omitempty"`

	// Free text reason in addition to `reason_code`. Maximum length 105 characters for both inbound and outbound direct debit decisions.
	Reason string `json:"reason,omitempty"`

	// Reason for a rejected decision. Required when answer is rejected, ignored otherwise. Has to be a valid SEPA direct debit decision reason code.
	ReasonCode string `json:"reason_code,omitempty"`
}

func DirectDebitDecisionAttributesWithDefaults(defaults client.Defaults) *DirectDebitDecisionAttributes {
	return &DirectDebitDecisionAttributes{

		Answer: defaults.GetString("DirectDebitDecisionAttributes", "answer"),

		Reason: defaults.GetString("DirectDebitDecisionAttributes", "reason"),

		ReasonCode: defaults.GetString("DirectDebitDecisionAttributes", "reason_code"),
	}
}

func (m *DirectDebitDecisionAttributes) WithAnswer(answer string) *DirectDebitDecisionAttributes {

	m.Answer = answer

	return m
}

func (m *DirectDebitDecisionAttributes) WithReason(reason string) *DirectDebitDecisionAttributes {

	m.Reason = reason

	return m
}

func (m *DirectDebitDecisionAttributes) WithReasonCode(reasonCode string) *DirectDebitDecisionAttributes {

	m.ReasonCode = reasonCode

	return m
}

// Validate validates this direct debit decision attributes
func (m *DirectDebitDecisionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnswer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var directDebitDecisionAttributesTypeAnswerPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rejected"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		directDebitDecisionAttributesTypeAnswerPropEnum = append(directDebitDecisionAttributesTypeAnswerPropEnum, v)
	}
}

const (

	// DirectDebitDecisionAttributesAnswerRejected captures enum value "rejected"
	DirectDebitDecisionAttributesAnswerRejected string = "rejected"
)

// prop value enum
func (m *DirectDebitDecisionAttributes) validateAnswerEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, directDebitDecisionAttributesTypeAnswerPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DirectDebitDecisionAttributes) validateAnswer(formats strfmt.Registry) error {

	if swag.IsZero(m.Answer) { // not required
		return nil
	}

	// value enum
	if err := m.validateAnswerEnum("attributes"+"."+"answer", "body", m.Answer); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitDecisionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitDecisionAttributes) UnmarshalBinary(b []byte) error {
	var res DirectDebitDecisionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitDecisionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitDecisionRelationships direct debit decision relationships
// swagger:model DirectDebitDecisionRelationships
type DirectDebitDecisionRelationships struct {

	// direct debit
	DirectDebit *DirectDebitDecisionRelationshipsDirectDebit `json:"direct_debit,omitempty"`

	// direct debit decision admission
	DirectDebitDecisionAdmission *DirectDebitDecisionRelationshipsDirectDebitDecisionAdmission `json:"direct_debit_decision_admission,omitempty"`

	// direct debit decision submission
	DirectDebitDecisionSubmission *DirectDebitDecisionRelationshipsDirectDebitDecisionSubmission `json:"direct_debit_decision_submission,omitempty"`
}

func DirectDebitDecisionRelationshipsWithDefaults(defaults client.Defaults) *DirectDebitDecisionRelationships {
	return &DirectDebitDecisionRelationships{

		DirectDebit: DirectDebitDecisionRelationshipsDirectDebitWithDefaults(defaults),

		DirectDebitDecisionAdmission: DirectDebitDecisionRelationshipsDirectDebitDecisionAdmissionWithDefaults(defaults),

		DirectDebitDecisionSubmission: DirectDebitDecisionRelationshipsDirectDebitDecisionSubmissionWithDefaults(defaults),
	}
}

func (m *DirectDebitDecisionRelationships) WithDirectDebit(directDebit DirectDebitDecisionRelationshipsDirectDebit) *DirectDebitDecisionRelationships {

	m.DirectDebit = &directDebit

	return m
}

func (m *DirectDebitDecisionRelationships) WithoutDirectDebit() *DirectDebitDecisionRelationships {
	m.DirectDebit = nil
	return m
}

func (m *DirectDebitDecisionRelationships) WithDirectDebitDecisionAdmission(directDebitDecisionAdmission DirectDebitDecisionRelationshipsDirectDebitDecisionAdmission) *DirectDebitDecisionRelationships {

	m.DirectDebitDecisionAdmission = &directDebitDecisionAdmission

	return m
}

func (m *DirectDebitDecisionRelationships) WithoutDirectDebitDecisionAdmission() *DirectDebitDecisionRelationships {
	m.DirectDebitDecisionAdmission = nil
	return m
}

func (m *DirectDebitDecisionRelationships) WithDirectDebitDecisionSubmission(directDebitDecisionSubmission DirectDebitDecisionRelationshipsDirectDebitDecisionSubmission) *DirectDebitDecisionRelationships {

	m.DirectDebitDecisionSubmission = &directDebitDecisionSubmission

	return m
}

func (m *DirectDebitDecisionRelationships) WithoutDirectDebitDecisionSubmission() *DirectDebitDecisionRelationships {
	m.DirectDebitDecisionSubmission = nil
	return m
}

// Validate validates this direct debit decision relationships
func (m *DirectDebitDecisionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectDebit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitDecisionAdmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitDecisionSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitDecisionRelationships) validateDirectDebit(formats strfmt.Registry) error {

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

func (m *DirectDebitDecisionRelationships) validateDirectDebitDecisionAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitDecisionAdmission) { // not required
		return nil
	}

	if m.DirectDebitDecisionAdmission != nil {
		if err := m.DirectDebitDecisionAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit_decision_admission")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitDecisionRelationships) validateDirectDebitDecisionSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitDecisionSubmission) { // not required
		return nil
	}

	if m.DirectDebitDecisionSubmission != nil {
		if err := m.DirectDebitDecisionSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit_decision_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitDecisionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitDecisionRelationships) UnmarshalBinary(b []byte) error {
	var res DirectDebitDecisionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitDecisionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitDecisionRelationshipsDirectDebit direct debit decision relationships direct debit
// swagger:model DirectDebitDecisionRelationshipsDirectDebit
type DirectDebitDecisionRelationshipsDirectDebit struct {

	// data
	Data []*DirectDebit `json:"data"`
}

func DirectDebitDecisionRelationshipsDirectDebitWithDefaults(defaults client.Defaults) *DirectDebitDecisionRelationshipsDirectDebit {
	return &DirectDebitDecisionRelationshipsDirectDebit{

		Data: make([]*DirectDebit, 0),
	}
}

func (m *DirectDebitDecisionRelationshipsDirectDebit) WithData(data []*DirectDebit) *DirectDebitDecisionRelationshipsDirectDebit {

	m.Data = data

	return m
}

// Validate validates this direct debit decision relationships direct debit
func (m *DirectDebitDecisionRelationshipsDirectDebit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitDecisionRelationshipsDirectDebit) validateData(formats strfmt.Registry) error {

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
func (m *DirectDebitDecisionRelationshipsDirectDebit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitDecisionRelationshipsDirectDebit) UnmarshalBinary(b []byte) error {
	var res DirectDebitDecisionRelationshipsDirectDebit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitDecisionRelationshipsDirectDebit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitDecisionRelationshipsDirectDebitDecisionAdmission direct debit decision relationships direct debit decision admission
// swagger:model DirectDebitDecisionRelationshipsDirectDebitDecisionAdmission
type DirectDebitDecisionRelationshipsDirectDebitDecisionAdmission struct {

	// data
	Data []*DirectDebitDecisionAdmission `json:"data"`
}

func DirectDebitDecisionRelationshipsDirectDebitDecisionAdmissionWithDefaults(defaults client.Defaults) *DirectDebitDecisionRelationshipsDirectDebitDecisionAdmission {
	return &DirectDebitDecisionRelationshipsDirectDebitDecisionAdmission{

		Data: make([]*DirectDebitDecisionAdmission, 0),
	}
}

func (m *DirectDebitDecisionRelationshipsDirectDebitDecisionAdmission) WithData(data []*DirectDebitDecisionAdmission) *DirectDebitDecisionRelationshipsDirectDebitDecisionAdmission {

	m.Data = data

	return m
}

// Validate validates this direct debit decision relationships direct debit decision admission
func (m *DirectDebitDecisionRelationshipsDirectDebitDecisionAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitDecisionRelationshipsDirectDebitDecisionAdmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "direct_debit_decision_admission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitDecisionRelationshipsDirectDebitDecisionAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitDecisionRelationshipsDirectDebitDecisionAdmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitDecisionRelationshipsDirectDebitDecisionAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitDecisionRelationshipsDirectDebitDecisionAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitDecisionRelationshipsDirectDebitDecisionSubmission direct debit decision relationships direct debit decision submission
// swagger:model DirectDebitDecisionRelationshipsDirectDebitDecisionSubmission
type DirectDebitDecisionRelationshipsDirectDebitDecisionSubmission struct {

	// data
	Data []*DirectDebitDecisionSubmission `json:"data"`
}

func DirectDebitDecisionRelationshipsDirectDebitDecisionSubmissionWithDefaults(defaults client.Defaults) *DirectDebitDecisionRelationshipsDirectDebitDecisionSubmission {
	return &DirectDebitDecisionRelationshipsDirectDebitDecisionSubmission{

		Data: make([]*DirectDebitDecisionSubmission, 0),
	}
}

func (m *DirectDebitDecisionRelationshipsDirectDebitDecisionSubmission) WithData(data []*DirectDebitDecisionSubmission) *DirectDebitDecisionRelationshipsDirectDebitDecisionSubmission {

	m.Data = data

	return m
}

// Validate validates this direct debit decision relationships direct debit decision submission
func (m *DirectDebitDecisionRelationshipsDirectDebitDecisionSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitDecisionRelationshipsDirectDebitDecisionSubmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "direct_debit_decision_submission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitDecisionRelationshipsDirectDebitDecisionSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitDecisionRelationshipsDirectDebitDecisionSubmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitDecisionRelationshipsDirectDebitDecisionSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitDecisionRelationshipsDirectDebitDecisionSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
