package main

import (
	"fmt"
	"io"
	"net/http"
)

// var masterPassword []byte
var masterPassword []byte

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func handleEncrypt(w http.ResponseWriter, r *http.Request) {
	if(!checkPost(w, r)) {
		return
	}
	fmt.Printf("post - encrypt\n")
	reqBody, err := io.ReadAll(r.Body)
    if err != nil {
    	fmt.Printf(err.Error())
    }
    fmt.Printf("%s", encrypt(reqBody))
}

func handleDecrypt(w http.ResponseWriter, r *http.Request) {
	if(!checkPost(w, r)) {
		return
	}
	fmt.Printf("post - decrypt\n")
	reqBody, err := io.ReadAll(r.Body)
    if err != nil {
    	fmt.Printf(err.Error())
    }
    fmt.Printf("%s", decrypt(reqBody))
}

func checkPost(w http.ResponseWriter, r *http.Request) bool {
	if(r.Method != "POST"){
		w.WriteHeader(http.StatusBadRequest)
    	w.Write([]byte(""))
		return false
	}
	return true
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/encrypt", handleEncrypt)
	http.HandleFunc("/decrypt", handleDecrypt)

	http.ListenAndServe(":3333", nil)
}