package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	// Title
	w := a.NewWindow("Init")

	// Resize
	w.Resize(fyne.NewSize(400, 400))

	// Button - 1 - text, 2 - function
	button := widget.NewButton("My button", func() {
		fmt.Println("Button clicked")
	})

	// Label Text
	w.SetContent(button)

	w.ShowAndRun()
}
