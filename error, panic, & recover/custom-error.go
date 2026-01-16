package main

import (
	"errors"
	"fmt"
	"strings"
)

func Validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

func main() {
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := Validate(name); valid {
		fmt.Println("halo", name)
	} else {
		fmt.Println(err.Error())
	}
}
