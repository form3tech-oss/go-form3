// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v5/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReturnAdmission return admission
// swagger:model ReturnAdmission
type ReturnAdmission struct {

	// attributes
	Attributes *ReturnAdmissionAttributes `json:"attributes,omitempty"`

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
	Relationships *ReturnAdmissionRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ReturnAdmissionWithDefaults(defaults client.Defaults) *ReturnAdmission {
	return &ReturnAdmission{

		Attributes: ReturnAdmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("ReturnAdmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("ReturnAdmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("ReturnAdmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("ReturnAdmission", "organisation_id"),

		Relationships: ReturnAdmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReturnAdmission", "type"),

		Version: defaults.GetInt64Ptr("ReturnAdmission", "version"),
	}
}

func (m *ReturnAdmission) WithAttributes(attributes ReturnAdmissionAttributes) *ReturnAdmission {

	m.Attributes = &attributes

	return m
}

func (m *ReturnAdmission) WithoutAttributes() *ReturnAdmission {
	m.Attributes = nil
	return m
}

func (m *ReturnAdmission) WithCreatedOn(createdOn strfmt.DateTime) *ReturnAdmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *ReturnAdmission) WithoutCreatedOn() *ReturnAdmission {
	m.CreatedOn = nil
	return m
}

func (m *ReturnAdmission) WithID(id strfmt.UUID) *ReturnAdmission {

	m.ID = &id

	return m
}

func (m *ReturnAdmission) WithoutID() *ReturnAdmission {
	m.ID = nil
	return m
}

func (m *ReturnAdmission) WithModifiedOn(modifiedOn strfmt.DateTime) *ReturnAdmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *ReturnAdmission) WithoutModifiedOn() *ReturnAdmission {
	m.ModifiedOn = nil
	return m
}

func (m *ReturnAdmission) WithOrganisationID(organisationID strfmt.UUID) *ReturnAdmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *ReturnAdmission) WithoutOrganisationID() *ReturnAdmission {
	m.OrganisationID = nil
	return m
}

func (m *ReturnAdmission) WithRelationships(relationships ReturnAdmissionRelationships) *ReturnAdmission {

	m.Relationships = &relationships

	return m
}

func (m *ReturnAdmission) WithoutRelationships() *ReturnAdmission {
	m.Relationships = nil
	return m
}

func (m *ReturnAdmission) WithType(typeVar string) *ReturnAdmission {

	m.Type = typeVar

	return m
}

func (m *ReturnAdmission) WithVersion(version int64) *ReturnAdmission {

	m.Version = &version

	return m
}

func (m *ReturnAdmission) WithoutVersion() *ReturnAdmission {
	m.Version = nil
	return m
}

