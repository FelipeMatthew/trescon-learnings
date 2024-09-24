package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Submenu")
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.CenterOnScreen()

	menuItemOne := fyne.NewMenuItem("New", func() { fmt.Println("New") })
	menuItemTwo := fyne.NewMenuItem("Copy", func() { fmt.Println("Copy") })
	menuItemThree := fyne.NewMenuItem("Paste", func() { fmt.Println("Paste") })
	menuItemFour := fyne.NewMenuItem("Delete", func() { fmt.Println("Delete") })

	newMenuOne := fyne.NewMenu("File", menuItemOne, menuItemTwo, menuItemThree, menuItemFour)
	newMenuTwo := fyne.NewMenu("About", menuItemOne, menuItemTwo, menuItemThree, menuItemFour)
	newMenuThree := fyne.NewMenu("Help", menuItemOne, menuItemTwo, menuItemThree, menuItemFour)

	mainMenu := fyne.NewMainMenu(newMenuOne, newMenuTwo, newMenuThree)

	myWindow.SetMainMenu(mainMenu)

	myWindow.ShowAndRun()
}
