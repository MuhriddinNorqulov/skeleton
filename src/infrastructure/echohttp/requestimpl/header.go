package requestimpl

import (
	"net/http"

	"example.com/PROJECT_NAME/src/core/domain/ports/httpport/request"
)

type HeaderImpl struct {
	header http.Header
}

func NewHeaderImpl(header http.Header) request.Header {
	return &HeaderImpl{header: header}
}

func (this *HeaderImpl) Get(key string) string {
	return this.header.Get(key)
}

func (this *HeaderImpl) Set(key string, value string) {
	this.header.Set(key, value)
}

func (this *HeaderImpl) Add(key, value string) {
	this.header.Add(key, value)
}
