package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Main menu")

	myWindow.Resize(fyne.NewSize(400, 400))

	menuItem := &fyne.Menu{
		Label: "File",
		Items: nil,
	}

	menu := fyne.NewMainMenu(menuItem)

	myWindow.SetMainMenu(menu)
	myWindow.ShowAndRun()
}
