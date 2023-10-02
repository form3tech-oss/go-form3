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

// ReturnSubmissionUpdate return submission update
// swagger:model ReturnSubmissionUpdate
type ReturnSubmissionUpdate struct {

	// attributes
	Attributes *ReturnSubmissionUpdateAttributes `json:"attributes,omitempty"`

	// Unique resource ID
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// relationships
	Relationships *ReturnSubmissionUpdateRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Required: true
	// Minimum: 0
	Version *int64 `json:"version"`
}

func ReturnSubmissionUpdateWithDefaults(defaults client.Defaults) *ReturnSubmissionUpdate {
	return &ReturnSubmissionUpdate{

		Attributes: ReturnSubmissionUpdateAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUIDPtr("ReturnSubmissionUpdate", "id"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("ReturnSubmissionUpdate", "organisation_id"),

		Relationships: ReturnSubmissionUpdateRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReturnSubmissionUpdate", "type"),

		Version: defaults.GetInt64Ptr("ReturnSubmissionUpdate", "version"),
	}
}

func (m *ReturnSubmissionUpdate) WithAttributes(attributes ReturnSubmissionUpdateAttributes) *ReturnSubmissionUpdate {

	m.Attributes = &attributes

	return m
}

func (m *ReturnSubmissionUpdate) WithoutAttributes() *ReturnSubmissionUpdate {
	m.Attributes = nil
	return m
}

func (m *ReturnSubmissionUpdate) WithID(id strfmt.UUID) *ReturnSubmissionUpdate {

	m.ID = &id

	return m
}

func (m *ReturnSubmissionUpdate) WithoutID() *ReturnSubmissionUpdate {
	m.ID = nil
	return m
}

func (m *ReturnSubmissionUpdate) WithOrganisationID(organisationID strfmt.UUID) *ReturnSubmissionUpdate {

	m.OrganisationID = &organisationID

	return m
}

func (m *ReturnSubmissionUpdate) WithoutOrganisationID() *ReturnSubmissionUpdate {
	m.OrganisationID = nil
	return m
}

func (m *ReturnSubmissionUpdate) WithRelationships(relationships ReturnSubmissionUpdateRelationships) *ReturnSubmissionUpdate {

	m.Relationships = &relationships

	return m
}

func (m *ReturnSubmissionUpdate) WithoutRelationships() *ReturnSubmissionUpdate {
	m.Relationships = nil
	return m
}

func (m *ReturnSubmissionUpdate) WithType(typeVar string) *ReturnSubmissionUpdate {

	m.Type = typeVar

	return m
}

func (m *ReturnSubmissionUpdate) WithVersion(version int64) *ReturnSubmissionUpdate {

	m.Version = &version

	return m
}

func (m *ReturnSubmissionUpdate) WithoutVersion() *ReturnSubmissionUpdate {
	m.Version = nil
	return m
}

// Validate validates this return submission update
func (m *ReturnSubmissionUpdate) Validate(formats strfmt.Registry) error {
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

func (m *ReturnSubmissionUpdate) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReturnSubmissionUpdate) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionUpdate) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionUpdate) validateRelationships(formats strfmt.Registry) error {

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

func (m *ReturnSubmissionUpdate) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionUpdate) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnSubmissionUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnSubmissionUpdate) UnmarshalBinary(b []byte) error {
	var res ReturnSubmissionUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnSubmissionUpdate) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReturnSubmissionUpdateAttributes return submission update attributes
// swagger:model ReturnSubmissionUpdateAttributes
type ReturnSubmissionUpdateAttributes struct {

	// Details of the account to which funds are redirected (if applicable)
	RedirectedAccountNumber string `json:"redirected_account_number,omitempty"`

	// Details of the bank to which funds are redirected (if applicable)
	RedirectedBankID string `json:"redirected_bank_id,omitempty"`

	// Route taken for a return
	// Enum: [on_us]
	Route string `json:"route,omitempty"`

	// Scheme-specific status (if submission has been submitted to a scheme)
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// [Description](http://api-docs.form3.tech/api.html#enumerations-scheme-status-codes-for-bacs) of `scheme_status_code`
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// Cycle in which the payment will be settled
	// Minimum: 0
	SettlementCycle *int64 `json:"settlement_cycle,omitempty"`

	// Date that the payment will be settled
	// Format: date
	SettlementDate *strfmt.Date `json:"settlement_date,omitempty"`

	// status
	Status ReturnSubmissionStatus `json:"status,omitempty"`

	// Description of the submission status
	StatusReason string `json:"status_reason,omitempty"`
}

