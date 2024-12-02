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

// SchemeFileAttributes scheme file attributes
// swagger:model SchemeFileAttributes
type SchemeFileAttributes struct {

	// The format of the file that will be submitted to Form3
	// Required: true
	// Enum: ["xml"]
	FileFormat *string `json:"file_format"`

	// Hashed content of the file
	// Required: true
	// Min Length: 1
	FileHash *string `json:"file_hash"`

	// The size of the file to be uploaded - number of bytes. Max size is 1.4 Gigabyte
	// Required: true
	// Maximum: 1.4e+09
	FileSize *int64 `json:"file_size"`

	// The file type
	// Required: true
	// Enum: ["switch","ISA"]
	FileType *string `json:"file_type"`

	// The algorithm used to generate the signature
	// Required: true
	// Enum: ["SHA256"]
	HashingAlgorithm *string `json:"hashing_algorithm"`

	// The count of chunks to be uploaded to the resource
	// Required: true
	// Minimum: 1
	NumberOfParts *int64 `json:"number_of_parts"`

	// Scheme/gateway that the file is to be processed by
	// Required: true
	// Enum: ["CASS","CISA"]
	PaymentScheme *string `json:"payment_scheme"`
}

func SchemeFileAttributesWithDefaults(defaults client.Defaults) *SchemeFileAttributes {
	return &SchemeFileAttributes{

		FileFormat: defaults.GetStringPtr("SchemeFileAttributes", "file_format"),

		FileHash: defaults.GetStringPtr("SchemeFileAttributes", "file_hash"),

		FileSize: defaults.GetInt64Ptr("SchemeFileAttributes", "file_size"),

		FileType: defaults.GetStringPtr("SchemeFileAttributes", "file_type"),

		HashingAlgorithm: defaults.GetStringPtr("SchemeFileAttributes", "hashing_algorithm"),

		NumberOfParts: defaults.GetInt64Ptr("SchemeFileAttributes", "number_of_parts"),

		PaymentScheme: defaults.GetStringPtr("SchemeFileAttributes", "payment_scheme"),
	}
}

func (m *SchemeFileAttributes) WithFileFormat(fileFormat string) *SchemeFileAttributes {

	m.FileFormat = &fileFormat

	return m
}

func (m *SchemeFileAttributes) WithoutFileFormat() *SchemeFileAttributes {
	m.FileFormat = nil
	return m
}

func (m *SchemeFileAttributes) WithFileHash(fileHash string) *SchemeFileAttributes {

	m.FileHash = &fileHash

	return m
}

func (m *SchemeFileAttributes) WithoutFileHash() *SchemeFileAttributes {
	m.FileHash = nil
	return m
}

func (m *SchemeFileAttributes) WithFileSize(fileSize int64) *SchemeFileAttributes {

	m.FileSize = &fileSize

	return m
}

func (m *SchemeFileAttributes) WithoutFileSize() *SchemeFileAttributes {
	m.FileSize = nil
	return m
}

func (m *SchemeFileAttributes) WithFileType(fileType string) *SchemeFileAttributes {

	m.FileType = &fileType

	return m
}

func (m *SchemeFileAttributes) WithoutFileType() *SchemeFileAttributes {
	m.FileType = nil
	return m
}

func (m *SchemeFileAttributes) WithHashingAlgorithm(hashingAlgorithm string) *SchemeFileAttributes {

	m.HashingAlgorithm = &hashingAlgorithm

	return m
}

func (m *SchemeFileAttributes) WithoutHashingAlgorithm() *SchemeFileAttributes {
	m.HashingAlgorithm = nil
	return m
}

func (m *SchemeFileAttributes) WithNumberOfParts(numberOfParts int64) *SchemeFileAttributes {

	m.NumberOfParts = &numberOfParts

	return m
}

func (m *SchemeFileAttributes) WithoutNumberOfParts() *SchemeFileAttributes {
	m.NumberOfParts = nil
	return m
}

func (m *SchemeFileAttributes) WithPaymentScheme(paymentScheme string) *SchemeFileAttributes {

	m.PaymentScheme = &paymentScheme

	return m
}

func (m *SchemeFileAttributes) WithoutPaymentScheme() *SchemeFileAttributes {
	m.PaymentScheme = nil
	return m
}

