package main

import (
	"fmt"
	"github.com/form3tech-oss/go-form3/pkg/form3"
)

func main() {

	f3 := form3.New()
	units := f3.Subscriptions.ListSubscriptions().MustDo()
	for _, unit := range units.Data {
		fmt.Printf("Deleting subscription %s to %s\n", unit.ID, *unit.Attributes.CallbackURI)
		f3.Subscriptions.DeleteSubscription().
			WithID(*unit.ID).
			WithVersion(*unit.Version).
			MustDo()

	}
}
