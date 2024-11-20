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

// TransactionFileAdmissionRelationships transaction file admission relationships
// swagger:model TransactionFileAdmissionRelationships
type TransactionFileAdmissionRelationships struct {

	// transaction files
	TransactionFiles *TransactionFileAdmissionRelationshipsTransactionFiles `json:"transaction_files,omitempty"`
}

func TransactionFileAdmissionRelationshipsWithDefaults(defaults client.Defaults) *TransactionFileAdmissionRelationships {
	return &TransactionFileAdmissionRelationships{

		TransactionFiles: TransactionFileAdmissionRelationshipsTransactionFilesWithDefaults(defaults),
	}
}

func (m *TransactionFileAdmissionRelationships) WithTransactionFiles(transactionFiles TransactionFileAdmissionRelationshipsTransactionFiles) *TransactionFileAdmissionRelationships {

	m.TransactionFiles = &transactionFiles

	return m
}

func (m *TransactionFileAdmissionRelationships) WithoutTransactionFiles() *TransactionFileAdmissionRelationships {
	m.TransactionFiles = nil
	return m
}

// Validate validates this transaction file admission relationships
func (m *TransactionFileAdmissionRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransactionFiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionFileAdmissionRelationships) validateTransactionFiles(formats strfmt.Registry) error {

	if swag.IsZero(m.TransactionFiles) { // not required
		return nil
	}

	if m.TransactionFiles != nil {
		if err := m.TransactionFiles.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transaction_files")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionFileAdmissionRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionFileAdmissionRelationships) UnmarshalBinary(b []byte) error {
	var res TransactionFileAdmissionRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *TransactionFileAdmissionRelationships) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// TransactionFileAdmissionRelationshipsTransactionFiles transaction file admission relationships transaction files
// swagger:model TransactionFileAdmissionRelationshipsTransactionFiles
type TransactionFileAdmissionRelationshipsTransactionFiles struct {

	// data
	Data []*TransactionFile `json:"data"`
}

func TransactionFileAdmissionRelationshipsTransactionFilesWithDefaults(defaults client.Defaults) *TransactionFileAdmissionRelationshipsTransactionFiles {
	return &TransactionFileAdmissionRelationshipsTransactionFiles{

		Data: make([]*TransactionFile, 0),
	}
}

func (m *TransactionFileAdmissionRelationshipsTransactionFiles) WithData(data []*TransactionFile) *TransactionFileAdmissionRelationshipsTransactionFiles {

	m.Data = data

	return m
}

// Validate validates this transaction file admission relationships transaction files
func (m *TransactionFileAdmissionRelationshipsTransactionFiles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionFileAdmissionRelationshipsTransactionFiles) validateData(formats strfmt.Registry) error {

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
					return ve.ValidateName("transaction_files" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionFileAdmissionRelationshipsTransactionFiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionFileAdmissionRelationshipsTransactionFiles) UnmarshalBinary(b []byte) error {
	var res TransactionFileAdmissionRelationshipsTransactionFiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *TransactionFileAdmissionRelationshipsTransactionFiles) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
