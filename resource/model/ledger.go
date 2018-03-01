package model

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"strings"
)

// Ledger represents the transaction ledger
type Ledger struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description,omitempty"`
	Currency    string              `json:"currency"`
	Accounts    map[string]*Account `json:"accounts"`
	entries     []*Entry
}

// MarshalJSON renders a simplified
func (l Ledger) MarshalJSON() ([]byte, error) {
	sAppend := func(a []string, name, value string) []string {
		return append(a, fmt.Sprintf("\"%s\":\"%s\"", name, value))
	}
	b := bytes.NewBufferString("{")
	s := []string{}

	// top level entries
	s = sAppend(s, "id", l.ID)
	s = sAppend(s, "name", l.Name)
	s = sAppend(s, "currency", l.Currency)

	if l.Description != "" {
		s = sAppend(s, "description", l.Description)
	}

	// accounts
	sa := []string{}
	for _, i := range l.Accounts {
		sa = append(sa, i.ID)
	}

	sj, _ := json.Marshal(sa)
	s = append(s, fmt.Sprintf("\"accounts\":%s", string(sj)))

	// entries
	sj, _ = json.Marshal(l.entries)
	s = append(s, fmt.Sprintf("\"entries\":%s", string(sj)))

	b.WriteString(strings.Join(s, ",") + "}")
	return b.Bytes(), nil
}

// NewLedger creates a new Account
func NewLedger(id, name, currency, desc string) *Ledger {
	l := new(Ledger)

	l.ID = id
	l.Name = name
	l.Description = desc
	l.Currency = currency
	l.Accounts = map[string]*Account{}
	l.entries = make([]*Entry, 0)

	return l
}

// AddAccount adds an account to the ledger
func (l *Ledger) AddAccount(account *Account) *Ledger {
	l.Accounts[account.ID] = account
	l.Accounts[account.ID].entryHandler = func(operation string, e *Entry) {
		if operation == "add" {
			l.entries = append(l.entries, e)
		}
	}
	return l
}

// LoadCSV loads information from a CSV formatted string
func (l *Ledger) LoadCSV(s string) error {
	r := csv.NewReader(strings.NewReader(s))
	entries, err := r.ReadAll()

	if err != nil {
		return fmt.Errorf("Error in parsing csv: %s", err)
	}

	// At most, just the header row
	if len(entries) < 2 {
		return nil
	}

	fields := entries[0]
	colFunc := getColumnParser(fields)

	for i := 1; i < len(entries); i++ {
		e := Entry{References: Reference{}}
		entry := entries[i]
		for j := 0; j < len(fields); j++ {
			if err = colFunc[j](&e, entry[j]); err != nil {
				return fmt.Errorf("Error parsing [%s] row: %d %v", fields[j], i, err)
			}
		}

		if _, ok := e.References[AccountReferenceName]; e.ID == "" || !ok {
			return fmt.Errorf("Error parsing file at line %d: `id` and `account` cannot be empty", i)
		}
		account := e.References[AccountReferenceName]
		// fmt.Println(e)
		if _, ok := l.Accounts[account]; !ok {
			l.AddAccount(NewAccount(account))
		}

		l.Accounts[account].AddEntry(e)
	}

	return nil
}
