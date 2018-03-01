package model_test

import (
	"fmt"

	"github.com/gianebao/ledger/resource/model"
)

func ExampleNewAmountFromString() {
	a, _ := model.NewAmountFromString("1000.06")

	fmt.Printf("Amount: %s\n", a)

	_, err := model.NewAmountFromString("a")

	fmt.Printf("Error: %v\n", err)
	// Output:
	// Amount: 1000.0600
	// Error: model.NewAmountFromString: strconv.ParseFloat: parsing "a": invalid syntax
}
