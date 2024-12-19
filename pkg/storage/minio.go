package storage

import (
	"context"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type minioStorage struct {
	client *minio.Client
}

func ConnectToMinIO() *minioStorage {
	endpoint := "localhost:9000"
	accessKeyID := os.Getenv("MINIO_ROOT_USER")
	secretAccessKey := os.Getenv("MINIO_ROOT_PASSWORD")
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("Failed to initialize MinIO client: %v", err)
	}

	log.Println("Connected to MinIO")

	return &minioStorage{client: minioClient}
}

func (storage *minioStorage) GetClient() *minio.Client {
	return storage.client
}

func (storage *minioStorage) MakeBucket(ctx context.Context, bucketName string) {
	if bucketName == "" {
		log.Fatalf("bucket name cannot be empty")
	}

	err := storage.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "ap-southeast-1"})
	if err != nil {
		// Check if bucket already exists
		exists, errBucketExists := storage.client.BucketExists(ctx, bucketName)
		if errBucketExists != nil {
			log.Fatalf("error checking if bucket exists: %v", errBucketExists)
		}
		if !exists {
			log.Fatalf("failed to create bucket: %v", err)
		}
		log.Printf("Bucket %s already exists.\n", bucketName)
	} else {
		log.Printf("Successfully created bucket %s.\n", bucketName)
	}
}
