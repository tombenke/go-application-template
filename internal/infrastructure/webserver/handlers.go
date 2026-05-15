package webserver

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/tombenke/go-12f-common/v2/log"
)

type errorResponse struct {
	Message string `json:"message"`
}

type pageData struct {
	Title       string
	Status      string
	Error       string
	SessionID   string
	GeneratedAt string
}

func (ws *WebServer) handleIndex(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := ws.getLogger(ctx)
	logger.Debug("GetStatus called")
	status, _ := ws.controller.GetStatus(ctx)
	statusText := status.ExecutionState
	if statusText == "" {
		statusText = "stopped"
	}

	data := pageData{
		Title:       "Propel Scenario Management",
		Status:      statusText,
		SessionID:   SessionIDFromContext(ctx),
		GeneratedAt: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := ws.templates.ExecuteTemplate(w, "layout", data); err != nil {
		ws.writeError(w, http.StatusInternalServerError, err)
		return
	}
}

func (*WebServer) writeJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}

func (ws *WebServer) writeError(w http.ResponseWriter, statusCode int, err error) {
	ws.getLogger(context.Background()).Debug("HTTP error", "status", statusCode, "error", err)
	ws.writeJSON(w, statusCode, errorResponse{Message: err.Error()})
}

func (*WebServer) getLogger(ctx context.Context) *slog.Logger {
	return log.GetFromContextOrDefault(ctx).With("component", "WebServer")
}
