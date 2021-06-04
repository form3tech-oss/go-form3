package form3_test

import (
	"io/ioutil"
	"testing"

	"github.com/form3tech-oss/go-form3/v2/pkg/form3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryPayments(t *testing.T) {
	f3, err := form3.NewFromEnv()
	require.NoError(t, err)

	p, err := f3.Payments.ListPayments().
		WithFilterAmount("2500.00").
		WithoutFilterOrganisationID().
		Do()
	assert.Nil(t, err)
	assert.NotEmpty(t, p.Data)
}

func TestPaymentFromJsonFile(t *testing.T) {
	f3, err := form3.NewFromEnv()
	require.NoError(t, err)

	j, err := ioutil.ReadFile("../../test/payment.json")
	require.NoError(t, err)

	p, err := f3.Payments.CreatePayment().FromJson(string(j))
	require.NoError(t, err)
	assert.Equal(t, "603.00", p.Data.Attributes.Amount)
}

func TestPaymentToJson(t *testing.T) {
	f3, err := form3.NewFromEnv()
	require.NoError(t, err)

	j, err := ioutil.ReadFile("../../test/payment.json")
	require.NoError(t, err)

	p, err := f3.Payments.CreatePayment().FromJson(string(j))
	require.NoError(t, err)

	p2, err := f3.Payments.CreatePayment().FromJson(p.Json())
	assert.Equal(t, p, p2)
}
