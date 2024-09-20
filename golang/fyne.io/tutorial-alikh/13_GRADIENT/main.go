package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Gradient")
	w.Resize(fyne.NewSize(400, 400))

	// Degree angle last param
	gradient := canvas.NewLinearGradient(color.Black, color.White, 50)
	// gradient1 := canvas.NewVerticalGradient(color.Black, color.White)
	// gradient2 := canvas.NewHorizontalGradient(color.Black, color.White)
	// gradient3 := canvas.NewRadialGradient(color.Black, color.White)

	w.SetContent(gradient)

	w.ShowAndRun()
}
