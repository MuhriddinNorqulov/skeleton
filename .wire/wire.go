package wire

import (
	"example.com/PROJECT_NAME/src/entrypoint/asynctask"
	"example.com/PROJECT_NAME/src/entrypoint/http"

	"github.com/google/wire"
)

func InitHttpApp() *http.App {
	wire.Build(ProviderSet)
	return nil
}

func InitAsyncApp() *asynctask.App {
	wire.Build(ProviderSet)
	return nil
}
