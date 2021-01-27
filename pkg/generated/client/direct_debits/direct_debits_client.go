// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package direct_debits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/v2/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new direct debits API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for direct debits API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
create direct debit API
*/
func (a *CreateDirectDebitRequest) Do() (*CreateDirectDebitCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDirectDebit",
		Method:             "POST",
		PathPattern:        "/transaction/directdebits",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateDirectDebitReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDirectDebitCreated), nil

}

func (a *CreateDirectDebitRequest) MustDo() *CreateDirectDebitCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
create direct debit return API
*/
func (a *CreateDirectDebitReturnRequest) Do() (*CreateDirectDebitReturnCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDirectDebitReturn",
		Method:             "POST",
		PathPattern:        "/transaction/directdebits/{id}/returns",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateDirectDebitReturnReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDirectDebitReturnCreated), nil

}

func (a *CreateDirectDebitReturnRequest) MustDo() *CreateDirectDebitReturnCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
create direct debit return submission API
*/
func (a *CreateDirectDebitReturnSubmissionRequest) Do() (*CreateDirectDebitReturnSubmissionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDirectDebitReturnSubmission",
		Method:             "POST",
		PathPattern:        "/transaction/directdebits/{id}/returns/{returnId}/submissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateDirectDebitReturnSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDirectDebitReturnSubmissionCreated), nil

}

func (a *CreateDirectDebitReturnSubmissionRequest) MustDo() *CreateDirectDebitReturnSubmissionCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
create direct debit reversal API
*/
func (a *CreateDirectDebitReversalRequest) Do() (*CreateDirectDebitReversalCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDirectDebitReversal",
		Method:             "POST",
		PathPattern:        "/transaction/directdebits/{id}/reversals",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateDirectDebitReversalReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDirectDebitReversalCreated), nil

}

func (a *CreateDirectDebitReversalRequest) MustDo() *CreateDirectDebitReversalCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
create direct debit submission API
*/
func (a *CreateDirectDebitSubmissionRequest) Do() (*CreateDirectDebitSubmissionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateDirectDebitSubmission",
		Method:             "POST",
		PathPattern:        "/transaction/directdebits/{id}/submissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreateDirectDebitSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDirectDebitSubmissionCreated), nil

}

func (a *CreateDirectDebitSubmissionRequest) MustDo() *CreateDirectDebitSubmissionCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get direct debit API
*/
func (a *GetDirectDebitRequest) Do() (*GetDirectDebitOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDirectDebit",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetDirectDebitReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDirectDebitOK), nil

}

func (a *GetDirectDebitRequest) MustDo() *GetDirectDebitOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get direct debit admission API
*/
func (a *GetDirectDebitAdmissionRequest) Do() (*GetDirectDebitAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDirectDebitAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetDirectDebitAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDirectDebitAdmissionOK), nil

}

func (a *GetDirectDebitAdmissionRequest) MustDo() *GetDirectDebitAdmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get direct debit return API
*/
func (a *GetDirectDebitReturnRequest) Do() (*GetDirectDebitReturnOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDirectDebitReturn",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}/returns/{returnId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetDirectDebitReturnReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDirectDebitReturnOK), nil

}

func (a *GetDirectDebitReturnRequest) MustDo() *GetDirectDebitReturnOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get direct debit return admission API
*/
func (a *GetDirectDebitReturnAdmissionRequest) Do() (*GetDirectDebitReturnAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDirectDebitReturnAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}/returns/{returnId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetDirectDebitReturnAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDirectDebitReturnAdmissionOK), nil

}

func (a *GetDirectDebitReturnAdmissionRequest) MustDo() *GetDirectDebitReturnAdmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get direct debit return reversal API
*/
func (a *GetDirectDebitReturnReversalRequest) Do() (*GetDirectDebitReturnReversalOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDirectDebitReturnReversal",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}/returns/{returnId}/reversals/{reversalId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetDirectDebitReturnReversalReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDirectDebitReturnReversalOK), nil

}

