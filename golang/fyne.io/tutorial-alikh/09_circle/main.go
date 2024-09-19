package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Circle")
	w.Resize(fyne.NewSize(400, 400))

	circle1 := canvas.NewCircle(color.Black)

	w.SetContent(circle1)
	w.ShowAndRun()
}
