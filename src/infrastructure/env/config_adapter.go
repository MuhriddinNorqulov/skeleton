package env

import (
	"example.com/PROJECT_NAME/src/core/domain/ports/config"
	"time"
)

type ConfigAdapter struct {
	env *Env
}

// @inject
func NewConfigAdapter(env *Env) config.ConfigProvider {
	return &ConfigAdapter{env: env}
}

func (this *ConfigAdapter) GetStorageBaseURL() string { return this.env.StorageURL }
func (this *ConfigAdapter) GetAccessTokenExpireMinutes() int64 {
	return this.env.AccessTimeExpireMinutes
}
func (this *ConfigAdapter) GetRefreshTokenExpireMinutes() int64 {
	return this.env.RefreshTokenExpireMinutes
}
func (this *ConfigAdapter) GetHmacSecret() string                  { return this.env.HmacSecret }
func (this *ConfigAdapter) GetWsReconnectBaseDelay() time.Duration { return 1 * time.Second }
func (this *ConfigAdapter) GetWsReconnectMaxDelay() time.Duration  { return 30 * time.Second }
func (this *ConfigAdapter) GetWsMessageTimeout() time.Duration     { return 10 * time.Second }
func (this *ConfigAdapter) GetWsIdleTimeout() time.Duration        { return 600 * time.Second }
