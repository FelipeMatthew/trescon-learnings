package utils

import (
	"os"
	"path/filepath"
)

func GetImageFilePath(fileName string) string {
	currentDirectory, _ := os.Getwd()
	imagePath := filepath.Join(currentDirectory, "assets", "images", fileName)
	return imagePath
}
