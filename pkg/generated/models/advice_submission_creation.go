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
)

// AdviceSubmissionCreation advice submission creation
// swagger:model AdviceSubmissionCreation
type AdviceSubmissionCreation struct {

	// data
	Data *NewAdviceSubmission `json:"data,omitempty"`
}

func AdviceSubmissionCreationWithDefaults(defaults client.Defaults) *AdviceSubmissionCreation {
	return &AdviceSubmissionCreation{

		Data: NewAdviceSubmissionWithDefaults(defaults),
	}
}

func (m *AdviceSubmissionCreation) WithData(data NewAdviceSubmission) *AdviceSubmissionCreation {

	m.Data = &data

	return m
}

func (m *AdviceSubmissionCreation) WithoutData() *AdviceSubmissionCreation {
	m.Data = nil
	return m
}

// Validate validates this advice submission creation
func (m *AdviceSubmissionCreation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdviceSubmissionCreation) validateData(formats strfmt.Registry) error {

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
func (m *AdviceSubmissionCreation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdviceSubmissionCreation) UnmarshalBinary(b []byte) error {
	var res AdviceSubmissionCreation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AdviceSubmissionCreation) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
