package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int)
	// go func() {
	// 	ch <- 100
	// }()
	// hasil := <-ch
	// fmt.Println(hasil)
	c := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		c <- 100
	}()
	value := <-c
	fmt.Println("Value from channel c in main: ", value)

}
