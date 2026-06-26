package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	JWTSecret          string
	JWTIssuer          string
	JWTAccessTokenTTL  time.Duration
	JWTRefreshTokenTTL time.Duration
}

func Load() *Config {
	cfg := &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "root"),
		DBName:     getEnv("DB_NAME", "go_clean"),
		JWTSecret:  getEnv("JWT_SECRET", "secret"),
		JWTIssuer:  getEnv("JWT_ISSUER", ""),
	}

	accessTTL, _ := strconv.Atoi(getEnv("JWT_ACCESS_TOKEN_TTL", "3600"))
	refreshTTL, _ := strconv.Atoi(getEnv("JWT_REFRESH_TOKEN_TTL", "604800"))

	cfg.JWTAccessTokenTTL = time.Duration(accessTTL) * time.Second
	cfg.JWTRefreshTokenTTL = time.Duration(refreshTTL) * time.Second

	log.Println("Config loaded")
	return cfg
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
