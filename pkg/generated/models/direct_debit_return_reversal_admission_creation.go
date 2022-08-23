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
)

// DirectDebitReturnReversalAdmissionCreation direct debit return reversal admission creation
// swagger:model DirectDebitReturnReversalAdmissionCreation
type DirectDebitReturnReversalAdmissionCreation struct {

	// data
	Data *DirectDebitReturnReversalAdmission `json:"data,omitempty"`
}

func DirectDebitReturnReversalAdmissionCreationWithDefaults(defaults client.Defaults) *DirectDebitReturnReversalAdmissionCreation {
	return &DirectDebitReturnReversalAdmissionCreation{

		Data: DirectDebitReturnReversalAdmissionWithDefaults(defaults),
	}
}

func (m *DirectDebitReturnReversalAdmissionCreation) WithData(data DirectDebitReturnReversalAdmission) *DirectDebitReturnReversalAdmissionCreation {

	m.Data = &data

	return m
}

func (m *DirectDebitReturnReversalAdmissionCreation) WithoutData() *DirectDebitReturnReversalAdmissionCreation {
	m.Data = nil
	return m
}

// Validate validates this direct debit return reversal admission creation
func (m *DirectDebitReturnReversalAdmissionCreation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnReversalAdmissionCreation) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectDebitReturnReversalAdmissionCreation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnReversalAdmissionCreation) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnReversalAdmissionCreation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnReversalAdmissionCreation) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
