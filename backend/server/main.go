package main

import (
	"fmt"
	"io"
	"net/http"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
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

func main() {
	myApp = app.New()
	myWindow = myApp.NewWindow("Main Window")

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/encrypt", handleEncrypt)
	http.HandleFunc("/decrypt", handleDecrypt)

	http.ListenAndServe(":3333", nil)

	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("Server Running..."),
	))
	myWindow.ShowAndRun()
}

func requirePassword() []byte {
    myApp.QueueMain(func() {
		popup()
	})
    return make([]byte, 32)
}

func popup() {
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