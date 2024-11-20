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

// RecallDecisionSubmissionCreation recall decision submission creation
// swagger:model RecallDecisionSubmissionCreation
type RecallDecisionSubmissionCreation struct {

	// data
	Data *NewRecallDecisionSubmission `json:"data,omitempty"`
}

func RecallDecisionSubmissionCreationWithDefaults(defaults client.Defaults) *RecallDecisionSubmissionCreation {
	return &RecallDecisionSubmissionCreation{

		Data: NewRecallDecisionSubmissionWithDefaults(defaults),
	}
}

func (m *RecallDecisionSubmissionCreation) WithData(data NewRecallDecisionSubmission) *RecallDecisionSubmissionCreation {

	m.Data = &data

	return m
}

func (m *RecallDecisionSubmissionCreation) WithoutData() *RecallDecisionSubmissionCreation {
	m.Data = nil
	return m
}

// Validate validates this recall decision submission creation
func (m *RecallDecisionSubmissionCreation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecallDecisionSubmissionCreation) validateData(formats strfmt.Registry) error {

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
func (m *RecallDecisionSubmissionCreation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecisionSubmissionCreation) UnmarshalBinary(b []byte) error {
	var res RecallDecisionSubmissionCreation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallDecisionSubmissionCreation) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
