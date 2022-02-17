package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CevdetAkiah/Registration/pkg/config"
	"github.com/CevdetAkiah/Registration/pkg/models"
	"github.com/CevdetAkiah/Registration/pkg/render"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	Repo              *Repository
	dbUsers                     = make(map[string]models.User)    //user ID, user
	dbSessions                  = make(map[string]models.Session) //session ID, user ID
	sessionLength     int       = 600                             //10 minute session
	dbSessionsCleaned time.Time                                   //records the last session clean up
)

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
	u := getUser(w, r)
	render.RenderTemplate(w, r, "index.page.html", &models.TemplateData{User: u})
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
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}
	var u models.User

	if r.Method == http.MethodPost {
		un := r.FormValue("email")
		p := r.FormValue("pwd")
		//username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		//create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)

		dbSessions[c.Value] = models.Session{Uname: un, LastActivity: time.Now()}

		//store the user in the database
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = models.User{UserName: un, Password: bs}
		dbUsers[un] = u

		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}

}

//LogOut logs the user out and closes any sessions that are over 10 minutes old
func (m *Repository) LogOut(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}

	c, _ := r.Cookie("session")

	//delete session
	delete(dbSessions, c.Value)

	//delete cookie
	c.MaxAge = -1
	http.SetCookie(w, c)

	//clean up dbSessions
	if time.Since(dbSessionsCleaned) > (time.Minute * 10) {
		go cleanSessions()
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
