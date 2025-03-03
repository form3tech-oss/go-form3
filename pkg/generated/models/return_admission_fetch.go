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

// ReturnAdmissionFetch return admission fetch
// swagger:model ReturnAdmissionFetch
type ReturnAdmissionFetch struct {

	// attributes
	Attributes *ReturnAdmissionFetchAttributes `json:"attributes,omitempty"`

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
	Relationships *ReturnAdmissionFetchRelationships `json:"relationships,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func ReturnAdmissionFetchWithDefaults(defaults client.Defaults) *ReturnAdmissionFetch {
	return &ReturnAdmissionFetch{

		Attributes: ReturnAdmissionFetchAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("ReturnAdmissionFetch", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("ReturnAdmissionFetch", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("ReturnAdmissionFetch", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("ReturnAdmissionFetch", "organisation_id"),

		Relationships: ReturnAdmissionFetchRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("ReturnAdmissionFetch", "type"),

		Version: defaults.GetInt64Ptr("ReturnAdmissionFetch", "version"),
	}
}

func (m *ReturnAdmissionFetch) WithAttributes(attributes ReturnAdmissionFetchAttributes) *ReturnAdmissionFetch {

	m.Attributes = &attributes

	return m
}

func (m *ReturnAdmissionFetch) WithoutAttributes() *ReturnAdmissionFetch {
	m.Attributes = nil
	return m
}

func (m *ReturnAdmissionFetch) WithCreatedOn(createdOn strfmt.DateTime) *ReturnAdmissionFetch {

	m.CreatedOn = &createdOn

	return m
}

func (m *ReturnAdmissionFetch) WithoutCreatedOn() *ReturnAdmissionFetch {
	m.CreatedOn = nil
	return m
}

func (m *ReturnAdmissionFetch) WithID(id strfmt.UUID) *ReturnAdmissionFetch {

	m.ID = &id

	return m
}

func (m *ReturnAdmissionFetch) WithoutID() *ReturnAdmissionFetch {
	m.ID = nil
	return m
}

func (m *ReturnAdmissionFetch) WithModifiedOn(modifiedOn strfmt.DateTime) *ReturnAdmissionFetch {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *ReturnAdmissionFetch) WithoutModifiedOn() *ReturnAdmissionFetch {
	m.ModifiedOn = nil
	return m
}

func (m *ReturnAdmissionFetch) WithOrganisationID(organisationID strfmt.UUID) *ReturnAdmissionFetch {

	m.OrganisationID = &organisationID

	return m
}

func (m *ReturnAdmissionFetch) WithoutOrganisationID() *ReturnAdmissionFetch {
	m.OrganisationID = nil
	return m
}

func (m *ReturnAdmissionFetch) WithRelationships(relationships ReturnAdmissionFetchRelationships) *ReturnAdmissionFetch {

	m.Relationships = &relationships

	return m
}

func (m *ReturnAdmissionFetch) WithoutRelationships() *ReturnAdmissionFetch {
	m.Relationships = nil
	return m
}

func (m *ReturnAdmissionFetch) WithType(typeVar string) *ReturnAdmissionFetch {

	m.Type = typeVar

	return m
}

func (m *ReturnAdmissionFetch) WithVersion(version int64) *ReturnAdmissionFetch {

	m.Version = &version

	return m
}

func (m *ReturnAdmissionFetch) WithoutVersion() *ReturnAdmissionFetch {
	m.Version = nil
	return m
}

// Validate validates this return admission fetch
func (m *ReturnAdmissionFetch) Validate(formats strfmt.Registry) error {
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

func (m *ReturnAdmissionFetch) validateAttributes(formats strfmt.Registry) error {

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

func (m *ReturnAdmissionFetch) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionFetch) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionFetch) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionFetch) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionFetch) validateRelationships(formats strfmt.Registry) error {

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

func (m *ReturnAdmissionFetch) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", m.Type, `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionFetch) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnAdmissionFetch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnAdmissionFetch) UnmarshalBinary(b []byte) error {
	var res ReturnAdmissionFetch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnAdmissionFetch) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// ReturnAdmissionFetchAttributes return admission fetch attributes
