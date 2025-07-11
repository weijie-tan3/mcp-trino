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
make test            # Run unit tests with race detection
make run-dev         # Run from source code (go run ./cmd)
make run             # Run built binary
make clean           # Clean build artifacts
make lint            # Run linting (same as CI: golangci-lint + go mod tidy)

# Docker development
make docker-compose-up   # Start with Docker Compose
make docker-compose-down # Stop Docker Compose
make run-docker          # Build and run Docker image locally

# Release and packaging
make release-snapshot    # Create snapshot release with GoReleaser
make build-dxt          # Build platform-specific binaries for DXT
make pack-dxt           # Package DXT extension

# Testing individual components
go test ./internal/config    # Test configuration package
go test ./internal/trino     # Test Trino client package
go test ./internal/handlers  # Test MCP handlers package
```

## Architecture

### Core Components

1. **Main Entry Point** (`cmd/main.go:26-134`): 
   - Server initialization and Trino connection testing
   - Transport selection (STDIO vs HTTP with SSE)
   - Graceful shutdown with signal handling
   - CORS support for web clients

2. **Configuration Layer** (`internal/config/config.go:11-74`): 
   - Environment-based configuration with validation
   - Security defaults (HTTPS, read-only queries)
   - Timeout configuration with validation

3. **Client Layer** (`internal/trino/client.go:16-233`): 
   - Database connection management with connection pooling
   - SQL injection protection via read-only query enforcement
   - Context-based timeout handling for queries

4. **Handler Layer** (`internal/handlers/trino_handlers.go:13-161`): 
   - MCP tool implementations with JSON response formatting
   - Parameter validation and error handling
   - Consistent logging for debugging

### Transport Support

- **STDIO Transport**: Direct MCP client integration (default)
- **HTTP Transport**: REST API with SSE support on `/sse` endpoint
- **Message Endpoint**: POST `/api/v1` for HTTP-based MCP communication
- **Status Endpoint**: GET `/` returns server status and version

### SQL Security Architecture

The security model centers around `internal/trino/client.go:73-90` with `isReadOnlyQuery()` function:
- Allows: SELECT, SHOW, DESCRIBE, EXPLAIN, WITH (CTEs)
- Blocks: INSERT, UPDATE, DELETE, CREATE, DROP, ALTER by default
- Override: Set `TRINO_ALLOW_WRITE_QUERIES=true` to bypass (logs warning)

### Available MCP Tools

All tools return JSON-formatted responses and handle parameter validation:
- `execute_query`: Execute SQL queries with security restrictions
- `list_catalogs`: Discover available data catalogs
- `list_schemas`: List schemas within catalogs (optional catalog param)
- `list_tables`: List tables within schemas (optional catalog/schema params)
- `get_table_schema`: Retrieve table structure (required table param)

## Configuration

Environment variables for connection and security:
- `TRINO_HOST`, `TRINO_PORT`, `TRINO_USER`, `TRINO_PASSWORD`
- `TRINO_SCHEME` (http/https), `TRINO_SSL`, `TRINO_SSL_INSECURE`
- `TRINO_ALLOW_WRITE_QUERIES` (default: false for security)
- `TRINO_QUERY_TIMEOUT` (default: 30 seconds, validated > 0)
- `MCP_TRANSPORT` (stdio/http), `MCP_PORT` (default: 9097), `MCP_HOST`

Key defaults and behaviors:
- HTTPS scheme forces SSL=true regardless of TRINO_SSL setting
- Invalid timeout values fall back to 30 seconds with warning
- Connection pool: 10 max open, 5 max idle, 5min max lifetime

## CI/CD Pipeline

The GitHub Actions workflow (`.github/workflows/build.yml`) includes:

**Code Quality** (`verify` job):
- Go mod tidy verification
- golangci-lint with 5m timeout
- Dependency verification

**Security** (`security` job):
- govulncheck for Go vulnerability scanning
- Trivy SARIF scanning (CRITICAL/HIGH/MEDIUM severity)
- SBOM generation with SPDX format
- Security results uploaded to GitHub Security tab

**Testing** (`test` job):
- Race detection enabled (`go test -race`)
- Code coverage with atomic mode
- Coverage uploaded to Codecov

**Build/Release**:
- Multi-platform builds via GoReleaser
- Docker images published to GHCR
- Automated releases on main branch pushes
- SLSA provenance generation for supply chain security

## Manual Testing

- Use `examples/test_query.go` for HTTP API testing
- Docker Compose setup includes real Trino server
- Set `MCP_TRANSPORT=http` and test SSE endpoint at `http://localhost:9097/sse`

## Build and Release

- **Multi-platform Support**: Uses GoReleaser for linux/darwin/windows on amd64/arm64/arm
- **Version Injection**: `-ldflags "-X main.Version=$(VERSION)"` sets version from git tags
- **Docker**: Multi-stage build with scratch base image for minimal size
- **Distribution**: GitHub Releases, GHCR, Homebrew tap (`tuannvm/mcp`)