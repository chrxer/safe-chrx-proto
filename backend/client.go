package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
    url := "http://localhost:3333/encrypt"
    data := []byte("veryImportantPassword")

    req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
    if err != nil {
        fmt.Printf(err.Error())
		return
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Printf(err.Error())
		return
    }
    defer resp.Body.Close()

    fmt.Printf("%s\n", resp.Status)
    fmt.Printf("%s\n", resp.Header)
    fmt.Printf("%s", resp.Body)
}