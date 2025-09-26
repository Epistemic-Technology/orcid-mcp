package tools

import (
	"context"
	"os"

	"github.com/Epistemic-Technology/orcid/orcid"
	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type GetOrcidRecordQuery struct {
	OrcidID string `json:"orcid_id"`
}

func GetOrcidRecordTool() *mcp.Tool {
	inputSchema, err := jsonschema.For[GetOrcidRecordQuery](nil)
	if err != nil {
		panic(err)
	}
	getOrcidRecordTool := mcp.Tool{
		Name:        "Get ORCID Record",
		Description: "Retrieves the ORCID record for a given ORCID ID.",
		InputSchema: inputSchema,
	}
	return &getOrcidRecordTool
}

func GetOrcidRecordHandler(ctx context.Context, req *mcp.CallToolRequest, query GetOrcidRecordQuery) (*mcp.CallToolResult, *orcid.Record, error) {
	bearerToken := os.Getenv("ORCID_BEARER_TOKEN")
	client := orcid.NewClient(orcid.WithBearerToken(bearerToken))
	record, err := client.GetRecord(ctx, query.OrcidID)
	if err != nil {
		return nil, nil, err
	}
	return &mcp.CallToolResult{}, record, nil
}
