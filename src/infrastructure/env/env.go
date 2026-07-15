package env

import (
	"os"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Env struct {
	HttpAddress string `env:"HTTP_ADDRESS,required"`
	DBHost      string `env:"DB_HOST,required"`
	DBPort      string `env:"DB_PORT,required"`
	DBUser      string `env:"DB_USER,required"`
	DBPassword  string `env:"DB_PASSWORD,required"`
	DBName      string `env:"DB_NAME,required"`
	DBSSLMode   string `env:"DB_SSLMODE,required"`

	RedisAddress  string `env:"REDIS_ADDRESS,required"`
	RedisPassword string `env:"REDIS_PASSWORD,required"`

	CorsAllowCredentials bool     `env:"CORS_ALLOW_CREDENTIALS,required"`
	CorsAllowOrigin      []string `env:"CORS_ALLOW_ORIGIN,required"`

	JwtSecret  string `env:"JWT_SECRET,required"`
	HmacSecret string `env:"HMAC_SECRET,required"`

	AccessTimeExpireMinutes   int64 `env:"ACCESS_EXPIRE_MINUTES,required"`
	RefreshTokenExpireMinutes int64 `env:"REFRESH_EXPIRE_MINUTES,required"`

	StoragePath string `env:"STORAGE_PATH" envDefault:"storage"`
	StorageURL  string `env:"STORAGE_URL,required"`
}

// @inject
func NewEnv() *Env {
	return parseEnv()
}

func parseEnv() *Env {
	loadEnv()
	cfg := &Env{}
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}

func loadEnv() {
	if os.Getenv("CONTAINER_MODE") != "1" {
		_ = godotenv.Load("env/.env.local")
		_ = godotenv.Load("env/.env")
	}
}
