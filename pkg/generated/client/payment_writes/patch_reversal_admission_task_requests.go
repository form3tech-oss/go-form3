// Code generated by go-swagger; DO NOT EDIT.

package payment_writes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/form3tech-oss/go-form3/v6/pkg/generated/models"
)

// Client.PatchReversalAdmissionTask creates a new PatchReversalAdmissionTaskRequest object
// with the default values initialized.
func (c *Client) PatchReversalAdmissionTask() *PatchReversalAdmissionTaskRequest {
	var ()
	return &PatchReversalAdmissionTaskRequest{

		ReversalAdmissionTaskAmendment: models.ReversalAdmissionTaskAmendmentWithDefaults(c.Defaults),

		AdmissionID: c.Defaults.GetStrfmtUUID("PatchReversalAdmissionTask", "admissionId"),

		ID: c.Defaults.GetStrfmtUUID("PatchReversalAdmissionTask", "id"),

		ReversalID: c.Defaults.GetStrfmtUUID("PatchReversalAdmissionTask", "reversalId"),

		TaskID: c.Defaults.GetStrfmtUUID("PatchReversalAdmissionTask", "taskId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type PatchReversalAdmissionTaskRequest struct {

	/*ReversalAdmissionTaskPatchRequest*/

	*models.ReversalAdmissionTaskAmendment

	/*AdmissionID      Admission Id      */

	AdmissionID strfmt.UUID

	/*ID      Payment Id      */

	ID strfmt.UUID

	/*ReversalID      Reversal Id      */

	ReversalID strfmt.UUID

	/*TaskID      Reversal Admission Task Id      */

	TaskID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *PatchReversalAdmissionTaskRequest) FromJson(j string) (*PatchReversalAdmissionTaskRequest, error) {

	var m models.ReversalAdmissionTaskAmendment
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %w", err)
	}

	o.ReversalAdmissionTaskAmendment = &m

	return o, nil
}

func (o *PatchReversalAdmissionTaskRequest) WithReversalAdmissionTaskPatchRequest(reversalAdmissionTaskPatchRequest models.ReversalAdmissionTaskAmendment) *PatchReversalAdmissionTaskRequest {

	o.ReversalAdmissionTaskAmendment = &reversalAdmissionTaskPatchRequest

	return o
}

func (o *PatchReversalAdmissionTaskRequest) WithoutReversalAdmissionTaskPatchRequest() *PatchReversalAdmissionTaskRequest {

	o.ReversalAdmissionTaskAmendment = &models.ReversalAdmissionTaskAmendment{}

	return o
}

func (o *PatchReversalAdmissionTaskRequest) WithAdmissionID(admissionID strfmt.UUID) *PatchReversalAdmissionTaskRequest {

	o.AdmissionID = admissionID

	return o
}

func (o *PatchReversalAdmissionTaskRequest) WithID(id strfmt.UUID) *PatchReversalAdmissionTaskRequest {

	o.ID = id

	return o
}

func (o *PatchReversalAdmissionTaskRequest) WithReversalID(reversalID strfmt.UUID) *PatchReversalAdmissionTaskRequest {

	o.ReversalID = reversalID

	return o
}

func (o *PatchReversalAdmissionTaskRequest) WithTaskID(taskID strfmt.UUID) *PatchReversalAdmissionTaskRequest {

	o.TaskID = taskID

	return o
}

// ////////////////
// WithContext adds the context to the patch reversal admission task Request
func (o *PatchReversalAdmissionTaskRequest) WithContext(ctx context.Context) *PatchReversalAdmissionTaskRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the patch reversal admission task Request
func (o *PatchReversalAdmissionTaskRequest) WithHTTPClient(client *http.Client) *PatchReversalAdmissionTaskRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *PatchReversalAdmissionTaskRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// ISBODYPARAM
	if o.ReversalAdmissionTaskAmendment != nil {
		if err := r.SetBodyParam(o.ReversalAdmissionTaskAmendment); err != nil {
			return err
		}
	}

	// path param admissionId
	if err := r.SetPathParam("admissionId", o.AdmissionID.String()); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param reversalId
	if err := r.SetPathParam("reversalId", o.ReversalID.String()); err != nil {
		return err
	}

	// path param taskId
	if err := r.SetPathParam("taskId", o.TaskID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
