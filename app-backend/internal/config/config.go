package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName       string
	Env           string
	HTTPPort      string
	GRPCPort      string
	POSTGRES_URL  string
	MIGRATE_DB    bool
	RedisURL      string
	AppBackendURL string
}

var (
	cfg  *Config
	once sync.Once
)

// Load returns the singleton config instance
func Load() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, using environment variables")
		}

		cfg = &Config{
			AppName:       getEnv("APP_NAME"),
			Env:           getEnv("ENV"),
			HTTPPort:      getEnv("HTTP_PORT"),
			GRPCPort:      getEnv("GRPC_PORT"),
			POSTGRES_URL:  getEnv("POSTGRES_URL"),
			MIGRATE_DB:    getEnvBool("MIGRATE_DB", false),
			RedisURL:      getEnv("REDIS_URL"),
			AppBackendURL: getEnv("APP_BACKEND_URL"),
		}
	})

	return cfg
}
