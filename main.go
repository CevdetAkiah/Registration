package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

type user struct {
	uName    string
	pWord    []byte
	loggedIn bool
}

type users struct {
	people user
}

var (
	tpl        *template.Template
	cybernauts users
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	//Handles the URLs
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/register", reg)
	http.HandleFunc("/register-details", regDets)
	http.HandleFunc("/login", login)
	http.HandleFunc("/post-login", postLogin)
	http.HandleFunc("/myaccount", myAccount)

	//Listens on port :8080 for any requests and handles the error.
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//Executes the index html page
func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

//Executes the register html page
func reg(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "register.html", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "post-login.html", nil)
}

func myAccount(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "myaccount.html", nil)
}

func regDets(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	userName := r.Form["uname"][0]
	password := r.Form["pword"][0]

	epw := encrypt(password)

	newUser := user{userName, epw, false}
	cybernauts.people = newUser
	tpl.ExecuteTemplate(w, "regDetails.html", nil)
}

//Simple rot13 encryption. Can only use a-z to create a password for this to work properly.
func encrypt(pw string) []byte {
	var epw = make([]byte, len(pw))

	bpw := []byte(strings.ToLower(pw))

	for i, v := range bpw {
		//ascii values are 97-122 for lower case a-z.
		if v <= 109 {
			epw[i] = v + 13
		} else if v >= 110 {
			epw[i] = v - 13
		}
	}
	return epw
}
