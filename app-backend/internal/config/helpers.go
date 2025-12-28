package config

import (
	"log"
	"os"
	"strconv"
)

func getEnv(key string, fallback ...string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	if len(fallback) > 0 {
		return fallback[0]
	}
	return ""
}

func getEnvBool(key string, fallback bool) bool {
	if v := os.Getenv(key); v != "" {
		if parsed, err := strconv.ParseBool(v); err == nil {
			return parsed
		}
		log.Printf("Warning: Invalid boolean value for %s: %s, using default: %v", key, v, fallback)
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil {
			return parsed
		}
		log.Printf("Warning: Invalid integer value for %s: %s, using default: %d", key, v, fallback)
	}
	return fallback
}
