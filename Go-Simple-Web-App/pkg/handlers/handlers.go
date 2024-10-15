package handlers

import (
	"myGoWebApp/pkg/config"
	"myGoWebApp/pkg/models"
	"myGoWebApp/pkg/render"
	"net/http"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the reposity type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets repo for new handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page value
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, from a map"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
