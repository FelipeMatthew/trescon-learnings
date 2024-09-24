package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())

	w := a.NewWindow("Dark & Light theme")
	w.Resize(fyne.NewSize(400, 400))

	label := widget.NewLabel("Choose your system color")
	label.TextStyle.Bold = true

	btnLight := widget.NewButton("Light", func() {
		a.Settings().SetTheme(theme.LightTheme())
	})

	btnDark := widget.NewButton("Dark", func() {
		a.Settings().SetTheme(theme.DarkTheme())

	})

	btnExit := widget.NewButton("Exit", func() {
		a.Quit()
	})

	w.SetContent(
		container.NewVBox(
			label,
			btnLight,
			btnDark,
			btnExit,
		),
	)

	w.ShowAndRun()
}
