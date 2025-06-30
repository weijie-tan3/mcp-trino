#!/usr/bin/env bash
set -euo pipefail

# Check for required commands
for cmd in curl tar; do
  if ! command -v "$cmd" >/dev/null 2>&1; then
    echo "Error: $cmd is not installed. Please install $cmd to proceed."
    exit 1
  fi
done

REPO="tuannvm/mcp-trino"
BINARY="mcp-trino"

# Detect OS
sysOS="$(uname | tr '[:upper:]' '[:lower:]')"
case "$sysOS" in
  linux)   OS="linux" ;;
  darwin)  OS="darwin" ;;
  *)
    echo "Unsupported OS: $sysOS"
    echo "Please download manually from: https://github.com/$REPO/releases/latest"
    exit 1
    ;;
esac

# Detect ARCH
ARCH="$(uname -m)"
case "$ARCH" in
  x86_64|amd64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *)
    echo "Unsupported architecture: $ARCH"
    echo "Please download manually from: https://github.com/$REPO/releases/latest"
    exit 1
    ;;
esac

# Get latest version tag from GitHub API, Use GITHUB_TOKEN if available to avoid potential rate limit
if [ -n "${GITHUB_TOKEN:-}" ]; then
  VERSION=$(curl -fsSL -H "Authorization: Bearer $GITHUB_TOKEN" "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
else
  VERSION=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
fi

if [ -z "$VERSION" ]; then
  echo "Error: Unable to get latest version from GitHub API"
  exit 1
fi

# Remove 'v' prefix from version for archive name
VERSION_NO_V="${VERSION#v}"

echo "Installing $BINARY $VERSION for $OS/$ARCH..."

# Create server directory if it doesn't exist
mkdir -p server

# Construct download URL based on GoReleaser naming convention
ARCHIVE_NAME="${BINARY}_${VERSION_NO_V}_${OS}_${ARCH}.tar.gz"
DOWNLOAD_URL="https://github.com/$REPO/releases/download/$VERSION/$ARCHIVE_NAME"

# Download and extract archive
echo "Downloading from: $DOWNLOAD_URL"
if ! curl -fsSL "$DOWNLOAD_URL" | tar -xz -C server --strip-components=0 "$BINARY"; then
  echo "Error: Failed to download and extract binary"
  exit 1
fi

# Make binary executable
chmod +x "server/$BINARY"

echo "✅ Successfully installed $BINARY $VERSION to server/$BINARY"
echo "You can now use the mcp-trino extension in Claude Desktop"
