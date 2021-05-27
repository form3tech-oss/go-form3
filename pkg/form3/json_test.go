package form3_test

import (
	"testing"

	"github.com/form3tech-oss/go-form3/v2/pkg/form3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJsonConversion(t *testing.T) {
	f3, err := form3.NewFromEnv()
	require.NoError(t, err)

	p := f3.Payments.CreatePayment().FromJson("{\n  \"data\": {\n    \"type\": \"payments\",\n    \"id\": \"758ea5e2-3725-48c7-8ef1-2a6d045f2037\",\n    \"version\": 0,\n     \"attributes\": {\n      \"amount\": \"603.00\",\n      \"beneficiary_party\": {\n        \"account_name\": \"Mrs Receiving Test\",\n        \"account_number\": \"12345678\",\n        \"account_number_code\": \"BBAN\",\n        \"account_type\": 1,\n        \"account_with\": {\n          \"bank_id\": \"123456\",\n          \"bank_id_code\": \"GBDSC\" },\n        \"country\": \"GB\" },\n      \"currency\": \"GBP\",\n      \"debtor_party\": {\n        \"account_name\": \"Mr Sending Test\",\n        \"account_number\": \"87654321\",\n        \"account_number_code\": \"BBAN\",\n        \"account_with\": {\n          \"bank_id\": \"333333\",\n          \"bank_id_code\": \"GBDSC\" },\n        \"country\": \"GB\" },\n      \"numeric_reference\": \"0001\",\n         \"reference\": \"ref2\",\n      \"scheme_payment_type\": \"Credit\",\n      \"payment_scheme\": \"Bacs\"\n    }\n  }\n}")
	assert.Equal(t, "603.00", p.Data.Attributes.Amount)
	p2 := f3.Payments.CreatePayment().FromJson(p.Json())
	assert.Equal(t, "603.00", p2.Data.Attributes.Amount)
}
