// Package config loads and holds all application settings.
package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// Config holds all settings for the application, grouped by area.
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	AWS      AWSConfig
	Upload   UploadConfig
}

// ServerConfig holds the HTTP server settings.
type ServerConfig struct {
	Port    string // Port the server listens on, e.g. "8080"
	GinMode string // Gin framework mode: "debug" or "release"
}

// DatabaseConfig holds the PostgreSQL connection settings.
type DatabaseConfig struct {
	Host     string // e.g. "localhost"
	Port     string // e.g. "5432"
	User     string // e.g. "postgres"
	Password string
	Name     string // Database name, e.g. "ecommerce_shop"
	SSLMode  string // "disable" for local development
}

// JWTConfig holds the settings for login tokens.
type JWTConfig struct {
	Secret          string        // Secret key used to sign tokens
	ExpiresIn       time.Duration // How long an access token is valid, e.g. 15 * time.Minute
	RefreshTokenExp time.Duration // How long a refresh token is valid, e.g. 7 * 24 * time.Hour
}

// AWSConfig holds the settings for connecting to AWS (or LocalStack in development).
type AWSConfig struct {
	Region          string // AWS region, e.g. "us-east-1"
	AccessKeyID     string // AWS access key (use "test" for LocalStack)
	SecretAccessKey string // AWS secret key (use "test" for LocalStack)
	S3Bucket        string // Name of the S3 bucket where files are stored
	S3Endpoint      string // Custom S3 URL, e.g. "http://localhost:4566" for LocalStack; empty for real AWS
}

// UploadConfig holds the settings for file uploads.
type UploadConfig struct {
	Path        string // Folder where uploaded files are saved
	MaxFileSize int64  // Maximum allowed file size in bytes (e.g. 10 << 20 = 10 MB)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

// Load reads settings from the .env file / environment variables and returns a Config.
func Load() (*Config, error) {
	_ = godotenv.Load()

	jwtExpiresIn, _ := time.ParseDuration(getEnv("JWT_EXPIRES_IN", "24h"))
	refreshTokenExpires, _ := time.ParseDuration(getEnv("REFRESH_TOKEN_EXPIRES_IN", "720h"))
	maxUploadSize, _ := strconv.ParseInt(getEnv("MAX_UPLOAD_SIZE", "10485760"), 10, 64)

	return &Config{
		Server: ServerConfig{
			Port:    getEnv("PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "password"),
			Name:     getEnv("DB_NAME", "ecommerce"),
			SSLMode:  getEnv("DB_SSL_MODE", "disable"),
		},
		JWT: JWTConfig{
			Secret:          getEnv("JWT_SECRET", "your-super-secret-jwt-key"),
			ExpiresIn:       jwtExpiresIn,
			RefreshTokenExp: refreshTokenExpires,
		},
		AWS: AWSConfig{
			Region:          getEnv("AWS_REGION", "us-east-1"),
			AccessKeyID:     getEnv("AWS_ACCESS_KEY_ID", "test"),
			SecretAccessKey: getEnv("AWS_SECRET_ACCESS_KEY", "test"),
			S3Bucket:        getEnv("AWS_S3_BUCKET", "ecommerce-uploads"),
			S3Endpoint:      getEnv("AWS_S3_ENDPOINT", "http://localhost:4566"),
		},
		Upload: UploadConfig{
			Path:        getEnv("UPLOAD_PATH", "./uploads"),
			MaxFileSize: maxUploadSize,
		},
	}, nil
}
