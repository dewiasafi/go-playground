# Channel di Go (Golang)

## 🤔 Apa itu Channel?
Channel itu dipakai buat **ngehubungin satu goroutine ke goroutine lain** lewat mekanisme **send-receive data**.

Jadi intinya ✨: 
- Goroutine A **ngirim data** 📤
- Goroutine B **nerima data** 📥

Channel ini jadi **jembatan / pipa** buat data pindah antar goroutine dengan aman.

Go punya prinsip terkenal:
> **Don't communicate by sharing memory, share memory by communicating**

---
## Kenapa Perlu Channel?
Karena goroutine itu jalan **concurrent**.

Kalau banyak goroutine akses data bareng tanpa aturan:
- Bisa bentrok
- Bisa race condition
- Bisa bug aneh

Channel bantu:
- Sinkronisasi data
- Hindari race condition
- Tanpa ribet pakai `sync.Mutex`

---

## Channel itu variabel
Channel itu variabel dengan tipe `chan`. Tapi channel ini jadi **media komunikasi antar goroutine**, jadi bukan sekedar variabel biasa.

**chan** adalah tipe data khusus di Go.

Ciri-cirinya:
- Strongly typed
- Satu channel cuma buat satu tipe data
- Aman buat komunikasi antar goroutine

example: `chan int`, `chan string`, `chan bool`

`make` **(Wajib)**

ketika buat channel wajib menggunakan built-in function `make`. Kalo buat tanpa deklarasi, 

misalnya:

 `var ch chan int` nilainya **nil**.

 akibatnya:

 - kirim data ❌
 - nerima data ❌
 - program deadlock

Makanya ketika buat channel wajib pakai `make`

`ch := make(chan int)`

## Operator `<-`
Operator `<-` dipakai buat kirim dan nerima data.

**kirim data**: `ch <-100`, kirim data 100 ke channel ch

**nerima data** : `hasil := <-ch`, terima data dari channel ch dan simpan di variabel *hasil*

---

## Sifat channel
### Sebagai perantara 
Channel itu kayak **pipa**:
- Ada yang **ngirim (sender)**
- Ada yang **nerima (receiver)**

### Blocking / Synchronous
Pengiriman dan penerimaan data di channel bersifat **blocking** (sinkron):
- Sender bakal **nunggu** sampai ada yang nerima data
- Receiver bakal **nunggu** sampai ada yang kirim data

Kode setelah operasi channel **baru jalan** setelah proses kirim–terima selesai.
```go
func main() {
	ch := make(chan int)
	ch <-10 // DEADLOCK
    hasil := <-ch
} 
```
 Kenapa deadlock? karena ga ada goroutine lain yang nerima datanya, proses kirim–terima data berlangsung di 1 goroutine.

 **Solusinya pakai goroutine**
 ```go 
 	ch := make(chan int)

	go func() {
		ch <- 100 //Mengirim data ke channel
	}()
	hasil := <-ch //Menerima data dari channel
	fmt.Println(hasil)
}
 ```
Sekarang: 
- Goroutine ngirim data 100 ke channel ch
- Main goroutine terima datanya dan simpan di variabel hasil

**Ilustrasi Alur Channel**
```css
[Goroutine A] ---- data ----> [ Channel ] ---- data ----> [Goroutine B]
```

---
## Apa itu send-only dan receive-only channel?

Di Go, **channel bisa dibatasi arahnya**:

- Send-only → cuma bisa ngirim

- Receive-only → cuma bisa nerima

Ini **bukan tipe channel baru**, tapi **pembatasan akses**.

- chan<- → panah ke channel (kirim)
- <-chan → panah dari channel (terima)

### Kenapa ini ada? 🤔

**Tujuan utamanya:**
- 👉 Compile-time safety
- 👉 Desain API yang jelas
- 👉 Mencegah bug concurrency

## Buffer & Unbuffer Channel

### Unbuffer Channel 

**Unbuffer** channel itu channel biasa yang proses nya harus tunggu ada sender dan receiver

```go
ch := make(chan int)
ch<-10
<-ch //DEADLOCK
```
**Proses**
- send ➜ nunggu receive
- receive ➜ nunggu send
- **Harus ketemu di waktu yang sama**

### Buffer Channel

**Buffered** channel adalah channel yang **punya kapasitas** untuk menyimpan beberapa data **tanpa harus langsung di-receive**.

```go
ch := make(chan int, 3) // buffer size = 3
```
**Artinya:**
- Channel bisa menampung 3 data
- send tidak langsung nge-block selama buffer belum penuh
- receive tidak nge-block selama buffer belum kosong

#### Prosesnya:
- send ➜ langsung masuk buffer (kalau belum penuh)
- receive ➜ ambil dari buffer (kalau belum kosong)
- Tidak harus barengan

**Note:** Proses akan terus **berjalan (tidak deadlock)** selama **jumlahnya masih sesuai kapasitas** channel, kalo sudah **melebihi** kapasitas maka **deadlock** 

```go
ch := make(chan int, 2)

ch <- 10
ch <- 20

fmt.Println(<-ch) // 10
fmt.Println(<-ch) // 20

```
#### 📌 Cara Receive Data Buffer Channel
- Ambil data FIFO (First In First Out)
- Mengosongkan buffer sedikit demi sedikit

---

## Select Channel

`select` dipakai untuk **menunggu beberapa operasi channel sekaligus**.

➡️ Mirip switch, tapi **khusus channel**

➡️ Dipakai untuk:
- Menghindari deadlock
- Timeout
- Non-blocking receive/send
- Fan-in / worker

### 🧠 select itu nunggu “channel yang hidup & siap”

- Channel kosong → ditunggu
- Channel closed → langsung ready
- Channel nil → **diabaikan**

📌 Makanya:
- Channel mati → close
- Channel selesai → set nil