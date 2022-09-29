// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserDefinedDataForPayment user defined data for payment
// swagger:model UserDefinedDataForPayment
type UserDefinedDataForPayment struct {

	// Key of the pair
	// Required: true
	// Max Length: 45
	// Min Length: 1
	Key *string `json:"key"`

	// Value of the pair
	// Required: true
	// Max Length: 45
	// Min Length: 1
	Value *string `json:"value"`
}

func UserDefinedDataForPaymentWithDefaults(defaults client.Defaults) *UserDefinedDataForPayment {
	return &UserDefinedDataForPayment{

		Key: defaults.GetStringPtr("UserDefinedDataForPayment", "key"),

		Value: defaults.GetStringPtr("UserDefinedDataForPayment", "value"),
	}
}

func (m *UserDefinedDataForPayment) WithKey(key string) *UserDefinedDataForPayment {

	m.Key = &key

	return m
}

func (m *UserDefinedDataForPayment) WithoutKey() *UserDefinedDataForPayment {
	m.Key = nil
	return m
}

func (m *UserDefinedDataForPayment) WithValue(value string) *UserDefinedDataForPayment {

	m.Value = &value

	return m
}

func (m *UserDefinedDataForPayment) WithoutValue() *UserDefinedDataForPayment {
	m.Value = nil
	return m
}

// Validate validates this user defined data for payment
func (m *UserDefinedDataForPayment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserDefinedDataForPayment) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	if err := validate.MinLength("key", "body", string(*m.Key), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("key", "body", string(*m.Key), 45); err != nil {
		return err
	}

	return nil
}

func (m *UserDefinedDataForPayment) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	if err := validate.MinLength("value", "body", string(*m.Value), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("value", "body", string(*m.Value), 45); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserDefinedDataForPayment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserDefinedDataForPayment) UnmarshalBinary(b []byte) error {
	var res UserDefinedDataForPayment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *UserDefinedDataForPayment) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
