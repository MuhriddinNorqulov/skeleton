package config

import "time"

type ConfigProvider interface {
	GetAccessTokenExpireMinutes() int64
	GetRefreshTokenExpireMinutes() int64
	GetHmacSecret() string
	GetStorageBaseURL() string
	GetWsReconnectBaseDelay() time.Duration
	GetWsReconnectMaxDelay() time.Duration
	GetWsMessageTimeout() time.Duration
	GetWsIdleTimeout() time.Duration
}
