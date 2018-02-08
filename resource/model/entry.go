package model

import "time"

type Entry struct {
	ID          string      `json:"id"`
	DateTime    time.Time   `json:"time"`
	Description string      `json:"description"`
	References  []Reference `json:"references"`
	Amount      Amount      `json:"amount"`
}

func NewEntry(id string, time time.Time, description string, amount Amount) Entry {
	e := Entry{
		ID:          id,
		DateTime:    time,
		Description: description,
		References:  []Reference{},
		Amount:      amount,
	}

	return e
}

func (e Entry) AddReference(reference ...Reference) Entry {
	e.References = append(e.References, reference...)
	return e
}
