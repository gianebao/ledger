package model_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gianebao/ledger/resource/model"
)

func ExampleNewLedger() {
	l := model.NewLedger("DW001", "DigiWallet", "DWC", "Sample digital wallet system")

	l.AddAccount(model.NewAccount("AC00001"))
	l.AddAccount(model.NewAccount("AC00002"))

	l.Accounts["AC00001"].AddEntry(model.NewEntry("E0001", time.Unix(1517386944, 0), "Initialize account", 500))
	l.Accounts["AC00001"].Entries[0].References["receipt-number"] = "xx123456789"

	j, err := json.MarshalIndent(l, "", "\t")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(j))
	// Output:
	// 	{
	// 	"id": "DW001",
	// 	"name": "DigiWallet",
	// 	"currency": "DWC",
	// 	"description": "Sample digital wallet system",
	// 	"accounts": [
	// 		"AC00001",
	// 		"AC00002"
	// 	],
	// 	"entries": [
	// 		{
	// 			"id": "E0001",
	// 			"time": "2018-01-31T16:22:24+08:00",
	// 			"description": "Initialize account",
	// 			"references": {
	// 				"account": "AC00001",
	// 				"receipt-number": "xx123456789"
	// 			},
	// 			"amount": 500.0000
	// 		}
	// 	]
	// }
}

func ExampleLedger_LoadCSV() {
	l := model.NewLedger("DW001", "DigiWallet", "DWC", "Sample digital wallet system")

	content, err := ioutil.ReadFile("sample-data/ledger-20180902.csv")
	if err != nil {
		panic(err)
	}

	if err = l.LoadCSV(string(content)); err != nil {
		panic(err)
	}

	fmt.Printf("Balance for account AC00001: %.4f\n", l.Accounts["AC00001"].Entries.Balance())
	fmt.Printf("Balance for account AC00002: %.4f\n", l.Accounts["AC00002"].Entries.Balance())
	fmt.Printf("Balance for account AC00003: %.4f\n", l.Accounts["AC00003"].Entries.Balance())
	fmt.Printf("Balance for account AC00004: %.4f\n", l.Accounts["AC00004"].Entries.Balance())
	fmt.Printf("Balance for account AC00005: %.4f\n", l.Accounts["AC00005"].Entries.Balance())

	l = model.NewLedger("DW002", "DigiWallet2", "DC2", "Sample digital wallet system version 2")
	if err = l.LoadCSV("id,account,amount,date\n" +
		"XX0010000001,AC00001,1000,2018-02-09T08:38:56+08:00"); err != nil {
		panic(err.Error())
	}

	fmt.Printf("[%s] txn: %s %s -- %.2f",
		l.Accounts["AC00001"].Entries[0].DateTime,
		l.Accounts["AC00001"].Entries[0].ID,
		l.Accounts["AC00001"].Entries[0].References["account"],
		l.Accounts["AC00001"].Entries[0].Amount)
	// Output:
	// Balance for account AC00001: 322.3500
	// Balance for account AC00002: 337.4900
	// Balance for account AC00003: 819.4000
	// Balance for account AC00004: 1039.0100
	// Balance for account AC00005: 976.2500
	// [2018-02-09 08:38:56 +0800 +08] txn: XX0010000001 AC00001 -- 1000.00
}
