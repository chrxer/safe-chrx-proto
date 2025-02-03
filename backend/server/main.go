package main

import (
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	if(r.Method == "POST"){
		fmt.Printf("post\n")
		reqBody, err := io.ReadAll(r.Body)
    	if err != nil {
        	fmt.Printf(err.Error())
    	}
    	fmt.Printf("%s", reqBody)

	} else {
		fmt.Printf("smth else (GET)\n")
	}
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	http.ListenAndServe(":3333", nil)
}