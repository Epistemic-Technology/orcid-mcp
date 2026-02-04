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

**For sandbox testing**, use `https://sandbox.orcid.org/oauth/token` instead.

The response will contain your bearer token:
```json
{
  "access_token": "your-bearer-token-here",
  "token_type": "bearer",
  "expires_in": 631138518,
  "scope": "/read-public"
}
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

### Claude Desktop Integration

Add the server to your Claude Desktop configuration:

1. Edit the Claude Desktop config file:
   - **Linux**: `~/.config/Claude/claude_desktop_config.json`
   - **macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
   - **Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

2. Add the ORCID MCP server to the `mcpServers` section:

```json
{
  "mcpServers": {
    "orcid-mcp": {
      "type": "stdio",
      "command": "/absolute/path/to/orcid-mcp/bin/orcid-mcp-local-server",
      "args": [],
      "env": {
        "ORCID_BEARER_TOKEN": "your-bearer-token-here"
      }
    }
  }
}
```

3. Replace `/absolute/path/to/orcid-mcp/bin/orcid-mcp-local-server` with the actual path to your compiled binary
4. Replace `your-bearer-token-here` with your actual bearer token from the API setup
5. Restart the Claude Desktop application

### Claude Code Integration

```bash
make cc-add-mcp        # Add to Claude Code (project scope)
make cc-remove-mcp     # Remove from Claude Code
```

For user-wide Claude Code access:

```bash
claude mcp add orcid-mcp-local-server --scope user -- /path/to/orcid-mcp/bin/orcid-mcp-local-server
```

Note: You may need to manually add the `ORCID_BEARER_TOKEN` environment variable to the generated configuration file.

## Available Tools

- **Get ORCID Record**: Retrieves ORCID records by ID

## Testing

```bash
make test
```

## Notes

- `/read-public` scope provides read-only access to public ORCID data
- Tokens are long-lived (approximately 20 years for read-public scope)
- No user authorization required for client credentials flow
- Production endpoint: `https://orcid.org/oauth/token`
- Sandbox endpoint: `https://sandbox.orcid.org/oauth/token`

## Security

- Never commit your bearer token or client credentials to version control
- Add `.env` files and credentials to `.gitignore`
- Consider regenerating API credentials periodically
- Use environment variables or secure secret management for production deployments
