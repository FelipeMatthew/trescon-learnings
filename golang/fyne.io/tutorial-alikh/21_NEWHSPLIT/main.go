package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("NewHSplit")
	w.Resize(fyne.NewSize(400, 400))

	labelOne := widget.NewLabel("New label")
	textOne := canvas.NewText("Text one", color.White)

	icon := widget.NewIcon(theme.CancelIcon())
	textTwo := canvas.NewText("Right side from", color.White)
	button := widget.NewButton("Click", func() {})

	w.SetContent(
		container.NewHSplit(
			container.NewVBox(
				labelOne,
				textOne,
			),
			container.NewVBox(
				icon, textTwo, button,
			),
		),
	)

	w.ShowAndRun()
}
