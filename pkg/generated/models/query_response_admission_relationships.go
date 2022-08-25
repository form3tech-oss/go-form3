// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v5/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// QueryResponseAdmissionRelationships query response admission relationships
// swagger:model QueryResponseAdmissionRelationships
type QueryResponseAdmissionRelationships struct {

	// payment
	Payment *RelationshipsPayment `json:"payment,omitempty"`

	// query
	Query *RelationshipsQuery `json:"query,omitempty"`

	// query response
	QueryResponse *RelationshipsQueryResponse `json:"query_response,omitempty"`

	// recall
	Recall *RelationshipsPaymentRecall `json:"recall,omitempty"`
}

func QueryResponseAdmissionRelationshipsWithDefaults(defaults client.Defaults) *QueryResponseAdmissionRelationships {
	return &QueryResponseAdmissionRelationships{

		Payment: RelationshipsPaymentWithDefaults(defaults),

		Query: RelationshipsQueryWithDefaults(defaults),

		QueryResponse: RelationshipsQueryResponseWithDefaults(defaults),

		Recall: RelationshipsPaymentRecallWithDefaults(defaults),
	}
}

func (m *QueryResponseAdmissionRelationships) WithPayment(payment RelationshipsPayment) *QueryResponseAdmissionRelationships {

	m.Payment = &payment

	return m
}

func (m *QueryResponseAdmissionRelationships) WithoutPayment() *QueryResponseAdmissionRelationships {
	m.Payment = nil
	return m
}

func (m *QueryResponseAdmissionRelationships) WithQuery(query RelationshipsQuery) *QueryResponseAdmissionRelationships {

	m.Query = &query

	return m
}

func (m *QueryResponseAdmissionRelationships) WithoutQuery() *QueryResponseAdmissionRelationships {
	m.Query = nil
	return m
}

func (m *QueryResponseAdmissionRelationships) WithQueryResponse(queryResponse RelationshipsQueryResponse) *QueryResponseAdmissionRelationships {

	m.QueryResponse = &queryResponse

	return m
}

func (m *QueryResponseAdmissionRelationships) WithoutQueryResponse() *QueryResponseAdmissionRelationships {
	m.QueryResponse = nil
	return m
}

func (m *QueryResponseAdmissionRelationships) WithRecall(recall RelationshipsPaymentRecall) *QueryResponseAdmissionRelationships {

	m.Recall = &recall

	return m
}

func (m *QueryResponseAdmissionRelationships) WithoutRecall() *QueryResponseAdmissionRelationships {
	m.Recall = nil
	return m
}

// Validate validates this query response admission relationships
func (m *QueryResponseAdmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueryResponse(formats); err != nil {
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

func (m *QueryResponseAdmissionRelationships) validatePayment(formats strfmt.Registry) error {

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

func (m *QueryResponseAdmissionRelationships) validateQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.Query) { // not required
		return nil
	}

	if m.Query != nil {
		if err := m.Query.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query")
			}
			return err
		}
	}

	return nil
}

func (m *QueryResponseAdmissionRelationships) validateQueryResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.QueryResponse) { // not required
		return nil
	}

	if m.QueryResponse != nil {
		if err := m.QueryResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query_response")
			}
			return err
		}
	}

	return nil
}

func (m *QueryResponseAdmissionRelationships) validateRecall(formats strfmt.Registry) error {

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
func (m *QueryResponseAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResponseAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res QueryResponseAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryResponseAdmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
