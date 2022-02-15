package handlers

import "net/http"

func alreadyLoggedIn(req *http.Request) bool {
	//get the cookie
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	//check if the UUID exists in the DB and grab the user name.
	un := Repo.App.DbSessions[c.Value]
	_, ok := Repo.App.DbUsers[un]
	return ok
}
