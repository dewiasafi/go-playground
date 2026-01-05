package main

import "fmt"

// Interface adalah tipe yg didefinisikan sebagai set dari method signatures
// Nilai yg mau di assign ke dalam interface bisa apa aja asalkan mengimplementasi method-method nya
// Interface ini tidak berisi data atau cara kerja, tapi berisi BEHAVIOR yg harus dimiliki oleh struct
// Interface ini tuh "Kontrak" atau "Perjanjian"
// Interface digunakan agar kode kita menjadi fleksibel dan generic. Dengan interface, kita tidak peduli siapa yang mengerjakan tugasnya, yang penting dia bisa melakukan tugas tersebut.

type person struct {
	name string
	age  int
}

type makhlukHidup interface { // harus didefinisikan apa aja yg masuk ke dalam makhlukHidup ini
	breath()
	move() // misalnya disini ada 2 method ini adalah syarat atau METHOD SIGNATURES dari interface makhlukHidup
	updateName(string)
	//Jadi struct apapun yg mengimplementasi kedua method tersebut disebut sebagai makhlukHidup
}

func (p person) breath() {
	fmt.Println(p.name, "Breath")
}

func (p person) move() {
	fmt.Println(p.name, "Move")
}

func (p *person) updateName(name string) {
	p.name = name
}

func main() {
	orang1 := person{
		name: "orang 1",
		age:  17,
	}
	// Dalam penggunaan nya kita tidak bisa inisiasi objek langsung dari interface
	// Contoh:
	// m1:= makhlukHidup{} ---> ini tidak bisa
	// cara inisiasinya adalah
	var m makhlukHidup
	// m disini pakai pointer karena di dalam interface makhlukHidup ada method update yang pakai params name untuk update p.name
	// jadi pakai & untuk pointer agar name di struct person, bukan buat data baru
	m = &orang1 // ini akan error jika tidak struct person tidak diimplementasi ke dalam method breath dan move
	m.breath()
	m.move()
	// fmt.Println(c.name) ini tidak bisa dilakukan karena interface semacam membungkus data dan menyembunyikan detail implementasi (Encapsulate)
	// Jadi nantinya data yg dipassing tidak bisa di
	// Cuma bisa  c.method aja

	// MUTASI data dari interface
	m.updateName("orang baru")
	fmt.Println(orang1.name)

	//Nanti kalo ada struct baru, dan mau masuk ke dalam interface maka struct tersebut harus digunakan dalam method breath, move, dan updateName
	// Contoh Penggunaan

	// 1. Inisiasi struct metode pembayaran
	metodeOvo := OVO{}
	metodeBank := Bank{}

	fmt.Println("--- Transaksi Pertama ---")
	// 2. Masukkan ke fungsi ProsesTransaksi
	// Ini bisa masuk karena OVO sudah punya fungsi Bayar(int)
	ProsesTransaksi(metodeOvo, 50000)

	fmt.Println("\n--- Transaksi Kedua ---")
	// Ini juga bisa masuk karena Bank punya fungsi Bayar(int)
	ProsesTransaksi(metodeBank, 150000)

	// 3. Contoh menggunakan Interface sebagai tipe data variabel
	var pilihanPembayaran Pembayaran

	pilihanPembayaran = OVO{}
	pilihanPembayaran.Bayar(10000)
}
