package logger

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
)

type GatewayLogger struct {
	*zap.Logger
}

// @inject
func NewGatewayLogger() *GatewayLogger {
	return &GatewayLogger{Logger: NewDailyFileZapLogger("logs/gateway/")}
}

func (this *GatewayLogger) LogRequest(method, url, contentType string, body []byte) {
	this.Info("gateway request",
		zap.String("method", method),
		zap.String("url", url),
		zap.String("content_type", contentType),
		zap.String("body", requestBody(contentType, body)),
	)
}

func (this *GatewayLogger) LogResponse(method, url string, statusCode int, contentType string, body []byte, duration time.Duration) {
	this.Info("gateway response",
		zap.String("method", method),
		zap.String("url", url),
		zap.Int("status", statusCode),
		zap.Duration("duration", duration),
		zap.String("body", responseBody(contentType, body)),
	)
}

func (this *GatewayLogger) LogError(method, url string, err error, duration time.Duration) {
	this.Error("gateway error",
		zap.String("method", method),
		zap.String("url", url),
		zap.Duration("duration", duration),
		zap.Error(err),
	)
}

const maxBodyLog = 512

func requestBody(contentType string, body []byte) string {
	if strings.HasPrefix(contentType, "multipart/") {
		return "[multipart form-data]"
	}
	if isBinary(contentType) {
		return "[binary]"
	}
	return truncateLog(string(body))
}

func responseBody(contentType string, body []byte) string {
	if isBinary(contentType) {
		return "[binary " + bytesSize(len(body)) + "]"
	}
	return truncateLog(string(body))
}

func truncateLog(s string) string {
	runes := []rune(s)
	if len(runes) <= maxBodyLog {
		return s
	}
	return string(runes[:maxBodyLog]) + fmt.Sprintf("... [+%d chars]", len(runes)-maxBodyLog)
}

func isBinary(contentType string) bool {
	for _, t := range []string{"application/pdf", "application/octet-stream", "image/"} {
		if strings.HasPrefix(contentType, t) {
			return true
		}
	}
	return false
}

func bytesSize(n int) string {
	switch {
	case n >= 1024*1024:
		return fmt.Sprintf("%d MB", n/(1024*1024))
	case n >= 1024:
		return fmt.Sprintf("%d KB", n/1024)
	default:
		return fmt.Sprintf("%d B", n)
	}
}
