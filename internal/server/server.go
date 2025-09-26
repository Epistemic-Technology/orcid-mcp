package server

import (
	"github.com/Epistemic-Technology/orcid-mcp/internal/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func CreateServer() *mcp.Server {
	server := mcp.NewServer(&mcp.Implementation{Name: "orcid-mcp", Version: "v0.0.1"}, nil)
	mcp.AddTool(server, tools.GetOrcidRecordTool(), tools.GetOrcidRecordHandler)
	return server
}
