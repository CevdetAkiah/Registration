package handlers

import (
	"net/http"

	"github.com/CevdetAkiah/Registration/pkg/config"
	"github.com/CevdetAkiah/Registration/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

//NewRepo sends the Repository type to main providing a container for the appConfig.
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers allows main to provide the app config including the current template cache to the handlers.
func NewHandlers(r *Repository) {
	Repo = r
}

//Index renders the index html page
func (m *Repository) Index(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "index.page.html", nil)
}

//About renders the about html page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.html", nil)
}

//Register renders the register html page
func (m *Repository) Register(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "register.page.html", nil)
}
