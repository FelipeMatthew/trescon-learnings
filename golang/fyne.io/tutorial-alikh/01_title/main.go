package main

import "fyne.io/fyne/v2/app"

func main() {
	a := app.New()

	title := "i'm a title"
	w := a.NewWindow(title)

	w.ShowAndRun()
}
