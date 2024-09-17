package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	// Define o formato ao qual vai exibir o relógio
	formatted := time.Now().Format("The time is: 03:04:05")
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Clock")

	clock := widget.NewLabel("")
	// Define o valor no setText
	updateTime(clock)

	w.SetContent(clock)
	// Atualizar o temporizador por segundo
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.ShowAndRun()
}
