package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

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
	files := []string{
		"C:\\Users\\felipe.nascimento\\Desktop\\Codes\\trescon-learnings\\minIO\\getting-start\\testfiles\\hello.txt",
		"C:\\Users\\felipe.nascimento\\Desktop\\Codes\\trescon-learnings\\minIO\\getting-start\\testfiles\\file1.txt",
		"C:\\Users\\felipe.nascimento\\Desktop\\Codes\\trescon-learnings\\minIO\\getting-start\\testfiles\\file2.txt",
	}

	for _, myFilePath := range files {
		// Extract the file path name
		objectName := filepath.Base(myFilePath)
		contentType := "application/text"

		// Inserting file
		info, err := minioClient.FPutObject(context.Background(), bucketName, objectName, myFilePath, minio.PutObjectOptions{ContentType: contentType})
		if err != nil {
			log.Printf("Failed to upload %s: %v\n", myFilePath, err)
			continue
		}
		fmt.Printf("File %s added successfully with size %d bytes\n", objectName, info.Size)
	}
}
