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
	"github.com/go-openapi/validate"
)

// RecallAdmissionAttributes recall admission attributes
// swagger:model RecallAdmissionAttributes
type RecallAdmissionAttributes struct {

	// Date and time the recall admission was created
	// Read Only: true
	// Format: date-time
	AdmissionDatetime *strfmt.DateTime `json:"admission_datetime,omitempty"`

	// Additional payment reference assigned by the scheme
	ReferenceID string `json:"reference_id,omitempty"`

	// source gateway
	SourceGateway string `json:"source_gateway,omitempty"`

	// status
	Status RecallAdmissionStatus `json:"status,omitempty"`

	// Human-readable reason for failure if admission status is failed
	StatusReason string `json:"status_reason,omitempty"`
}

func RecallAdmissionAttributesWithDefaults(defaults client.Defaults) *RecallAdmissionAttributes {
	return &RecallAdmissionAttributes{

		AdmissionDatetime: defaults.GetStrfmtDateTimePtr("RecallAdmissionAttributes", "admission_datetime"),

		ReferenceID: defaults.GetString("RecallAdmissionAttributes", "reference_id"),

		SourceGateway: defaults.GetString("RecallAdmissionAttributes", "source_gateway"),

		// TODO Status: RecallAdmissionStatus,

		StatusReason: defaults.GetString("RecallAdmissionAttributes", "status_reason"),
	}
}

func (m *RecallAdmissionAttributes) WithAdmissionDatetime(admissionDatetime strfmt.DateTime) *RecallAdmissionAttributes {

	m.AdmissionDatetime = &admissionDatetime

	return m
}

func (m *RecallAdmissionAttributes) WithoutAdmissionDatetime() *RecallAdmissionAttributes {
	m.AdmissionDatetime = nil
	return m
}

func (m *RecallAdmissionAttributes) WithReferenceID(referenceID string) *RecallAdmissionAttributes {

	m.ReferenceID = referenceID

	return m
}

func (m *RecallAdmissionAttributes) WithSourceGateway(sourceGateway string) *RecallAdmissionAttributes {

	m.SourceGateway = sourceGateway

	return m
}

func (m *RecallAdmissionAttributes) WithStatus(status RecallAdmissionStatus) *RecallAdmissionAttributes {

	m.Status = status

	return m
}

func (m *RecallAdmissionAttributes) WithStatusReason(statusReason string) *RecallAdmissionAttributes {

	m.StatusReason = statusReason

	return m
}

// Validate validates this recall admission attributes
func (m *RecallAdmissionAttributes) Validate(formats strfmt.Registry) error {
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

func (m *RecallAdmissionAttributes) validateAdmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.AdmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("admission_datetime", "body", "date-time", m.AdmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RecallAdmissionAttributes) validateStatus(formats strfmt.Registry) error {

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
func (m *RecallAdmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallAdmissionAttributes) UnmarshalBinary(b []byte) error {
	var res RecallAdmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallAdmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
