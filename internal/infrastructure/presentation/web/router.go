// The central router component of the web presentation layer.
package web

import (
	"net/http"

	"github.com/tombenke/go-application-template/internal/infrastructure/presentation/web/components/gtd"
)

func NewRouter(controller *gtd.Controller) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", controller.HandleIndex)
	mux.HandleFunc("GET /contacts/new", controller.HandleNewContact)
	mux.HandleFunc("GET /contacts", controller.HandleContactsGet)
	mux.HandleFunc("POST /contacts", controller.HandleContactsPost)
	mux.HandleFunc("PUT /contacts/{id}", controller.HandleContactsPut)
	mux.HandleFunc("DELETE /contacts/{id}", controller.HandleContactsDelete)
	mux.HandleFunc("GET /help", controller.HandleHelp)
	return mux
}
