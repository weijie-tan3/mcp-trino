# Trino MCP Server in Go

Model Context Protocol (MCP) server for Trino, reimplemented in Go.

## Overview

This project implements a Model Context Protocol (MCP) server for Trino in Go. It enables AI assistants to access Trino's distributed SQL query engine through standardized MCP tools.

Trino (formerly PrestoSQL) is a powerful distributed SQL query engine designed for fast analytics on large datasets.

## Features

- ✅ MCP server implementation in Go
- ✅ Trino SQL query execution through MCP tools
- ✅ Catalog, schema, and table discovery
- ✅ Docker container support
- ✅ Supports both STDIO and HTTP transports
- ✅ Compatible with Cursor, Claude Desktop, Windsurf, ChatWise, and any MCP-compatible clients.

## Prerequisites

- Go 1.24 or later
- Docker and Docker Compose (for containerized usage)
- A running Trino server (or use the provided Docker setup)

## Installation

1. Clone the repository:

```bash
git clone https://github.com/tuannvm/mcp-trino.git
cd mcp-trino
```

2. Build the server:

```bash
make build
```

## MCP Integration

This MCP server can be integrated with several AI applications:

### Cursor

To use with [Cursor](https://cursor.sh/), create or edit `~/.cursor/mcp.json`:

```json
{
  "mcpServers": {
    "mcp-trino": {
      "command": "/path/to/mcp-trino",
      "args": [],
      "env": {
        "TRINO_HOST": "localhost",
        "TRINO_PORT": "8080",
        "TRINO_USER": "trino",
        "TRINO_PASSWORD": ""
      }
    }
  }
}
```

Replace the path and environment variables with your specific Trino configuration.

### Claude Desktop

To use with [Claude Desktop](https://claude.ai/desktop), edit your Claude configuration file:

- macOS: `~/Library/Application Support/Claude/claude_desktop_config.json`
- Windows: `%APPDATA%\Claude\claude_desktop_config.json`

```json
{
  "mcpServers": {
    "mcp-trino": {
      "command": "/path/to/mcp-trino",
      "args": [],
      "env": {
        "TRINO_HOST": "localhost",
        "TRINO_PORT": "8080",
        "TRINO_USER": "trino",
        "TRINO_PASSWORD": ""
      }
    }
  }
}
```

After updating the configuration, restart Claude Desktop. You should see the MCP tools available in the tools menu.

### Windsurf

To use with [Windsurf](https://windsurf.com/refer?referral_code=sjqdvqozgx2wyi7r), create or edit your `mcp_config.json`:

```json
{
  "mcpServers": {
    "mcp-trino": {
      "command": "/path/to/mcp-trino",
      "args": [],
      "env": {
        "TRINO_HOST": "localhost",
        "TRINO_PORT": "8080",
        "TRINO_USER": "trino",
        "TRINO_PASSWORD": ""
      }
    }
  }
}
```

Restart Windsurf to apply the changes. The Trino MCP tools will be available to the Cascade AI.

### ChatWise

To use with [ChatWise](https://chatwise.app?atp=uo1wzc), follow these steps:

1. Open ChatWise and go to Settings
2. Navigate to the Tools section
3. Click the "+" icon to add a new tool
4. Select "Command Line MCP"
5. Configure with the following details:
   - ID: `mcp-trino` (or any name you prefer)
   - Command: `/path/to/mcp-trino` (full path to the mcp-trino binary)
   - Args: (leave empty)
   - Env: Add the following environment variables:
     ```
     TRINO_HOST=localhost
     TRINO_PORT=8080
     TRINO_USER=trino
     TRINO_PASSWORD=
     ```

Alternatively, you can import the configuration from JSON:

1. Copy this JSON to your clipboard:
   ```json
   {
     "mcpServers": {
       "mcp-trino": {
         "command": "/path/to/mcp-trino",
         "args": [],
         "env": {
           "TRINO_HOST": "localhost",
           "TRINO_PORT": "8080",
           "TRINO_USER": "trino",
           "TRINO_PASSWORD": ""
         }
       }
     }
   }
   ```
2. In ChatWise Settings > Tools, click the "+" icon
3. Select "Import JSON from Clipboard"
4. Toggle the switch next to the tool to enable it

Once enabled, click the hammer icon below the input box in ChatWise to access Trino MCP tools.

## Available MCP Tools

The server provides the following MCP tools:

1. `execute_query` - Execute a SQL query against Trino
2. `list_catalogs` - List all catalogs available in the Trino server
3. `list_schemas` - List all schemas in a catalog
4. `list_tables` - List all tables in a schema
5. `get_table_schema` - Get the schema of a table

## Configuration

The server can be configured using the following environment variables:

| Variable           | Description                   | Default   |
| ------------------ | ----------------------------- | --------- |
| TRINO_HOST         | Trino server hostname         | localhost |
| TRINO_PORT         | Trino server port             | 8080      |
| TRINO_USER         | Trino user                    | trino     |
| TRINO_PASSWORD     | Trino password                | (empty)   |
| TRINO_CATALOG      | Default catalog               | memory    |
| TRINO_SCHEMA       | Default schema                | default   |
| TRINO_SCHEME       | Connection scheme (http/https)| https     |
| TRINO_SSL          | Enable SSL                    | true      |
| TRINO_SSL_INSECURE | Allow insecure SSL            | true      |
| MCP_TRANSPORT      | Transport method (stdio/http) | stdio     |
| MCP_PORT           | HTTP port for http transport  | 9097      |

> **Note**: When `TRINO_SCHEME` is set to "https", `TRINO_SSL` is automatically set to true regardless of the provided value.

> **Important**: The default connection mode is HTTPS. If you're using an HTTP-only Trino server, you must set `TRINO_SCHEME=http` in your environment variables.

## Standalone Usage

Run the server with STDIO transport (for direct LLM integration):

```bash
./bin/mcp-trino
```

Or with HTTP transport:

```bash
MCP_TRANSPORT=http MCP_PORT=9097 ./bin/mcp-trino
```

## Docker Usage

Start the server with Docker Compose:

```bash
docker-compose up -d
```

Verify the API is working:

```bash
curl -X POST "http://localhost:9097/api/query" \
     -H "Content-Type: application/json" \
     -d '{"query": "SELECT 1 AS test"}'
```

## Development

```bash
# Install dependencies
go mod download

# Run tests
go test ./...

# Run linters
make lint
```

## License

MIT

## CI/CD and Releases

This project uses GitHub Actions for continuous integration and GoReleaser for automated releases.

### Makefile Commands

```bash
# Build the application
make build

# Run tests
make test

# Run linters (same as CI)
make lint

# Clean build artifacts
make clean

# Run in development mode
make run-dev

# Test GoReleaser locally
make release-snapshot

# Run with Docker
make docker-compose-up
make docker-compose-down
```
