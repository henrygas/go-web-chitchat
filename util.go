package main

import (
	"errors"
	"go-web-chitchat/data"
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) (session data.Session, err error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return
	}

	sess := data.Session{
		Uuid: cookie.Value,
	}
	if ok, _ := sess.Check(); !ok {
		err = errors.New("Invalid session")
	}
	return
}
