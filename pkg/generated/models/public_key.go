// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PublicKey public key
// swagger:model PublicKey
type PublicKey struct {

	// attributes
	Attributes *PublicKeyAttributes `json:"attributes,omitempty"`

	// created on
	// Format: date-time
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// Unique resource ID
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// modified by
	// Format: uuid
	ModifiedBy *strfmt.UUID `json:"modified_by,omitempty"`

	// modified on
	// Format: date-time
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Format: uuid
	OrganisationID strfmt.UUID `json:"organisation_id,omitempty"`

	// Name of the resource type
	Type string `json:"type,omitempty"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`
}

func PublicKeyWithDefaults(defaults client.Defaults) *PublicKey {
	return &PublicKey{

		Attributes: PublicKeyAttributesWithDefaults(defaults),

		CreatedOn: defaults.GetStrfmtDateTimePtr("PublicKey", "created_on"),

		ID: defaults.GetStrfmtUUID("PublicKey", "id"),

		ModifiedBy: defaults.GetStrfmtUUIDPtr("PublicKey", "modified_by"),

		ModifiedOn: defaults.GetStrfmtDateTimePtr("PublicKey", "modified_on"),

		OrganisationID: defaults.GetStrfmtUUID("PublicKey", "organisation_id"),

		Type: defaults.GetString("PublicKey", "type"),

		Version: defaults.GetInt64Ptr("PublicKey", "version"),
	}
}

func (m *PublicKey) WithAttributes(attributes PublicKeyAttributes) *PublicKey {

	m.Attributes = &attributes

	return m
}

func (m *PublicKey) WithoutAttributes() *PublicKey {
	m.Attributes = nil
	return m
}

func (m *PublicKey) WithCreatedOn(createdOn strfmt.DateTime) *PublicKey {

	m.CreatedOn = &createdOn

	return m
}

func (m *PublicKey) WithoutCreatedOn() *PublicKey {
	m.CreatedOn = nil
	return m
}

func (m *PublicKey) WithID(id strfmt.UUID) *PublicKey {

	m.ID = id

	return m
}

func (m *PublicKey) WithModifiedBy(modifiedBy strfmt.UUID) *PublicKey {

	m.ModifiedBy = &modifiedBy

	return m
}

func (m *PublicKey) WithoutModifiedBy() *PublicKey {
	m.ModifiedBy = nil
	return m
}

func (m *PublicKey) WithModifiedOn(modifiedOn strfmt.DateTime) *PublicKey {

	m.ModifiedOn = &modifiedOn

	return m
}

func (m *PublicKey) WithoutModifiedOn() *PublicKey {
	m.ModifiedOn = nil
	return m
}

func (m *PublicKey) WithOrganisationID(organisationID strfmt.UUID) *PublicKey {

	m.OrganisationID = organisationID

	return m
}

func (m *PublicKey) WithType(typeVar string) *PublicKey {

	m.Type = typeVar

	return m
}

func (m *PublicKey) WithVersion(version int64) *PublicKey {

	m.Version = &version

	return m
}

func (m *PublicKey) WithoutVersion() *PublicKey {
	m.Version = nil
	return m
}

// Validate validates this public key
func (m *PublicKey) Validate(formats strfmt.Registry) error {
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

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
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

func (m *PublicKey) validateAttributes(formats strfmt.Registry) error {

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

func (m *PublicKey) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PublicKey) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PublicKey) validateModifiedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_by", "body", "uuid", m.ModifiedBy.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PublicKey) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PublicKey) validateOrganisationID(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganisationID) { // not required
		return nil
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PublicKey) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicKey) UnmarshalBinary(b []byte) error {
	var res PublicKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PublicKey) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PublicKeyAttributes public key attributes
// swagger:model PublicKeyAttributes
type PublicKeyAttributes struct {

	// expires on
	// Format: date-time
	ExpiresOn strfmt.DateTime `json:"expires_on,omitempty"`

	// fingerprint
	Fingerprint string `json:"fingerprint,omitempty"`

	// public key
	PublicKey string `json:"public_key,omitempty"`
}

func PublicKeyAttributesWithDefaults(defaults client.Defaults) *PublicKeyAttributes {
	return &PublicKeyAttributes{

		ExpiresOn: defaults.GetStrfmtDateTime("PublicKeyAttributes", "expires_on"),

		Fingerprint: defaults.GetString("PublicKeyAttributes", "fingerprint"),

		PublicKey: defaults.GetString("PublicKeyAttributes", "public_key"),
	}
}

func (m *PublicKeyAttributes) WithExpiresOn(expiresOn strfmt.DateTime) *PublicKeyAttributes {

	m.ExpiresOn = expiresOn

	return m
}

func (m *PublicKeyAttributes) WithFingerprint(fingerprint string) *PublicKeyAttributes {

	m.Fingerprint = fingerprint

	return m
}

func (m *PublicKeyAttributes) WithPublicKey(publicKey string) *PublicKeyAttributes {

	m.PublicKey = publicKey

	return m
}

// Validate validates this public key attributes
func (m *PublicKeyAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiresOn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicKeyAttributes) validateExpiresOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpiresOn) { // not required
		return nil
	}

	if err := validate.FormatOf("attributes"+"."+"expires_on", "body", "date-time", m.ExpiresOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicKeyAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicKeyAttributes) UnmarshalBinary(b []byte) error {
	var res PublicKeyAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PublicKeyAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
