package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SplashPage(window fyne.Window) fyne.CanvasObject {
	label := widget.NewLabel("Developed by: 3CON")

	navigateBtn := widget.NewButton("Continuar", func() {
		homePage := HomePage(window)
		window.SetContent(homePage)
	})

	content := container.NewVBox(
		label,
		navigateBtn,
	)

	return content
}
