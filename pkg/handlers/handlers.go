package handlers

import (
	"net/http"

	"github.com/CevdetAkiah/Registration/pkg/config"
	"github.com/CevdetAkiah/Registration/pkg/models"
	"github.com/CevdetAkiah/Registration/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

//Executes the index html page
func (m *Repository) Index(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "index.page.gohtml", nil)

}

//Executes the register gohtml page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//test
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	render.RenderTemplate(w, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})

}
