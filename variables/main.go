package main

import "fmt"

func main() {
	const title = "Data Siswa" // title is constanta, so we can't change the value
	var nama string            // Nilai default nya ""
	var usia uint8             // bisa int, int16, int32, int64, uint. Nilai default nya 0
	var tinggi float32         // bisa float32, float 64
	var aktif bool             // Nilai defaultnya false

	nama = "Budi"
	usia = 18
	tinggi = 173.5
	aktif = true
	fmt.Printf("%s\nNama: %s \nUsia: %d \nTinggi: %.2f cm \nAktif: %t\n", title, nama, usia, tinggi, aktif)

	// cara cepat deklarasi variable
	kelas := 12 //cuma bisa dipakai kalo deklarasi di dalam function
	fmt.Printf("Kelas: %v", kelas)
}
