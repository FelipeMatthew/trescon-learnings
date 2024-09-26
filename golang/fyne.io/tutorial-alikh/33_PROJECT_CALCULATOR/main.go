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
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculator")
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.CenterOnScreen()

	var valueOne int = 0
	var valueTwo int = 0
	var symbol string = ""
	var total string = ""
	var entryText string = ""

	entryOne := widget.NewEntry()
	entryOne.TextStyle = fyne.TextStyle{Bold: true}
	entryOne.SetPlaceHolder("Calculator...")

	btnOne := widget.NewButton("1", func() {
		total += "1"
		entryText += "1"
		entryOne.SetText(fmt.Sprint(entryText))
	})

	btnTwo := widget.NewButton("2", func() {
		total += "2"
		entryText += "2"
		entryOne.SetText(fmt.Sprint(entryText))
	})

	btnThree := widget.NewButton("3", func() {
		total += "3"
		entryText += "3"
		entryOne.SetText(fmt.Sprint(entryText))
	})

	btnFour := widget.NewButton("4", func() {
		total += "4"
		entryText += "4"
		entryOne.SetText(fmt.Sprint(entryText))
	})

	btnFive := widget.NewButton("5", func() {
		total += "5"
		entryText += "5"
		entryOne.SetText(fmt.Sprint(entryText))
	})

	btnSix := widget.NewButton("6", func() {
		total += "6"
		entryText += "6"
		entryOne.SetText(fmt.Sprint(entryText))
	})

	btnSeven := widget.NewButton("7", func() {
		total += "7"
		entryText += "7"
		entryOne.SetText(fmt.Sprint(entryText))
	})

	btnEight := widget.NewButton("8", func() {
		total += "8"
		entryText += "8"
		entryOne.SetText(fmt.Sprint(entryText))
	})

	btnNine := widget.NewButton("9", func() {
		total += "9"
		entryText += "9"
		entryOne.SetText(fmt.Sprint(entryText))
	})

	btnZero := widget.NewButton("0", func() {
		total += "0"
		entryText += "0"
		entryOne.SetText(fmt.Sprint(entryText))
	})

	result := canvas.NewText("Result", color.White)
	result.Alignment = fyne.TextAlignCenter
	result.TextSize = 24

	btnClear := widget.NewButton("Clear", func() {
		total = ""
		entryText = ""
		entryOne.SetText("")
		result.Text = ""
	})

	btnPlus := widget.NewButton("+", func() {
		valueOne, _ = strconv.Atoi(total)
		symbol = "+"
		entryText = total + "+"
		entryOne.SetText(fmt.Sprint(entryText))
		total = ""
	})

	btnMinus := widget.NewButton("-", func() {
		valueOne, _ = strconv.Atoi(total)
		symbol = "-"
		total = ""
		entryText = total + "-"
		entryOne.SetText(fmt.Sprint(entryText))
	})

	btnMultiply := widget.NewButton("*", func() {
		valueOne, _ = strconv.Atoi(total)
		symbol = "*"
		total = ""
		entryText = entryText + "*"
		entryOne.SetText(fmt.Sprint(entryText))
	})

	btnDivide := widget.NewButton("/", func() {
		valueOne, _ = strconv.Atoi(total)
		symbol = "/"
		total = ""
		entryText = entryText + "/"
		entryOne.SetText(fmt.Sprint(entryText))
	})

	btnEqual := widget.NewButton("=", func() {
		valueTwo, _ = strconv.Atoi(total)
		if symbol == "+" {
			myAnswer := (valueOne) + (valueTwo)
			result.Text = fmt.Sprint(myAnswer)
			result.Refresh()
		} else if symbol == "-" {
			myAnswer := (valueOne) - (valueTwo)
			result.Text = fmt.Sprint(myAnswer)
			result.Refresh()
		} else if symbol == "*" {
			myAnswer := (valueOne) * (valueTwo)
			result.Text = fmt.Sprint(myAnswer)
			result.Refresh()
		} else if symbol == "/" {
			myAnswer := (valueOne) / (valueTwo)
			result.Text = fmt.Sprint(myAnswer)
			result.Refresh()
		}
		total = ""
	})

	myWindow.SetContent(
		container.NewVBox(
			entryOne,
			result,
			container.NewGridWithColumns(
				3,
				btnOne,
				btnTwo,
				btnThree,
				btnFour,
				btnFive,
				btnSix,
				btnSeven,
				btnEight,
				btnNine,
				btnZero,
				btnEqual,
				btnPlus,
				btnMinus,
				btnMultiply,
				btnDivide,
				btnClear,
			),
		),
	)

	myWindow.ShowAndRun()
}
