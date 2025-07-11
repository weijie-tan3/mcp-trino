package config

import (
	"os"
	"testing"
	"time"
)

func TestNewTrinoConfig_OAuth(t *testing.T) {
	// Save original env vars to restore later
	originalEnvs := map[string]string{
		"TRINO_EXTERNAL_AUTHENTICATION": os.Getenv("TRINO_EXTERNAL_AUTHENTICATION"),
		"TRINO_ACCESS_TOKEN":            os.Getenv("TRINO_ACCESS_TOKEN"),
		"TRINO_USER":                    os.Getenv("TRINO_USER"),
		"TRINO_PASSWORD":                os.Getenv("TRINO_PASSWORD"),
	}
	defer func() {
		for key, value := range originalEnvs {
			if value == "" {
				os.Unsetenv(key)
			} else {
				os.Setenv(key, value)
			}
		}
	}()

	tests := []struct {
		name                       string
		externalAuth               string
		accessToken                string
		expectedExternalAuth       bool
		expectedAccessToken        string
		expectedUser               string
		expectedPassword           string
	}{
		{
			name:                 "OAuth disabled, traditional auth",
			externalAuth:         "false",
			accessToken:          "",
			expectedExternalAuth: false,
			expectedAccessToken:  "",
			expectedUser:         "trino",
			expectedPassword:     "",
		},
		{
			name:                 "OAuth enabled with access token",
			externalAuth:         "true",
			accessToken:          "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test.token",
			expectedExternalAuth: true,
			expectedAccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test.token",
			expectedUser:         "trino",
			expectedPassword:     "",
		},
		{
			name:                 "OAuth enabled without access token",
			externalAuth:         "true",
			accessToken:          "",
			expectedExternalAuth: true,
			expectedAccessToken:  "",
			expectedUser:         "trino",
			expectedPassword:     "",
		},
		{
			name:                 "OAuth disabled with access token present",
			externalAuth:         "false",
			accessToken:          "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test.token",
			expectedExternalAuth: false,
			expectedAccessToken:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test.token",
			expectedUser:         "trino",
			expectedPassword:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables for this test case
			os.Setenv("TRINO_EXTERNAL_AUTHENTICATION", tt.externalAuth)
			os.Setenv("TRINO_ACCESS_TOKEN", tt.accessToken)

			config := NewTrinoConfig()

			if config.ExternalAuthentication != tt.expectedExternalAuth {
				t.Errorf("ExternalAuthentication = %v, want %v", config.ExternalAuthentication, tt.expectedExternalAuth)
			}

			if config.AccessToken != tt.expectedAccessToken {
				t.Errorf("AccessToken = %q, want %q", config.AccessToken, tt.expectedAccessToken)
			}

			if config.User != tt.expectedUser {
				t.Errorf("User = %q, want %q", config.User, tt.expectedUser)
			}

			if config.Password != tt.expectedPassword {
				t.Errorf("Password = %q, want %q", config.Password, tt.expectedPassword)
			}

			// Verify default values are still set correctly
			if config.Host != "localhost" {
				t.Errorf("Host = %q, want %q", config.Host, "localhost")
			}

			if config.Port != 8080 {
				t.Errorf("Port = %d, want %d", config.Port, 8080)
			}

			if config.QueryTimeout != 30*time.Second {
				t.Errorf("QueryTimeout = %v, want %v", config.QueryTimeout, 30*time.Second)
			}
		})
	}
}

func TestNewTrinoConfig_Defaults(t *testing.T) {
	// Clear all environment variables to test defaults
	envVars := []string{
		"TRINO_EXTERNAL_AUTHENTICATION",
		"TRINO_ACCESS_TOKEN",
		"TRINO_HOST",
		"TRINO_PORT",
		"TRINO_USER",
		"TRINO_PASSWORD",
		"TRINO_CATALOG",
		"TRINO_SCHEMA",
		"TRINO_SCHEME",
		"TRINO_SSL",
		"TRINO_SSL_INSECURE",
		"TRINO_ALLOW_WRITE_QUERIES",
		"TRINO_QUERY_TIMEOUT",
	}

	// Save original values
	originalValues := make(map[string]string)
	for _, env := range envVars {
		originalValues[env] = os.Getenv(env)
		os.Unsetenv(env)
	}

	// Restore original values after test
	defer func() {
		for env, value := range originalValues {
			if value == "" {
				os.Unsetenv(env)
			} else {
				os.Setenv(env, value)
			}
		}
	}()

	config := NewTrinoConfig()

	// Test OAuth defaults
	if config.ExternalAuthentication != false {
		t.Errorf("ExternalAuthentication = %v, want %v", config.ExternalAuthentication, false)
	}

	if config.AccessToken != "" {
		t.Errorf("AccessToken = %q, want empty string", config.AccessToken)
	}

	// Test other defaults
	if config.Host != "localhost" {
		t.Errorf("Host = %q, want %q", config.Host, "localhost")
	}

	if config.User != "trino" {
		t.Errorf("User = %q, want %q", config.User, "trino")
	}

	if config.Scheme != "https" {
		t.Errorf("Scheme = %q, want %q", config.Scheme, "https")
	}
}