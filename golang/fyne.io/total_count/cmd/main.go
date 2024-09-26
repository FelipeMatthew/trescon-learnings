package main

import (
	"total_count/internal/services"
	"total_count/internal/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()

	myWindow := myApp.NewWindow("Total count")
	myWindow.Resize(fyne.NewSize(700, 500))
	myWindow.CenterOnScreen()

	imgPath := services.GetImageFilePath("half_logo.png")
	icon, _ := fyne.LoadResourceFromPath(imgPath)
	myWindow.SetIcon(icon)

	// Page navigation
	splashPage := ui.SplashPage(myWindow)

	myWindow.SetContent(splashPage)

	myWindow.ShowAndRun()
}
