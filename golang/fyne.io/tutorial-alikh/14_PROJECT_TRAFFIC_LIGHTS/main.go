package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Traffic lights")
	w.Resize(fyne.NewSize(500, 500))
	fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())

	rectOne := canvas.NewCircle(color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	rectOne.Resize(fyne.NewSize(50, 50))
	rectTwo := canvas.NewCircle(color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	rectTwo.Resize(fyne.NewSize(50, 50))
	rectThree := canvas.NewCircle(color.NRGBA{R: 255, G: 255, B: 255, A: 255})
	rectThree.Resize(fyne.NewSize(50, 50))

	btnRed := widget.NewButton("Red", func() {
		reset(rectOne, rectTwo, rectThree)
		rectOne.FillColor = color.NRGBA{R: 255, G: 0, B: 0, A: 255}
		rectOne.Refresh()
	})
	btnYellow := widget.NewButton("Yellow", func() {
		reset(rectOne, rectTwo, rectThree)
		rectTwo.FillColor = color.NRGBA{R: 255, G: 255, B: 0, A: 255}
		rectTwo.Refresh()
	})

	btnGreen := widget.NewButton("Green", func() {
		reset(rectOne, rectTwo, rectThree)
		rectThree.FillColor = color.NRGBA{R: 0, G: 255, B: 0, A: 255}
		rectThree.Refresh()
	})

	btnReset := widget.NewButton("Reset", func() {
		rectOne.FillColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		rectTwo.FillColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		rectThree.FillColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
		rectOne.Refresh()
		rectTwo.Refresh()
		rectThree.Refresh()
	})

	// Container part
	w.SetContent(
		container.NewHSplit(
			container.NewGridWithRows(
				7,
				layout.NewSpacer(),
				rectOne,
				layout.NewSpacer(),
				rectTwo,
				layout.NewSpacer(),
				rectThree,
			),

			container.NewGridWithRows(
				9,
				layout.NewSpacer(),
				btnRed,
				btnYellow,
				btnGreen,
				layout.NewSpacer(),
				btnReset,
			),
		),
	)

	w.ShowAndRun()
}

func incrementMove(x uint8) (a uint8) {
	if x == 255 {
		x = 0
	} else if x == 0 {
		x = 255
	}
	return x
}

func reset(r1, r2, r3 *canvas.Circle) {
	r1.FillColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	r2.FillColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	r3.FillColor = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	r1.Refresh()
	r2.Refresh()
	r3.Refresh()
}
