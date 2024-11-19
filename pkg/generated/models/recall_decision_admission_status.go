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

// RecallDecisionAdmissionStatus [Status](http://draft-api-docs.form3.tech/api.html#enumerations-payment-admission-status) of the admission
// swagger:model RecallDecisionAdmissionStatus
type RecallDecisionAdmissionStatus string

const (

	// RecallDecisionAdmissionStatusConfirmed captures enum value "confirmed"
	RecallDecisionAdmissionStatusConfirmed RecallDecisionAdmissionStatus = "confirmed"

	// RecallDecisionAdmissionStatusFailed captures enum value "failed"
	RecallDecisionAdmissionStatusFailed RecallDecisionAdmissionStatus = "failed"
)

// for schema
var recallDecisionAdmissionStatusEnum []interface{}

func init() {
	var res []RecallDecisionAdmissionStatus
	if err := json.Unmarshal([]byte(`["confirmed","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recallDecisionAdmissionStatusEnum = append(recallDecisionAdmissionStatusEnum, v)
	}
}

func (m RecallDecisionAdmissionStatus) validateRecallDecisionAdmissionStatusEnum(path, location string, value RecallDecisionAdmissionStatus) error {
	if err := validate.Enum(path, location, value, recallDecisionAdmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this recall decision admission status
func (m RecallDecisionAdmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRecallDecisionAdmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecallDecisionAdmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
