// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TransactionFileRelationships transaction file relationships
// swagger:model TransactionFileRelationships
type TransactionFileRelationships struct {

	// reports
	Reports *ThinRelationship `json:"reports,omitempty"`

	// transaction file admissions
	TransactionFileAdmissions *TransactionFileRelationshipsTransactionFileAdmissions `json:"transaction_file_admissions,omitempty"`

	// transaction file submissions
	TransactionFileSubmissions *TransactionFileRelationshipsTransactionFileSubmissions `json:"transaction_file_submissions,omitempty"`
}

func TransactionFileRelationshipsWithDefaults(defaults client.Defaults) *TransactionFileRelationships {
	return &TransactionFileRelationships{

		Reports: ThinRelationshipWithDefaults(defaults),

		TransactionFileAdmissions: TransactionFileRelationshipsTransactionFileAdmissionsWithDefaults(defaults),

		TransactionFileSubmissions: TransactionFileRelationshipsTransactionFileSubmissionsWithDefaults(defaults),
	}
}

func (m *TransactionFileRelationships) WithReports(reports ThinRelationship) *TransactionFileRelationships {

	m.Reports = &reports

	return m
}

func (m *TransactionFileRelationships) WithoutReports() *TransactionFileRelationships {
	m.Reports = nil
	return m
}

func (m *TransactionFileRelationships) WithTransactionFileAdmissions(transactionFileAdmissions TransactionFileRelationshipsTransactionFileAdmissions) *TransactionFileRelationships {

	m.TransactionFileAdmissions = &transactionFileAdmissions

	return m
}

func (m *TransactionFileRelationships) WithoutTransactionFileAdmissions() *TransactionFileRelationships {
	m.TransactionFileAdmissions = nil
	return m
}

func (m *TransactionFileRelationships) WithTransactionFileSubmissions(transactionFileSubmissions TransactionFileRelationshipsTransactionFileSubmissions) *TransactionFileRelationships {

	m.TransactionFileSubmissions = &transactionFileSubmissions

	return m
}

func (m *TransactionFileRelationships) WithoutTransactionFileSubmissions() *TransactionFileRelationships {
	m.TransactionFileSubmissions = nil
	return m
}

// Validate validates this transaction file relationships
func (m *TransactionFileRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReports(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionFileAdmissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransactionFileSubmissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionFileRelationships) validateReports(formats strfmt.Registry) error {

	if swag.IsZero(m.Reports) { // not required
		return nil
	}

	if m.Reports != nil {
		if err := m.Reports.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reports")
			}
			return err
		}
	}

	return nil
}

func (m *TransactionFileRelationships) validateTransactionFileAdmissions(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionFileAdmissions) { // not required
		return nil
	}

	if m.TransactionFileAdmissions != nil {
		if err := m.TransactionFileAdmissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transaction_file_admissions")
			}
			return err
		}
	}

	return nil
}

func (m *TransactionFileRelationships) validateTransactionFileSubmissions(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionFileSubmissions) { // not required
		return nil
	}

	if m.TransactionFileSubmissions != nil {
		if err := m.TransactionFileSubmissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transaction_file_submissions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionFileRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionFileRelationships) UnmarshalBinary(b []byte) error {
	var res TransactionFileRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *TransactionFileRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// TransactionFileRelationshipsTransactionFileAdmissions transaction file relationships transaction file admissions
// swagger:model TransactionFileRelationshipsTransactionFileAdmissions
type TransactionFileRelationshipsTransactionFileAdmissions struct {

	// data
	Data []*TransactionFileAdmission `json:"data"`
}

func TransactionFileRelationshipsTransactionFileAdmissionsWithDefaults(defaults client.Defaults) *TransactionFileRelationshipsTransactionFileAdmissions {
	return &TransactionFileRelationshipsTransactionFileAdmissions{

		Data: make([]*TransactionFileAdmission, 0),
	}
}

func (m *TransactionFileRelationshipsTransactionFileAdmissions) WithData(data []*TransactionFileAdmission) *TransactionFileRelationshipsTransactionFileAdmissions {

	m.Data = data

	return m
}

// Validate validates this transaction file relationships transaction file admissions
func (m *TransactionFileRelationshipsTransactionFileAdmissions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionFileRelationshipsTransactionFileAdmissions) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("transaction_file_admissions" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionFileRelationshipsTransactionFileAdmissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionFileRelationshipsTransactionFileAdmissions) UnmarshalBinary(b []byte) error {
	var res TransactionFileRelationshipsTransactionFileAdmissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *TransactionFileRelationshipsTransactionFileAdmissions) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// TransactionFileRelationshipsTransactionFileSubmissions transaction file relationships transaction file submissions
// swagger:model TransactionFileRelationshipsTransactionFileSubmissions
type TransactionFileRelationshipsTransactionFileSubmissions struct {

	// data
	Data []*TransactionFileSubmission `json:"data"`
}

func TransactionFileRelationshipsTransactionFileSubmissionsWithDefaults(defaults client.Defaults) *TransactionFileRelationshipsTransactionFileSubmissions {
	return &TransactionFileRelationshipsTransactionFileSubmissions{

		Data: make([]*TransactionFileSubmission, 0),
	}
}

func (m *TransactionFileRelationshipsTransactionFileSubmissions) WithData(data []*TransactionFileSubmission) *TransactionFileRelationshipsTransactionFileSubmissions {

	m.Data = data

	return m
}

// Validate validates this transaction file relationships transaction file submissions
func (m *TransactionFileRelationshipsTransactionFileSubmissions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionFileRelationshipsTransactionFileSubmissions) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("transaction_file_submissions" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionFileRelationshipsTransactionFileSubmissions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionFileRelationshipsTransactionFileSubmissions) UnmarshalBinary(b []byte) error {
	var res TransactionFileRelationshipsTransactionFileSubmissions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *TransactionFileRelationshipsTransactionFileSubmissions) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
