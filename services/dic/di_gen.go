//go:build !templater && !di
// +build !templater,!di

// Code generated by "templater"; DO NOT EDIT.
// source: di.go
package dic

// gen
import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/google/wire"
	"github.com/spf13/cobra"
	"github.com/zerops-dev/di/app"
	"github.com/zerops-dev/di/appRunner"
	configurator "github.com/zerops-dev/di/di"
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
	"github.com/zeropsio/zerops-go/sdkBase"
)

type Config struct {
	NameLoggerConfig1      logger.Config
	NameHttpClientConfig14 httpClient.Config
	NameSdkBaseConfig17    sdkBase.Config
	NameServerConfig18     server.Config
}

func NewConfig(prefix string,
	configurator *configurator.Handler,
) *Config {
	c := &Config{
		NameLoggerConfig1:      (logger.Config)(logger.NewConfig()),
		NameHttpClientConfig14: (httpClient.Config)(httpClient.NewConfig()),
		NameSdkBaseConfig17:    (sdkBase.Config)(zeropsSdk.NewConfig()),
		NameServerConfig18:     (server.Config)(server.NewConfig()),
	}
	configurator.Register(concat(prefix, "log"), &c.NameLoggerConfig1)
	configurator.Register(concat(prefix, "httpClient"), &c.NameHttpClientConfig14)
	configurator.Register(concat(prefix, "sdk"), &c.NameSdkBaseConfig17)
	configurator.Register(concat(prefix, "server"), &c.NameServerConfig18)

	return c
}

type NameLoggerConfig1 logger.Config

func GetNameLoggerConfig1(c Config) NameLoggerConfig1 {
	return NameLoggerConfig1(c.NameLoggerConfig1)
}

type NameHttpClientConfig14 httpClient.Config

func GetNameHttpClientConfig14(c Config) NameHttpClientConfig14 {
	return NameHttpClientConfig14(c.NameHttpClientConfig14)
}

type NameSdkBaseConfig17 sdkBase.Config

func GetNameSdkBaseConfig17(c Config) NameSdkBaseConfig17 {
	return NameSdkBaseConfig17(c.NameSdkBaseConfig17)
}

type NameServerConfig18 server.Config

func GetNameServerConfig18(c Config) NameServerConfig18 {
	return NameServerConfig18(c.NameServerConfig18)
}

type AppConfig struct {
	configurator *configurator.Handler
	config       *Config
}

func NewAppConfig(prefix string, cmd *cobra.Command) *AppConfig {
	configurator := configurator.New(cmd)
	return &AppConfig{
		configurator: configurator,
		config:       NewConfig(prefix, configurator),
	}
}

func (c *AppConfig) Read(app *app.ApplicationSetup, cmd *cobra.Command) error {
	return c.configurator.Read(app, cmd)
}

func (c *AppConfig) Configurator() *configurator.Handler {
	return c.configurator
}

func (c *AppConfig) Register(prefix string, in interface{}) {
	c.configurator.Register(prefix, in)
}

func (c *AppConfig) ConfigCommand(app *app.ApplicationSetup) *cobra.Command {
	return c.configurator.ConfigCommand(app)
}

func (c *AppConfig) Config() Config {
	return *c.config
}

func concat(parts ...string) string {
	var result []string
	for _, part := range parts {
		p := strings.TrimPrefix(part, "<")
		if strings.HasPrefix(part, "<") && len(result) > 0 {
			result[len(result)-1] = p
		} else {
			result = append(result, p)
		}
	}
	return strings.Join(result, ".")
}

type NameServiceStartHandler10 serviceStart.Handler

func GetNameServiceStartHandler10(arg0 *NameZeropsSdkHandler16,
) *NameServiceStartHandler10 {

	return (*NameServiceStartHandler10)(serviceStart.New((*zeropsSdk.Handler)(arg0)))

}

type NameServiceCreateHandler8 serviceCreate.Handler

func GetNameServiceCreateHandler8(arg0 *NameZeropsSdkHandler16,
) *NameServiceCreateHandler8 {

	return (*NameServiceCreateHandler8)(serviceCreate.New((*zeropsSdk.Handler)(arg0)))

}

type NameAppRunnerRegister3 appRunner.Register

func GetNameAppRunnerRegister3(arg0 *NameAppRunnerHandler2,
) NameAppRunnerRegister3 {

	return (NameAppRunnerRegister3)(appRunner.DiBind((*appRunner.Handler)(arg0)))

}

type NameServerHandler4 server.Handler

func GetNameServerHandler4(arg0 NameServerConfig18,
	arg1 NameAppRunnerRegister3,
) (*NameServerHandler4, error) {

	v, err := server.New((server.Config)(arg0),
		(appRunner.Register)(arg1),
	)
	return (*NameServerHandler4)(v), err

}

type NameZeropsSdkHandler16 zeropsSdk.Handler

