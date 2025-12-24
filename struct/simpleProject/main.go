package main

import (
	"fmt"
	"strings"
)

type Barang struct {
	Nama   string
	Harga  float64
	Jumlah int
}

func main() {
	var daftarBelanja []Barang // Ini adalah Slice untuk menyimpan banyak barang
	var isMember bool
	var inputMember string
	var namaMember string

	fmt.Println("=== SISTEM KASIR GO ===")
	fmt.Print("Nama Pelanggan: ")
	fmt.Scanln(&namaMember)

	// Perulangan untuk input barang belanjaan
	for {
		var namaBarang string
		var hargaBarang float64

		fmt.Printf("\nMasukkan nama barang (ketik 'done' untuk hitung): ")
		fmt.Scanln(&namaBarang)

		if namaBarang == "done" {
			break
		}

		fmt.Printf("Masukkan harga %s: ", namaBarang)
		fmt.Scanln(&hargaBarang)

		// Masukkan ke dalam daftar belanja
		itemBaru := Barang{Nama: namaBarang, Harga: hargaBarang}
		daftarBelanja = append(daftarBelanja, itemBaru)
	}

	fmt.Print("\nApakah punya kartu member?")
	fmt.Scanln(&inputMember)

	if strings.ToUpper(inputMember) == "Y" {
		isMember = true
	} else {
		isMember = false
	}
	cetakStruk(namaMember, daftarBelanja, isMember)
}

func cetakStruk(pelanggan string, barang []Barang, member bool) {
	var totalAwal float64
	fmt.Println("\n===============================")
	fmt.Printf("Struk Belanja: %s\n", pelanggan)
	fmt.Println("-------------------------------")

	for _, item := range barang {
		fmt.Printf("%-15s : Rp %10.2f\n", item.Nama, item.Harga)
		totalAwal += item.Harga
	}

	DiskonPersen(member, totalAwal)
}
