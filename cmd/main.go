package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jansaidl/zerops-mcp/tools/container"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {

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
	mcpServer := server.NewMCPServer(
		"zerops",
		"1.0.0.",
		server.WithHooks(hooks),
	)

	{
		c := &container.ReadFile{}
		mcpServer.AddTool(c.Tools(), mcp.NewTypedToolHandler(c.Handle))
	}

	{
		c := &container.WriteFile{}
		mcpServer.AddTool(c.Tools(), mcp.NewTypedToolHandler(c.Handle))
	}
	{
		c := &container.ReadDir{}
		mcpServer.AddTool(c.Tools(), mcp.NewTypedToolHandler(c.Handle))
	}

	if err := server.ServeStdio(mcpServer); err != nil {
		log.Fatal(err)
	}

}
