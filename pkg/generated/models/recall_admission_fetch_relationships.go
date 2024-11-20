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

// RecallAdmissionFetchRelationships recall admission fetch relationships
// swagger:model RecallAdmissionFetchRelationships
type RecallAdmissionFetchRelationships struct {

	// payment
	Payment *RelationshipLinks `json:"payment,omitempty"`

	// recall
	Recall *RelationshipLinks `json:"recall,omitempty"`
}

func RecallAdmissionFetchRelationshipsWithDefaults(defaults client.Defaults) *RecallAdmissionFetchRelationships {
	return &RecallAdmissionFetchRelationships{

		Payment: RelationshipLinksWithDefaults(defaults),

		Recall: RelationshipLinksWithDefaults(defaults),
	}
}

func (m *RecallAdmissionFetchRelationships) WithPayment(payment RelationshipLinks) *RecallAdmissionFetchRelationships {

	m.Payment = &payment

	return m
}

func (m *RecallAdmissionFetchRelationships) WithoutPayment() *RecallAdmissionFetchRelationships {
	m.Payment = nil
	return m
}

func (m *RecallAdmissionFetchRelationships) WithRecall(recall RelationshipLinks) *RecallAdmissionFetchRelationships {

	m.Recall = &recall

	return m
}

func (m *RecallAdmissionFetchRelationships) WithoutRecall() *RecallAdmissionFetchRelationships {
	m.Recall = nil
	return m
}

// Validate validates this recall admission fetch relationships
func (m *RecallAdmissionFetchRelationships) Validate(formats strfmt.Registry) error {
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

func (m *RecallAdmissionFetchRelationships) validatePayment(formats strfmt.Registry) error {

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

func (m *RecallAdmissionFetchRelationships) validateRecall(formats strfmt.Registry) error {

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
func (m *RecallAdmissionFetchRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecallAdmissionFetchRelationships) UnmarshalBinary(b []byte) error {
	var res RecallAdmissionFetchRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RecallAdmissionFetchRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