func (a *GetDirectDebitReturnReversalRequest) MustDo() *GetDirectDebitReturnReversalOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get direct debit return reversal admission API
*/
func (a *GetDirectDebitReturnReversalAdmissionRequest) Do() (*GetDirectDebitReturnReversalAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDirectDebitReturnReversalAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}/returns/{returnId}/reversals/{reversalId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetDirectDebitReturnReversalAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDirectDebitReturnReversalAdmissionOK), nil

}

func (a *GetDirectDebitReturnReversalAdmissionRequest) MustDo() *GetDirectDebitReturnReversalAdmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get direct debit return submission API
*/
func (a *GetDirectDebitReturnSubmissionRequest) Do() (*GetDirectDebitReturnSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDirectDebitReturnSubmission",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}/returns/{returnId}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetDirectDebitReturnSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDirectDebitReturnSubmissionOK), nil

}

func (a *GetDirectDebitReturnSubmissionRequest) MustDo() *GetDirectDebitReturnSubmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get direct debit reversal API
*/
func (a *GetDirectDebitReversalRequest) Do() (*GetDirectDebitReversalOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDirectDebitReversal",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}/reversals/{reversalId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetDirectDebitReversalReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDirectDebitReversalOK), nil

}

func (a *GetDirectDebitReversalRequest) MustDo() *GetDirectDebitReversalOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get direct debit reversal admission API
*/
func (a *GetDirectDebitReversalAdmissionRequest) Do() (*GetDirectDebitReversalAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDirectDebitReversalAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}/reversals/{reversalId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetDirectDebitReversalAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDirectDebitReversalAdmissionOK), nil

}

func (a *GetDirectDebitReversalAdmissionRequest) MustDo() *GetDirectDebitReversalAdmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get direct debit submission API
*/
func (a *GetDirectDebitSubmissionRequest) Do() (*GetDirectDebitSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDirectDebitSubmission",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetDirectDebitSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDirectDebitSubmissionOK), nil

}

func (a *GetDirectDebitSubmissionRequest) MustDo() *GetDirectDebitSubmissionOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get transaction directdebits ID decisions decision ID API
*/
func (a *GetTransactionDirectdebitsIDDecisionsDecisionIDRequest) Do() (*GetTransactionDirectdebitsIDDecisionsDecisionIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionDirectdebitsIDDecisionsDecisionID",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}/decisions/{decisionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionDirectdebitsIDDecisionsDecisionIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionDirectdebitsIDDecisionsDecisionIDOK), nil

}

func (a *GetTransactionDirectdebitsIDDecisionsDecisionIDRequest) MustDo() *GetTransactionDirectdebitsIDDecisionsDecisionIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get transaction directdebits ID decisions decision ID admissions admission ID API
*/
func (a *GetTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsAdmissionIDRequest) Do() (*GetTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsAdmissionIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsAdmissionID",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}/decisions/{decisionId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsAdmissionIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsAdmissionIDOK), nil

}

func (a *GetTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsAdmissionIDRequest) MustDo() *GetTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsAdmissionIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get transaction directdebits ID decisions decision ID submissions submission ID API
*/
func (a *GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest) Do() (*GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionID",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}/decisions/{decisionId}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDOK), nil

}

func (a *GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDRequest) MustDo() *GetTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsSubmissionIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get transaction directdebits ID recalls recall ID API
*/
func (a *GetTransactionDirectdebitsIDRecallsRecallIDRequest) Do() (*GetTransactionDirectdebitsIDRecallsRecallIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionDirectdebitsIDRecallsRecallID",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}/recalls/{recallID}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionDirectdebitsIDRecallsRecallIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionDirectdebitsIDRecallsRecallIDOK), nil

}

