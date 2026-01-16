package main

import (
	"crypto/rand"
	"fmt"
)

func CryptoRand() {
	b := make([]byte, 16)
	rand.Read(b)
	fmt.Printf("Crypto random: %x\n", b)
}
