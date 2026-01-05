package main

import "fmt"

// Ini adalah Kontraknya
type Pembayaran interface {
	Bayar(jumlah int) // Siapapun yang mau jadi metode pembayaran, WAJIB punya fungsi Bayar
}

// Struct OVO
type OVO struct{}

func (o OVO) Bayar(jumlah int) {
	fmt.Printf("Membayar %d menggunakan OVO\n", jumlah)
}

// Struct Bank
type Bank struct{}

func (b Bank) Bayar(jumlah int) {
	fmt.Printf("Membayar %d via Transfer Bank\n", jumlah)
}

// Fungsi ini sekarang bisa menerima APAPUN selama dia punya method Bayar()
func ProsesTransaksi(p Pembayaran, harga int) {
	p.Bayar(harga)
}
