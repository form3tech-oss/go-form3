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

// PaymentResourceType payment resource type
// swagger:model PaymentResourceType
type PaymentResourceType string

const (

	// PaymentResourceTypePayments captures enum value "payments"
	PaymentResourceTypePayments PaymentResourceType = "payments"
)

// for schema
var paymentResourceTypeEnum []interface{}

func init() {
	var res []PaymentResourceType
	if err := json.Unmarshal([]byte(`["payments"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		paymentResourceTypeEnum = append(paymentResourceTypeEnum, v)
	}
}

func (m PaymentResourceType) validatePaymentResourceTypeEnum(path, location string, value PaymentResourceType) error {
	if err := validate.Enum(path, location, value, paymentResourceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this payment resource type
func (m PaymentResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePaymentResourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentResourceType) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
