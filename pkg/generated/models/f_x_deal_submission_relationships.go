// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FXDealSubmissionRelationships f x deal submission relationships
// swagger:model FXDealSubmissionRelationships
type FXDealSubmissionRelationships struct {

	// fx deal
	// Required: true
	FxDeal *FXDealSubmissionRelationshipFXDeal `json:"fx_deal"`
}

func FXDealSubmissionRelationshipsWithDefaults(defaults client.Defaults) *FXDealSubmissionRelationships {
	return &FXDealSubmissionRelationships{

		FxDeal: FXDealSubmissionRelationshipFXDealWithDefaults(defaults),
	}
}

func (m *FXDealSubmissionRelationships) WithFxDeal(fxDeal FXDealSubmissionRelationshipFXDeal) *FXDealSubmissionRelationships {

	m.FxDeal = &fxDeal

	return m
}

func (m *FXDealSubmissionRelationships) WithoutFxDeal() *FXDealSubmissionRelationships {
	m.FxDeal = nil
	return m
}

// Validate validates this f x deal submission relationships
func (m *FXDealSubmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFxDeal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FXDealSubmissionRelationships) validateFxDeal(formats strfmt.Registry) error {

	if err := validate.Required("fx_deal", "body", m.FxDeal); err != nil {
		return err
	}

	if m.FxDeal != nil {
		if err := m.FxDeal.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fx_deal")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FXDealSubmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FXDealSubmissionRelationships) UnmarshalBinary(b []byte) error {
	var res FXDealSubmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *FXDealSubmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
