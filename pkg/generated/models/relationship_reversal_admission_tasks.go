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

// RelationshipReversalAdmissionTasks relationship reversal admission tasks
// swagger:model RelationshipReversalAdmissionTasks
type RelationshipReversalAdmissionTasks struct {

	// data
	Data []*ReversalAdmissionTask `json:"data"`
}

func RelationshipReversalAdmissionTasksWithDefaults(defaults client.Defaults) *RelationshipReversalAdmissionTasks {
	return &RelationshipReversalAdmissionTasks{

		Data: make([]*ReversalAdmissionTask, 0),
	}
}

func (m *RelationshipReversalAdmissionTasks) WithData(data []*ReversalAdmissionTask) *RelationshipReversalAdmissionTasks {

	m.Data = data

	return m
}

// Validate validates this relationship reversal admission tasks
func (m *RelationshipReversalAdmissionTasks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelationshipReversalAdmissionTasks) validateData(formats strfmt.Registry) error {

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
func (m *RelationshipReversalAdmissionTasks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationshipReversalAdmissionTasks) UnmarshalBinary(b []byte) error {
	var res RelationshipReversalAdmissionTasks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RelationshipReversalAdmissionTasks) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
