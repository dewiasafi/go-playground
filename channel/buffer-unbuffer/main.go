package main

import "fmt"

func main() {
	// un := make(chan int) // unbuffer channel
	// un <- 10             // deadlock karena tidak ada receivernya

	buff := make(chan int, 2)    // buffer channel
	buff <- 40                   // kapasitas 1 terpakai
	buff <- 20                   // kapasitas 2 terpakai, aman karena buff adalah buffer channel
	value := <-buff              // kapasitas buff diambil
	fmt.Println("value ", value) // yang muncul disini adalah 40
	fmt.Println("buff ", <-buff) // yang muncul disini adalah 20

	buff <- 30 // deadlock jika kapasitas buff melebih 2

	/*
		Ingat Channel di Go itu FIFO (First In First Out)
		Artinya: yang masuk duluan → keluar duluan
		jadi ketika <-buff itu diambil data pertama yg masuk yaitu 40
	*/

}
