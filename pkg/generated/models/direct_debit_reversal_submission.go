// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DirectDebitReversalSubmission direct debit reversal submission
// swagger:model DirectDebitReversalSubmission
type DirectDebitReversalSubmission struct {

	// attributes
	Attributes *DirectDebitReversalSubmissionAttributes `json:"attributes,omitempty"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// id
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
	Relationships *DirectDebitReversalSubmissionRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func DirectDebitReversalSubmissionWithDefaults(defaults client.Defaults) *DirectDebitReversalSubmission {
	return &DirectDebitReversalSubmission{

		Attributes: DirectDebitReversalSubmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("DirectDebitReversalSubmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("DirectDebitReversalSubmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("DirectDebitReversalSubmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("DirectDebitReversalSubmission", "organisation_id"),

		Relationships: DirectDebitReversalSubmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("DirectDebitReversalSubmission", "type"),

		Version: defaults.GetInt64Ptr("DirectDebitReversalSubmission", "version"),
	}
}

func (m *DirectDebitReversalSubmission) WithAttributes(attributes DirectDebitReversalSubmissionAttributes) *DirectDebitReversalSubmission {

	m.Attributes = &attributes

	return m
}

func (m *DirectDebitReversalSubmission) WithoutAttributes() *DirectDebitReversalSubmission {
	m.Attributes = nil
	return m
}

func (m *DirectDebitReversalSubmission) WithCreatedOn(createdOn strfmt.DateTime) *DirectDebitReversalSubmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *DirectDebitReversalSubmission) WithoutCreatedOn() *DirectDebitReversalSubmission {
	m.CreatedOn = nil
	return m
}

func (m *DirectDebitReversalSubmission) WithID(id strfmt.UUID) *DirectDebitReversalSubmission {

	m.ID = &id

	return m
}

func (m *DirectDebitReversalSubmission) WithoutID() *DirectDebitReversalSubmission {
	m.ID = nil
	return m
}

func (m *DirectDebitReversalSubmission) WithModifiedOn(modifiedOn strfmt.DateTime) *DirectDebitReversalSubmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *DirectDebitReversalSubmission) WithoutModifiedOn() *DirectDebitReversalSubmission {
	m.ModifiedOn = nil
	return m
}

func (m *DirectDebitReversalSubmission) WithOrganisationID(organisationID strfmt.UUID) *DirectDebitReversalSubmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *DirectDebitReversalSubmission) WithoutOrganisationID() *DirectDebitReversalSubmission {
	m.OrganisationID = nil
	return m
}

func (m *DirectDebitReversalSubmission) WithRelationships(relationships DirectDebitReversalSubmissionRelationships) *DirectDebitReversalSubmission {

	m.Relationships = &relationships

	return m
}

func (m *DirectDebitReversalSubmission) WithoutRelationships() *DirectDebitReversalSubmission {
	m.Relationships = nil
	return m
}

func (m *DirectDebitReversalSubmission) WithType(typeVar string) *DirectDebitReversalSubmission {

	m.Type = typeVar

	return m
}

func (m *DirectDebitReversalSubmission) WithVersion(version int64) *DirectDebitReversalSubmission {

	m.Version = &version

	return m
}

func (m *DirectDebitReversalSubmission) WithoutVersion() *DirectDebitReversalSubmission {
	m.Version = nil
	return m
}

// Validate validates this direct debit reversal submission
func (m *DirectDebitReversalSubmission) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitReversalSubmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *DirectDebitReversalSubmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReversalSubmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReversalSubmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReversalSubmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReversalSubmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *DirectDebitReversalSubmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReversalSubmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReversalSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReversalSubmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitReversalSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReversalSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReversalSubmissionAttributes direct debit reversal submission attributes
// swagger:model DirectDebitReversalSubmissionAttributes
type DirectDebitReversalSubmissionAttributes struct {

	// destination gateway
	DestinationGateway string `json:"destination_gateway,omitempty"`

	// scheme status code
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// scheme status code description
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// status
	Status DirectDebitReversalSubmissionStatus `json:"status,omitempty"`

	// status reason
	StatusReason string `json:"status_reason,omitempty"`

	// submission datetime
	// Read Only: true
	// Format: date-time
	SubmissionDatetime strfmt.DateTime `json:"submission_datetime,omitempty"`

	// transaction start datetime
	// Read Only: true
	// Format: date-time
	TransactionStartDatetime strfmt.DateTime `json:"transaction_start_datetime,omitempty"`
}

func DirectDebitReversalSubmissionAttributesWithDefaults(defaults client.Defaults) *DirectDebitReversalSubmissionAttributes {
	return &DirectDebitReversalSubmissionAttributes{

		DestinationGateway: defaults.GetString("DirectDebitReversalSubmissionAttributes", "destination_gateway"),

		SchemeStatusCode: defaults.GetString("DirectDebitReversalSubmissionAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("DirectDebitReversalSubmissionAttributes", "scheme_status_code_description"),

		// TODO Status: DirectDebitReversalSubmissionStatus,

		StatusReason: defaults.GetString("DirectDebitReversalSubmissionAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTime("DirectDebitReversalSubmissionAttributes", "submission_datetime"),

		TransactionStartDatetime: defaults.GetStrfmtDateTime("DirectDebitReversalSubmissionAttributes", "transaction_start_datetime"),
	}
}

func (m *DirectDebitReversalSubmissionAttributes) WithDestinationGateway(destinationGateway string) *DirectDebitReversalSubmissionAttributes {

	m.DestinationGateway = destinationGateway

	return m
}

func (m *DirectDebitReversalSubmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *DirectDebitReversalSubmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *DirectDebitReversalSubmissionAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *DirectDebitReversalSubmissionAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *DirectDebitReversalSubmissionAttributes) WithStatus(status DirectDebitReversalSubmissionStatus) *DirectDebitReversalSubmissionAttributes {

	m.Status = status

	return m
}

func (m *DirectDebitReversalSubmissionAttributes) WithStatusReason(statusReason string) *DirectDebitReversalSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *DirectDebitReversalSubmissionAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *DirectDebitReversalSubmissionAttributes {

	m.SubmissionDatetime = submissionDatetime

	return m
}

func (m *DirectDebitReversalSubmissionAttributes) WithTransactionStartDatetime(transactionStartDatetime strfmt.DateTime) *DirectDebitReversalSubmissionAttributes {

	m.TransactionStartDatetime = transactionStartDatetime

	return m
}

// Validate validates this direct debit reversal submission attributes
func (m *DirectDebitReversalSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionStartDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReversalSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

func (m *DirectDebitReversalSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReversalSubmissionAttributes) validateTransactionStartDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionStartDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"transaction_start_datetime", "body", "date-time", m.TransactionStartDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReversalSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReversalSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res DirectDebitReversalSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReversalSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReversalSubmissionRelationships direct debit reversal submission relationships
// swagger:model DirectDebitReversalSubmissionRelationships
type DirectDebitReversalSubmissionRelationships struct {

	// direct debit
	DirectDebit *DirectDebitReversalSubmissionRelationshipsDirectDebit `json:"direct_debit,omitempty"`

	// direct debit reversal
	DirectDebitReversal *DirectDebitReversalSubmissionRelationshipsDirectDebitReversal `json:"direct_debit_reversal,omitempty"`
}

func DirectDebitReversalSubmissionRelationshipsWithDefaults(defaults client.Defaults) *DirectDebitReversalSubmissionRelationships {
	return &DirectDebitReversalSubmissionRelationships{

		DirectDebit: DirectDebitReversalSubmissionRelationshipsDirectDebitWithDefaults(defaults),

		DirectDebitReversal: DirectDebitReversalSubmissionRelationshipsDirectDebitReversalWithDefaults(defaults),
	}
}

func (m *DirectDebitReversalSubmissionRelationships) WithDirectDebit(directDebit DirectDebitReversalSubmissionRelationshipsDirectDebit) *DirectDebitReversalSubmissionRelationships {

	m.DirectDebit = &directDebit

	return m
}

func (m *DirectDebitReversalSubmissionRelationships) WithoutDirectDebit() *DirectDebitReversalSubmissionRelationships {
	m.DirectDebit = nil
	return m
}

func (m *DirectDebitReversalSubmissionRelationships) WithDirectDebitReversal(directDebitReversal DirectDebitReversalSubmissionRelationshipsDirectDebitReversal) *DirectDebitReversalSubmissionRelationships {

	m.DirectDebitReversal = &directDebitReversal

	return m
}

func (m *DirectDebitReversalSubmissionRelationships) WithoutDirectDebitReversal() *DirectDebitReversalSubmissionRelationships {
	m.DirectDebitReversal = nil
	return m
}

// Validate validates this direct debit reversal submission relationships
func (m *DirectDebitReversalSubmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectDebit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitReversal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReversalSubmissionRelationships) validateDirectDebit(formats strfmt.Registry) error {

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

func (m *DirectDebitReversalSubmissionRelationships) validateDirectDebitReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitReversal) { // not required
		return nil
	}

	if m.DirectDebitReversal != nil {
		if err := m.DirectDebitReversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit_reversal")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReversalSubmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReversalSubmissionRelationships) UnmarshalBinary(b []byte) error {
	var res DirectDebitReversalSubmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReversalSubmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReversalSubmissionRelationshipsDirectDebit direct debit reversal submission relationships direct debit
// swagger:model DirectDebitReversalSubmissionRelationshipsDirectDebit
type DirectDebitReversalSubmissionRelationshipsDirectDebit struct {

	// data
	Data []*DirectDebit `json:"data"`
}

func DirectDebitReversalSubmissionRelationshipsDirectDebitWithDefaults(defaults client.Defaults) *DirectDebitReversalSubmissionRelationshipsDirectDebit {
	return &DirectDebitReversalSubmissionRelationshipsDirectDebit{

		Data: make([]*DirectDebit, 0),
	}
}

func (m *DirectDebitReversalSubmissionRelationshipsDirectDebit) WithData(data []*DirectDebit) *DirectDebitReversalSubmissionRelationshipsDirectDebit {

	m.Data = data

	return m
}

// Validate validates this direct debit reversal submission relationships direct debit
func (m *DirectDebitReversalSubmissionRelationshipsDirectDebit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReversalSubmissionRelationshipsDirectDebit) validateData(formats strfmt.Registry) error {

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
func (m *DirectDebitReversalSubmissionRelationshipsDirectDebit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReversalSubmissionRelationshipsDirectDebit) UnmarshalBinary(b []byte) error {
	var res DirectDebitReversalSubmissionRelationshipsDirectDebit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReversalSubmissionRelationshipsDirectDebit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReversalSubmissionRelationshipsDirectDebitReversal direct debit reversal submission relationships direct debit reversal
// swagger:model DirectDebitReversalSubmissionRelationshipsDirectDebitReversal
type DirectDebitReversalSubmissionRelationshipsDirectDebitReversal struct {

	// data
	Data []*DirectDebitReversal `json:"data"`
}

func DirectDebitReversalSubmissionRelationshipsDirectDebitReversalWithDefaults(defaults client.Defaults) *DirectDebitReversalSubmissionRelationshipsDirectDebitReversal {
	return &DirectDebitReversalSubmissionRelationshipsDirectDebitReversal{

		Data: make([]*DirectDebitReversal, 0),
	}
}

func (m *DirectDebitReversalSubmissionRelationshipsDirectDebitReversal) WithData(data []*DirectDebitReversal) *DirectDebitReversalSubmissionRelationshipsDirectDebitReversal {

	m.Data = data

	return m
}

// Validate validates this direct debit reversal submission relationships direct debit reversal
func (m *DirectDebitReversalSubmissionRelationshipsDirectDebitReversal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReversalSubmissionRelationshipsDirectDebitReversal) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "direct_debit_reversal" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReversalSubmissionRelationshipsDirectDebitReversal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReversalSubmissionRelationshipsDirectDebitReversal) UnmarshalBinary(b []byte) error {
	var res DirectDebitReversalSubmissionRelationshipsDirectDebitReversal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReversalSubmissionRelationshipsDirectDebitReversal) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
