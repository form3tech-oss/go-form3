// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DirectDebitReturnReversal direct debit return reversal
// swagger:model DirectDebitReturnReversal
type DirectDebitReturnReversal struct {

	// attributes
	Attributes interface{} `json:"attributes,omitempty"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// modified on
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

	// organisation id
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// relationships
	Relationships *DirectDebitReturnReversalRelationships `json:"relationships,omitempty"`

	// type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`

	// version
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

// line 140

func DirectDebitReturnReversalWithDefaults(defaults client.Defaults) *DirectDebitReturnReversal {
	return &DirectDebitReturnReversal{

		// TODO Attributes: interface{},

		CreatedOn: defaults.GetStrfmtDateTimePtr("DirectDebitReturnReversal", "created_on"),

		ID: defaults.GetStrfmtUUID("DirectDebitReturnReversal", "id"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("DirectDebitReturnReversal", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("DirectDebitReturnReversal", "organisation_id"),

		Relationships: DirectDebitReturnReversalRelationshipsWithDefaults(defaults),

		Type: defaults.GetString("DirectDebitReturnReversal", "type"),

		Version: defaults.GetInt64Ptr("DirectDebitReturnReversal", "version"),
	}
}

func (m *DirectDebitReturnReversal) WithAttributes(attributes interface{}) *DirectDebitReturnReversal {

	m.Attributes = attributes

	return m
}

func (m *DirectDebitReturnReversal) WithCreatedOn(createdOn strfmt.DateTime) *DirectDebitReturnReversal {

	m.CreatedOn = &createdOn

	return m
}

func (m *DirectDebitReturnReversal) WithoutCreatedOn() *DirectDebitReturnReversal {
	m.CreatedOn = nil
	return m
}

func (m *DirectDebitReturnReversal) WithID(id strfmt.UUID) *DirectDebitReturnReversal {

	m.ID = id

	return m
}

func (m *DirectDebitReturnReversal) WithModifiedOn(modifiedOn strfmt.DateTime) *DirectDebitReturnReversal {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *DirectDebitReturnReversal) WithoutModifiedOn() *DirectDebitReturnReversal {
	m.ModifiedOn = nil
	return m
}

func (m *DirectDebitReturnReversal) WithOrganisationID(organisationID strfmt.UUID) *DirectDebitReturnReversal {

	m.OrganisationID = organisationID

	return m
}

func (m *DirectDebitReturnReversal) WithRelationships(relationships DirectDebitReturnReversalRelationships) *DirectDebitReturnReversal {

	m.Relationships = &relationships

	return m
}

func (m *DirectDebitReturnReversal) WithoutRelationships() *DirectDebitReturnReversal {
	m.Relationships = nil
	return m
}

func (m *DirectDebitReturnReversal) WithType(typeVar string) *DirectDebitReturnReversal {

	m.Type = typeVar

	return m
}

func (m *DirectDebitReturnReversal) WithVersion(version int64) *DirectDebitReturnReversal {

	m.Version = &version

	return m
}

func (m *DirectDebitReturnReversal) WithoutVersion() *DirectDebitReturnReversal {
	m.Version = nil
	return m
}

// Validate validates this direct debit return reversal
func (m *DirectDebitReturnReversal) Validate(formats strfmt.Registry) error {
	var res []error

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

func (m *DirectDebitReturnReversal) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnReversal) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnReversal) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnReversal) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnReversal) validateRelationships(formats strfmt.Registry) error {

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

func (m *DirectDebitReturnReversal) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *DirectDebitReturnReversal) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnReversal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnReversal) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnReversal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnReversal) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnReversalRelationships direct debit return reversal relationships
// swagger:model DirectDebitReturnReversalRelationships
type DirectDebitReturnReversalRelationships struct {

	// direct debit
	DirectDebit *DirectDebitReturnReversalRelationshipsDirectDebit `json:"direct_debit,omitempty"`

	// direct debit return
	DirectDebitReturn *DirectDebitReturnReversalRelationshipsDirectDebitReturn `json:"direct_debit_return,omitempty"`

	// direct debit return reversal admission
	DirectDebitReturnReversalAdmission *DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmission `json:"direct_debit_return_reversal_admission,omitempty"`
}

// line 140

func DirectDebitReturnReversalRelationshipsWithDefaults(defaults client.Defaults) *DirectDebitReturnReversalRelationships {
	return &DirectDebitReturnReversalRelationships{

		DirectDebit: DirectDebitReturnReversalRelationshipsDirectDebitWithDefaults(defaults),

		DirectDebitReturn: DirectDebitReturnReversalRelationshipsDirectDebitReturnWithDefaults(defaults),

		DirectDebitReturnReversalAdmission: DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmissionWithDefaults(defaults),
	}
}

func (m *DirectDebitReturnReversalRelationships) WithDirectDebit(directDebit DirectDebitReturnReversalRelationshipsDirectDebit) *DirectDebitReturnReversalRelationships {

	m.DirectDebit = &directDebit

	return m
}

func (m *DirectDebitReturnReversalRelationships) WithoutDirectDebit() *DirectDebitReturnReversalRelationships {
	m.DirectDebit = nil
	return m
}

