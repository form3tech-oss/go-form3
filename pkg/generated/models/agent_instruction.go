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

// AgentInstruction agent instruction
// swagger:model AgentInstruction
type AgentInstruction struct {

	// instructions
	Instructions []*Instruction `json:"instructions"`

	// role
	Role AgentInstructionRole `json:"role,omitempty"`
}

func AgentInstructionWithDefaults(defaults client.Defaults) *AgentInstruction {
	return &AgentInstruction{

		Instructions: make([]*Instruction, 0),

		// TODO Role: AgentInstructionRole,

	}
}

func (m *AgentInstruction) WithInstructions(instructions []*Instruction) *AgentInstruction {

	m.Instructions = instructions

	return m
}

func (m *AgentInstruction) WithRole(role AgentInstructionRole) *AgentInstruction {

	m.Role = role

	return m
}

// Validate validates this agent instruction
func (m *AgentInstruction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstructions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgentInstruction) validateInstructions(formats strfmt.Registry) error {

	if swag.IsZero(m.Instructions) { // not required
		return nil
	}

	for i := 0; i < len(m.Instructions); i++ {
		if swag.IsZero(m.Instructions[i]) { // not required
			continue
		}

		if m.Instructions[i] != nil {
			if err := m.Instructions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instructions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AgentInstruction) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if err := m.Role.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("role")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AgentInstruction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgentInstruction) UnmarshalBinary(b []byte) error {
	var res AgentInstruction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *AgentInstruction) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}