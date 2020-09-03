package main

import (
	"go-web-chitchat/data"
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) (session data.Session, err error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return
	}

	session = data.Session{
		Uuid: cookie.Value,
	}
	if ok, _ := session.Check(); !ok {
		err = errSessionInvalid
	}
	return
}
