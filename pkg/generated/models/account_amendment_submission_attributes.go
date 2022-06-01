// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v4/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountAmendmentSubmissionAttributes account amendment submission attributes
// swagger:model AccountAmendmentSubmissionAttributes
type AccountAmendmentSubmissionAttributes struct {

	// status
	Status SubmissionStatus `json:"status,omitempty"`

	// Description of the submission status
	StatusReason string `json:"status_reason,omitempty"`

	// submission datetime
	// Read Only: true
	// Format: date-time
	SubmissionDatetime *strfmt.DateTime `json:"submission_datetime,omitempty"`
}

func AccountAmendmentSubmissionAttributesWithDefaults(defaults client.Defaults) *AccountAmendmentSubmissionAttributes {
	return &AccountAmendmentSubmissionAttributes{

		// TODO Status: SubmissionStatus,

		StatusReason: defaults.GetString("AccountAmendmentSubmissionAttributes", "status_reason"),

		SubmissionDatetime: defaults.GetStrfmtDateTimePtr("AccountAmendmentSubmissionAttributes", "submission_datetime"),
	}
}

func (m *AccountAmendmentSubmissionAttributes) WithStatus(status SubmissionStatus) *AccountAmendmentSubmissionAttributes {

	m.Status = status

	return m
}

func (m *AccountAmendmentSubmissionAttributes) WithStatusReason(statusReason string) *AccountAmendmentSubmissionAttributes {

	m.StatusReason = statusReason

	return m
}

func (m *AccountAmendmentSubmissionAttributes) WithSubmissionDatetime(submissionDatetime strfmt.DateTime) *AccountAmendmentSubmissionAttributes {

	m.SubmissionDatetime = &submissionDatetime

	return m
}

func (m *AccountAmendmentSubmissionAttributes) WithoutSubmissionDatetime() *AccountAmendmentSubmissionAttributes {
	m.SubmissionDatetime = nil
	return m
}

// Validate validates this account amendment submission attributes
func (m *AccountAmendmentSubmissionAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmissionDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountAmendmentSubmissionAttributes) validateStatus(formats strfmt.Registry) error {

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

func (m *AccountAmendmentSubmissionAttributes) validateSubmissionDatetime(formats strfmt.Registry) error {

	if swag.IsZero(m.SubmissionDatetime) { // not required
		return nil
	}

	if err := validate.FormatOf("submission_datetime", "body", "date-time", m.SubmissionDatetime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountAmendmentSubmissionAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountAmendmentSubmissionAttributes) UnmarshalBinary(b []byte) error {
	var res AccountAmendmentSubmissionAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AccountAmendmentSubmissionAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
