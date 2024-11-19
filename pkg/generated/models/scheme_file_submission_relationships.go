// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SchemeFileSubmissionRelationships scheme file submission relationships
// swagger:model SchemeFileSubmissionRelationships
type SchemeFileSubmissionRelationships struct {

	// scheme file
	SchemeFile *SchemeFileSubmissionRelationshipsSchemeFile `json:"scheme_file,omitempty"`
}

func SchemeFileSubmissionRelationshipsWithDefaults(defaults client.Defaults) *SchemeFileSubmissionRelationships {
	return &SchemeFileSubmissionRelationships{

		SchemeFile: SchemeFileSubmissionRelationshipsSchemeFileWithDefaults(defaults),
	}
}

func (m *SchemeFileSubmissionRelationships) WithSchemeFile(schemeFile SchemeFileSubmissionRelationshipsSchemeFile) *SchemeFileSubmissionRelationships {

	m.SchemeFile = &schemeFile

	return m
}

func (m *SchemeFileSubmissionRelationships) WithoutSchemeFile() *SchemeFileSubmissionRelationships {
	m.SchemeFile = nil
	return m
}

// Validate validates this scheme file submission relationships
func (m *SchemeFileSubmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchemeFile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemeFileSubmissionRelationships) validateSchemeFile(formats strfmt.Registry) error {

	if swag.IsZero(m.SchemeFile) { // not required
		return nil
	}

	if m.SchemeFile != nil {
		if err := m.SchemeFile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheme_file")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemeFileSubmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeFileSubmissionRelationships) UnmarshalBinary(b []byte) error {
	var res SchemeFileSubmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeFileSubmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// SchemeFileSubmissionRelationshipsSchemeFile scheme file submission relationships scheme file
// swagger:model SchemeFileSubmissionRelationshipsSchemeFile
type SchemeFileSubmissionRelationshipsSchemeFile struct {

	// data
	Data []*SchemeFile `json:"data"`
}

func SchemeFileSubmissionRelationshipsSchemeFileWithDefaults(defaults client.Defaults) *SchemeFileSubmissionRelationshipsSchemeFile {
	return &SchemeFileSubmissionRelationshipsSchemeFile{

		Data: make([]*SchemeFile, 0),
	}
}

func (m *SchemeFileSubmissionRelationshipsSchemeFile) WithData(data []*SchemeFile) *SchemeFileSubmissionRelationshipsSchemeFile {

	m.Data = data

	return m
}

// Validate validates this scheme file submission relationships scheme file
func (m *SchemeFileSubmissionRelationshipsSchemeFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemeFileSubmissionRelationshipsSchemeFile) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("scheme_file" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemeFileSubmissionRelationshipsSchemeFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeFileSubmissionRelationshipsSchemeFile) UnmarshalBinary(b []byte) error {
	var res SchemeFileSubmissionRelationshipsSchemeFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeFileSubmissionRelationshipsSchemeFile) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
