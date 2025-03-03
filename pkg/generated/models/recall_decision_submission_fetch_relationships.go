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

// RecallDecisionSubmissionFetchRelationships recall decision submission fetch relationships
// swagger:model RecallDecisionSubmissionFetchRelationships
type RecallDecisionSubmissionFetchRelationships struct {

	// payment
	Payment *RelationshipLinks `json:"payment,omitempty"`

	// recall
	Recall *RelationshipLinks `json:"recall,omitempty"`

	// recall decision
	RecallDecision *RelationshipLinks `json:"recall_decision,omitempty"`
}

func RecallDecisionSubmissionFetchRelationshipsWithDefaults(defaults client.Defaults) *RecallDecisionSubmissionFetchRelationships {
	return &RecallDecisionSubmissionFetchRelationships{

		Payment: RelationshipLinksWithDefaults(defaults),

		Recall: RelationshipLinksWithDefaults(defaults),

		RecallDecision: RelationshipLinksWithDefaults(defaults),
	}
}

func (m *RecallDecisionSubmissionFetchRelationships) WithPayment(payment RelationshipLinks) *RecallDecisionSubmissionFetchRelationships {

	m.Payment = &payment

	return m
}

func (m *RecallDecisionSubmissionFetchRelationships) WithoutPayment() *RecallDecisionSubmissionFetchRelationships {
	m.Payment = nil
	return m
}

func (m *RecallDecisionSubmissionFetchRelationships) WithRecall(recall RelationshipLinks) *RecallDecisionSubmissionFetchRelationships {

	m.Recall = &recall

	return m
}

func (m *RecallDecisionSubmissionFetchRelationships) WithoutRecall() *RecallDecisionSubmissionFetchRelationships {
	m.Recall = nil
	return m
}

func (m *RecallDecisionSubmissionFetchRelationships) WithRecallDecision(recallDecision RelationshipLinks) *RecallDecisionSubmissionFetchRelationships {

	m.RecallDecision = &recallDecision

	return m
}

func (m *RecallDecisionSubmissionFetchRelationships) WithoutRecallDecision() *RecallDecisionSubmissionFetchRelationships {
	m.RecallDecision = nil
	return m
}

// Validate validates this recall decision submission fetch relationships
func (m *RecallDecisionSubmissionFetchRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecallDecision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecallDecisionSubmissionFetchRelationships) validatePayment(formats strfmt.Registry) error {

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

func (m *RecallDecisionSubmissionFetchRelationships) validateRecall(formats strfmt.Registry) error {

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

func (m *RecallDecisionSubmissionFetchRelationships) validateRecallDecision(formats strfmt.Registry) error {

	if swag.IsZero(m.RecallDecision) { // not required
		return nil
	}

	if m.RecallDecision != nil {
		if err := m.RecallDecision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recall_decision")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecallDecisionSubmissionFetchRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallDecisionSubmissionFetchRelationships) UnmarshalBinary(b []byte) error {
	var res RecallDecisionSubmissionFetchRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallDecisionSubmissionFetchRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
