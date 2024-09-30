package ui

import (
	"image/color"
	"log"
	"total_count/internal/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func DashboardPage(window fyne.Window) fyne.CanvasObject {

	titleText := canvas.NewText("Dashboard: ", color.White)
	titleText.TextSize = 24
	titleText.TextStyle.Bold = true

	// Fetching API data
	imageData, err := services.FetchImagesData()
	if err != nil {
		log.Printf("Erro ao buscar dados da API: %v", err)
		return container.NewVBox(titleText, canvas.NewText("Erro ao carregar os dados", color.RGBA{255, 0, 0, 255}))
	}

	// Slice para armazenar os cards

	var cardList []fyne.CanvasObject

	for _, image := range imageData {
		// Criar os cards com base nos dados da API
		card := Card(image.Code, image.Description, image.Count, image.Timestamp, window)
		cardList = append(cardList, container.NewPadded(card))
	}

	cards := container.NewVBox(
		container.NewGridWithColumns(3, cardList...),
	)

	content := container.NewVBox(
		titleText,
		cards,
	)

	return content
}
