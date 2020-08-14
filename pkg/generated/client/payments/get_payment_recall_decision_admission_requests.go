// Code generated by go-swagger; DO NOT EDIT.

package payments

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// Client.GetPaymentRecallDecisionAdmission creates a new GetPaymentRecallDecisionAdmissionRequest object
// with the default values initialized.
func (c *Client) GetPaymentRecallDecisionAdmission() *GetPaymentRecallDecisionAdmissionRequest {
	var ()
	return &GetPaymentRecallDecisionAdmissionRequest{

		AdmissionID: c.Defaults.GetStrfmtUUID("GetPaymentRecallDecisionAdmission", "admissionId"),

		DecisionID: c.Defaults.GetStrfmtUUID("GetPaymentRecallDecisionAdmission", "decisionId"),

		ID: c.Defaults.GetStrfmtUUID("GetPaymentRecallDecisionAdmission", "id"),

		RecallID: c.Defaults.GetStrfmtUUID("GetPaymentRecallDecisionAdmission", "recallId"),

		timeout: cr.DefaultTimeout,

		transport: c.transport,
		formats:   c.formats,
	}
}

type GetPaymentRecallDecisionAdmissionRequest struct {

	/*AdmissionID      Admission Id      */

	AdmissionID strfmt.UUID

	/*DecisionID      Decision Id      */

	DecisionID strfmt.UUID

	/*ID      Payment Id      */

	ID strfmt.UUID

	/*RecallID      Recall Id      */

	RecallID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client

	transport runtime.ClientTransport
	formats   strfmt.Registry
}

func (o *GetPaymentRecallDecisionAdmissionRequest) FromJson(j string) *GetPaymentRecallDecisionAdmissionRequest {

	return o
}

func (o *GetPaymentRecallDecisionAdmissionRequest) WithAdmissionID(admissionID strfmt.UUID) *GetPaymentRecallDecisionAdmissionRequest {

	o.AdmissionID = admissionID

	return o
}

func (o *GetPaymentRecallDecisionAdmissionRequest) WithDecisionID(decisionID strfmt.UUID) *GetPaymentRecallDecisionAdmissionRequest {

	o.DecisionID = decisionID

	return o
}

func (o *GetPaymentRecallDecisionAdmissionRequest) WithID(id strfmt.UUID) *GetPaymentRecallDecisionAdmissionRequest {

	o.ID = id

	return o
}

func (o *GetPaymentRecallDecisionAdmissionRequest) WithRecallID(recallID strfmt.UUID) *GetPaymentRecallDecisionAdmissionRequest {

	o.RecallID = recallID

	return o
}

//////////////////
// WithContext adds the context to the get payment recall decision admission Request
func (o *GetPaymentRecallDecisionAdmissionRequest) WithContext(ctx context.Context) *GetPaymentRecallDecisionAdmissionRequest {
	o.Context = ctx
	return o
}

// WithHTTPClient adds the HTTPClient to the get payment recall decision admission Request
func (o *GetPaymentRecallDecisionAdmissionRequest) WithHTTPClient(client *http.Client) *GetPaymentRecallDecisionAdmissionRequest {
	o.HTTPClient = client
	return o
}

// WriteToRequest writes these Request to a swagger request
func (o *GetPaymentRecallDecisionAdmissionRequest) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param admissionId
	if err := r.SetPathParam("admissionId", o.AdmissionID.String()); err != nil {
		return err
	}

	// path param decisionId
	if err := r.SetPathParam("decisionId", o.DecisionID.String()); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// path param recallId
	if err := r.SetPathParam("recallId", o.RecallID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
