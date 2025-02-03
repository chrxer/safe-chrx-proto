package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
    url := "http://localhost:3333/hello"
    data := []byte(`{"name": "John", "age": 30}`)

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

    // do something with the response
}