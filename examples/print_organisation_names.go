package main

import (
	"fmt"

	"github.com/form3tech-oss/go-form3/v2/pkg/form3"
)

func main() {
	f3, _ := form3.NewFromEnv()

	units := f3.Organisations.ListUnits().MustDo()
	for _, unit := range units.Data {
		fmt.Printf("%s - %s\n", unit.ID, unit.Attributes.Name)
	}
}
