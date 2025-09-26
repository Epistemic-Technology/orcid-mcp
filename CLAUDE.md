# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Build
```bash
make build              # Build the orcid-mcp-local-server binary to bin/
make clean              # Remove build artifacts
```

### Testing
```bash
make test               # Run all tests with verbose output
```

### Development
```bash
make run                # Run the server in development mode
make inspect            # Run the MCP inspector on local server
```

### Claude Code Integration
```bash
make cc-add-mcp         # Add local server to Claude Code
make cc-remove-mcp      # Remove MCP servers from Claude Code
```

## Architecture

This is a Model Context Protocol (MCP) server implementation for ORCID integration, built in Go.

### Core Components

- **cmd/orcid-mcp-local-server/main.go**: Entry point that initializes and runs the MCP server using stdio transport
- **internal/server/server.go**: Creates the MCP server instance and registers available tools
- **internal/tools/get-tool.go**: Implements the "Get ORCID Record" tool that fetches ORCID records using the ORCID API

### Key Dependencies

- `github.com/modelcontextprotocol/go-sdk/mcp`: MCP SDK for Go
- `github.com/Epistemic-Technology/orcid`: ORCID API client library

### Environment Variables

- `ORCID_BEARER_TOKEN`: Required for authenticating with the ORCID API

### MCP Tool Implementation

The server currently provides one tool:
- **Get ORCID Record**: Retrieves ORCID records by ID, requires bearer token authentication