package network

import (
	"encoding/json"
	"example.com/PROJECT_NAME/src/core/utils"
	"net/http"
)

func GetBody[T any](resp *http.Response) (*T, error) {
	var body T
	defer utils.Close(resp.Body)
	err := json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func GetErrorBody(resp *http.Response) map[string]any {
	var body map[string]any
	defer utils.Close(resp.Body)
	_ = json.NewDecoder(resp.Body).Decode(&body)
	return body
}
