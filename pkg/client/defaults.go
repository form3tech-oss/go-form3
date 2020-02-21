package client

import "github.com/go-openapi/strfmt"

type Defaults interface {
	GetStrfmtUUID(objectName, attributeName string) strfmt.UUID
	GetStrfmtUUIDPtr(objectName, attributeName string) *strfmt.UUID

	GetStrfmtDateTime(objectName, attributeName string) strfmt.DateTime
	GetStrfmtDateTimePtr(objectName, attributeName string) *strfmt.DateTime

	GetStrfmtDate(objectName, attributeName string) strfmt.Date
	GetStrfmtDatePtr(objectName, attributeName string) *strfmt.Date

	GetStrfmtURI(objectName, attributeName string) strfmt.URI
	GetStrfmtURIPtr(objectName, attributeName string) *strfmt.URI

	GetString(objectName, attributeName string) string
	GetStringArray(objectName, attributeName string) []string
	GetStringPtr(objectName, attributeName string) *string

	GetBool(objectName, attributeName string) bool
	GetBoolPtr(objectName, attributeName string) *bool

	GetInt64(objectName, attributeName string) int64
	GetInt64Ptr(objectName, attributeName string) *int64

	GetInt32(objectName, attributeName string) int32

	GetFloat64(objectName, attibuteName string) float64
	GetFloat64Ptr(objectName, attibuteName string) *float64
}
