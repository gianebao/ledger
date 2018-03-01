package model_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"testing"
	"time"

	"github.com/gianebao/ledger/resource/model"
	"github.com/stretchr/testify/assert"
)

func ExampleNewLedger() {
	l := model.NewLedger("DW001", "DigiWallet", "DWC", "Sample digital wallet system")

	l.AddAccount(model.NewAccount("AC00001"))
	l.AddAccount(model.NewAccount("AC00002"))

	// Setting entry timezone
	loc, _ := time.LoadLocation("Asia/Singapore")

	l.Accounts["AC00001"].AddEntry(model.NewEntry("E0001", time.Unix(1517386944, 0).In(loc), "Initialize account", 500))
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

	for i := 1; i <= 5; i++ {
		fmt.Printf("Balance for account AC0000%d: %.4f\n", i, l.Accounts["AC0000"+strconv.Itoa(i)].Entries.Balance())
	}

	l = model.NewLedger("DW002", "DigiWallet2", "DC2", "Sample digital wallet system version 2")
	if err = l.LoadCSV("id,account,amount,date\n" +
		"XX0010000001,AC00001,1000,2018-02-09T08:38:56+08:00"); err != nil {
		panic(err.Error())
	}

	fmt.Printf("[%s] txn: %s %s -- %.2f",
		l.Accounts["AC00001"].Entries[0].DateTime.Format(model.DefaultCSVTimeFormat),
		l.Accounts["AC00001"].Entries[0].ID,
		l.Accounts["AC00001"].Entries[0].References["account"],
		l.Accounts["AC00001"].Entries[0].Amount)
	// Output:
	// Balance for account AC00001: 322.3500
	// Balance for account AC00002: 337.4900
	// Balance for account AC00003: 819.4000
	// Balance for account AC00004: 1039.0100
	// Balance for account AC00005: 976.2500
	// [2018-02-09T08:38:56+08:00] txn: XX0010000001 AC00001 -- 1000.00
}

func TestLedger_LoadCSV_Edgecases(t *testing.T) {
	l := model.NewLedger("DW001", "DigiWallet", "DWC", "Sample digital wallet system")
	err := l.LoadCSV("id,account,amount,date\n" +
		"XX0010000001")

	// Test for invalid CSV
	assert.Error(t, err)

	err = l.LoadCSV("id,account,amount,date\n")

	// Test for CSV with just the header row
	assert.Nil(t, err)

	err = l.LoadCSV("id,account,amount,date\n" +
		"XX0010000001,AC00001,1000,2018-0a-09T08:38:56+08:00")

	// Test for invalid entry
	assert.EqualError(t, err, "Error parsing [date] row: 1 parsing time \"2018-0a-09T08:38:56+08:00\": month out of range")

	err = l.LoadCSV("id,account,amount,date\n" +
		",AC00001,1000,2018-08-09T08:38:56+08:00")

	// Test for invalid entry
	assert.EqualError(t, err, "Error parsing file at line 1: `id` and `account` cannot be empty")
}
