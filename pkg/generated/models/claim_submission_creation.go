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

// ClaimSubmissionCreation claim submission creation
// swagger:model ClaimSubmissionCreation
type ClaimSubmissionCreation struct {

	// data
	// Required: true
	Data *ClaimSubmission `json:"data"`
}

func ClaimSubmissionCreationWithDefaults(defaults client.Defaults) *ClaimSubmissionCreation {
	return &ClaimSubmissionCreation{

		Data: ClaimSubmissionWithDefaults(defaults),
	}
}

func (m *ClaimSubmissionCreation) WithData(data ClaimSubmission) *ClaimSubmissionCreation {

	m.Data = &data

	return m
}

func (m *ClaimSubmissionCreation) WithoutData() *ClaimSubmissionCreation {
	m.Data = nil
	return m
}

// Validate validates this claim submission creation
func (m *ClaimSubmissionCreation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClaimSubmissionCreation) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
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
func (m *ClaimSubmissionCreation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClaimSubmissionCreation) UnmarshalBinary(b []byte) error {
	var res ClaimSubmissionCreation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ClaimSubmissionCreation) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
