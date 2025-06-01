package container

import (
	"bytes"
	"context"

	"github.com/mark3labs/mcp-go/mcp"

	mcpServer "github.com/mark3labs/mcp-go/server"
)

func NewReadFile() *ReadFile {
	return &ReadFile{}
}

type ReadFile struct {
}

type ReadFileParams struct {
	Container string `json:"container"`
	Filename  string `json:"filename"`
}

func (h *ReadFile) McpTool() (mcp.Tool, mcpServer.ToolHandlerFunc) {
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

func (h *ReadFile) Handle(_ context.Context, _ mcp.CallToolRequest, params ReadFileParams) (*mcp.CallToolResult, error) {
	buf := bytes.NewBuffer(nil)
	if err := sshReadFile(params.Container, params.Filename, buf); err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(buf.String()), nil
}
