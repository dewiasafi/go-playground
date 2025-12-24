package main

import "fmt"

func DiskonPersen(member bool, totalAwal float64) {
	diskon := 0.0

	if member {
		diskon = 0.1
	}

	if totalAwal > 500000 {
		diskon += 0.05
	}

	potongan := totalAwal * diskon
	totalAkhir := totalAwal - potongan

	fmt.Println("-------------------------------")
	fmt.Printf("Total Awal      : Rp %10.2f\n", totalAwal)
	fmt.Printf("Diskon (%.0f%%) : Rp %10.2f\n", diskon*100, potongan)
	fmt.Printf("TOTAL BAYAR     : Rp %10.2f\n", totalAkhir)
	fmt.Println("===============================")
}
