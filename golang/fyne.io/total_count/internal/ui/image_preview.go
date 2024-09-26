package ui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ImagePreviewPage(window fyne.Window, imagePath string) fyne.CanvasObject {

	selectedImage := canvas.NewImageFromFile(imagePath)
	selectedImage.FillMode = canvas.ImageFillContain
	selectedImage.SetMinSize(fyne.NewSize(400, 400))

	textTitle := canvas.NewText("Imagem selecionada: ", color.White)
	textTitle.Alignment = fyne.TextAlignCenter
	textTitle.TextSize = 24

	backBtn := widget.NewButtonWithIcon("Voltar", theme.NavigateBackIcon(), func() {
		window.SetContent(UploadImagePage(window))
	})

	confirmButton := widget.NewButtonWithIcon("Confirmar", theme.ConfirmIcon(), func() {
		prodCode := widget.NewEntry()
		prodCode.PlaceHolder = "Ex: 1DC4D32A"

		description := widget.NewEntry()
		description.PlaceHolder = "Produto da loja A"

		form := dialog.NewForm("Confirmação de imagem", "Enviar", "Cancelar", []*widget.FormItem{

			widget.NewFormItem("Código do produto", prodCode),
			widget.NewFormItem("Descrição", description),
		}, func(confirm bool) {
			if confirm {
				fmt.Println("Confirmed: true")
				// TODO: convert image to base64 and send it to backend
			}
		}, window)

		form.Resize(fyne.NewSize(350, 200))

		form.Show()
	})

	content := container.NewVBox(
		textTitle,
		selectedImage,
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			backBtn,
			layout.NewSpacer(),
			confirmButton,
			layout.NewSpacer(),
		),
	)

	return content
}
