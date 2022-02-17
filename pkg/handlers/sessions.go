package handlers

import (
	"net/http"
	"time"

	"github.com/CevdetAkiah/Registration/pkg/models"
	uuid "github.com/satori/go.uuid"
)

func alreadyLoggedIn(req *http.Request) bool {
	//get the cookie
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	//check if the UUID exists in the DB and grab the user name.
	session := dbSessions[c.Value]
	_, ok := dbUsers[session.Uname]
	return ok
}

//getUser grabs the user from the session database and sets the session length
func getUser(w http.ResponseWriter, r *http.Request) models.User {
	//get cookie
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	var U models.User
	if session, ok := dbSessions[c.Value]; ok {
		U = dbUsers[session.Uname]
	}
	return U
}

//cleanSessions deletes any sessions older than 10 minutes
func cleanSessions() {
	for k, v := range dbSessions {
		if time.Since(v.LastActivity) > (time.Minute * 10) {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
}
