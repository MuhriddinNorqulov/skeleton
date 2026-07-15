package context

import (
	"context"
	"encoding/base64"
	"errors"
	"io"
	"strings"

	"example.com/PROJECT_NAME/src/core/application/response"
	"example.com/PROJECT_NAME/src/core/domain/entity"
	"example.com/PROJECT_NAME/src/core/domain/ports/httpport/ctx"
	"example.com/PROJECT_NAME/src/core/domain/ports/httpport/request"
	"example.com/PROJECT_NAME/src/infrastructure/echohttp/requestimpl"

	"github.com/labstack/echo/v4"
)

type EchoContext struct {
	user           *entity.UserEntity
	IdempotencyKey string
	echo.Context
}

func NewEchoContext(c echo.Context) ctx.Context {
	return &EchoContext{Context: c}
}

func (this *EchoContext) GetContext() context.Context {
	return this.Context.Request().Context()
}

func (this *EchoContext) Param(name string) string {
	return this.Context.Param(name)
}

func (this *EchoContext) QueryParam(name string) string {
	return this.Context.QueryParam(name)
}

func (this *EchoContext) SetUser(user *entity.UserEntity) {
	this.user = user
}

func (this *EchoContext) User() *entity.UserEntity {
	return this.user
}

func (this *EchoContext) JSON(status int, i any) error {
	return this.Context.JSON(status, i)
}

func (this *EchoContext) Success(status int, i any) error {
	payload := response.NewResponse(response.CodeSuccess, true, i, "")
	return this.JSON(status, payload)
}

func (this *EchoContext) GetRequest() *request.Request {
	header := this.Context.Request().Header
	h := requestimpl.NewHeaderImpl(header)
	return &request.Request{Header: h}
}

func (this *EchoContext) GetIdempotencyKey() string {
	return this.IdempotencyKey
}

func (this *EchoContext) SetIdempotencyKey(key string) {
	this.IdempotencyKey = key
}

func (this *EchoContext) StreamFile(filename, contentType string, r io.Reader) error {
	this.Context.Response().Header().Set("Content-Disposition", `attachment; filename="`+filename+`"`)
	return this.Context.Stream(200, contentType, r)
}

func (this *EchoContext) GetBasicCredential() (string, string, error) {
	auth := this.Context.Request().Header.Get("Authorization")
	if auth == "" || !strings.HasPrefix(auth, "Basic ") {
		return "", "", errors.New("authorization header missing or invalid")
	}
	decoded, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(auth, "Basic "))
	if err != nil {
		return "", "", errors.New("invalid authorization header")
	}
	pair := strings.SplitN(string(decoded), ":", 2)
	if len(pair) != 2 {
		return "", "", errors.New("invalid authorization header")
	}
	return pair[0], pair[1], nil
}
