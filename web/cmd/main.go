package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/CevdetAkiah/Registration/pkg/handlers"
)





func main() {
	//Handles the URLs
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))

	http.HandleFunc("/", handlers.Index)

	http.HandleFunc("/register", handlers.reg)

	http.HandleFunc("/register-details", handlers.regDets)

	http.HandleFunc("/login", handlers.login)

	http.HandleFunc("/post-login", handlers.postLogin)

	http.HandleFunc("/myaccount", handlers.myAccount)

	//Listens on port :8080 for any requests and handles the error.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
