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

// ReversalAdmissionTaskRelationships reversal admission task relationships
// swagger:model ReversalAdmissionTaskRelationships
type ReversalAdmissionTaskRelationships struct {

	// payment
	Payment *RelationshipPayments `json:"payment,omitempty"`

	// reversal
	Reversal *RelationshipReversals `json:"reversal,omitempty"`

	// reversal admission
	ReversalAdmission *RelationshipReversalAdmissions `json:"reversal_admission,omitempty"`
}

func ReversalAdmissionTaskRelationshipsWithDefaults(defaults client.Defaults) *ReversalAdmissionTaskRelationships {
	return &ReversalAdmissionTaskRelationships{

		Payment: RelationshipPaymentsWithDefaults(defaults),

		Reversal: RelationshipReversalsWithDefaults(defaults),

		ReversalAdmission: RelationshipReversalAdmissionsWithDefaults(defaults),
	}
}

func (m *ReversalAdmissionTaskRelationships) WithPayment(payment RelationshipPayments) *ReversalAdmissionTaskRelationships {

	m.Payment = &payment

	return m
}

func (m *ReversalAdmissionTaskRelationships) WithoutPayment() *ReversalAdmissionTaskRelationships {
	m.Payment = nil
	return m
}

func (m *ReversalAdmissionTaskRelationships) WithReversal(reversal RelationshipReversals) *ReversalAdmissionTaskRelationships {

	m.Reversal = &reversal

	return m
}

func (m *ReversalAdmissionTaskRelationships) WithoutReversal() *ReversalAdmissionTaskRelationships {
	m.Reversal = nil
	return m
}

func (m *ReversalAdmissionTaskRelationships) WithReversalAdmission(reversalAdmission RelationshipReversalAdmissions) *ReversalAdmissionTaskRelationships {

	m.ReversalAdmission = &reversalAdmission

	return m
}

func (m *ReversalAdmissionTaskRelationships) WithoutReversalAdmission() *ReversalAdmissionTaskRelationships {
	m.ReversalAdmission = nil
	return m
}

// Validate validates this reversal admission task relationships
func (m *ReversalAdmissionTaskRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReversal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReversalAdmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReversalAdmissionTaskRelationships) validatePayment(formats strfmt.Registry) error {

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

func (m *ReversalAdmissionTaskRelationships) validateReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.Reversal) { // not required
		return nil
	}

	if m.Reversal != nil {
		if err := m.Reversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reversal")
			}
			return err
		}
	}

	return nil
}

func (m *ReversalAdmissionTaskRelationships) validateReversalAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.ReversalAdmission) { // not required
		return nil
	}

	if m.ReversalAdmission != nil {
		if err := m.ReversalAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reversal_admission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReversalAdmissionTaskRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReversalAdmissionTaskRelationships) UnmarshalBinary(b []byte) error {
	var res ReversalAdmissionTaskRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *ReversalAdmissionTaskRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
