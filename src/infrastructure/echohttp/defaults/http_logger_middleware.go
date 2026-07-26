package defaults

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/muhriddinnorqulov/skeleton/src/core/application/response"
	"github.com/muhriddinnorqulov/skeleton/src/infrastructure/logger"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type HttpLoggerMiddleware struct {
	logger *logger.HttpLogger
}

// @inject
func NewHttpLoggerMiddleware(logger *logger.HttpLogger) *HttpLoggerMiddleware {
	return &HttpLoggerMiddleware{logger: logger}
}

func (this *HttpLoggerMiddleware) Wrap(next echo.HandlerFunc) echo.HandlerFunc {
	const maxBodyBytes = 64 * 1024

	return func(c echo.Context) error {
		req := c.Request()
		start := time.Now()

		// --- Request ID ---
		reqID := req.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.NewString()
		}

		// --- Logger ---
		lg := this.logger.With(zap.String("request_id", reqID))

		// --- Helpers ---
		baseFields := func() []zap.Field {
			return []zap.Field{
				zap.String("request_id", reqID),
				zap.String("method", req.Method),
				zap.String("path", req.URL.Path),
				zap.String("query", req.URL.RawQuery),
				zap.String("ip", c.RealIP()),
				zap.String("user_agent", req.UserAgent()),
				zap.Any("scope", "http"),
			}
		}

		// --- Request body ---
		var reqBody any
		contentType := req.Header.Get(echo.HeaderContentType)
		if req.Body != nil && req.Body != http.NoBody {
			if strings.HasPrefix(contentType, "multipart/") {
				reqBody = summarizeMultipart(req)
			} else {
				b, _ := io.ReadAll(io.LimitReader(req.Body, maxBodyBytes))
				reqBody = parseJSON(contentType, b)
				req.Body = io.NopCloser(bytes.NewReader(b))
			}
		}

		// --- Response recorder ---
		rec := &responseBodyRecorder{
			ResponseWriter: c.Response().Writer,
			limit:          maxBodyBytes,
		}
		c.Response().Writer = rec

		// --- Execute handler ---
		err := next(c)

		latency := time.Since(start)

		status := c.Response().Status
		if err != nil && !c.Response().Committed {
			if he, ok := errors.AsType[*echo.HTTPError](err); ok {
				status = he.Code
			} else {
				status = http.StatusInternalServerError
			}
		}

		resBody := parseJSON(c.Response().Header().Get(echo.HeaderContentType), rec.body.Bytes())

		fields := append(baseFields(),
			zap.Int("status", status),
			zap.Duration("latency", latency),
			zap.Any("req_body", reqBody),
			zap.Any("res_body", resBody),
		)

		isGet := req.Method == http.MethodGet
		hasError := err != nil || status >= 400

		if isGet && !hasError {
			return nil
		}

		if err != nil {
			if se, ok := errors.AsType[*response.SafeError](err); ok {
				lg.Error("http error", append(fields,
					zap.String("error_code", string(se.Code)),
					zap.String("caller", se.Caller),
					zap.Error(se.Internal),
				)...)
			} else {
				lg.Error("http error", append(fields, zap.Error(err))...)
			}
			return err
		}

		switch {
		case status >= 500:
			lg.Error("http", fields...)
		case status >= 400:
			lg.Warn("http", fields...)
		default:
			lg.Info("http", fields...)
		}

		return nil
	}
}

type responseBodyRecorder struct {
	http.ResponseWriter
	body  bytes.Buffer
	limit int64
}

func (r *responseBodyRecorder) Write(b []byte) (int, error) {
	if int64(r.body.Len()) < r.limit {
		remain := r.limit - int64(r.body.Len())
		if int64(len(b)) > remain {
			r.body.Write(b[:remain])
		} else {
			r.body.Write(b)
		}
	}
	return r.ResponseWriter.Write(b)
}

func (r *responseBodyRecorder) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	h, ok := r.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, http.ErrNotSupported
	}
	return h.Hijack()
}

func parseJSON(contentType string, raw []byte) any {
	if len(raw) == 0 {
		return nil
	}
	if !strings.Contains(strings.ToLower(contentType), "application/json") {
		return truncateStr(string(raw), 2000)
	}
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return truncateStr(string(raw), 2000)
	}
	return v
}

func truncateStr(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "...(truncated)"
}

func summarizeMultipart(req *http.Request) any {
	ct := req.Header.Get("Content-Type")
	_, params, err := mime.ParseMediaType(ct)
	if err != nil {
		return "multipart (unparseable)"
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "multipart (read error)"
	}
	req.Body = io.NopCloser(bytes.NewReader(body))

	reader := multipart.NewReader(bytes.NewReader(body), params["boundary"])
	var parts []map[string]string
	for {
		part, err := reader.NextPart()
		if err != nil {
			break
		}
		info := map[string]string{
			"name": part.FormName(),
		}
		if fn := part.FileName(); fn != "" {
			info["filename"] = fn
			info["content_type"] = part.Header.Get("Content-Type")
			size, _ := io.Copy(io.Discard, part)
			info["size"] = fmt.Sprintf("%d", size)
		}
		_ = part.Close()
		parts = append(parts, info)
	}

	return parts
}
