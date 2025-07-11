# Trino MCP Server in Go

A high-performance Model Context Protocol (MCP) server for Trino implemented in Go. This project enables AI assistants to seamlessly interact with Trino's distributed SQL query engine through standardized MCP tools.

[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/tuannvm/mcp-trino/build.yml?branch=main&label=CI%2FCD&logo=github)](https://github.com/tuannvm/mcp-trino/actions/workflows/build.yml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/tuannvm/mcp-trino?logo=go)](https://github.com/tuannvm/mcp-trino/blob/main/go.mod)
[![Trivy Scan](https://img.shields.io/github/actions/workflow/status/tuannvm/mcp-trino/build.yml?branch=main&label=Trivy%20Security%20Scan&logo=aquasec)](https://github.com/tuannvm/mcp-trino/actions/workflows/build.yml)
[![SLSA 3](https://slsa.dev/images/gh-badge-level3.svg)](https://slsa.dev)
[![Go Report Card](https://goreportcard.com/badge/github.com/tuannvm/mcp-trino)](https://goreportcard.com/report/github.com/tuannvm/mcp-trino)
[![Go Reference](https://pkg.go.dev/badge/github.com/tuannvm/mcp-trino.svg)](https://pkg.go.dev/github.com/tuannvm/mcp-trino)
[![Docker Image](https://img.shields.io/github/v/release/tuannvm/mcp-trino?sort=semver&label=GHCR&logo=docker)](https://github.com/tuannvm/mcp-trino/pkgs/container/mcp-trino)
[![GitHub Release](https://img.shields.io/github/v/release/tuannvm/mcp-trino?sort=semver)](https://github.com/tuannvm/mcp-trino/releases/latest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)


## Overview

This project implements a Model Context Protocol (MCP) server for Trino in Go. It enables AI assistants to access Trino's distributed SQL query engine through standardized MCP tools.

Trino (formerly PrestoSQL) is a powerful distributed SQL query engine designed for fast analytics on large datasets.

## Features

- ✅ MCP server implementation in Go
- ✅ Trino SQL query execution through MCP tools
- ✅ Catalog, schema, and table discovery
- ✅ Docker container support
- ✅ Supports both STDIO and HTTP transports
- ✅ Server-Sent Events (SSE) support for Cursor and other MCP clients
- ✅ Compatible with Cursor, Claude Desktop, Windsurf, ChatWise, and any MCP-compatible clients.

## Installation

### Quick Install (One-liner)

For macOS and Linux, install with a single command:

```bash
curl -fsSL https://raw.githubusercontent.com/tuannvm/mcp-trino/main/install.sh -o install.sh && chmod +x install.sh && ./install.sh
```

### Homebrew (macOS and Linux)

The easiest way to install mcp-trino is using Homebrew:

```bash
# Add the tap repository
brew tap tuannvm/mcp

# Install mcp-trino
brew install mcp-trino
```

To update to the latest version:

```bash
brew update && brew upgrade mcp-trino
```

### Alternative Installation Methods

#### Manual Download

1. Download the appropriate binary for your platform from the [GitHub Releases](https://github.com/tuannvm/mcp-trino/releases) page.
2. Place the binary in a directory included in your PATH (e.g., `/usr/local/bin` on Linux/macOS)
3. Make it executable (`chmod +x mcp-trino` on Linux/macOS)

#### From Source

```bash
git clone https://github.com/tuannvm/mcp-trino.git
cd mcp-trino
make build
# Binary will be in ./bin/
```

## Downloads

You can download pre-built binaries for your platform:

| Platform | Architecture | Download Link |
|----------|--------------|---------------|
| macOS | x86_64 (Intel) | [Download](https://github.com/tuannvm/mcp-trino/releases/latest/download/mcp-trino-darwin-amd64) |
| macOS | ARM64 (Apple Silicon) | [Download](https://github.com/tuannvm/mcp-trino/releases/latest/download/mcp-trino-darwin-arm64) |
| Linux | x86_64 | [Download](https://github.com/tuannvm/mcp-trino/releases/latest/download/mcp-trino-linux-amd64) |
| Linux | ARM64 | [Download](https://github.com/tuannvm/mcp-trino/releases/latest/download/mcp-trino-linux-arm64) |
| Windows | x86_64 | [Download](https://github.com/tuannvm/mcp-trino/releases/latest/download/mcp-trino-windows-amd64.exe) |

Or see all available downloads on the [GitHub Releases](https://github.com/tuannvm/mcp-trino/releases) page.

## VS Code Dev Container

For easy testing and development, you can use the VS Code dev container feature:

### Using the Dev Container Feature

Add the following to your `devcontainer.json`:

```json
{
    "name": "Trino MCP Development",
    "image": "mcr.microsoft.com/devcontainers/go:1.24",
    "features": {
        "ghcr.io/tuannvm/devcontainer-features/trino-mcp-server:1.0.0": {
            "trinoHost": "your-trino-host",
            "trinoCatalog": "your-catalog"
        }
    }
}
```

This will automatically:
- Set up the Go development environment
- Clone and build the Trino MCP server
- Configure VS Code with MCP settings
- Create environment variables for Trino connection

See `examples/devcontainer/` for a complete example configuration.

### Installation Troubleshooting

If you encounter issues during installation:

**Common Issues:**
- **Binary not found in PATH**: The install script installs to `~/.local/bin` by default. Make sure this directory is in your PATH:
  ```bash
  export PATH="$HOME/.local/bin:$PATH"
  ```
  Add this to your shell profile (`.bashrc`, `.zshrc`, etc.) to make it permanent.

- **Permission denied**: If you get permission errors, ensure the install directory is writable:
  ```bash
  mkdir -p ~/.local/bin
  chmod 755 ~/.local/bin
  ```

- **Claude configuration not found**: If the install script doesn't detect your Claude installation:
  - For Claude Desktop: Check if the config file exists at the expected location
  - For Claude Code: Verify the `claude` command is available in PATH
  - Use the manual configuration instructions provided by the script

- **GitHub API rate limiting**: If you're hitting GitHub API rate limits:
  ```bash
  export GITHUB_TOKEN=your_github_token
  curl -fsSL https://raw.githubusercontent.com/tuannvm/mcp-trino/main/install.sh | bash
  ```

**Getting Help:**
- Check the [GitHub Issues](https://github.com/tuannvm/mcp-trino/issues) for similar problems
- Run the install script with `--help` for usage information
- Use manual installation methods if the automated script fails
## MCP Integration

This MCP server can be integrated with several AI applications:

### Using Docker Image

To use the Docker image instead of a local binary:

#### Traditional Authentication
```json
{
  "mcpServers": {
    "mcp-trino": {
      "command": "docker",
      "args": ["run", "--rm", "-i",
               "-e", "TRINO_HOST=<HOST>",
               "-e", "TRINO_PORT=<PORT>",
               "-e", "TRINO_USER=<USERNAME>",
               "-e", "TRINO_PASSWORD=<PASSWORD>",
               "-e", "TRINO_SCHEME=http",
               "ghcr.io/tuannvm/mcp-trino:latest"],
      "env": {}
    }
  }
}
```

#### OAuth Authentication
```json
{
  "mcpServers": {
    "mcp-trino": {
      "command": "docker",
      "args": ["run", "--rm", "-i",
               "-e", "TRINO_HOST=<HOST>",
               "-e", "TRINO_PORT=<PORT>",
               "-e", "TRINO_EXTERNAL_AUTHENTICATION=true",
               "-e", "TRINO_ACCESS_TOKEN=<YOUR_JWT_TOKEN>",
               "-e", "TRINO_SCHEME=https",
               "ghcr.io/tuannvm/mcp-trino:latest"],
      "env": {}
    }
  }
}
```

> **Note**: The `host.docker.internal` special DNS name allows the container to connect to services running on the host machine. If your Trino server is running elsewhere, replace with the appropriate host.

This Docker configuration can be used in any of the below applications.

### Cursor

To use with [Cursor](https://cursor.sh/), create or edit `~/.cursor/mcp.json`:

#### Traditional Authentication
```json
{
  "mcpServers": {
    "mcp-trino": {
      "command": "mcp-trino",
      "args": [],
      "env": {
        "TRINO_HOST": "<HOST>",
        "TRINO_PORT": "<PORT>",
        "TRINO_USER": "<USERNAME>",
        "TRINO_PASSWORD": "<PASSWORD>"
      }
    }
  }
}
```

#### OAuth Authentication
```json
{
  "mcpServers": {
    "mcp-trino": {
      "command": "mcp-trino",
      "args": [],
      "env": {
        "TRINO_HOST": "<HOST>",
        "TRINO_PORT": "<PORT>",
        "TRINO_EXTERNAL_AUTHENTICATION": "true",
        "TRINO_ACCESS_TOKEN": "<YOUR_JWT_TOKEN>"
      }
    }
  }
}
```

Replace the environment variables with your specific Trino configuration.

For HTTP+SSE transport mode (supported for Cursor integration):

```json
{
  "mcpServers": {
    "mcp-trino-http": {
      "url": "http://localhost:9097/sse"
    }
  }
}
```

Then start the server in a separate terminal with:

#### Traditional Authentication
```bash
MCP_TRANSPORT=http TRINO_HOST=<HOST> TRINO_PORT=<PORT> TRINO_USER=<USERNAME> TRINO_PASSWORD=<PASSWORD> mcp-trino
```

#### OAuth Authentication
```bash
MCP_TRANSPORT=http TRINO_HOST=<HOST> TRINO_PORT=<PORT> TRINO_EXTERNAL_AUTHENTICATION=true TRINO_ACCESS_TOKEN=<YOUR_JWT_TOKEN> mcp-trino
```

### Claude Desktop

To use with [Claude Desktop](https://claude.ai/desktop), the easiest way is to use the install script which will automatically configure it for you. Alternatively, you can manually edit your Claude configuration file:

- macOS: `~/Library/Application Support/Claude/claude_desktop_config.json`
- Windows: `%APPDATA%\Claude\claude_desktop_config.json`
- Linux: `~/.config/Claude/claude_desktop_config.json`

#### Traditional Authentication
```json
{
  "mcpServers": {
    "mcp-trino": {
      "command": "mcp-trino",
      "args": [],
      "env": {
        "TRINO_HOST": "<HOST>",
        "TRINO_PORT": "<PORT>",
        "TRINO_USER": "<USERNAME>",
        "TRINO_PASSWORD": "<PASSWORD>"
      }
    }
  }
}
```

#### OAuth Authentication
```json
{
  "mcpServers": {
    "mcp-trino": {
      "command": "mcp-trino",
      "args": [],
      "env": {
        "TRINO_HOST": "<HOST>",
        "TRINO_PORT": "<PORT>",
        "TRINO_EXTERNAL_AUTHENTICATION": "true",
        "TRINO_ACCESS_TOKEN": "<YOUR_JWT_TOKEN>"
      }
    }
  }
}
```

After updating the configuration, restart Claude Desktop. You should see the MCP tools available in the tools menu.

### Claude Code

To use with [Claude Code](https://claude.ai/code), the install script will automatically configure it for you. Alternatively, you can manually add the MCP server:

```bash
claude mcp add mcp-trino mcp-trino
```

Then set your environment variables:

```bash
export TRINO_HOST=<HOST>
export TRINO_PORT=<PORT>
export TRINO_USER=<USERNAME>
export TRINO_PASSWORD=<PASSWORD>
```

Restart Claude Code to see the MCP tools available.

### Windsurf

To use with [Windsurf](https://windsurf.com/refer?referral_code=sjqdvqozgx2wyi7r), create or edit your `mcp_config.json`:

```json
{
  "mcpServers": {
    "mcp-trino": {
      "command": "mcp-trino",
      "args": [],
      "env": {
        "TRINO_HOST": "<HOST>",
        "TRINO_PORT": "<PORT>",
        "TRINO_USER": "<USERNAME>",
        "TRINO_PASSWORD": "<PASSWORD>"
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
   - Command: `mcp-trino`
   - Args: (leave empty)
   - Env: Add the following environment variables:
     ```
     TRINO_HOST=<HOST>
     TRINO_PORT=<PORT>
     TRINO_USER=<USERNAME>
     TRINO_PASSWORD=<PASSWORD>
     ```

Alternatively, you can import the configuration from JSON:

1. Copy this JSON to your clipboard:
   ```json
   {
     "mcpServers": {
       "mcp-trino": {
         "command": "mcp-trino",
         "args": [],
         "env": {
           "TRINO_HOST": "<HOST>",
           "TRINO_PORT": "<PORT>",
           "TRINO_USER": "<USERNAME>",
           "TRINO_PASSWORD": "<PASSWORD>"
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

### execute_query

Execute a SQL query against Trino with full SQL support for complex analytical queries.

**Sample Prompt:**
> "How many customers do we have per region? Can you show them in descending order?"

**Example:**
```json
{
  "query": "SELECT region, COUNT(*) as customer_count FROM tpch.tiny.customer GROUP BY region ORDER BY customer_count DESC"
}
```

**Response:**
```json
{
  "columns": ["region", "customer_count"],
  "data": [
    ["AFRICA", 5],
    ["AMERICA", 5],
    ["ASIA", 5],
    ["EUROPE", 5],
    ["MIDDLE EAST", 5]
  ]
}
```

### list_catalogs

List all catalogs available in the Trino server, providing a comprehensive view of your data ecosystem.

**Sample Prompt:**
> "What databases do we have access to in our Trino environment?"

**Example:**
```json
{}
```

**Response:**
```json
{
  "catalogs": ["tpch", "memory", "system", "jmx"]
}
```

### list_schemas

List all schemas in a catalog, helping you navigate through the data hierarchy efficiently.

**Sample Prompt:**
> "What schemas or datasets are available in the tpch catalog?"

**Example:**
```json
{
  "catalog": "tpch"
}
```

**Response:**
```json
{
  "schemas": ["information_schema", "sf1", "sf100", "sf1000", "tiny"]
}
```

### list_tables

List all tables in a schema, giving you visibility into available datasets.

**Sample Prompt:**
> "What tables are available in the tpch tiny schema? I need to know what data we can query."

**Example:**
```json
{
  "catalog": "tpch",
  "schema": "tiny"
}
```

**Response:**
```json
{
  "tables": ["customer", "lineitem", "nation", "orders", "part", "partsupp", "region", "supplier"]
}
```

### get_table_schema

Get the schema of a table, understanding the structure of your data for better query planning.

**Sample Prompt:**
> "What columns are in the customer table? I need to know the data types and structure before writing my query."

**Example:**
```json
{
  "catalog": "tpch",
  "schema": "tiny",
  "table": "customer"
}
```

**Response:**
```json
{
  "columns": [
    {
      "name": "custkey",
      "type": "bigint",
      "nullable": false
    },
    {
      "name": "name",
      "type": "varchar",
      "nullable": false
    },
    {
      "name": "address",
      "type": "varchar",
      "nullable": false
    },
    {
      "name": "nationkey",
      "type": "bigint",
      "nullable": false
    },
    {
      "name": "phone",
      "type": "varchar",
      "nullable": false
    },
    {
      "name": "acctbal",
      "type": "double",
      "nullable": false
    },
    {
      "name": "mktsegment",
      "type": "varchar",
      "nullable": false
    },
    {
      "name": "comment",
      "type": "varchar",
      "nullable": false
    }
  ]
}
```

This information is invaluable for understanding the column names, data types, and nullability constraints before writing queries against the table.

## End-to-End Example

Here's a complete interaction example showing how an AI assistant might use these tools to answer a business question:

**User Query:** "Can you help me analyze our biggest customers? I want to know the top 5 customers with the highest account balances."

**AI Assistant's workflow:**
1. First, discover available catalogs
   ```
   > Using list_catalogs tool
   > Discovers tpch catalog
   ```

2. Then, find available schemas
   ```
   > Using list_schemas tool with catalog "tpch"
   > Discovers "tiny" schema
   ```

3. Explore available tables
   ```
   > Using list_tables tool with catalog "tpch" and schema "tiny"
   > Finds "customer" table
   ```

4. Check the customer table schema
   ```
   > Using get_table_schema tool
   > Discovers "custkey", "name", "acctbal" and other columns
   ```

5. Finally, execute the query
   ```
   > Using execute_query tool with:
   > "SELECT custkey, name, acctbal FROM tpch.tiny.customer ORDER BY acctbal DESC LIMIT 5"
   ```

6. Returns the results to the user:
   ```
   The top 5 customers with highest account balances are:
   1. Customer #65 (Customer#000000065): $9,222.78
   2. Customer #13 (Customer#000000013): $8,270.47
   3. Customer #89 (Customer#000000089): $7,990.56
   4. Customer #11 (Customer#000000011): $7,912.91
   5. Customer #82 (Customer#000000082): $7,629.41
   ```

This seamless workflow demonstrates how the MCP tools enable AI assistants to explore and query data in a conversational manner.

## Configuration

The server can be configured using the following environment variables:

### Connection Parameters

| Variable               | Description                       | Default   |
| ---------------------- | --------------------------------- | --------- |
| TRINO_HOST             | Trino server hostname             | localhost |
| TRINO_PORT             | Trino server port                 | 8080      |
| TRINO_USER             | Trino user                        | trino     |
| TRINO_PASSWORD         | Trino password                    | (empty)   |
| TRINO_CATALOG          | Default catalog                   | memory    |
| TRINO_SCHEMA           | Default schema                    | default   |
| TRINO_SCHEME           | Connection scheme (http/https)    | https     |
| TRINO_SSL              | Enable SSL                        | true      |
| TRINO_SSL_INSECURE     | Allow insecure SSL                | true      |

### Authentication Parameters

| Variable                   | Description                              | Default |
| -------------------------- | ---------------------------------------- | ------- |
| TRINO_EXTERNAL_AUTHENTICATION | Enable OAuth/external authentication  | false   |
| TRINO_ACCESS_TOKEN         | JWT access token for OAuth authentication | (empty) |

### Security and Query Parameters

| Variable               | Description                       | Default   |
| ---------------------- | --------------------------------- | --------- |
| TRINO_ALLOW_WRITE_QUERIES | Allow non-read-only SQL queries | false     |
| TRINO_QUERY_TIMEOUT    | Query timeout in seconds          | 30        |

### MCP Transport Parameters

| Variable               | Description                       | Default   |
| ---------------------- | --------------------------------- | --------- |
| MCP_TRANSPORT          | Transport method (stdio/http)     | stdio     |
| MCP_PORT               | HTTP port for http transport      | 9097      |
| MCP_HOST               | Host for HTTP callbacks           | localhost |

### Authentication Modes

The server supports two authentication modes:

#### Traditional Authentication (Username/Password)
This is the default mode using username and password authentication:

```bash
export TRINO_HOST=localhost
export TRINO_PORT=8080
export TRINO_USER=myuser
export TRINO_PASSWORD=mypassword
```

#### OAuth Authentication (External Authentication)
For OAuth/external authentication using JWT tokens:

```bash
export TRINO_HOST=localhost
export TRINO_PORT=8080
export TRINO_EXTERNAL_AUTHENTICATION=true
export TRINO_ACCESS_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

When OAuth is enabled with a valid access token, the username and password are ignored, and the access token is used for authentication via the `Authorization: Bearer` header.

> **Note**: When `TRINO_SCHEME` is set to "https", `TRINO_SSL` is automatically set to true regardless of the provided value.

> **Important**: The default connection mode is HTTPS. If you're using an HTTP-only Trino server, you must set `TRINO_SCHEME=http` in your environment variables.

> **Security Note**: By default, only read-only queries (SELECT, SHOW, DESCRIBE, EXPLAIN) are allowed to prevent SQL injection. If you need to execute write operations or other non-read queries, set `TRINO_ALLOW_WRITE_QUERIES=true`, but be aware this bypasses this security protection.

> **OAuth Note**: When `TRINO_EXTERNAL_AUTHENTICATION=true` is set but no `TRINO_ACCESS_TOKEN` is provided, the server will log a warning and fall back to traditional username/password authentication.

> **For Cursor Integration**: When using with Cursor, set `MCP_TRANSPORT=http` and connect to the `/sse` endpoint. The server will automatically handle SSE (Server-Sent Events) connections.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## CI/CD and Releases

This project uses GitHub Actions for continuous integration and GoReleaser for automated releases.

### Continuous Integration Checks

Our CI pipeline performs the following checks on all PRs and commits to the main branch:

#### Code Quality
- **Linting**: Using golangci-lint to check for common code issues and style violations
- **Go Module Verification**: Ensuring go.mod and go.sum are properly maintained
- **Formatting**: Verifying code is properly formatted with gofmt

#### Security
- **Vulnerability Scanning**: Using govulncheck to check for known vulnerabilities in dependencies
- **Dependency Scanning**: Using Trivy to scan for vulnerabilities in dependencies (CRITICAL, HIGH, and MEDIUM)
- **SBOM Generation**: Creating a Software Bill of Materials for dependency tracking
- **SLSA Provenance**: Creating verifiable build provenance for supply chain security

#### Testing
- **Unit Tests**: Running tests with race detection and code coverage reporting
- **Build Verification**: Ensuring the codebase builds successfully

#### CI/CD Security
- **Least Privilege**: Workflows run with minimum required permissions
- **Pinned Versions**: All GitHub Actions use specific versions to prevent supply chain attacks
- **Dependency Updates**: Automated dependency updates via Dependabot

### Release Process

When changes are merged to the main branch:

1. CI checks are run to validate code quality and security
2. If successful, a new release is automatically created with:
   - Semantic versioning based on commit messages
   - Binary builds for multiple platforms
   - Docker image publishing to GitHub Container Registry
   - SBOM and provenance attestation
