// Todo: API connection + communication

package api

type PostImage struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Image       string `json:"image"` // BASE64 image
}

// Opicional por enquanto: 27/09
