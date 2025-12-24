package main

import "fmt"

// kalo ada function yg dari luar file main.go
// jalankan nya bukan go run main.go, tapi go run .
func main() {
	sayHello()
	fmt.Println(sumCalculator(10, 32))
	fmt.Println(bagi(10, 4))
	fmt.Println(hitungLuas(5, 10))
	fmt.Println(jumlahkanSemua(1, 2, 3, 4, 5))
	testing()

	// Penulisan nama function sangat berpengaruh pada scope function
	// Jika nama function diawali dengan huruf kapital scopenya public, tapi jika diawali dengan huruf kecil scope nya private
	SimpleProject() // simple project ini function dari file simpleProject.go

}

// basic function
func sayHello() {
	fmt.Println("Hello Guys")
}

//function parameter
func sumCalculator(a int, b int) int {
	return a + b
}

// function dengan 2 atau lebih return
func bagi(a, b float32) (float32, string) {
	if b == 0 {
		return 0, "Error: Tidak bisa membagi dengan nol"
	}
	return a / b, "Success" // pastikan tipe data a dan b sama dengan tipe data return
}

// function dengan return yg sudah dinamai
func hitungLuas(panjang, lebar int) (luas int) {
	luas = panjang * lebar
	return
}

// function dengan parameter yg tidak terbatas
// jumlah parameter ini bisa beda-beda di setiap pemanggilan function.
// tapi tipe datanya harus sama
func jumlahkanSemua(angka ...int) int {
	total := 0
	for _, n := range angka {
		total += n
	}
	return total
}

// function tanpa nama
// function yang bisa dibuat di dalam function lain atau langsung dijalankan
// sering dipakai untuk closure atau goroutines
func testing() {
	// function langsung dijalankan (IIFE)
	func(pesan string) {
		fmt.Println("test: ", pesan)
	}("Halo dari anonymous function!")
}
