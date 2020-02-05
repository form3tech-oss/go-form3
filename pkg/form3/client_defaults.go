package form3

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type ClientDefaults struct {
	OrganisationId *strfmt.UUID
}

func (*ClientDefaults) GetStrfmtUUID(objectName, attributeName string) strfmt.UUID {
	var a strfmt.UUID
	return a
}

func (d *ClientDefaults) GetStrfmtUUIDPtr(objectName, attributeName string) *strfmt.UUID {
	if "organisation_id" == attributeName {
		return d.OrganisationId
	}
	if "id" == attributeName {
		u := strfmt.UUID(uuid.Must(uuid.NewRandom()).String())
		return &u
	}
	return nil
}

func (*ClientDefaults) GetStrfmtDateTime(objectName, attributeName string) strfmt.DateTime {
	var a strfmt.DateTime
	return a
}

func (*ClientDefaults) GetStrfmtDateTimePtr(objectName, attributeName string) *strfmt.DateTime {
	return nil
}

func (*ClientDefaults) GetStrfmtDate(objectName, attributeName string) strfmt.Date {
	var a strfmt.Date
	return a
}

func (*ClientDefaults) GetStrfmtDatePtr(objectName, attributeName string) *strfmt.Date {
	return nil
}

func (*ClientDefaults) GetStrfmtURI(objectName, attributeName string) strfmt.URI {
	var a strfmt.URI
	return a
}

func (*ClientDefaults) GetStrfmtURIPtr(objectName, attributeName string) *strfmt.URI {
	return nil
}

func (*ClientDefaults) GetString(objectName, attributeName string) string {
	return ""
}

func (*ClientDefaults) GetStringArray(objectName, attributeName string) []string {
	var a []string
	return a
}

func (*ClientDefaults) GetStringPtr(objectName, attributeName string) *string {
	return nil
}

func (*ClientDefaults) GetBool(objectName, attributeName string) bool {
	return true
}

func (*ClientDefaults) GetBoolPtr(objectName, attributeName string) *bool {
	return nil
}

func (*ClientDefaults) GetInt64(objectName, attributeName string) int64 {
	return int64(44)
}

func (*ClientDefaults) GetInt64Ptr(objectName, attributeName string) *int64 {
	return nil
}

func (*ClientDefaults) GetInt32(objectName, attributeName string) int32 {
	return int32(44)
}

func NewClientDefaults() *ClientDefaults {
	return &ClientDefaults{}
}
