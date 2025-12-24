package main

import "fmt"

// Value receiver -- hanya baca data struct
func (b Barang) ReadData() {
	fmt.Printf("%-15s | Rp %10.2f | Stok: %d\n", b.Nama, b.Harga, b.Stok)
}

// Pointer receiver -- mengubah data asli
func (b *Barang) TambahStok(jumlah int) {
	b.Stok += jumlah
	fmt.Printf("--> Stok %s berhasil ditambah %d\n", b.Nama, jumlah)
}

// ini function biasa yg mengembalikan nilai total
func HitungAset(daftar []Barang) float64 {
	total := 0.0
	for _, b := range daftar {
		total += b.Harga * float64(b.Stok)
	}
	return total
}
