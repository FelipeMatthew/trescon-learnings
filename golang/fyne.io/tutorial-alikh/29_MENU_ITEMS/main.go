package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Menu items")
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.CenterOnScreen()

	// Menu items
	menuItemOne := fyne.NewMenuItem("New", func() { fmt.Println("New") })
	menuItemTwo := fyne.NewMenuItem("Save", func() { fmt.Println("Save") })
	menuItemThree := fyne.NewMenuItem("Edit", func() { fmt.Println("Edit") })

	newMenu := fyne.NewMenu("File", menuItemOne, menuItemTwo, menuItemThree)

	menu := fyne.NewMainMenu(newMenu)

	myWindow.SetMainMenu(menu)
	myWindow.SetContent(widget.NewLabel("My content"))

	myWindow.ShowAndRun()
}
