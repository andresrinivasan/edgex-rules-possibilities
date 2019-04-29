package main

import (
	"encoding/json"
	"fmt"

	"github.com/diegoholiveira/jsonlogic"
)

func main() {
	var logic interface{}
	var data interface{}
	var result interface{}

	_ = json.Unmarshal([]byte(`{
		"filter": [
			{ "var": "events" },
			{ "and": 
			  { "=": [{ "var": "device" },"562114e9e4b0385849b96cd8"],
					">": [{ "var": "RPM" }, 1200]
				}
			}
		]
	}`), &logic)

	_ = json.Unmarshal([]byte(`{
		"events": [
			{"device": "562114e9e4b0385849b96cd8", "RPM": 1200},
			{"device": "562114e9e4b0385849b96cd8", "RPM": 1201},
			{"device": "111", "RPM": 1200},
			{"device": "222", "RPM": 1201}
		]
	}`), &data)

	err := jsonlogic.Apply(logic, data, &result)
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	fmt.Println("Too fast")
	for _, _user := range result.([]interface{}) {
		user := _user.(map[string]interface{})

		fmt.Printf("    - %v\n", user)
	}
}

//  {
//    "name": "motortoofastsignal",
//    "condition": {
//      "device": "562114e9e4b0385849b96cd8",
//      "checks": [
//        {
//          "parameter": "RPM",
//          "operand1": "Integer.parseInt(value)",
//          "operation": ">",
//          "operand2": "1200"
//        }
//      ]
//    },
//    "action": {
//      "device": "56325f7ee4b05eaae5a89ce1",
//      "command": "56325f6de4b05eaae5a89cdc",
//      "body": "{\\\"value\\\":\\\"3\\\"}"
//    },
//    "log": "Patlite warning triggered for engine speed too high"
//  }
