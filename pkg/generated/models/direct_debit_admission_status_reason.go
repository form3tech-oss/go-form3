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

// DirectDebitAdmissionStatusReason direct debit admission status reason
// swagger:model DirectDebitAdmissionStatusReason
type DirectDebitAdmissionStatusReason string

const (

	// DirectDebitAdmissionStatusReasonAccepted captures enum value "accepted"
	DirectDebitAdmissionStatusReasonAccepted DirectDebitAdmissionStatusReason = "accepted"

	// DirectDebitAdmissionStatusReasonInvalidBeneficiaryDetails captures enum value "invalid_beneficiary_details"
	DirectDebitAdmissionStatusReasonInvalidBeneficiaryDetails DirectDebitAdmissionStatusReason = "invalid_beneficiary_details"

	// DirectDebitAdmissionStatusReasonBankidNotProvisioned captures enum value "bankid_not_provisioned"
	DirectDebitAdmissionStatusReasonBankidNotProvisioned DirectDebitAdmissionStatusReason = "bankid_not_provisioned"

	// DirectDebitAdmissionStatusReasonUnknownAccountnumber captures enum value "unknown_accountnumber"
	DirectDebitAdmissionStatusReasonUnknownAccountnumber DirectDebitAdmissionStatusReason = "unknown_accountnumber"

	// DirectDebitAdmissionStatusReasonPendingSettlement captures enum value "pending_settlement"
	DirectDebitAdmissionStatusReasonPendingSettlement DirectDebitAdmissionStatusReason = "pending_settlement"
)

// for schema
var directDebitAdmissionStatusReasonEnum []interface{}

func init() {
	var res []DirectDebitAdmissionStatusReason
	if err := json.Unmarshal([]byte(`["accepted","invalid_beneficiary_details","bankid_not_provisioned","unknown_accountnumber","pending_settlement"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		directDebitAdmissionStatusReasonEnum = append(directDebitAdmissionStatusReasonEnum, v)
	}
}

func (m DirectDebitAdmissionStatusReason) validateDirectDebitAdmissionStatusReasonEnum(path, location string, value DirectDebitAdmissionStatusReason) error {
	if err := validate.Enum(path, location, value, directDebitAdmissionStatusReasonEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this direct debit admission status reason
func (m DirectDebitAdmissionStatusReason) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDirectDebitAdmissionStatusReasonEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectDebitAdmissionStatusReason) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
