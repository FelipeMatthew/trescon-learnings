package main

import (
	"os"
	"path/filepath"
	"total_count/internal/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
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

	// Page navigation
	splashPage := ui.SplashPage(myWindow)

	myWindow.SetContent(splashPage)

	myWindow.ShowAndRun()
}

func imageFilePath(fileName string) string {
	currentDirectory, _ := os.Getwd()
	imagePath := filepath.Join(currentDirectory, "assets", "images", fileName)
	return imagePath
}
