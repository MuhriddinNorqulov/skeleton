package db

import (
	"fmt"
	"github.com/muhriddinnorqulov/skeleton/src/infrastructure/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// @inject
func NewGormDB(env *env.Env) *gorm.DB {
	db, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
				env.DBHost, env.DBUser, env.DBPassword, env.DBName, env.DBPort, env.DBSSLMode,
			),
		),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	if err != nil {
		panic(fmt.Errorf("[NewGormDB] failed to connect database: %w", err))
	}
	return db
}
