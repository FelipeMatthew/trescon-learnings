package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"total_count/internal/api/models"
	"total_count/internal/api/routes"
)

func PostImage(code, description, base64Image string) (*models.CountItemsResponse, error) {
	// cria os dados com base nos parametros
	item := models.CountItems{
		Code:        code,
		Description: description,
		Image:       base64Image,
	}

	// obj to json
	itemJson, err := json.Marshal(item)
	if err != nil {
		return nil, fmt.Errorf("error ao converter obj para json: %v", err)
	}

	// API URL
	countImageRoute := routes.NewApiRoutes().CountItens

	response, err := http.Post(countImageRoute, "application/json", bytes.NewBuffer(itemJson))
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("falha ao enviar imagem: %v", response.StatusCode)
	}

	var serverResult models.CountItemsResponse
	json.NewDecoder(response.Body).Decode(&serverResult)

	return &serverResult, nil
}
