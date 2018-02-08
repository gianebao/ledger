package model

type Marshaler interface {
	Amount() ([]byte, error)
}
