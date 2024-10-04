package ui

import (
	"fmt"
	"image/color"
	"total_count/internal/api/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func DashboardDataPage(imageData models.ImageData, window fyne.Window) fyne.CanvasObject {

	originalImgText := canvas.NewText("Imagem original:", color.White)
	originalImgText.TextStyle.Bold = true
	originalImgText.TextSize = 14

	processedImgText := canvas.NewText(fmt.Sprintf("Imagem processada: %v", imageData.RawImagePath), color.White)
	processedImgText.TextStyle.Bold = true
	processedImgText.TextSize = 14

	// todo: show image

	editedImgText := canvas.NewText(fmt.Sprintf("Imagem atualizada: %v", imageData.Count), color.White)
	editedImgText.TextStyle.Bold = true
	editedImgText.TextSize = 14

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
