package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Card(prodCode, description string, qtd int, timestamp string, window fyne.Window) *fyne.Container {

	// Code
	productCode := canvas.NewText(fmt.Sprintf("Código: %v", prodCode), color.White)
	productCode.TextSize = 18
	productCode.TextStyle.Bold = true

	// Description
	descriptionText := canvas.NewText(fmt.Sprintf("Descrição: %v", description), color.White)
	descriptionText.TextSize = 14

	// Date
	dataText := canvas.NewText(fmt.Sprintf("Data: %v", timestamp), color.White)
	dataText.TextSize = 12
	dataText.Color = color.RGBA{150, 150, 150, 255}

	// QTD
	quantityText := canvas.NewText(fmt.Sprintf("Quantidade encontrada: %v", qtd), color.White)
	quantityText.TextSize = 12

	// Card
	cardBackground := canvas.NewRectangle(color.RGBA{30, 30, 30, 255})
	cardBackground.SetMinSize(fyne.NewSize(200, 100))

	// Invisible btn
	cardBtn := widget.NewButton("", func() {
		window.SetContent(DashboardDataPage(window))
	})
	cardBtn.Importance = widget.LowImportance

	cardContent := container.NewVBox(
		productCode,
		descriptionText,
		quantityText,
		dataText,
	)

	card := container.NewStack(
		cardBackground,
		container.NewPadded(cardContent),
		cardBtn,
	)

	return card
}
