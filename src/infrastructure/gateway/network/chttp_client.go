package network

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"example.com/PROJECT_NAME/src/infrastructure/logger"
)

type CHTTpClient struct {
	*http.Client
	log *logger.GatewayLogger
}

// @inject
func NewCHTTpClient(client *http.Client, log *logger.GatewayLogger) *CHTTpClient {
	return &CHTTpClient{Client: client, log: log}
}

func (this *CHTTpClient) Get(ctx context.Context, url string, auth *BasicAuth, header map[string]string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	if auth != nil {
		req.SetBasicAuth(auth.Username, auth.Password)
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	this.log.LogRequest("GET", url, "", nil)
	return this.do(req)
}

func (this *CHTTpClient) Post(ctx context.Context, url string, body any, auth *BasicAuth, header map[string]string) (*http.Response, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	if auth != nil {
		req.SetBasicAuth(auth.Username, auth.Password)
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	cType := req.Header.Get("Content-Type")
	if cType == "" {
		cType = "application/json"
		req.Header.Set("Content-Type", cType)
	}
	this.log.LogRequest("POST", url, cType, payload)
	return this.do(req)
}

func (this *CHTTpClient) PostMultipart(ctx context.Context, url string, form MultipartForm, header map[string]string) (*http.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for _, f := range form.Files {
		part, err := writer.CreateFormFile(f.Field, f.Name)
		if err != nil {
			return nil, err
		}
		if _, err = io.Copy(part, f.Reader); err != nil {
			return nil, err
		}
	}

	for k, v := range form.Fields {
		if err := writer.WriteField(k, v); err != nil {
			return nil, err
		}
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return nil, err
	}

	contentType := writer.FormDataContentType()
	req.Header.Set("Content-Type", contentType)
	for k, v := range header {
		req.Header.Set(k, v)
	}

	this.log.LogRequest("POST", url, contentType, nil)
	return this.do(req)
}

func (this *CHTTpClient) do(req *http.Request) (*http.Response, error) {
	start := time.Now()
	resp, err := this.Do(req)
	duration := time.Since(start)

	if err != nil {
		this.log.LogError(req.Method, req.URL.String(), err, duration)
		return nil, err
	}

	body, _ := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	resp.Body = io.NopCloser(bytes.NewReader(body))

	this.log.LogResponse(req.Method, req.URL.String(), resp.StatusCode, resp.Header.Get("Content-Type"), body, duration)
	return resp, nil
}
