package config

import (
	"os"
	"strconv"
	"strings"
)

// TrinoConfig holds Trino connection parameters
type TrinoConfig struct {
	Host        string
	Port        int
	User        string
	Password    string
	Catalog     string
	Schema      string
	Scheme      string
	SSL         bool
	SSLInsecure bool
}

// NewTrinoConfig creates a new TrinoConfig with values from environment variables or defaults
func NewTrinoConfig() *TrinoConfig {
	port, _ := strconv.Atoi(getEnv("TRINO_PORT", "8080"))
	ssl, _ := strconv.ParseBool(getEnv("TRINO_SSL", "true"))
	sslInsecure, _ := strconv.ParseBool(getEnv("TRINO_SSL_INSECURE", "true"))
	scheme := getEnv("TRINO_SCHEME", "https")
	
	// If using HTTPS, force SSL to true
	if strings.ToLower(scheme) == "https" {
		ssl = true
	}

	return &TrinoConfig{
		Host:        getEnv("TRINO_HOST", "localhost"),
		Port:        port,
		User:        getEnv("TRINO_USER", "trino"),
		Password:    getEnv("TRINO_PASSWORD", ""),
		Catalog:     getEnv("TRINO_CATALOG", "memory"),
		Schema:      getEnv("TRINO_SCHEMA", "default"),
		Scheme:      scheme,
		SSL:         ssl,
		SSLInsecure: sslInsecure,
	}
}

// getEnv retrieves an environment variable or returns a default value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
