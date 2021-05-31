package main

import (
	"fmt"
	"time"

	"github.com/form3tech-oss/go-form3/v2/pkg/form3"
	"github.com/go-openapi/strfmt"
)

// Demonstrating different ways of creating payment resources
func main() {
	f3, _ := form3.NewFromEnv()
	{
		fmt.Println("Create using defaults: ")

		p := f3.Payments.CreatePayment()
		fmt.Println(p.Json())
	}
	{
		fmt.Println("Import from JSON: ")

		p, _ := f3.Payments.CreatePayment().FromJson("{\n  \"data\": {\n    \"type\": \"payments\",\n    \"id\": \"{{$guid}}\",\n    \"version\": 0,\n     \"attributes\": {\n      \"amount\": \"603.00\",\n      \"beneficiary_party\": {\n        \"account_name\": \"Mrs Receiving Test\",\n        \"account_number\": \"12345678\",\n        \"account_number_code\": \"BBAN\",\n        \"account_type\": 1,\n        \"account_with\": {\n          \"bank_id\": \"123456\",\n          \"bank_id_code\": \"GBDSC\" },\n        \"country\": \"GB\" },\n      \"currency\": \"GBP\",\n      \"debtor_party\": {\n        \"account_name\": \"Mr Sending Test\",\n        \"account_number\": \"87654321\",\n        \"account_number_code\": \"BBAN\",\n        \"account_with\": {\n          \"bank_id\": \"333333\",\n          \"bank_id_code\": \"GBDSC\" },\n        \"country\": \"GB\" },\n      \"numeric_reference\": \"0001\",\n         \"reference\": \"ref2\",\n      \"scheme_payment_type\": \"Credit\",\n      \"payment_scheme\": \"Bacs\"\n    }\n  }\n}")

		fmt.Println(p.Json())
	}
	{
		fmt.Println("Create using defaults + customisation: ")

		p := f3.Payments.CreatePayment()
		p.Data.Attributes.
			WithAmount("200").
			WithCurrency("GBP").
			WithReference("D/1234567890123456").
			WithNumericReference("0001").
			WithSchemePaymentType("ImmediatePayment").
			WithPaymentScheme("FPS").
			WithProcessingDate(strfmt.Date(time.Now())).
			WithoutFx()

		p.Data.Attributes.BeneficiaryParty.
			WithAccountName("Mrs Receiving Test").
			WithAccountNumber("12345678").
			WithAccountNumberCode("BBAN").
			WithAccountType(1).
			WithCountry("GB")

		p.Data.Attributes.BeneficiaryParty.AccountWith.
			WithBankID("112233").
			WithBankIDCode("GBDSC")

		fmt.Println(p.Json())
	}
}
