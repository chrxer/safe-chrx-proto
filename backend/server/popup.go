package main

import (
	"fmt"

	"fyne.io/fyne/dialog"
	"fyne.io/fyne/v2/widget"
)

var inp string;

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

// Install/configure framework https://github.com/fyne-io/fyne
// Make GUI for user to give masterpassword