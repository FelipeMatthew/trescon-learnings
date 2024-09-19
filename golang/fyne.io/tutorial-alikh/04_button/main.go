package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Window name")
	w.Resize(fyne.NewSize(400, 400))
	// 1 - name of button / 2 - function
	button := widget.NewButton("Click", func() {
		fmt.Println("button clicked")
	})
	w.SetContent(button)
	// w.SetContent(widget.NewLabel("New label text"))
	w.ShowAndRun()
}