func (a *GetTransactionDirectdebitsIDRecallsRecallIDRequest) MustDo() *GetTransactionDirectdebitsIDRecallsRecallIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get transaction directdebits ID recalls recall ID admissions admission ID API
*/
func (a *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest) Do() (*GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionID",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}/recalls/{recallId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDOK), nil

}

func (a *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDRequest) MustDo() *GetTransactionDirectdebitsIDRecallsRecallIDAdmissionsAdmissionIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get transaction directdebits ID recalls recall ID submissions submission ID API
*/
func (a *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest) Do() (*GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionID",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}/recalls/{recallId}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK), nil

}

func (a *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDRequest) MustDo() *GetTransactionDirectdebitsIDRecallsRecallIDSubmissionsSubmissionIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
get transaction directdebits ID reversals reversal ID submissions submission ID API
*/
func (a *GetTransactionDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest) Do() (*GetTransactionDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetTransactionDirectdebitsIDReversalsReversalIDSubmissionsSubmissionID",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits/{id}/reversals/{reversalId}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetTransactionDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK), nil

}

func (a *GetTransactionDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDRequest) MustDo() *GetTransactionDirectdebitsIDReversalsReversalIDSubmissionsSubmissionIDOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
list direct debits API
*/
func (a *ListDirectDebitsRequest) Do() (*ListDirectDebitsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListDirectDebits",
		Method:             "GET",
		PathPattern:        "/transaction/directdebits",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListDirectDebitsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListDirectDebitsOK), nil

}

func (a *ListDirectDebitsRequest) MustDo() *ListDirectDebitsOK {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post transaction directdebits ID decisions API
*/
func (a *PostTransactionDirectdebitsIDDecisionsRequest) Do() (*PostTransactionDirectdebitsIDDecisionsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostTransactionDirectdebitsIDDecisions",
		Method:             "POST",
		PathPattern:        "/transaction/directdebits/{id}/decisions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostTransactionDirectdebitsIDDecisionsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostTransactionDirectdebitsIDDecisionsCreated), nil

}

func (a *PostTransactionDirectdebitsIDDecisionsRequest) MustDo() *PostTransactionDirectdebitsIDDecisionsCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post transaction directdebits ID decisions decision ID admissions API
*/
func (a *PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsRequest) Do() (*PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissions",
		Method:             "POST",
		PathPattern:        "/transaction/directdebits/{id}/decisions/{decisionId}/admissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsCreated), nil

}

func (a *PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsRequest) MustDo() *PostTransactionDirectdebitsIDDecisionsDecisionIDAdmissionsCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post transaction directdebits ID decisions decision ID submissions API
*/
func (a *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest) Do() (*PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissions",
		Method:             "POST",
		PathPattern:        "/transaction/directdebits/{id}/decisions/{decisionId}/submissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsCreated), nil

}

func (a *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsRequest) MustDo() *PostTransactionDirectdebitsIDDecisionsDecisionIDSubmissionsCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post transaction directdebits ID recalls API
*/
func (a *PostTransactionDirectdebitsIDRecallsRequest) Do() (*PostTransactionDirectdebitsIDRecallsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostTransactionDirectdebitsIDRecalls",
		Method:             "POST",
		PathPattern:        "/transaction/directdebits/{id}/recalls",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostTransactionDirectdebitsIDRecallsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostTransactionDirectdebitsIDRecallsCreated), nil

}

func (a *PostTransactionDirectdebitsIDRecallsRequest) MustDo() *PostTransactionDirectdebitsIDRecallsCreated {
	r0, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r0
}

/*
post transaction directdebits ID reversals reversal ID submissions API
*/
func (a *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest) Do() (*PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostTransactionDirectdebitsIDReversalsReversalIDSubmissions",
		Method:             "POST",
		PathPattern:        "/transaction/directdebits/{id}/reversals/{reversalId}/submissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsCreated), nil

}

func (a *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsRequest) MustDo() *PostTransactionDirectdebitsIDReversalsReversalIDSubmissionsCreated {
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
