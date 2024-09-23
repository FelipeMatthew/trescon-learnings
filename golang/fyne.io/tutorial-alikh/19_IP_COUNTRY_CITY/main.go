package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("IP Finder App")
	w.Resize(fyne.NewSize(400, 400))

	// UI
	labelOne := widget.NewLabel("What's my IP?")
	labelTwo := widget.NewLabel("Find your IP:")

	// Output
	labelIp := widget.NewLabel("IP: ")
	labelCity := widget.NewLabel("City: ")
	labelCountry := widget.NewLabel("Country: ")

	btn := widget.NewButton("Find", func() {
		labelIp.Text = fmt.Sprintf("IP: %v", myIP())
		labelCity.Text = fmt.Sprintf("City: %v", myCity())
		labelCountry.Text = fmt.Sprintf("Country: %v", myCountry())

		labelIp.Refresh()
		labelCity.Refresh()
		labelCountry.Refresh()
	})

	w.SetContent(
		container.NewVBox(
			labelOne,
			labelTwo,
			btn,
			labelIp,
			labelCity,
			labelCountry,
		),
	)

	w.ShowAndRun()
}

func myIP() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var data Data
	json.Unmarshal(body, &data)
	return data.Query
}

func myCountry() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var data Data
	json.Unmarshal(body, &data)
	return data.Country
}

func myCity() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}

	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var data Data
	json.Unmarshal(body, &data)
	return data.City
}

type Data struct {
	Query   string
	Country string
	City    string
}
