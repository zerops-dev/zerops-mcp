package serviceDelete

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpServer "github.com/mark3labs/mcp-go/server"
	"github.com/zerops-dev/zerops-mcp/services/zeropsSdk"
	"github.com/zeropsio/zerops-go/dto/input/path"
	"github.com/zeropsio/zerops-go/types/uuid"
)

func New(sdk *zeropsSdk.Handler) *Handler {
	return &Handler{
		sdk: sdk,
	}
}

type Handler struct {
	sdk *zeropsSdk.Handler
}

type Params struct {
	Token          string              `json:"token"`
	ServiceStackId uuid.ServiceStackId `json:"serviceStackId"`
}

func (h *Handler) McpTool() (mcp.Tool, mcpServer.ToolHandlerFunc) {
	return mcp.NewTool(
		"service_delete",
		mcp.WithDescription("Delete service"),
		mcp.WithString("token",
			mcp.Description("sdk access token"),
			mcp.Required(),
		),
		mcp.WithString("serviceStackId",
			mcp.Description("service id"),
			mcp.Required(),
		),
	), mcp.NewTypedToolHandler(h.Handle)
}

func (h *Handler) Handle(ctx context.Context, _ mcp.CallToolRequest, params Params) (*mcp.CallToolResult, error) {
	_, err := h.sdk.Authorized(params.Token).DeleteServiceStack(ctx, path.ServiceStackId{
		Id: params.ServiceStackId,
	})
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText("success"), nil
}
