package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

// TrinoConfig holds Trino connection parameters
type TrinoConfig struct {
	Host              string
	Port              int
	User              string
	Password          string
	Catalog           string
	Schema            string
	Scheme            string
	SSL               bool
	SSLInsecure       bool
	AllowWriteQueries bool // Controls whether non-read-only SQL queries are allowed
}

// NewTrinoConfig creates a new TrinoConfig with values from environment variables or defaults
func NewTrinoConfig() *TrinoConfig {
	port, _ := strconv.Atoi(getEnv("TRINO_PORT", "8080"))
	ssl, _ := strconv.ParseBool(getEnv("TRINO_SSL", "true"))
	sslInsecure, _ := strconv.ParseBool(getEnv("TRINO_SSL_INSECURE", "true"))
	scheme := getEnv("TRINO_SCHEME", "https")
	allowWriteQueries, _ := strconv.ParseBool(getEnv("TRINO_ALLOW_WRITE_QUERIES", "false"))

	// If using HTTPS, force SSL to true
	if strings.EqualFold(scheme, "https") {
		ssl = true
	}

	// Log a warning if write queries are allowed
	if allowWriteQueries {
		log.Println("WARNING: Write queries are enabled (TRINO_ALLOW_WRITE_QUERIES=true). SQL injection protection is bypassed.")
	}

	return &TrinoConfig{
		Host:              getEnv("TRINO_HOST", "localhost"),
		Port:              port,
		User:              getEnv("TRINO_USER", "trino"),
		Password:          getEnv("TRINO_PASSWORD", ""),
		Catalog:           getEnv("TRINO_CATALOG", "memory"),
		Schema:            getEnv("TRINO_SCHEMA", "default"),
		Scheme:            scheme,
		SSL:               ssl,
		SSLInsecure:       sslInsecure,
		AllowWriteQueries: allowWriteQueries,
	}
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
