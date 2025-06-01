package container

import (
	"context"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	mcpServer "github.com/mark3labs/mcp-go/server"
)

func NewWriteFile() *WriteFile {
	return &WriteFile{}
}

type WriteFile struct {
}

type WriteFileParams struct {
	Container string `json:"container"`
	Filename  string `json:"filename"`
	Content   string `json:"content"`
}

func (h *WriteFile) McpTool() (mcp.Tool, mcpServer.ToolHandlerFunc) {
	return mcp.NewTool(
		"container_file_write",
		mcp.WithDescription("Write a single file or directory"),
		mcp.WithString("container",
			mcp.Description("target container name"),
			mcp.Required(),
		),
		mcp.WithString("filename",
			mcp.Description("target filename"),
			mcp.Required(),
		),
		mcp.WithString("content",
			mcp.Description("target filename content"),
			mcp.Required(),
		),
	), mcp.NewTypedToolHandler(h.Handle)
}

func (h *WriteFile) Handle(_ context.Context, _ mcp.CallToolRequest, params WriteFileParams) (*mcp.CallToolResult, error) {
	if err := sshWriteFile(params.Container, strings.NewReader(params.Content), params.Filename); err != nil {
		return nil, err
	}
	return mcp.NewToolResultText("success"), nil
}
