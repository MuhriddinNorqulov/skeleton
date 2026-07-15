package file

import "context"

type Html2Pdf interface {
	ConvertA4(ctx context.Context, content string) (*Stream, error)
}