func GetNameZeropsSdkHandler16(arg0 NameSdkBaseConfig17,
	arg1 *NameHttpClient13,
) *NameZeropsSdkHandler16 {

	return (*NameZeropsSdkHandler16)(zeropsSdk.New((sdkBase.Config)(arg0),
		(*http.Client)(arg1),
	))

}

type NameSdkHandler15 sdk.Handler

func GetNameSdkHandler15(arg0 NameSdkBaseConfig17,
	arg1 *NameHttpClient13,
) NameSdkHandler15 {

	return (NameSdkHandler15)(sdk.New((sdkBase.Config)(arg0),
		(*http.Client)(arg1),
	))

}

type NameAppRunnerHandler2 appRunner.Handler

func GetNameAppRunnerHandler2(arg0 *NameSlogLogger0,
) *NameAppRunnerHandler2 {

	return (*NameAppRunnerHandler2)(appRunner.New((*slog.Logger)(arg0)))

}

type NameServiceRestartHandler11 serviceRestart.Handler

func GetNameServiceRestartHandler11(arg0 *NameZeropsSdkHandler16,
) *NameServiceRestartHandler11 {

	return (*NameServiceRestartHandler11)(serviceRestart.New((*zeropsSdk.Handler)(arg0)))

}

type NameHttpClient13 http.Client

func GetNameHttpClient13(arg0 NameHttpClientConfig14,
) *NameHttpClient13 {

	return (*NameHttpClient13)(httpClient.New((httpClient.Config)(arg0)))

}

type NameSlogLogger0 slog.Logger

func GetNameSlogLogger0(arg0 NameLoggerConfig1,
) *NameSlogLogger0 {

	return (*NameSlogLogger0)(logger.New((logger.Config)(arg0)))

}

type NameServiceStopHandler9 serviceStop.Handler

func GetNameServiceStopHandler9(arg0 *NameZeropsSdkHandler16,
) *NameServiceStopHandler9 {

	return (*NameServiceStopHandler9)(serviceStop.New((*zeropsSdk.Handler)(arg0)))

}

type NameServiceDeleteHandler12 serviceDelete.Handler

func GetNameServiceDeleteHandler12(arg0 *NameZeropsSdkHandler16,
) *NameServiceDeleteHandler12 {

	return (*NameServiceDeleteHandler12)(serviceDelete.New((*zeropsSdk.Handler)(arg0)))

}

type NameContainerReadDirectoryHandler5 containerReadDirectory.Handler

func GetNameContainerReadDirectoryHandler5() *NameContainerReadDirectoryHandler5 {

	return (*NameContainerReadDirectoryHandler5)(containerReadDirectory.New())

}

type NameContainerWriteFileHandler6 containerWriteFile.Handler

func GetNameContainerWriteFileHandler6() *NameContainerWriteFileHandler6 {

	return (*NameContainerWriteFileHandler6)(containerWriteFile.New())

}

type NameContainerReadFileHandler7 containerReadFile.Handler

func GetNameContainerReadFileHandler7() *NameContainerReadFileHandler7 {

	return (*NameContainerReadFileHandler7)(containerReadFile.New())

}

type setter_0 struct{}
type setter_0_Value server.Tools

func getSetter_0_Value(target *NameServiceCreateHandler8) (setter_0_Value, error) {

	return (setter_0_Value)((*serviceCreate.Handler)(target)), nil

}

func setSetter_0(source *NameServerHandler4, target setter_0_Value) *setter_0 {

	server.RegisterTools((*server.Handler)(source), target)
	return &setter_0{}

}

type setter_1 struct{}
type setter_1_Value server.Tools

func getSetter_1_Value(target *NameServiceStopHandler9) (setter_1_Value, error) {

	return (setter_1_Value)((*serviceStop.Handler)(target)), nil

}

func setSetter_1(source *NameServerHandler4, target setter_1_Value) *setter_1 {

	server.RegisterTools((*server.Handler)(source), target)
	return &setter_1{}

}

type setter_2 struct{}
type setter_2_Value server.Tools

func getSetter_2_Value(target *NameServiceStartHandler10) (setter_2_Value, error) {

	return (setter_2_Value)((*serviceStart.Handler)(target)), nil

}

func setSetter_2(source *NameServerHandler4, target setter_2_Value) *setter_2 {

	server.RegisterTools((*server.Handler)(source), target)
	return &setter_2{}

}

type setter_3 struct{}
type setter_3_Value server.Tools

func getSetter_3_Value(target *NameServiceRestartHandler11) (setter_3_Value, error) {

	return (setter_3_Value)((*serviceRestart.Handler)(target)), nil

}

func setSetter_3(source *NameServerHandler4, target setter_3_Value) *setter_3 {

	server.RegisterTools((*server.Handler)(source), target)
	return &setter_3{}

}

type setter_4 struct{}
type setter_4_Value server.Tools

