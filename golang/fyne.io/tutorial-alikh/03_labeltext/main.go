package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	title := "Another title"
	w := a.NewWindow(title)
	w.Resize(fyne.NewSize(400, 400))
	w.SetContent(widget.NewLabel("Hiii im a label"))
	w.ShowAndRun()
}
