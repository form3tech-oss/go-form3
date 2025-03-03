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

// TransactionFileAdmissionStatus Status of the transaction file admission
// swagger:model TransactionFileAdmissionStatus
type TransactionFileAdmissionStatus string

const (

	// TransactionFileAdmissionStatusConfirmed captures enum value "confirmed"
	TransactionFileAdmissionStatusConfirmed TransactionFileAdmissionStatus = "confirmed"

	// TransactionFileAdmissionStatusFailed captures enum value "failed"
	TransactionFileAdmissionStatusFailed TransactionFileAdmissionStatus = "failed"

	// TransactionFileAdmissionStatusPending captures enum value "pending"
	TransactionFileAdmissionStatusPending TransactionFileAdmissionStatus = "pending"
)

// for schema
var transactionFileAdmissionStatusEnum []interface{}

func init() {
	var res []TransactionFileAdmissionStatus
	if err := json.Unmarshal([]byte(`["confirmed","failed","pending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transactionFileAdmissionStatusEnum = append(transactionFileAdmissionStatusEnum, v)
	}
}

func (m TransactionFileAdmissionStatus) validateTransactionFileAdmissionStatusEnum(path, location string, value TransactionFileAdmissionStatus) error {
	if err := validate.Enum(path, location, value, transactionFileAdmissionStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this transaction file admission status
func (m TransactionFileAdmissionStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTransactionFileAdmissionStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionFileAdmissionStatus) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
