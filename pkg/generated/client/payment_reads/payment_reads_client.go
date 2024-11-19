// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package payment_reads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/v6/pkg/client"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new payment reads API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for payment reads API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
get payment API
*/
func (a *GetPaymentRequest) Do() (*GetPaymentOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPayment",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentOK), nil

}

func (a *GetPaymentRequest) MustDo() *GetPaymentOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment admission API
*/
func (a *GetPaymentAdmissionRequest) Do() (*GetPaymentAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentAdmissionOK), nil

}

func (a *GetPaymentAdmissionRequest) MustDo() *GetPaymentAdmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment admission task API
*/
func (a *GetPaymentAdmissionTaskRequest) Do() (*GetPaymentAdmissionTaskOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentAdmissionTask",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/admissions/{admissionId}/tasks/{taskId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentAdmissionTaskReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentAdmissionTaskOK), nil

}

func (a *GetPaymentAdmissionTaskRequest) MustDo() *GetPaymentAdmissionTaskOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment advice API
*/
func (a *GetPaymentAdviceRequest) Do() (*GetPaymentAdviceOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentAdvice",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/advices/{adviceId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentAdviceReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentAdviceOK), nil

}

func (a *GetPaymentAdviceRequest) MustDo() *GetPaymentAdviceOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment advice submission API
*/
func (a *GetPaymentAdviceSubmissionRequest) Do() (*GetPaymentAdviceSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentAdviceSubmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/advices/{adviceId}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentAdviceSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentAdviceSubmissionOK), nil

}

func (a *GetPaymentAdviceSubmissionRequest) MustDo() *GetPaymentAdviceSubmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment recall API
*/
func (a *GetPaymentRecallRequest) Do() (*GetPaymentRecallOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRecall",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentRecallReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentRecallOK), nil

}

func (a *GetPaymentRecallRequest) MustDo() *GetPaymentRecallOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment recall admission API
*/
func (a *GetPaymentRecallAdmissionRequest) Do() (*GetPaymentRecallAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRecallAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentRecallAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentRecallAdmissionOK), nil

}

func (a *GetPaymentRecallAdmissionRequest) MustDo() *GetPaymentRecallAdmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment recall decision API
*/
func (a *GetPaymentRecallDecisionRequest) Do() (*GetPaymentRecallDecisionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRecallDecision",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/decisions/{decisionId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentRecallDecisionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentRecallDecisionOK), nil

}

func (a *GetPaymentRecallDecisionRequest) MustDo() *GetPaymentRecallDecisionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment recall decision admission API
*/
func (a *GetPaymentRecallDecisionAdmissionRequest) Do() (*GetPaymentRecallDecisionAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRecallDecisionAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/decisions/{decisionId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentRecallDecisionAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentRecallDecisionAdmissionOK), nil

}

func (a *GetPaymentRecallDecisionAdmissionRequest) MustDo() *GetPaymentRecallDecisionAdmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment recall decision submission API
*/
func (a *GetPaymentRecallDecisionSubmissionRequest) Do() (*GetPaymentRecallDecisionSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRecallDecisionSubmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/decisions/{decisionId}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentRecallDecisionSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentRecallDecisionSubmissionOK), nil

}

func (a *GetPaymentRecallDecisionSubmissionRequest) MustDo() *GetPaymentRecallDecisionSubmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment recall reversal API
*/
func (a *GetPaymentRecallReversalRequest) Do() (*GetPaymentRecallReversalOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRecallReversal",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/reversals/{reversalId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentRecallReversalReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentRecallReversalOK), nil

}

func (a *GetPaymentRecallReversalRequest) MustDo() *GetPaymentRecallReversalOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment recall reversal admission API
*/
func (a *GetPaymentRecallReversalAdmissionRequest) Do() (*GetPaymentRecallReversalAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRecallReversalAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/reversals/{reversalId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentRecallReversalAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentRecallReversalAdmissionOK), nil

}

func (a *GetPaymentRecallReversalAdmissionRequest) MustDo() *GetPaymentRecallReversalAdmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment recall submission API
*/
func (a *GetPaymentRecallSubmissionRequest) Do() (*GetPaymentRecallSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRecallSubmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentRecallSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentRecallSubmissionOK), nil

}

func (a *GetPaymentRecallSubmissionRequest) MustDo() *GetPaymentRecallSubmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment return API
*/
func (a *GetPaymentReturnRequest) Do() (*GetPaymentReturnOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentReturn",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/returns/{returnId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReturnReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentReturnOK), nil

}

func (a *GetPaymentReturnRequest) MustDo() *GetPaymentReturnOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment return admission API
*/
func (a *GetPaymentReturnAdmissionRequest) Do() (*GetPaymentReturnAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentReturnAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/returns/{returnId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReturnAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentReturnAdmissionOK), nil

}

func (a *GetPaymentReturnAdmissionRequest) MustDo() *GetPaymentReturnAdmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment return reversal API
*/
func (a *GetPaymentReturnReversalRequest) Do() (*GetPaymentReturnReversalOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentReturnReversal",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/returns/{returnId}/reversals/{reversalId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReturnReversalReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentReturnReversalOK), nil

}

