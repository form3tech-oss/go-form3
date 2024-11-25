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

// DirectDebitReturnReversalAdmission direct debit return reversal admission
// swagger:model DirectDebitReturnReversalAdmission
type DirectDebitReturnReversalAdmission struct {

	// attributes
	// Required: true
	Attributes *DirectDebitReturnReversalAdmissionAttributes `json:"attributes"`

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
	Relationships *DirectDebitReturnReversalAdmissionRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func DirectDebitReturnReversalAdmissionWithDefaults(defaults client.Defaults) *DirectDebitReturnReversalAdmission {
	return &DirectDebitReturnReversalAdmission{

		Attributes: DirectDebitReturnReversalAdmissionAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("DirectDebitReturnReversalAdmission", "created_on"),

		ID: defaults.GetStrfmtUUIDPtr("DirectDebitReturnReversalAdmission", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("DirectDebitReturnReversalAdmission", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("DirectDebitReturnReversalAdmission", "organisation_id"),

		Relationships: DirectDebitReturnReversalAdmissionRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("DirectDebitReturnReversalAdmission", "type"),

		Version: defaults.GetInt64Ptr("DirectDebitReturnReversalAdmission", "version"),
	}
}

func (m *DirectDebitReturnReversalAdmission) WithAttributes(attributes DirectDebitReturnReversalAdmissionAttributes) *DirectDebitReturnReversalAdmission {

	m.Attributes = &attributes

	return m
}

func (m *DirectDebitReturnReversalAdmission) WithoutAttributes() *DirectDebitReturnReversalAdmission {
	m.Attributes = nil
	return m
}

func (m *DirectDebitReturnReversalAdmission) WithCreatedOn(createdOn strfmt.DateTime) *DirectDebitReturnReversalAdmission {

	m.CreatedOn = &createdOn

	return m
}

func (m *DirectDebitReturnReversalAdmission) WithoutCreatedOn() *DirectDebitReturnReversalAdmission {
	m.CreatedOn = nil
	return m
}

func (m *DirectDebitReturnReversalAdmission) WithID(id strfmt.UUID) *DirectDebitReturnReversalAdmission {

	m.ID = &id

	return m
}

func (m *DirectDebitReturnReversalAdmission) WithoutID() *DirectDebitReturnReversalAdmission {
	m.ID = nil
	return m
}

func (m *DirectDebitReturnReversalAdmission) WithModifiedOn(modifiedOn strfmt.DateTime) *DirectDebitReturnReversalAdmission {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *DirectDebitReturnReversalAdmission) WithoutModifiedOn() *DirectDebitReturnReversalAdmission {
	m.ModifiedOn = nil
	return m
}

func (m *DirectDebitReturnReversalAdmission) WithOrganisationID(organisationID strfmt.UUID) *DirectDebitReturnReversalAdmission {

	m.OrganisationID = &organisationID

	return m
}

func (m *DirectDebitReturnReversalAdmission) WithoutOrganisationID() *DirectDebitReturnReversalAdmission {
	m.OrganisationID = nil
	return m
}

func (m *DirectDebitReturnReversalAdmission) WithRelationships(relationships DirectDebitReturnReversalAdmissionRelationships) *DirectDebitReturnReversalAdmission {

	m.Relationships = &relationships

	return m
}

func (m *DirectDebitReturnReversalAdmission) WithoutRelationships() *DirectDebitReturnReversalAdmission {
	m.Relationships = nil
	return m
}

func (m *DirectDebitReturnReversalAdmission) WithType(typeVar string) *DirectDebitReturnReversalAdmission {

	m.Type = typeVar

	return m
}

func (m *DirectDebitReturnReversalAdmission) WithVersion(version int64) *DirectDebitReturnReversalAdmission {

	m.Version = &version

	return m
}

func (m *DirectDebitReturnReversalAdmission) WithoutVersion() *DirectDebitReturnReversalAdmission {
	m.Version = nil
	return m
}

// Validate validates this direct debit return reversal admission
func (m *DirectDebitReturnReversalAdmission) Validate(formats strfmt.Registry) error {
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

func (m *DirectDebitReturnReversalAdmission) validateAttributes(formats strfmt.Registry) error {

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

func (m *DirectDebitReturnReversalAdmission) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnReversalAdmission) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnReversalAdmission) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnReversalAdmission) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnReversalAdmission) validateRelationships(formats strfmt.Registry) error {

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

func (m *DirectDebitReturnReversalAdmission) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnReversalAdmission) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnReversalAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnReversalAdmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnReversalAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnReversalAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnReversalAdmissionAttributes direct debit return reversal admission attributes
// swagger:model DirectDebitReturnReversalAdmissionAttributes
type DirectDebitReturnReversalAdmissionAttributes struct {

	// scheme status code
	// Required: true
	// Min Length: 1
	SchemeStatusCode *string `json:"scheme_status_code"`

	// scheme status code description
	// Required: true
	// Min Length: 1
	SchemeStatusCodeDescription *string `json:"scheme_status_code_description"`
}

func DirectDebitReturnReversalAdmissionAttributesWithDefaults(defaults client.Defaults) *DirectDebitReturnReversalAdmissionAttributes {
	return &DirectDebitReturnReversalAdmissionAttributes{

		SchemeStatusCode: defaults.GetStringPtr("DirectDebitReturnReversalAdmissionAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetStringPtr("DirectDebitReturnReversalAdmissionAttributes", "scheme_status_code_description"),
	}
}

func (m *DirectDebitReturnReversalAdmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *DirectDebitReturnReversalAdmissionAttributes {

	m.SchemeStatusCode = &schemeStatusCode

	return m
}

func (m *DirectDebitReturnReversalAdmissionAttributes) WithoutSchemeStatusCode() *DirectDebitReturnReversalAdmissionAttributes {
	m.SchemeStatusCode = nil
	return m
}

func (m *DirectDebitReturnReversalAdmissionAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *DirectDebitReturnReversalAdmissionAttributes {

	m.SchemeStatusCodeDescription = &schemeStatusCodeDescription

	return m
}

func (m *DirectDebitReturnReversalAdmissionAttributes) WithoutSchemeStatusCodeDescription() *DirectDebitReturnReversalAdmissionAttributes {
	m.SchemeStatusCodeDescription = nil
	return m
}

// Validate validates this direct debit return reversal admission attributes
func (m *DirectDebitReturnReversalAdmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchemeStatusCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemeStatusCodeDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnReversalAdmissionAttributes) validateSchemeStatusCode(formats strfmt.Registry) error {

	if err := validate.Required("attributes"+"."+"scheme_status_code", "body", m.SchemeStatusCode); err != nil {
		return err
	}

	if err := validate.MinLength("attributes"+"."+"scheme_status_code", "body", string(*m.SchemeStatusCode), 1); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnReversalAdmissionAttributes) validateSchemeStatusCodeDescription(formats strfmt.Registry) error {

	if err := validate.Required("attributes"+"."+"scheme_status_code_description", "body", m.SchemeStatusCodeDescription); err != nil {
		return err
	}

	if err := validate.MinLength("attributes"+"."+"scheme_status_code_description", "body", string(*m.SchemeStatusCodeDescription), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnReversalAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnReversalAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnReversalAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnReversalAdmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnReversalAdmissionRelationships direct debit return reversal admission relationships
// swagger:model DirectDebitReturnReversalAdmissionRelationships
type DirectDebitReturnReversalAdmissionRelationships struct {

	// direct debit
	DirectDebit *DirectDebitReturnReversalAdmissionRelationshipsDirectDebit `json:"direct_debit,omitempty"`

	// direct debit return
	DirectDebitReturn *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturn `json:"direct_debit_return,omitempty"`

	// direct debit return reversal
	DirectDebitReturnReversal *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversal `json:"direct_debit_return_reversal,omitempty"`
}

func DirectDebitReturnReversalAdmissionRelationshipsWithDefaults(defaults client.Defaults) *DirectDebitReturnReversalAdmissionRelationships {
	return &DirectDebitReturnReversalAdmissionRelationships{

		DirectDebit: DirectDebitReturnReversalAdmissionRelationshipsDirectDebitWithDefaults(defaults),

		DirectDebitReturn: DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnWithDefaults(defaults),

		DirectDebitReturnReversal: DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversalWithDefaults(defaults),
	}
}

func (m *DirectDebitReturnReversalAdmissionRelationships) WithDirectDebit(directDebit DirectDebitReturnReversalAdmissionRelationshipsDirectDebit) *DirectDebitReturnReversalAdmissionRelationships {

	m.DirectDebit = &directDebit

	return m
}

func (m *DirectDebitReturnReversalAdmissionRelationships) WithoutDirectDebit() *DirectDebitReturnReversalAdmissionRelationships {
	m.DirectDebit = nil
	return m
}

func (m *DirectDebitReturnReversalAdmissionRelationships) WithDirectDebitReturn(directDebitReturn DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturn) *DirectDebitReturnReversalAdmissionRelationships {

	m.DirectDebitReturn = &directDebitReturn

	return m
}

func (m *DirectDebitReturnReversalAdmissionRelationships) WithoutDirectDebitReturn() *DirectDebitReturnReversalAdmissionRelationships {
	m.DirectDebitReturn = nil
	return m
}

func (m *DirectDebitReturnReversalAdmissionRelationships) WithDirectDebitReturnReversal(directDebitReturnReversal DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversal) *DirectDebitReturnReversalAdmissionRelationships {

	m.DirectDebitReturnReversal = &directDebitReturnReversal

	return m
}

func (m *DirectDebitReturnReversalAdmissionRelationships) WithoutDirectDebitReturnReversal() *DirectDebitReturnReversalAdmissionRelationships {
	m.DirectDebitReturnReversal = nil
	return m
}

// Validate validates this direct debit return reversal admission relationships
func (m *DirectDebitReturnReversalAdmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectDebit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitReturn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitReturnReversal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnReversalAdmissionRelationships) validateDirectDebit(formats strfmt.Registry) error {

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

func (m *DirectDebitReturnReversalAdmissionRelationships) validateDirectDebitReturn(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitReturn) { // not required
		return nil
	}

	if m.DirectDebitReturn != nil {
		if err := m.DirectDebitReturn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit_return")
			}
			return err
		}
	}

	return nil
}

func (m *DirectDebitReturnReversalAdmissionRelationships) validateDirectDebitReturnReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitReturnReversal) { // not required
		return nil
	}

	if m.DirectDebitReturnReversal != nil {
		if err := m.DirectDebitReturnReversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit_return_reversal")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnReversalAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnReversalAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnReversalAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnReversalAdmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnReversalAdmissionRelationshipsDirectDebit direct debit return reversal admission relationships direct debit
// swagger:model DirectDebitReturnReversalAdmissionRelationshipsDirectDebit
type DirectDebitReturnReversalAdmissionRelationshipsDirectDebit struct {

	// data
	Data []*DirectDebit `json:"data"`
}

func DirectDebitReturnReversalAdmissionRelationshipsDirectDebitWithDefaults(defaults client.Defaults) *DirectDebitReturnReversalAdmissionRelationshipsDirectDebit {
	return &DirectDebitReturnReversalAdmissionRelationshipsDirectDebit{

		Data: make([]*DirectDebit, 0),
	}
}

func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebit) WithData(data []*DirectDebit) *DirectDebitReturnReversalAdmissionRelationshipsDirectDebit {

	m.Data = data

	return m
}

// Validate validates this direct debit return reversal admission relationships direct debit
func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebit) validateData(formats strfmt.Registry) error {

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
func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebit) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnReversalAdmissionRelationshipsDirectDebit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturn direct debit return reversal admission relationships direct debit return
// swagger:model DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturn
type DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturn struct {

	// data
	Data []*DirectDebitReturn `json:"data"`
}

func DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnWithDefaults(defaults client.Defaults) *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturn {
	return &DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturn{

		Data: make([]*DirectDebitReturn, 0),
	}
}

func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturn) WithData(data []*DirectDebitReturn) *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturn {

	m.Data = data

	return m
}

// Validate validates this direct debit return reversal admission relationships direct debit return
func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturn) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "direct_debit_return" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturn) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturn) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversal direct debit return reversal admission relationships direct debit return reversal
// swagger:model DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversal
type DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversal struct {

	// data
	Data []*DirectDebitReturnReversal `json:"data"`
}

func DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversalWithDefaults(defaults client.Defaults) *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversal {
	return &DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversal{

		Data: make([]*DirectDebitReturnReversal, 0),
	}
}

func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversal) WithData(data []*DirectDebitReturnReversal) *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversal {

	m.Data = data

	return m
}

// Validate validates this direct debit return reversal admission relationships direct debit return reversal
func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversal) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "direct_debit_return_reversal" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversal) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnReversalAdmissionRelationshipsDirectDebitReturnReversal) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
