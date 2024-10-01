package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"total_count/internal/api/models"
	"total_count/internal/api/routes"
)

func PostImage(code, description, base64Image string) error {
	// cria os dados com base nos parametros
	item := models.CountItems{
		Code:        code,
		Description: description,
		Image:       base64Image,
	}

	// obj to json
	itemJson, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("error ao converter obj para json: %v", err)
	}

	// API URL
	countImageRoute := routes.NewApiRoutes().CountItens

	response, err := http.Post(countImageRoute, "application/json", bytes.NewBuffer(itemJson))
	if err != nil {
		return fmt.Errorf("erro ao fazer requisição: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("falha ao enviar imagem: %v", response.StatusCode)
	}

	var serverResult map[string]interface{}
	json.NewDecoder(response.Body).Decode(&serverResult)
	fmt.Println("Server result: ", serverResult)

	return nil
}
