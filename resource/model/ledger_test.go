package model_test

import (
	"github.com/gianebao/ledger/resource/model"
)

func ExampleNewLedger() {
	l := model.NewLedger("DW001", "DigiWallet", "DWC", "Sample digital wallet system")

	l.AddAccount(model.NewAccount("AC00001"))
	l.AddAccount(model.NewAccount("AC00002"))
}
