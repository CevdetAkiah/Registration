package handlers

import (
	"fmt"
	"net/http"

	"github.com/CevdetAkiah/Registration/pkg/config"
	"github.com/CevdetAkiah/Registration/pkg/models"
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
	render.RenderTemplate(w, r, "index.page.html", &models.TemplateData{})
}

//About renders the about html page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{})
}

//Register renders the register html page
func (m *Repository) Register(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "register.page.html", &models.TemplateData{})
}

//PostRegister renders the register html page
func (m *Repository) PostRegister(w http.ResponseWriter, r *http.Request) {

	email := r.Form.Get("email")
	password := r.Form.Get("pwd")
	fmt.Printf("Email is %s and password is %s", email, password)
}
