package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetImageFromHttp(imagePath string) (string, error) {

	httpImagePath := fmt.Sprintf("http://%v", imagePath)

	imageResp, err := http.Get(httpImagePath)
	if err != nil {
		return "erro ao fazer http request", err
	}
	defer imageResp.Body.Close()

	// trabalhando com imagem temporária
	if imageResp.StatusCode != 200 {
		return "erro ao buscar imagem", err
	}

	tempFile, err := os.CreateTemp("", "image-*.png")
	if err != nil {
		return "erro ao criar arquivo temporário", err
	}
	defer os.Remove(tempFile.Name())

	_, err = io.Copy(tempFile, imageResp.Body)
	if err != nil {
		return "erro ao copiar arquivo", err
	}

	return tempFile.Name(), nil
}
