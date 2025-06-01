package main

import (
	"context"
	"path"
	"strings"

	"github.com/zerops-dev/zerops-mcp/services/cmd"
	runApp "github.com/zerops-dev/di/app"
)

var (
	Name            = "zerops-mcp"
	Service         = "zerops-mcp"
	Exec            = "/usr/local/bin/zerops-mcp"
	BuildTime       = "time"
	Commit          = "commit"
	Version         = "v0.0.0"
	Description     = "Zerops mcp"
	DescriptionLong = "Zerops mcp"
)

func main() {

	applicationSetup := runApp.New(context.Background(), Name)
	applicationSetup.Exec = strings.Join([]string{Exec, "run", "--config", path.Join("/etc", Service, "config.yml")}, " ")
	applicationSetup.Version = Version
	applicationSetup.Commit = Commit
	applicationSetup.Service = Service
	applicationSetup.Description = Description
	applicationSetup.DescriptionLong = DescriptionLong
	applicationSetup.SetBuildTime(BuildTime)

	applicationSetup.RegisterCommands()
	cmd.AddCommands(applicationSetup, applicationSetup.RootCommand())
	applicationSetup.Run()
}
