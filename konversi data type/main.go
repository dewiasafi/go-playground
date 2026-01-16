package main

import (
	"fmt"
	"strconv"
)

func main() {
	// strconv.Atoi() digunakan untuk konversi data dari tipe string ke int
	var str1 = "124"
	var num1, err1 = strconv.Atoi(str1)
	if err1 == nil {
		fmt.Println(num1) // 124
	}

	//strconv.Itoa() digunakan untuk konversi data dari tipe int ke string.
	var num2 = 124
	var str2 = strconv.Itoa(num2)
	fmt.Println(str2)

	//strconv.ParseInt() digunakan untuk konversi string berbentuk numerik dengan basis tertentu ke tipe numerik non-desimal dengan lebar data bisa ditentukan.
	var str3 = "124"
	var num3, err = strconv.ParseInt(str3, 10, 64) //(value, basis data, type data )
	if err == nil {
		fmt.Println(num3) // 124
	}

	var text1 = "halo"
	var b = []byte(text1)

	fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])
	// 104 97 108 111
}
