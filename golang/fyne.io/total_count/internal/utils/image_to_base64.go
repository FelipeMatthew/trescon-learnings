package utils

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

// abrir o arquivo, ler o conteúdo em binário e convertê-lo para uma string Base64.
func ConvertImageToBase64(imagePath string) (string, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return "", fmt.Errorf("erro ao abrir arquivo: %v", err)
	}
	defer file.Close()

	// Pega os bytes da imagem e assim podendo converter os bytes em string64
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("erro ao ler bytes da imagem: %v", err)
	}

	base64string := base64.StdEncoding.EncodeToString(fileBytes)
	return base64string, nil
}
