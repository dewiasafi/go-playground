package main

import "fmt"

func main() {
	a := 1
	// 'a' menyimpan nilai 1.
	fmt.Println("Nilai a:", a)

	// '&a' mengambil alamat memori tempat 1 disimpan.
	fmt.Println("Alamat memori a:", &a)

	// b adalah variabel tipe pointer (*int) yang menyimpan alamat a.
	b := &a
	fmt.Println("Isi variabel b (alamat a):", b)
	fmt.Println("Nilai yang ditunjuk b (dereference):", *b)

	fmt.Println("--- Eksekusi Fungsi ---")
	fmt.Println("Alamat asli a di main:", &a)

	// printA menerima alamat, jadi alamatnya AKAN SAMA dengan di main.
	printA(&a)

	// printB menerima nilai, Go akan membuat salinan di alamat baru.
	// Jadi alamatnya AKAN BERBEDA dengan di main.
	printB(a)
	fmt.Println("Nilai setelah di increment di printA", a)
}

// Menerima pointer ke int
func printA(a *int) {
	(*a)++
	fmt.Println("Alamat a di func printA (Pass by Reference):", a)
}

// Menerima nilai int biasa
func printB(a int) {
	a += 2
	// Karena di-copy, &a di sini adalah alamat memori baru milik fungsi ini saja.
	fmt.Println("Alamat a di func printB (Pass by Value):", &a)
}
