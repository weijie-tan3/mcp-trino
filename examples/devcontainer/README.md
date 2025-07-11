# Devcontainer Example

This directory contains an example VS Code devcontainer configuration that uses the Trino MCP Server devcontainer feature.

## Usage

1. Copy the `devcontainer.json` file to your project's `.devcontainer/` directory
2. Modify the Trino connection settings in the feature configuration:
   - `trinoHost`: Your Trino server hostname
   - `trinoCatalog`: The catalog you want to use
   - Other settings as needed
3. Open your project in VS Code
4. When prompted, click "Reopen in Container" or use Command Palette: "Dev Containers: Reopen in Container"

## Configuration

The example configuration:
- Uses the official Go 1.24 dev container as the base image
- Installs the Trino MCP Server feature with sample settings
- Forwards ports 8080 (Trino) and 9097 (MCP Server) 
- Includes the Go extension for VS Code

## Testing with Local Trino

To test with a local Trino instance, you can use the Docker Compose setup in the repository root:

1. Start Trino: `docker-compose up -d`
2. Use `trinoHost: "localhost"` in your devcontainer configuration
3. The memory catalog should be available for testing

## Customization

You can customize the feature by modifying the options in the `features` section:

```json
"ghcr.io/tuannvm/devcontainer-features/trino-mcp-server:1.0.0": {
    "trinoHost": "your-trino-host.com",
    "trinoPort": "8080",
    "trinoCatalog": "your-catalog",
    "trinoSchema": "your-schema"
}
```