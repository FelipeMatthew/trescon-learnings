package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Horizontal and verical box")
	w.Resize(fyne.NewSize(500, 500))

	btnOne := widget.NewButton("Click", func() {

	})
	labelOne := widget.NewLabel("My text here")

	// vericalBox := container.NewVBox(btnOne, labelOne)
	horizontalBox := container.NewHBox(btnOne, labelOne)

	w.SetContent(
		horizontalBox,
	)

	w.ShowAndRun()
}
