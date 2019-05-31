package main

import (
	"fmt"
	"github.com/form3tech-oss/go-form3/pkg/form3"
)

func main() {

	f3 := form3.New()
	units := f3.Subscriptions.ListSubscriptions().MustDo()
	for _, sub := range units.Data {
		fmt.Printf("%s - %s - %s\n", sub.ID, *sub.Attributes.RecordType, *sub.Attributes.CallbackURI)
	}
}
