package network

import "io"

type MultipartFile struct {
	Field  string
	Name   string
	Reader io.Reader
}

type MultipartForm struct {
	Files  []MultipartFile
	Fields map[string]string
}
