package dic

import (
	"github.com/zerops-dev/di/appRunner"
	"github.com/zerops-dev/di/di/s"
	"github.com/zerops-dev/di/logger"
	"github.com/zerops-dev/zerops-mcp/services/httpClient"
	"github.com/zerops-dev/zerops-mcp/services/server"
	"github.com/zerops-dev/zerops-mcp/services/zeropsSdk"
	"github.com/zerops-dev/zerops-mcp/tools/containerReadDirectory"
	"github.com/zerops-dev/zerops-mcp/tools/containerReadFile"
	"github.com/zerops-dev/zerops-mcp/tools/containerWriteFile"
	"github.com/zerops-dev/zerops-mcp/tools/serviceCreate"
	"github.com/zerops-dev/zerops-mcp/tools/serviceDelete"
	"github.com/zerops-dev/zerops-mcp/tools/serviceRestart"
	"github.com/zerops-dev/zerops-mcp/tools/serviceStart"
	"github.com/zerops-dev/zerops-mcp/tools/serviceStop"
	"github.com/zeropsio/zerops-go/sdk"
	_ "golang.org/x/tools/imports"
)

//go:generate templater -tags "di" -templateTags "!templater,!di"
var _ = func() *s.Di {
	di := s.NewDi("zerops-mcp",
		s.WithCommand(
			"run",
			"run command",
			``,
		),
	)

	di.Add(
		s.Scope(
			logger.DiScope(),
			appRunner.DiScope(),
			s.Service(server.New, s.WithSetter(server.RegisterTools)),

			s.Service(containerReadDirectory.New),
			s.Service(containerWriteFile.New),
			s.Service(containerReadFile.New),

			s.Service(serviceCreate.New),
			s.Service(serviceStop.New),
			s.Service(serviceStart.New),
			s.Service(serviceRestart.New),
			s.Service(serviceDelete.New),

			httpClient.DiScope(),

			s.Service(sdk.New),

			s.Service(zeropsSdk.New),
			s.Config(zeropsSdk.NewConfig, "sdk"),

			s.Config(server.NewConfig, "server"),
		),
	)

	return di
}
