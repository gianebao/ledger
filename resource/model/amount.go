package model

import (
	"bytes"
	"fmt"
	"strconv"
)

// Amount represents an amount value in a ledger entry
type Amount float64

// AmountDisplayFormat defines the format of the amount when displayed in JSON
var AmountDisplayFormat = "%.4f"

// String returns a formatted string
func (a Amount) String() string {
	return fmt.Sprintf(AmountDisplayFormat, a)
}

// MarshalJSON renders the amount according to `AmountDisplayFormat`
func (a Amount) MarshalJSON() ([]byte, error) {
	return bytes.NewBufferString(a.String()).Bytes(), nil
}

// NewAmountFromString creates a new amount from a `string`
func NewAmountFromString(value string) (Amount, error) {
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}

	return Amount(v), nil
}
