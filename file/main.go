package main

import (
	"fmt"
	"io"
	"os"
)

var path = "./test.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	return (err != nil)
}

func createFile() {
	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("==> file berhasil dibuat", path)
}

func writeFile() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString("Halo Guys \n")
	if isError(err) {
		return
	}
	_, err = file.WriteString("Ayo Belajar \n")
	if isError(err) {
		return
	}

	err = file.Sync()
	if isError(err) {
		return
	}
	fmt.Println("==> file berhasil di isi")
}

func readFile() {
	var file, err = os.OpenFile(path, os.O_RDONLY, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}
	fmt.Println("==> file berhasil dibaca")
	fmt.Println(string(text))
}

func deleteFile() {
	var err = os.Remove(path)
	if isError(err) {
		return
	}
	fmt.Println("==> file berhasil di delete")

}

func main() {
	createFile()
	writeFile()
	readFile()
	deleteFile()
}
