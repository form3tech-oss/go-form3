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

// RecallDecisionAnswer Answer to the recall request.
// swagger:model RecallDecisionAnswer
type RecallDecisionAnswer string

const (

	// RecallDecisionAnswerAccepted captures enum value "accepted"
	RecallDecisionAnswerAccepted RecallDecisionAnswer = "accepted"

	// RecallDecisionAnswerRejected captures enum value "rejected"
	RecallDecisionAnswerRejected RecallDecisionAnswer = "rejected"

	// RecallDecisionAnswerPending captures enum value "pending"
	RecallDecisionAnswerPending RecallDecisionAnswer = "pending"

	// RecallDecisionAnswerPartiallyAccepted captures enum value "partially_accepted"
	RecallDecisionAnswerPartiallyAccepted RecallDecisionAnswer = "partially_accepted"
)

// for schema
var recallDecisionAnswerEnum []interface{}

func init() {
	var res []RecallDecisionAnswer
	if err := json.Unmarshal([]byte(`["accepted","rejected","pending","partially_accepted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recallDecisionAnswerEnum = append(recallDecisionAnswerEnum, v)
	}
}

func (m RecallDecisionAnswer) validateRecallDecisionAnswerEnum(path, location string, value RecallDecisionAnswer) error {
	if err := validate.Enum(path, location, value, recallDecisionAnswerEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this recall decision answer
func (m RecallDecisionAnswer) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateRecallDecisionAnswerEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecallDecisionAnswer) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
