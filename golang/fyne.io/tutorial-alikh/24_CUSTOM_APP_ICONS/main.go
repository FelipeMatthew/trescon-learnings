package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("How to change App ICON")

	// *
	currentDir, _ := os.Getwd()
	imagePath := filepath.Join(currentDir, "24_CUSTOM_APP_ICONS", "assets", "images", "cat.png")

	r, _ := LoadResourceFromPath(imagePath)
	w.SetIcon(r)

	w.Resize(fyne.NewSize(400, 400))
	w.ShowAndRun()
}

type Resource interface {
	Name() string
	Content() []byte
}

type StaticResource struct {
	StaticName    string
	StaticContent []byte
}

func (r *StaticResource) Name() string {
	return r.StaticName
}

func (r *StaticResource) Content() []byte {
	return r.StaticContent
}

func NewStaticResource(name string, content []byte) *StaticResource {
	return &StaticResource{
		StaticName:    name,
		StaticContent: content,
	}
}

func LoadResourceFromPath(path string) (Resource, error) {
	bytes, err := ioutil.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	name := filepath.Base(path)
	return NewStaticResource(name, bytes), nil
}
