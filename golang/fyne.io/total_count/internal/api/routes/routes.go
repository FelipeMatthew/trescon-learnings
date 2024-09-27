package routes

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ApiRoutes struct {
	ApiBaseUrl    string
	GetImagesData string
	GetImage      string
	CountItens    string
}

func NewApiRoutes() *ApiRoutes {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	apiBaseUrl := os.Getenv("API_URL")

	return &ApiRoutes{
		ApiBaseUrl:    apiBaseUrl,
		GetImagesData: apiBaseUrl + "/get_images_data",
		GetImage:      apiBaseUrl + "/get_image",
		CountItens:    apiBaseUrl + "/count_itens",
	}
}
