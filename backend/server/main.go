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

/* GLOBAL VARIABLES */

var masterKey []byte
var wg sync.WaitGroup
var userPassword string

var myApp fyne.App
var myWindow fyne.Window

/* SERVER ENDPOINT */

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "See https://github.com/chrxer/safe-chrx-proto/tree/main/backend/server\n")
}

func handleEncrypt(w http.ResponseWriter, r *http.Request) {
	if !isPost(w, r) {
		return
	}
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	w.Write(encrypt(reqBody))
}

func handleDecrypt(w http.ResponseWriter, r *http.Request) {
	if !isPost(w, r) {
		return
	}
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	w.Write(decrypt(reqBody))
}

func isPost(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(""))
		return false
	}
	return true
}

func serve() {
		http.HandleFunc("/", getRoot)
		http.HandleFunc("/encrypt", handleEncrypt)
		http.HandleFunc("/decrypt", handleDecrypt)
		err := http.ListenAndServe("localhost:3333", nil)
		if err != nil {
			fmt.Println("Server error: ", err)
			// Handle the error appropriately
		}
	}

/* MAIN */

func main() {
	myApp = app.New()
	drv := myApp.Driver()

	if drv, ok := drv.(desktop.Driver); ok {
		myWindow = drv.CreateSplashWindow()

		// Hide instead of close -> Closing stops the entire program
		myWindow.SetCloseIntercept(func() {
			wg.Done() // See getMasterPassword() in crypt.go
			myWindow.Hide()
		})

		var content *fyne.Container

		// Is there already a master password set?
		if(len(fetchHash()) == 0) {
			content = createPswdSetterWindow()
		} else {
			content = createPswdQueryWindow()
		}
		myWindow.SetContent(container.NewPadded(content))
		myWindow.Resize(content.Size())
	} else {
		fmt.Println("Failed to create splash window")
	}

	go serve()

	myWindow.Hide()
	myApp.Run() // Due to being hidden it runs in the background
}

func isValid(pswd string) bool {
	hash := fetchHash()
	return argonCheckPswd(pswd, hash)
}

func createPswdQueryWindow() *fyne.Container {
	// entry = input box
	entry := widget.NewPasswordEntry()
	entry.SetPlaceHolder("Enter password...")

	// errorLabel will be used to display errors (in red text)
	errorLabel := canvas.NewText("Please enter master password...", color.Black)
	submitButton := widget.NewButton("OK", func() {
		pswd := entry.Text
		if len(pswd) < 8 {
			errorLabel.Text = "Password is too short"
		} else if len(pswd) > 32 {
			errorLabel.Text = "Password is too long"
		} else { 
			if isValid(pswd) {
				userPassword = pswd;
				myWindow.Hide()
				wg.Done() // See getMasterPassword() in crypt.go
				return
			} else {
				errorLabel.Text = "Invalid password"
			}
		}
		// Set error color (red)
		errorLabel.Color = color.RGBA{R: 255, G: 80, B: 80, A: 255}
		errorLabel.Refresh()
	})
	
	// Creates padded container (horizontal -> Entry + button) within a container (vertical -> Error label + previous container)
	content := container.NewPadded(container.NewBorder(errorLabel, nil, nil, nil, container.NewBorder(nil, nil, nil, submitButton, entry)))
	
	return content
}

func createPswdSetterWindow() *fyne.Container {
	// Entry = input box
	entry1 := widget.NewPasswordEntry()
	entry1.SetPlaceHolder("Enter password...")

	entry2 := widget.NewPasswordEntry()
	entry2.SetPlaceHolder("Confirm password...")

	// errorLabel will be used to display errors (in red text)
	errorLabel := canvas.NewText("Please create master password", color.Black)
	submitButton := widget.NewButton("OK", func() {
		pswd1 := entry1.Text
		pswd2 := entry2.Text
		if pswd1 != pswd2 {
			errorLabel.Text = "Passwords need to be identical"
		} else if (len(pswd1) < 8) {
			errorLabel.Text = "Password is too short"
		} else if len(pswd1) > 32 {
			errorLabel.Text = "Password is too long"
		} else { 
			writeHash(argonHash(pswd1))
			userPassword = pswd1;
			myWindow.Hide()
			wg.Done() // See getMasterPassword() in crypt.go
			return
		}
		// Set error color (red)
		errorLabel.Color = color.RGBA{R: 255, G: 80, B: 80, A: 255}
		errorLabel.Refresh()
	})

	// Creates padded container (vertical -> Entry + entry + button) within a container (vertical -> Error label + previous container)
	content := container.NewPadded(container.NewBorder(errorLabel, nil, nil, nil, container.NewVBox(entry1, entry2, submitButton)))
	
	return content
}