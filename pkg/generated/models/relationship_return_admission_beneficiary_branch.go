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

// RelationshipReturnAdmissionBeneficiaryBranch relationship return admission beneficiary branch
// swagger:model RelationshipReturnAdmissionBeneficiaryBranch
type RelationshipReturnAdmissionBeneficiaryBranch struct {

	// data
	Data []*ReturnAdmissionBeneficiaryBranch `json:"data"`
}

func RelationshipReturnAdmissionBeneficiaryBranchWithDefaults(defaults client.Defaults) *RelationshipReturnAdmissionBeneficiaryBranch {
	return &RelationshipReturnAdmissionBeneficiaryBranch{

		Data: make([]*ReturnAdmissionBeneficiaryBranch, 0),
	}
}

func (m *RelationshipReturnAdmissionBeneficiaryBranch) WithData(data []*ReturnAdmissionBeneficiaryBranch) *RelationshipReturnAdmissionBeneficiaryBranch {

	m.Data = data

	return m
}

// Validate validates this relationship return admission beneficiary branch
func (m *RelationshipReturnAdmissionBeneficiaryBranch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelationshipReturnAdmissionBeneficiaryBranch) validateData(formats strfmt.Registry) error {

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
func (m *RelationshipReturnAdmissionBeneficiaryBranch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationshipReturnAdmissionBeneficiaryBranch) UnmarshalBinary(b []byte) error {
	var res RelationshipReturnAdmissionBeneficiaryBranch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *RelationshipReturnAdmissionBeneficiaryBranch) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
