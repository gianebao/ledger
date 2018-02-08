package model

// Account represent a financial account detail based on a general transaction ledger
type Account struct {
	ID      string  `json:"id"`
	Entries Entries `json:"entries"`
}

// NewAccount creates a new account
func NewAccount(id string) *Account {
	a := new(Account)

	a.ID = id
	a.Entries = Entries{}

	return a
}

// AddEntry appends a new entry to Account
func (a *Account) AddEntry(entry ...Entry) *Account {
	a.Entries = append(a.Entries, entry...)
	return a
}
