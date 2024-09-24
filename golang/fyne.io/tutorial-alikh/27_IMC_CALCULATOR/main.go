package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("IMC Calculator")
	w.Resize(fyne.NewSize(400, 400))

	titleText := canvas.NewText("Calculate your IMC", color.White)
	titleText.Alignment = fyne.TextAlignCenter
	titleText.TextStyle.Bold = true

	inputWeight := widget.NewEntry()
	inputWeight.SetPlaceHolder("88.9 kg")

	inputHeight := widget.NewEntry()
	inputHeight.SetPlaceHolder("170 cm")

	result := widget.NewLabel("Your imc is:  and you are: ")
	calculateBtn := widget.NewButton("Calculate", func() {
		height, _ := strconv.ParseFloat(inputHeight.Text, 32)
		weight, _ := strconv.ParseFloat(inputWeight.Text, 32)

		result.Text = imcCalc(float32(weight), float32(height))
		result.Refresh()
	})

	mainContainer := container.NewVBox(titleText, inputHeight, inputWeight, calculateBtn, result)

	w.SetContent(mainContainer)

	w.ShowAndRun()
}

func imcCalc(weight, height float32) string {
	height = height / 100
	imc := weight / (height * height)
	result := ""

	switch {
	case imc < 18.5:
		result = "you are underweight"
	case imc < 24.9:
		result = "you are normal, congrats"
	case imc < 29.9:
		result = "you are underweight"
	case imc > 29.9:
		result = "you have obesity, watch out!"
	default:
		result = "invalid IMC"
	}

	return fmt.Sprintf("Your imc is %.2f, %s", imc, result)
}
