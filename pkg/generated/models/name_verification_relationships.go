// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NameVerificationRelationships name verification relationships
// swagger:model NameVerificationRelationships
type NameVerificationRelationships struct {

	// name verification admission
	NameVerificationAdmission *NameVerificationRelationshipsNameVerificationAdmission `json:"name_verification_admission,omitempty"`

	// name verification submission
	NameVerificationSubmission *NameVerificationRelationshipsNameVerificationSubmission `json:"name_verification_submission,omitempty"`
}

func NameVerificationRelationshipsWithDefaults(defaults client.Defaults) *NameVerificationRelationships {
	return &NameVerificationRelationships{

		NameVerificationAdmission: NameVerificationRelationshipsNameVerificationAdmissionWithDefaults(defaults),

		NameVerificationSubmission: NameVerificationRelationshipsNameVerificationSubmissionWithDefaults(defaults),
	}
}

func (m *NameVerificationRelationships) WithNameVerificationAdmission(nameVerificationAdmission NameVerificationRelationshipsNameVerificationAdmission) *NameVerificationRelationships {

	m.NameVerificationAdmission = &nameVerificationAdmission

	return m
}

func (m *NameVerificationRelationships) WithoutNameVerificationAdmission() *NameVerificationRelationships {
	m.NameVerificationAdmission = nil
	return m
}

func (m *NameVerificationRelationships) WithNameVerificationSubmission(nameVerificationSubmission NameVerificationRelationshipsNameVerificationSubmission) *NameVerificationRelationships {

	m.NameVerificationSubmission = &nameVerificationSubmission

	return m
}

func (m *NameVerificationRelationships) WithoutNameVerificationSubmission() *NameVerificationRelationships {
	m.NameVerificationSubmission = nil
	return m
}

// Validate validates this name verification relationships
func (m *NameVerificationRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNameVerificationAdmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameVerificationSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameVerificationRelationships) validateNameVerificationAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.NameVerificationAdmission) { // not required
		return nil
	}

	if m.NameVerificationAdmission != nil {
		if err := m.NameVerificationAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name_verification_admission")
			}
			return err
		}
	}

	return nil
}

func (m *NameVerificationRelationships) validateNameVerificationSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.NameVerificationSubmission) { // not required
		return nil
	}

	if m.NameVerificationSubmission != nil {
		if err := m.NameVerificationSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name_verification_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NameVerificationRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameVerificationRelationships) UnmarshalBinary(b []byte) error {
	var res NameVerificationRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NameVerificationRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// NameVerificationRelationshipsNameVerificationAdmission name verification relationships name verification admission
// swagger:model NameVerificationRelationshipsNameVerificationAdmission
type NameVerificationRelationshipsNameVerificationAdmission struct {

	// data
	// Max Items: 1
	// Min Items: 0
	Data []*NameVerificationAdmission `json:"data"`
}

func NameVerificationRelationshipsNameVerificationAdmissionWithDefaults(defaults client.Defaults) *NameVerificationRelationshipsNameVerificationAdmission {
	return &NameVerificationRelationshipsNameVerificationAdmission{

		Data: make([]*NameVerificationAdmission, 0),
	}
}

func (m *NameVerificationRelationshipsNameVerificationAdmission) WithData(data []*NameVerificationAdmission) *NameVerificationRelationshipsNameVerificationAdmission {

	m.Data = data

	return m
}

// Validate validates this name verification relationships name verification admission
func (m *NameVerificationRelationshipsNameVerificationAdmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameVerificationRelationshipsNameVerificationAdmission) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	iDataSize := int64(len(m.Data))

	if err := validate.MinItems("name_verification_admission"+"."+"data", "body", iDataSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("name_verification_admission"+"."+"data", "body", iDataSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("name_verification_admission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NameVerificationRelationshipsNameVerificationAdmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameVerificationRelationshipsNameVerificationAdmission) UnmarshalBinary(b []byte) error {
	var res NameVerificationRelationshipsNameVerificationAdmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NameVerificationRelationshipsNameVerificationAdmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// NameVerificationRelationshipsNameVerificationSubmission name verification relationships name verification submission
// swagger:model NameVerificationRelationshipsNameVerificationSubmission
type NameVerificationRelationshipsNameVerificationSubmission struct {

	// data
	// Max Items: 1
	// Min Items: 0
	Data []*NameVerificationSubmission `json:"data"`
}

func NameVerificationRelationshipsNameVerificationSubmissionWithDefaults(defaults client.Defaults) *NameVerificationRelationshipsNameVerificationSubmission {
	return &NameVerificationRelationshipsNameVerificationSubmission{

		Data: make([]*NameVerificationSubmission, 0),
	}
}

func (m *NameVerificationRelationshipsNameVerificationSubmission) WithData(data []*NameVerificationSubmission) *NameVerificationRelationshipsNameVerificationSubmission {

	m.Data = data

	return m
}

// Validate validates this name verification relationships name verification submission
func (m *NameVerificationRelationshipsNameVerificationSubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameVerificationRelationshipsNameVerificationSubmission) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	iDataSize := int64(len(m.Data))

	if err := validate.MinItems("name_verification_submission"+"."+"data", "body", iDataSize, 0); err != nil {
		return err
	}

	if err := validate.MaxItems("name_verification_submission"+"."+"data", "body", iDataSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("name_verification_submission" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NameVerificationRelationshipsNameVerificationSubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NameVerificationRelationshipsNameVerificationSubmission) UnmarshalBinary(b []byte) error {
	var res NameVerificationRelationshipsNameVerificationSubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *NameVerificationRelationshipsNameVerificationSubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
