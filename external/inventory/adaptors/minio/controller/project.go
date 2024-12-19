package inventory

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	ports "project-management-service/external/_ports"

	"github.com/minio/minio-go/v7"
)

type projectInventoryMinIO struct {
	client *minio.Client
}

func NewProjectInventoryMinIO(client *minio.Client) ports.ProjectInventory {
	return &projectInventoryMinIO{
		client: client,
	}
}

func (s *projectInventoryMinIO) UploadProjectImage(file multipart.File, projID string) error {
	if file == nil {
		return fmt.Errorf("file cannot be nil")
	}

	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "upload-*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up the temporary file after use

	// Copy the file content into the temporary file
	_, err = io.Copy(tmpFile, file)
	if err != nil {
		return fmt.Errorf("failed to copy file content to temp file: %w", err)
	}

	// Generate a unique object name
	objectName := fmt.Sprintf("images/%s.png", projID)

	// Upload the temporary file to MinIO using FPutObject
	_, err = s.client.FPutObject(context.Background(), "project", objectName, tmpFile.Name(), minio.PutObjectOptions{
		ContentType: "application/octet-stream", // Set the content type, this can be dynamic based on the file
	})
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	return nil
}
