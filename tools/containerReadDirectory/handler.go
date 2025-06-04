package containerReadDirectory

import (
	"context"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	mcpServer "github.com/mark3labs/mcp-go/server"
	"github.com/zerops-dev/zerops-mcp/services/sshUtil"
)

func New() *Handler {
	return &Handler{}
}

type Handler struct {
}

type Params struct {
	Container string `json:"container"`
	Directory string `json:"directory"`
}

func (h *Handler) McpTool() (mcp.Tool, mcpServer.ToolHandlerFunc) {
	return mcp.NewTool(
		"container_file_read",
		mcp.WithDescription("Read a single file"),
		mcp.WithString("container",
			mcp.Description("source container name"),
			mcp.Required(),
		),
		mcp.WithString("filename",
			mcp.Description("source filename"),
			mcp.Required(),
		),
	), mcp.NewTypedToolHandler(h.Handle)
}

func (h *Handler) Handle(_ context.Context, _ mcp.CallToolRequest, params Params) (*mcp.CallToolResult, error) {
	results, err := sshUtil.ReaddDirectory(params.Container, params.Directory)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(strings.Join(results, "\n")), nil
}