// swagger:model ReturnAdmissionFetchAttributes
type ReturnAdmissionFetchAttributes struct {

	// Date and time the payment admission was created
	// Read Only: true
	// Format: date-time
	AdmissionDatetime strfmt.DateTime `json:"admission_datetime,omitempty"`

	// posting status
	PostingStatus PostingStatus `json:"posting_status,omitempty"`

	// Additional payment reference assigned by the scheme
	ReferenceID string `json:"reference_id,omitempty"`

	// Route taken for a return
	// Enum: ["on_us","xp"]
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

func ReturnAdmissionFetchAttributesWithDefaults(defaults client.Defaults) *ReturnAdmissionFetchAttributes {
	return &ReturnAdmissionFetchAttributes{

		AdmissionDatetime: defaults.GetStrfmtDateTime("ReturnAdmissionFetchAttributes", "admission_datetime"),

		// TODO PostingStatus: PostingStatus,

		ReferenceID: defaults.GetString("ReturnAdmissionFetchAttributes", "reference_id"),

		Route: defaults.GetString("ReturnAdmissionFetchAttributes", "route"),

		SchemeStatusCode: defaults.GetString("ReturnAdmissionFetchAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("ReturnAdmissionFetchAttributes", "scheme_status_code_description"),

		SettlementCycle: defaults.GetInt64Ptr("ReturnAdmissionFetchAttributes", "settlement_cycle"),

		SettlementDate: defaults.GetStrfmtDatePtr("ReturnAdmissionFetchAttributes", "settlement_date"),

		// TODO Status: ReturnAdmissionStatus,

		StatusReason: defaults.GetString("ReturnAdmissionFetchAttributes", "status_reason"),

		UniqueSchemeID: defaults.GetString("ReturnAdmissionFetchAttributes", "unique_scheme_id"),
	}
}

func (m *ReturnAdmissionFetchAttributes) WithAdmissionDatetime(admissionDatetime strfmt.DateTime) *ReturnAdmissionFetchAttributes {

	m.AdmissionDatetime = admissionDatetime

	return m
}

func (m *ReturnAdmissionFetchAttributes) WithPostingStatus(postingStatus PostingStatus) *ReturnAdmissionFetchAttributes {

	m.PostingStatus = postingStatus

	return m
}

func (m *ReturnAdmissionFetchAttributes) WithReferenceID(referenceID string) *ReturnAdmissionFetchAttributes {

	m.ReferenceID = referenceID

	return m
}

func (m *ReturnAdmissionFetchAttributes) WithRoute(route string) *ReturnAdmissionFetchAttributes {

	m.Route = route

	return m
}

func (m *ReturnAdmissionFetchAttributes) WithSchemeStatusCode(schemeStatusCode string) *ReturnAdmissionFetchAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *ReturnAdmissionFetchAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *ReturnAdmissionFetchAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *ReturnAdmissionFetchAttributes) WithSettlementCycle(settlementCycle int64) *ReturnAdmissionFetchAttributes {

	m.SettlementCycle = &settlementCycle

	return m
}

func (m *ReturnAdmissionFetchAttributes) WithoutSettlementCycle() *ReturnAdmissionFetchAttributes {
	m.SettlementCycle = nil
	return m
}

func (m *ReturnAdmissionFetchAttributes) WithSettlementDate(settlementDate strfmt.Date) *ReturnAdmissionFetchAttributes {

	m.SettlementDate = &settlementDate

	return m
}

func (m *ReturnAdmissionFetchAttributes) WithoutSettlementDate() *ReturnAdmissionFetchAttributes {
	m.SettlementDate = nil
	return m
}

func (m *ReturnAdmissionFetchAttributes) WithStatus(status ReturnAdmissionStatus) *ReturnAdmissionFetchAttributes {

	m.Status = status

	return m
}

func (m *ReturnAdmissionFetchAttributes) WithStatusReason(statusReason string) *ReturnAdmissionFetchAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *ReturnAdmissionFetchAttributes) WithUniqueSchemeID(uniqueSchemeID string) *ReturnAdmissionFetchAttributes {

	m.UniqueSchemeID = uniqueSchemeID

	return m
}

// Validate validates this return admission fetch attributes
func (m *ReturnAdmissionFetchAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostingStatus(formats); err != nil {
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

func (m *ReturnAdmissionFetchAttributes) validateAdmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.AdmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"admission_datetime", "body", "date-time", m.AdmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionFetchAttributes) validatePostingStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.PostingStatus) { // not required
		return nil
	}

	if err := m.PostingStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("attributes" + "." + "posting_status")
		}
		return err
	}

	return nil
}

var returnAdmissionFetchAttributesTypeRoutePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["on_us","xp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		returnAdmissionFetchAttributesTypeRoutePropEnum = append(returnAdmissionFetchAttributesTypeRoutePropEnum, v)
	}
}

const (

	// ReturnAdmissionFetchAttributesRouteOnUs captures enum value "on_us"
	ReturnAdmissionFetchAttributesRouteOnUs string = "on_us"

	// ReturnAdmissionFetchAttributesRouteXp captures enum value "xp"
	ReturnAdmissionFetchAttributesRouteXp string = "xp"
)

// prop value enum
func (m *ReturnAdmissionFetchAttributes) validateRouteEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, returnAdmissionFetchAttributesTypeRoutePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ReturnAdmissionFetchAttributes) validateRoute(formats strfmt.Registry) error {

	if swag.IsZero(m.Route) { // not required
		return nil
	}

	// value enum
	if err := m.validateRouteEnum("attributes"+"."+"route", "body", m.Route); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionFetchAttributes) validateSettlementCycle(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementCycle) { // not required
		return nil
	}

	if err := validate.MinimumInt("attributes"+"."+"settlement_cycle", "body", int64(*m.SettlementCycle), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionFetchAttributes) validateSettlementDate(formats strfmt.Registry) error {

	if swag.IsZero(m.SettlementDate) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"settlement_date", "body", "date", m.SettlementDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReturnAdmissionFetchAttributes) validateStatus(formats strfmt.Registry) error {

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
func (m *ReturnAdmissionFetchAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnAdmissionFetchAttributes) UnmarshalBinary(b []byte) error {
	var res ReturnAdmissionFetchAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnAdmissionFetchAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
