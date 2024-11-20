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

// PaymentAdmissionBeneficiaryBranch payment admission beneficiary branch
// swagger:model PaymentAdmissionBeneficiaryBranch
type PaymentAdmissionBeneficiaryBranch struct {

	// attributes
	Attributes *PaymentAdmissionBeneficiaryBranchAttributes `json:"attributes,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Name of the resource type
	// Pattern: ^[A-Za-z_]*$
	Type string `json:"type,omitempty"`
}

func PaymentAdmissionBeneficiaryBranchWithDefaults(defaults client.Defaults) *PaymentAdmissionBeneficiaryBranch {
	return &PaymentAdmissionBeneficiaryBranch{

		Attributes: PaymentAdmissionBeneficiaryBranchAttributesWithDefaults(defaults),

		ID: defaults.GetStrfmtUUID("PaymentAdmissionBeneficiaryBranch", "id"),

		Type: defaults.GetString("PaymentAdmissionBeneficiaryBranch", "type"),
	}
}

func (m *PaymentAdmissionBeneficiaryBranch) WithAttributes(attributes PaymentAdmissionBeneficiaryBranchAttributes) *PaymentAdmissionBeneficiaryBranch {

	m.Attributes = &attributes

	return m
}

func (m *PaymentAdmissionBeneficiaryBranch) WithoutAttributes() *PaymentAdmissionBeneficiaryBranch {
	m.Attributes = nil
	return m
}

func (m *PaymentAdmissionBeneficiaryBranch) WithID(id strfmt.UUID) *PaymentAdmissionBeneficiaryBranch {

	m.ID = id

	return m
}

func (m *PaymentAdmissionBeneficiaryBranch) WithType(typeVar string) *PaymentAdmissionBeneficiaryBranch {

	m.Type = typeVar

	return m
}

// Validate validates this payment admission beneficiary branch
func (m *PaymentAdmissionBeneficiaryBranch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAdmissionBeneficiaryBranch) validateAttributes(formats strfmt.Registry) error {

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

func (m *PaymentAdmissionBeneficiaryBranch) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PaymentAdmissionBeneficiaryBranch) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := validate.Pattern("type", "body", string(m.Type), `^[A-Za-z_]*$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAdmissionBeneficiaryBranch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdmissionBeneficiaryBranch) UnmarshalBinary(b []byte) error {
	var res PaymentAdmissionBeneficiaryBranch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdmissionBeneficiaryBranch) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// PaymentAdmissionBeneficiaryBranchAttributes payment admission beneficiary branch attributes
// swagger:model PaymentAdmissionBeneficiaryBranchAttributes
type PaymentAdmissionBeneficiaryBranchAttributes struct {

	// All purpose list of key-value pairs specific data stored on the associated beneficiary branch.
	// Max Items: 5
	UserDefinedData []*UserDefinedData `json:"user_defined_data"`
}

func PaymentAdmissionBeneficiaryBranchAttributesWithDefaults(defaults client.Defaults) *PaymentAdmissionBeneficiaryBranchAttributes {
	return &PaymentAdmissionBeneficiaryBranchAttributes{

		UserDefinedData: make([]*UserDefinedData, 0),
	}
}

func (m *PaymentAdmissionBeneficiaryBranchAttributes) WithUserDefinedData(userDefinedData []*UserDefinedData) *PaymentAdmissionBeneficiaryBranchAttributes {

	m.UserDefinedData = userDefinedData

	return m
}

// Validate validates this payment admission beneficiary branch attributes
func (m *PaymentAdmissionBeneficiaryBranchAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserDefinedData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAdmissionBeneficiaryBranchAttributes) validateUserDefinedData(formats strfmt.Registry) error {

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
func (m *PaymentAdmissionBeneficiaryBranchAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdmissionBeneficiaryBranchAttributes) UnmarshalBinary(b []byte) error {
	var res PaymentAdmissionBeneficiaryBranchAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdmissionBeneficiaryBranchAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
