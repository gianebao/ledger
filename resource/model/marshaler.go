package model

// Marshaler defines how marshal works for custom data types
type Marshaler interface {
	Amount() ([]byte, error)
	Ledger() ([]byte, error)
}
