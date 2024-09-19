package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Checkbox")
	w.Resize(fyne.NewSize(400, 400))

	chk := widget.NewCheck("Checkbox label", func(b bool) {
		if b {
			fmt.Println("True")
		} else {
			fmt.Println("False")
		}

	})

	w.SetContent(chk)
	w.ShowAndRun()
}
