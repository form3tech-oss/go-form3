package form3_test

import (
	"testing"

	"github.com/form3tech-oss/go-form3/v7/pkg/generated/client/payment_writes"
	"github.com/go-openapi/runtime"
	"github.com/stretchr/testify/assert"
)

// Ensure the OpenAPI runtime.ClientResponseStatus interface is being met by our generated code.
// Since we use custom templates, this does not happen out-of-the-box & we must update our templates to meet these.
func Test_ResponseInterface(t *testing.T) {
	response := payment_writes.NewCreatePaymentCreated()
	assert.Implements(t, (*runtime.ClientResponseStatus)(nil), response)
}
