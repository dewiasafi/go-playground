package main

import (
	"fmt"
	"runtime"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func main() {
	runtime.GOMAXPROCS(2) //digunakan untuk menentukan jumlah core yang diaktifkan untuk eksekusi program

	go print(3, "Hello")      // go didepan print ini artinya function print dijalankan di thread lain
	go print(3, "Apa kabar?") // print yg ini dijalankan didalam thread main

	var input string
	fmt.Scanln(&input) // Blocking agar proses diselesaikan ketika user klik enter
}
