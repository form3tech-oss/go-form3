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

// DirectDebitReturnReversalCreationResponse direct debit return reversal creation response
// swagger:model DirectDebitReturnReversalCreationResponse
type DirectDebitReturnReversalCreationResponse struct {

	// data
	Data *DirectDebitReturnReversal `json:"data,omitempty"`
}

func DirectDebitReturnReversalCreationResponseWithDefaults(defaults client.Defaults) *DirectDebitReturnReversalCreationResponse {
	return &DirectDebitReturnReversalCreationResponse{

		Data: DirectDebitReturnReversalWithDefaults(defaults),
	}
}

func (m *DirectDebitReturnReversalCreationResponse) WithData(data DirectDebitReturnReversal) *DirectDebitReturnReversalCreationResponse {

	m.Data = &data

	return m
}

func (m *DirectDebitReturnReversalCreationResponse) WithoutData() *DirectDebitReturnReversalCreationResponse {
	m.Data = nil
	return m
}

// Validate validates this direct debit return reversal creation response
func (m *DirectDebitReturnReversalCreationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitReturnReversalCreationResponse) validateData(formats strfmt.Registry) error {

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
func (m *DirectDebitReturnReversalCreationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectDebitReturnReversalCreationResponse) UnmarshalBinary(b []byte) error {
	var res DirectDebitReturnReversalCreationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *DirectDebitReturnReversalCreationResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
