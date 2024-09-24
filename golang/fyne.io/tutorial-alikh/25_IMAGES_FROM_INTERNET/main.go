package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Images from internet")
	w.Resize(fyne.NewSize(400, 400))

	label := widget.NewLabel("Getting icon from internet")

	image, _ := fyne.LoadResourceFromURLString("https://static-00.iconduck.com/assets.00/golang-icon-1594x2048-0xixr8zr.png")

	img := canvas.NewImageFromResource(image)
	img.SetMinSize(fyne.NewSize(300, 300))

	w.SetContent(container.NewVBox(label, img))

	w.ShowAndRun()
}
