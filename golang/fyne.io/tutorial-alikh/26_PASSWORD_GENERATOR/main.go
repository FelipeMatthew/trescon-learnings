package main

import (
	"image/color"
	"math/rand"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Password generator")
	w.Resize(fyne.NewSize(400, 400))

	textOne := canvas.NewText("Password generator", color.White)
	textOne.TextSize = 25

	// Password length
	input := widget.NewEntry()
	input.SetPlaceHolder("Password length")

	result := canvas.NewText("", color.White)

	btn := widget.NewButton("Generate password", func() {
		passlength, _ := strconv.Atoi(input.Text)
		password := RandomCharacteres(passlength)
		result.Text = password
		result.Refresh()
	})

	w.SetContent(container.NewVBox(textOne, input, btn, result))

	w.ShowAndRun()
}

func RandomCharacteres(passLength int) string {
	// Lógica:
	// Vai pegar e fazer um for e que dentro desse for vai ter as 4 categorias: lowercase, uppercase...
	// Após isso ele vai ter um switch que cada vez que ele fizer ele vai adicionar um caracter na senha e assim vai se somando...

	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	specialChar := "~!@#$%^&*()_-?><,."

	passwordLenght := passLength
	password := ""

	for n := 0; n < passwordLenght; n++ {
		randNum := rand.Intn(4)

		switch randNum {
		case 0:
			randNum := rand.Intn(len(lowerCase))
			password = password + string(lowerCase[randNum])
		case 1:
			randNum := rand.Intn(len(upperCase))
			password = password + string(upperCase[randNum])
		case 2:
			randNum := rand.Intn(len(numbers))
			password = password + string(numbers[randNum])
		case 3:
			randNum := rand.Intn(len(specialChar))
			password = password + string(specialChar[randNum])
		}
	}

	return password

}
