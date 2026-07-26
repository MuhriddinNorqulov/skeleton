package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/muhriddinnorqulov/skeleton/src/core/application/response"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/storage"
	"github.com/muhriddinnorqulov/skeleton/src/core/utils"
	"github.com/muhriddinnorqulov/skeleton/src/infrastructure/env"
)

type LocalFileStorage struct {
	bucketPath string
}

// @inject
func NewLocalFileStorage(env *env.Env) storage.FileStorage {
	abs, err := filepath.Abs(env.StoragePath)
	if err != nil {
		panic(fmt.Sprintf("invalid storage path: %v", err))
	}
	if err := os.MkdirAll(abs, 0755); err != nil {
		panic(fmt.Sprintf("cannot create storage directory: %v", err))
	}
	return &LocalFileStorage{bucketPath: abs}
}

func (this *LocalFileStorage) resolve(objectPath string) (string, error) {
	cleaned := filepath.Clean(objectPath)
	full := filepath.Join(this.bucketPath, cleaned)
	if !strings.HasPrefix(full, this.bucketPath+string(filepath.Separator)) && full != this.bucketPath {
		return "", response.NewSafeError(response.CodeFileError,
			fmt.Errorf("access denied: path %q is outside storage bucket", objectPath), utils.CallerPath(1))
	}
	return full, nil
}

func (this *LocalFileStorage) ReadFile(_ context.Context, objectPath string) ([]byte, error) {
	full, err := this.resolve(objectPath)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(full)
}

func (this *LocalFileStorage) FileSize(_ context.Context, objectPath string) (int64, error) {
	full, err := this.resolve(objectPath)
	if err != nil {
		return 0, err
	}
	info, err := os.Stat(full)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

func (this *LocalFileStorage) PutByteObject(_ context.Context, objectName string, b []byte, _ string) error {
	full, err := this.resolve(objectName)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		return err
	}
	return os.WriteFile(full, b, 0644)
}

func (this *LocalFileStorage) PutStream(_ context.Context, objectPath string, r io.Reader, _ int64, _ string) error {
	full, err := this.resolve(objectPath)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		return err
	}
	f, err := os.Create(full)
	if err != nil {
		return err
	}
	defer utils.Close(f)
	if _, err = io.Copy(f, r); err != nil {
		return response.NewSafeError(response.CodeFileError,
			fmt.Errorf("cannot write file %q: %w", objectPath, err), utils.CallerPath(1))
	}
	return nil
}

func (this *LocalFileStorage) TempFile(pattern string) (string, error) {
	tmpDir := filepath.Join(this.bucketPath, "tmp")
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return "", err
	}
	f, err := os.CreateTemp(tmpDir, pattern)
	if err != nil {
		return "", err
	}
	return f.Name(), f.Close()
}

func (this *LocalFileStorage) RemoveFile(path string) error {
	return os.Remove(path)
}

func (this *LocalFileStorage) AbsPath(objectPath string) string {
	full, _ := this.resolve(objectPath)
	return full
}
