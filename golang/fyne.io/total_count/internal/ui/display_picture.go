package ui

import (
	"fmt"
	"image/color"
	"log"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func DisplayPictureScreen(imagePath string, count, confidence int, window fyne.Window) fyne.CanvasObject {

	imageResp, err := http.Get(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer imageResp.Body.Close()

	image := canvas.NewImageFromReader(imageResp.Body, "image")
	image.FillMode = canvas.ImageFillContain

	titleText := canvas.NewText("Imagem processada", color.White)
	titleText.TextSize = 24
	titleText.TextStyle.Bold = true

	confidenceText := canvas.NewText(fmt.Sprintf("Confiança: %v%%", confidence), color.White)
	confidenceText.TextSize = 16

	countText := canvas.NewText(fmt.Sprintf("Unidades encontradas: %v", count), color.White)
	countText.TextSize = 16

	content := container.NewVBox(titleText, confidenceText, countText, image)

	return content
}
