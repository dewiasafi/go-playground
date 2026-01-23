package main

import "fmt"

func main() {
	total1 := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println("total int:", total1)

	total2 := Sum([]float32{2.5, 7.2})
	fmt.Println("total float32:", total2)

	total3 := Sum([]float64{1.23, 6.33, 12.6})
	fmt.Println("total float64:", total3)
}
