# Asyncronous Function di Golang

## Go punya Async Function 🤔?

di Go **ngga ada keyword `async`/  `await`** seperti di Javascript. 

Tapi...

👉 **Go tetap mendukung asynchronous programming**, bahkan jadi salah satu kekuatan utamanya.

Di Go:
> **Asynchronous = goroutine (+ channel / waitgroup)**

---

## Goroutine = Async di Go
Goroutine adalah function yang dijalankan **secara asynchronous**.

*Example*
```go
func doTask() {
    fmt.Println("DOING TASK...")
	time.Sleep(2 * time.Second)
	fmt.Println("DONE")
}
```
```go
func main(){
    doTask()
    fmt.Println("Test")
}
```
Tunggu `doTask()` selesai baru eksekusi `fmt.Println("Test")`, artinya: proses berjalan secara **syncronous**, tapi kalo seperti ini.

```go
func main(){
    go doTask() 
    fmt.Println("Test")}
```
artinya: 
- `doTask()` jalan di goroutine lain
- `fmt.Println("Test")` lanjut jalan sebagai program utamanya
- Non-blocking

** *tapi...* proses ini akan jadi masalah kalo doTask() mengembalikan nilai tertentu**. 

**KENAPA?**

Karena `doTask()` berjalan di goroutine lain, `main` tidak bisa menangkap nilai dari `return` oleh  `doTask()`. 

Jadi, untuk **async yang mengembalikan nilai itu bukan pakai `return` langsung**, tapi lewat **channel** 

**NOTE: Async di Go yang punya nilai balik = goroutine + channel**

*Example:*
```go
func asyncHitung(x int, ch chan int) {
	ch <- x * 2
}

func main() {
	ch := make(chan int)
	go asyncHitung(10, ch) // async function

	hasil := <-ch // nunggu hasil (mirip await)
	fmt.Println(hasil)
}
```
**Biar Lebih Rapi (Function Return Channel)**

Biasanya function-nya ngembaliin channel, bukan nerima channel. jadi proses Asyncronous-nya berjalan di function `asyncHitung`

```go
func asyncHitung(x int) chan int { 
	ch := make(chan int)

	go func() {
		ch <- x * 2
	}()

	return ch
}

func main() {
	result := asyncHitung(10)
	fmt.Println(<-result)
}
```
---
## Kesalahan pada pola Async Function

Pada program `main.go` terdapat function yang namanya **doSomethingAsyncKatanya** yang secara hasil mirip dengan *Async Function* padahal bukan.

**WHY ?** 
```go
func doSomethingAsyncKatanya(ch chan int) chan int {
	fmt.Println("DOING SOMETHING By doSomethingAsyncKatanya ...")
	time.Sleep(2 * time.Second)
	fmt.Println("ASYNC By doSomethingAsyncKatanya DONE")
	ch <- 1
	return ch
}
```
**1️⃣ Kesalahan di doSomethingAsyncKatanya**

❌ Tidak async
- tidak ada keyword `go` di dalam function *doSomethingAsyncKatanya* 
- function ini tetap **Blocking**
❌ Nerima `chan` tapi juga return `chan`
- Redundan
- Caller sudah pegang `ch`, 
- return `ch` itu nggak ada gunanya
❌ Disebut async, tapi cara panggilnya `go doSomethingAsyncKatanya(c)`

**Kesimpulannya:**

 `doSomethingAsyncKatanya` **bukan** async function, tapi **sync function yang dipaksa jalan di goroutine.**

---
## RULE OF THUMB (INI PENTING 🔑) 

Kalo mau dikatakan sebagai async function di Go maka **Harus**

**✅ Boleh**
- Buat sendiri goroutine `(go)` di dalam functionnya *(bukan ketika ditambahkan go ketika call function)*
- Channel dikelola oleh function dan di return untuk caller 
- Channel boleh dikelola oleh main, asalkan function tidak memiliki return nilai
- Caller yang ambil hasil `<-ch`, kalo dipanggil tanpa `<-` nanti hasilnya cuma alamat memory aja

#### Prinsip KERAS async function di Go🔒

**❗ Goroutine TIDAK BISA mengembalikan nilai langsung**

Kalau function:
- pakai go
- dan mau “balikin hasil” **👉 HARUS lewat channel / shared state / callback**


**❌ Jangan**
- Nerima channel + return channel
- Nerima channel + return value (`int`, `string`, `dll`)
- Ngaku async tapi ngga ada `go` di dalam function
