package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func DisplayPictureScreen(window fyne.Window) fyne.CanvasObject {

	titleText := canvas.NewText("Visualizar imagem", color.White)

	content := container.NewVBox(titleText)

	return content
}
