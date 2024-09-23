package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// https://api.openweathermap.org/data/2.5/weather?lat=44.34&lon=10.99&appid=e231a0e0896da40d1781e95dab337841
func main() {
	a := app.New()
	w := a.NewWindow("Weather App")
	w.Resize(fyne.NewSize(300, 400))

	_weather := HttpRequest()

	// UI
	labelOne := canvas.NewText("Weather API & Fyne.io", color.White)
	labelOne.TextStyle = fyne.TextStyle{Bold: true}
	labelOne.TextSize = 20

	labelTwo := canvas.NewText(fmt.Sprintf("Country: %s", _weather.Sys.Country), color.White)
	labelThree := canvas.NewText(fmt.Sprintf("Wind speed: %.2f", _weather.Wind.Speed), color.White)
	labelFour := canvas.NewText(fmt.Sprintf("Temperature fahrenheit: %.2f", _weather.Main.Temp), color.White)
	labelFive := canvas.NewText(fmt.Sprintf("Location: %s", _weather.Name), color.White)

	w.SetContent(
		container.NewVBox(
			labelOne,
			labelTwo,
			labelThree,
			labelFour,
			labelFive,
		),
	)

	w.ShowAndRun()
}

func HttpRequest() Weather_ {
	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?lat=-23.53&lon=-46.62&appid=e231a0e0896da40d1781e95dab337841")
	if err != nil {
		fmt.Print(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
	}

	weather, err := UnmarshalWeather(body)
	if err != nil {
		fmt.Print(err)
	}

	return weather
}

// https://app.quicktype.io/#l=Go
func UnmarshalWeather(data []byte) (Weather_, error) {
	var r Weather_
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather_) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather_ struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int64     `json:"timezone"`
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Cod        int64     `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
	SeaLevel  int64   `json:"sea_level"`
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type Weather struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
	Gust  float64 `json:"gust"`
}
