package ports

import (
	"mime/multipart"
)

type ProjectInventory interface {
	UploadProjectImage(file multipart.File, projID string) error
}