// Validate validates this return admission
func (m *ReturnAdmission) Validate(formats strfmt.Registry) error {
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

func (m *ReturnAdmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReturnAdmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *ReturnAdmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnAdmission) UnmarshalBinary(b []byte) error {
	var res ReturnAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReturnAdmissionAttributes return admission attributes
// swagger:model ReturnAdmissionAttributes
type ReturnAdmissionAttributes struct {

	// Date and time the payment admission was created
	// Read Only: true
	// Format: date-time
	AdmissionDatetime strfmt.DateTime `json:"admission_datetime,omitempty"`

	// Route taken for a return
	// Enum: [on_us]
	Route string `json:"route,omitempty"`

	// Refer to individual scheme where applicable
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// [Description](http://api-docs.form3.tech/api.html#enumerations-scheme-status-codes-for-bacs) of `scheme_status_code`
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// Cycle in which the payment will be settled
	// Minimum: 0
	SettlementCycle *int64 `json:"settlement_cycle,omitempty"`

	// Date on which the payment will be settled
	// Format: date
	SettlementDate *strfmt.Date `json:"settlement_date,omitempty"`

	// status
	Status ReturnAdmissionStatus `json:"status,omitempty"`

	// Further explanation of the status. [See here](http://api-docs.form3.tech/api.html#enumerations-payment-admission-status-reasons) for possible values.
	StatusReason string `json:"status_reason,omitempty"`

	// Scheme-specific unique ID (42 character string)
	UniqueSchemeID string `json:"unique_scheme_id,omitempty"`
}

func ReturnAdmissionAttributesWithDefaults(defaults client.Defaults) *ReturnAdmissionAttributes {
	return &ReturnAdmissionAttributes{

		AdmissionDatetime: defaults.GetStrfmtDateTime("ReturnAdmissionAttributes", "admission_datetime"),

		Route: defaults.GetString("ReturnAdmissionAttributes", "route"),

		SchemeStatusCode: defaults.GetString("ReturnAdmissionAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("ReturnAdmissionAttributes", "scheme_status_code_description"),

		SettlementCycle: defaults.GetInt64Ptr("ReturnAdmissionAttributes", "settlement_cycle"),

		SettlementDate: defaults.GetStrfmtDatePtr("ReturnAdmissionAttributes", "settlement_date"),

		// TODO Status: ReturnAdmissionStatus,

		StatusReason: defaults.GetString("ReturnAdmissionAttributes", "status_reason"),

		UniqueSchemeID: defaults.GetString("ReturnAdmissionAttributes", "unique_scheme_id"),
	}
}

func (m *ReturnAdmissionAttributes) WithAdmissionDatetime(admissionDatetime strfmt.DateTime) *ReturnAdmissionAttributes {

	m.AdmissionDatetime = admissionDatetime

	return m
}

func (m *ReturnAdmissionAttributes) WithRoute(route string) *ReturnAdmissionAttributes {

	m.Route = route

	return m
}

func (m *ReturnAdmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *ReturnAdmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *ReturnAdmissionAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *ReturnAdmissionAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *ReturnAdmissionAttributes) WithSettlementCycle(settlementCycle int64) *ReturnAdmissionAttributes {

	m.SettlementCycle = &settlementCycle

	return m
}

func (m *ReturnAdmissionAttributes) WithoutSettlementCycle() *ReturnAdmissionAttributes {
	m.SettlementCycle = nil
	return m
}

func (m *ReturnAdmissionAttributes) WithSettlementDate(settlementDate strfmt.Date) *ReturnAdmissionAttributes {

	m.SettlementDate = &settlementDate

	return m
}

func (m *ReturnAdmissionAttributes) WithoutSettlementDate() *ReturnAdmissionAttributes {
	m.SettlementDate = nil
	return m
}

func (m *ReturnAdmissionAttributes) WithStatus(status ReturnAdmissionStatus) *ReturnAdmissionAttributes {

	m.Status = status

	return m
}

func (m *ReturnAdmissionAttributes) WithStatusReason(statusReason string) *ReturnAdmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *ReturnAdmissionAttributes) WithUniqueSchemeID(uniqueSchemeID string) *ReturnAdmissionAttributes {

	m.UniqueSchemeID = uniqueSchemeID

	return m
}

// Validate validates this return admission attributes
func (m *ReturnAdmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

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

func (m *ReturnAdmissionAttributes) validateAdmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.AdmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"admission_datetime", "body", "date-time", m.AdmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

var returnAdmissionAttributesTypeRoutePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on_us"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		returnAdmissionAttributesTypeRoutePropEnum = append(returnAdmissionAttributesTypeRoutePropEnum, v)
	}
}

const (

	// ReturnAdmissionAttributesRouteOnUs captures enum value "on_us"
	ReturnAdmissionAttributesRouteOnUs string = "on_us"
)

// prop value enum
func (m *ReturnAdmissionAttributes) validateRouteEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, returnAdmissionAttributesTypeRoutePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ReturnAdmissionAttributes) validateRoute(formats strfmt.Registry) error {

	if swag.IsZero(m.Route) { // not required
		return nil
	}

	// value enum
	if err := m.validateRouteEnum("attributes"+"."+"route", "body", m.Route); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionAttributes) validateSettlementCycle(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementCycle) { // not required
		return nil
	}

	if err := validate.MinimumInt("attributes"+"."+"settlement_cycle", "body", int64(*m.SettlementCycle), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionAttributes) validateSettlementDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"settlement_date", "body", "date", m.SettlementDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionAttributes) validateStatus(formats strfmt.Registry) error {

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
func (m *ReturnAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res ReturnAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnAdmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
