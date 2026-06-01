package common

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"

	"github.com/tombenke/go-12f-common/v2/must"
)

type TemplateRegistry struct {
	pageTemplates       map[string]*template.Template
	componentsTemplates map[string]*template.Template
}

func NewTemplateRegistry(fss fs.FS) (*TemplateRegistry, error) {
	pageTemplates := must.MustVal(buildPageTemplates(fss))
	componentsTemplates := must.MustVal(buildComponentTemplates(fss))

	return &TemplateRegistry{pageTemplates: pageTemplates, componentsTemplates: componentsTemplates}, nil
}

func buildPageTemplates(fss fs.FS) (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)
	pages, err := fs.Glob(fss, "templates/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFS(fss, "templates/layouts/*.html", "templates/pages/"+name, "templates/partials/*.html")
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}

func buildComponentTemplates(fss fs.FS) (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)
	components, err := fs.Glob(fss, "templates/components/*.html")
	if err != nil {
		return nil, err
	}

	for _, component := range components {
		name := filepath.Base(component)

		ts, err := template.ParseFS(fss, "templates/layouts/*.html", "templates/components/"+name, "templates/partials/*.html")
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}

func (reg *TemplateRegistry) RenderPage(w http.ResponseWriter, layoutName, pageName string, data any) {
	tmpl, exists := reg.pageTemplates[pageName]
	if !exists {
		http.Error(w, fmt.Sprintf("Page '%s' not found", pageName), http.StatusNotFound)
		return
	}

	renderTemplate(w, tmpl, layoutName, data)
}

func (reg *TemplateRegistry) RenderComponent(w http.ResponseWriter, layoutName, componentName string, data any) {
	tmpl, exists := reg.componentsTemplates[componentName]
	if !exists {
		http.Error(w, fmt.Sprintf("Component '%s' not found", componentName), http.StatusNotFound)
		return
	}

	renderTemplate(w, tmpl, layoutName, data)
}

func renderTemplate(w http.ResponseWriter, tmpl *template.Template, templateName string, data any) {
	err := tmpl.ExecuteTemplate(w, templateName, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
