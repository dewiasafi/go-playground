package main

import (
	"fmt"
	"time"
)

func doSomethingAsyncKatanya(ch chan int) int {
	fmt.Println("DOING SOMETHING By doSomethingAsyncKatanya ...")
	time.Sleep(2 * time.Second)
	fmt.Println("ASYNC By doSomethingAsyncKatanya DONE")
	ch <- 1
	return <-ch
}

func doSomethingAsync() chan int {
	ch := make(chan int)
	go func() {
		fmt.Println("DOING SOMETHING By doSomethingAsync...")
		time.Sleep(2 * time.Second)
		fmt.Println("ASYNCRONOUS By doSomethingAsync DONE")
		ch <- 1
	}()
	return ch
}

func doSometing() int {
	fmt.Println("DOING SOMETHING By doSomething...")
	time.Sleep(2 * time.Second)
	fmt.Println("By doSomething DONE")
	return 1
}

func doSomethingAsyncronous() chan int {
	ch := make(chan int)

	go func() {
		ch <- doSometing()
	}()

	return ch
}

func main() {
	// eksekusi a, tunggu selesai
	a := doSometing()
	b := doSometing()
	// tunggu a selesai, baru eksekusi
	fmt.Println("Tanpa goroutine + channel: ", a, b) // berjalan secara sinkron

	// Kalau mau proses doSomething() jalan secara asyncronous bisa pakai `go`
	// Tapi karena doSomething() punya return, tidak bisa langsung seperti itu
	// harus ada channel untuk simpan nilai return nya

	c := make(chan int)
	d := make(chan int)
	go doSomethingAsyncKatanya(c)
	go doSomethingAsyncKatanya(d)
	fmt.Println("DO 1: ", <-c, <-d)
	fmt.Println("---")

	e := doSomethingAsync()
	f := doSomethingAsync()
	fmt.Println("DO 2: ", <-e, <-f)
	fmt.Println("---")

	// Bisa juga menggunakan func return biasa jadi async function
	g := doSomethingAsyncronous()
	h := doSomethingAsyncronous()
	fmt.Println("DO 3: ", <-g, h)
	fmt.Println("---")

}
