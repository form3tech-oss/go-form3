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
	"github.com/go-openapi/validate"
)

// SchemeMessageAdmissionAttributes scheme message admission attributes
// swagger:model SchemeMessageAdmissionAttributes
type SchemeMessageAdmissionAttributes struct {

	// admission datetime
	// Format: date-time
	AdmissionDatetime strfmt.DateTime `json:"admission_datetime,omitempty"`

	// scheme status code
	SchemeStatusCode string `json:"scheme_status_code,omitempty"`

	// scheme status code description
	SchemeStatusCodeDescription string `json:"scheme_status_code_description,omitempty"`

	// status
	Status SchemeMessageAdmissionStatus `json:"status,omitempty"`
}

func SchemeMessageAdmissionAttributesWithDefaults(defaults client.Defaults) *SchemeMessageAdmissionAttributes {
	return &SchemeMessageAdmissionAttributes{

		AdmissionDatetime: defaults.GetStrfmtDateTime("SchemeMessageAdmissionAttributes", "admission_datetime"),

		SchemeStatusCode: defaults.GetString("SchemeMessageAdmissionAttributes", "scheme_status_code"),

		SchemeStatusCodeDescription: defaults.GetString("SchemeMessageAdmissionAttributes", "scheme_status_code_description"),

		// TODO Status: SchemeMessageAdmissionStatus,

	}
}

func (m *SchemeMessageAdmissionAttributes) WithAdmissionDatetime(admissionDatetime strfmt.DateTime) *SchemeMessageAdmissionAttributes {

	m.AdmissionDatetime = admissionDatetime

	return m
}

func (m *SchemeMessageAdmissionAttributes) WithSchemeStatusCode(schemeStatusCode string) *SchemeMessageAdmissionAttributes {

	m.SchemeStatusCode = schemeStatusCode

	return m
}

func (m *SchemeMessageAdmissionAttributes) WithSchemeStatusCodeDescription(schemeStatusCodeDescription string) *SchemeMessageAdmissionAttributes {

	m.SchemeStatusCodeDescription = schemeStatusCodeDescription

	return m
}

func (m *SchemeMessageAdmissionAttributes) WithStatus(status SchemeMessageAdmissionStatus) *SchemeMessageAdmissionAttributes {

	m.Status = status

	return m
}

// Validate validates this scheme message admission attributes
func (m *SchemeMessageAdmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemeMessageAdmissionAttributes) validateAdmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.AdmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("admission_datetime", "body", "date-time", m.AdmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SchemeMessageAdmissionAttributes) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemeMessageAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeMessageAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res SchemeMessageAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeMessageAdmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
