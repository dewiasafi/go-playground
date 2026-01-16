package main

import (
	"fmt"
	"math/rand"
	"time"
)

// untuk menghasilkan random di golang ada 2 cara:
// 1. Pakai math/rand: u/ game, simulasi, testing, data dummy
// 2. Pakai crypto/rand: u/token, password, OTP, security

var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	// math/rand
	randomizer := rand.New(rand.NewSource(122))          // 122 adalah seed (bibit angka random), kalo seednya sama ketika di run ulang angka random nya pasti sama kayak sebelumnya
	fmt.Println("random float32:", randomizer.Float32()) // Float32 adalah tipe data angka random yg mau ditampilkan
	fmt.Println("random int:", randomizer.Int())
	fmt.Println("random uint:", randomizer.Uint32())
	fmt.Println("---------------------------")

	// Random unik number
	randomizerUnix := rand.New(rand.NewSource(time.Now().UTC().UnixNano())) // seed pakai time agar angka randomnya setiap render beda-beda
	fmt.Println("random unik float32:", randomizerUnix.Float32())
	fmt.Println("random unik unib int:", randomizerUnix.Int())
	fmt.Println("random unik uint:", randomizerUnix.Uint32())
	fmt.Println("---------------------------")

	// Random String
	fmt.Println("random string 5 karakter:", randomString(5)) // 5 disini bukan seed tapi panjang random yang mau dihasilkan

	// crypto/rand
	CryptoRand()

}
