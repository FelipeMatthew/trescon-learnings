package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Child menu")
	myWindow.Resize(fyne.NewSize(400, 400))
	myWindow.CenterOnScreen()

	menuItemOne := fyne.NewMenuItem("Edit", nil)

	menuItemOne.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Text", nil),
		fyne.NewMenuItem("Phrase", nil),
		fyne.NewMenuItem("Poem", nil),
		fyne.NewMenuItem("Dialog", nil))

	menuItemTwo := fyne.NewMenuItem("Copy", nil)

	menuItemTwo.ChildMenu = fyne.NewMenu("",
		fyne.NewMenuItem("Magazine", nil),
		fyne.NewMenuItem("HQ", nil),
		fyne.NewMenuItem("PDF", nil),
		fyne.NewMenuItem("Jonjon", nil))

	menuItemThree := fyne.NewMenuItem("Past", nil)
	menuItemFour := fyne.NewMenuItem("About", nil)

	newMenuItemOne := fyne.NewMenu("file", menuItemOne, menuItemTwo, menuItemThree, menuItemFour)
	newMenuItemTwo := fyne.NewMenu("About", menuItemThree, menuItemFour)

	newMainMenu := fyne.NewMainMenu(newMenuItemOne, newMenuItemTwo)

	myWindow.SetMainMenu(newMainMenu)

	myWindow.ShowAndRun()
}
