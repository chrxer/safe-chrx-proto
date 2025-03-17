package main

import (
	"fmt"
	"io"
	"net/http"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/driver/desktop"
)

var masterKey []byte

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "See https://github.com/chrxer/safe-chrx-proto/tree/main/backend/server\n")
}

func handleEncrypt(w http.ResponseWriter, r *http.Request) {
	if !testPost(w, r) {
		return
	}
	fmt.Printf("Encrypting..\n")
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	w.Write(encrypt(reqBody))
}

func handleDecrypt(w http.ResponseWriter, r *http.Request) {
	if !testPost(w, r) {
		return
	}
	fmt.Printf("Decrypting..\n")
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	w.Write(decrypt(reqBody))
}

func testPost(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(""))
		return false
	}
	return true
}

var myApp fyne.App
var myWindow fyne.Window

func serve() {
		http.HandleFunc("/", getRoot)
		http.HandleFunc("/encrypt", handleEncrypt)
		http.HandleFunc("/decrypt", handleDecrypt)
		err := http.ListenAndServe("localhost:3333", nil)
		if err != nil {
			fmt.Printf("Server error")
			// Handle the error appropriately
		}
	}

func main() {
	myApp = app.New()
	// var myWindow fyne.Window
	drv := myApp.Driver()
	if drv, ok := drv.(desktop.Driver); ok {
		myWindow = drv.CreateSplashWindow()
	} else {
		myWindow = myApp.NewWindow("Couldn't create splash window")
	}

	go serve()
	// go myWindow.Hide() // put this one into a function with a slight delay (after myWindow hsa started up)
	myWindow.ShowAndRun()
}


func requirePassword() []byte {
	POPUP()
	return []byte("a")
}


func POPUP() {
	myWindow.Hide()
	win := myApp.NewWindow("Popup Window")
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter something...")

	dialogBox := dialog.NewCustomConfirm("Input Needed", "OK", "Cancel", entry,
		func(confirm bool) {
			if confirm {
				fmt.Println("User entered:", entry.Text)
			} else {
				fmt.Println("User cancelled input")
			}
			win.Close()
		}, win)

	dialogBox.Show()
	win.Show()
}