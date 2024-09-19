package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Change color")
	w.Resize(fyne.NewSize(300, 300))

	colorx := color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	textX := canvas.NewText("my text", colorx)

	w.SetContent(textX)

	w.ShowAndRun()
}
