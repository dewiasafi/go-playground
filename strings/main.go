package main

import (
	"fmt"
	"strings"
)

func main() {
	//strings.Contains() cek apakah string tersebut
	var isExists = strings.Contains("john wick", "wick") // (dataSumber, dataDicek)
	fmt.Println(isExists)

	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1) // true

	var isPrefix2 = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2) // false

	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1) // false

	var isSuffix2 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2) // true

	var howMany = strings.Count("ethan hunt", "t")
	fmt.Println(howMany) // 2

	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1) // 2

	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"

	var str = strings.Repeat("na", 4)
	fmt.Println(str) // "nananana"

	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1) // output: ["the", "dark", "knight"]

	var string2 = strings.Split("batman", "")
	fmt.Println(string2) // output: ["b", "a", "t", "m", "a", "n"]

	var data = []string{"banana", "papaya", "tomato"}
	var str1 = strings.Join(data, "-")
	fmt.Println(str1) // "banana-papaya-tomato"

	var str2 = strings.ToLower("aLAy")
	fmt.Println(str2) // "alay"

	var str3 = strings.ToUpper("eat!")
	fmt.Println(str3) // "EAT!"
}
