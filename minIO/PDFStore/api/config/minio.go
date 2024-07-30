package config

import (
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func MinioInit() {
	var err error
	endpoint := os.Getenv("MINIO_URL")
	accessKeyID := os.Getenv("MINIO_AKEY")
	secretAccessKey := os.Getenv("MINIO_SKEY")
	useSSL := false

	MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalf("Failed to initialize MinIO client: %v", err)
	}
}

// TODO
func GenerateBucket() {

}
