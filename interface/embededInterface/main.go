package main

import "fmt"

// Interface kecil 1
type Reader interface {
	Read() string
}

// Interface kecil 2
type Writer interface {
	Write(data string)
}

// Embedded Interface: Menggabungkan Reader dan Writer
// Sekarang, siapapun yang ingin jadi ReadWriter WAJIB punya fungsi Read DAN Write.
type ReadWriter interface {
	Reader
	Writer
}

// --- Implementasi ---

type Dokumen struct {
	Isi string
}

func (d *Dokumen) Read() string {
	return d.Isi
}

func (d *Dokumen) Write(text string) {
	d.Isi = text
	fmt.Println("Dokumen berhasil diperbarui.")
}

func main() {
	doc := &Dokumen{Isi: "Halo Dunia"}

	// Variabel ini bertipe ReadWriter (gabungan)
	var rw ReadWriter = doc

	fmt.Println("Isi:", rw.Read())
	rw.Write("Halo Go")
}
