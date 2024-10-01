package ui

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func UploadImagePage(window fyne.Window) fyne.CanvasObject {
	textTitle := canvas.NewText("Enviar imagem para contagem de tubos:", color.White)
	textTitle.Alignment = fyne.TextAlignCenter
	textTitle.TextSize = 28

	folderIcon := widget.NewButtonWithIcon("Escolher imagem", theme.FolderIcon(), func() {

		// TODO: Validar se retorno é imagem
		// TODO: Converter imagem para base64
		dialog.NewFileOpen(func(ufs fyne.URIReadCloser, err error) {
			if err == nil || ufs != nil {
				defer ufs.Close()

				imgPath := ufs.URI().Path()

				validExtensions := []string{".png", ".jpg", ".jpeg"}
				isValidImage := false

				for _, ext := range validExtensions {
					// Valida se no caminho da imagem tem sufixo indicado acima
					// 1. string para ser validada, 2. o parametro
					if strings.HasSuffix(strings.ToLower(imgPath), ext) {
						isValidImage = true
						break
					}
				}

				if !isValidImage {
					dialog.ShowError(fmt.Errorf("arquivo inválido. Apenas imagens png e jpg são aceitas."), window)
					return
				}

				window.SetContent(ImagePreviewPage(window, imgPath))
			}
		}, window).Show()
	})

	pictureIcon := widget.NewButtonWithIcon("Dashboard", theme.ComputerIcon(), func() {
		window.SetContent(DashboardPage(window))
	})

	// TODO: Fix container layout
	content :=
		container.NewVBox(

			textTitle,
			layout.NewSpacer(),

			container.NewCenter(
				container.NewHBox(
					layout.NewSpacer(),
					folderIcon,
					layout.NewSpacer(),
					pictureIcon,
					layout.NewSpacer(),
				),
			),
		)

	return content
}
