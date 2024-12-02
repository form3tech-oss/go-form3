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

// ResourceSignature Structure containing information about cryptographic signature of the payment.
// swagger:model ResourceSignature
type ResourceSignature struct {

	// The key identifier used for signature verification.
	// Format: uuid
	KeyID strfmt.UUID `json:"key_id,omitempty"`

	// Reference information associated with the signature.
	// Max Items: 10
	References []string `json:"references"`

	// Indicates who is signing the payment.
	// Enum: ["customer"]
	Role string `json:"role,omitempty"`

	// The cryptographic signature of the payment.
	// Max Length: 300
	Signature string `json:"signature,omitempty"`
}

func ResourceSignatureWithDefaults(defaults client.Defaults) *ResourceSignature {
	return &ResourceSignature{

		KeyID: defaults.GetStrfmtUUID("ResourceSignature", "key_id"),

		References: make([]string, 0),

		Role: defaults.GetString("ResourceSignature", "role"),

		Signature: defaults.GetString("ResourceSignature", "signature"),
	}
}

func (m *ResourceSignature) WithKeyID(keyID strfmt.UUID) *ResourceSignature {

	m.KeyID = keyID

	return m
}

func (m *ResourceSignature) WithReferences(references []string) *ResourceSignature {

	m.References = references

	return m
}

func (m *ResourceSignature) WithRole(role string) *ResourceSignature {

	m.Role = role

	return m
}

func (m *ResourceSignature) WithSignature(signature string) *ResourceSignature {

	m.Signature = signature

	return m
}

// Validate validates this resource signature
func (m *ResourceSignature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferences(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceSignature) validateKeyID(formats strfmt.Registry) error {

	if swag.IsZero(m.KeyID) { // not required
		return nil
	}

	if err := validate.FormatOf("key_id", "body", "uuid", m.KeyID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ResourceSignature) validateReferences(formats strfmt.Registry) error {

	if swag.IsZero(m.References) { // not required
		return nil
	}

	iReferencesSize := int64(len(m.References))

	if err := validate.MaxItems("references", "body", iReferencesSize, 10); err != nil {
		return err
	}

	for i := 0; i < len(m.References); i++ {

		if err := validate.MaxLength("references"+"."+strconv.Itoa(i), "body", m.References[i], 50); err != nil {
			return err
		}

	}

	return nil
}

var resourceSignatureTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["customer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceSignatureTypeRolePropEnum = append(resourceSignatureTypeRolePropEnum, v)
	}
}

const (

	// ResourceSignatureRoleCustomer captures enum value "customer"
	ResourceSignatureRoleCustomer string = "customer"
)

// prop value enum
func (m *ResourceSignature) validateRoleEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resourceSignatureTypeRolePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ResourceSignature) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

func (m *ResourceSignature) validateSignature(formats strfmt.Registry) error {

	if swag.IsZero(m.Signature) { // not required
		return nil
	}

	if err := validate.MaxLength("signature", "body", m.Signature, 300); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceSignature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceSignature) UnmarshalBinary(b []byte) error {
	var res ResourceSignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ResourceSignature) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