func (m *DirectDebitReturnReversalRelationships) WithDirectDebitReturn(directDebitReturn DirectDebitReturnReversalRelationshipsDirectDebitReturn) *DirectDebitReturnReversalRelationships {

	m.DirectDebitReturn = &directDebitReturn

	return m
}

func (m *DirectDebitReturnReversalRelationships) WithoutDirectDebitReturn() *DirectDebitReturnReversalRelationships {
	m.DirectDebitReturn = nil
	return m
}

func (m *DirectDebitReturnReversalRelationships) WithDirectDebitReturnReversalAdmission(directDebitReturnReversalAdmission DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmission) *DirectDebitReturnReversalRelationships {

	m.DirectDebitReturnReversalAdmission = &directDebitReturnReversalAdmission

	return m
}

func (m *DirectDebitReturnReversalRelationships) WithoutDirectDebitReturnReversalAdmission() *DirectDebitReturnReversalRelationships {
	m.DirectDebitReturnReversalAdmission = nil
	return m
}

// Validate validates this direct debit return reversal relationships
func (m *DirectDebitReturnReversalRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectDebit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitReturn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirectDebitReturnReversalAdmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnReversalRelationships) validateDirectDebit(formats strfmt.Registry) error {

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

func (m *DirectDebitReturnReversalRelationships) validateDirectDebitReturn(formats strfmt.Registry) error {

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

func (m *DirectDebitReturnReversalRelationships) validateDirectDebitReturnReversalAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.DirectDebitReturnReversalAdmission) { // not required
		return nil
	}

	if m.DirectDebitReturnReversalAdmission != nil {
		if err := m.DirectDebitReturnReversalAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "direct_debit_return_reversal_admission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnReversalRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnReversalRelationships) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnReversalRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnReversalRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnReversalRelationshipsDirectDebit direct debit return reversal relationships direct debit
// swagger:model DirectDebitReturnReversalRelationshipsDirectDebit
type DirectDebitReturnReversalRelationshipsDirectDebit struct {

	// data
	Data []*DirectDebit `json:"data"`
}

// line 140

func DirectDebitReturnReversalRelationshipsDirectDebitWithDefaults(defaults client.Defaults) *DirectDebitReturnReversalRelationshipsDirectDebit {
	return &DirectDebitReturnReversalRelationshipsDirectDebit{

		Data: make([]*DirectDebit, 0),
	}
}

func (m *DirectDebitReturnReversalRelationshipsDirectDebit) WithData(data []*DirectDebit) *DirectDebitReturnReversalRelationshipsDirectDebit {

	m.Data = data

	return m
}

// Validate validates this direct debit return reversal relationships direct debit
func (m *DirectDebitReturnReversalRelationshipsDirectDebit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnReversalRelationshipsDirectDebit) validateData(formats strfmt.Registry) error {

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
func (m *DirectDebitReturnReversalRelationshipsDirectDebit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnReversalRelationshipsDirectDebit) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnReversalRelationshipsDirectDebit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnReversalRelationshipsDirectDebit) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnReversalRelationshipsDirectDebitReturn direct debit return reversal relationships direct debit return
// swagger:model DirectDebitReturnReversalRelationshipsDirectDebitReturn
type DirectDebitReturnReversalRelationshipsDirectDebitReturn struct {

	// data
	Data []*DirectDebitReturn `json:"data"`
}

// line 140

func DirectDebitReturnReversalRelationshipsDirectDebitReturnWithDefaults(defaults client.Defaults) *DirectDebitReturnReversalRelationshipsDirectDebitReturn {
	return &DirectDebitReturnReversalRelationshipsDirectDebitReturn{

		Data: make([]*DirectDebitReturn, 0),
	}
}

func (m *DirectDebitReturnReversalRelationshipsDirectDebitReturn) WithData(data []*DirectDebitReturn) *DirectDebitReturnReversalRelationshipsDirectDebitReturn {

	m.Data = data

	return m
}

// Validate validates this direct debit return reversal relationships direct debit return
func (m *DirectDebitReturnReversalRelationshipsDirectDebitReturn) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnReversalRelationshipsDirectDebitReturn) validateData(formats strfmt.Registry) error {

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
func (m *DirectDebitReturnReversalRelationshipsDirectDebitReturn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnReversalRelationshipsDirectDebitReturn) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnReversalRelationshipsDirectDebitReturn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnReversalRelationshipsDirectDebitReturn) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmission direct debit return reversal relationships direct debit return reversal admission
// swagger:model DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmission
type DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmission struct {

	// data
	Data []*DirectDebitReturnReversalAdmission `json:"data"`
}

// line 140

func DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmissionWithDefaults(defaults client.Defaults) *DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmission {
	return &DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmission{

		Data: make([]*DirectDebitReturnReversalAdmission, 0),
	}
}

func (m *DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmission) WithData(data []*DirectDebitReturnReversalAdmission) *DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmission {

	m.Data = data

	return m
}

// Validate validates this direct debit return reversal relationships direct debit return reversal admission
func (m *DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmission) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("relationships" + "." + "direct_debit_return_reversal_admission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmission) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnReversalRelationshipsDirectDebitReturnReversalAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
