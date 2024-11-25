// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RecallRelationships recall relationships
// swagger:model RecallRelationships
type RecallRelationships struct {

	// ID of the payment resource related to the recall
	Payment *RelationshipLinks `json:"payment,omitempty"`

	// ID of the recall admission resource related to the recall
	RecallAdmission *RelationshipLinks `json:"recall_admission,omitempty"`

	// ID of the recall decision resource related to the recall
	RecallDecisions *RelationshipLinks `json:"recall_decisions,omitempty"`

	// ID of the recall reversal resource related to the recall
	RecallReversal *RelationshipLinks `json:"recall_reversal,omitempty"`

	// ID of the recall submission resource related to the recall
	RecallSubmission *RelationshipLinks `json:"recall_submission,omitempty"`
}

func RecallRelationshipsWithDefaults(defaults client.Defaults) *RecallRelationships {
	return &RecallRelationships{

		Payment: RelationshipLinksWithDefaults(defaults),

		RecallAdmission: RelationshipLinksWithDefaults(defaults),

		RecallDecisions: RelationshipLinksWithDefaults(defaults),

		RecallReversal: RelationshipLinksWithDefaults(defaults),

		RecallSubmission: RelationshipLinksWithDefaults(defaults),
	}
}

func (m *RecallRelationships) WithPayment(payment RelationshipLinks) *RecallRelationships {

	m.Payment = &payment

	return m
}

func (m *RecallRelationships) WithoutPayment() *RecallRelationships {
	m.Payment = nil
	return m
}

func (m *RecallRelationships) WithRecallAdmission(recallAdmission RelationshipLinks) *RecallRelationships {

	m.RecallAdmission = &recallAdmission

	return m
}

func (m *RecallRelationships) WithoutRecallAdmission() *RecallRelationships {
	m.RecallAdmission = nil
	return m
}

func (m *RecallRelationships) WithRecallDecisions(recallDecisions RelationshipLinks) *RecallRelationships {

	m.RecallDecisions = &recallDecisions

	return m
}

func (m *RecallRelationships) WithoutRecallDecisions() *RecallRelationships {
	m.RecallDecisions = nil
	return m
}

func (m *RecallRelationships) WithRecallReversal(recallReversal RelationshipLinks) *RecallRelationships {

	m.RecallReversal = &recallReversal

	return m
}

func (m *RecallRelationships) WithoutRecallReversal() *RecallRelationships {
	m.RecallReversal = nil
	return m
}

func (m *RecallRelationships) WithRecallSubmission(recallSubmission RelationshipLinks) *RecallRelationships {

	m.RecallSubmission = &recallSubmission

	return m
}

func (m *RecallRelationships) WithoutRecallSubmission() *RecallRelationships {
	m.RecallSubmission = nil
	return m
}

// Validate validates this recall relationships
func (m *RecallRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecallAdmission(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecallDecisions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecallReversal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecallSubmission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecallRelationships) validatePayment(formats strfmt.Registry) error {

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

func (m *RecallRelationships) validateRecallAdmission(formats strfmt.Registry) error {

	if swag.IsZero(m.RecallAdmission) { // not required
		return nil
	}

	if m.RecallAdmission != nil {
		if err := m.RecallAdmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recall_admission")
			}
			return err
		}
	}

	return nil
}

func (m *RecallRelationships) validateRecallDecisions(formats strfmt.Registry) error {

	if swag.IsZero(m.RecallDecisions) { // not required
		return nil
	}

	if m.RecallDecisions != nil {
		if err := m.RecallDecisions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recall_decisions")
			}
			return err
		}
	}

	return nil
}

func (m *RecallRelationships) validateRecallReversal(formats strfmt.Registry) error {

	if swag.IsZero(m.RecallReversal) { // not required
		return nil
	}

	if m.RecallReversal != nil {
		if err := m.RecallReversal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recall_reversal")
			}
			return err
		}
	}

	return nil
}

func (m *RecallRelationships) validateRecallSubmission(formats strfmt.Registry) error {

	if swag.IsZero(m.RecallSubmission) { // not required
		return nil
	}

	if m.RecallSubmission != nil {
		if err := m.RecallSubmission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recall_submission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallRelationships) UnmarshalBinary(b []byte) error {
	var res RecallRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
