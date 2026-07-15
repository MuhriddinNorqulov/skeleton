package storage

import (
	"context"
	"io"
)

type FileStorage interface {
	ReadFile(ctx context.Context, objectPath string) ([]byte, error)
	FileSize(ctx context.Context, objectPath string) (int64, error)
	PutByteObject(ctx context.Context, objectName string, b []byte, contentType string) error
	PutStream(ctx context.Context, objectPath string, r io.Reader, size int64, contentType string) error
	TempFile(pattern string) (string, error)
	RemoveFile(path string) error
	AbsPath(objectPath string) string
}
