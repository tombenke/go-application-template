package webserver

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"sync"
	"time"

	"github.com/tombenke/go-12f-common/v2/apprun"
	"github.com/tombenke/go-12f-common/v2/healthcheck"
	"github.com/tombenke/go-12f-common/v2/log"
	appservice "github.com/tombenke/go-application-template/internal/app_service"
)

//go:embed templates/*.html
var templatesFS embed.FS

type WebServer struct {
	config     *Config
	controller appservice.ControllerService
	sessions   *SessionStore
	templates  *template.Template
	router     http.Handler
	httpServer *http.Server
	err        error
	ready      bool
	mu         sync.RWMutex
}

var _ apprun.ComponentLifecycleManager = (*WebServer)(nil)

func NewWebServer(cfg *Config, controller appservice.ControllerService) (*WebServer, error) {
	tmpl, err := template.ParseFS(templatesFS, "templates/*.html")
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates: %w", err)
	}

	ws := &WebServer{
		config:     cfg,
		controller: controller,
		sessions:   NewSessionStore(30 * time.Minute),
		templates:  tmpl,
		err:        healthcheck.ServiceNotAvailableError{},
	}

	ws.registerRoutes()

	return ws, nil
}

func (ws *WebServer) Startup(ctx context.Context, _ *sync.WaitGroup) error {
	log.GetFromContextOrDefault(ctx).Debug("WebServer: Startup", "port", ws.config.Port)

	ws.httpServer = &http.Server{
		Addr:         fmt.Sprintf(":%d", ws.config.Port),
		Handler:      ws.router,
		ReadTimeout:  ws.config.ReadTimeout,
		WriteTimeout: ws.config.WriteTimeout,
		IdleTimeout:  ws.config.IdleTimeout,
	}

	ws.mu.Lock()
	ws.ready = true
	ws.err = nil
	ws.mu.Unlock()

	go func() {
		if err := ws.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			ws.mu.Lock()
			ws.err = err
			ws.ready = false
			ws.mu.Unlock()
		}
	}()

	return nil
}

func (ws *WebServer) Shutdown(ctx context.Context) error {
	log.GetFromContextOrDefault(ctx).Debug("WebServer: Shutdown")

	ws.mu.Lock()
	ws.ready = false
	ws.err = healthcheck.ServiceNotAvailableError{}
	ws.mu.Unlock()

	if ws.httpServer == nil {
		return nil
	}

	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return ws.httpServer.Shutdown(shutdownCtx)
}

func (ws *WebServer) Check(ctx context.Context) error {
	log.GetFromContextOrDefault(ctx).Debug("WebServer: Check")
	ws.mu.RLock()
	defer ws.mu.RUnlock()
	return ws.err
}

func (ws *WebServer) Readiness(_ context.Context) error {
	ws.mu.RLock()
	defer ws.mu.RUnlock()
	if ws.ready {
		return nil
	}
	return healthcheck.ServiceNotAvailableError{}
}

func (ws *WebServer) registerRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", ws.handleIndex)

	ws.router = ws.withSession(mux)

	return ws.router
}
