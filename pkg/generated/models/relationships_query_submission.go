// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v7/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RelationshipsQuerySubmission relationships query submission
// swagger:model RelationshipsQuerySubmission
type RelationshipsQuerySubmission struct {

	// data
	Data []*QuerySubmission `json:"data"`
}

func RelationshipsQuerySubmissionWithDefaults(defaults client.Defaults) *RelationshipsQuerySubmission {
	return &RelationshipsQuerySubmission{

		Data: make([]*QuerySubmission, 0),
	}
}

func (m *RelationshipsQuerySubmission) WithData(data []*QuerySubmission) *RelationshipsQuerySubmission {

	m.Data = data

	return m
}

// Validate validates this relationships query submission
func (m *RelationshipsQuerySubmission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelationshipsQuerySubmission) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelationshipsQuerySubmission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationshipsQuerySubmission) UnmarshalBinary(b []byte) error {
	var res RelationshipsQuerySubmission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RelationshipsQuerySubmission) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
