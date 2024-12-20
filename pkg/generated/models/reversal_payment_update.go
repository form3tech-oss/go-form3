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

// ReversalPaymentUpdate reversal payment update
// swagger:model ReversalPaymentUpdate
type ReversalPaymentUpdate struct {

	// attributes
	Attributes *ReversalPaymentUpdateAttributes `json:"attributes,omitempty"`

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
	Relationships *ReversalPaymentUpdateRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Required: true
	// Minimum: 0
	Version *int64 `json:"version"`
}

func ReversalPaymentUpdateWithDefaults(defaults client.Defaults) *ReversalPaymentUpdate {
	return &ReversalPaymentUpdate{

		Attributes: ReversalPaymentUpdateAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("ReversalPaymentUpdate", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("ReversalPaymentUpdate", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("ReversalPaymentUpdate", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("ReversalPaymentUpdate", "organisation_id"),

		Relationships: ReversalPaymentUpdateRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReversalPaymentUpdate", "type"),

		Version: defaults.GetInt64Ptr("ReversalPaymentUpdate", "version"),
	}
}

func (m *ReversalPaymentUpdate) WithAttributes(attributes ReversalPaymentUpdateAttributes) *ReversalPaymentUpdate {

	m.Attributes = &attributes

	return m
}

func (m *ReversalPaymentUpdate) WithoutAttributes() *ReversalPaymentUpdate {
	m.Attributes = nil
	return m
}

func (m *ReversalPaymentUpdate) WithCreatedOn(createdOn strfmt.DateTime) *ReversalPaymentUpdate {

	m.CreatedOn = &createdOn

	return m
}

func (m *ReversalPaymentUpdate) WithoutCreatedOn() *ReversalPaymentUpdate {
	m.CreatedOn = nil
	return m
}

func (m *ReversalPaymentUpdate) WithID(id strfmt.UUID) *ReversalPaymentUpdate {

	m.ID = &id

	return m
}

func (m *ReversalPaymentUpdate) WithoutID() *ReversalPaymentUpdate {
	m.ID = nil
	return m
}

func (m *ReversalPaymentUpdate) WithModifiedOn(modifiedOn strfmt.DateTime) *ReversalPaymentUpdate {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *ReversalPaymentUpdate) WithoutModifiedOn() *ReversalPaymentUpdate {
	m.ModifiedOn = nil
	return m
}

func (m *ReversalPaymentUpdate) WithOrganisationID(organisationID strfmt.UUID) *ReversalPaymentUpdate {

	m.OrganisationID = &organisationID

	return m
}

func (m *ReversalPaymentUpdate) WithoutOrganisationID() *ReversalPaymentUpdate {
	m.OrganisationID = nil
	return m
}

func (m *ReversalPaymentUpdate) WithRelationships(relationships ReversalPaymentUpdateRelationships) *ReversalPaymentUpdate {

	m.Relationships = &relationships

	return m
}

func (m *ReversalPaymentUpdate) WithoutRelationships() *ReversalPaymentUpdate {
	m.Relationships = nil
	return m
}

func (m *ReversalPaymentUpdate) WithType(typeVar string) *ReversalPaymentUpdate {

	m.Type = typeVar

	return m
}

func (m *ReversalPaymentUpdate) WithVersion(version int64) *ReversalPaymentUpdate {

	m.Version = &version

	return m
}

func (m *ReversalPaymentUpdate) WithoutVersion() *ReversalPaymentUpdate {
	m.Version = nil
	return m
}

// Validate validates this reversal payment update
func (m *ReversalPaymentUpdate) Validate(formats strfmt.Registry) error {
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

func (m *ReversalPaymentUpdate) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReversalPaymentUpdate) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalPaymentUpdate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalPaymentUpdate) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalPaymentUpdate) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReversalPaymentUpdate) validateRelationships(formats strfmt.Registry) error {

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

func (m *ReversalPaymentUpdate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", m.Type, `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ReversalPaymentUpdate) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalPaymentUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalPaymentUpdate) UnmarshalBinary(b []byte) error {
	var res ReversalPaymentUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalPaymentUpdate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReversalPaymentUpdateAttributes reversal payment update attributes
// swagger:model ReversalPaymentUpdateAttributes
type ReversalPaymentUpdateAttributes struct {

	// All purpose list of key-value pairs.
	// Max Items: 5
	UserDefinedData []*UserDefinedData `json:"user_defined_data,omitempty"`
}

func ReversalPaymentUpdateAttributesWithDefaults(defaults client.Defaults) *ReversalPaymentUpdateAttributes {
	return &ReversalPaymentUpdateAttributes{

		UserDefinedData: make([]*UserDefinedData, 0),
	}
}

func (m *ReversalPaymentUpdateAttributes) WithUserDefinedData(userDefinedData []*UserDefinedData) *ReversalPaymentUpdateAttributes {

	m.UserDefinedData = userDefinedData

	return m
}

// Validate validates this reversal payment update attributes
func (m *ReversalPaymentUpdateAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserDefinedData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReversalPaymentUpdateAttributes) validateUserDefinedData(formats strfmt.Registry) error {

	if swag.IsZero(m.UserDefinedData) { // not required
		return nil
	}

	iUserDefinedDataSize := int64(len(m.UserDefinedData))

	if err := validate.MaxItems("attributes"+"."+"user_defined_data", "body", iUserDefinedDataSize, 5); err != nil {
		return err
	}

	for i := 0; i < len(m.UserDefinedData); i++ {
		if swag.IsZero(m.UserDefinedData[i]) { // not required
			continue
		}

		if m.UserDefinedData[i] != nil {
			if err := m.UserDefinedData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + "user_defined_data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalPaymentUpdateAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalPaymentUpdateAttributes) UnmarshalBinary(b []byte) error {
	var res ReversalPaymentUpdateAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalPaymentUpdateAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReversalPaymentUpdateRelationships reversal payment update relationships
// swagger:model ReversalPaymentUpdateRelationships
type ReversalPaymentUpdateRelationships struct {

	// ID of the payment resource related to the reversal
	Payment *RelationshipLinks `json:"payment,omitempty"`

	// reversal admission
	ReversalAdmission *ReversalPaymentUpdateRelationshipsReversalAdmission `json:"reversal_admission,omitempty"`

	// ID of the reversal submission resource related to the reversal
	ReversalSubmission *RelationshipLinks `json:"reversal_submission,omitempty"`
}

func ReversalPaymentUpdateRelationshipsWithDefaults(defaults client.Defaults) *ReversalPaymentUpdateRelationships {
	return &ReversalPaymentUpdateRelationships{

		Payment: RelationshipLinksWithDefaults(defaults),

		ReversalAdmission: ReversalPaymentUpdateRelationshipsReversalAdmissionWithDefaults(defaults),

		ReversalSubmission: RelationshipLinksWithDefaults(defaults),
	}
}

func (m *ReversalPaymentUpdateRelationships) WithPayment(payment RelationshipLinks) *ReversalPaymentUpdateRelationships {

	m.Payment = &payment

	return m
}

func (m *ReversalPaymentUpdateRelationships) WithoutPayment() *ReversalPaymentUpdateRelationships {
	m.Payment = nil
	return m
}

func (m *ReversalPaymentUpdateRelationships) WithReversalAdmission(reversalAdmission ReversalPaymentUpdateRelationshipsReversalAdmission) *ReversalPaymentUpdateRelationships {

	m.ReversalAdmission = &reversalAdmission

	return m
}

func (m *ReversalPaymentUpdateRelationships) WithoutReversalAdmission() *ReversalPaymentUpdateRelationships {
	m.ReversalAdmission = nil
	return m
}

func (m *ReversalPaymentUpdateRelationships) WithReversalSubmission(reversalSubmission RelationshipLinks) *ReversalPaymentUpdateRelationships {

	m.ReversalSubmission = &reversalSubmission

	return m
}

func (m *ReversalPaymentUpdateRelationships) WithoutReversalSubmission() *ReversalPaymentUpdateRelationships {
	m.ReversalSubmission = nil
	return m
}

// Validate validates this reversal payment update relationships
func (m *ReversalPaymentUpdateRelationships) Validate(formats strfmt.Registry) error {
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

func (m *ReversalPaymentUpdateRelationships) validatePayment(formats strfmt.Registry) error {

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

func (m *ReversalPaymentUpdateRelationships) validateReversalAdmission(formats strfmt.Registry) error {

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

func (m *ReversalPaymentUpdateRelationships) validateReversalSubmission(formats strfmt.Registry) error {

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
func (m *ReversalPaymentUpdateRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalPaymentUpdateRelationships) UnmarshalBinary(b []byte) error {
	var res ReversalPaymentUpdateRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalPaymentUpdateRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReversalPaymentUpdateRelationshipsReversalAdmission reversal payment update relationships reversal admission
// swagger:model ReversalPaymentUpdateRelationshipsReversalAdmission
type ReversalPaymentUpdateRelationshipsReversalAdmission struct {

	// Array of Reversal Admission resources related to the reversal
	Data []*ReversalAdmission `json:"data"`
}

func ReversalPaymentUpdateRelationshipsReversalAdmissionWithDefaults(defaults client.Defaults) *ReversalPaymentUpdateRelationshipsReversalAdmission {
	return &ReversalPaymentUpdateRelationshipsReversalAdmission{

		Data: make([]*ReversalAdmission, 0),
	}
}

func (m *ReversalPaymentUpdateRelationshipsReversalAdmission) WithData(data []*ReversalAdmission) *ReversalPaymentUpdateRelationshipsReversalAdmission {

	m.Data = data

	return m
}

// Validate validates this reversal payment update relationships reversal admission
func (m *ReversalPaymentUpdateRelationshipsReversalAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReversalPaymentUpdateRelationshipsReversalAdmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "reversal_admission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalPaymentUpdateRelationshipsReversalAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalPaymentUpdateRelationshipsReversalAdmission) UnmarshalBinary(b []byte) error {
	var res ReversalPaymentUpdateRelationshipsReversalAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalPaymentUpdateRelationshipsReversalAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
