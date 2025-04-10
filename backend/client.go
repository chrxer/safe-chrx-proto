package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const url = "http://localhost:3333/"

func _req(s []byte, endpoint string) []byte {
	endpoint_url := url + endpoint
	
	req, err := http.NewRequest(http.MethodPost, endpoint_url, bytes.NewBuffer(s))
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

	//fmt.Printf("\n%s\n", resp.Body)

	if resp.StatusCode != 200 {
		fmt.Printf("Excpected status code 200, but got: %s", strconv.Itoa(resp.StatusCode))
		return []byte("")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return []byte("")
	}

	return body
}

func encrypt(s []byte) []byte {
	return _req(s, "encrypt")
}

func decrypt(s []byte) []byte {
	return _req(s, "decrypt")
}

func main() {
	data := []byte("password")
	encrypted := encrypt(data)
	fmt.Printf("\nPlaintext: %s\nEncrypted (hex): %x (%s)\n", data, encrypted,encrypted)
	var decrypted = []byte("")
	decrypt(encrypted)
	decrypt(encrypted)
	decrypted = decrypt(encrypted)

	fmt.Printf("Decrypted: %s\n\n", decrypted)
}
