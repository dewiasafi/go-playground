package main

import (
	"encoding/json"
	"fmt"
)

func ObjectToJson() {
	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println("=> Object to Json")
	fmt.Println(jsonString)
}
