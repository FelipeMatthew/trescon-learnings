package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"total_count/internal/api/models"
	"total_count/internal/api/routes"
)

func FetchImagesData() ([]models.ImageData, error) {
	getImagesRoute := routes.NewApiRoutes().GetImagesData

	// Fazendo a requisição HTTP
	response, err := http.Get(getImagesRoute)
	if err != nil {
		return nil, fmt.Errorf("error to get images data: %v", err)
	}
	defer response.Body.Close()

	// Validando o status da resposta
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %v", response.StatusCode)
	}

	// Lendo o corpo da resposta
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("body error: %v", err)
	}

	// Estrutura para armazenar o JSON completo
	var result map[string]interface{}
	if err := json.Unmarshal(responseData, &result); err != nil {
		return nil, fmt.Errorf("error parsing json: %v", err)
	}

	// Extraindo o campo "images" e validando
	imagesInterface, ok := result["images"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid or missing 'images' field")
	}

	var imageList []models.ImageData

	// Percorrendo a lista de imagens
	for _, image := range imagesInterface {
		imageMap, ok := image.(map[string]interface{})
		if !ok {
			continue
		}

		// Extraindo as coordenadas
		coordinatesMap, ok := imageMap["coordinates"].(map[string]interface{})
		coordinates := make(map[string]models.Coordinates)

		if ok {
			for key, coord := range coordinatesMap {
				coordMap := coord.(map[string]interface{})
				coordinates[key] = models.Coordinates{
					X: int(coordMap["x"].(float64)),
					Y: int(coordMap["y"].(float64)),
					W: int(coordMap["w"].(float64)),
					H: int(coordMap["h"].(float64)),
				}
			}
		}

		imageList = append(imageList, models.ImageData{
			UID:                fmt.Sprintf("%v", imageMap["uid"]),
			Code:               fmt.Sprintf("%v", imageMap["code"]),
			Description:        fmt.Sprintf("%v", imageMap["description"]),
			Timestamp:          fmt.Sprintf("%v", imageMap["timestamp"]),
			Count:              int(imageMap["count"].(float64)),
			Confidence:         imageMap["confidence"].(float64),
			RawImagePath:       fmt.Sprintf("%v", imageMap["raw_image_path"]),
			PredictedImagePath: fmt.Sprintf("%v", imageMap["predicted_image_path"]),
			EditedImagePath:    fmt.Sprintf("%v", imageMap["edited_image_path"]),
			Coordinates:        coordinates,
		})
	}

	sort.Slice(imageList, func(i, j int) bool {
		return imageList[i].Timestamp > imageList[j].Timestamp
	})

	return imageList, nil
}
