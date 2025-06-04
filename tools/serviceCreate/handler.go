package serviceCreate

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	mcpServer "github.com/mark3labs/mcp-go/server"
	"github.com/zerops-dev/zerops-mcp/services/zeropsSdk"
	"github.com/zeropsio/zerops-go/dto/input/body"
	"github.com/zeropsio/zerops-go/dto/input/path"
	"github.com/zeropsio/zerops-go/types"
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
	Token     string         `json:"token"`
	ProjectId uuid.ProjectId `json:"projectId"`
	Name      types.String   `json:"projectId"`
}

func (h *Handler) McpTool() (mcp.Tool, mcpServer.ToolHandlerFunc) {
	return mcp.NewTool(
		"service_create",
		mcp.WithDescription("Create new service"),
		mcp.WithString("token",
			mcp.Description("sdk access token"),
			mcp.Required(),
		),
		mcp.WithString("projectId",
			mcp.Description("project id"),
			mcp.Required(),
		),
		mcp.WithString("name",
			mcp.Description("service name"),
			mcp.Required(),
		),
	), mcp.NewTypedToolHandler(h.Handle)
}

func (h *Handler) Handle(ctx context.Context, _ mcp.CallToolRequest, params Params) (*mcp.CallToolResult, error) {
	_, err := h.sdk.Authorized(params.Token).PostServiceStack(ctx, path.ServiceStackServiceStackTypeVersionId{
		ServiceStackTypeVersionId: "alpine_v3_21",
	}, body.PostStandardServiceStack{
		ProjectId:        params.ProjectId,
		Name:             params.Name,
		StartWithoutCode: types.NewBoolNull(true),
	})
	if err != nil {
		return nil, err
	}
	return mcp.NewToolResultText("success"), nil
}
