package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Dice game")
	w.Resize(fyne.NewSize(500, 500))

	// Image
	currentDir, _ := os.Getwd()
	imagePath := filepath.Join(currentDir, "assets", "images", "01.png")

	titleLabel := canvas.NewText("Roll the dicee!", color.White)
	titleLabel.TextSize = 30

	result := canvas.NewImageFromFile(imagePath)
	result.FillMode = canvas.ImageFillContain
	result.SetMinSize(fyne.NewSize(250, 250))

	rollBtn := widget.NewButton("Play", func() {
		randNum := rand.Intn(6) + 1 // Ignoring 0
		newImagePath := filepath.Join(currentDir, "assets", "images", fmt.Sprintf("0%d.png", randNum))
		result.File = newImagePath
		result.Refresh()
	})

	w.SetContent(
		container.NewVBox(
			titleLabel,
			rollBtn,
			result,
		),
	)

	w.ShowAndRun()

}
