package main

import "fmt"

// Struct = benda atau data
// Method = aksi atau perilaku

// Method adalah function receiver
// Cara buat method
// method itu seperti function biasa
// tapi punya receiver (penerima) yang terletak di antara func dan nama function

type Barang struct { // struct barang ini bisa dipakai di file go lain, tapi harus dalam 1 package
	Nama  string
	Harga float64
	Stok  int
}

// ini adalah method
// (b Barang) adalah receivernya
func (b Barang) dataBarang() {
	fmt.Printf("Produk: %s | Harga: Rp%.2f | Stok: %d\n", b.Nama, b.Harga, b.Stok)
}

func main() {
	item := Barang{"Beras", 100000, 20}
	// Panggil method-nya pakai titik (.)
	item.dataBarang()
	// Receiver ada 2 jenis: value dan pointer

	// Value Receiver --> membuat salinan dari struct, jika data diubah di method data asli tidak berubah (hanya baca)
	//ketika passing data ke dalam function go secara default akan mengkopi nilainya dan menyimpan di alamat baru
	// Pointer Receiver --> menggunakan alamat memori aslinya, jika data diubah di method data asli ikut berubah (jika mau edit isi data struct)

	inventaris := []Barang{
		{Nama: "Kopi Arabika", Harga: 25000, Stok: 10},
		{Nama: "Gula Aren", Harga: 15000, Stok: 20},
		{Nama: "Susu UHT", Harga: 18000, Stok: 5},
	}
	fmt.Println("=== DAFTAR INVENTARIS AWAL ===")
	for _, item := range inventaris {
		item.ReadData()
	}

	fmt.Println("\n=== UPDATE STOK ===")
	// Menggunakan & untuk mengambil alamat memori agar bisa diubah Pointer
	// Kita tambah stok Susu UHT (indeks ke-2)
	inventaris[2].TambahStok(5)

	fmt.Println("\n=== TOTAL ASET TOKO ===")
	totalAset := HitungAset(inventaris)
	fmt.Printf("Total nilai barang di gudang: Rp %.2f\n", totalAset)
}
