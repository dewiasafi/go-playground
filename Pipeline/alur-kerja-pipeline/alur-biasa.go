package main

import (
	"fmt"
	"time"
)

func AlurBiasa() {
	nums := []int{1, 2, 3, 4}
	start := time.Now()

	duration := time.Since(start)
	for _, n := range nums {
		result := n * n
		fmt.Println("Alur biasa: ", result)
	}

	fmt.Println("done in", duration.Seconds(), "seconds")

}
