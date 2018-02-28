package model

import (
	"strings"
	"time"
)

// Entry represents a transaction ledger entry
type Entry struct {
	ID          string    `json:"id"`
	DateTime    time.Time `json:"time"`
	Description string    `json:"description"`
	References  Reference `json:"references"`
	Amount      Amount    `json:"amount"`
}

const (
	// DefaultCSVTimeFormat defines the date format when reading a CSV
	DefaultCSVTimeFormat = time.RFC3339

	// DefaultCSVTimezone defines the default timezone
	DefaultCSVTimezone = "Europe/Dublin"
)

// NewEntry creates a new entry
func NewEntry(id string, t time.Time, description string, amount Amount) Entry {
	e := Entry{
		ID:          id,
		DateTime:    t,
		Description: description,
		References:  Reference{},
		Amount:      amount,
	}

	return e
}

// setID sets the value of ID
func setID(e *Entry, v string) error {
	e.ID = v
	return nil
}

// setDescription sets the value of Description
func setDescription(e *Entry, v string) error {
	e.Description = v
	return nil
}

// setAmount sets the value of Amount
func setAmount(e *Entry, v string) error {
	var err error
	e.Amount, err = NewAmountFromString(v)
	return err
}

// setDefaultDate sets the value of DateTime based on `DefaultCSVTimeFormat`
func setDefaultDate(e *Entry, v string) error {
	var err error
	e.DateTime, err = time.Parse(DefaultCSVTimeFormat, v)
	return err
}

// getColumnParser returns a list of functions when parsing CSV
func getColumnParser(fields []string) []func(e *Entry, v string) error {
	var err error
	colFunc := []func(e *Entry, v string) error{}

	for i := 0; i < len(fields); i++ {
		switch fields[i] {
		case "id":
			colFunc = append(colFunc, setID)

		case "description":
			colFunc = append(colFunc, setDescription)

		case "amount":
			colFunc = append(colFunc, setAmount)

		case "date":
			colFunc = append(colFunc, setDefaultDate)

		default:
			if 0 == strings.Index(fields[i], "date ") {
				format := fields[i][6:]
				colFunc = append(colFunc, func(e *Entry, v string) error {
					e.DateTime, err = time.Parse(format, v)
					return err
				})

				continue
			} else {
				field := fields[i]
				colFunc = append(colFunc, func(e *Entry, v string) error {
					e.References[field] = v
					return nil
				})
			}
		}
	}

	return colFunc
}
