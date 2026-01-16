package main

import (
	"encoding/json"
	"fmt"
)

// Contoh cara decoding json ke bentuk objek.

func JsonToStruct() {

	var data User
	// Proses decode json string menggunakan json.Unmarshal()
	var err = json.Unmarshal(JsonData, &data) // jsonData harus dalam bentuk []byte

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("=> TO Struct")

	fmt.Println("user :", data.FullName)
	fmt.Println("age  :", data.Age)
}
