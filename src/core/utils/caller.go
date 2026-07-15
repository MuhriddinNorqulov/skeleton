package utils

import (
	"fmt"
	"runtime"
)

// CallerPath chaqiruvchining fayl yo'li va satr raqamini qaytaradi.
func CallerPath(skip int) string {
	_, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return "unknown"
	}
	return fmt.Sprintf("file://%s:%d", file, line)
}
