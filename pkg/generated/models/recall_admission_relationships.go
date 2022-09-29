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

// RecallAdmissionRelationships recall admission relationships
// swagger:model RecallAdmissionRelationships
type RecallAdmissionRelationships struct {

	// payment
	Payment *RelationshipPayments `json:"payment,omitempty"`

	// recall
	Recall *RelationshipRecalls `json:"recall,omitempty"`
}

func RecallAdmissionRelationshipsWithDefaults(defaults client.Defaults) *RecallAdmissionRelationships {
	return &RecallAdmissionRelationships{

		Payment: RelationshipPaymentsWithDefaults(defaults),

		Recall: RelationshipRecallsWithDefaults(defaults),
	}
}

func (m *RecallAdmissionRelationships) WithPayment(payment RelationshipPayments) *RecallAdmissionRelationships {

	m.Payment = &payment

	return m
}

func (m *RecallAdmissionRelationships) WithoutPayment() *RecallAdmissionRelationships {
	m.Payment = nil
	return m
}

func (m *RecallAdmissionRelationships) WithRecall(recall RelationshipRecalls) *RecallAdmissionRelationships {

	m.Recall = &recall

	return m
}

func (m *RecallAdmissionRelationships) WithoutRecall() *RecallAdmissionRelationships {
	m.Recall = nil
	return m
}

// Validate validates this recall admission relationships
func (m *RecallAdmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecall(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecallAdmissionRelationships) validatePayment(formats strfmt.Registry) error {

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

func (m *RecallAdmissionRelationships) validateRecall(formats strfmt.Registry) error {

	if swag.IsZero(m.Recall) { // not required
		return nil
	}

	if m.Recall != nil {
		if err := m.Recall.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recall")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res RecallAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallAdmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
