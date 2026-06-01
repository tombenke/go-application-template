// The presenter (view) component of the contact page in the web presentation layer.
package gtd

import (
	"net/http"

	"embed"

	"github.com/tombenke/go-12f-common/v2/must"
	"github.com/tombenke/go-application-template/internal/infrastructure/presentation/web/components/common"
)

//go:embed templates/**/*.html
var ComponentTemplatesFS embed.FS

type Presenter struct {
	templates *common.TemplateRegistry
}

func NewPresenter() *Presenter {
	return &Presenter{
		templates: must.MustVal(common.NewTemplateRegistry(common.NewMultiFS(common.CommonTemplatesFS, ComponentTemplatesFS))),
	}
}

func (p *Presenter) PresentContacts(w http.ResponseWriter, data contactsPageData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	p.templates.RenderPage(w, "base", "contacts.html", data)
}

func (p *Presenter) PresentNewContact(w http.ResponseWriter, data newContactPageData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	p.templates.RenderPage(w, "base", "new_contact.html", data)
}

func (p *Presenter) PresentIndex(w http.ResponseWriter, data indexPageData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	p.templates.RenderPage(w, "base", "index.html", data)
}

func (p *Presenter) PresentHelp(w http.ResponseWriter, data common.PageData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	p.templates.RenderPage(w, "base", "help.html", data)
}
