package main

import (
	"encoding/json"
	"fmt"
)

func JsonToInterface() {
	// decoding data json jadi bervariable map[string]interface{}.
	var data1 map[string]interface{}
	json.Unmarshal(JsonData, &data1)

	fmt.Println("=> TO map[string]interface")
	fmt.Println("user :", data1["Name"])
	fmt.Println("age  :", data1["Age"])

	// decoding json ke interface
	// untuk akses nilai properti, data harus di casting jadi map[string]interface

	var data2 interface{}
	json.Unmarshal(JsonData, &data2)

	decodeData := data2.(map[string]interface{})

	fmt.Println("=> TO interface")
	fmt.Println("user: ", decodeData["Name"])
	fmt.Println("age  :", decodeData["Age"])

}