func ReturnSubmissionUpdateAttributesWithDefaults(defaults client.Defaults) *ReturnSubmissionUpdateAttributes {
	return &ReturnSubmissionUpdateAttributes{

		RedirectedAccountNumber: defaults.GetString("ReturnSubmissionUpdateAttributes", "redirected_account_number"),

		RedirectedBankID: defaults.GetString("ReturnSubmissionUpdateAttributes", "redirected_bank_id"),

		Route: defaults.GetString("ReturnSubmissionUpdateAttributes", "route"),

		SchemeStatusCode: defaults.GetString("ReturnSubmissionUpdateAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("ReturnSubmissionUpdateAttributes", "scheme_status_code_description"),

		SettlementCycle: defaults.GetInt64Ptr("ReturnSubmissionUpdateAttributes", "settlement_cycle"),

		SettlementDate: defaults.GetStrfmtDatePtr("ReturnSubmissionUpdateAttributes", "settlement_date"),

		// TODO Status: ReturnSubmissionStatus,

		StatusReason: defaults.GetString("ReturnSubmissionUpdateAttributes", "status_reason"),
	}
}

func (m *ReturnSubmissionUpdateAttributes) WithRedirectedAccountNumber(redirectedAccountNumber string) *ReturnSubmissionUpdateAttributes {

	m.RedirectedAccountNumber = redirectedAccountNumber

	return m
}

func (m *ReturnSubmissionUpdateAttributes) WithRedirectedBankID(redirectedBankID string) *ReturnSubmissionUpdateAttributes {

	m.RedirectedBankID = redirectedBankID

	return m
}

func (m *ReturnSubmissionUpdateAttributes) WithRoute(route string) *ReturnSubmissionUpdateAttributes {

	m.Route = route

	return m
}

func (m *ReturnSubmissionUpdateAttributes) WithSchemeStatusCode(schemeStatusCode string) *ReturnSubmissionUpdateAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *ReturnSubmissionUpdateAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *ReturnSubmissionUpdateAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *ReturnSubmissionUpdateAttributes) WithSettlementCycle(settlementCycle int64) *ReturnSubmissionUpdateAttributes {

	m.SettlementCycle = &settlementCycle

	return m
}

func (m *ReturnSubmissionUpdateAttributes) WithoutSettlementCycle() *ReturnSubmissionUpdateAttributes {
	m.SettlementCycle = nil
	return m
}

func (m *ReturnSubmissionUpdateAttributes) WithSettlementDate(settlementDate strfmt.Date) *ReturnSubmissionUpdateAttributes {

	m.SettlementDate = &settlementDate

	return m
}

func (m *ReturnSubmissionUpdateAttributes) WithoutSettlementDate() *ReturnSubmissionUpdateAttributes {
	m.SettlementDate = nil
	return m
}

func (m *ReturnSubmissionUpdateAttributes) WithStatus(status ReturnSubmissionStatus) *ReturnSubmissionUpdateAttributes {

	m.Status = status

	return m
}

func (m *ReturnSubmissionUpdateAttributes) WithStatusReason(statusReason string) *ReturnSubmissionUpdateAttributes {

	m.StatusReason = statusReason

	return m
}

// Validate validates this return submission update attributes
func (m *ReturnSubmissionUpdateAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRoute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettlementCycle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettlementDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var returnSubmissionUpdateAttributesTypeRoutePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on_us"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		returnSubmissionUpdateAttributesTypeRoutePropEnum = append(returnSubmissionUpdateAttributesTypeRoutePropEnum, v)
	}
}

const (

	// ReturnSubmissionUpdateAttributesRouteOnUs captures enum value "on_us"
	ReturnSubmissionUpdateAttributesRouteOnUs string = "on_us"
)

// prop value enum
func (m *ReturnSubmissionUpdateAttributes) validateRouteEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, returnSubmissionUpdateAttributesTypeRoutePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ReturnSubmissionUpdateAttributes) validateRoute(formats strfmt.Registry) error {

	if swag.IsZero(m.Route) { // not required
		return nil
	}

	// value enum
	if err := m.validateRouteEnum("attributes"+"."+"route", "body", m.Route); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionUpdateAttributes) validateSettlementCycle(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementCycle) { // not required
		return nil
	}

	if err := validate.MinimumInt("attributes"+"."+"settlement_cycle", "body", int64(*m.SettlementCycle), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionUpdateAttributes) validateSettlementDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"settlement_date", "body", "date", m.SettlementDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnSubmissionUpdateAttributes) validateStatus(formats strfmt.Registry) error {

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
func (m *ReturnSubmissionUpdateAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnSubmissionUpdateAttributes) UnmarshalBinary(b []byte) error {
	var res ReturnSubmissionUpdateAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnSubmissionUpdateAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
