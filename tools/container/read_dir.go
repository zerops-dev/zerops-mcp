package container

import (
	"context"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	mcpServer "github.com/mark3labs/mcp-go/server"
)

func NewReadDir() *ReadDir {
	return &ReadDir{}
}

type ReadDir struct{}

type ReadDirParams struct {
	Container string `json:"container"`
	Directory string `json:"directory"`
}

func (h *ReadDir) McpTool() (mcp.Tool, mcpServer.ToolHandlerFunc) {
	return mcp.NewTool(
		"container_list_directory",
		mcp.WithDescription("List directory"),
		mcp.WithString("container",
			mcp.Description("source container name"),
			mcp.Required(),
		),
		mcp.WithString("directory",
			mcp.Description("source directory"),
			mcp.Required(),
		),
	), mcp.NewTypedToolHandler(h.Handle)
}

func (h *ReadDir) Handle(_ context.Context, _ mcp.CallToolRequest, params ReadFileParams) (*mcp.CallToolResult, error) {
	results, err := sshReaddDirectory(params.Container, params.Filename)
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText(strings.Join(results, "\n")), nil
}
