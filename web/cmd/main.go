package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CevdetAkiah/Registration/pkg/config"
	"github.com/CevdetAkiah/Registration/pkg/handlers"
	"github.com/CevdetAkiah/Registration/pkg/render"
)

const portNumber = ":8080"

var (
	app config.AppConfig
)

func main() {

	//Create the template cache.
	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Can't create template cache ", err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplates(&app) //send the template cache to the render package.

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo) //send the template cache to the handler package.

	//Handles the URLs
	http.HandleFunc("/", handlers.Repo.Index)
	http.HandleFunc("/about", handlers.Repo.About)

	// srv := &http.Server{
	// 	Addr:    portNumber,
	// 	Handler: routes(&app),
	// }

	//Listens on port :8080 for any requests and handles the error.
	log.Fatal(http.ListenAndServe(portNumber, nil))
}
