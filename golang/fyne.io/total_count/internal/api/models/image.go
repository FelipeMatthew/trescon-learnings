package models

// Get all images
type ImageData struct {
	UID                string                 `json:"uid"`
	Code               string                 `json:"code"`
	Description        string                 `json:"description"`
	Timestamp          string                 `json:"timestamp"`
	Count              int                    `json:"count"`
	Confidence         int                    `json:"confidence"`
	RawImagePath       string                 `json:"raw_image_path"`
	PredictedImagePath string                 `json:"predicted_image_path"`
	EditedImagePath    string                 `json:"edited_image_path"`
	Coordinates        map[string]Coordinates `json:"coordinates"`
}

type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

// Post image
type CountItems struct {
	Code        string `json:"code"`
	Description string `json:"description"`
	Image       string `json:"image"` // BASE64 image
}

// Repsonse posted image
type CountItemsResponse struct {
	UID         string                 `json:"uid"`
	Count       int                    `json:"count"`
	ImagePath   string                 `json:"image_path"`
	Coordinates map[string]Coordinates `json:"coordinates"`
	Confidence  int                    `json:"confidence"`
	Message     string                 `json:"message"`
}

// Update Image
type EditImage struct {
	UID         string                 `json:"uid"`
	Coordinates map[string]Coordinates `json:"coordinates"`
}
