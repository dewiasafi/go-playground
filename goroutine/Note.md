# Penjelasan Eksekusi Goroutine di Go

Dokumen ini menjelaskan **alur eksekusi program Go** ketika menggunakan `goroutine`, berdasarkan contoh pemanggilan fungsi `print` sebanyak dua kali di `main.go`.

---

## Apa Itu Goroutine?

```go
Goroutine itu sering disebut sebagai "mini thread"
Dieksekusi secara asynchronous (tidak menunggu proses lain selesai)

Goroutine merupakan bagian dari concurrent programming di Go
Untuk mengeksekusi sebuah proses sebagai goroutine, proses tersebut
WAJIB dibungkus dalam sebuah function

Cara pakainya mirip seperti memanggil function biasa,
tetapi dengan menambahkan keyword `go` di depannya
```

### Penjelasan Sederhana

* **Goroutine** adalah unit eksekusi ringan (lebih ringan dari thread OS)
* Go runtime yang mengatur penjadwalan goroutine, bukan OS secara langsung
* Goroutine berjalan **asynchronous**, artinya tidak saling menunggu

Contoh:

```go
go print(3, "Hello") // dijalankan sebagai goroutine
print(3, "Apa Kabar?") // dijalankan di main goroutine
```

Pada contoh di atas:

* `print("Hello")` dijalankan di goroutine terpisah
* `print("Apa Kabar?")` tetap dijalankan di goroutine `main`

---

## Studi Kasus

Di dalam `main.go` terdapat dua pemanggilan fungsi:

* **A** → `go print(3, "Hello")`
* **B** → `print(3, "Apa Kabar?")`

Pertanyaannya:

* Apakah **A harus dijalankan lebih dulu** karena ditulis sebelum **B**?
* Kenapa hasil output bisa tidak berurutan?

---

## Urutan Penulisan ≠ Urutan Eksekusi

Secara **urutan kode**, memang:

1. A ditulis lebih dulu
2. B ditulis setelahnya

Namun, **urutan penulisan kode tidak menjamin urutan eksekusi**, terutama ketika melibatkan `goroutine`.

---

## Apa yang Terjadi Saat `go` Digunakan?

Ketika kita menambahkan keyword `go` di depan pemanggilan fungsi:

* Fungsi tersebut **tidak dijalankan di thread utama (main thread)**
* Fungsi akan dijalankan secara **asynchronous** di goroutine lain
* Eksekusi program **tidak menunggu** fungsi tersebut selesai

Artinya:

* **A** dijalankan di goroutine terpisah
* **B** tetap dijalankan di goroutine utama (`main`)

---

## A dan B Berjalan Bersamaan (Concurrent)

Karena:

* A berjalan di goroutine lain
* B berjalan di goroutine `main`

Maka:

> **A dan B berjalan secara bersamaan (concurrent), bukan bergantian**

Akibatnya:

* Bisa saja **B selesai lebih dulu dari A**
* Urutan output **tidak bisa diprediksi**

Contoh kemungkinan output:

* `Apa Kabar?` muncul lebih dulu
* atau `Hello` muncul lebih dulu

Semua tergantung **scheduler Go**.

---

## Masalah Utama: Program Terlalu Cepat Selesai

Perlu dipahami satu hal penting:

> Jika **fungsi `main()` selesai**, maka **seluruh goroutine akan langsung dihentikan**.

Jadi jika:

* Proses di `main` (B) selesai lebih cepat
* `main()` berakhir

Maka:

* Goroutine A **belum tentu sempat menyelesaikan tugasnya**
* Program langsung berhenti

---

## Solusi Sederhana: Menahan Program dengan `fmt.Scanln`

Salah satu cara paling sederhana untuk mencegah `main()` selesai terlalu cepat adalah:

* Menambahkan `fmt.Scanln(&input)` di akhir `main()`

Efeknya:

* Program **menunggu input user (Enter)**
* `main()` tidak langsung selesai
* Goroutine lain punya waktu untuk menyelesaikan pekerjaannya

---

## Kenapa `fmt.Scanln` Bisa Bekerja?

Karena:

* `fmt.Scanln` bersifat **blocking**
* Goroutine `main` akan **berhenti sementara**
* Goroutine lain tetap berjalan

Dengan begitu:

> Semua goroutine diberi kesempatan untuk selesai sebelum program berakhir.
**NOTE:** `fmt.Scanln` **bukan solusi profesional** untuk sinkronisasi goroutine.
---

## Kesimpulan

* Urutan penulisan kode **tidak menjamin urutan eksekusi**
* `go` membuat fungsi berjalan di goroutine terpisah
* Goroutine berjalan **bersamaan (concurrent)**
* Jika `main()` selesai, semua goroutine akan dihentikan
* `fmt.Scanln` dapat digunakan untuk menahan `main()` sementara
