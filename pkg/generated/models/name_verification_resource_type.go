// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NameVerificationResourceType name verification resource type
// swagger:model NameVerificationResourceType
type NameVerificationResourceType string

const (

	// NameVerificationResourceTypeNameVerifications captures enum value "name_verifications"
	NameVerificationResourceTypeNameVerifications NameVerificationResourceType = "name_verifications"

	// NameVerificationResourceTypeNameVerificationAdmissions captures enum value "name_verification_admissions"
	NameVerificationResourceTypeNameVerificationAdmissions NameVerificationResourceType = "name_verification_admissions"

	// NameVerificationResourceTypeNameVerificationSubmissions captures enum value "name_verification_submissions"
	NameVerificationResourceTypeNameVerificationSubmissions NameVerificationResourceType = "name_verification_submissions"
)

// for schema
var nameVerificationResourceTypeEnum []interface{}

func init() {
	var res []NameVerificationResourceType
	if err := json.Unmarshal([]byte(`["name_verifications","name_verification_admissions","name_verification_submissions"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nameVerificationResourceTypeEnum = append(nameVerificationResourceTypeEnum, v)
	}
}

func (m NameVerificationResourceType) validateNameVerificationResourceTypeEnum(path, location string, value NameVerificationResourceType) error {
	if err := validate.Enum(path, location, value, nameVerificationResourceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this name verification resource type
func (m NameVerificationResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateNameVerificationResourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameVerificationResourceType) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
