package ui

import (
	"image/color"
	"log"
	"time"
	"total_count/internal/api/services"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func DashboardPage(window fyne.Window) fyne.CanvasObject {

	titleText := canvas.NewText("Dashboard: ", color.White)
	titleText.TextSize = 24
	titleText.TextStyle.Bold = true

	cardsContainer := container.NewVBox()

	returnBtn := widget.NewButtonWithIcon("Voltar", theme.NavigateBackIcon(), func() {
		window.SetContent(UploadImagePage(window))
	})

	realtimeData := func() {
		// Fetching API data
		imageData, err := services.FetchImagesData()
		if err != nil {
			log.Printf("Erro ao buscar dados da API: %v", err)
			return
		}

		// TODO: Exibir texto de card vazio
		cardsContainer.Objects = nil

		// Slice para armazenar os cards
		var cardList []fyne.CanvasObject
		for _, image := range imageData {
			// Criar os cards com base nos dados da API
			card := Card(image.Code, image.Description, image.Count, image.Timestamp, window)
			cardList = append(cardList, container.NewPadded(card))
		}

		cardsContainer.Add(container.NewGridWithColumns(2, cardList...))
		cardsContainer.Refresh()
	}

	// Faz um gorotine que vai pegar e consultar a api a cada 5 segundos e trazer os dados atualizados
	// TODO: aprofundar estudos em go func
	refreshApiReq := time.NewTicker(10 * time.Second)
	go func() {
		for range refreshApiReq.C {
			realtimeData()
		}
	}()

	realtimeData()

	verticalContainer := container.NewVScroll(cardsContainer)
	verticalContainer.SetMinSize(fyne.NewSize(700, 470))

	content := container.NewVBox(
		container.NewHBox(
			returnBtn,
			titleText,
		),
		verticalContainer,
	)

	return content
}
