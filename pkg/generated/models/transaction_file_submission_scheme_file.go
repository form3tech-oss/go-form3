// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v4/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TransactionFileSubmissionSchemeFile transaction file submission scheme file
// swagger:model TransactionFileSubmissionSchemeFile
type TransactionFileSubmissionSchemeFile struct {

	// Service User Number (SUN) of the payment originator
	// Pattern: .{6}$
	ClearingID string `json:"clearing_id,omitempty"`

	// Number of the file sent to the scheme
	SchemeFileID string `json:"scheme_file_id,omitempty"`

	// Submission serial number
	SchemeSubmissionID string `json:"scheme_submission_id,omitempty"`

	// The count of transactions submitted in a given scheme file
	TransactionCount int64 `json:"transaction_count,omitempty"`

	// The sum of transactions submitted in a given scheme file
	TransactionSum string `json:"transaction_sum,omitempty"`
}

func TransactionFileSubmissionSchemeFileWithDefaults(defaults client.Defaults) *TransactionFileSubmissionSchemeFile {
	return &TransactionFileSubmissionSchemeFile{

		ClearingID: defaults.GetString("TransactionFileSubmissionSchemeFile", "clearing_id"),

		SchemeFileID: defaults.GetString("TransactionFileSubmissionSchemeFile", "scheme_file_id"),

		SchemeSubmissionID: defaults.GetString("TransactionFileSubmissionSchemeFile", "scheme_submission_id"),

		TransactionCount: defaults.GetInt64("TransactionFileSubmissionSchemeFile", "transaction_count"),

		TransactionSum: defaults.GetString("TransactionFileSubmissionSchemeFile", "transaction_sum"),
	}
}

func (m *TransactionFileSubmissionSchemeFile) WithClearingID(clearingID string) *TransactionFileSubmissionSchemeFile {

	m.ClearingID = clearingID

	return m
}

func (m *TransactionFileSubmissionSchemeFile) WithSchemeFileID(schemeFileID string) *TransactionFileSubmissionSchemeFile {

	m.SchemeFileID = schemeFileID

	return m
}

func (m *TransactionFileSubmissionSchemeFile) WithSchemeSubmissionID(schemeSubmissionID string) *TransactionFileSubmissionSchemeFile {

	m.SchemeSubmissionID = schemeSubmissionID

	return m
}

func (m *TransactionFileSubmissionSchemeFile) WithTransactionCount(transactionCount int64) *TransactionFileSubmissionSchemeFile {

	m.TransactionCount = transactionCount

	return m
}

func (m *TransactionFileSubmissionSchemeFile) WithTransactionSum(transactionSum string) *TransactionFileSubmissionSchemeFile {

	m.TransactionSum = transactionSum

	return m
}

// Validate validates this transaction file submission scheme file
func (m *TransactionFileSubmissionSchemeFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClearingID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionFileSubmissionSchemeFile) validateClearingID(formats strfmt.Registry) error {

	if swag.IsZero(m.ClearingID) { // not required
		return nil
	}

	if err := validate.Pattern("clearing_id", "body", string(m.ClearingID), `.{6}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionFileSubmissionSchemeFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionFileSubmissionSchemeFile) UnmarshalBinary(b []byte) error {
	var res TransactionFileSubmissionSchemeFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *TransactionFileSubmissionSchemeFile) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
