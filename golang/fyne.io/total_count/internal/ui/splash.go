package ui

import (
	"image/color"
	"total_count/internal/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func SplashPage(window fyne.Window) fyne.CanvasObject {
	text := canvas.NewText("Desenvolvido pela 3CON", color.White)
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = 24

	logoPath := services.GetImageFilePath("main-logo.png")

	logoImage := canvas.NewImageFromFile(logoPath)
	logoImage.SetMinSize(fyne.NewSize(200, 300))

	navigateBtn := widget.NewButton("Continuar", func() {
		homePage := HomePage(window)
		window.SetContent(homePage)
	})

	content := container.NewCenter(
		container.NewVBox(
			layout.NewSpacer(),
			text,
			logoImage,
			layout.NewSpacer(),
			navigateBtn,
			layout.NewSpacer(),
		),
	)

	return content
}
