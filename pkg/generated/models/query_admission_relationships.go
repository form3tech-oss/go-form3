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
	"github.com/go-openapi/validate"
)

// QueryAdmissionRelationships query admission relationships
// swagger:model QueryAdmissionRelationships
type QueryAdmissionRelationships struct {

	// query
	// Required: true
	Query *RelationshipsFullQuery `json:"query"`
}

func QueryAdmissionRelationshipsWithDefaults(defaults client.Defaults) *QueryAdmissionRelationships {
	return &QueryAdmissionRelationships{

		Query: RelationshipsFullQueryWithDefaults(defaults),
	}
}

func (m *QueryAdmissionRelationships) WithQuery(query RelationshipsFullQuery) *QueryAdmissionRelationships {

	m.Query = &query

	return m
}

func (m *QueryAdmissionRelationships) WithoutQuery() *QueryAdmissionRelationships {
	m.Query = nil
	return m
}

// Validate validates this query admission relationships
func (m *QueryAdmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryAdmissionRelationships) validateQuery(formats strfmt.Registry) error {

	if err := validate.Required("query", "body", m.Query); err != nil {
		return err
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

// MarshalBinary interface implementation
func (m *QueryAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res QueryAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *QueryAdmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
