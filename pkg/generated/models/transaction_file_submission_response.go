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

// TransactionFileSubmissionResponse transaction file submission response
// swagger:model TransactionFileSubmissionResponse
type TransactionFileSubmissionResponse struct {

	// data
	Data *TransactionFileSubmission `json:"data,omitempty"`

	// links
	Links *TransactionFileLinks `json:"links,omitempty"`
}

func TransactionFileSubmissionResponseWithDefaults(defaults client.Defaults) *TransactionFileSubmissionResponse {
	return &TransactionFileSubmissionResponse{

		Data: TransactionFileSubmissionWithDefaults(defaults),

		// TODO Links: TransactionFileLinks,

	}
}

func (m *TransactionFileSubmissionResponse) WithData(data TransactionFileSubmission) *TransactionFileSubmissionResponse {

	m.Data = &data

	return m
}

func (m *TransactionFileSubmissionResponse) WithoutData() *TransactionFileSubmissionResponse {
	m.Data = nil
	return m
}

func (m *TransactionFileSubmissionResponse) WithLinks(links TransactionFileLinks) *TransactionFileSubmissionResponse {

	m.Links = &links

	return m
}

func (m *TransactionFileSubmissionResponse) WithoutLinks() *TransactionFileSubmissionResponse {
	m.Links = nil
	return m
}

// Validate validates this transaction file submission response
func (m *TransactionFileSubmissionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionFileSubmissionResponse) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *TransactionFileSubmissionResponse) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionFileSubmissionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionFileSubmissionResponse) UnmarshalBinary(b []byte) error {
	var res TransactionFileSubmissionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *TransactionFileSubmissionResponse) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
