package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func DashboardDataPage(uid string, window fyne.Window) fyne.CanvasObject {

	originalImgText := canvas.NewText("Imagem original:", color.White)
	originalImgText.TextStyle.Bold = true
	originalImgText.TextSize = 14

	// todo: show image

	processedImgText := canvas.NewText(fmt.Sprintf("O MEU UID: %v", uid), color.White)
	processedImgText.TextStyle.Bold = true
	processedImgText.TextSize = 14

	// todo: show image

	content := container.NewVBox(
		originalImgText,
		// img
		layout.NewSpacer(),
		processedImgText,
		// img
	)

	return content

}
