package main

import (
	"fmt"
	"github.com/form3tech-oss/go-form3/pkg/form3"
)

func main() {

	f3 := form3.New()

	payment, _ := f3.Payments.GetPayment().
		WithID("87cb3e94-103f-4b91-8c31-c4bea0ee9cf5").
		Do()

	sub := f3.Payments.GetPaymentSubmission().
		WithID(*payment.Data.ID).
		WithSubmissionID(*payment.Data.Relationships.PaymentSubmission.Data[0].ID).MustDo()

	fmt.Printf("%s - %s\n", payment.Data.Attributes.Reference, sub.Data.Attributes.Status)

	audit := f3.Audit.GetAuditEntry().WithRecordType(sub.Data.Type).WithID(*sub.Data.ID).MustDo()

	for _, entry := range audit.Data {
		fmt.Printf("%v\n", entry)
	}
}
