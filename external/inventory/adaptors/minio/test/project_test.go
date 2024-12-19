package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	inventory "project-management-service/external/inventory/adaptors/minio/controller"
	"project-management-service/pkg"
	"project-management-service/pkg/storage"

	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/assert"
)

func TestUploadImage(t *testing.T) {
	pkg.LoadEnv(".env.test")

	// Setup: Connect to MinIO and create the necessary bucket
	minioStorage := storage.ConnectToMinIO()
	minioStorage.MakeBucket(context.Background(), "project")
	projectInv := inventory.NewProjectInventoryMinIO(minioStorage.GetClient())

	// Specify the path to the local image file
	filePath := "test2.png" // Replace this with the actual path to your image file

	// Open the file from the local filesystem
	file, err := os.Open(filePath)
	assert.NoError(t, err, "Error opening file")
	defer file.Close()

	projectID := "9b4f5650-5aa7-46f9-a594-60bcbe81acc0"

	// Call the method and validate the result
	err = projectInv.UploadProjectImage(file, projectID)
	assert.NoError(t, err)

	// Generate object name based on timestamp for uniqueness
	objectName := fmt.Sprintf("%s.png", projectID)

	// Check if the object exists in the bucket
	client := minioStorage.GetClient()
	_, err = client.StatObject(context.Background(), "project", objectName, minio.StatObjectOptions{})
	assert.NoError(t, err, "The file should be uploaded to the MinIO bucket.")
}
