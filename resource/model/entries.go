package model

// Entries represents a group of ledger entries that share the same currency
type Entries []Entry

// Balance calculates the total running balance of the ledger
func (es Entries) Balance() Amount {
	var total Amount

	total = 0

	for _, e := range es {
		total = total + e.Amount
	}

	return total
}
