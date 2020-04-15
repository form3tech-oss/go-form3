package client

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type ClientDefaults struct {
	EmptyDefaults
	OrganisationId *strfmt.UUID
}

func (d *ClientDefaults) GetStrfmtUUID(objectName, attributeName string) strfmt.UUID {
	var a strfmt.UUID
	if "organisation_id" == attributeName && d.OrganisationId != nil {
		a = *d.OrganisationId
	} else if "id" == attributeName {
		a = strfmt.UUID(uuid.Must(uuid.NewRandom()).String())
	}
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

func NewClientDefaults() *ClientDefaults {
	return &ClientDefaults{
		EmptyDefaults: *NewEmptyDefaults(),
	}
}
