package main

import (
	"fmt"
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
	w := a.NewWindow("Random number generator")
	w.Resize(fyne.NewSize(400, 400))

	textOne := canvas.NewText("Random number generator", color.White)
	textOne.TextSize = 30

	// Logic part
	result := canvas.NewText("", color.White)
	result.TextSize = 26

	btnOne := widget.NewButton("Generate", func() {
		randNum := rand.Intn(10000)
		result.Text = fmt.Sprint(randNum)
		// Important to current reload
		result.Refresh()
	})

	w.SetContent(container.NewVBox(
		textOne,
		btnOne,
		result,
	),
	)

	w.ShowAndRun()
}
