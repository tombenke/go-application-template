package controller

import (
	"context"
	"log/slog"

	"github.com/tombenke/go-12f-common/v2/log"
	appservice "github.com/tombenke/go-application-template/internal/app_service"
)

type Controller struct {
	//scenarioService domain.ScenarioService
}

var _ appservice.ControllerService = (*Controller)(nil)

func NewController( /*inject services here*/ ) *Controller {
	return &Controller{
		//scenarioService: scenarioService,
	}
}

func (c *Controller) GetStatus(ctx context.Context) (appservice.Status, error) {
	logger := c.getLogger(ctx)
	logger.Debug("GetStatus called")

	// status, err := c.scenarioService.Status(ctx)
	// if err != nil {
	// 	return appservice.Status{}, fmt.Errorf("failed to get scenario status: %w", err)
	// }
	status := appservice.Status{
		ExecutionState: "running",
		Progress:       50,
	}

	return status, nil
}

func (*Controller) getLogger(ctx context.Context) *slog.Logger {
	return log.GetFromContextOrDefault(ctx).With("component", "Controller")
}
