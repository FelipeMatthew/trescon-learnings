package ui

import (
	"image/color"
	"total_count/internal/api/services"
	"total_count/internal/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func SplashPage(window fyne.Window) fyne.CanvasObject {

	title := canvas.NewText("Total count", color.White)
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true
	title.TextSize = 30

	text := canvas.NewText("Desenvolvido por 3CON", color.White)
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = 24

	logoPath := utils.GetImageFilePath("main-logo.png")

	logoImage := canvas.NewImageFromFile(logoPath)
	logoImage.SetMinSize(fyne.NewSize(350, 150))
	logoImage.FillMode = canvas.ImageFillContain

	navigateBtn := widget.NewButton("Continuar", func() {
		uploadImagePage := UploadImagePage(window)
		window.SetContent(uploadImagePage)
	})

	btnContainer := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		navigateBtn,
		layout.NewSpacer(),
	)

	content := container.NewVBox(
		layout.NewSpacer(),
		title,
		text,
		layout.NewSpacer(),
		logoImage,
		layout.NewSpacer(),
		btnContainer,
		layout.NewSpacer(),
	)

	services.PostImage()

	return content
}
