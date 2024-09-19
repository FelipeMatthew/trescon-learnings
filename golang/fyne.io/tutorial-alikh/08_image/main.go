package main

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Image")
	w.Resize(fyne.NewSize(500, 500))

	currDir, _ := os.Getwd()

	imagePath := filepath.Join(currDir, "assets", "images", "dog.png")

	image := canvas.NewImageFromFile(imagePath)

	w.SetContent(image)

	w.ShowAndRun()
}
