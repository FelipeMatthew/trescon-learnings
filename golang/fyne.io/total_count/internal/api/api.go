// Todo: API connection + communication

package api

type PostImage struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Image       string `json:"image"` // BASE64 image
}

// Opicional por enquanto: 27/09
type EditImage struct {
	UID         string                 `json:"uid"`
	Coordinates map[string]Coordinates `json:"coordinates"`
}

type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}
