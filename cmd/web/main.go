package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/CevdetAkiah/Registration/pkg/config"
	"github.com/CevdetAkiah/Registration/pkg/handlers"
	"github.com/CevdetAkiah/Registration/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var (
	app config.AppConfig
)

func main() {

	app.InProduction = false //Will use this when encrypting cookies for the session. Set to false when in dev mode.

	session := scs.New()                           //Sessions persist user data, eg is a user logged in?
	session.Lifetime = 24 * time.Hour              //sets how long the session lasts for.
	session.Cookie.Persist = true                  //sets if the cookie persists past the closure of the browser.
	session.Cookie.SameSite = http.SameSiteLaxMode // sets if a cookie can be sent to other sites with cross site requests
	session.Cookie.Secure = app.InProduction       //sets if the cookie is encrypted or not. In dev mode this is set to false, but in production mode set to true.

	app.Session = session

	//Create the template cache.
	tc, err := render.CreateTemplateCache()
	if err != nil {
		fmt.Println("Can't create template cache ", err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplates(&app) //send the template cache to the render package.

	repo := handlers.NewRepo(&app) //Create a repo using the app config ready to be sent back to the handlers package for use.
	handlers.NewHandlers(repo)     //send the template cache to the handler package.

	//Sets the port to listen on and provides the routing for the handlers.
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	//Listens on port :8080 for any requests and handles the error.
	log.Fatal(srv.ListenAndServe())
	log.Println("Application is running on port: ", portNumber)
}
