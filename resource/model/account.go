package model

// Account represent a financial account detail based on a general transaction ledger
type Account struct {
	ID           string  `json:"id"`
	Entries      Entries `json:"entries"`
	entryHandler func(flag string, e *Entry)
}

const (
	// AccountReferenceName defines the name for account as reference to an entry
	AccountReferenceName = "account"
)

// NewAccount creates a new account
func NewAccount(id string) *Account {
	a := new(Account)

	a.ID = id
	a.Entries = Entries{}
	a.entryHandler = func(flag string, e *Entry) {}

	return a
}

// AddEntry appends a new entry to Account
func (a *Account) AddEntry(entry ...Entry) *Account {
	for _, e := range entry {
		e.References[AccountReferenceName] = a.ID
		a.Entries = append(a.Entries, e)
		a.entryHandler("add", &e)
	}

	return a
}
