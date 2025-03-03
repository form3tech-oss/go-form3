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

// MandateAdmissionStatusReason mandate admission status reason
// swagger:model MandateAdmissionStatusReason
type MandateAdmissionStatusReason string

const (

	// MandateAdmissionStatusReasonAccepted captures enum value "accepted"
	MandateAdmissionStatusReasonAccepted MandateAdmissionStatusReason = "accepted"

	// MandateAdmissionStatusReasonInvalidBeneficiaryDetails captures enum value "invalid_beneficiary_details"
	MandateAdmissionStatusReasonInvalidBeneficiaryDetails MandateAdmissionStatusReason = "invalid_beneficiary_details"

	// MandateAdmissionStatusReasonBankidNotProvisioned captures enum value "bankid_not_provisioned"
	MandateAdmissionStatusReasonBankidNotProvisioned MandateAdmissionStatusReason = "bankid_not_provisioned"

	// MandateAdmissionStatusReasonUnknownAccountnumber captures enum value "unknown_accountnumber"
	MandateAdmissionStatusReasonUnknownAccountnumber MandateAdmissionStatusReason = "unknown_accountnumber"

	// MandateAdmissionStatusReasonMandateCancelled captures enum value "mandate_cancelled"
	MandateAdmissionStatusReasonMandateCancelled MandateAdmissionStatusReason = "mandate_cancelled"

	// MandateAdmissionStatusReasonMandateExpired captures enum value "mandate_expired"
	MandateAdmissionStatusReasonMandateExpired MandateAdmissionStatusReason = "mandate_expired"

	// MandateAdmissionStatusReasonMandateFailed captures enum value "mandate_failed"
	MandateAdmissionStatusReasonMandateFailed MandateAdmissionStatusReason = "mandate_failed"

	// MandateAdmissionStatusReasonUnknownInstruction captures enum value "unknown_instruction"
	MandateAdmissionStatusReasonUnknownInstruction MandateAdmissionStatusReason = "unknown_instruction"

	// MandateAdmissionStatusReasonPayerReferenceNotUnique captures enum value "payer_reference_not_unique"
	MandateAdmissionStatusReasonPayerReferenceNotUnique MandateAdmissionStatusReason = "payer_reference_not_unique"

	// MandateAdmissionStatusReasonInvalidReference captures enum value "invalid_reference"
	MandateAdmissionStatusReasonInvalidReference MandateAdmissionStatusReason = "invalid_reference"

	// MandateAdmissionStatusReasonNotEligibleForReinstatementByOriginator captures enum value "not_eligible_for_reinstatement_by_originator"
	MandateAdmissionStatusReasonNotEligibleForReinstatementByOriginator MandateAdmissionStatusReason = "not_eligible_for_reinstatement_by_originator"
)

// for schema
var mandateAdmissionStatusReasonEnum []interface{}

func init() {
	var res []MandateAdmissionStatusReason
	if err := json.Unmarshal([]byte(`["accepted","invalid_beneficiary_details","bankid_not_provisioned","unknown_accountnumber","mandate_cancelled","mandate_expired","mandate_failed","unknown_instruction","payer_reference_not_unique","invalid_reference","not_eligible_for_reinstatement_by_originator"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mandateAdmissionStatusReasonEnum = append(mandateAdmissionStatusReasonEnum, v)
	}
}

func (m MandateAdmissionStatusReason) validateMandateAdmissionStatusReasonEnum(path, location string, value MandateAdmissionStatusReason) error {
	if err := validate.Enum(path, location, value, mandateAdmissionStatusReasonEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this mandate admission status reason
func (m MandateAdmissionStatusReason) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMandateAdmissionStatusReasonEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MandateAdmissionStatusReason) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
