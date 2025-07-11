# Trino MCP Server Dev Container Feature

This dev container feature installs and configures a Trino Model Context Protocol (MCP) server for development purposes.

## Description

This feature sets up everything needed to run and develop with a Trino MCP server:

- Installs the MCP server and its dependencies
- Sets up Go development environment  
- Configures the necessary environment for development with Trino databases

## Usage

```json
"features": {
    "ghcr.io/tuannvm/devcontainer-features/trino-mcp-server:1.0.0": {
        "version": "latest",
        "trinoHost": "trino.example.com",
        "trinoPort": "8080", 
        "trinoUser": "trino",
        "trinoCatalog": "memory",
        "trinoSchema": "default",
        "trinoMcpRepo": "https://github.com/tuannvm/mcp-trino"
    }
}
```

## Dependencies

This feature automatically installs the following dependencies:
- `ghcr.io/devcontainers/features/go` - For Go development environment (version 1.24)

You don't need to explicitly include this in your devcontainer.json file.

## Options

| Option         | Default                                    | Description                                                       |
|----------------|--------------------------------------------|-------------------------------------------------------------------|
| version        | "latest"                                   | Version of the Trino MCP server to install                       |
| trinoHost      | ""                                         | Trino host address (must be specified at runtime)                |
| trinoPort      | "8080"                                     | Trino port number                                                 |
| trinoUser      | "trino"                                    | Trino user name                                                   |
| trinoCatalog   | ""                                         | Trino catalog name (must be specified at runtime)                |
| trinoSchema    | "default"                                  | Trino schema name                                                 |
| trinoMcpRepo   | "https://github.com/tuannvm/mcp-trino"    | Trino MCP repository URL                                          |

## Environment Variables

The feature creates a `.trino-mcp-env` file in the user's home directory with the following variables:

- `TRINO_HOST` - Trino server host
- `TRINO_PORT` - Trino server port  
- `TRINO_USER` - Trino username
- `TRINO_CATALOG` - Trino catalog to use
- `TRINO_SCHEMA` - Trino schema to use
- `TRINO_SCHEME` - Connection scheme (defaults to http)
- `MCP_TRANSPORT` - MCP transport type (defaults to stdio)

## Development Workflow

1. The feature automatically clones the Trino MCP repository to `/opt/trino-mcp-server`
2. Go dependencies are downloaded during installation
3. The binary is built via the `postCreateCommand` 
4. VS Code is configured to use the MCP server with the appropriate environment file

## VS Code Integration

The feature automatically configures VS Code with MCP server settings that:
- Use the built binary at `/opt/trino-mcp-server/trino-mcp`
- Load environment variables from `~/.trino-mcp-env`
- Use STDIO transport for communication

## Example devcontainer.json

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

## Troubleshooting

### Connection Issues
- Ensure `trinoHost` and `trinoCatalog` are correctly configured
- Check that the Trino server is accessible from the dev container
- Verify environment variables in `~/.trino-mcp-env`

### Build Issues  
- The Go toolchain should be automatically installed
- If build fails, check that dependencies downloaded correctly
- Try rebuilding manually: `cd /opt/trino-mcp-server && go build -o trino-mcp ./cmd`

## License

See the [LICENSE](../../LICENSE) file for details.

## More Information

For more information about the Trino MCP server, see the [main repository](https://github.com/tuannvm/mcp-trino).