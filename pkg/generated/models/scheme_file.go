// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"log"

	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SchemeFile scheme file
// swagger:model SchemeFile
type SchemeFile struct {

	// id
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// type
	// Enum: [scheme_files]
	Type string `json:"type,omitempty"`

	// Unique ID of the organisation this resource is created by
	// Required: true
	// Format: uuid
	OrganisationID *strfmt.UUID `json:"organisation_id"`

	// Version number
	// Minimum: 0
	Version *int64 `json:"version,omitempty"`

	// created on
	// Read Only: true
	// Format: date-time
	CreatedOn strfmt.DateTime `json:"created_on,omitempty"`

	// modified on
	// Read Only: true
	// Format: date-time
	ModifiedOn strfmt.DateTime `json:"modified_on,omitempty"`

	// attributes
	// Required: true
	Attributes *SchemeFileAttributes `json:"attributes"`

	// relationships
	Relationships *SchemeFileRelationships `json:"relationships,omitempty"`
}

func SchemeFileWithDefaults(defaults client.Defaults) *SchemeFile {
	return &SchemeFile{

		ID: defaults.GetStrfmtUUIDPtr("SchemeFile", "id"),

		Type: defaults.GetString("SchemeFile", "type"),

		OrganisationID: defaults.GetStrfmtUUIDPtr("SchemeFile", "organisation_id"),

		Version: defaults.GetInt64Ptr("SchemeFile", "version"),

		CreatedOn: defaults.GetStrfmtDateTime("SchemeFile", "created_on"),

		ModifiedOn: defaults.GetStrfmtDateTime("SchemeFile", "modified_on"),

		Attributes: SchemeFileAttributesWithDefaults(defaults),

		Relationships: SchemeFileRelationshipsWithDefaults(defaults),
	}
}

func (m *SchemeFile) WithID(id strfmt.UUID) *SchemeFile {

	m.ID = &id

	return m
}

func (m *SchemeFile) WithoutID() *SchemeFile {
	m.ID = nil
	return m
}

func (m *SchemeFile) WithType(typeVar string) *SchemeFile {

	m.Type = typeVar

	return m
}

func (m *SchemeFile) WithOrganisationID(organisationID strfmt.UUID) *SchemeFile {

	m.OrganisationID = &organisationID

	return m
}

func (m *SchemeFile) WithoutOrganisationID() *SchemeFile {
	m.OrganisationID = nil
	return m
}

func (m *SchemeFile) WithVersion(version int64) *SchemeFile {

	m.Version = &version

	return m
}

func (m *SchemeFile) WithoutVersion() *SchemeFile {
	m.Version = nil
	return m
}

func (m *SchemeFile) WithCreatedOn(createdOn strfmt.DateTime) *SchemeFile {

	m.CreatedOn = createdOn

	return m
}

func (m *SchemeFile) WithModifiedOn(modifiedOn strfmt.DateTime) *SchemeFile {

	m.ModifiedOn = modifiedOn

	return m
}

func (m *SchemeFile) WithAttributes(attributes SchemeFileAttributes) *SchemeFile {

	m.Attributes = &attributes

	return m
}

func (m *SchemeFile) WithoutAttributes() *SchemeFile {
	m.Attributes = nil
	return m
}

func (m *SchemeFile) WithRelationships(relationships SchemeFileRelationships) *SchemeFile {

	m.Relationships = &relationships

	return m
}

func (m *SchemeFile) WithoutRelationships() *SchemeFile {
	m.Relationships = nil
	return m
}

// Validate validates this scheme file
func (m *SchemeFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganisationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemeFile) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

var schemeFileTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["scheme_files"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schemeFileTypeTypePropEnum = append(schemeFileTypeTypePropEnum, v)
	}
}

const (

	// SchemeFileTypeSchemeFiles captures enum value "scheme_files"
	SchemeFileTypeSchemeFiles string = "scheme_files"
)

// prop value enum
func (m *SchemeFile) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, schemeFileTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SchemeFile) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFile) validateOrganisationID(formats strfmt.Registry) error {

	if err := validate.Required("organisation_id", "body", m.OrganisationID); err != nil {
		return err
	}

	if err := validate.FormatOf("organisation_id", "body", "uuid", m.OrganisationID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFile) validateVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if err := validate.MinimumInt("version", "body", int64(*m.Version), 0, false); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFile) validateCreatedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("created_on", "body", "date-time", m.CreatedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFile) validateModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("modified_on", "body", "date-time", m.ModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFile) validateAttributes(formats strfmt.Registry) error {

	if err := validate.Required("attributes", "body", m.Attributes); err != nil {
		return err
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attributes")
			}
			return err
		}
	}

	return nil
}

