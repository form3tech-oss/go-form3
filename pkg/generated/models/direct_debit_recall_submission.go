// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DirectDebitRecallSubmission direct debit recall submission
// swagger:model DirectDebitRecallSubmission
type DirectDebitRecallSubmission struct {

	// attributes
	Attributes *DirectDebitRecallSubmissionAttributes `json:"attributes,omitempty"`

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
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *DirectDebitRecallSubmissionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func DirectDebitRecallSubmissionWithDefaults(defaults client.Defaults) *DirectDebitRecallSubmission {
	return &DirectDebitRecallSubmission{

		Attributes: DirectDebitRecallSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("DirectDebitRecallSubmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("DirectDebitRecallSubmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("DirectDebitRecallSubmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("DirectDebitRecallSubmission", "organisation_id"),

		Relationships: DirectDebitRecallSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("DirectDebitRecallSubmission", "type"),

		Version: defaults.GetInt64Ptr("DirectDebitRecallSubmission", "version"),
	}
}

func (m *DirectDebitRecallSubmission) WithAttributes(attributes DirectDebitRecallSubmissionAttributes) *DirectDebitRecallSubmission {

	m.Attributes = &attributes

	return m
}

func (m *DirectDebitRecallSubmission) WithoutAttributes() *DirectDebitRecallSubmission {
	m.Attributes = nil
	return m
}

func (m *DirectDebitRecallSubmission) WithCreatedOn(createdOn strfmt.DateTime) *DirectDebitRecallSubmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *DirectDebitRecallSubmission) WithoutCreatedOn() *DirectDebitRecallSubmission {
	m.CreatedOn = nil
	return m
}

func (m *DirectDebitRecallSubmission) WithID(id strfmt.UUID) *DirectDebitRecallSubmission {

	m.ID = &id

	return m
}

func (m *DirectDebitRecallSubmission) WithoutID() *DirectDebitRecallSubmission {
	m.ID = nil
	return m
}

func (m *DirectDebitRecallSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *DirectDebitRecallSubmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *DirectDebitRecallSubmission) WithoutModifiedOn() *DirectDebitRecallSubmission {
	m.ModifiedOn = nil
	return m
}

func (m *DirectDebitRecallSubmission) WithOrganisationID(organisationID strfmt.UUID) *DirectDebitRecallSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *DirectDebitRecallSubmission) WithoutOrganisationID() *DirectDebitRecallSubmission {
	m.OrganisationID = nil
	return m
}

func (m *DirectDebitRecallSubmission) WithRelationships(relationships DirectDebitRecallSubmissionRelationships) *DirectDebitRecallSubmission {

	m.Relationships = &relationships

	return m
}

func (m *DirectDebitRecallSubmission) WithoutRelationships() *DirectDebitRecallSubmission {
	m.Relationships = nil
	return m
}

func (m *DirectDebitRecallSubmission) WithType(typeVar string) *DirectDebitRecallSubmission {

	m.Type = typeVar

	return m
}

func (m *DirectDebitRecallSubmission) WithVersion(version int64) *DirectDebitRecallSubmission {

	m.Version = &version

	return m
}

func (m *DirectDebitRecallSubmission) WithoutVersion() *DirectDebitRecallSubmission {
	m.Version = nil
	return m
}

// Validate validates this direct debit recall submission
func (m *DirectDebitRecallSubmission) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitRecallSubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *DirectDebitRecallSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitRecallSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitRecallSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitRecallSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitRecallSubmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *DirectDebitRecallSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitRecallSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitRecallSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRecallSubmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitRecallSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRecallSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitRecallSubmissionAttributes direct debit recall submission attributes
// swagger:model DirectDebitRecallSubmissionAttributes
type DirectDebitRecallSubmissionAttributes struct {

	// destination gateway
	DestinationGateway string `json:"destination_gateway,omitempty"`

	// Scheme-specific status (if submission has been submitted to a scheme)
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// status
	Status DirectDebitRecallSubmissionStatus `json:"status,omitempty"`

	// Reason for submission failure if status is `delivery_failed`
	StatusReason string `json:"status_reason,omitempty"`

	// Date and time of the submission
	// Read Only: true
	// Format: date-time
	SubmissionDatetime *strfmt.DateTime `json:"submission_datetime,omitempty"`
}

func DirectDebitRecallSubmissionAttributesWithDefaults(defaults client.Defaults) *DirectDebitRecallSubmissionAttributes {
	return &DirectDebitRecallSubmissionAttributes{

		DestinationGateway: defaults.GetString("DirectDebitRecallSubmissionAttributes", "destination_gateway"),

		SchemeStatusCode: defaults.GetString("DirectDebitRecallSubmissionAttributes", "scheme_status_code"),

		// TODO Status: DirectDebitRecallSubmissionStatus,

		StatusReason: defaults.GetString("DirectDebitRecallSubmissionAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTimePtr("DirectDebitRecallSubmissionAttributes", "submission_datetime"),
	}
}

func (m *DirectDebitRecallSubmissionAttributes) WithDestinationGateway(destinationGateway string) *DirectDebitRecallSubmissionAttributes {

	m.DestinationGateway = destinationGateway

	return m
}

func (m *DirectDebitRecallSubmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *DirectDebitRecallSubmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *DirectDebitRecallSubmissionAttributes) WithStatus(status DirectDebitRecallSubmissionStatus) *DirectDebitRecallSubmissionAttributes {

	m.Status = status

	return m
}

func (m *DirectDebitRecallSubmissionAttributes) WithStatusReason(statusReason string) *DirectDebitRecallSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *DirectDebitRecallSubmissionAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *DirectDebitRecallSubmissionAttributes {

	m.SubmissionDatetime = &submissionDatetime

	return m
}

func (m *DirectDebitRecallSubmissionAttributes) WithoutSubmissionDatetime() *DirectDebitRecallSubmissionAttributes {
	m.SubmissionDatetime = nil
	return m
}

// Validate validates this direct debit recall submission attributes
func (m *DirectDebitRecallSubmissionAttributes) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitRecallSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

func (m *DirectDebitRecallSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitRecallSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRecallSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res DirectDebitRecallSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRecallSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitRecallSubmissionRelationships direct debit recall submission relationships
// swagger:model DirectDebitRecallSubmissionRelationships
type DirectDebitRecallSubmissionRelationships struct {

	// direct debit
	DirectDebit *DirectDebitRecallSubmissionRelationshipsDirectDebit `json:"direct_debit,omitempty"`

	// direct debit recall
	DirectDebitRecall *DirectDebitRecallSubmissionRelationshipsDirectDebitRecall `json:"direct_debit_recall,omitempty"`
}

func DirectDebitRecallSubmissionRelationshipsWithDefaults(defaults client.Defaults) *DirectDebitRecallSubmissionRelationships {
	return &DirectDebitRecallSubmissionRelationships{

		DirectDebit: DirectDebitRecallSubmissionRelationshipsDirectDebitWithDefaults(defaults),

		DirectDebitRecall: DirectDebitRecallSubmissionRelationshipsDirectDebitRecallWithDefaults(defaults),
	}
}

func (m *DirectDebitRecallSubmissionRelationships) WithDirectDebit(directDebit DirectDebitRecallSubmissionRelationshipsDirectDebit) *DirectDebitRecallSubmissionRelationships {

	m.DirectDebit = &directDebit

	return m
}

func (m *DirectDebitRecallSubmissionRelationships) WithoutDirectDebit() *DirectDebitRecallSubmissionRelationships {
	m.DirectDebit = nil
	return m
}

func (m *DirectDebitRecallSubmissionRelationships) WithDirectDebitRecall(directDebitRecall DirectDebitRecallSubmissionRelationshipsDirectDebitRecall) *DirectDebitRecallSubmissionRelationships {

	m.DirectDebitRecall = &directDebitRecall

	return m
}

func (m *DirectDebitRecallSubmissionRelationships) WithoutDirectDebitRecall() *DirectDebitRecallSubmissionRelationships {
	m.DirectDebitRecall = nil
	return m
}

// Validate validates this direct debit recall submission relationships
func (m *DirectDebitRecallSubmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectDebit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitRecall(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitRecallSubmissionRelationships) validateDirectDebit(formats strfmt.Registry) error {

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

func (m *DirectDebitRecallSubmissionRelationships) validateDirectDebitRecall(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitRecall) { // not required
		return nil
	}

	if m.DirectDebitRecall != nil {
		if err := m.DirectDebitRecall.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit_recall")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitRecallSubmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRecallSubmissionRelationships) UnmarshalBinary(b []byte) error {
	var res DirectDebitRecallSubmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRecallSubmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitRecallSubmissionRelationshipsDirectDebit direct debit recall submission relationships direct debit
// swagger:model DirectDebitRecallSubmissionRelationshipsDirectDebit
type DirectDebitRecallSubmissionRelationshipsDirectDebit struct {

	// data
	Data []*DirectDebit `json:"data"`
}

func DirectDebitRecallSubmissionRelationshipsDirectDebitWithDefaults(defaults client.Defaults) *DirectDebitRecallSubmissionRelationshipsDirectDebit {
	return &DirectDebitRecallSubmissionRelationshipsDirectDebit{

		Data: make([]*DirectDebit, 0),
	}
}

func (m *DirectDebitRecallSubmissionRelationshipsDirectDebit) WithData(data []*DirectDebit) *DirectDebitRecallSubmissionRelationshipsDirectDebit {

	m.Data = data

	return m
}

// Validate validates this direct debit recall submission relationships direct debit
func (m *DirectDebitRecallSubmissionRelationshipsDirectDebit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitRecallSubmissionRelationshipsDirectDebit) validateData(formats strfmt.Registry) error {

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
func (m *DirectDebitRecallSubmissionRelationshipsDirectDebit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRecallSubmissionRelationshipsDirectDebit) UnmarshalBinary(b []byte) error {
	var res DirectDebitRecallSubmissionRelationshipsDirectDebit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRecallSubmissionRelationshipsDirectDebit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitRecallSubmissionRelationshipsDirectDebitRecall direct debit recall submission relationships direct debit recall
// swagger:model DirectDebitRecallSubmissionRelationshipsDirectDebitRecall
type DirectDebitRecallSubmissionRelationshipsDirectDebitRecall struct {

	// data
	Data []*DirectDebitRecall `json:"data"`
}

func DirectDebitRecallSubmissionRelationshipsDirectDebitRecallWithDefaults(defaults client.Defaults) *DirectDebitRecallSubmissionRelationshipsDirectDebitRecall {
	return &DirectDebitRecallSubmissionRelationshipsDirectDebitRecall{

		Data: make([]*DirectDebitRecall, 0),
	}
}

func (m *DirectDebitRecallSubmissionRelationshipsDirectDebitRecall) WithData(data []*DirectDebitRecall) *DirectDebitRecallSubmissionRelationshipsDirectDebitRecall {

	m.Data = data

	return m
}

// Validate validates this direct debit recall submission relationships direct debit recall
func (m *DirectDebitRecallSubmissionRelationshipsDirectDebitRecall) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitRecallSubmissionRelationshipsDirectDebitRecall) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "direct_debit_recall" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitRecallSubmissionRelationshipsDirectDebitRecall) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitRecallSubmissionRelationshipsDirectDebitRecall) UnmarshalBinary(b []byte) error {
	var res DirectDebitRecallSubmissionRelationshipsDirectDebitRecall
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitRecallSubmissionRelationshipsDirectDebitRecall) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
