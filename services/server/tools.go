package server

import "github.com/mark3labs/mcp-go/server"

type Tool interface {
	Register(server *server.MCPServer)
}

func (h *Handler) RegisterTools(in ...Tool) {
	for _, tool := range in {
		tool.Register(h.mcpServer)
	}
}
