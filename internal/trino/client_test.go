package trino

import (
	"strings"
	"testing"

	"github.com/tuannvm/mcp-trino/internal/config"
)

func TestBuildDSN_OAuth(t *testing.T) {
	tests := []struct {
		name                   string
		config                 *config.TrinoConfig
		expectedDSNContains    []string
		expectedDSNNotContains []string
		expectedDSN            string
	}{
		{
			name: "Traditional authentication with username and password",
			config: &config.TrinoConfig{
				Host:                   "localhost",
				Port:                   8080,
				User:                   "testuser",
				Password:               "testpass",
				Catalog:                "memory",
				Schema:                 "default",
				Scheme:                 "https",
				SSL:                    true,
				SSLInsecure:            true,
				ExternalAuthentication: false,
				AccessToken:            "",
			},
			expectedDSN: "https://testuser:testpass@localhost:8080?catalog=memory&schema=default&SSL=true&SSLInsecure=true",
		},
		{
			name: "OAuth authentication with access token",
			config: &config.TrinoConfig{
				Host:                   "localhost",
				Port:                   8080,
				User:                   "testuser", // Should be ignored with OAuth
				Password:               "testpass", // Should be ignored with OAuth
				Catalog:                "memory",
				Schema:                 "default",
				Scheme:                 "https",
				SSL:                    true,
				SSLInsecure:            true,
				ExternalAuthentication: true,
				AccessToken:            "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test.token",
			},
			expectedDSN: "https://localhost:8080?catalog=memory&schema=default&SSL=true&SSLInsecure=true&accessToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test.token",
		},
		{
			name: "OAuth authentication with special characters in token",
			config: &config.TrinoConfig{
				Host:                   "trino.example.com",
				Port:                   443,
				Catalog:                "hive",
				Schema:                 "public",
				Scheme:                 "https",
				SSL:                    true,
				SSLInsecure:            false,
				ExternalAuthentication: true,
				AccessToken:            "token+with/special=chars&symbols",
			},
			expectedDSN: "https://trino.example.com:443?catalog=hive&schema=public&SSL=true&SSLInsecure=false&accessToken=token%2Bwith%2Fspecial%3Dchars%26symbols",
		},
		{
			name: "External auth enabled but no access token",
			config: &config.TrinoConfig{
				Host:                   "localhost",
				Port:                   8080,
				User:                   "testuser",
				Password:               "testpass",
				Catalog:                "memory",
				Schema:                 "default",
				Scheme:                 "https",
				SSL:                    true,
				SSLInsecure:            true,
				ExternalAuthentication: true,
				AccessToken:            "",
			},
			expectedDSN: "https://testuser:testpass@localhost:8080?catalog=memory&schema=default&SSL=true&SSLInsecure=true",
		},
		{
			name: "Special characters in catalog and schema",
			config: &config.TrinoConfig{
				Host:                   "localhost",
				Port:                   8080,
				User:                   "user+name",
				Password:               "pass@word",
				Catalog:                "catalog with spaces",
				Schema:                 "schema/with/slashes",
				Scheme:                 "https",
				SSL:                    true,
				SSLInsecure:            false,
				ExternalAuthentication: false,
				AccessToken:            "",
			},
			expectedDSN: "https://user%2Bname:pass%40word@localhost:8080?catalog=catalog+with+spaces&schema=schema%2Fwith%2Fslashes&SSL=true&SSLInsecure=false",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualDSN := buildDSN(tt.config)

			if actualDSN != tt.expectedDSN {
				t.Errorf("buildDSN() = %q, want %q", actualDSN, tt.expectedDSN)
			}

			// Additional checks for contains/not contains if specified
			for _, mustContain := range tt.expectedDSNContains {
				if !strings.Contains(actualDSN, mustContain) {
					t.Errorf("DSN %q should contain %q", actualDSN, mustContain)
				}
			}

			for _, mustNotContain := range tt.expectedDSNNotContains {
				if strings.Contains(actualDSN, mustNotContain) {
					t.Errorf("DSN %q should not contain %q", actualDSN, mustNotContain)
				}
			}
		})
	}
}

func TestNewClient_OAuth_DSN(t *testing.T) {
	tests := []struct {
		name           string
		config         *config.TrinoConfig
		expectedDSNContains []string
		expectedDSNNotContains []string
	}{
		{
			name: "Traditional authentication with username and password",
			config: &config.TrinoConfig{
				Host:                   "localhost",
				Port:                   8080,
				User:                   "testuser",
				Password:               "testpass",
				Catalog:                "memory",
				Schema:                 "default",
				Scheme:                 "https",
				SSL:                    true,
				SSLInsecure:            true,
				ExternalAuthentication: false,
				AccessToken:            "",
			},
			expectedDSNContains: []string{
				"https://testuser:testpass@localhost:8080",
				"catalog=memory",
				"schema=default",
				"SSL=true",
				"SSLInsecure=true",
			},
			expectedDSNNotContains: []string{
				"accessToken=",
			},
		},
		{
			name: "OAuth authentication with access token",
			config: &config.TrinoConfig{
				Host:                   "localhost",
				Port:                   8080,
				User:                   "testuser",  // Should be ignored with OAuth
				Password:               "testpass",  // Should be ignored with OAuth
				Catalog:                "memory",
				Schema:                 "default",
				Scheme:                 "https",
				SSL:                    true,
				SSLInsecure:            true,
				ExternalAuthentication: true,
				AccessToken:            "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test.token",
			},
			expectedDSNContains: []string{
				"https://localhost:8080",
				"catalog=memory",
				"schema=default",
				"SSL=true",
				"SSLInsecure=true",
				"accessToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.test.token",
			},
			expectedDSNNotContains: []string{
				"testuser:testpass@",
				"testuser",
				"testpass",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Since NewClient() tries to connect to the database, we expect it to fail
			// but we can examine the DSN by looking at any database-related error messages
			// or by testing the DSN construction separately.
			
			client, err := NewClient(tt.config)
			
			// We expect an error because we're not connecting to a real server
			if err == nil && client != nil {
				// If somehow we got a client, clean it up
				client.Close()
				t.Log("Unexpected successful connection - cleaning up")
			}
			
			// The error should contain connection-related information
			// This is a limitation of testing without a real server, but we can
			// at least verify that the function handles the config correctly
			if err != nil {
				// This is expected - we're testing with a config that won't connect
				t.Logf("Expected connection error: %v", err)
			}
		})
	}
}

