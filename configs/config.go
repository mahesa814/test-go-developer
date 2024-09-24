package configs

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	AppPort                  = getEnv("APP_PORT", "8080")
	AppMode                  = getEnv("APP_MODE", "development")
	DatabaseHost             = getEnv("DATABASE_HOST", "localhost")
	DatabasePort             = getEnv("DATABASE_PORT", "5432")
	DatabaseUser             = getEnv("DATABASE_USERNAME", "postgres")
	DatabaseName             = getEnv("DATABASE_NAME", "postgres")
	DatabasePassword         = getEnv("DATABASE_PASSWORD", "postgres")
	DatabaseSSL              = getEnv("DATABASE_SSL", "disable")
	DatabaseLog              = getEnv("DATABASE_LOG", "disable")
	DatabaseSecretKey        = getEnv("DATABASE_SECRET_KEY", "your-secret-key")
	JwtSecret                = getEnv("JWT_SECRET", "your-secret-key")
	JwtTTL                   = getEnv("JWT_TTL", "3600")
	JwtRefreshTTL            = getEnv("JWT_REFRESH_TTL", "3600")
)

func getEnv(key, fallback string) string {
	LoadEnv()
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}
