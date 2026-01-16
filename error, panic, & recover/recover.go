package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}

//Recover berguna untuk meng-handle panic error. Pada saat panic error muncul, recover men-take-over goroutine yang sedang panic dan efek sampingnya pesan panic tidak muncul dan eksekusi program adalah tidak error.

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured: ", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

//Untuk menggunakan recover, fungsi/closure/IIFE di mana recover() berada harus dieksekusi dengan cara di-defer.

func main() {
	defer catch()

	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}
