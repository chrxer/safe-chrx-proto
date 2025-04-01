package main

import (
	"fmt"
	"image/color"
	"io"
	"net/http"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

var masterKey []byte
var wg sync.WaitGroup
var userPassword string

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

		myWindow.SetCloseIntercept(func() {
			wg.Done()
			myWindow.Hide()
		})

		entry := widget.NewPasswordEntry()
		entry.SetPlaceHolder("Enter password...")

		errorLabel := canvas.NewText("Please enter master password...", color.Black)
		submitButton := widget.NewButton("OK", func() {
			pswd := entry.Text
			if len(pswd) < 8 {
				errorLabel = canvas.NewText("Password is too short", color.RGBA{R: 255, G: 128, B: 128, A: 255})
			} else if len(pswd) > 32 {
				errorLabel = canvas.NewText("Password is too long", color.RGBA{R: 255, G: 128, B: 128, A: 255})
			} else { 
				if isValid(pswd) {
					userPassword = pswd;
					myWindow.Hide()
					wg.Done()
				} else {
					errorLabel = canvas.NewText("Invalid password", color.RGBA{R: 255, G: 128, B: 128, A: 255})
				}
			}
			// Refresh canvas.text somehow?? -> errorLabel.Refresh() nor content.Refresh() work.
		})

		entry.Resize(fyne.NewSize(100, 0))
		
		content := container.NewPadded(container.NewBorder(errorLabel, nil, nil, nil, container.NewBorder(nil, nil, nil, submitButton, entry)))

		myWindow.SetContent(container.NewPadded(content))
		myWindow.Resize(content.Size())

	} else {
		myWindow = myApp.NewWindow("Couldn't create splash window")
	}

	go serve()
	myWindow.Hide()
	myWindow.SetMaster()
	myApp.Run()
	fmt.Println("App closing..")
}

func isValid(pswd string) bool {
    // -> Golang Argon2
	if len(pswd) == 0 {
		return false
	}
	return true
}