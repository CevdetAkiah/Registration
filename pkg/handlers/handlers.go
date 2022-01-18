package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

var (
	tpl   *template.Template
	users []User
)

type User struct {
	uName    string
	pWord    []byte
	loggedIn bool
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

//Executes the index html page
func Index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}

//Executes the register html page
func Reg(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "register.html", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "login.html", nil)
}

func MyAccount(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "myaccount.html", nil)
}

func RegDets(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	userName := r.Form["uname"][0]
	password := r.Form["pword"][0]

	epw := Encrypt(password)

	users = append(users, User{userName, epw, false})

	tpl.ExecuteTemplate(w, "regDetails.html", nil)
}

//Simple rot13 encryption. Can only use a-z to create a password for this to work properly.
func Encrypt(pw string) []byte {
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

func PostLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(w, "post-login.html", nil)
}
