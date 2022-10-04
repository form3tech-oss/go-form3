// Code generated by go-swagger; DO NOT EDIT.

package scheme_file_api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// Client.CreateSchemeFile creates a new CreateSchemeFileRequest object
// with the default values initialized.
func (c *Client) CreateSchemeFile() *CreateSchemeFileRequest {
	var ()
	return &CreateSchemeFileRequest{

		SchemeFileCreation: models.SchemeFileCreationWithDefaults(c.Defaults),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type CreateSchemeFileRequest struct {

	/*SchemeFileCreationRequest*/

	*models.SchemeFileCreation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *CreateSchemeFileRequest) FromJson(j string) (*CreateSchemeFileRequest, error) {

	var m models.SchemeFileCreation
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.SchemeFileCreation = &m

	return o, nil
}

func (o *CreateSchemeFileRequest) WithSchemeFileCreationRequest(schemeFileCreationRequest models.SchemeFileCreation) *CreateSchemeFileRequest {

	o.SchemeFileCreation = &schemeFileCreationRequest

	return o
}

func (o *CreateSchemeFileRequest) WithoutSchemeFileCreationRequest() *CreateSchemeFileRequest {

	o.SchemeFileCreation = &models.SchemeFileCreation{}

	return o
}

//////////////////
// WithContext adds the context to the create scheme file Request
func (o *CreateSchemeFileRequest) WithContext(ctx context.Context) *CreateSchemeFileRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the create scheme file Request
func (o *CreateSchemeFileRequest) WithHTTPClient(client *http.Client) *CreateSchemeFileRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *CreateSchemeFileRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.SchemeFileCreation != nil {
		if err := r.SetBodyParam(o.SchemeFileCreation); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
