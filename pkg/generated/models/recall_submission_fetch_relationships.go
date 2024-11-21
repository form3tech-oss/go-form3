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

// RecallSubmissionFetchRelationships recall submission fetch relationships
// swagger:model RecallSubmissionFetchRelationships
type RecallSubmissionFetchRelationships struct {

	// payment
	Payment *RelationshipLinks `json:"payment,omitempty"`

	// recall
	Recall *RelationshipLinks `json:"recall,omitempty"`
}

func RecallSubmissionFetchRelationshipsWithDefaults(defaults client.Defaults) *RecallSubmissionFetchRelationships {
	return &RecallSubmissionFetchRelationships{

		Payment: RelationshipLinksWithDefaults(defaults),

		Recall: RelationshipLinksWithDefaults(defaults),
	}
}

func (m *RecallSubmissionFetchRelationships) WithPayment(payment RelationshipLinks) *RecallSubmissionFetchRelationships {

	m.Payment = &payment

	return m
}

func (m *RecallSubmissionFetchRelationships) WithoutPayment() *RecallSubmissionFetchRelationships {
	m.Payment = nil
	return m
}

func (m *RecallSubmissionFetchRelationships) WithRecall(recall RelationshipLinks) *RecallSubmissionFetchRelationships {

	m.Recall = &recall

	return m
}

func (m *RecallSubmissionFetchRelationships) WithoutRecall() *RecallSubmissionFetchRelationships {
	m.Recall = nil
	return m
}

// Validate validates this recall submission fetch relationships
func (m *RecallSubmissionFetchRelationships) Validate(formats strfmt.Registry) error {
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

func (m *RecallSubmissionFetchRelationships) validatePayment(formats strfmt.Registry) error {

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

func (m *RecallSubmissionFetchRelationships) validateRecall(formats strfmt.Registry) error {

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
func (m *RecallSubmissionFetchRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallSubmissionFetchRelationships) UnmarshalBinary(b []byte) error {
	var res RecallSubmissionFetchRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallSubmissionFetchRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}