package main

import (
	"fmt"

	"github.com/form3tech-oss/go-form3/pkg/form3"
	"github.com/form3tech-oss/go-form3/pkg/generated/models"

	. "github.com/ahmetb/go-linq"
)

func main() {
	f3 := form3.NewFromEnv()

	payments := f3.Payments.ListPayments().
		WithFilterAmount("2500.00").
		WithoutFilterOrganisationID().
		MustDo()

	for _, p := range payments.Data {
		fmt.Printf("%s - %s - %s\n", p.ID, p.Attributes.Amount, p.Attributes.ProcessingDate)
	}

	// or fancier options with go-linq
	From(payments.Data).
		OrderByDescendingT(func(p *models.Payment) string { return p.Attributes.Amount }).
		ForEachT(func(p *models.Payment) {
			fmt.Printf("%s - %s - %s\n", p.ID, p.Attributes.Amount, p.Attributes.ProcessingDate)
		})

	From(payments.Data).
		WhereT(func(p *models.Payment) bool { return len(p.Relationships.PaymentSubmission.Data) == 1 }).
		WhereT(func(p *models.Payment) bool {
			return p.Relationships.PaymentSubmission.Data[0].Attributes.Status == models.PaymentSubmissionStatusDeliveryConfirmed
		}).
		SelectT(func(p *models.Payment) string { return p.Attributes.Amount }).
		SumInts()
}
