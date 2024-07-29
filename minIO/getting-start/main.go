package main

import (
	"context"
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	// MinIO - Config
	endpoint := "localhost:9000"
	acessKeyID := "admin"
	secretAccessKey := "password"
	useSSL := false

	// Getting start minio client
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(acessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("MinIO client initialized successfully")

	// creating a bucket
	bucketName := "mynewbucket"
	location := "us-east-1"

	// Inserting bucket
	err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check if bucket already exist
		exist, errBucketExist := minioClient.BucketExists(context.Background(), bucketName)
		if errBucketExist == nil && exist {
			log.Printf("Bucket already exists\n", bucketName)
		} else {
			log.Fatalln(err) // return the created bucket
			fmt.Println("Bucket created successfully \n", bucketName)
		}
	}

	// creating a file to insert on bucket
	objectName := "hello.txt"
	filePath := "C:\\Users\\felipe.nascimento\\Desktop\\Codes\\trescon-learnings\\minIO\\getting-start\\testfiles\\hello.txt"
	contentType := "application/text"

	// Inserting file
	info, err := minioClient.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("File added successfully: \nFile: %s of size: &d", objectName, info.Size)

}
