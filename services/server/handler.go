package server

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zerops-dev/di/appRunner"
)

type Tools interface {
	McpTool() (mcp.Tool, server.ToolHandlerFunc)
}

func RegisterTools(r *Handler, p Tools) {
	r.mcpServer.AddTool(p.McpTool())
}

func New(
	config Config,
	register appRunner.Register,
) (h *Handler, _ error) {
	defer func() { register.Add(h) }()
	h = &Handler{
		config: config,
	}
	return h, h.initServer()
}

type Handler struct {
	config    Config
	mcpServer *server.MCPServer
}

func (h *Handler) initServer() error {
	hooks := &server.Hooks{}
	hooks.AddBeforeAny(func(ctx context.Context, id any, method mcp.MCPMethod, message any) {
		fmt.Printf("beforeAny: %s, %v, %v\n", method, id, message)
	})
	hooks.AddOnSuccess(func(ctx context.Context, id any, method mcp.MCPMethod, message any, result any) {
		fmt.Printf("onSuccess: %s, %v, %v, %v\n", method, id, message, result)
	})
	hooks.AddOnError(func(ctx context.Context, id any, method mcp.MCPMethod, message any, err error) {
		fmt.Printf("onError: %s, %v, %v, %v\n", method, id, message, err)
	})
	hooks.AddBeforeInitialize(func(ctx context.Context, id any, message *mcp.InitializeRequest) {
		fmt.Printf("beforeInitialize: %v, %v\n", id, message)
	})
	hooks.AddOnRequestInitialization(func(ctx context.Context, id any, message any) error {
		fmt.Printf("AddOnRequestInitialization: %v, %v\n", id, message)
		// authorization verification and other preprocessing tasks are performed.
		return nil
	})
	hooks.AddAfterInitialize(func(ctx context.Context, id any, message *mcp.InitializeRequest, result *mcp.InitializeResult) {
		fmt.Printf("afterInitialize: %v, %v, %v\n", id, message, result)
	})
	hooks.AddAfterCallTool(func(ctx context.Context, id any, message *mcp.CallToolRequest, result *mcp.CallToolResult) {
		fmt.Printf("afterCallTool: %v, %v, %v\n", id, message, result)
	})
	hooks.AddBeforeCallTool(func(ctx context.Context, id any, message *mcp.CallToolRequest) {
		fmt.Printf("beforeCallTool: %v, %v\n", id, message)
	})

	h.mcpServer = server.NewMCPServer(
		"zerops",
		"1.0.0.",
		server.WithHooks(hooks),
	)

	return nil
}

func (h *Handler) Run(_ context.Context) error {

	if err := server.ServeStdio(h.mcpServer); err != nil {
		return err
	}
	return nil
}
