# orcid-mcp

Model Context Protocol (MCP) server for ORCID integration. This server provides tools to interact with ORCID records through the MCP interface.

## Features

- Fetch ORCID records by ID
- Built with Go using the MCP SDK

## Prerequisites

- Go 1.20 or later
- ORCID API bearer token

## Installation

```bash
make build
```

This will build the `orcid-mcp-local-server` binary in the `bin/` directory.

## ORCID API Setup

### Register for API credentials

1. Visit [ORCID Developer Tools](https://orcid.org/developer-tools)
2. Sign in to your ORCID account and navigate to Developer Tools
3. Click "Register for the free ORCID public API"
4. Complete the registration form:
   - **Application name**: `Personal ORCID MCP Server` (or similar)
   - **Application URL**: Your institutional URL or `https://orcid.org`
   - **Application description**: `Personal MCP server for accessing ORCID records`
   - **Redirect URI**: `https://orcid.org/oauth/playground`
5. Save the form to receive your **Client ID** and **Client Secret**

### Obtain bearer token

Exchange your credentials for a bearer token using client credentials flow:
```bash
curl -i -L -H 'Accept: application/json' \
  -d 'client_id=YOUR_CLIENT_ID' \
  -d 'client_secret=YOUR_CLIENT_SECRET' \
  -d 'scope=/read-public' \
  -d 'grant_type=client_credentials' \
  'https://orcid.org/oauth/token'
```

### Configure environment

Set the bearer token as an environment variable:
```bash
export ORCID_BEARER_TOKEN="your-access-token-here"
```

For persistent configuration, add to your shell config (`~/.bashrc`, `~/.zshrc`, etc.):
```bash
# ORCID MCP Server
export ORCID_BEARER_TOKEN="your-access-token-here"
```

## Usage

### Development

```bash
make run                # Run the server
make inspect           # Test with MCP inspector
```

### Claude Code Integration

```bash
make cc-add-mcp        # Add to Claude Code
make cc-remove-mcp     # Remove from Claude Code
```

## Available Tools

- **Get ORCID Record**: Retrieves ORCID records by ID

## Testing

```bash
make test
```
