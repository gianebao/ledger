package model_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gianebao/ledger/resource/model"
)

func ExampleAccount_AddEntry() {
	a := model.NewAccount("x0001")

	a.AddEntry(model.NewEntry("t0001", time.Unix(1517386944, 0), "Initial transaction load", 10000))

	j, err := json.MarshalIndent(a, "", "\t")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(string(j))
	// Output:
	// 	{
	// 	"id": "x0001",
	// 	"entries": [
	// 		{
	// 			"id": "t0001",
	// 			"time": "2018-01-31T16:22:24+08:00",
	// 			"description": "Initial transaction load",
	// 			"references": [],
	// 			"amount": 10000.0000
	// 		}
	// 	]
	// }
}
