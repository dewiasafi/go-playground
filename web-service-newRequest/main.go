package main

import "fmt"

var BaseUrl = "http://localhost:8080"

type Student struct {
	ID    string
	Name  string
	Grade int
}

// http.NewRequest() → CLIENT SIDE
// Ini dipakai untuk MEMBUAT request ke server.
/* Biasanya dipakai kalau:
   	- Aplikasi Go kamu mau memanggil API lain
	- Bikin HTTP client
	- Integrasi ke service eksternal
*/

// http.NewRequest digunakan untuk mengirim request, bukan menerima
func main() {
	// users, err := FetchUsers()
	// if err != nil {
	// 	fmt.Println("Error!", err.Error())
	// 	return
	// }

	// for _, each := range users {
	// 	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	// }

	user1, err := FetchUser("E001")
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade)

}
