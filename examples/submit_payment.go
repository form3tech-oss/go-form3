package main

import (
	"fmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/form3"
)

func main() {
	f3, _ := form3.NewFromEnv()

	p := f3.Payments.CreatePayment()
	p.Data.WithType("payments")
	p.Data.Attributes.Reference = "123"
	p.Data.Attributes.
		WithAmount("200").
		WithReference("Testing")
	created, err := p.Do()

	fmt.Println(err)

	response := f3.Payments.CreatePaymentSubmission().
		WithID(*created.Data.ID).MustDo()
	fmt.Printf("%+v", response)
}
