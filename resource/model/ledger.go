package model

type Ledger struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description,omitempty"`
	Currency    string              `json:"currency"`
	Accounts    map[string]*Account `json:"accounts"`
	Entries     []Entry             `json:"entries"`
}

// NewLedger creates a new Account
func NewLedger(id, name, currency, desc string) *Ledger {
	l := new(Ledger)

	l.ID = id
	l.Name = name
	l.Description = desc
	l.Currency = currency
	l.Accounts = map[string]*Account{}
	l.Entries = Entries{}

	return l
}

// AddAccount adds an account to the ledger
func (l *Ledger) AddAccount(account *Account) *Ledger {
	l.Accounts[account.ID] = account
	return l
}
