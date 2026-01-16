package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func FetchUsers() ([]Student, error) {
	var err error
	var client = &http.Client{}
	var data []Student

	request, err := http.NewRequest("GET", BaseUrl+"/users", nil)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request) // jalankan request
	if err != nil {
		return nil, err
	}

	defer response.Body.Close() // close ketika sudah tidak digunakan
	// Data response body tersedia via property Body dalam tipe []byte

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	// Gunakan JSON Decoder untuk mengkonversinya menjadi bentuk JSON.
	// Contohnya bisa dilihat di kode di atas, json.NewDecoder(response.Body).Decode(&data)
	return data, nil
}

func FetchUserPost(ID string) (Student, error) {
	var err error
	var client = &http.Client{}
	var data Student

	// url.Values{} akan menghasilkan objek yang nantinya digunakan sebagai form data request.
	var params = url.Values{}

	params.Set("id", ID)
	var payload = bytes.NewBufferString(params.Encode())
	// Proses encoding ada data params u/ diubah jadi bytes.Buffer
	// Data buffer disisipkan ke parameter ketiga pemanggila fungsi http.NewRequest()

	request, err := http.NewRequest("POST", BaseUrl+"/user", payload)
	if err != nil {
		return data, err
	}

	// Karena data yang akan dikirim adalah encoded, maka pada header perlu dituliskan juga tipe encoding-nya.
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request) // jalankan request
	if err != nil {
		return data, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil

}

func FetchUser(ID string) (Student, error) {
	var data Student

	url := BaseUrl + "/user?id=" + ID

	response, err := http.Get(url)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(response.Body)
		return data, fmt.Errorf("server error: %s", string(body))
	}

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		return data, err
	}
	return data, nil
}
