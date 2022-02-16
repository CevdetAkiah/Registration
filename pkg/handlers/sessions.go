package handlers

import (
	"net/http"

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
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}

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
	http.SetCookie(w, c)
	var U models.User
	if un, ok := dbSessions[c.Value]; ok {
		U = dbUsers[un]
	}
	return U
}
