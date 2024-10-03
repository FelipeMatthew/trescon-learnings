package ui

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func DisplayPictureScreen(imagePath string, count int, confidence float64, window fyne.Window) fyne.CanvasObject {

	httpImagePath := fmt.Sprintf("http://%v", imagePath)

	imageResp, err := http.Get(httpImagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer imageResp.Body.Close()

	// trabalhando com imagem temporária
	if imageResp.StatusCode != 200 {
		log.Fatalf("erro ao buscar imagem, erro: %d", err)
	}

	tempFile, err := os.CreateTemp("", "image-*.png")
	if err != nil {
		log.Fatalf("erro ao criar arquivo temporario: %d", err)
	}
	defer os.Remove(tempFile.Name())

	_, err = io.Copy(tempFile, imageResp.Body)
	if err != nil {
		log.Fatalf("erro ao copiar arquivo temporario: %d", err)
	}

	image := canvas.NewImageFromFile(tempFile.Name())
	image.SetMinSize(fyne.NewSize(300, 300))
	image.FillMode = canvas.ImageFillContain

	titleText := canvas.NewText("Imagem processada", color.White)
	titleText.TextSize = 24
	titleText.TextStyle.Bold = true

	confidenceText := canvas.NewText(fmt.Sprintf("Confiança: %v%%", confidence), color.White)
	confidenceText.TextSize = 16
	confidenceText.Alignment = fyne.TextAlignCenter

	countText := canvas.NewText(fmt.Sprintf("Unidades encontradas: %v", count), color.White)
	countText.TextSize = 16
	countText.Alignment = fyne.TextAlignCenter

	dashboardBtn := widget.NewButtonWithIcon("Dashboard", theme.ComputerIcon(), func() {
		window.SetContent(DashboardPage(window))
	})

	uploadImage := widget.NewButtonWithIcon("Nova imagem", theme.MediaPhotoIcon(), func() {
		window.SetContent(UploadImagePage(window))
	})

	content := container.NewVBox(
		titleText,
		layout.NewSpacer(),
		image,
		confidenceText,
		countText,
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			dashboardBtn,
			layout.NewSpacer(),
			uploadImage,
			layout.NewSpacer(),
		),
	)

	return content
}
