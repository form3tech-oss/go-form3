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
)

// RecallDecisionCreationResponse recall decision creation response
// swagger:model RecallDecisionCreationResponse
type RecallDecisionCreationResponse struct {

	// data
	Data *RecallDecision `json:"data,omitempty"`
}

func RecallDecisionCreationResponseWithDefaults(defaults client.Defaults) *RecallDecisionCreationResponse {
	return &RecallDecisionCreationResponse{

		Data: RecallDecisionWithDefaults(defaults),
	}
}

func (m *RecallDecisionCreationResponse) WithData(data RecallDecision) *RecallDecisionCreationResponse {

	m.Data = &data

	return m
}

func (m *RecallDecisionCreationResponse) WithoutData() *RecallDecisionCreationResponse {
	m.Data = nil
	return m
}

// Validate validates this recall decision creation response
func (m *RecallDecisionCreationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecallDecisionCreationResponse) validateData(formats strfmt.Registry) error {

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
func (m *RecallDecisionCreationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecisionCreationResponse) UnmarshalBinary(b []byte) error {
	var res RecallDecisionCreationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallDecisionCreationResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
