package main

import "fmt"

// struct adalah salah satu konsep u/ mengelompokan variable dengan tipe data yg berbeda-beda.
// go bukan bahasa OOP murni (ga ada class), jadi struct dipakai untuk kelompokan data

type Barang struct {
	Nama  string
	Harga float64
	Stok  int
}

func main() {
	// Cara 1: Menyebutkan nama field (Sangat disarankan)
	item1 := Barang{
		Nama:  "Kopi Arabika",
		Harga: 25000,
		Stok:  10,
	}

	// Cara 2: Tanpa nama field (harus urut)
	item2 := Barang{"Gula Pasir", 20000, 15} // karena nama field nya tidak disebutkan maka nilai setiap field yg dimasukan harus urut

	// Cara 3: Kosongan dulu, isi belakangan
	var item3 Barang
	item3.Nama = "Susu UHT" // isi struct Barang dengan field Nama jadi "Susu UHT"
	item3.Harga = 18000
	item3.Stok = 5

	// Cara akses data struct
	// sebuah struct dapat diambil/diubah datanya pakai (.) mirip kyk akses field di objek

	fmt.Println("Program Berhasil Dijalankan!")
	fmt.Printf("Item 1: %s, Harga: %.2f, Stok:%d\n", item1.Nama, item1.Harga, item1.Stok)
	fmt.Printf("Item 2: %s, Harga: %.2f, Stok:%d\n", item2.Nama, item2.Harga, item2.Stok)
	fmt.Printf("Item 2: %s, Harga: %.2f, Stok:%d\n", item3.Nama, item3.Harga)

	// Struct di dalam struct (embedded struct)
	// masukan struct ke dalam struct lain. digunakan ketika data komplek
	type Alamat struct {
		Kota     string
		Provinsi string
	}

	type Pelanggan struct {
		Nama   string
		Member bool
		Lokasi Alamat // lokasi ini diambil dari struct Alamat
	}

	// Tag JSON di Struct (Penting untuk Projek Kamu)
	// ketika mau simpan data struct ke dalam json, perlu menambahkan nama kunci JSON-nya menggunakan Tags.
	type Transaksi struct {
		ID     int     `json:"transaction_id"`
		Total  float64 `json:"total_payment"`
		Status string  `json:"status"`
	}
	// tanpa tags ini, saat disimpan nanti pakai nama field nya.
	// misal nya field ID maka tersimpan ID di json
}
