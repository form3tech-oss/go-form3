package client

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type ClientDefaults struct {
	EmptyDefaults
	OrganisationID *strfmt.UUID
}

func (d *ClientDefaults) GetStrfmtUUID(objectName, attributeName string) strfmt.UUID {
	var a strfmt.UUID
	if attributeName == "organisation_id" && d.OrganisationID != nil {
		a = *d.OrganisationID
	} else if attributeName == "id" {
		a = strfmt.UUID(uuid.Must(uuid.NewRandom()).String())
	}

	return a
}

func (d *ClientDefaults) GetStrfmtUUIDPtr(objectName, attributeName string) *strfmt.UUID {
	if attributeName == "organisation_id" {
		return d.OrganisationID
	}

	if attributeName == "id" {
		u := strfmt.UUID(uuid.Must(uuid.NewRandom()).String())
		return &u
	}

	return nil
}

func NewClientDefaults() *ClientDefaults {
	return &ClientDefaults{
		EmptyDefaults: *NewEmptyDefaults(),
	}
}
