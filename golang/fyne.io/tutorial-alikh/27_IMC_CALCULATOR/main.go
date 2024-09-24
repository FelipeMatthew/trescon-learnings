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
	w := a.NewWindow("IMC Calculator")
	w.Resize(fyne.NewSize(400, 800))

	titleText := canvas.NewText("Calculate your IMC", color.White)

	inputWeight := widget.NewEntry()
	inputWeight.SetPlaceHolder("88.9 kg")

	inputHeight := widget.NewEntry()
	inputHeight.SetPlaceHolder("1.70 m")

	calculateBtn := widget.NewButton("Calculate", func() {

	})

	result := widget.NewLabel("You are ...")

	mainContainer := container.NewVBox(titleText, inputHeight, inputWeight, calculateBtn, result)

	w.SetContent(mainContainer)

	w.ShowAndRun()
}
