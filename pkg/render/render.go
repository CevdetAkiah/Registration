package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/CevdetAkiah/Registration/pkg/config"
	"github.com/CevdetAkiah/Registration/pkg/models"
)

var (
	functions = template.FuncMap{}
	app       *config.AppConfig
)

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc = make(map[string]*template.Template)

	// get the template cache from the app config
	//If UseCache is disabled it means we can make changes on the fly to HTML pages (this is called development mode)
	//If enabled we're in production mode, and the currently parsed cache will be in use (no on the fly changes).
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Can't find template in the cache.")
	}

	td = AddDefaultData(td)

	// buf := new(bytes.Buffer)
	t.Execute(w, td)

	// _, err := buf.WriteTo(w)
	// if err != nil {
	// 	log.Println("Error writing template to browser. ", err)
	// }

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)
	pages, err := filepath.Glob("../../templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("../../templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, err
}
