package model

import (
	"bytes"
	"fmt"
)

// Amount represents an amount value in a ledger entry
type Amount float64

// AmountDisplayFormat defines the format of the amount when displayed in JSON
var AmountDisplayFormat = "%.4f"

// MarshalJSON renders the amount according to `AmountDisplayFormat`
func (a Amount) MarshalJSON() ([]byte, error) {
	return bytes.NewBufferString(fmt.Sprintf(AmountDisplayFormat, a)).Bytes(), nil
}
