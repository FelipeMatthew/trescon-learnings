package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("")

	label := widget.NewLabel("hi")
	w.SetContent(label)
	// Caso essa janela seja fechada, para a aplicação
	w.SetMaster()
	// Exibir duas páginas ao mesmo tempo
	w.Show()

	w2 := a.NewWindow("Larger")
	w2.SetContent(widget.NewButton("Open new", func() {
		w3 := a.NewWindow("third")
		w3.SetContent(widget.NewLabel("third page"))
		w3.Show()
	}))
	w2.Resize(fyne.NewSize(100, 100))
	w2.Show()

	// Roda toda a aplicação
	a.Run()
}

func openWindow(a fyne.App) {
	w3 := a.NewWindow("third")
	w3.SetContent(widget.NewLabel("third page"))
	w3.Show()
}