func (a *GetPaymentReturnReversalRequest) MustDo() *GetPaymentReturnReversalOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment return reversal admission API
*/
func (a *GetPaymentReturnReversalAdmissionRequest) Do() (*GetPaymentReturnReversalAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentReturnReversalAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/returns/{returnId}/reversals/{reversalId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReturnReversalAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentReturnReversalAdmissionOK), nil

}

func (a *GetPaymentReturnReversalAdmissionRequest) MustDo() *GetPaymentReturnReversalAdmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment return submission API
*/
func (a *GetPaymentReturnSubmissionRequest) Do() (*GetPaymentReturnSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentReturnSubmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/returns/{returnId}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReturnSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentReturnSubmissionOK), nil

}

func (a *GetPaymentReturnSubmissionRequest) MustDo() *GetPaymentReturnSubmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment reversal API
*/
func (a *GetPaymentReversalRequest) Do() (*GetPaymentReversalOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentReversal",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/reversals/{reversalId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReversalReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentReversalOK), nil

}

func (a *GetPaymentReversalRequest) MustDo() *GetPaymentReversalOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment reversal admission API
*/
func (a *GetPaymentReversalAdmissionRequest) Do() (*GetPaymentReversalAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentReversalAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/reversals/{reversalId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReversalAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentReversalAdmissionOK), nil

}

func (a *GetPaymentReversalAdmissionRequest) MustDo() *GetPaymentReversalAdmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment reversal submission API
*/
func (a *GetPaymentReversalSubmissionRequest) Do() (*GetPaymentReversalSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentReversalSubmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/reversals/{reversalId}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReversalSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentReversalSubmissionOK), nil

}

func (a *GetPaymentReversalSubmissionRequest) MustDo() *GetPaymentReversalSubmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment submission API
*/
func (a *GetPaymentSubmissionRequest) Do() (*GetPaymentSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentSubmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentSubmissionOK), nil

}

func (a *GetPaymentSubmissionRequest) MustDo() *GetPaymentSubmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get payment submission task API
*/
func (a *GetPaymentSubmissionTaskRequest) Do() (*GetPaymentSubmissionTaskOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentSubmissionTask",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/submissions/{submissionId}/tasks/{taskId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentSubmissionTaskReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentSubmissionTaskOK), nil

}

func (a *GetPaymentSubmissionTaskRequest) MustDo() *GetPaymentSubmissionTaskOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get positions API
*/
func (a *GetPositionsRequest) Do() (*GetPositionsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPositions",
		Method:             "GET",
		PathPattern:        "/organisation/positions",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPositionsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPositionsOK), nil

}

func (a *GetPositionsRequest) MustDo() *GetPositionsOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get return submission task API
*/
func (a *GetReturnSubmissionTaskRequest) Do() (*GetReturnSubmissionTaskOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetReturnSubmissionTask",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{paymentId}/returns/{returnId}/submissions/{returnSubmissionId}/tasks/{taskId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetReturnSubmissionTaskReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReturnSubmissionTaskOK), nil

}

func (a *GetReturnSubmissionTaskRequest) MustDo() *GetReturnSubmissionTaskOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get reversal admission task API
*/
func (a *GetReversalAdmissionTaskRequest) Do() (*GetReversalAdmissionTaskOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetReversalAdmissionTask",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/reversals/{reversalId}/admissions/{admissionId}/tasks/{taskId}",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetReversalAdmissionTaskReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetReversalAdmissionTaskOK), nil

}

func (a *GetReversalAdmissionTaskRequest) MustDo() *GetReversalAdmissionTaskOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
list payment admission tasks API
*/
func (a *ListPaymentAdmissionTasksRequest) Do() (*ListPaymentAdmissionTasksOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListPaymentAdmissionTasks",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/admissions/{admissionId}/tasks",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListPaymentAdmissionTasksReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListPaymentAdmissionTasksOK), nil

}

func (a *ListPaymentAdmissionTasksRequest) MustDo() *ListPaymentAdmissionTasksOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
list payments API
*/
func (a *ListPaymentsRequest) Do() (*ListPaymentsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListPayments",
		Method:             "GET",
		PathPattern:        "/transaction/payments",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListPaymentsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListPaymentsOK), nil

}

func (a *ListPaymentsRequest) MustDo() *ListPaymentsOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
list return submission tasks API
*/
func (a *ListReturnSubmissionTasksRequest) Do() (*ListReturnSubmissionTasksOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListReturnSubmissionTasks",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{paymentId}/returns/{returnId}/submissions/{returnSubmissionId}/tasks",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListReturnSubmissionTasksReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListReturnSubmissionTasksOK), nil

}

func (a *ListReturnSubmissionTasksRequest) MustDo() *ListReturnSubmissionTasksOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
list reversal admission tasks API
*/
func (a *ListReversalAdmissionTasksRequest) Do() (*ListReversalAdmissionTasksOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListReversalAdmissionTasks",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/reversals/{reversalId}/admissions/{admissionId}/tasks",
		ProducesMediaTypes: []string{"application/vnd.api+json", "application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListReversalAdmissionTasksReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListReversalAdmissionTasksOK), nil

}

func (a *ListReversalAdmissionTasksRequest) MustDo() *ListReversalAdmissionTasksOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/////////

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
