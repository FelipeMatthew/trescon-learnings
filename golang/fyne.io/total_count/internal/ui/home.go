package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func HomePage(window fyne.Window) fyne.CanvasObject {
	label := widget.NewLabel("Welcome to 3CON")

	content := container.NewVBox(
		label,
	)

	return content
}
