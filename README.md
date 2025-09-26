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

## Configuration

Set your ORCID API bearer token as an environment variable:

```bash
export ORCID_BEARER_TOKEN="your-token-here"
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