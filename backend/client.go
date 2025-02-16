package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const url = "http://localhost:3333/"

func req(s []byte, endpoint string) []byte {
	req, err := http.NewRequest(http.MethodPost, url+endpoint, bytes.NewBuffer(s))
	if err != nil {
		fmt.Printf("%s", err.Error())
		return []byte("")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return []byte("")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Excpected status code 200, but got: %s", strconv.Itoa(resp.StatusCode))
		return []byte("")
	}

	body, err := io.ReadAll(req.Body)

	if err != nil {
		fmt.Printf("%s", err.Error())
		return []byte("")
	}

	return body
}

func encrypt(s []byte) []byte {
	return req(s, "encrypt")
}

func decrypt(s []byte) []byte {
	return req(s, "decrypt")
}

func main() {
	data := []byte("veryImportantPassword")
	encrypted := encrypt(data)
	decrypted := decrypt(encrypted)

	fmt.Printf("\nData: %s\nEncrypted: %s\nDecrypted: %s", data, encrypted, decrypted)
}
