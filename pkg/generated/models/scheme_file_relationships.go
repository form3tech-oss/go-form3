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

// SchemeFileRelationships scheme file relationships
// swagger:model SchemeFileRelationships
type SchemeFileRelationships struct {

	// scheme file submission
	SchemeFileSubmission *SchemeFileSubmission `json:"scheme_file_submission,omitempty"`
}

func SchemeFileRelationshipsWithDefaults(defaults client.Defaults) *SchemeFileRelationships {
	return &SchemeFileRelationships{

		SchemeFileSubmission: SchemeFileSubmissionWithDefaults(defaults),
	}
}

func (m *SchemeFileRelationships) WithSchemeFileSubmission(schemeFileSubmission SchemeFileSubmission) *SchemeFileRelationships {

	m.SchemeFileSubmission = &schemeFileSubmission

	return m
}

func (m *SchemeFileRelationships) WithoutSchemeFileSubmission() *SchemeFileRelationships {
	m.SchemeFileSubmission = nil
	return m
}

// Validate validates this scheme file relationships
func (m *SchemeFileRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchemeFileSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemeFileRelationships) validateSchemeFileSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.SchemeFileSubmission) { // not required
		return nil
	}

	if m.SchemeFileSubmission != nil {
		if err := m.SchemeFileSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheme_file_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemeFileRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeFileRelationships) UnmarshalBinary(b []byte) error {
	var res SchemeFileRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeFileRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
