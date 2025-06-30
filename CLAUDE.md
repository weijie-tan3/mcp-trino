# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**mcp-trino** is a Model Context Protocol (MCP) server that enables AI assistants to interact with Trino's distributed SQL query engine. It acts as a bridge between AI applications and Trino databases, allowing conversational access to data analytics through standardized MCP tools.

## Tech Stack

- **Language:** Go 1.24.2
- **Key Dependencies:** 
  - `github.com/mark3labs/mcp-go` v0.25.0 (MCP protocol)
  - `github.com/trinodb/trino-go-client` v0.323.0 (Trino client)
- **Build Tools:** GoReleaser, Docker, GitHub Actions, golangci-lint

## Development Commands

```bash
# Core development
make build           # Build binary to ./bin/mcp-trino
make test            # Run unit tests
make run-dev         # Run from source code
make run             # Run built binary
make clean           # Clean build artifacts
make lint            # Run linting (same as CI)

# Docker development
make docker-compose-up   # Start with Docker Compose
make docker-compose-down # Stop Docker Compose

# Release
make release-snapshot    # Create snapshot release
```

## Architecture

### Core Components

1. **Main Entry Point** (`cmd/main.go`): Server initialization, transport selection, graceful shutdown
2. **Configuration Layer** (`internal/config/config.go`): Environment-based configuration with validation
3. **Client Layer** (`internal/trino/client.go`): Database connection management with security features
4. **Handler Layer** (`internal/handlers/trino_handlers.go`): MCP tool implementations for Trino operations

### Transport Support

- **STDIO Transport**: For direct MCP client integration
- **HTTP Transport**: REST API with Server-Sent Events (SSE) support
- **CORS Support**: For web-based MCP clients

### Available MCP Tools

- `execute_query`: Execute SQL queries with security restrictions
- `list_catalogs`: Discover available data catalogs
- `list_schemas`: List schemas within catalogs
- `list_tables`: List tables within schemas
- `get_table_schema`: Retrieve table structure and column information

## Configuration

Environment variables for connection and security:
- `TRINO_HOST`, `TRINO_PORT`, `TRINO_USER`, `TRINO_PASSWORD`
- `TRINO_SCHEME` (http/https), `TRINO_SSL`, `TRINO_SSL_INSECURE`
- `TRINO_ALLOW_WRITE_QUERIES` (default: false for security)
- `TRINO_QUERY_TIMEOUT` (default: 30 seconds)
- `MCP_TRANSPORT` (stdio/http), `MCP_PORT`, `MCP_HOST`

## Security Features

- **SQL Injection Protection**: Read-only query enforcement by default
- **Query Timeout**: Configurable timeout to prevent resource exhaustion
- **SSL/TLS Support**: Secure connections to Trino
- **Environment-based Configuration**: No hardcoded credentials

## Testing

- **Code Quality**: Uses golangci-lint for static analysis
- **Security**: govulncheck for vulnerability scanning, Trivy for dependencies
- **Manual Testing**: `examples/test_query.go` for HTTP API testing
- **Integration**: Docker Compose setup with real Trino server

## Build and Release

- **Multi-platform Support**: linux/amd64, linux/arm64, linux/arm/v7, linux/arm/v6
- **Docker**: Multi-stage build with optimized images
- **CI/CD**: Automated builds, security scanning, and releases via GitHub Actions
- **Distribution**: GitHub Releases, GitHub Container Registry, Homebrew tap