// Validate validates this scheme file attributes
func (m *SchemeFileAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFileFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHashingAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumberOfParts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentScheme(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var schemeFileAttributesTypeFileFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["xml"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schemeFileAttributesTypeFileFormatPropEnum = append(schemeFileAttributesTypeFileFormatPropEnum, v)
	}
}

const (

	// SchemeFileAttributesFileFormatXML captures enum value "xml"
	SchemeFileAttributesFileFormatXML string = "xml"
)

// prop value enum
func (m *SchemeFileAttributes) validateFileFormatEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, schemeFileAttributesTypeFileFormatPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SchemeFileAttributes) validateFileFormat(formats strfmt.Registry) error {

	if err := validate.Required("file_format", "body", m.FileFormat); err != nil {
		return err
	}

	// value enum
	if err := m.validateFileFormatEnum("file_format", "body", *m.FileFormat); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFileAttributes) validateFileHash(formats strfmt.Registry) error {

	if err := validate.Required("file_hash", "body", m.FileHash); err != nil {
		return err
	}

	if err := validate.MinLength("file_hash", "body", *m.FileHash, 1); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFileAttributes) validateFileSize(formats strfmt.Registry) error {

	if err := validate.Required("file_size", "body", m.FileSize); err != nil {
		return err
	}

	if err := validate.MaximumInt("file_size", "body", int64(*m.FileSize), 1.4e+09, false); err != nil {
		return err
	}

	return nil
}

var schemeFileAttributesTypeFileTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["switch","ISA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schemeFileAttributesTypeFileTypePropEnum = append(schemeFileAttributesTypeFileTypePropEnum, v)
	}
}

const (

	// SchemeFileAttributesFileTypeSwitch captures enum value "switch"
	SchemeFileAttributesFileTypeSwitch string = "switch"

	// SchemeFileAttributesFileTypeISA captures enum value "ISA"
	SchemeFileAttributesFileTypeISA string = "ISA"
)

// prop value enum
func (m *SchemeFileAttributes) validateFileTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, schemeFileAttributesTypeFileTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SchemeFileAttributes) validateFileType(formats strfmt.Registry) error {

	if err := validate.Required("file_type", "body", m.FileType); err != nil {
		return err
	}

	// value enum
	if err := m.validateFileTypeEnum("file_type", "body", *m.FileType); err != nil {
		return err
	}

	return nil
}

var schemeFileAttributesTypeHashingAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SHA256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schemeFileAttributesTypeHashingAlgorithmPropEnum = append(schemeFileAttributesTypeHashingAlgorithmPropEnum, v)
	}
}

const (

	// SchemeFileAttributesHashingAlgorithmSHA256 captures enum value "SHA256"
	SchemeFileAttributesHashingAlgorithmSHA256 string = "SHA256"
)

// prop value enum
func (m *SchemeFileAttributes) validateHashingAlgorithmEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, schemeFileAttributesTypeHashingAlgorithmPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SchemeFileAttributes) validateHashingAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("hashing_algorithm", "body", m.HashingAlgorithm); err != nil {
		return err
	}

	// value enum
	if err := m.validateHashingAlgorithmEnum("hashing_algorithm", "body", *m.HashingAlgorithm); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFileAttributes) validateNumberOfParts(formats strfmt.Registry) error {

	if err := validate.Required("number_of_parts", "body", m.NumberOfParts); err != nil {
		return err
	}

	if err := validate.MinimumInt("number_of_parts", "body", int64(*m.NumberOfParts), 1, false); err != nil {
		return err
	}

	return nil
}

var schemeFileAttributesTypePaymentSchemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CASS","CISA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schemeFileAttributesTypePaymentSchemePropEnum = append(schemeFileAttributesTypePaymentSchemePropEnum, v)
	}
}

const (

	// SchemeFileAttributesPaymentSchemeCASS captures enum value "CASS"
	SchemeFileAttributesPaymentSchemeCASS string = "CASS"

	// SchemeFileAttributesPaymentSchemeCISA captures enum value "CISA"
	SchemeFileAttributesPaymentSchemeCISA string = "CISA"
)

// prop value enum
func (m *SchemeFileAttributes) validatePaymentSchemeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, schemeFileAttributesTypePaymentSchemePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SchemeFileAttributes) validatePaymentScheme(formats strfmt.Registry) error {

	if err := validate.Required("payment_scheme", "body", m.PaymentScheme); err != nil {
		return err
	}

	// value enum
	if err := m.validatePaymentSchemeEnum("payment_scheme", "body", *m.PaymentScheme); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemeFileAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeFileAttributes) UnmarshalBinary(b []byte) error {
	var res SchemeFileAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeFileAttributes) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