func getSetter_4_Value(target *NameServiceDeleteHandler12) (setter_4_Value, error) {

	return (setter_4_Value)((*serviceDelete.Handler)(target)), nil

}

func setSetter_4(source *NameServerHandler4, target setter_4_Value) *setter_4 {

	server.RegisterTools((*server.Handler)(source), target)
	return &setter_4{}

}

type setter_5 struct{}
type setter_5_Value server.Tools

func getSetter_5_Value(target *NameContainerReadDirectoryHandler5) (setter_5_Value, error) {

	return (setter_5_Value)((*containerReadDirectory.Handler)(target)), nil

}

func setSetter_5(source *NameServerHandler4, target setter_5_Value) *setter_5 {

	server.RegisterTools((*server.Handler)(source), target)
	return &setter_5{}

}

type setter_6 struct{}
type setter_6_Value server.Tools

func getSetter_6_Value(target *NameContainerWriteFileHandler6) (setter_6_Value, error) {

	return (setter_6_Value)((*containerWriteFile.Handler)(target)), nil

}

func setSetter_6(source *NameServerHandler4, target setter_6_Value) *setter_6 {

	server.RegisterTools((*server.Handler)(source), target)
	return &setter_6{}

}

type setter_7 struct{}
type setter_7_Value server.Tools

func getSetter_7_Value(target *NameContainerReadFileHandler7) (setter_7_Value, error) {

	return (setter_7_Value)((*containerReadFile.Handler)(target)), nil

}

func setSetter_7(source *NameServerHandler4, target setter_7_Value) *setter_7 {

	server.RegisterTools((*server.Handler)(source), target)
	return &setter_7{}

}

var DepSet = wire.NewSet(
	app.Set,
	NewApplication, getSetter_0_Value, setSetter_0,
	getSetter_1_Value, setSetter_1,
	getSetter_2_Value, setSetter_2,
	getSetter_3_Value, setSetter_3,
	getSetter_4_Value, setSetter_4,
	getSetter_5_Value, setSetter_5,
	getSetter_6_Value, setSetter_6,
	getSetter_7_Value, setSetter_7,
	GetNameServiceStartHandler10,
	GetNameServiceCreateHandler8,
	GetNameAppRunnerRegister3,
	GetNameServerHandler4,
	GetNameZeropsSdkHandler16,
	GetNameSdkHandler15,
	GetNameAppRunnerHandler2,
	GetNameServiceRestartHandler11,
	GetNameHttpClient13,
	GetNameSlogLogger0,
	GetNameServiceStopHandler9,
	GetNameServiceDeleteHandler12,
	GetNameContainerReadDirectoryHandler5,
	GetNameContainerWriteFileHandler6,
	GetNameContainerReadFileHandler7,

	GetNameLoggerConfig1,
	GetNameHttpClientConfig14,
	GetNameSdkBaseConfig17,
	GetNameServerConfig18,
)

func NewApplication(
	systemApp *app.ApplicationSetup, _ *setter_0,
	_ *setter_1,
	_ *setter_2,
	_ *setter_3,
	_ *setter_4,
	_ *setter_5,
	_ *setter_6,
	_ *setter_7,
	_ *NameServiceStartHandler10,
	_ *NameServiceCreateHandler8,
	_ NameAppRunnerRegister3,
	_ *NameServerHandler4,
	_ *NameZeropsSdkHandler16,
	_ NameSdkHandler15,
	runner26 *NameAppRunnerHandler2,
	_ *NameServiceRestartHandler11,
	_ *NameHttpClient13,
	logger09 *NameSlogLogger0,
	_ *NameServiceStopHandler9,
	_ *NameServiceDeleteHandler12,
	_ *NameContainerReadDirectoryHandler5,
	_ *NameContainerWriteFileHandler6,
	_ *NameContainerReadFileHandler7,
) *App {
	return &App{
		app:    systemApp,
		runner: (*appRunner.Handler)(runner26),
		logger: (*slog.Logger)(logger09),
	}
}

type App struct {
	app    *app.ApplicationSetup
	runner *appRunner.Handler
	logger *slog.Logger
}

func (h *App) GetRunner() *appRunner.Handler {
	return h.runner
}
func (h *App) GetLogger() *slog.Logger {
	return h.logger
}
func (h *App) Run() error {
	return h.runner.RunWithSigTerm(h.app.Context, h.app.Cancel)
}

func CreateCommand(app *app.ApplicationSetup) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "run",
		Short: "run command",
		Long:  "",
	}

	cfg := NewAppConfig("zerops-mcp", cmd)

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if err := cfg.Read(app, cmd); err != nil {
			return err
		}

		application, err := CreateApp(app, cfg.Config())
		if err != nil {
			panic(err)
		}
		return application.Run()
	}
	cmd.AddCommand(cfg.ConfigCommand(app))

	return cmd
}