func (m *SchemeFile) validateRelationships(formats strfmt.Registry) error {

	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	if m.Relationships != nil {
		if err := m.Relationships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemeFile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemeFile) UnmarshalBinary(b []byte) error {
	var res SchemeFile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
func (m *SchemeFile) Json() string {
	json, err := json.MarshalIndent(m, "  ", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}

// SchemeFileAttributes scheme file attributes
// swagger:model SchemeFileAttributes
type SchemeFileAttributes struct {

	// The format of the file that will be submitted to Form3
	// Required: true
	// Enum: [xml]
	FileFormat *string `json:"file_format"`

	// Hashed content of the file
	// Required: true
	// Min Length: 1
	FileHash *string `json:"file_hash"`

	// The size of the file to be uploaded - number of bytes. Max size is 77MB
	// Required: true
	// Maximum: 8.0740352e+07
	FileSize *int64 `json:"file_size"`

	// The file type
	// Required: true
	// Enum: [switch]
	FileType *string `json:"file_type"`

	// The algorithm used to generate the signature
	// Required: true
	// Enum: [SHA256]
	HashingAlgorithm *string `json:"hashing_algorithm"`

	// The count of chunks to be uploaded to the resource
	// Required: true
	// Minimum: 1
	NumberOfParts *int64 `json:"number_of_parts"`

	// Scheme/gateway that the file is to be processed by
	// Required: true
	// Enum: [CASS]
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

	if err := validate.Required("attributes"+"."+"file_format", "body", m.FileFormat); err != nil {
		return err
	}

	// value enum
	if err := m.validateFileFormatEnum("attributes"+"."+"file_format", "body", *m.FileFormat); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFileAttributes) validateFileHash(formats strfmt.Registry) error {

	if err := validate.Required("attributes"+"."+"file_hash", "body", m.FileHash); err != nil {
		return err
	}

	if err := validate.MinLength("attributes"+"."+"file_hash", "body", string(*m.FileHash), 1); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFileAttributes) validateFileSize(formats strfmt.Registry) error {

	if err := validate.Required("attributes"+"."+"file_size", "body", m.FileSize); err != nil {
		return err
	}

	if err := validate.MaximumInt("attributes"+"."+"file_size", "body", int64(*m.FileSize), 8.0740352e+07, false); err != nil {
		return err
	}

	return nil
}

var schemeFileAttributesTypeFileTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["switch"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schemeFileAttributesTypeFileTypePropEnum = append(schemeFileAttributesTypeFileTypePropEnum, v)
	}
}

const (

	// SchemeFileAttributesFileTypeSwitch captures enum value "switch"
	SchemeFileAttributesFileTypeSwitch string = "switch"
)

// prop value enum
func (m *SchemeFileAttributes) validateFileTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, schemeFileAttributesTypeFileTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SchemeFileAttributes) validateFileType(formats strfmt.Registry) error {

	if err := validate.Required("attributes"+"."+"file_type", "body", m.FileType); err != nil {
		return err
	}

	// value enum
	if err := m.validateFileTypeEnum("attributes"+"."+"file_type", "body", *m.FileType); err != nil {
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

	if err := validate.Required("attributes"+"."+"hashing_algorithm", "body", m.HashingAlgorithm); err != nil {
		return err
	}

	// value enum
	if err := m.validateHashingAlgorithmEnum("attributes"+"."+"hashing_algorithm", "body", *m.HashingAlgorithm); err != nil {
		return err
	}

	return nil
}

func (m *SchemeFileAttributes) validateNumberOfParts(formats strfmt.Registry) error {

	if err := validate.Required("attributes"+"."+"number_of_parts", "body", m.NumberOfParts); err != nil {
		return err
	}

	if err := validate.MinimumInt("attributes"+"."+"number_of_parts", "body", int64(*m.NumberOfParts), 1, false); err != nil {
		return err
	}

	return nil
}

var schemeFileAttributesTypePaymentSchemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CASS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schemeFileAttributesTypePaymentSchemePropEnum = append(schemeFileAttributesTypePaymentSchemePropEnum, v)
	}
}

const (

	// SchemeFileAttributesPaymentSchemeCASS captures enum value "CASS"
	SchemeFileAttributesPaymentSchemeCASS string = "CASS"
)

// prop value enum
func (m *SchemeFileAttributes) validatePaymentSchemeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, schemeFileAttributesTypePaymentSchemePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SchemeFileAttributes) validatePaymentScheme(formats strfmt.Registry) error {

	if err := validate.Required("attributes"+"."+"payment_scheme", "body", m.PaymentScheme); err != nil {
		return err
	}

	// value enum
	if err := m.validatePaymentSchemeEnum("attributes"+"."+"payment_scheme", "body", *m.PaymentScheme); err != nil {
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
