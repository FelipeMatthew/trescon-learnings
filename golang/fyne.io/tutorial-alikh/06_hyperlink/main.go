package main

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hyperlink window")
	w.Resize(fyne.NewSize(600, 600))

	url, _ := url.Parse("https://www.google.com")
	hyperlink := widget.NewHyperlink("visit me", url)

	w.SetContent(hyperlink)

	w.ShowAndRun()
}
