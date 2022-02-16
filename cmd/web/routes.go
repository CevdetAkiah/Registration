package main

import (
	"net/http"

	"github.com/CevdetAkiah/Registration/pkg/config"
	"github.com/CevdetAkiah/Registration/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	mux.Get("/index", handlers.Repo.Index)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/register", handlers.Repo.Register)
	mux.Post("/register-post", handlers.Repo.PostRegister)
	mux.Get("/logout", handlers.Repo.LogOut)

	//fileServer serves all static files
	fileServer := http.FileServer(http.Dir("../../static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
