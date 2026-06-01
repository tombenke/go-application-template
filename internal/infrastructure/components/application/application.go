package application

import (
	"context"
	"log/slog"
	"sync"

	"github.com/tombenke/go-12f-common/v2/apprun"
	"github.com/tombenke/go-12f-common/v2/buildinfo"
	"github.com/tombenke/go-12f-common/v2/log"
	"github.com/tombenke/go-12f-common/v2/must"
	appservices "github.com/tombenke/go-application-template/internal/application"
	"github.com/tombenke/go-application-template/internal/infrastructure/components/webserver"
	"github.com/tombenke/go-application-template/internal/infrastructure/presentation/web"
	"github.com/tombenke/go-application-template/internal/infrastructure/presentation/web/components/gtd"
	inmemory "github.com/tombenke/go-application-template/internal/infrastructure/repositories/in_memory"
)

var _ apprun.Application = (*Application)(nil)

// Application represents the overall application object.
type Application struct {
	config *Config
	web    *webserver.WebServer

	gtd appservices.GTDManager
	// The internal components of the application
	components []apprun.ComponentLifecycleManager
}

// Creates a new application instance.
func NewApplication(config *Config) (apprun.Application, error) {

	gtdRepository := inmemory.NewGTDRepository()
	gtdManager := appservices.NewGTDManager(gtdRepository)
	contactsController := gtd.NewController(gtdManager)

	router := web.NewRouter(contactsController)
	webComponent := must.MustVal(webserver.NewWebServer(&config.webserver, router))

	// Create and return the application object
	return &Application{
		config: config,
		gtd:    gtdManager,
		web:    webComponent,
		components: []apprun.ComponentLifecycleManager{
			webComponent,
		},
	}, nil
}

func (a *Application) Components(_ context.Context) []apprun.ComponentLifecycleManager {
	return a.components
}

// Startup the application.
func (a *Application) AfterStartup(ctx context.Context, _ *sync.WaitGroup) error {
	a.getLogger(ctx).Info("AfterStartup")
	a.getLogger(ctx).Debug("BuildInfo", "AppName", buildinfo.AppName(), "Version", buildinfo.Version())

	return nil
}

// Shutdown the application.
func (a *Application) BeforeShutdown(ctx context.Context) error {
	a.getLogger(ctx).Info("BeforeShutdown")
	return nil
}

func (*Application) getLogger(ctx context.Context) *slog.Logger {
	return log.GetFromContextOrDefault(ctx).With("app", "Application")
}