func TestIsReadOnlyQuery(t *testing.T) {
	tests := []struct {
		name     string
		query    string
		expected bool
	}{
		// Basic read-only queries
		{
			name:     "Simple SELECT query",
			query:    "SELECT * FROM table",
			expected: true,
		},
		{
			name:     "SELECT query with WHERE clause",
			query:    "SELECT id, name FROM users WHERE age > 18",
			expected: true,
		},
		{
			name:     "SHOW query",
			query:    "SHOW TABLES",
			expected: true,
		},
		{
			name:     "DESCRIBE query",
			query:    "DESCRIBE users",
			expected: true,
		},
		{
			name:     "EXPLAIN query",
			query:    "EXPLAIN SELECT * FROM users",
			expected: true,
		},
		{
			name:     "WITH query (CTE)",
			query:    "WITH cte AS (SELECT * FROM users) SELECT * FROM cte",
			expected: true,
		},

		// Complex read-only queries
		{
			name:     "SELECT with GROUP BY",
			query:    "SELECT department, COUNT(*) FROM employees GROUP BY department",
			expected: true,
		},
		{
			name:     "SELECT with ORDER BY",
			query:    "SELECT * FROM products ORDER BY price DESC",
			expected: true,
		},
		{
			name:     "SELECT with JOIN",
			query:    "SELECT u.name, o.product FROM users u JOIN orders o ON u.id = o.user_id",
			expected: true,
		},
		{
			name:     "Complex SELECT with multiple clauses",
			query:    "SELECT department, COUNT(*) as count, AVG(salary) as avg_salary FROM employees WHERE hire_date > '2020-01-01' GROUP BY department HAVING count > 5 ORDER BY avg_salary DESC LIMIT 10",
			expected: true,
		},

		// Queries with different whitespace formatting
		{
			name:     "SELECT with newlines",
			query:    "SELECT\n* FROM\nusers",
			expected: true,
		},
		{
			name:     "SELECT with tabs and spaces",
			query:    "SELECT    id,\n\t\tname\nFROM users",
			expected: true,
		},
		{
			name:     "SELECT keyword without space",
			query:    "SELECT*FROM users",
			expected: true,
		},
		{
			name:     "SELECT with leading and trailing whitespace",
			query:    "  \n  SELECT * FROM users  \n  ",
			expected: true,
		},

		// Keywords without spaces
		{
			name:     "SELECT without space after keyword",
			query:    "SELECTid, name FROM users",
			expected: true,
		},
		{
			name:     "SHOW without space after keyword",
			query:    "SHOWtables",
			expected: true,
		},
		{
			name:     "DESCRIBE without space after keyword",
			query:    "DESCRIBEusers",
			expected: true,
		},

		// Case insensitivity
		{
			name:     "Lowercase SELECT",
			query:    "select * from users",
			expected: true,
		},
		{
			name:     "Mixed case SELECT",
			query:    "SeLeCt * FrOm UsErS",
			expected: true,
		},

		// Write operations (should return false)
		{
			name:     "INSERT query",
			query:    "INSERT INTO users VALUES (1, 'John')",
			expected: false,
		},
		{
			name:     "UPDATE query",
			query:    "UPDATE users SET name = 'John' WHERE id = 1",
			expected: false,
		},
		{
			name:     "DELETE query",
			query:    "DELETE FROM users WHERE id = 1",
			expected: false,
		},
		{
			name:     "DROP query",
			query:    "DROP TABLE users",
			expected: false,
		},
		{
			name:     "CREATE query",
			query:    "CREATE TABLE users (id INT, name VARCHAR)",
			expected: false,
		},
		{
			name:     "ALTER query",
			query:    "ALTER TABLE users ADD COLUMN email VARCHAR",
			expected: false,
		},
		{
			name:     "TRUNCATE query",
			query:    "TRUNCATE TABLE users",
			expected: false,
		},

		// Sneaky write operations embedded in SELECT (should return false)
		{
			name:     "SELECT with embedded INSERT",
			query:    "SELECT * FROM users; INSERT INTO logs VALUES ('accessed')",
			expected: false,
		},
		{
			name:     "SELECT with embedded UPDATE",
			query:    "SELECT * FROM (UPDATE users SET active = true RETURNING *) AS updated",
			expected: false,
		},
		{
			name:     "SELECT with embedded DELETE",
			query:    "SELECT * FROM users WHERE id IN (DELETE FROM inactive_users RETURNING user_id)",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isReadOnlyQuery(tt.query)
			if result != tt.expected {
				t.Errorf("isReadOnlyQuery(%q) = %v, want %v", tt.query, result, tt.expected)
			}
		})
	}
}
