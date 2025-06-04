package containerReadFile

import (
	"bytes"
	"context"

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
	Filename  string `json:"filename"`
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
	buf := bytes.NewBuffer(nil)
	if err := sshUtil.ReadFile(params.Container, params.Filename, buf); err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(buf.String()), nil
}
