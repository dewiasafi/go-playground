package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	doneChan := make(chan struct{})

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- 1
		close(c1)
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- 2
		close(c2)
	}()

	go func() {
		time.Sleep(3 * time.Second) // waktu tunggunya harus lebih besar dari yg lain
		close(doneChan)
	}()

	for { // loop agar bisa terus menunggu event dari banyak channel
		select { // mengeksekusi case channel yang siap lebih dulu
		case data, ok := <-c1:
			if !ok {
				c1 = nil
				continue
			}
			fmt.Println("AMBIL DARI CHANNEL C1: ", data)
		case data, ok := <-c2:
			if !ok {
				c1 = nil
				continue
			}
			fmt.Println("AMBIL DARI CHANNEL C2: ", data)
		case <-doneChan:
			fmt.Println("SELESAI")
			return // keluar dari fungsi jadi nya selesai, makanya delaynya harus lebih besar supaya c1 dan c2 selesai dulu
		}
	}
}
