package common

import (
	"embed"
)

//go:embed templates/**/*.html
var CommonTemplatesFS embed.FS

const CompanyName = "Demo Company"
const ApplicationName = "GTD"

type ErrorResponse struct {
	Message string `json:"message"`
}

type PageData struct {
	Title           string
	Year            int
	CompanyName     string
	ApplicationName string
}

func NewPageData(title string) PageData {
	return PageData{
		Title:           title,
		Year:            2024,
		CompanyName:     CompanyName,
		ApplicationName: ApplicationName,
	}
}

/*
TODO: These methods are not used in the current implementation, but they can be useful for handling JSON responses and errors in the future. We can keep them here for now, and if we find that they are not needed, we can remove them later.

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
*/
