package ui

import (
	"image/color"

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

		dialog.NewFileOpen(func(ufs fyne.URIReadCloser, err error) {
			if err == nil || ufs != nil {
				defer ufs.Close()
				imgPath := ufs.URI().Path()

				window.SetContent(ImagePreviewPage(window, imgPath))
			}
		}, window).Show()
	})

	pictureIcon := widget.NewButtonWithIcon("Tirar foto", theme.MediaPhotoIcon(), func() {
		// TODO: Criar função para gerar nome da imagem aleatóriamente
		// TODO: Validação se há ou não camera disponível
		// TODO: FIX IT
		// imageName := "/tmp/captured_image.png"
		// cmd := exec.Command("fswebcam", "--no-banner", imageName)
		// err := cmd.Run()
		// if err != nil {
		// 	dialog.ShowError(err, window)
		// 	return
		// }

		// window.SetContent(ImagePreviewPage(window, imageName))
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
