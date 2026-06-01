// The controller component of the GTD page in the web presentation layer.
package gtd

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/tombenke/go-12f-common/v2/log"
	"github.com/tombenke/go-application-template/internal/application"
	"github.com/tombenke/go-application-template/internal/infrastructure/presentation/web/components/common"
)

type Controller struct {
	gtdManager application.GTDManager
	presenter  *Presenter
}

func NewController(gtdManager application.GTDManager) *Controller {
	return &Controller{
		gtdManager: gtdManager,
		presenter:  NewPresenter(),
	}
}

func (c *Controller) HandleIndex(w http.ResponseWriter, _ *http.Request) {

	data := indexPageData{
		PageData: common.NewPageData("Demo Contacts Application - Home"),
	}

	c.presenter.PresentIndex(w, data)
}

func (c *Controller) HandleHelp(w http.ResponseWriter, _ *http.Request) {
	data := common.NewPageData("Demo Contacts Application - Help")
	c.presenter.PresentHelp(w, data)
}

func (c *Controller) HandleNewContact(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := c.getLogger(ctx)
	logger.Debug("Handling new contact request", "method", r.Method)

	data := newContactPageData{
		PageData: common.NewPageData("A Demo Contacts Application - New Contact"),
		Contact:  &application.ContactDTO{},
	}

	c.presenter.PresentNewContact(w, data)
}

func (c *Controller) HandleContactsGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := c.getLogger(ctx)
	logger.Debug("Handling contacts request", "method", r.Method)
	contacts, _ := c.gtdManager.GetContacts(ctx)

	data := contactsPageData{
		PageData: common.NewPageData("A Demo Contacts Application - Contacts"),
		Contacts: contacts,
	}

	c.presenter.PresentContacts(w, data)
}

func (c *Controller) HandleContactsPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := c.getLogger(ctx)
	logger.Debug("Handling contacts request", "method", r.Method)
	id, err := uuid.Parse(r.FormValue("id"))
	if err != nil {
		logger.Error("Failed to parse contact ID", "error", err)
		id = uuid.New() // fallback to new UUID
	}

	contactDTO := application.ContactDTO{
		ID:        id,
		Email:     r.FormValue("email"),
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Phone:     r.FormValue("phone"),
	}
	// TODO validation
	logger.Debug("Received new contact data", "contact", contactDTO)

	err = c.gtdManager.AddContact(ctx, contactDTO)
	if err != nil {
		logger.Error("Failed to add contact", "error", err)
		http.Error(w, "Failed to add contact", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func (c *Controller) HandleContactsPut(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := c.getLogger(ctx)
	logger.Debug("Handling contacts request", "method", r.Method)
	id, err := uuid.Parse(r.FormValue("id"))
	if err != nil {
		logger.Error("Failed to parse contact ID", "error", err)
		id = uuid.New() // fallback to new UUID
	}
	contactDTO := application.ContactDTO{
		ID:        id,
		Email:     r.FormValue("email"),
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Phone:     r.FormValue("phone"),
	}

	// TODO validation
	logger.Debug("Received contact data to update", "contact", contactDTO)

	err = c.gtdManager.UpdateContact(ctx, contactDTO)
	if err != nil {
		logger.Error("Failed to update contact", "error", err)
		http.Error(w, "Failed to update contact", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func (c *Controller) HandleContactsDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := c.getLogger(ctx)
	logger.Debug("Handling contacts request", "method", r.Method)
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		logger.Error("Failed to parse contact ID", "error", err)
		id = uuid.New() // fallback to new UUID
	}

	// TODO validation
	logger.Debug("Received contact data to delete", "contact", id)

	err = c.gtdManager.DeleteContact(ctx, id)
	if err != nil {
		logger.Error("Failed to delete contact", "error", err)
		http.Error(w, "Failed to delete contact", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

func (*Controller) getLogger(ctx context.Context) *slog.Logger {
	return log.GetFromContextOrDefault(ctx).With("controller", "ContactController")
}
