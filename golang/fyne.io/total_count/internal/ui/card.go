package ui

import (
	"fmt"
	"image/color"
	"log"
	"total_count/internal/api/models"
	"total_count/internal/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func Card(imageData models.ImageData, window fyne.Window) *fyne.Container {

	// Code
	productCode := canvas.NewText(fmt.Sprintf("Código: %v", imageData.Code), color.White)
	productCode.TextSize = 18
	productCode.TextStyle.Bold = true

	// Description
	descriptionText := canvas.NewText(fmt.Sprintf("Descrição: %v", imageData.Description), color.White)
	descriptionText.TextSize = 16

	// Date
	// Converter timestamp
	formattedTime, err := utils.TimestampToDateTime(imageData.Timestamp)
	if err != nil {
		log.Printf("Erro ao converter timestamp: %v", err)
	}

	dataText := canvas.NewText(fmt.Sprintf("Data de criação: %v", formattedTime), color.White)
	dataText.TextSize = 14
	dataText.Color = color.RGBA{150, 150, 150, 255}

	// QTD
	quantityText := canvas.NewText(fmt.Sprintf("Quantidade encontrada: %v", imageData.Count), color.White)
	quantityText.TextSize = 12

	// Card
	cardBackground := canvas.NewRectangle(color.RGBA{30, 30, 30, 255})
	cardBackground.SetMinSize(fyne.NewSize(250, 150))

	// Invisible btn
	cardBtn := widget.NewButtonWithIcon("", theme.SearchIcon(), func() {
		window.SetContent(DashboardDataPage(imageData, window))
	})

	cardContent := container.NewVBox(
		productCode,
		layout.NewSpacer(),
		descriptionText,
		layout.NewSpacer(),

		quantityText,
		dataText,

		container.NewHBox(
			layout.NewSpacer(),
			cardBtn,
		),
	)

	card := container.NewStack(
		cardBackground,
		container.NewPadded(cardContent),
	)

	return card
}
