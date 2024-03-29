// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReturnSubmissionTaskAmendment return submission task amendment
// swagger:model ReturnSubmissionTaskAmendment
type ReturnSubmissionTaskAmendment struct {

	// data
	Data *ReturnSubmissionTaskUpdate `json:"data,omitempty"`
}

func ReturnSubmissionTaskAmendmentWithDefaults(defaults client.Defaults) *ReturnSubmissionTaskAmendment {
	return &ReturnSubmissionTaskAmendment{

		Data: ReturnSubmissionTaskUpdateWithDefaults(defaults),
	}
}

func (m *ReturnSubmissionTaskAmendment) WithData(data ReturnSubmissionTaskUpdate) *ReturnSubmissionTaskAmendment {

	m.Data = &data

	return m
}

func (m *ReturnSubmissionTaskAmendment) WithoutData() *ReturnSubmissionTaskAmendment {
	m.Data = nil
	return m
}

// Validate validates this return submission task amendment
func (m *ReturnSubmissionTaskAmendment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnSubmissionTaskAmendment) validateData(formats strfmt.Registry) error {

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
func (m *ReturnSubmissionTaskAmendment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnSubmissionTaskAmendment) UnmarshalBinary(b []byte) error {
	var res ReturnSubmissionTaskAmendment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnSubmissionTaskAmendment) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
