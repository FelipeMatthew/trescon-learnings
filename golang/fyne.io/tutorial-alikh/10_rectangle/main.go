package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Rectangle")
	w.Resize(fyne.NewSize(400, 400))

	rectangle := canvas.NewRectangle(color.Black)
	rectangle.StrokeWidth = 10
	rectangle.StrokeColor = color.NRGBA{R: 255, G: 255, B: 0, A: 255}

	w.SetContent(rectangle)

	w.ShowAndRun()

}
