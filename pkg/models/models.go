package models

import "time"

//User holds the log in details for a user
type User struct {
	UserName string
	Password []byte
}

//Session
type Session struct {
	Uname        string    //username
	LastActivity time.Time //track last cookie use
}
