package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Env           string
	Addr          string
	BaseURL       string
	SessionSecret string
	DBDSN         string
	GoogleID      string
	GoogleSecret  string
	AppleID       string
	AppleSecret   string
	CleanupEvery  time.Duration
}

func Load() Config {
	cleanupHours, _ := strconv.Atoi(getEnv("CLEANUP_INTERVAL_HOURS", "6"))
	return Config{
		Env:           getEnv("APP_ENV", "development"),
		Addr:          getEnv("APP_ADDR", ":8080"),
		BaseURL:       getEnv("APP_BASE_URL", "http://localhost:8080"),
		SessionSecret: getEnv("SESSION_SECRET", "change-me"),
		DBDSN:         fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4,utf8", getEnv("DB_USER", "root"), getEnv("DB_PASS", "root"), getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306"), getEnv("DB_NAME", "stay_tene_life")),
		GoogleID:      os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
		GoogleSecret:  os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
		AppleID:       os.Getenv("APPLE_OAUTH_CLIENT_ID"),
		AppleSecret:   os.Getenv("APPLE_OAUTH_CLIENT_SECRET"),
		CleanupEvery:  time.Duration(cleanupHours) * time.Hour,
	}
}

func getEnv(k, fallback string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return fallback
}
