package main

import (
	"fmt"
	"time"
)

// package time milik golang
// Meskipun nama package-nya adalah time, yang dicakup adalah date dan time, jadi bukan hanya waktu saja.
func main() {
	var time1 = time.Now() // waktu sekarang
	fmt.Printf("time now %v\n", time1)

	// time.Date(tahun, bulan, tanggal, jam, menit, detik, nanodetik, timezone)
	var time2 = time.Date(2025, 12, 24, 10, 20, 0, 0, time.UTC) // waktu custom
	fmt.Printf("time date %v\n", time2)

}
