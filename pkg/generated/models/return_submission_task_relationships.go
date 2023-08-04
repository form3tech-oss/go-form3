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
)

// ReturnSubmissionTaskRelationships return submission task relationships
// swagger:model ReturnSubmissionTaskRelationships
type ReturnSubmissionTaskRelationships struct {

	// payment
	Payment *RelationshipPayments `json:"payment,omitempty"`

	// payment return
	PaymentReturn *RelationshipReturns `json:"payment_return,omitempty"`

	// return submission
	ReturnSubmission *RelationshipReturnSubmissions `json:"return_submission,omitempty"`
}

func ReturnSubmissionTaskRelationshipsWithDefaults(defaults client.Defaults) *ReturnSubmissionTaskRelationships {
	return &ReturnSubmissionTaskRelationships{

		Payment: RelationshipPaymentsWithDefaults(defaults),

		PaymentReturn: RelationshipReturnsWithDefaults(defaults),

		ReturnSubmission: RelationshipReturnSubmissionsWithDefaults(defaults),
	}
}

func (m *ReturnSubmissionTaskRelationships) WithPayment(payment RelationshipPayments) *ReturnSubmissionTaskRelationships {

	m.Payment = &payment

	return m
}

func (m *ReturnSubmissionTaskRelationships) WithoutPayment() *ReturnSubmissionTaskRelationships {
	m.Payment = nil
	return m
}

func (m *ReturnSubmissionTaskRelationships) WithPaymentReturn(paymentReturn RelationshipReturns) *ReturnSubmissionTaskRelationships {

	m.PaymentReturn = &paymentReturn

	return m
}

func (m *ReturnSubmissionTaskRelationships) WithoutPaymentReturn() *ReturnSubmissionTaskRelationships {
	m.PaymentReturn = nil
	return m
}

func (m *ReturnSubmissionTaskRelationships) WithReturnSubmission(returnSubmission RelationshipReturnSubmissions) *ReturnSubmissionTaskRelationships {

	m.ReturnSubmission = &returnSubmission

	return m
}

func (m *ReturnSubmissionTaskRelationships) WithoutReturnSubmission() *ReturnSubmissionTaskRelationships {
	m.ReturnSubmission = nil
	return m
}

// Validate validates this return submission task relationships
func (m *ReturnSubmissionTaskRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentReturn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturnSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReturnSubmissionTaskRelationships) validatePayment(formats strfmt.Registry) error {

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

func (m *ReturnSubmissionTaskRelationships) validatePaymentReturn(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentReturn) { // not required
		return nil
	}

	if m.PaymentReturn != nil {
		if err := m.PaymentReturn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("payment_return")
			}
			return err
		}
	}

	return nil
}

func (m *ReturnSubmissionTaskRelationships) validateReturnSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.ReturnSubmission) { // not required
		return nil
	}

	if m.ReturnSubmission != nil {
		if err := m.ReturnSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("return_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReturnSubmissionTaskRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReturnSubmissionTaskRelationships) UnmarshalBinary(b []byte) error {
	var res ReturnSubmissionTaskRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReturnSubmissionTaskRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}