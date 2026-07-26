package ctx

import (
	"context"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/entity"
	"github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/httpport/request"
	"io"
	"mime/multipart"
)

type Context interface {
	GetContext() context.Context
	User() *entity.UserEntity
	SetUser(user *entity.UserEntity)

	GetIdempotencyKey() string
	SetIdempotencyKey(key string)

	JSON(status int, i any) error
	Success(status int, i any) error
	StreamFile(filename, contentType string, r io.Reader) error

	Param(name string) string
	QueryParam(name string) string
	FormFile(name string) (*multipart.FileHeader, error)

	Bind(i any) error
	Validate(i any) error

	GetRequest() *request.Request
	GetBasicCredential() (string, string, error)
}
