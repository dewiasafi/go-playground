package main

import "fmt"

func sendOnly(c chan<- int) {
	go func() {
		c <- 10
	}()
}

func receiveOnly(c <-chan int) {
	fmt.Println("Value from Receive Only Channel: ", <-c)
}

func main() {
	c := make(chan int)
	sendOnly(c)
	receiveOnly(c)
}
