// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// NameVerificationSubmissionReasonCode name verification submission reason code
// swagger:model NameVerificationSubmissionReasonCode
type NameVerificationSubmissionReasonCode string

const (

	// NameVerificationSubmissionReasonCodeANNM captures enum value "ANNM"
	NameVerificationSubmissionReasonCodeANNM NameVerificationSubmissionReasonCode = "ANNM"

	// NameVerificationSubmissionReasonCodeMBAM captures enum value "MBAM"
	NameVerificationSubmissionReasonCodeMBAM NameVerificationSubmissionReasonCode = "MBAM"

	// NameVerificationSubmissionReasonCodeBANM captures enum value "BANM"
	NameVerificationSubmissionReasonCodeBANM NameVerificationSubmissionReasonCode = "BANM"

	// NameVerificationSubmissionReasonCodePANM captures enum value "PANM"
	NameVerificationSubmissionReasonCodePANM NameVerificationSubmissionReasonCode = "PANM"

	// NameVerificationSubmissionReasonCodeBAMM captures enum value "BAMM"
	NameVerificationSubmissionReasonCodeBAMM NameVerificationSubmissionReasonCode = "BAMM"

	// NameVerificationSubmissionReasonCodePAMM captures enum value "PAMM"
	NameVerificationSubmissionReasonCodePAMM NameVerificationSubmissionReasonCode = "PAMM"

	// NameVerificationSubmissionReasonCodeAC01 captures enum value "AC01"
	NameVerificationSubmissionReasonCodeAC01 NameVerificationSubmissionReasonCode = "AC01"

	// NameVerificationSubmissionReasonCodeIVCR captures enum value "IVCR"
	NameVerificationSubmissionReasonCodeIVCR NameVerificationSubmissionReasonCode = "IVCR"

	// NameVerificationSubmissionReasonCodeCASS captures enum value "CASS"
	NameVerificationSubmissionReasonCodeCASS NameVerificationSubmissionReasonCode = "CASS"

	// NameVerificationSubmissionReasonCodeSCNS captures enum value "SCNS"
	NameVerificationSubmissionReasonCodeSCNS NameVerificationSubmissionReasonCode = "SCNS"

	// NameVerificationSubmissionReasonCodeACNS captures enum value "ACNS"
	NameVerificationSubmissionReasonCodeACNS NameVerificationSubmissionReasonCode = "ACNS"

	// NameVerificationSubmissionReasonCodeOPTO captures enum value "OPTO"
	NameVerificationSubmissionReasonCodeOPTO NameVerificationSubmissionReasonCode = "OPTO"
)

// for schema
var nameVerificationSubmissionReasonCodeEnum []interface{}

func init() {
	var res []NameVerificationSubmissionReasonCode
	if err := json.Unmarshal([]byte(`["ANNM","MBAM","BANM","PANM","BAMM","PAMM","AC01","IVCR","CASS","SCNS","ACNS","OPTO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nameVerificationSubmissionReasonCodeEnum = append(nameVerificationSubmissionReasonCodeEnum, v)
	}
}

func (m NameVerificationSubmissionReasonCode) validateNameVerificationSubmissionReasonCodeEnum(path, location string, value NameVerificationSubmissionReasonCode) error {
	if err := validate.Enum(path, location, value, nameVerificationSubmissionReasonCodeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this name verification submission reason code
func (m NameVerificationSubmissionReasonCode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateNameVerificationSubmissionReasonCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NameVerificationSubmissionReasonCode) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}