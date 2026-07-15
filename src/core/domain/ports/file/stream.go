package file

import "io"

type Stream struct {
	Body        io.ReadCloser
	Size        int64
	ContentType string
}
