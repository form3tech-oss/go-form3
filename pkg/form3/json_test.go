package form3_test

import (
	"io/ioutil"
	"testing"

	"github.com/form3tech-oss/go-form3/v2/pkg/form3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaymentFromJsonFile(t *testing.T) {
	f3, err := form3.NewFromEnv()
	require.NoError(t, err)

	j, err := ioutil.ReadFile("payment.json")
	require.NoError(t, err)

	p := f3.Payments.CreatePayment().FromJson(string(j))
	assert.Equal(t, "603.00", p.Data.Attributes.Amount)
}

func TestPaymentToJson(t *testing.T) {
	f3, err := form3.NewFromEnv()
	require.NoError(t, err)

	j, err := ioutil.ReadFile("payment.json")
	require.NoError(t, err)

	p := f3.Payments.CreatePayment().FromJson(string(j))
	p2 := f3.Payments.CreatePayment().FromJson(p.Json())
	assert.Equal(t, p, p2)
}
