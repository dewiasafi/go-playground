package main

import (
	"fmt"
	"strings"
)

// Ini adalah project sederhana untuk sistem kasir untuk hitung diskon
func SimpleProject() {
	var nama string
	var totalBelanja float64
	var isMember bool
	var inputMember string

	fmt.Print("Masukan nama pelanggan: ")
	// %s untuk membaca string (sampai ketemu spasi)
	fmt.Scanf("%s\n", &nama)

	fmt.Print("Masukkan Total Belanja: ")
	// %f untuk membaca angka desimal (float)
	fmt.Scanf("%f\n", &totalBelanja)

	fmt.Print("Apakah member? (Y/N): ")
	// %t untuk membaca boolean
	fmt.Scanf("%t\n", &inputMember)

	if strings.ToUpper(inputMember) == "Y" {
		isMember = true
	} else {
		isMember = false
	}

	totalBayar, hemat := hitungTotalBayar(totalBelanja, isMember)

	fmt.Println("\n--- Hasil Perhitungan ---")
	fmt.Printf("Pelanggan: %s\n", strings.Title(strings.ToLower(nama))) // ubah dulu nama jadi lowerCase, baru buat jadi kapital diawal
	fmt.Printf("Total Belanja: Rp %.2f \nTotal Bayar: Rp %.2f (Hemat: %.2f)\n", totalBelanja, totalBayar, hemat)
}

func hitungTotalBayar(harga float64, member bool) (float64, float64) {
	diskon := 0.0
	if member && harga > 500000 {
		diskon = 0.1
	}
	if harga > 500000 {
		diskon += 0.05
	}

	potonganHarga := harga * diskon
	return harga - potonganHarga, potonganHarga
}
