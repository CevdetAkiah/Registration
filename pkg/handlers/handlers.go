package handlers

import (
	"log"
	"net/http"

	"github.com/CevdetAkiah/Registration/pkg/config"
	"github.com/CevdetAkiah/Registration/pkg/forms"
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
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	reservation := models.Registration{
		Email:    r.Form.Get("email"),
		Password: r.Form.Get("pwd"),
	}

	form := forms.New(r.PostForm)

	form.Has("email", r)
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "register.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})

		return
	}
}
