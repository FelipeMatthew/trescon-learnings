package main

import (
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())

	myWindow := myApp.NewWindow("Total count")
	myWindow.Resize(fyne.NewSize(700, 500))
	myWindow.CenterOnScreen()

	imgPath := imageFilePath("half_logo.png")
	icon, _ := fyne.LoadResourceFromPath(imgPath)

	myWindow.SetIcon(icon)

	myWindow.SetContent(widget.NewLabel("Testing air toml"))

	myWindow.ShowAndRun()
}

func imageFilePath(fileName string) string {
	currentDirectory, _ := os.Getwd()
	imagePath := filepath.Join(currentDirectory, "assets", "images", fileName)
	return imagePath
}
