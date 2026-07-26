package wire

import (
	"github.com/muhriddinnorqulov/skeleton/src/entrypoint/asynctask"
	"github.com/muhriddinnorqulov/skeleton/src/entrypoint/http"

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
