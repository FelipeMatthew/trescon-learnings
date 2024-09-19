package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()

	// Title
	title := "Hi im a title"

	w := a.NewWindow(title)
	// Resize
	w.Resize(fyne.NewSize(400, 400))

	w.ShowAndRun()
}
