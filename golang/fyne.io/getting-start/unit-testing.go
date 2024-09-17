package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// widget.entry = input - textfield

func makeUI() (*widget.Label, *widget.Entry) {
	out, in := widget.NewLabel("Hello world!"), widget.NewEntry()

	// Assim que receber o texto alterado
	in.OnChanged = func(content string) {
		// Vai pegar o conteudo do texto e alterar no label inicial
		out.SetText("Hello " + content + "!")
	}

	return out, in
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello person")

	// Container
	w.SetContent(container.NewVBox(makeUI()))

	w.ShowAndRun()
}
