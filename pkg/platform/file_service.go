package platform

import (
	"context"
	"mime/multipart"
)

type FileService interface {
	Upload(ctx context.Context, filePath string, file *multipart.FileHeader) error
	Delete(ctx context.Context, filePath string) error
}
