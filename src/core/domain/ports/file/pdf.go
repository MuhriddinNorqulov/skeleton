package file

import "io"

type PdfService interface {
	TotalPages(path string) (int, error)
	TotalPagesFromReader(rs io.ReadSeeker) (int, error)
}
