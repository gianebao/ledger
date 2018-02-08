package model_test

import (
	"testing"
	"time"

	"github.com/gianebao/ledger/resource/model"
	"github.com/stretchr/testify/assert"
)

func TestEntries_Balance(t *testing.T) {
	test := assert.New(t)
	a := model.NewAccount("x0001")

	a.AddEntry(model.NewEntry("t0001", time.Unix(1517386944, 0), "Initial transaction load", 10000))
	a.AddEntry(model.NewEntry("t0002", time.Unix(1517387044, 0), "Purchased - Some thing", -300))
	a.AddEntry(model.NewEntry("t0003", time.Unix(1517387144, 0), "Top up - Paymuch", +75.05))

	test.Equal(9775.05, float64(a.Entries.Balance()))
}
