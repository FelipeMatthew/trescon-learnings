package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Card")
	w.Resize(fyne.NewSize(400, 400))

	card := widget.NewCard(
		"Hello from card",
		"In card you can write everything",
		canvas.NewCircle(color.Black),
	)

	w.SetContent(
		container.NewHBox(card),
	)

	w.ShowAndRun()
}
