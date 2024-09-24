package main

import (
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Random color generator")
	w.Resize(fyne.NewSize(800, 400))

	// Lógic
	colorx := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	rectangle := canvas.NewRectangle(colorx)
	rectangle.SetMinSize(fyne.NewSize(300, 300))

	btnRandom := widget.NewButton("Generate random color", func() {
		rectangle.FillColor = color.NRGBA{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 255}
		rectangle.Refresh()
	})

	btnRed := widget.NewButton("Random RED", func() {
		rectangle.FillColor = color.NRGBA{R: uint8(rand.Intn(255)), G: 0, B: 0, A: 255}
		rectangle.Refresh()
	})

	btnGreen := widget.NewButton("Random GREEN", func() {
		rectangle.FillColor = color.NRGBA{R: 0, G: uint8(rand.Intn(255)), B: 0, A: 255}
		rectangle.Refresh()
	})

	btnBlue := widget.NewButton("Random BLUE", func() {
		rectangle.FillColor = color.NRGBA{R: 0, G: 0, B: uint8(rand.Intn(255)), A: 255}
		rectangle.Refresh()
	})

	w.SetContent(container.NewVBox(
		rectangle,
		btnRandom,
		btnRed,
		btnGreen,
		btnBlue,
	))

	w.ShowAndRun()
}
