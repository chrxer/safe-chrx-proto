package main

import (
	"bytes"
	"flag"
	"fmt"
	"image/color"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"runtime/debug"
	"strconv"
	"sync"
	"syscall"
	"unicode"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

/* GLOBAL VARIABLES */

var masterKey []byte
var wg sync.WaitGroup
var userPassword string

var myApp fyne.App
var myWindow fyne.Window

/* SERVER ENDPOINT */

func getRoot(w http.ResponseWriter, is_locked bool) {
	var response string;
	response = "See https://github.com/chrxer/safe-chrx-proto/tree/main/backend/server\n"
	if is_locked{
		response += "Status: Locked"
	}else{
		response += "Status: Unlocked"
	}
	io.WriteString(w, response)
}

func handleEncrypt(w http.ResponseWriter, r *http.Request, key []byte) {
	if !isPost(w, r) {
		return
	}
	var reqBody []byte
	var err error
	reqBody, err = io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		reqBody = []byte("")
	}

	// encrypt
	reqBody = decrypt(reqBody, key)

	var encrypted []byte
	encrypted = encrypt(reqBody, []byte(""))
	encrypted = encrypt(encrypted, key)

	w.Write(encrypted)
}

func handleDecrypt(w http.ResponseWriter, r *http.Request, key []byte) {
	if !isPost(w, r) {
		return
	}
	var reqBody []byte
	var err error
	reqBody, err = io.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		reqBody = []byte("")
	}

	// decryption
	reqBody = decrypt(reqBody, key)

	var decrypted []byte
	decrypted = decrypt(reqBody,[]byte(""))
	
	decrypted = encrypt(decrypted, key)

	w.Write(decrypted)
}

func isPost(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(""))
		return false
	}
	return true
}

func serve(port int, key []byte) {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
			getRoot(w,bytes.Equal(masterKey,[]byte("")))
		})
		http.HandleFunc("/encrypt", func(w http.ResponseWriter, r *http.Request){handleEncrypt(w,r,key)})
		http.HandleFunc("/decrypt", func(w http.ResponseWriter, r *http.Request){handleDecrypt(w,r,key)})
		err := http.ListenAndServe("localhost:"+strconv.Itoa(port), nil)
		if err != nil {
			panic(err.Error())
		}
	}

/* MAIN */

func main() {
	// Set up logging file first
	tmpFile, err := os.OpenFile(path.Join(os.TempDir(), "chrxCryptServerLog.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Could not create temporary log file", err)
	}
	log.SetOutput(tmpFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	fmt.Printf("Logging to file at :\"%s\"\n", tmpFile.Name())

	var port *int
	// Ensure cleanup on exit
	defer func() {
		log.Printf("Exiting (port: %d)\n", *port)
		if r := recover(); r != nil {
			s:= fmt.Sprintf("%s: %s", r, debug.Stack())
			log.Println(s)
			println(s)
		}
		
		tmpFile.Close()
	}()

	// Catch OS signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		s := <-sigChan
		log.Printf("Caught signal: %s, contintuing anyways\n", s)
	}()

	// CLI flags
	port = flag.Int("port", 3333, "port to serve on")
	reset := flag.Bool("reset", false, "Reset the password. All currently encrypted data will be lost")
	flag.Parse()
	if *reset {
		writeHash("")
		log.Printf("Password reset by user\n")
		return
	}
	
	key := readAESKeyFromStdin()

	myApp = app.New()
	drv := myApp.Driver()

	log.Printf("Started on %d\n",*port)

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
		panic("Failed to create splash window")
	}

	go serve(*port, key)

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
	errorLabel := canvas.NewText("Please enter master password...", theme.Color(theme.ColorNameForeground))
	submitButton := widget.NewButton("OK", func() {
		pswd := entry.Text
		if len(pswd) < 8 {
			errorLabel.Text = "Password must be at least 8 characters long"
		} else if len(pswd) > 32 {
			errorLabel.Text = "Password cannot be over 32 characters long"
		} else if isValid(pswd) {
			userPassword = pswd;
			myWindow.Hide()
			wg.Done() // See getMasterPassword() in crypt.go
			return
		} else {
			errorLabel.Text = "Invalid password"
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
	errorLabel := canvas.NewText("Please create master password", theme.Color(theme.ColorNameForeground))
	submitButton := widget.NewButton("OK", func() {
		pswd1 := entry1.Text
		pswd2 := entry2.Text
		if pswd1 != pswd2 {
			errorLabel.Text = "Passwords must be identical"
		} else if (len(pswd1) < 8) {
			errorLabel.Text = "Password must be at least 8 characters long"
		} else if len(pswd1) > 32 {
			errorLabel.Text = "Password cannot be over 32 characters long"
		} else {
			for _, char := range pswd1 {
				if !unicode.IsPrint(char) || unicode.IsSpace(char) {
					errorLabel.Text = "Password can only contain printable characters"
					return
				}
			}

			firstChar := pswd1[0]
			allSame := true
			for i := 1; i < len(pswd1); i++ {
				if pswd1[i] != firstChar {
					allSame = false
					break
				}
			}
			if allSame {
				errorLabel.Text = "Password cannot have all characters the same"
				return
			}
			
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