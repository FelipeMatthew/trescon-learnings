package ui

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func DashboardPage(window fyne.Window) fyne.CanvasObject {

	titleText := canvas.NewText("Dashboard: ", color.White)
	titleText.TextSize = 24
	titleText.TextStyle.Bold = true

	// todo: make it dynamic

	cardOne := Card("123", "Produto 01", 3, time.Now(), window)
	cardTwo := Card("12ADV", "Produto 02", 255, time.Now(), window)
	cardThree := Card("134DDAA", "Produto 03", 245, time.Now(), window)
	cardFour := Card("13AD3", "Produto 04", 215, time.Now(), window)
	cardFive := Card("1AVV93", "Produto 05", 245, time.Now(), window)
	cardSix := Card("9D9AA", "Produto 06", 125, time.Now(), window)

	cards := container.NewVBox(
		container.NewGridWithColumns(3,
			container.NewPadded(cardOne),
			container.NewPadded(cardTwo),
			container.NewPadded(cardThree),
		),
		layout.NewSpacer(),
		container.NewGridWithColumns(3,
			container.NewPadded(cardFour),
			container.NewPadded(cardFive),
			container.NewPadded(cardSix),
		),
	)

	content := container.NewVBox(
		titleText,
		cards,
	)

	return content
}
