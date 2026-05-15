package appservice

import (
	"context"
)

type Status struct {
	ExecutionState string
	Progress       int
}

type ControllerService interface {
	// synchronizes its clock with the central controller clock.
	// default: 1s time base
	// responds with a status description message
	GetStatus(ctx context.Context) (Status, error)
}
