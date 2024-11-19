// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PaymentAdmissionTaskRelationships payment admission task relationships
// swagger:model PaymentAdmissionTaskRelationships
type PaymentAdmissionTaskRelationships struct {

	// payment
	Payment *RelationshipPayments `json:"payment,omitempty"`

	// payment admission
	PaymentAdmission *RelationshipPaymentAdmissions `json:"payment_admission,omitempty"`

	// payment admission task
	PaymentAdmissionTask *RelationshipPaymentAdmissionTask `json:"payment_admission_task,omitempty"`
}

func PaymentAdmissionTaskRelationshipsWithDefaults(defaults client.Defaults) *PaymentAdmissionTaskRelationships {
	return &PaymentAdmissionTaskRelationships{

		Payment: RelationshipPaymentsWithDefaults(defaults),

		PaymentAdmission: RelationshipPaymentAdmissionsWithDefaults(defaults),

		PaymentAdmissionTask: RelationshipPaymentAdmissionTaskWithDefaults(defaults),
	}
}

func (m *PaymentAdmissionTaskRelationships) WithPayment(payment RelationshipPayments) *PaymentAdmissionTaskRelationships {

	m.Payment = &payment

	return m
}

func (m *PaymentAdmissionTaskRelationships) WithoutPayment() *PaymentAdmissionTaskRelationships {
	m.Payment = nil
	return m
}

func (m *PaymentAdmissionTaskRelationships) WithPaymentAdmission(paymentAdmission RelationshipPaymentAdmissions) *PaymentAdmissionTaskRelationships {

	m.PaymentAdmission = &paymentAdmission

	return m
}

func (m *PaymentAdmissionTaskRelationships) WithoutPaymentAdmission() *PaymentAdmissionTaskRelationships {
	m.PaymentAdmission = nil
	return m
}

func (m *PaymentAdmissionTaskRelationships) WithPaymentAdmissionTask(paymentAdmissionTask RelationshipPaymentAdmissionTask) *PaymentAdmissionTaskRelationships {

	m.PaymentAdmissionTask = &paymentAdmissionTask

	return m
}

func (m *PaymentAdmissionTaskRelationships) WithoutPaymentAdmissionTask() *PaymentAdmissionTaskRelationships {
	m.PaymentAdmissionTask = nil
	return m
}

// Validate validates this payment admission task relationships
func (m *PaymentAdmissionTaskRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentAdmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentAdmissionTask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PaymentAdmissionTaskRelationships) validatePayment(formats strfmt.Registry) error {

	if swag.IsZero(m.Payment) { // not required
		return nil
	}

	if m.Payment != nil {
		if err := m.Payment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAdmissionTaskRelationships) validatePaymentAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentAdmission) { // not required
		return nil
	}

	if m.PaymentAdmission != nil {
		if err := m.PaymentAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_admission")
			}
			return err
		}
	}

	return nil
}

func (m *PaymentAdmissionTaskRelationships) validatePaymentAdmissionTask(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentAdmissionTask) { // not required
		return nil
	}

	if m.PaymentAdmissionTask != nil {
		if err := m.PaymentAdmissionTask.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_admission_task")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PaymentAdmissionTaskRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PaymentAdmissionTaskRelationships) UnmarshalBinary(b []byte) error {
	var res PaymentAdmissionTaskRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *PaymentAdmissionTaskRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
