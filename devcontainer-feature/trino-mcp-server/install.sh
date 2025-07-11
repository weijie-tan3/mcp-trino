#!/usr/bin/env bash
set -e

# The 'install.sh' entrypoint script is always executed as the root user

# Variable declarations from feature options
VERSION="${VERSION:-"latest"}"
TRINO_HOST="${TRINOHOST}"
TRINO_PORT="${TRINOPORT:-"8080"}"
TRINO_USER="${TRINOUSER:-"trino"}"
TRINO_CATALOG="${TRINOCATALOG}"
TRINO_SCHEMA="${TRINOSCHEMA:-"default"}"
TRINO_MCP_REPO="${TRINOMCPREPO}"

echo "Setting up Trino MCP Server..."

# Setup the server directory
mkdir -p /opt/trino-mcp-server
cd /opt/trino-mcp-server

if [ "${VERSION}" = "latest" ]; then
    # Clone the repository if not in the workspace
    echo "Cloning the latest version of Trino MCP Server from ${TRINO_MCP_REPO}..."
    git clone ${TRINO_MCP_REPO} .
else
    # Clone a specific version
    echo "Cloning version ${VERSION} of Trino MCP Server from ${TRINO_MCP_REPO}..."
    git clone --depth 1 --branch ${VERSION} ${TRINO_MCP_REPO} .
fi

# Make sure Go is available and install dependencies
echo "Installing Go dependencies..."
go mod download

INSTALL_WITH_SUDO="false"
if command -v sudo >/dev/null 2>&1; then
    if [ "root" != "$_REMOTE_USER" ]; then
        INSTALL_WITH_SUDO="true"
    fi
fi

ENV_PATH=/home/$_REMOTE_USER/.trino-mcp-env

if [ "${INSTALL_WITH_SUDO}" = "true" ]; then
    sudo -u ${_REMOTE_USER} bash -c "echo 'TRINO_HOST=$TRINO_HOST' >> $ENV_PATH"
    sudo -u ${_REMOTE_USER} bash -c "echo 'TRINO_PORT=$TRINO_PORT' >> $ENV_PATH"
    sudo -u ${_REMOTE_USER} bash -c "echo 'TRINO_USER=$TRINO_USER' >> $ENV_PATH"
    sudo -u ${_REMOTE_USER} bash -c "echo 'TRINO_CATALOG=$TRINO_CATALOG' >> $ENV_PATH"
    sudo -u ${_REMOTE_USER} bash -c "echo 'TRINO_SCHEMA=$TRINO_SCHEMA' >> $ENV_PATH"
    sudo -u ${_REMOTE_USER} bash -c "echo 'TRINO_SCHEME=http' >> $ENV_PATH"
    sudo -u ${_REMOTE_USER} bash -c "echo 'MCP_TRANSPORT=stdio' >> $ENV_PATH"
else
    echo "TRINO_HOST=$TRINO_HOST" >> $ENV_PATH || true
    echo "TRINO_PORT=$TRINO_PORT" >> $ENV_PATH || true
    echo "TRINO_USER=$TRINO_USER" >> $ENV_PATH || true
    echo "TRINO_CATALOG=$TRINO_CATALOG" >> $ENV_PATH || true
    echo "TRINO_SCHEMA=$TRINO_SCHEMA" >> $ENV_PATH || true
    echo "TRINO_SCHEME=http" >> $ENV_PATH || true
    echo "MCP_TRANSPORT=stdio" >> $ENV_PATH || true
fi

echo "Trino MCP Server installation complete!"
echo "Environment variables have been written to $ENV_PATH"
echo "The server will be built automatically via postCreateCommand."