package config

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func MinioInit() {
	var err error
	endpoint := "localhost:9000"
	acessKeyID := "admin"
	secretAccessKey := "password"
	useSSL := false

	MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(acessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalf("Failed to initialize MinIO client: %v", err)
	}
}

// TODO
func GenerateBucket() {

}
