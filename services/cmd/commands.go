package cmd

import (
	"github.com/zerops-dev/zerops-mcp/services/dic"
	"github.com/spf13/cobra"
	runApp "github.com/zerops-dev/di/app"
)

func AddCommands(applicationSetup *runApp.ApplicationSetup, rootCommand *cobra.Command) {
	rootCommand.AddCommand(dic.CreateCommand(applicationSetup))
}
