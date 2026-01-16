package main

// Di Go, data json dituliskan sebagai string.
// Dengan menggunakan json.Unmarshal, json string bisa dikonversi menjadi 2 bentuk objek:
// 1. Dalam bentuk map[string]interface{}
// 2. Objek struct

// Struct User ini digunakan untuk membuat variabel baru penampung hasil decode json string
// Pada operasi decoding data json string ke variabel objek struct, semua level akses property struct penampung harus publik.
// Jika tidak datanya ga masuk
var JsonString = `{"Name": "john wick", "Age": 27}`
var JsonData = []byte(JsonString)

type User struct {
	FullName string `json:"Name"` // "Name" harus sama dengan properti untuk memasukkan nilai FullName di jsonString
	Age      int
}

func main() {
	JsonToStruct()
	JsonToInterface()
	JsonToArray()
	ObjectToJson()
}
