package main

import (
	"fmt"
	"net/http"
)

// Web Service API adalah sebuah web yang menerima request dari client dan menghasilkan response,
// biasa berupa JSON/XML atau format lainnya.

type Student struct {
	ID    string
	Name  string
	Grade int
}

var Data = []Student{
	{"E001", "ethan", 11},
	{"W001", "wick", 12},
	{"B001", "bourne", 10},
	{"B002", "bond", 11},
}

//  Ini dipanggil oleh server saat ada request masuk.
/*Fungsinya:
- Baca request dari client
- Proses data
- Kirim response
*/

func main() {
	http.HandleFunc("/users", GetUsers)
	http.HandleFunc("/user", GetUserById)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
