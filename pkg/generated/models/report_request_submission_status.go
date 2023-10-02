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

// ReportRequestSubmissionStatus report request submission status
// swagger:model ReportRequestSubmissionStatus
type ReportRequestSubmissionStatus string

const (

	// ReportRequestSubmissionStatusDeliveryConfirmed captures enum value "delivery_confirmed"
	ReportRequestSubmissionStatusDeliveryConfirmed ReportRequestSubmissionStatus = "delivery_confirmed"

	// ReportRequestSubmissionStatusDeliveryFailed captures enum value "delivery_failed"
	ReportRequestSubmissionStatusDeliveryFailed ReportRequestSubmissionStatus = "delivery_failed"

	// ReportRequestSubmissionStatusPending captures enum value "pending"
	ReportRequestSubmissionStatusPending ReportRequestSubmissionStatus = "pending"
)

// for schema
var reportRequestSubmissionStatusEnum []interface{}

func init() {
	var res []ReportRequestSubmissionStatus
	if err := json.Unmarshal([]byte(`["delivery_confirmed","delivery_failed","pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportRequestSubmissionStatusEnum = append(reportRequestSubmissionStatusEnum, v)
	}
}

func (m ReportRequestSubmissionStatus) validateReportRequestSubmissionStatusEnum(path, location string, value ReportRequestSubmissionStatus) error {
	if err := validate.Enum(path, location, value, reportRequestSubmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this report request submission status
func (m ReportRequestSubmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateReportRequestSubmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportRequestSubmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
